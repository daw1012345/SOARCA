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

// Adversary adversary
//
// swagger:model Adversary
type Adversary struct {

	// adversary id
	AdversaryID string `json:"adversary_id,omitempty"`

	// atomic ordering
	AtomicOrdering []string `json:"atomic_ordering"`

	// description
	Description string `json:"description,omitempty"`

	// has repeatable abilities
	// Read Only: true
	HasRepeatableAbilities *bool `json:"has_repeatable_abilities,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// objective
	Objective string `json:"objective,omitempty"`

	// plugin
	Plugin *string `json:"plugin,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this adversary
func (m *Adversary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this adversary based on the context it is used
func (m *Adversary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHasRepeatableAbilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Adversary) contextValidateHasRepeatableAbilities(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "has_repeatable_abilities", "body", m.HasRepeatableAbilities); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Adversary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Adversary) UnmarshalBinary(b []byte) error {
	var res Adversary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
