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

// Operation operation
//
// swagger:model Operation
type Operation struct {

	// adversary
	Adversary *Adversary `json:"adversary,omitempty"`

	// auto close
	AutoClose bool `json:"auto_close,omitempty"`

	// autonomous
	Autonomous int64 `json:"autonomous,omitempty"`

	// chain
	// Read Only: true
	Chain interface{} `json:"chain,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// host group
	// Read Only: true
	HostGroup []*Agent `json:"host_group"`

	// id
	ID string `json:"id,omitempty"`

	// jitter
	Jitter string `json:"jitter,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// obfuscator
	Obfuscator string `json:"obfuscator,omitempty"`

	// planner
	Planner *Planner `json:"planner,omitempty"`

	// source
	Source *Source `json:"source,omitempty"`

	// start
	// Read Only: true
	Start string `json:"start,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// use learning parsers
	UseLearningParsers bool `json:"use_learning_parsers,omitempty"`

	// visibility
	Visibility int64 `json:"visibility,omitempty"`
}

// Validate validates this operation
func (m *Operation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdversary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Operation) validateAdversary(formats strfmt.Registry) error {
	if swag.IsZero(m.Adversary) { // not required
		return nil
	}

	if m.Adversary != nil {
		if err := m.Adversary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adversary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adversary")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) validateHostGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.HostGroup) { // not required
		return nil
	}

	for i := 0; i < len(m.HostGroup); i++ {
		if swag.IsZero(m.HostGroup[i]) { // not required
			continue
		}

		if m.HostGroup[i] != nil {
			if err := m.HostGroup[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("host_group" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("host_group" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Operation) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Operation) validatePlanner(formats strfmt.Registry) error {
	if swag.IsZero(m.Planner) { // not required
		return nil
	}

	if m.Planner != nil {
		if err := m.Planner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("planner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("planner")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this operation based on the context it is used
func (m *Operation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdversary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlanner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Operation) contextValidateAdversary(ctx context.Context, formats strfmt.Registry) error {

	if m.Adversary != nil {

		if swag.IsZero(m.Adversary) { // not required
			return nil
		}

		if err := m.Adversary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adversary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adversary")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) contextValidateHostGroup(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "host_group", "body", []*Agent(m.HostGroup)); err != nil {
		return err
	}

	for i := 0; i < len(m.HostGroup); i++ {

		if m.HostGroup[i] != nil {

			if swag.IsZero(m.HostGroup[i]) { // not required
				return nil
			}

			if err := m.HostGroup[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("host_group" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("host_group" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Operation) contextValidatePlanner(ctx context.Context, formats strfmt.Registry) error {

	if m.Planner != nil {

		if swag.IsZero(m.Planner) { // not required
			return nil
		}

		if err := m.Planner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("planner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("planner")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if swag.IsZero(m.Source) { // not required
			return nil
		}

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *Operation) contextValidateStart(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "start", "body", string(m.Start)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Operation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Operation) UnmarshalBinary(b []byte) error {
	var res Operation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
