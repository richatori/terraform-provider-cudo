// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// UpdateProjectReader is a Reader for the UpdateProject structure.
type UpdateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProjectOK creates a UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {
	return &UpdateProjectOK{}
}

/*
UpdateProjectOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateProjectOK struct {
	Payload *models.Project
}

// IsSuccess returns true when this update project o k response has a 2xx status code
func (o *UpdateProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project o k response has a 3xx status code
func (o *UpdateProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project o k response has a 4xx status code
func (o *UpdateProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project o k response has a 5xx status code
func (o *UpdateProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update project o k response a status code equal to that given
func (o *UpdateProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update project o k response
func (o *UpdateProjectOK) Code() int {
	return 200
}

func (o *UpdateProjectOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/projects/{project.id}][%d] updateProjectOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectOK) String() string {
	return fmt.Sprintf("[PATCH /v1/projects/{project.id}][%d] updateProjectOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *UpdateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectDefault creates a UpdateProjectDefault with default headers values
func NewUpdateProjectDefault(code int) *UpdateProjectDefault {
	return &UpdateProjectDefault{
		_statusCode: code,
	}
}

/*
UpdateProjectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateProjectDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this update project default response has a 2xx status code
func (o *UpdateProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update project default response has a 3xx status code
func (o *UpdateProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update project default response has a 4xx status code
func (o *UpdateProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update project default response has a 5xx status code
func (o *UpdateProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update project default response a status code equal to that given
func (o *UpdateProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update project default response
func (o *UpdateProjectDefault) Code() int {
	return o._statusCode
}

func (o *UpdateProjectDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/projects/{project.id}][%d] UpdateProject default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProjectDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/projects/{project.id}][%d] UpdateProject default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProjectDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *UpdateProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateProjectBody update project body
swagger:model UpdateProjectBody
*/
type UpdateProjectBody struct {

	// billing account Id
	// Required: true
	BillingAccountID *string `json:"billingAccountId"`

	// create by
	// Read Only: true
	CreateBy string `json:"createBy,omitempty"`

	// resource count
	// Read Only: true
	ResourceCount int32 `json:"resourceCount,omitempty"`
}

// Validate validates this update project body
func (o *UpdateProjectBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBillingAccountID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateProjectBody) validateBillingAccountID(formats strfmt.Registry) error {

	if err := validate.Required("project"+"."+"billingAccountId", "body", o.BillingAccountID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update project body based on the context it is used
func (o *UpdateProjectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCreateBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateResourceCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateProjectBody) contextValidateCreateBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "project"+"."+"createBy", "body", string(o.CreateBy)); err != nil {
		return err
	}

	return nil
}

func (o *UpdateProjectBody) contextValidateResourceCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "project"+"."+"resourceCount", "body", int32(o.ResourceCount)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateProjectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateProjectBody) UnmarshalBinary(b []byte) error {
	var res UpdateProjectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
