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

// Executor executor
//
// swagger:model Executor
type Executor struct {

	// additional info
	AdditionalInfo map[string]string `json:"additional_info,omitempty"`

	// build target
	BuildTarget *string `json:"build_target,omitempty"`

	// cleanup
	Cleanup []string `json:"cleanup"`

	// code
	Code *string `json:"code,omitempty"`

	// command
	Command *string `json:"command,omitempty"`

	// language
	Language *string `json:"language,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// parsers
	Parsers []*Parser `json:"parsers"`

	// payloads
	Payloads []string `json:"payloads"`

	// platform
	Platform *string `json:"platform,omitempty"`

	// timeout
	Timeout *int64 `json:"timeout,omitempty"`

	// uploads
	Uploads []string `json:"uploads"`

	// variations
	Variations []*Variation `json:"variations"`
}

// Validate validates this executor
func (m *Executor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Executor) validateParsers(formats strfmt.Registry) error {
	if swag.IsZero(m.Parsers) { // not required
		return nil
	}

	for i := 0; i < len(m.Parsers); i++ {
		if swag.IsZero(m.Parsers[i]) { // not required
			continue
		}

		if m.Parsers[i] != nil {
			if err := m.Parsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Executor) validateVariations(formats strfmt.Registry) error {
	if swag.IsZero(m.Variations) { // not required
		return nil
	}

	for i := 0; i < len(m.Variations); i++ {
		if swag.IsZero(m.Variations[i]) { // not required
			continue
		}

		if m.Variations[i] != nil {
			if err := m.Variations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this executor based on the context it is used
func (m *Executor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Executor) contextValidateParsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parsers); i++ {

		if m.Parsers[i] != nil {

			if swag.IsZero(m.Parsers[i]) { // not required
				return nil
			}

			if err := m.Parsers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Executor) contextValidateVariations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variations); i++ {

		if m.Variations[i] != nil {

			if swag.IsZero(m.Variations[i]) { // not required
				return nil
			}

			if err := m.Variations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Executor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Executor) UnmarshalBinary(b []byte) error {
	var res Executor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
