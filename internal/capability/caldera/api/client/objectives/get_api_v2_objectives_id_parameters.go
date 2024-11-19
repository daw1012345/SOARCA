// Code generated by go-swagger; DO NOT EDIT.

package objectives

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

// NewGetAPIV2ObjectivesIDParams creates a new GetAPIV2ObjectivesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2ObjectivesIDParams() *GetAPIV2ObjectivesIDParams {
	return &GetAPIV2ObjectivesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2ObjectivesIDParamsWithTimeout creates a new GetAPIV2ObjectivesIDParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2ObjectivesIDParamsWithTimeout(timeout time.Duration) *GetAPIV2ObjectivesIDParams {
	return &GetAPIV2ObjectivesIDParams{
		timeout: timeout,
	}
}

// NewGetAPIV2ObjectivesIDParamsWithContext creates a new GetAPIV2ObjectivesIDParams object
// with the ability to set a context for a request.
func NewGetAPIV2ObjectivesIDParamsWithContext(ctx context.Context) *GetAPIV2ObjectivesIDParams {
	return &GetAPIV2ObjectivesIDParams{
		Context: ctx,
	}
}

// NewGetAPIV2ObjectivesIDParamsWithHTTPClient creates a new GetAPIV2ObjectivesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2ObjectivesIDParamsWithHTTPClient(client *http.Client) *GetAPIV2ObjectivesIDParams {
	return &GetAPIV2ObjectivesIDParams{
		HTTPClient: client,
	}
}

/*
GetAPIV2ObjectivesIDParams contains all the parameters to send to the API endpoint

	for the get API v2 objectives ID operation.

	Typically these are written to a http.Request.
*/
type GetAPIV2ObjectivesIDParams struct {

	// Exclude.
	Exclude []string

	/* ID.

	   UUID of the objective to be retrieved
	*/
	ID string

	// Include.
	Include []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 objectives ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ObjectivesIDParams) WithDefaults() *GetAPIV2ObjectivesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 objectives ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ObjectivesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) WithTimeout(timeout time.Duration) *GetAPIV2ObjectivesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) WithContext(ctx context.Context) *GetAPIV2ObjectivesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) WithHTTPClient(client *http.Client) *GetAPIV2ObjectivesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) WithExclude(exclude []string) *GetAPIV2ObjectivesIDParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithID adds the id to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) WithID(id string) *GetAPIV2ObjectivesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) SetID(id string) {
	o.ID = id
}

// WithInclude adds the include to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) WithInclude(include []string) *GetAPIV2ObjectivesIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get API v2 objectives ID params
func (o *GetAPIV2ObjectivesIDParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2ObjectivesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Exclude != nil {

		// binding items for exclude
		joinedExclude := o.bindParamExclude(reg)

		// query array param exclude
		if err := r.SetQueryParam("exclude", joinedExclude...); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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

// bindParamGetAPIV2ObjectivesID binds the parameter exclude
func (o *GetAPIV2ObjectivesIDParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamGetAPIV2ObjectivesID binds the parameter include
func (o *GetAPIV2ObjectivesIDParams) bindParamInclude(formats strfmt.Registry) []string {
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
