// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListUserPermissionsResponse list user permissions response
//
// swagger:model ListUserPermissionsResponse
type ListUserPermissionsResponse struct {

	// billing account permissions
	// Required: true
	BillingAccountPermissions []*UserPermission `json:"billingAccountPermissions"`

	// data center permissions
	// Required: true
	DataCenterPermissions []*UserPermission `json:"dataCenterPermissions"`

	// project permissions
	// Required: true
	ProjectPermissions []*UserPermission `json:"projectPermissions"`
}

// Validate validates this list user permissions response
func (m *ListUserPermissionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingAccountPermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataCenterPermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectPermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListUserPermissionsResponse) validateBillingAccountPermissions(formats strfmt.Registry) error {

	if err := validate.Required("billingAccountPermissions", "body", m.BillingAccountPermissions); err != nil {
		return err
	}

	for i := 0; i < len(m.BillingAccountPermissions); i++ {
		if swag.IsZero(m.BillingAccountPermissions[i]) { // not required
			continue
		}

		if m.BillingAccountPermissions[i] != nil {
			if err := m.BillingAccountPermissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billingAccountPermissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("billingAccountPermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListUserPermissionsResponse) validateDataCenterPermissions(formats strfmt.Registry) error {

	if err := validate.Required("dataCenterPermissions", "body", m.DataCenterPermissions); err != nil {
		return err
	}

	for i := 0; i < len(m.DataCenterPermissions); i++ {
		if swag.IsZero(m.DataCenterPermissions[i]) { // not required
			continue
		}

		if m.DataCenterPermissions[i] != nil {
			if err := m.DataCenterPermissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataCenterPermissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataCenterPermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListUserPermissionsResponse) validateProjectPermissions(formats strfmt.Registry) error {

	if err := validate.Required("projectPermissions", "body", m.ProjectPermissions); err != nil {
		return err
	}

	for i := 0; i < len(m.ProjectPermissions); i++ {
		if swag.IsZero(m.ProjectPermissions[i]) { // not required
			continue
		}

		if m.ProjectPermissions[i] != nil {
			if err := m.ProjectPermissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectPermissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projectPermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list user permissions response based on the context it is used
func (m *ListUserPermissionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillingAccountPermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataCenterPermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectPermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListUserPermissionsResponse) contextValidateBillingAccountPermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BillingAccountPermissions); i++ {

		if m.BillingAccountPermissions[i] != nil {

			if swag.IsZero(m.BillingAccountPermissions[i]) { // not required
				return nil
			}

			if err := m.BillingAccountPermissions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billingAccountPermissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("billingAccountPermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListUserPermissionsResponse) contextValidateDataCenterPermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DataCenterPermissions); i++ {

		if m.DataCenterPermissions[i] != nil {

			if swag.IsZero(m.DataCenterPermissions[i]) { // not required
				return nil
			}

			if err := m.DataCenterPermissions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataCenterPermissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataCenterPermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListUserPermissionsResponse) contextValidateProjectPermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProjectPermissions); i++ {

		if m.ProjectPermissions[i] != nil {

			if swag.IsZero(m.ProjectPermissions[i]) { // not required
				return nil
			}

			if err := m.ProjectPermissions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectPermissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projectPermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListUserPermissionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListUserPermissionsResponse) UnmarshalBinary(b []byte) error {
	var res ListUserPermissionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
