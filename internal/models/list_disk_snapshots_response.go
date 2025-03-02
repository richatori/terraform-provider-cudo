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

// ListDiskSnapshotsResponse list disk snapshots response
//
// swagger:model ListDiskSnapshotsResponse
type ListDiskSnapshotsResponse struct {

	// snapshots
	// Required: true
	Snapshots []*Snapshot `json:"snapshots"`
}

// Validate validates this list disk snapshots response
func (m *ListDiskSnapshotsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDiskSnapshotsResponse) validateSnapshots(formats strfmt.Registry) error {

	if err := validate.Required("snapshots", "body", m.Snapshots); err != nil {
		return err
	}

	for i := 0; i < len(m.Snapshots); i++ {
		if swag.IsZero(m.Snapshots[i]) { // not required
			continue
		}

		if m.Snapshots[i] != nil {
			if err := m.Snapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list disk snapshots response based on the context it is used
func (m *ListDiskSnapshotsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDiskSnapshotsResponse) contextValidateSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Snapshots); i++ {

		if m.Snapshots[i] != nil {

			if swag.IsZero(m.Snapshots[i]) { // not required
				return nil
			}

			if err := m.Snapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListDiskSnapshotsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDiskSnapshotsResponse) UnmarshalBinary(b []byte) error {
	var res ListDiskSnapshotsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
