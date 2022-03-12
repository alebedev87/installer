// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudPlacementgroupsGetallParams creates a new PcloudPlacementgroupsGetallParams object
// with the default values initialized.
func NewPcloudPlacementgroupsGetallParams() *PcloudPlacementgroupsGetallParams {
	var ()
	return &PcloudPlacementgroupsGetallParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPlacementgroupsGetallParamsWithTimeout creates a new PcloudPlacementgroupsGetallParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudPlacementgroupsGetallParamsWithTimeout(timeout time.Duration) *PcloudPlacementgroupsGetallParams {
	var ()
	return &PcloudPlacementgroupsGetallParams{

		timeout: timeout,
	}
}

// NewPcloudPlacementgroupsGetallParamsWithContext creates a new PcloudPlacementgroupsGetallParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudPlacementgroupsGetallParamsWithContext(ctx context.Context) *PcloudPlacementgroupsGetallParams {
	var ()
	return &PcloudPlacementgroupsGetallParams{

		Context: ctx,
	}
}

// NewPcloudPlacementgroupsGetallParamsWithHTTPClient creates a new PcloudPlacementgroupsGetallParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudPlacementgroupsGetallParamsWithHTTPClient(client *http.Client) *PcloudPlacementgroupsGetallParams {
	var ()
	return &PcloudPlacementgroupsGetallParams{
		HTTPClient: client,
	}
}

/*PcloudPlacementgroupsGetallParams contains all the parameters to send to the API endpoint
for the pcloud placementgroups getall operation typically these are written to a http.Request
*/
type PcloudPlacementgroupsGetallParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud placementgroups getall params
func (o *PcloudPlacementgroupsGetallParams) WithTimeout(timeout time.Duration) *PcloudPlacementgroupsGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud placementgroups getall params
func (o *PcloudPlacementgroupsGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud placementgroups getall params
func (o *PcloudPlacementgroupsGetallParams) WithContext(ctx context.Context) *PcloudPlacementgroupsGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud placementgroups getall params
func (o *PcloudPlacementgroupsGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud placementgroups getall params
func (o *PcloudPlacementgroupsGetallParams) WithHTTPClient(client *http.Client) *PcloudPlacementgroupsGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud placementgroups getall params
func (o *PcloudPlacementgroupsGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud placementgroups getall params
func (o *PcloudPlacementgroupsGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPlacementgroupsGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud placementgroups getall params
func (o *PcloudPlacementgroupsGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPlacementgroupsGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}