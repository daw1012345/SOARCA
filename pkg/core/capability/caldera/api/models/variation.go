// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Variation variation
//
// swagger:model Variation
type Variation struct {

	// command
	Command string `json:"command,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this variation
func (m *Variation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this variation based on context it is used
func (m *Variation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Variation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Variation) UnmarshalBinary(b []byte) error {
	var res Variation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
