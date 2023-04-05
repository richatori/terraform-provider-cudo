// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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
	"github.com/go-openapi/swag"
)

// NewListAPIKeysParams creates a new ListAPIKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAPIKeysParams() *ListAPIKeysParams {
	return &ListAPIKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAPIKeysParamsWithTimeout creates a new ListAPIKeysParams object
// with the ability to set a timeout on a request.
func NewListAPIKeysParamsWithTimeout(timeout time.Duration) *ListAPIKeysParams {
	return &ListAPIKeysParams{
		timeout: timeout,
	}
}

// NewListAPIKeysParamsWithContext creates a new ListAPIKeysParams object
// with the ability to set a context for a request.
func NewListAPIKeysParamsWithContext(ctx context.Context) *ListAPIKeysParams {
	return &ListAPIKeysParams{
		Context: ctx,
	}
}

// NewListAPIKeysParamsWithHTTPClient creates a new ListAPIKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAPIKeysParamsWithHTTPClient(client *http.Client) *ListAPIKeysParams {
	return &ListAPIKeysParams{
		HTTPClient: client,
	}
}

/*
ListAPIKeysParams contains all the parameters to send to the API endpoint

	for the list Api keys operation.

	Typically these are written to a http.Request.
*/
type ListAPIKeysParams struct {

	// PageNumber.
	//
	// Format: int32
	PageNumber *int32

	// PageSize.
	//
	// Format: int32
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list Api keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAPIKeysParams) WithDefaults() *ListAPIKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list Api keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAPIKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list Api keys params
func (o *ListAPIKeysParams) WithTimeout(timeout time.Duration) *ListAPIKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list Api keys params
func (o *ListAPIKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list Api keys params
func (o *ListAPIKeysParams) WithContext(ctx context.Context) *ListAPIKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list Api keys params
func (o *ListAPIKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list Api keys params
func (o *ListAPIKeysParams) WithHTTPClient(client *http.Client) *ListAPIKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list Api keys params
func (o *ListAPIKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the list Api keys params
func (o *ListAPIKeysParams) WithPageNumber(pageNumber *int32) *ListAPIKeysParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the list Api keys params
func (o *ListAPIKeysParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the list Api keys params
func (o *ListAPIKeysParams) WithPageSize(pageSize *int32) *ListAPIKeysParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list Api keys params
func (o *ListAPIKeysParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *ListAPIKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
