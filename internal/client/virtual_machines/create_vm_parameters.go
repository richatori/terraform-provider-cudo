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

// NewCreateVMParams creates a new CreateVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVMParams() *CreateVMParams {
	return &CreateVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVMParamsWithTimeout creates a new CreateVMParams object
// with the ability to set a timeout on a request.
func NewCreateVMParamsWithTimeout(timeout time.Duration) *CreateVMParams {
	return &CreateVMParams{
		timeout: timeout,
	}
}

// NewCreateVMParamsWithContext creates a new CreateVMParams object
// with the ability to set a context for a request.
func NewCreateVMParamsWithContext(ctx context.Context) *CreateVMParams {
	return &CreateVMParams{
		Context: ctx,
	}
}

// NewCreateVMParamsWithHTTPClient creates a new CreateVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVMParamsWithHTTPClient(client *http.Client) *CreateVMParams {
	return &CreateVMParams{
		HTTPClient: client,
	}
}

/*
CreateVMParams contains all the parameters to send to the API endpoint

	for the create VM operation.

	Typically these are written to a http.Request.
*/
type CreateVMParams struct {

	// Body.
	Body CreateVMBody

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create VM params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMParams) WithDefaults() *CreateVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create VM params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create VM params
func (o *CreateVMParams) WithTimeout(timeout time.Duration) *CreateVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create VM params
func (o *CreateVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create VM params
func (o *CreateVMParams) WithContext(ctx context.Context) *CreateVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create VM params
func (o *CreateVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create VM params
func (o *CreateVMParams) WithHTTPClient(client *http.Client) *CreateVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create VM params
func (o *CreateVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create VM params
func (o *CreateVMParams) WithBody(body CreateVMBody) *CreateVMParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create VM params
func (o *CreateVMParams) SetBody(body CreateVMBody) {
	o.Body = body
}

// WithProjectID adds the projectID to the create VM params
func (o *CreateVMParams) WithProjectID(projectID string) *CreateVMParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create VM params
func (o *CreateVMParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
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
