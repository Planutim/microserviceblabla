// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Kind Kind A Kind represents the specific kind of type that a Type represents.
//
// The zero Kind is not a valid kind.
//
// swagger:model Kind
type Kind uint64

// Validate validates this kind
func (m Kind) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kind based on context it is used
func (m Kind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
