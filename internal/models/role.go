// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Role role
//
// swagger:model Role
type Role string

func NewRole(value Role) *Role {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Role.
func (m Role) Pointer() *Role {
	return &m
}

const (

	// RoleUNKNOWN captures enum value "UNKNOWN"
	RoleUNKNOWN Role = "UNKNOWN"

	// RoleOWNER captures enum value "OWNER"
	RoleOWNER Role = "OWNER"

	// RoleVIEWER captures enum value "VIEWER"
	RoleVIEWER Role = "VIEWER"

	// RoleEDITOR captures enum value "EDITOR"
	RoleEDITOR Role = "EDITOR"
)

// for schema
var roleEnum []interface{}

func init() {
	var res []Role
	if err := json.Unmarshal([]byte(`["UNKNOWN","OWNER","VIEWER","EDITOR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roleEnum = append(roleEnum, v)
	}
}

func (m Role) validateRoleEnum(path, location string, value Role) error {
	if err := validate.EnumCase(path, location, value, roleEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this role
func (m Role) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this role based on context it is used
func (m Role) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
