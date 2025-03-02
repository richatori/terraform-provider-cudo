// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

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

// NewDeletePrivateVMImageParams creates a new DeletePrivateVMImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePrivateVMImageParams() *DeletePrivateVMImageParams {
	return &DeletePrivateVMImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePrivateVMImageParamsWithTimeout creates a new DeletePrivateVMImageParams object
// with the ability to set a timeout on a request.
func NewDeletePrivateVMImageParamsWithTimeout(timeout time.Duration) *DeletePrivateVMImageParams {
	return &DeletePrivateVMImageParams{
		timeout: timeout,
	}
}

// NewDeletePrivateVMImageParamsWithContext creates a new DeletePrivateVMImageParams object
// with the ability to set a context for a request.
func NewDeletePrivateVMImageParamsWithContext(ctx context.Context) *DeletePrivateVMImageParams {
	return &DeletePrivateVMImageParams{
		Context: ctx,
	}
}

// NewDeletePrivateVMImageParamsWithHTTPClient creates a new DeletePrivateVMImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePrivateVMImageParamsWithHTTPClient(client *http.Client) *DeletePrivateVMImageParams {
	return &DeletePrivateVMImageParams{
		HTTPClient: client,
	}
}

/*
DeletePrivateVMImageParams contains all the parameters to send to the API endpoint

	for the delete private VM image operation.

	Typically these are written to a http.Request.
*/
type DeletePrivateVMImageParams struct {

	// ID.
	ID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete private VM image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePrivateVMImageParams) WithDefaults() *DeletePrivateVMImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete private VM image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePrivateVMImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete private VM image params
func (o *DeletePrivateVMImageParams) WithTimeout(timeout time.Duration) *DeletePrivateVMImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete private VM image params
func (o *DeletePrivateVMImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete private VM image params
func (o *DeletePrivateVMImageParams) WithContext(ctx context.Context) *DeletePrivateVMImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete private VM image params
func (o *DeletePrivateVMImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete private VM image params
func (o *DeletePrivateVMImageParams) WithHTTPClient(client *http.Client) *DeletePrivateVMImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete private VM image params
func (o *DeletePrivateVMImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete private VM image params
func (o *DeletePrivateVMImageParams) WithID(id string) *DeletePrivateVMImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete private VM image params
func (o *DeletePrivateVMImageParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete private VM image params
func (o *DeletePrivateVMImageParams) WithProjectID(projectID string) *DeletePrivateVMImageParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete private VM image params
func (o *DeletePrivateVMImageParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePrivateVMImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
