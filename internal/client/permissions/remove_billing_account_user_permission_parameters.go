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

// NewRemoveBillingAccountUserPermissionParams creates a new RemoveBillingAccountUserPermissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveBillingAccountUserPermissionParams() *RemoveBillingAccountUserPermissionParams {
	return &RemoveBillingAccountUserPermissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveBillingAccountUserPermissionParamsWithTimeout creates a new RemoveBillingAccountUserPermissionParams object
// with the ability to set a timeout on a request.
func NewRemoveBillingAccountUserPermissionParamsWithTimeout(timeout time.Duration) *RemoveBillingAccountUserPermissionParams {
	return &RemoveBillingAccountUserPermissionParams{
		timeout: timeout,
	}
}

// NewRemoveBillingAccountUserPermissionParamsWithContext creates a new RemoveBillingAccountUserPermissionParams object
// with the ability to set a context for a request.
func NewRemoveBillingAccountUserPermissionParamsWithContext(ctx context.Context) *RemoveBillingAccountUserPermissionParams {
	return &RemoveBillingAccountUserPermissionParams{
		Context: ctx,
	}
}

// NewRemoveBillingAccountUserPermissionParamsWithHTTPClient creates a new RemoveBillingAccountUserPermissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveBillingAccountUserPermissionParamsWithHTTPClient(client *http.Client) *RemoveBillingAccountUserPermissionParams {
	return &RemoveBillingAccountUserPermissionParams{
		HTTPClient: client,
	}
}

/*
RemoveBillingAccountUserPermissionParams contains all the parameters to send to the API endpoint

	for the remove billing account user permission operation.

	Typically these are written to a http.Request.
*/
type RemoveBillingAccountUserPermissionParams struct {

	// BillingAccountID.
	BillingAccountID string

	// Body.
	Body RemoveBillingAccountUserPermissionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove billing account user permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveBillingAccountUserPermissionParams) WithDefaults() *RemoveBillingAccountUserPermissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove billing account user permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveBillingAccountUserPermissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove billing account user permission params
func (o *RemoveBillingAccountUserPermissionParams) WithTimeout(timeout time.Duration) *RemoveBillingAccountUserPermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove billing account user permission params
func (o *RemoveBillingAccountUserPermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove billing account user permission params
func (o *RemoveBillingAccountUserPermissionParams) WithContext(ctx context.Context) *RemoveBillingAccountUserPermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove billing account user permission params
func (o *RemoveBillingAccountUserPermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove billing account user permission params
func (o *RemoveBillingAccountUserPermissionParams) WithHTTPClient(client *http.Client) *RemoveBillingAccountUserPermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove billing account user permission params
func (o *RemoveBillingAccountUserPermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingAccountID adds the billingAccountID to the remove billing account user permission params
func (o *RemoveBillingAccountUserPermissionParams) WithBillingAccountID(billingAccountID string) *RemoveBillingAccountUserPermissionParams {
	o.SetBillingAccountID(billingAccountID)
	return o
}

// SetBillingAccountID adds the billingAccountId to the remove billing account user permission params
func (o *RemoveBillingAccountUserPermissionParams) SetBillingAccountID(billingAccountID string) {
	o.BillingAccountID = billingAccountID
}

// WithBody adds the body to the remove billing account user permission params
func (o *RemoveBillingAccountUserPermissionParams) WithBody(body RemoveBillingAccountUserPermissionBody) *RemoveBillingAccountUserPermissionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove billing account user permission params
func (o *RemoveBillingAccountUserPermissionParams) SetBody(body RemoveBillingAccountUserPermissionBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveBillingAccountUserPermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
