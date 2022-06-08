package manifests

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	aiv1beta1 "github.com/openshift/assisted-service/api/v1beta1"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/mock"
	"github.com/openshift/installer/pkg/types"
)

func TestInfraEnv_Generate(t *testing.T) {

	t.Skip("Skipping asset generation test")

	installConfig := &installconfig.InstallConfig{
		Config: &types.InstallConfig{
			ObjectMeta: v1.ObjectMeta{
				Name:      "ocp-edge-cluster-0",
				Namespace: "cluster0",
			},
			PullSecret: "secret-agent",
			SSHKey:     "ssh-key",
		},
	}
	agentPullSecret := &AgentPullSecret{}

	parents := asset.Parents{}
	parents.Add(installConfig, agentPullSecret)

	asset := &InfraEnv{}
	err := asset.Generate(parents)
	assert.NoError(t, err)

	assert.NotEmpty(t, asset.Files())
	infraEnvFile := asset.Files()[0]
	assert.Equal(t, "cluster-manifests/infraenv.yml", infraEnvFile.Filename)

	infraEnv := asset.Config
	assert.Equal(t, "infraEnv", infraEnv.Name)
	assert.Equal(t, "cluster0", infraEnv.Namespace)
	assert.Equal(t, "ocp-edge-cluster-0", infraEnv.Spec.ClusterRef.Name)
	assert.Equal(t, "cluster0", infraEnv.Spec.ClusterRef.Namespace)
	assert.Equal(t, "ssh-key", infraEnv.Spec.SSHAuthorizedKey)
	assert.Equal(t, "pull-secret", infraEnv.Spec.PullSecretRef.Name)

}

func TestInfraEnv_LoadedFromDisk(t *testing.T) {

	cases := []struct {
		name           string
		data           string
		fetchError     error
		expectedFound  bool
		expectedError  string
		expectedConfig *aiv1beta1.InfraEnv
	}{
		{
			name: "valid-config-file",
			data: `
metadata:
  name: infraEnv
  namespace: cluster0
spec:
  clusterRef:
    name: ocp-edge-cluster-0
    namespace: cluster0
  nmStateConfigLabelSelector: 
    matchLabels:
      cluster0-nmstate-label-name: cluster0-nmstate-label-value
  pullSecretRef:
    name: pull-secret
  sshAuthorizedKey: |
    ssh-rsa AAAAmyKey`,
			expectedFound: true,
			expectedConfig: &aiv1beta1.InfraEnv{
				ObjectMeta: v1.ObjectMeta{
					Name:      "infraEnv",
					Namespace: "cluster0",
				},
				Spec: aiv1beta1.InfraEnvSpec{
					ClusterRef: &aiv1beta1.ClusterReference{
						Name:      "ocp-edge-cluster-0",
						Namespace: "cluster0",
					},
					NMStateConfigLabelSelector: v1.LabelSelector{
						MatchLabels: map[string]string{
							"cluster0-nmstate-label-name": "cluster0-nmstate-label-value",
						},
					},
					PullSecretRef: &corev1.LocalObjectReference{
						Name: "pull-secret",
					},
					SSHAuthorizedKey: "ssh-rsa AAAAmyKey",
				},
			},
		},
		{
			name:          "not-yaml",
			data:          `This is not a yaml file`,
			expectedError: "failed to unmarshal cluster-manifests/infraenv.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type v1beta1.InfraEnv",
		},
		{
			name:       "file-not-found",
			fetchError: &os.PathError{Err: os.ErrNotExist},
		},
		{
			name:          "error-fetching-file",
			fetchError:    errors.New("fetch failed"),
			expectedError: "failed to load cluster-manifests/infraenv.yaml file: fetch failed",
		},
		{
			name: "unknown-field",
			data: `
		metadata:
		  name: infraEnv
		  namespace: cluster0
		spec:
		  wrongField: wrongValue`,
			expectedError: "failed to unmarshal cluster-manifests/infraenv.yaml: error converting YAML to JSON: yaml: line 2: found character that cannot start any token",
		},
		{
			name: "empty-NMStateLabelSelector",
			data: `
metadata:
  name: infraEnv
  namespace: cluster0
spec:
  clusterRef:
    name: ocp-edge-cluster-0
    namespace: cluster0
  nmStateConfigLabelSelector: 
  pullSecretRef:
    name: pull-secret
  sshAuthorizedKey: |
    ssh-rsa AAAAmyKey`,
			expectedError: "invalid InfraEnv configuration: Spec.NMStateConfigLabelSelector.MatchLabels: Required value: at least one label must be set",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			fileFetcher := mock.NewMockFileFetcher(mockCtrl)
			fileFetcher.EXPECT().FetchByName(infraEnvFilename).
				Return(
					&asset.File{
						Filename: infraEnvFilename,
						Data:     []byte(tc.data)},
					tc.fetchError,
				)

			asset := &InfraEnv{}
			found, err := asset.Load(fileFetcher)
			if tc.expectedError != "" {
				assert.Equal(t, tc.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedFound, found, "unexpected found value returned from Load")
			if tc.expectedFound {
				assert.Equal(t, tc.expectedConfig, asset.Config, "unexpected Config in InfraEnv")
			}
		})
	}

}
