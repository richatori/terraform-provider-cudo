// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// ResizeVMDiskReader is a Reader for the ResizeVMDisk structure.
type ResizeVMDiskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResizeVMDiskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResizeVMDiskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResizeVMDiskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResizeVMDiskOK creates a ResizeVMDiskOK with default headers values
func NewResizeVMDiskOK() *ResizeVMDiskOK {
	return &ResizeVMDiskOK{}
}

/*
ResizeVMDiskOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResizeVMDiskOK struct {
	Payload models.ResizeVMDiskResponse
}

// IsSuccess returns true when this resize Vm disk o k response has a 2xx status code
func (o *ResizeVMDiskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resize Vm disk o k response has a 3xx status code
func (o *ResizeVMDiskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resize Vm disk o k response has a 4xx status code
func (o *ResizeVMDiskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resize Vm disk o k response has a 5xx status code
func (o *ResizeVMDiskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resize Vm disk o k response a status code equal to that given
func (o *ResizeVMDiskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resize Vm disk o k response
func (o *ResizeVMDiskOK) Code() int {
	return 200
}

func (o *ResizeVMDiskOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/projects/{projectId}/vms/{id}/disks][%d] resizeVmDiskOK  %+v", 200, o.Payload)
}

func (o *ResizeVMDiskOK) String() string {
	return fmt.Sprintf("[PATCH /v1/projects/{projectId}/vms/{id}/disks][%d] resizeVmDiskOK  %+v", 200, o.Payload)
}

func (o *ResizeVMDiskOK) GetPayload() models.ResizeVMDiskResponse {
	return o.Payload
}

func (o *ResizeVMDiskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResizeVMDiskDefault creates a ResizeVMDiskDefault with default headers values
func NewResizeVMDiskDefault(code int) *ResizeVMDiskDefault {
	return &ResizeVMDiskDefault{
		_statusCode: code,
	}
}

/*
ResizeVMDiskDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResizeVMDiskDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this resize VM disk default response has a 2xx status code
func (o *ResizeVMDiskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resize VM disk default response has a 3xx status code
func (o *ResizeVMDiskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resize VM disk default response has a 4xx status code
func (o *ResizeVMDiskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resize VM disk default response has a 5xx status code
func (o *ResizeVMDiskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resize VM disk default response a status code equal to that given
func (o *ResizeVMDiskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resize VM disk default response
func (o *ResizeVMDiskDefault) Code() int {
	return o._statusCode
}

func (o *ResizeVMDiskDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/projects/{projectId}/vms/{id}/disks][%d] ResizeVMDisk default  %+v", o._statusCode, o.Payload)
}

func (o *ResizeVMDiskDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/projects/{projectId}/vms/{id}/disks][%d] ResizeVMDisk default  %+v", o._statusCode, o.Payload)
}

func (o *ResizeVMDiskDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *ResizeVMDiskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
