// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

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

// NewGetSSHKeyParams creates a new GetSSHKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSSHKeyParams() *GetSSHKeyParams {
	return &GetSSHKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSSHKeyParamsWithTimeout creates a new GetSSHKeyParams object
// with the ability to set a timeout on a request.
func NewGetSSHKeyParamsWithTimeout(timeout time.Duration) *GetSSHKeyParams {
	return &GetSSHKeyParams{
		timeout: timeout,
	}
}

// NewGetSSHKeyParamsWithContext creates a new GetSSHKeyParams object
// with the ability to set a context for a request.
func NewGetSSHKeyParamsWithContext(ctx context.Context) *GetSSHKeyParams {
	return &GetSSHKeyParams{
		Context: ctx,
	}
}

// NewGetSSHKeyParamsWithHTTPClient creates a new GetSSHKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSSHKeyParamsWithHTTPClient(client *http.Client) *GetSSHKeyParams {
	return &GetSSHKeyParams{
		HTTPClient: client,
	}
}

/*
GetSSHKeyParams contains all the parameters to send to the API endpoint

	for the get Ssh key operation.

	Typically these are written to a http.Request.
*/
type GetSSHKeyParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Ssh key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSSHKeyParams) WithDefaults() *GetSSHKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Ssh key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSSHKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Ssh key params
func (o *GetSSHKeyParams) WithTimeout(timeout time.Duration) *GetSSHKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Ssh key params
func (o *GetSSHKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Ssh key params
func (o *GetSSHKeyParams) WithContext(ctx context.Context) *GetSSHKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Ssh key params
func (o *GetSSHKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Ssh key params
func (o *GetSSHKeyParams) WithHTTPClient(client *http.Client) *GetSSHKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Ssh key params
func (o *GetSSHKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get Ssh key params
func (o *GetSSHKeyParams) WithID(id string) *GetSSHKeyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get Ssh key params
func (o *GetSSHKeyParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSSHKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
