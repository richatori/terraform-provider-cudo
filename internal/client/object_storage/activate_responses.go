// Code generated by go-swagger; DO NOT EDIT.

package object_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// ActivateReader is a Reader for the Activate structure.
type ActivateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActivateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActivateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActivateOK creates a ActivateOK with default headers values
func NewActivateOK() *ActivateOK {
	return &ActivateOK{}
}

/*
ActivateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ActivateOK struct {
	Payload interface{}
}

// IsSuccess returns true when this activate o k response has a 2xx status code
func (o *ActivateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this activate o k response has a 3xx status code
func (o *ActivateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this activate o k response has a 4xx status code
func (o *ActivateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this activate o k response has a 5xx status code
func (o *ActivateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this activate o k response a status code equal to that given
func (o *ActivateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the activate o k response
func (o *ActivateOK) Code() int {
	return 200
}

func (o *ActivateOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/object-storage/activate][%d] activateOK  %+v", 200, o.Payload)
}

func (o *ActivateOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/object-storage/activate][%d] activateOK  %+v", 200, o.Payload)
}

func (o *ActivateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ActivateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateDefault creates a ActivateDefault with default headers values
func NewActivateDefault(code int) *ActivateDefault {
	return &ActivateDefault{
		_statusCode: code,
	}
}

/*
ActivateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ActivateDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this activate default response has a 2xx status code
func (o *ActivateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this activate default response has a 3xx status code
func (o *ActivateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this activate default response has a 4xx status code
func (o *ActivateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this activate default response has a 5xx status code
func (o *ActivateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this activate default response a status code equal to that given
func (o *ActivateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the activate default response
func (o *ActivateDefault) Code() int {
	return o._statusCode
}

func (o *ActivateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/object-storage/activate][%d] Activate default  %+v", o._statusCode, o.Payload)
}

func (o *ActivateDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/object-storage/activate][%d] Activate default  %+v", o._statusCode, o.Payload)
}

func (o *ActivateDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *ActivateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ActivateBody activate body
swagger:model ActivateBody
*/
type ActivateBody struct {

	// data center Id
	DataCenterID string `json:"dataCenterId,omitempty"`
}

// Validate validates this activate body
func (o *ActivateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this activate body based on context it is used
func (o *ActivateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ActivateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ActivateBody) UnmarshalBinary(b []byte) error {
	var res ActivateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
