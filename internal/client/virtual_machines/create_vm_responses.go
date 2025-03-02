// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

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

// CreateVMReader is a Reader for the CreateVM structure.
type CreateVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateVMDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateVMOK creates a CreateVMOK with default headers values
func NewCreateVMOK() *CreateVMOK {
	return &CreateVMOK{}
}

/*
CreateVMOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateVMOK struct {
	Payload *models.CreateVMResponse
}

// IsSuccess returns true when this create Vm o k response has a 2xx status code
func (o *CreateVMOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Vm o k response has a 3xx status code
func (o *CreateVMOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Vm o k response has a 4xx status code
func (o *CreateVMOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Vm o k response has a 5xx status code
func (o *CreateVMOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create Vm o k response a status code equal to that given
func (o *CreateVMOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create Vm o k response
func (o *CreateVMOK) Code() int {
	return 200
}

func (o *CreateVMOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vm][%d] createVmOK  %+v", 200, o.Payload)
}

func (o *CreateVMOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vm][%d] createVmOK  %+v", 200, o.Payload)
}

func (o *CreateVMOK) GetPayload() *models.CreateVMResponse {
	return o.Payload
}

func (o *CreateVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateVMResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVMDefault creates a CreateVMDefault with default headers values
func NewCreateVMDefault(code int) *CreateVMDefault {
	return &CreateVMDefault{
		_statusCode: code,
	}
}

/*
CreateVMDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateVMDefault struct {
	_statusCode int

	Payload *models.Status
}

// IsSuccess returns true when this create VM default response has a 2xx status code
func (o *CreateVMDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create VM default response has a 3xx status code
func (o *CreateVMDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create VM default response has a 4xx status code
func (o *CreateVMDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create VM default response has a 5xx status code
func (o *CreateVMDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create VM default response a status code equal to that given
func (o *CreateVMDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create VM default response
func (o *CreateVMDefault) Code() int {
	return o._statusCode
}

func (o *CreateVMDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vm][%d] CreateVM default  %+v", o._statusCode, o.Payload)
}

func (o *CreateVMDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/vm][%d] CreateVM default  %+v", o._statusCode, o.Payload)
}

func (o *CreateVMDefault) GetPayload() *models.Status {
	return o.Payload
}

func (o *CreateVMDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateVMBody create VM body
swagger:model CreateVMBody
*/
type CreateVMBody struct {

	// boot disk
	BootDisk *models.Disk `json:"bootDisk,omitempty"`

	// boot disk image Id
	// Required: true
	BootDiskImageID *string `json:"bootDiskImageId"`

	// cpu model
	CPUModel string `json:"cpuModel,omitempty"`

	// custom Ssh keys
	CustomSSHKeys []string `json:"customSshKeys"`

	// data center Id
	DataCenterID string `json:"dataCenterId,omitempty"`

	// gpu model
	GpuModel string `json:"gpuModel,omitempty"`

	// gpus
	Gpus int32 `json:"gpus,omitempty"`

	// machine type
	MachineType string `json:"machineType,omitempty"`

	// max price hr
	MaxPriceHr *models.Decimal `json:"maxPriceHr,omitempty"`

	// memory gib
	MemoryGib int32 `json:"memoryGib,omitempty"`

	// metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// nics
	Nics []*models.CreateVMRequestNIC `json:"nics"`

	// password
	Password string `json:"password,omitempty"`

	// ignored if any nics are provided
	SecurityGroupIds []string `json:"securityGroupIds"`

	// ssh key source
	SSHKeySource *models.SSHKeySource `json:"sshKeySource,omitempty"`

	// start script
	StartScript string `json:"startScript,omitempty"`

	// storage disk ids
	StorageDiskIds []string `json:"storageDiskIds"`

	// vcpus
	Vcpus int32 `json:"vcpus,omitempty"`

	// vm Id
	// Required: true
	VMID *string `json:"vmId"`
}

// Validate validates this create VM body
func (o *CreateVMBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBootDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBootDiskImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMaxPriceHr(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSSHKeySource(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVMID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVMBody) validateBootDisk(formats strfmt.Registry) error {
	if swag.IsZero(o.BootDisk) { // not required
		return nil
	}

	if o.BootDisk != nil {
		if err := o.BootDisk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "bootDisk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "bootDisk")
			}
			return err
		}
	}

	return nil
}

func (o *CreateVMBody) validateBootDiskImageID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"bootDiskImageId", "body", o.BootDiskImageID); err != nil {
		return err
	}

	return nil
}

func (o *CreateVMBody) validateMaxPriceHr(formats strfmt.Registry) error {
	if swag.IsZero(o.MaxPriceHr) { // not required
		return nil
	}

	if o.MaxPriceHr != nil {
		if err := o.MaxPriceHr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "maxPriceHr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "maxPriceHr")
			}
			return err
		}
	}

	return nil
}

func (o *CreateVMBody) validateNics(formats strfmt.Registry) error {
	if swag.IsZero(o.Nics) { // not required
		return nil
	}

	for i := 0; i < len(o.Nics); i++ {
		if swag.IsZero(o.Nics[i]) { // not required
			continue
		}

		if o.Nics[i] != nil {
			if err := o.Nics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateVMBody) validateSSHKeySource(formats strfmt.Registry) error {
	if swag.IsZero(o.SSHKeySource) { // not required
		return nil
	}

	if o.SSHKeySource != nil {
		if err := o.SSHKeySource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "sshKeySource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "sshKeySource")
			}
			return err
		}
	}

	return nil
}

func (o *CreateVMBody) validateVMID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"vmId", "body", o.VMID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create VM body based on the context it is used
func (o *CreateVMBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBootDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMaxPriceHr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSSHKeySource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVMBody) contextValidateBootDisk(ctx context.Context, formats strfmt.Registry) error {

	if o.BootDisk != nil {

		if swag.IsZero(o.BootDisk) { // not required
			return nil
		}

		if err := o.BootDisk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "bootDisk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "bootDisk")
			}
			return err
		}
	}

	return nil
}

func (o *CreateVMBody) contextValidateMaxPriceHr(ctx context.Context, formats strfmt.Registry) error {

	if o.MaxPriceHr != nil {

		if swag.IsZero(o.MaxPriceHr) { // not required
			return nil
		}

		if err := o.MaxPriceHr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "maxPriceHr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "maxPriceHr")
			}
			return err
		}
	}

	return nil
}

func (o *CreateVMBody) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Nics); i++ {

		if o.Nics[i] != nil {

			if swag.IsZero(o.Nics[i]) { // not required
				return nil
			}

			if err := o.Nics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateVMBody) contextValidateSSHKeySource(ctx context.Context, formats strfmt.Registry) error {

	if o.SSHKeySource != nil {

		if swag.IsZero(o.SSHKeySource) { // not required
			return nil
		}

		if err := o.SSHKeySource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "sshKeySource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "sshKeySource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateVMBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVMBody) UnmarshalBinary(b []byte) error {
	var res CreateVMBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
