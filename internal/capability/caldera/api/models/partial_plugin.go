// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PartialPlugin partial plugin
//
// swagger:model Partial-Plugin
type PartialPlugin struct {

	// access
	Access int64 `json:"access,omitempty"`

	// address
	Address string `json:"address,omitempty"`

	// data dir
	DataDir string `json:"data_dir,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this partial plugin
func (m *PartialPlugin) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this partial plugin based on context it is used
func (m *PartialPlugin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PartialPlugin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartialPlugin) UnmarshalBinary(b []byte) error {
	var res PartialPlugin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
