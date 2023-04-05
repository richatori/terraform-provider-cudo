// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cudo.org/v1/terraform-provider-cudo/internal/models"
)

// CreateProjectReader is a Reader for the CreateProject structure.
type CreateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProjectOK creates a CreateProjectOK with default headers values
func NewCreateProjectOK() *CreateProjectOK {
	return &CreateProjectOK{}
}

/*
	CreateProjectOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateProjectOK struct {
	Payload *models.Project
}

// IsSuccess returns true when this create project o k response has a 2xx status code
func (o *CreateProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create project o k response has a 3xx status code
func (o *CreateProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project o k response has a 4xx status code
func (o *CreateProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create project o k response has a 5xx status code
func (o *CreateProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create project o k response a status code equal to that given
func (o *CreateProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateProjectOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createProjectOK  %+v", 200, o.Payload)
}

func (o *CreateProjectOK) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createProjectOK  %+v", 200, o.Payload)
}

func (o *CreateProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *CreateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectDefault creates a CreateProjectDefault with default headers values
func NewCreateProjectDefault(code int) *CreateProjectDefault {
	return &CreateProjectDefault{
		_statusCode: code,
	}
}

/*
	CreateProjectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateProjectDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the create project default response
func (o *CreateProjectDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create project default response has a 2xx status code
func (o *CreateProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create project default response has a 3xx status code
func (o *CreateProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create project default response has a 4xx status code
func (o *CreateProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create project default response has a 5xx status code
func (o *CreateProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create project default response a status code equal to that given
func (o *CreateProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateProjectDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] CreateProject default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProjectDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] CreateProject default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProjectDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *CreateProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
