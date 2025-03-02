// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ObjectStorageBucket object storage bucket
//
// swagger:model ObjectStorageBucket
type ObjectStorageBucket struct {

	// billable bytes
	BillableBytes string `json:"billableBytes,omitempty"`

	// data center Id
	DataCenterID string `json:"dataCenterId,omitempty"`

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// object count
	ObjectCount string `json:"objectCount,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// size bytes
	SizeBytes string `json:"sizeBytes,omitempty"`
}

// Validate validates this object storage bucket
func (m *ObjectStorageBucket) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this object storage bucket based on context it is used
func (m *ObjectStorageBucket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectStorageBucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectStorageBucket) UnmarshalBinary(b []byte) error {
	var res ObjectStorageBucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
