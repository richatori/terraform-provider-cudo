// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListVMMachineTypesRequest list VM machine types request
//
// swagger:model ListVMMachineTypesRequest
type ListVMMachineTypesRequest struct {

	// cpu model
	CPUModel string `json:"cpuModel,omitempty"`

	// data center Id
	DataCenterID string `json:"dataCenterId,omitempty"`

	// gpu
	Gpu int32 `json:"gpu,omitempty"`

	// gpu model
	GpuModel string `json:"gpuModel,omitempty"`

	// memory gib
	// Required: true
	MemoryGib *int32 `json:"memoryGib"`

	// order by
	OrderBy string `json:"orderBy,omitempty"`

	// page number
	PageNumber int32 `json:"pageNumber,omitempty"`

	// page size
	PageSize int32 `json:"pageSize,omitempty"`

	// public Ipv4
	PublicIPV4 bool `json:"publicIpv4,omitempty"`

	// region Id
	RegionID string `json:"regionId,omitempty"`

	// storage gib
	StorageGib int32 `json:"storageGib,omitempty"`

	// vcpu
	// Required: true
	Vcpu *int32 `json:"vcpu"`
}

// Validate validates this list VM machine types request
func (m *ListVMMachineTypesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMemoryGib(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcpu(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListVMMachineTypesRequest) validateMemoryGib(formats strfmt.Registry) error {

	if err := validate.Required("memoryGib", "body", m.MemoryGib); err != nil {
		return err
	}

	return nil
}

func (m *ListVMMachineTypesRequest) validateVcpu(formats strfmt.Registry) error {

	if err := validate.Required("vcpu", "body", m.Vcpu); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list VM machine types request based on context it is used
func (m *ListVMMachineTypesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListVMMachineTypesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListVMMachineTypesRequest) UnmarshalBinary(b []byte) error {
	var res ListVMMachineTypesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
