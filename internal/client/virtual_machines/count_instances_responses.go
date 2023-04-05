// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cudo.org/v1/terraform-provider-cudo/internal/models"
)

// CountInstancesReader is a Reader for the CountInstances structure.
type CountInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCountInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCountInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCountInstancesOK creates a CountInstancesOK with default headers values
func NewCountInstancesOK() *CountInstancesOK {
	return &CountInstancesOK{}
}

/*
	CountInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type CountInstancesOK struct {
	Payload *models.CountInstancesResponse
}

// IsSuccess returns true when this count instances o k response has a 2xx status code
func (o *CountInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this count instances o k response has a 3xx status code
func (o *CountInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count instances o k response has a 4xx status code
func (o *CountInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this count instances o k response has a 5xx status code
func (o *CountInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this count instances o k response a status code equal to that given
func (o *CountInstancesOK) IsCode(code int) bool {
	return code == 200
}

func (o *CountInstancesOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/instance-count][%d] countInstancesOK  %+v", 200, o.Payload)
}

func (o *CountInstancesOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/instance-count][%d] countInstancesOK  %+v", 200, o.Payload)
}

func (o *CountInstancesOK) GetPayload() *models.CountInstancesResponse {
	return o.Payload
}

func (o *CountInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CountInstancesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCountInstancesDefault creates a CountInstancesDefault with default headers values
func NewCountInstancesDefault(code int) *CountInstancesDefault {
	return &CountInstancesDefault{
		_statusCode: code,
	}
}

/*
	CountInstancesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CountInstancesDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the count instances default response
func (o *CountInstancesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this count instances default response has a 2xx status code
func (o *CountInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this count instances default response has a 3xx status code
func (o *CountInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this count instances default response has a 4xx status code
func (o *CountInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this count instances default response has a 5xx status code
func (o *CountInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this count instances default response a status code equal to that given
func (o *CountInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CountInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/instance-count][%d] CountInstances default  %+v", o._statusCode, o.Payload)
}

func (o *CountInstancesDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/instance-count][%d] CountInstances default  %+v", o._statusCode, o.Payload)
}

func (o *CountInstancesDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *CountInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
