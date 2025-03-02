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

// NewListPublicVMImagesParams creates a new ListPublicVMImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPublicVMImagesParams() *ListPublicVMImagesParams {
	return &ListPublicVMImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPublicVMImagesParamsWithTimeout creates a new ListPublicVMImagesParams object
// with the ability to set a timeout on a request.
func NewListPublicVMImagesParamsWithTimeout(timeout time.Duration) *ListPublicVMImagesParams {
	return &ListPublicVMImagesParams{
		timeout: timeout,
	}
}

// NewListPublicVMImagesParamsWithContext creates a new ListPublicVMImagesParams object
// with the ability to set a context for a request.
func NewListPublicVMImagesParamsWithContext(ctx context.Context) *ListPublicVMImagesParams {
	return &ListPublicVMImagesParams{
		Context: ctx,
	}
}

// NewListPublicVMImagesParamsWithHTTPClient creates a new ListPublicVMImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPublicVMImagesParamsWithHTTPClient(client *http.Client) *ListPublicVMImagesParams {
	return &ListPublicVMImagesParams{
		HTTPClient: client,
	}
}

/*
ListPublicVMImagesParams contains all the parameters to send to the API endpoint

	for the list public VM images operation.

	Typically these are written to a http.Request.
*/
type ListPublicVMImagesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list public VM images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPublicVMImagesParams) WithDefaults() *ListPublicVMImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list public VM images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPublicVMImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list public VM images params
func (o *ListPublicVMImagesParams) WithTimeout(timeout time.Duration) *ListPublicVMImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list public VM images params
func (o *ListPublicVMImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list public VM images params
func (o *ListPublicVMImagesParams) WithContext(ctx context.Context) *ListPublicVMImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list public VM images params
func (o *ListPublicVMImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list public VM images params
func (o *ListPublicVMImagesParams) WithHTTPClient(client *http.Client) *ListPublicVMImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list public VM images params
func (o *ListPublicVMImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListPublicVMImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
