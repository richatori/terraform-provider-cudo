// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewGetNetworkParams creates a new GetNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkParams() *GetNetworkParams {
	return &GetNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkParamsWithTimeout creates a new GetNetworkParams object
// with the ability to set a timeout on a request.
func NewGetNetworkParamsWithTimeout(timeout time.Duration) *GetNetworkParams {
	return &GetNetworkParams{
		timeout: timeout,
	}
}

// NewGetNetworkParamsWithContext creates a new GetNetworkParams object
// with the ability to set a context for a request.
func NewGetNetworkParamsWithContext(ctx context.Context) *GetNetworkParams {
	return &GetNetworkParams{
		Context: ctx,
	}
}

// NewGetNetworkParamsWithHTTPClient creates a new GetNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkParamsWithHTTPClient(client *http.Client) *GetNetworkParams {
	return &GetNetworkParams{
		HTTPClient: client,
	}
}

/*
GetNetworkParams contains all the parameters to send to the API endpoint

	for the get network operation.

	Typically these are written to a http.Request.
*/
type GetNetworkParams struct {

	// ID.
	ID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkParams) WithDefaults() *GetNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network params
func (o *GetNetworkParams) WithTimeout(timeout time.Duration) *GetNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network params
func (o *GetNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network params
func (o *GetNetworkParams) WithContext(ctx context.Context) *GetNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network params
func (o *GetNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network params
func (o *GetNetworkParams) WithHTTPClient(client *http.Client) *GetNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network params
func (o *GetNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get network params
func (o *GetNetworkParams) WithID(id string) *GetNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network params
func (o *GetNetworkParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the get network params
func (o *GetNetworkParams) WithProjectID(projectID string) *GetNetworkParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get network params
func (o *GetNetworkParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
