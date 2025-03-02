// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// UpdateSecurityGroupReader is a Reader for the UpdateSecurityGroup structure.
type UpdateSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateSecurityGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSecurityGroupOK creates a UpdateSecurityGroupOK with default headers values
func NewUpdateSecurityGroupOK() *UpdateSecurityGroupOK {
	return &UpdateSecurityGroupOK{}
}

/*
UpdateSecurityGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateSecurityGroupOK struct {
	Payload *models.UpdateSecurityGroupResponse
}

// IsSuccess returns true when this update security group o k response has a 2xx status code
func (o *UpdateSecurityGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update security group o k response has a 3xx status code
func (o *UpdateSecurityGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update security group o k response has a 4xx status code
func (o *UpdateSecurityGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update security group o k response has a 5xx status code
func (o *UpdateSecurityGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update security group o k response a status code equal to that given
func (o *UpdateSecurityGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update security group o k response
func (o *UpdateSecurityGroupOK) Code() int {
	return 200
}

func (o *UpdateSecurityGroupOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/projects/{securityGroup.projectId}/networks/security-groups/{securityGroup.id}][%d] updateSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateSecurityGroupOK) String() string {
	return fmt.Sprintf("[PATCH /v1/projects/{securityGroup.projectId}/networks/security-groups/{securityGroup.id}][%d] updateSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateSecurityGroupOK) GetPayload() *models.UpdateSecurityGroupResponse {
	return o.Payload
}

func (o *UpdateSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateSecurityGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityGroupDefault creates a UpdateSecurityGroupDefault with default headers values
func NewUpdateSecurityGroupDefault(code int) *UpdateSecurityGroupDefault {
	return &UpdateSecurityGroupDefault{
		_statusCode: code,
	}
}

/*
UpdateSecurityGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateSecurityGroupDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this update security group default response has a 2xx status code
func (o *UpdateSecurityGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update security group default response has a 3xx status code
func (o *UpdateSecurityGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update security group default response has a 4xx status code
func (o *UpdateSecurityGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update security group default response has a 5xx status code
func (o *UpdateSecurityGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update security group default response a status code equal to that given
func (o *UpdateSecurityGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update security group default response
func (o *UpdateSecurityGroupDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSecurityGroupDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/projects/{securityGroup.projectId}/networks/security-groups/{securityGroup.id}][%d] UpdateSecurityGroup default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSecurityGroupDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/projects/{securityGroup.projectId}/networks/security-groups/{securityGroup.id}][%d] UpdateSecurityGroup default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSecurityGroupDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *UpdateSecurityGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateSecurityGroupBody update security group body
swagger:model UpdateSecurityGroupBody
*/
type UpdateSecurityGroupBody struct {

	// data center Id
	// Required: true
	DataCenterID *string `json:"dataCenterId"`

	// description
	Description string `json:"description,omitempty"`

	// rules
	Rules []*models.Rule `json:"rules"`
}

// Validate validates this update security group body
func (o *UpdateSecurityGroupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDataCenterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSecurityGroupBody) validateDataCenterID(formats strfmt.Registry) error {

	if err := validate.Required("securityGroup"+"."+"dataCenterId", "body", o.DataCenterID); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSecurityGroupBody) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(o.Rules) { // not required
		return nil
	}

	for i := 0; i < len(o.Rules); i++ {
		if swag.IsZero(o.Rules[i]) { // not required
			continue
		}

		if o.Rules[i] != nil {
			if err := o.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityGroup" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("securityGroup" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update security group body based on the context it is used
func (o *UpdateSecurityGroupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSecurityGroupBody) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {

			if swag.IsZero(o.Rules[i]) { // not required
				return nil
			}

			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityGroup" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("securityGroup" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSecurityGroupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSecurityGroupBody) UnmarshalBinary(b []byte) error {
	var res UpdateSecurityGroupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
