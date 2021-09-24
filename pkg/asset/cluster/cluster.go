package cluster

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/cluster/aws"
	"github.com/openshift/installer/pkg/asset/cluster/azure"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/password"
	"github.com/openshift/installer/pkg/asset/quota"
	"github.com/openshift/installer/pkg/metrics/timer"
	"github.com/openshift/installer/pkg/terraform"
	"github.com/openshift/installer/pkg/terraform/exec"
	platformstages "github.com/openshift/installer/pkg/terraform/stages/platform"
	typesaws "github.com/openshift/installer/pkg/types/aws"
	typesazure "github.com/openshift/installer/pkg/types/azure"
)

// Cluster uses the terraform executable to launch a cluster
// with the given terraform tfvar and generated templates.
type Cluster struct {
	FileList []*asset.File
}

var _ asset.WritableAsset = (*Cluster)(nil)

// Name returns the human-friendly name of the asset.
func (c *Cluster) Name() string {
	return "Cluster"
}

// Dependencies returns the direct dependency for launching
// the cluster.
func (c *Cluster) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.ClusterID{},
		&installconfig.InstallConfig{},
		// PlatformCredsCheck, PlatformPermsCheck and PlatformProvisionCheck
		// perform validations & check perms required to provision infrastructure.
		// We do not actually use them in this asset directly, hence
		// they are put in the dependencies but not fetched in Generate.
		&installconfig.PlatformCredsCheck{},
		&installconfig.PlatformPermsCheck{},
		&installconfig.PlatformProvisionCheck{},
		&quota.PlatformQuotaCheck{},
		&TerraformVariables{},
		&password.KubeadminPassword{},
	}
}

// Generate launches the cluster and generates the terraform state file on disk.
func (c *Cluster) Generate(parents asset.Parents) (err error) {
	clusterID := &installconfig.ClusterID{}
	installConfig := &installconfig.InstallConfig{}
	terraformVariables := &TerraformVariables{}
	parents.Get(clusterID, installConfig, terraformVariables)

	if installConfig.Config.Platform.None != nil {
		return errors.New("cluster cannot be created with platform set to 'none'")
	}

	if installConfig.Config.BootstrapInPlace != nil {
		return errors.New("cluster cannot be created with bootstrapInPlace set")
	}

	platform := installConfig.Config.Platform.Name()
	stages := platformstages.StagesForPlatform(platform)

	logrus.Infof("Creating infrastructure resources...")
	switch platform {
	case typesaws.Name:
		if err := aws.PreTerraform(context.TODO(), clusterID.InfraID, installConfig); err != nil {
			return err
		}
	case typesazure.Name:
		if err := azure.PreTerraform(context.TODO(), clusterID.InfraID, installConfig); err != nil {
			return err
		}
		if installConfig.Config.Platform.Azure.CloudName == typesazure.StackCloud {
			platform = "azurestack"
		}
	}

	tfvarsFiles := make([]*asset.File, 0, len(terraformVariables.Files())+len(stages))
	for _, file := range terraformVariables.Files() {
		tfvarsFiles = append(tfvarsFiles, file)
	}

	for _, stage := range stages {

		// Copy the terraform.tfvars to a temp directory where the terraform
		// will be invoked within.
		tmpDir, err := ioutil.TempDir("", fmt.Sprintf("openshift-install-%s-", stage.Name()))
		if err != nil {
			return errors.Wrap(err, "failed to create temp dir for terraform execution")
		}
		defer os.RemoveAll(tmpDir)

		var extraArgs []string
		for _, file := range tfvarsFiles {
			if err := ioutil.WriteFile(filepath.Join(tmpDir, file.Filename), file.Data, 0600); err != nil {
				return err
			}
			extraArgs = append(extraArgs, fmt.Sprintf("-var-file=%s", filepath.Join(tmpDir, file.Filename)))
		}

		outputs, err := c.applyTerraform(tmpDir, platform, stage, extraArgs)
		if err != nil {
			return err
		}

		tfvarsFiles = append(tfvarsFiles, outputs)
	}

	return nil
}

// Files returns the FileList generated by the asset.
func (c *Cluster) Files() []*asset.File {
	return c.FileList
}

// Load returns error if the tfstate file is already on-disk, because we want to
// prevent user from accidentally re-launching the cluster.
func (c *Cluster) Load(f asset.FileFetcher) (found bool, err error) {
	matches, err := filepath.Glob("terraform(.*)?.tfstate")
	if err != nil {
		return true, err
	}
	if len(matches) != 0 {
		return true, errors.Errorf("terraform state files alread exist.  There may already be a running cluster")
	}

	return false, nil
}

func (c *Cluster) applyTerraform(tmpDir string, platform string, stage terraform.Stage, extraArgs []string) (*asset.File, error) {
	timer.StartTimer(stage.Name())
	defer timer.StopTimer(stage.Name())

	stateFile, err := terraform.Apply(tmpDir, platform, stage, extraArgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to create cluster")
		if stateFile == "" {
			return nil, err
		}
		// Store the error from the apply, but continue with the
		// generation so that the Terraform state file is recovered from
		// the temporary directory.
	}

	data, err2 := ioutil.ReadFile(stateFile)
	if err2 == nil {
		c.FileList = append(c.FileList, &asset.File{
			Filename: stage.StateFilename(),
			Data:     data,
		})
	} else {
		logrus.Errorf("Failed to read tfstate: %v", err2)
		if err == nil {
			err2 = err
		}
	}
	if err != nil {
		return nil, err
	}

	outputs, err := exec.Outputs(stateFile)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get outputs from state file %q", stateFile)
	}

	outputsFile := &asset.File{
		Filename: stage.OutputsFilename(),
		Data:     outputs,
	}
	c.FileList = append(c.FileList, outputsFile)
	return outputsFile, nil
}
