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
	"github.com/go-openapi/swag"
)

// NewListPrivateVMImagesParams creates a new ListPrivateVMImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPrivateVMImagesParams() *ListPrivateVMImagesParams {
	return &ListPrivateVMImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPrivateVMImagesParamsWithTimeout creates a new ListPrivateVMImagesParams object
// with the ability to set a timeout on a request.
func NewListPrivateVMImagesParamsWithTimeout(timeout time.Duration) *ListPrivateVMImagesParams {
	return &ListPrivateVMImagesParams{
		timeout: timeout,
	}
}

// NewListPrivateVMImagesParamsWithContext creates a new ListPrivateVMImagesParams object
// with the ability to set a context for a request.
func NewListPrivateVMImagesParamsWithContext(ctx context.Context) *ListPrivateVMImagesParams {
	return &ListPrivateVMImagesParams{
		Context: ctx,
	}
}

// NewListPrivateVMImagesParamsWithHTTPClient creates a new ListPrivateVMImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPrivateVMImagesParamsWithHTTPClient(client *http.Client) *ListPrivateVMImagesParams {
	return &ListPrivateVMImagesParams{
		HTTPClient: client,
	}
}

/*
ListPrivateVMImagesParams contains all the parameters to send to the API endpoint

	for the list private VM images operation.

	Typically these are written to a http.Request.
*/
type ListPrivateVMImagesParams struct {

	// PageNumber.
	//
	// Format: int32
	PageNumber *int32

	// PageSize.
	//
	// Format: int32
	PageSize *int32

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list private VM images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPrivateVMImagesParams) WithDefaults() *ListPrivateVMImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list private VM images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPrivateVMImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list private VM images params
func (o *ListPrivateVMImagesParams) WithTimeout(timeout time.Duration) *ListPrivateVMImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list private VM images params
func (o *ListPrivateVMImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list private VM images params
func (o *ListPrivateVMImagesParams) WithContext(ctx context.Context) *ListPrivateVMImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list private VM images params
func (o *ListPrivateVMImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list private VM images params
func (o *ListPrivateVMImagesParams) WithHTTPClient(client *http.Client) *ListPrivateVMImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list private VM images params
func (o *ListPrivateVMImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the list private VM images params
func (o *ListPrivateVMImagesParams) WithPageNumber(pageNumber *int32) *ListPrivateVMImagesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the list private VM images params
func (o *ListPrivateVMImagesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the list private VM images params
func (o *ListPrivateVMImagesParams) WithPageSize(pageSize *int32) *ListPrivateVMImagesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list private VM images params
func (o *ListPrivateVMImagesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithProjectID adds the projectID to the list private VM images params
func (o *ListPrivateVMImagesParams) WithProjectID(projectID string) *ListPrivateVMImagesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list private VM images params
func (o *ListPrivateVMImagesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListPrivateVMImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
