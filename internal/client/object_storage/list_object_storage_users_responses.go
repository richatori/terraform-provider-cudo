// Code generated by go-swagger; DO NOT EDIT.

package object_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
)

// ListObjectStorageUsersReader is a Reader for the ListObjectStorageUsers structure.
type ListObjectStorageUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListObjectStorageUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListObjectStorageUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListObjectStorageUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListObjectStorageUsersOK creates a ListObjectStorageUsersOK with default headers values
func NewListObjectStorageUsersOK() *ListObjectStorageUsersOK {
	return &ListObjectStorageUsersOK{}
}

/*
ListObjectStorageUsersOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListObjectStorageUsersOK struct {
	Payload *models.ListObjectStorageUsersResponse
}

// IsSuccess returns true when this list object storage users o k response has a 2xx status code
func (o *ListObjectStorageUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list object storage users o k response has a 3xx status code
func (o *ListObjectStorageUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list object storage users o k response has a 4xx status code
func (o *ListObjectStorageUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list object storage users o k response has a 5xx status code
func (o *ListObjectStorageUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list object storage users o k response a status code equal to that given
func (o *ListObjectStorageUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list object storage users o k response
func (o *ListObjectStorageUsersOK) Code() int {
	return 200
}

func (o *ListObjectStorageUsersOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/object-storage/users][%d] listObjectStorageUsersOK  %+v", 200, o.Payload)
}

func (o *ListObjectStorageUsersOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/object-storage/users][%d] listObjectStorageUsersOK  %+v", 200, o.Payload)
}

func (o *ListObjectStorageUsersOK) GetPayload() *models.ListObjectStorageUsersResponse {
	return o.Payload
}

func (o *ListObjectStorageUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListObjectStorageUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListObjectStorageUsersDefault creates a ListObjectStorageUsersDefault with default headers values
func NewListObjectStorageUsersDefault(code int) *ListObjectStorageUsersDefault {
	return &ListObjectStorageUsersDefault{
		_statusCode: code,
	}
}

/*
ListObjectStorageUsersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListObjectStorageUsersDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this list object storage users default response has a 2xx status code
func (o *ListObjectStorageUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list object storage users default response has a 3xx status code
func (o *ListObjectStorageUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list object storage users default response has a 4xx status code
func (o *ListObjectStorageUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list object storage users default response has a 5xx status code
func (o *ListObjectStorageUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list object storage users default response a status code equal to that given
func (o *ListObjectStorageUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list object storage users default response
func (o *ListObjectStorageUsersDefault) Code() int {
	return o._statusCode
}

func (o *ListObjectStorageUsersDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/object-storage/users][%d] ListObjectStorageUsers default  %+v", o._statusCode, o.Payload)
}

func (o *ListObjectStorageUsersDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/object-storage/users][%d] ListObjectStorageUsers default  %+v", o._statusCode, o.Payload)
}

func (o *ListObjectStorageUsersDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *ListObjectStorageUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
