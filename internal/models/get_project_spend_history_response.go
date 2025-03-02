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

// GetProjectSpendHistoryResponse get project spend history response
//
// swagger:model GetProjectSpendHistoryResponse
type GetProjectSpendHistoryResponse struct {

	// project spend history
	// Required: true
	ProjectSpendHistory []*ProjectSpend `json:"projectSpendHistory"`
}

// Validate validates this get project spend history response
func (m *GetProjectSpendHistoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectSpendHistory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProjectSpendHistoryResponse) validateProjectSpendHistory(formats strfmt.Registry) error {

	if err := validate.Required("projectSpendHistory", "body", m.ProjectSpendHistory); err != nil {
		return err
	}

	for i := 0; i < len(m.ProjectSpendHistory); i++ {
		if swag.IsZero(m.ProjectSpendHistory[i]) { // not required
			continue
		}

		if m.ProjectSpendHistory[i] != nil {
			if err := m.ProjectSpendHistory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectSpendHistory" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projectSpendHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get project spend history response based on the context it is used
func (m *GetProjectSpendHistoryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProjectSpendHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProjectSpendHistoryResponse) contextValidateProjectSpendHistory(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProjectSpendHistory); i++ {

		if m.ProjectSpendHistory[i] != nil {
			if err := m.ProjectSpendHistory[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectSpendHistory" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projectSpendHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetProjectSpendHistoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProjectSpendHistoryResponse) UnmarshalBinary(b []byte) error {
	var res GetProjectSpendHistoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
