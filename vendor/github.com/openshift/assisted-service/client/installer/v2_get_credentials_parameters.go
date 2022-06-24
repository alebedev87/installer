// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewV2GetCredentialsParams creates a new V2GetCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2GetCredentialsParams() *V2GetCredentialsParams {
	return &V2GetCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2GetCredentialsParamsWithTimeout creates a new V2GetCredentialsParams object
// with the ability to set a timeout on a request.
func NewV2GetCredentialsParamsWithTimeout(timeout time.Duration) *V2GetCredentialsParams {
	return &V2GetCredentialsParams{
		timeout: timeout,
	}
}

// NewV2GetCredentialsParamsWithContext creates a new V2GetCredentialsParams object
// with the ability to set a context for a request.
func NewV2GetCredentialsParamsWithContext(ctx context.Context) *V2GetCredentialsParams {
	return &V2GetCredentialsParams{
		Context: ctx,
	}
}

// NewV2GetCredentialsParamsWithHTTPClient creates a new V2GetCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2GetCredentialsParamsWithHTTPClient(client *http.Client) *V2GetCredentialsParams {
	return &V2GetCredentialsParams{
		HTTPClient: client,
	}
}

/* V2GetCredentialsParams contains all the parameters to send to the API endpoint
   for the v2 get credentials operation.

   Typically these are written to a http.Request.
*/
type V2GetCredentialsParams struct {

	/* ClusterID.

	   The cluster whose admin credentials should be retrieved.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 get credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2GetCredentialsParams) WithDefaults() *V2GetCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 get credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2GetCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 get credentials params
func (o *V2GetCredentialsParams) WithTimeout(timeout time.Duration) *V2GetCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 get credentials params
func (o *V2GetCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 get credentials params
func (o *V2GetCredentialsParams) WithContext(ctx context.Context) *V2GetCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 get credentials params
func (o *V2GetCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 get credentials params
func (o *V2GetCredentialsParams) WithHTTPClient(client *http.Client) *V2GetCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 get credentials params
func (o *V2GetCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the v2 get credentials params
func (o *V2GetCredentialsParams) WithClusterID(clusterID strfmt.UUID) *V2GetCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the v2 get credentials params
func (o *V2GetCredentialsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *V2GetCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
