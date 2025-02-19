// Code generated by go-swagger; DO NOT EDIT.

package abilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteAPIV2AbilitiesAbilityIDParams creates a new DeleteAPIV2AbilitiesAbilityIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPIV2AbilitiesAbilityIDParams() *DeleteAPIV2AbilitiesAbilityIDParams {
	return &DeleteAPIV2AbilitiesAbilityIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV2AbilitiesAbilityIDParamsWithTimeout creates a new DeleteAPIV2AbilitiesAbilityIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAPIV2AbilitiesAbilityIDParamsWithTimeout(timeout time.Duration) *DeleteAPIV2AbilitiesAbilityIDParams {
	return &DeleteAPIV2AbilitiesAbilityIDParams{
		timeout: timeout,
	}
}

// NewDeleteAPIV2AbilitiesAbilityIDParamsWithContext creates a new DeleteAPIV2AbilitiesAbilityIDParams object
// with the ability to set a context for a request.
func NewDeleteAPIV2AbilitiesAbilityIDParamsWithContext(ctx context.Context) *DeleteAPIV2AbilitiesAbilityIDParams {
	return &DeleteAPIV2AbilitiesAbilityIDParams{
		Context: ctx,
	}
}

// NewDeleteAPIV2AbilitiesAbilityIDParamsWithHTTPClient creates a new DeleteAPIV2AbilitiesAbilityIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPIV2AbilitiesAbilityIDParamsWithHTTPClient(client *http.Client) *DeleteAPIV2AbilitiesAbilityIDParams {
	return &DeleteAPIV2AbilitiesAbilityIDParams{
		HTTPClient: client,
	}
}

/*
DeleteAPIV2AbilitiesAbilityIDParams contains all the parameters to send to the API endpoint

	for the delete API v2 abilities ability ID operation.

	Typically these are written to a http.Request.
*/
type DeleteAPIV2AbilitiesAbilityIDParams struct {

	/* AbilityID.

	   UUID of the Ability to be retrieved
	*/
	AbilityID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API v2 abilities ability ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIV2AbilitiesAbilityIDParams) WithDefaults() *DeleteAPIV2AbilitiesAbilityIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API v2 abilities ability ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIV2AbilitiesAbilityIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API v2 abilities ability ID params
func (o *DeleteAPIV2AbilitiesAbilityIDParams) WithTimeout(timeout time.Duration) *DeleteAPIV2AbilitiesAbilityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API v2 abilities ability ID params
func (o *DeleteAPIV2AbilitiesAbilityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API v2 abilities ability ID params
func (o *DeleteAPIV2AbilitiesAbilityIDParams) WithContext(ctx context.Context) *DeleteAPIV2AbilitiesAbilityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API v2 abilities ability ID params
func (o *DeleteAPIV2AbilitiesAbilityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API v2 abilities ability ID params
func (o *DeleteAPIV2AbilitiesAbilityIDParams) WithHTTPClient(client *http.Client) *DeleteAPIV2AbilitiesAbilityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API v2 abilities ability ID params
func (o *DeleteAPIV2AbilitiesAbilityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAbilityID adds the abilityID to the delete API v2 abilities ability ID params
func (o *DeleteAPIV2AbilitiesAbilityIDParams) WithAbilityID(abilityID string) *DeleteAPIV2AbilitiesAbilityIDParams {
	o.SetAbilityID(abilityID)
	return o
}

// SetAbilityID adds the abilityId to the delete API v2 abilities ability ID params
func (o *DeleteAPIV2AbilitiesAbilityIDParams) SetAbilityID(abilityID string) {
	o.AbilityID = abilityID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV2AbilitiesAbilityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ability_id
	if err := r.SetPathParam("ability_id", o.AbilityID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
