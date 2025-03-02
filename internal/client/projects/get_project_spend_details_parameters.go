// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewGetProjectSpendDetailsParams creates a new GetProjectSpendDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectSpendDetailsParams() *GetProjectSpendDetailsParams {
	return &GetProjectSpendDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectSpendDetailsParamsWithTimeout creates a new GetProjectSpendDetailsParams object
// with the ability to set a timeout on a request.
func NewGetProjectSpendDetailsParamsWithTimeout(timeout time.Duration) *GetProjectSpendDetailsParams {
	return &GetProjectSpendDetailsParams{
		timeout: timeout,
	}
}

// NewGetProjectSpendDetailsParamsWithContext creates a new GetProjectSpendDetailsParams object
// with the ability to set a context for a request.
func NewGetProjectSpendDetailsParamsWithContext(ctx context.Context) *GetProjectSpendDetailsParams {
	return &GetProjectSpendDetailsParams{
		Context: ctx,
	}
}

// NewGetProjectSpendDetailsParamsWithHTTPClient creates a new GetProjectSpendDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectSpendDetailsParamsWithHTTPClient(client *http.Client) *GetProjectSpendDetailsParams {
	return &GetProjectSpendDetailsParams{
		HTTPClient: client,
	}
}

/*
GetProjectSpendDetailsParams contains all the parameters to send to the API endpoint

	for the get project spend details operation.

	Typically these are written to a http.Request.
*/
type GetProjectSpendDetailsParams struct {

	// ProjectID.
	ProjectID string

	// SpendID.
	SpendID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project spend details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectSpendDetailsParams) WithDefaults() *GetProjectSpendDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project spend details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectSpendDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project spend details params
func (o *GetProjectSpendDetailsParams) WithTimeout(timeout time.Duration) *GetProjectSpendDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project spend details params
func (o *GetProjectSpendDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project spend details params
func (o *GetProjectSpendDetailsParams) WithContext(ctx context.Context) *GetProjectSpendDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project spend details params
func (o *GetProjectSpendDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project spend details params
func (o *GetProjectSpendDetailsParams) WithHTTPClient(client *http.Client) *GetProjectSpendDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project spend details params
func (o *GetProjectSpendDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get project spend details params
func (o *GetProjectSpendDetailsParams) WithProjectID(projectID string) *GetProjectSpendDetailsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project spend details params
func (o *GetProjectSpendDetailsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithSpendID adds the spendID to the get project spend details params
func (o *GetProjectSpendDetailsParams) WithSpendID(spendID string) *GetProjectSpendDetailsParams {
	o.SetSpendID(spendID)
	return o
}

// SetSpendID adds the spendId to the get project spend details params
func (o *GetProjectSpendDetailsParams) SetSpendID(spendID string) {
	o.SpendID = spendID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectSpendDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	// path param spendId
	if err := r.SetPathParam("spendId", o.SpendID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
