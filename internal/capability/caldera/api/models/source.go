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
)

// Source source
//
// swagger:model Source
type Source struct {

	// adjustments
	Adjustments []*Adjustment `json:"adjustments"`

	// facts
	Facts []*Fact `json:"facts"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// plugin
	Plugin *string `json:"plugin,omitempty"`

	// relationships
	Relationships []*Relationship `json:"relationships"`

	// rules
	Rules []*Rule `json:"rules"`
}

// Validate validates this source
func (m *Source) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdjustments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Source) validateAdjustments(formats strfmt.Registry) error {
	if swag.IsZero(m.Adjustments) { // not required
		return nil
	}

	for i := 0; i < len(m.Adjustments); i++ {
		if swag.IsZero(m.Adjustments[i]) { // not required
			continue
		}

		if m.Adjustments[i] != nil {
			if err := m.Adjustments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adjustments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("adjustments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Source) validateFacts(formats strfmt.Registry) error {
	if swag.IsZero(m.Facts) { // not required
		return nil
	}

	for i := 0; i < len(m.Facts); i++ {
		if swag.IsZero(m.Facts[i]) { // not required
			continue
		}

		if m.Facts[i] != nil {
			if err := m.Facts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("facts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("facts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Source) validateRelationships(formats strfmt.Registry) error {
	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	for i := 0; i < len(m.Relationships); i++ {
		if swag.IsZero(m.Relationships[i]) { // not required
			continue
		}

		if m.Relationships[i] != nil {
			if err := m.Relationships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Source) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this source based on the context it is used
func (m *Source) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdjustments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelationships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Source) contextValidateAdjustments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Adjustments); i++ {

		if m.Adjustments[i] != nil {

			if swag.IsZero(m.Adjustments[i]) { // not required
				return nil
			}

			if err := m.Adjustments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adjustments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("adjustments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Source) contextValidateFacts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Facts); i++ {

		if m.Facts[i] != nil {

			if swag.IsZero(m.Facts[i]) { // not required
				return nil
			}

			if err := m.Facts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("facts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("facts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Source) contextValidateRelationships(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Relationships); i++ {

		if m.Relationships[i] != nil {

			if swag.IsZero(m.Relationships[i]) { // not required
				return nil
			}

			if err := m.Relationships[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Source) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rules); i++ {

		if m.Rules[i] != nil {

			if swag.IsZero(m.Rules[i]) { // not required
				return nil
			}

			if err := m.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Source) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Source) UnmarshalBinary(b []byte) error {
	var res Source
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
