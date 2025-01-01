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

	"soarca/pkg/core/capability/caldera/api/models"
)

// NewPatchAPIV2AbilitiesAbilityIDParams creates a new PatchAPIV2AbilitiesAbilityIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIV2AbilitiesAbilityIDParams() *PatchAPIV2AbilitiesAbilityIDParams {
	return &PatchAPIV2AbilitiesAbilityIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIV2AbilitiesAbilityIDParamsWithTimeout creates a new PatchAPIV2AbilitiesAbilityIDParams object
// with the ability to set a timeout on a request.
func NewPatchAPIV2AbilitiesAbilityIDParamsWithTimeout(timeout time.Duration) *PatchAPIV2AbilitiesAbilityIDParams {
	return &PatchAPIV2AbilitiesAbilityIDParams{
		timeout: timeout,
	}
}

// NewPatchAPIV2AbilitiesAbilityIDParamsWithContext creates a new PatchAPIV2AbilitiesAbilityIDParams object
// with the ability to set a context for a request.
func NewPatchAPIV2AbilitiesAbilityIDParamsWithContext(ctx context.Context) *PatchAPIV2AbilitiesAbilityIDParams {
	return &PatchAPIV2AbilitiesAbilityIDParams{
		Context: ctx,
	}
}

// NewPatchAPIV2AbilitiesAbilityIDParamsWithHTTPClient creates a new PatchAPIV2AbilitiesAbilityIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIV2AbilitiesAbilityIDParamsWithHTTPClient(client *http.Client) *PatchAPIV2AbilitiesAbilityIDParams {
	return &PatchAPIV2AbilitiesAbilityIDParams{
		HTTPClient: client,
	}
}

/*
PatchAPIV2AbilitiesAbilityIDParams contains all the parameters to send to the API endpoint

	for the patch API v2 abilities ability ID operation.

	Typically these are written to a http.Request.
*/
type PatchAPIV2AbilitiesAbilityIDParams struct {

	/* AbilityID.

	   UUID of the Ability to be retrieved
	*/
	AbilityID string

	// Body.
	Body *models.PartialAbility1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API v2 abilities ability ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2AbilitiesAbilityIDParams) WithDefaults() *PatchAPIV2AbilitiesAbilityIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API v2 abilities ability ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2AbilitiesAbilityIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API v2 abilities ability ID params
func (o *PatchAPIV2AbilitiesAbilityIDParams) WithTimeout(timeout time.Duration) *PatchAPIV2AbilitiesAbilityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API v2 abilities ability ID params
func (o *PatchAPIV2AbilitiesAbilityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API v2 abilities ability ID params
func (o *PatchAPIV2AbilitiesAbilityIDParams) WithContext(ctx context.Context) *PatchAPIV2AbilitiesAbilityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API v2 abilities ability ID params
func (o *PatchAPIV2AbilitiesAbilityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API v2 abilities ability ID params
func (o *PatchAPIV2AbilitiesAbilityIDParams) WithHTTPClient(client *http.Client) *PatchAPIV2AbilitiesAbilityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API v2 abilities ability ID params
func (o *PatchAPIV2AbilitiesAbilityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAbilityID adds the abilityID to the patch API v2 abilities ability ID params
func (o *PatchAPIV2AbilitiesAbilityIDParams) WithAbilityID(abilityID string) *PatchAPIV2AbilitiesAbilityIDParams {
	o.SetAbilityID(abilityID)
	return o
}

// SetAbilityID adds the abilityId to the patch API v2 abilities ability ID params
func (o *PatchAPIV2AbilitiesAbilityIDParams) SetAbilityID(abilityID string) {
	o.AbilityID = abilityID
}

// WithBody adds the body to the patch API v2 abilities ability ID params
func (o *PatchAPIV2AbilitiesAbilityIDParams) WithBody(body *models.PartialAbility1) *PatchAPIV2AbilitiesAbilityIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch API v2 abilities ability ID params
func (o *PatchAPIV2AbilitiesAbilityIDParams) SetBody(body *models.PartialAbility1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIV2AbilitiesAbilityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ability_id
	if err := r.SetPathParam("ability_id", o.AbilityID); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
