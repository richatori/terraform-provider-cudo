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

// ListAPIKeysResponse list Api keys response
//
// swagger:model ListApiKeysResponse
type ListAPIKeysResponse struct {

	// api keys
	// Required: true
	APIKeys []*APIKey `json:"apiKeys"`

	// page number
	// Required: true
	PageNumber *int32 `json:"pageNumber"`

	// page size
	// Required: true
	PageSize *int32 `json:"pageSize"`

	// total count
	// Required: true
	TotalCount *int32 `json:"totalCount"`
}

// Validate validates this list Api keys response
func (m *ListAPIKeysResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListAPIKeysResponse) validateAPIKeys(formats strfmt.Registry) error {

	if err := validate.Required("apiKeys", "body", m.APIKeys); err != nil {
		return err
	}

	for i := 0; i < len(m.APIKeys); i++ {
		if swag.IsZero(m.APIKeys[i]) { // not required
			continue
		}

		if m.APIKeys[i] != nil {
			if err := m.APIKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apiKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListAPIKeysResponse) validatePageNumber(formats strfmt.Registry) error {

	if err := validate.Required("pageNumber", "body", m.PageNumber); err != nil {
		return err
	}

	return nil
}

func (m *ListAPIKeysResponse) validatePageSize(formats strfmt.Registry) error {

	if err := validate.Required("pageSize", "body", m.PageSize); err != nil {
		return err
	}

	return nil
}

func (m *ListAPIKeysResponse) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("totalCount", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this list Api keys response based on the context it is used
func (m *ListAPIKeysResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListAPIKeysResponse) contextValidateAPIKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.APIKeys); i++ {

		if m.APIKeys[i] != nil {

			if swag.IsZero(m.APIKeys[i]) { // not required
				return nil
			}

			if err := m.APIKeys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apiKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListAPIKeysResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAPIKeysResponse) UnmarshalBinary(b []byte) error {
	var res ListAPIKeysResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
