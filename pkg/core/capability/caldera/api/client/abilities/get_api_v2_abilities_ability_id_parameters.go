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

// NewGetAPIV2AbilitiesAbilityIDParams creates a new GetAPIV2AbilitiesAbilityIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2AbilitiesAbilityIDParams() *GetAPIV2AbilitiesAbilityIDParams {
	return &GetAPIV2AbilitiesAbilityIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2AbilitiesAbilityIDParamsWithTimeout creates a new GetAPIV2AbilitiesAbilityIDParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2AbilitiesAbilityIDParamsWithTimeout(timeout time.Duration) *GetAPIV2AbilitiesAbilityIDParams {
	return &GetAPIV2AbilitiesAbilityIDParams{
		timeout: timeout,
	}
}

// NewGetAPIV2AbilitiesAbilityIDParamsWithContext creates a new GetAPIV2AbilitiesAbilityIDParams object
// with the ability to set a context for a request.
func NewGetAPIV2AbilitiesAbilityIDParamsWithContext(ctx context.Context) *GetAPIV2AbilitiesAbilityIDParams {
	return &GetAPIV2AbilitiesAbilityIDParams{
		Context: ctx,
	}
}

// NewGetAPIV2AbilitiesAbilityIDParamsWithHTTPClient creates a new GetAPIV2AbilitiesAbilityIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2AbilitiesAbilityIDParamsWithHTTPClient(client *http.Client) *GetAPIV2AbilitiesAbilityIDParams {
	return &GetAPIV2AbilitiesAbilityIDParams{
		HTTPClient: client,
	}
}

/*
GetAPIV2AbilitiesAbilityIDParams contains all the parameters to send to the API endpoint

	for the get API v2 abilities ability ID operation.

	Typically these are written to a http.Request.
*/
type GetAPIV2AbilitiesAbilityIDParams struct {

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

// WithDefaults hydrates default values in the get API v2 abilities ability ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2AbilitiesAbilityIDParams) WithDefaults() *GetAPIV2AbilitiesAbilityIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 abilities ability ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2AbilitiesAbilityIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) WithTimeout(timeout time.Duration) *GetAPIV2AbilitiesAbilityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) WithContext(ctx context.Context) *GetAPIV2AbilitiesAbilityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) WithHTTPClient(client *http.Client) *GetAPIV2AbilitiesAbilityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAbilityID adds the abilityID to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) WithAbilityID(abilityID string) *GetAPIV2AbilitiesAbilityIDParams {
	o.SetAbilityID(abilityID)
	return o
}

// SetAbilityID adds the abilityId to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) SetAbilityID(abilityID string) {
	o.AbilityID = abilityID
}

// WithExclude adds the exclude to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) WithExclude(exclude []string) *GetAPIV2AbilitiesAbilityIDParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) WithInclude(include []string) *GetAPIV2AbilitiesAbilityIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get API v2 abilities ability ID params
func (o *GetAPIV2AbilitiesAbilityIDParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2AbilitiesAbilityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetAPIV2AbilitiesAbilityID binds the parameter exclude
func (o *GetAPIV2AbilitiesAbilityIDParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamGetAPIV2AbilitiesAbilityID binds the parameter include
func (o *GetAPIV2AbilitiesAbilityIDParams) bindParamInclude(formats strfmt.Registry) []string {
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
