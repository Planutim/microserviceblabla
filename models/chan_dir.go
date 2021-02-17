// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ChanDir ChanDir ChanDir represents a channel type's direction.
//
// swagger:model ChanDir
type ChanDir int64

// Validate validates this chan dir
func (m ChanDir) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chan dir based on context it is used
func (m ChanDir) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
