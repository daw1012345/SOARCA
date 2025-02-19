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
	"github.com/go-openapi/swag"
)

// NewHeadAPIV2AbilitiesAbilityIDParams creates a new HeadAPIV2AbilitiesAbilityIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadAPIV2AbilitiesAbilityIDParams() *HeadAPIV2AbilitiesAbilityIDParams {
	return &HeadAPIV2AbilitiesAbilityIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadAPIV2AbilitiesAbilityIDParamsWithTimeout creates a new HeadAPIV2AbilitiesAbilityIDParams object
// with the ability to set a timeout on a request.
func NewHeadAPIV2AbilitiesAbilityIDParamsWithTimeout(timeout time.Duration) *HeadAPIV2AbilitiesAbilityIDParams {
	return &HeadAPIV2AbilitiesAbilityIDParams{
		timeout: timeout,
	}
}

// NewHeadAPIV2AbilitiesAbilityIDParamsWithContext creates a new HeadAPIV2AbilitiesAbilityIDParams object
// with the ability to set a context for a request.
func NewHeadAPIV2AbilitiesAbilityIDParamsWithContext(ctx context.Context) *HeadAPIV2AbilitiesAbilityIDParams {
	return &HeadAPIV2AbilitiesAbilityIDParams{
		Context: ctx,
	}
}

// NewHeadAPIV2AbilitiesAbilityIDParamsWithHTTPClient creates a new HeadAPIV2AbilitiesAbilityIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadAPIV2AbilitiesAbilityIDParamsWithHTTPClient(client *http.Client) *HeadAPIV2AbilitiesAbilityIDParams {
	return &HeadAPIV2AbilitiesAbilityIDParams{
		HTTPClient: client,
	}
}

/*
HeadAPIV2AbilitiesAbilityIDParams contains all the parameters to send to the API endpoint

	for the head API v2 abilities ability ID operation.

	Typically these are written to a http.Request.
*/
type HeadAPIV2AbilitiesAbilityIDParams struct {

	/* AbilityID.

	   UUID of the Ability to be retrieved
	*/
	AbilityID string

	// Exclude.
	Exclude []string

	// Include.
	Include []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the head API v2 abilities ability ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2AbilitiesAbilityIDParams) WithDefaults() *HeadAPIV2AbilitiesAbilityIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the head API v2 abilities ability ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2AbilitiesAbilityIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) WithTimeout(timeout time.Duration) *HeadAPIV2AbilitiesAbilityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) WithContext(ctx context.Context) *HeadAPIV2AbilitiesAbilityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) WithHTTPClient(client *http.Client) *HeadAPIV2AbilitiesAbilityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAbilityID adds the abilityID to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) WithAbilityID(abilityID string) *HeadAPIV2AbilitiesAbilityIDParams {
	o.SetAbilityID(abilityID)
	return o
}

// SetAbilityID adds the abilityId to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) SetAbilityID(abilityID string) {
	o.AbilityID = abilityID
}

// WithExclude adds the exclude to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) WithExclude(exclude []string) *HeadAPIV2AbilitiesAbilityIDParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) WithInclude(include []string) *HeadAPIV2AbilitiesAbilityIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the head API v2 abilities ability ID params
func (o *HeadAPIV2AbilitiesAbilityIDParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *HeadAPIV2AbilitiesAbilityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ability_id
	if err := r.SetPathParam("ability_id", o.AbilityID); err != nil {
		return err
	}

	if o.Exclude != nil {

		// binding items for exclude
		joinedExclude := o.bindParamExclude(reg)

		// query array param exclude
		if err := r.SetQueryParam("exclude", joinedExclude...); err != nil {
			return err
		}
	}

	if o.Include != nil {

		// binding items for include
		joinedInclude := o.bindParamInclude(reg)

		// query array param include
		if err := r.SetQueryParam("include", joinedInclude...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamHeadAPIV2AbilitiesAbilityID binds the parameter exclude
func (o *HeadAPIV2AbilitiesAbilityIDParams) bindParamExclude(formats strfmt.Registry) []string {
	excludeIR := o.Exclude

	var excludeIC []string
	for _, excludeIIR := range excludeIR { // explode []string

		excludeIIV := excludeIIR // string as string
		excludeIC = append(excludeIC, excludeIIV)
	}

	// items.CollectionFormat: "multi"
	excludeIS := swag.JoinByFormat(excludeIC, "multi")

	return excludeIS
}

// bindParamHeadAPIV2AbilitiesAbilityID binds the parameter include
func (o *HeadAPIV2AbilitiesAbilityIDParams) bindParamInclude(formats strfmt.Registry) []string {
	includeIR := o.Include

	var includeIC []string
	for _, includeIIR := range includeIR { // explode []string

		includeIIV := includeIIR // string as string
		includeIC = append(includeIC, includeIIV)
	}

	// items.CollectionFormat: "multi"
	includeIS := swag.JoinByFormat(includeIC, "multi")

	return includeIS
}
