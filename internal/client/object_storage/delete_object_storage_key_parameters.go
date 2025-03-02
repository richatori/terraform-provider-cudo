// Code generated by go-swagger; DO NOT EDIT.

package object_storage

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

// NewDeleteObjectStorageKeyParams creates a new DeleteObjectStorageKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteObjectStorageKeyParams() *DeleteObjectStorageKeyParams {
	return &DeleteObjectStorageKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteObjectStorageKeyParamsWithTimeout creates a new DeleteObjectStorageKeyParams object
// with the ability to set a timeout on a request.
func NewDeleteObjectStorageKeyParamsWithTimeout(timeout time.Duration) *DeleteObjectStorageKeyParams {
	return &DeleteObjectStorageKeyParams{
		timeout: timeout,
	}
}

// NewDeleteObjectStorageKeyParamsWithContext creates a new DeleteObjectStorageKeyParams object
// with the ability to set a context for a request.
func NewDeleteObjectStorageKeyParamsWithContext(ctx context.Context) *DeleteObjectStorageKeyParams {
	return &DeleteObjectStorageKeyParams{
		Context: ctx,
	}
}

// NewDeleteObjectStorageKeyParamsWithHTTPClient creates a new DeleteObjectStorageKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteObjectStorageKeyParamsWithHTTPClient(client *http.Client) *DeleteObjectStorageKeyParams {
	return &DeleteObjectStorageKeyParams{
		HTTPClient: client,
	}
}

/*
DeleteObjectStorageKeyParams contains all the parameters to send to the API endpoint

	for the delete object storage key operation.

	Typically these are written to a http.Request.
*/
type DeleteObjectStorageKeyParams struct {

	// AccessKey.
	AccessKey *string

	// ID.
	ID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete object storage key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteObjectStorageKeyParams) WithDefaults() *DeleteObjectStorageKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete object storage key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteObjectStorageKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) WithTimeout(timeout time.Duration) *DeleteObjectStorageKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) WithContext(ctx context.Context) *DeleteObjectStorageKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) WithHTTPClient(client *http.Client) *DeleteObjectStorageKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKey adds the accessKey to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) WithAccessKey(accessKey *string) *DeleteObjectStorageKeyParams {
	o.SetAccessKey(accessKey)
	return o
}

// SetAccessKey adds the accessKey to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) SetAccessKey(accessKey *string) {
	o.AccessKey = accessKey
}

// WithID adds the id to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) WithID(id string) *DeleteObjectStorageKeyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) WithProjectID(projectID string) *DeleteObjectStorageKeyParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete object storage key params
func (o *DeleteObjectStorageKeyParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteObjectStorageKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessKey != nil {

		// query param accessKey
		var qrAccessKey string

		if o.AccessKey != nil {
			qrAccessKey = *o.AccessKey
		}
		qAccessKey := qrAccessKey
		if qAccessKey != "" {

			if err := r.SetQueryParam("accessKey", qAccessKey); err != nil {
				return err
			}
		}
	}

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
