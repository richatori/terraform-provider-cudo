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

// DataCenter data center
//
// swagger:model DataCenter
type DataCenter struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// region Id
	// Required: true
	RegionID *string `json:"regionId"`

	// renewable energy
	// Required: true
	RenewableEnergy *bool `json:"renewableEnergy"`

	// supplier name
	// Required: true
	SupplierName *string `json:"supplierName"`
}

// Validate validates this data center
func (m *DataCenter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenewableEnergy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupplierName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataCenter) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DataCenter) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("regionId", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

func (m *DataCenter) validateRenewableEnergy(formats strfmt.Registry) error {

	if err := validate.Required("renewableEnergy", "body", m.RenewableEnergy); err != nil {
		return err
	}

	return nil
}

func (m *DataCenter) validateSupplierName(formats strfmt.Registry) error {

	if err := validate.Required("supplierName", "body", m.SupplierName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this data center based on context it is used
func (m *DataCenter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataCenter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataCenter) UnmarshalBinary(b []byte) error {
	var res DataCenter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
