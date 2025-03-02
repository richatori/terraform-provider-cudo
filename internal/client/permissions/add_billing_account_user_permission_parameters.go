// Code generated by go-swagger; DO NOT EDIT.

package permissions

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

// NewAddBillingAccountUserPermissionParams creates a new AddBillingAccountUserPermissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddBillingAccountUserPermissionParams() *AddBillingAccountUserPermissionParams {
	return &AddBillingAccountUserPermissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddBillingAccountUserPermissionParamsWithTimeout creates a new AddBillingAccountUserPermissionParams object
// with the ability to set a timeout on a request.
func NewAddBillingAccountUserPermissionParamsWithTimeout(timeout time.Duration) *AddBillingAccountUserPermissionParams {
	return &AddBillingAccountUserPermissionParams{
		timeout: timeout,
	}
}

// NewAddBillingAccountUserPermissionParamsWithContext creates a new AddBillingAccountUserPermissionParams object
// with the ability to set a context for a request.
func NewAddBillingAccountUserPermissionParamsWithContext(ctx context.Context) *AddBillingAccountUserPermissionParams {
	return &AddBillingAccountUserPermissionParams{
		Context: ctx,
	}
}

// NewAddBillingAccountUserPermissionParamsWithHTTPClient creates a new AddBillingAccountUserPermissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddBillingAccountUserPermissionParamsWithHTTPClient(client *http.Client) *AddBillingAccountUserPermissionParams {
	return &AddBillingAccountUserPermissionParams{
		HTTPClient: client,
	}
}

/*
AddBillingAccountUserPermissionParams contains all the parameters to send to the API endpoint

	for the add billing account user permission operation.

	Typically these are written to a http.Request.
*/
type AddBillingAccountUserPermissionParams struct {

	// BillingAccountID.
	BillingAccountID string

	// Body.
	Body AddBillingAccountUserPermissionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add billing account user permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBillingAccountUserPermissionParams) WithDefaults() *AddBillingAccountUserPermissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add billing account user permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBillingAccountUserPermissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add billing account user permission params
func (o *AddBillingAccountUserPermissionParams) WithTimeout(timeout time.Duration) *AddBillingAccountUserPermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add billing account user permission params
func (o *AddBillingAccountUserPermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add billing account user permission params
func (o *AddBillingAccountUserPermissionParams) WithContext(ctx context.Context) *AddBillingAccountUserPermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add billing account user permission params
func (o *AddBillingAccountUserPermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add billing account user permission params
func (o *AddBillingAccountUserPermissionParams) WithHTTPClient(client *http.Client) *AddBillingAccountUserPermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add billing account user permission params
func (o *AddBillingAccountUserPermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingAccountID adds the billingAccountID to the add billing account user permission params
func (o *AddBillingAccountUserPermissionParams) WithBillingAccountID(billingAccountID string) *AddBillingAccountUserPermissionParams {
	o.SetBillingAccountID(billingAccountID)
	return o
}

// SetBillingAccountID adds the billingAccountId to the add billing account user permission params
func (o *AddBillingAccountUserPermissionParams) SetBillingAccountID(billingAccountID string) {
	o.BillingAccountID = billingAccountID
}

// WithBody adds the body to the add billing account user permission params
func (o *AddBillingAccountUserPermissionParams) WithBody(body AddBillingAccountUserPermissionBody) *AddBillingAccountUserPermissionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add billing account user permission params
func (o *AddBillingAccountUserPermissionParams) SetBody(body AddBillingAccountUserPermissionBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddBillingAccountUserPermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param billingAccountId
	if err := r.SetPathParam("billingAccountId", o.BillingAccountID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
