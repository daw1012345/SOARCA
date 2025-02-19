// Code generated by go-swagger; DO NOT EDIT.

package schedules

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

// NewHeadAPIV2SchedulesIDParams creates a new HeadAPIV2SchedulesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadAPIV2SchedulesIDParams() *HeadAPIV2SchedulesIDParams {
	return &HeadAPIV2SchedulesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadAPIV2SchedulesIDParamsWithTimeout creates a new HeadAPIV2SchedulesIDParams object
// with the ability to set a timeout on a request.
func NewHeadAPIV2SchedulesIDParamsWithTimeout(timeout time.Duration) *HeadAPIV2SchedulesIDParams {
	return &HeadAPIV2SchedulesIDParams{
		timeout: timeout,
	}
}

// NewHeadAPIV2SchedulesIDParamsWithContext creates a new HeadAPIV2SchedulesIDParams object
// with the ability to set a context for a request.
func NewHeadAPIV2SchedulesIDParamsWithContext(ctx context.Context) *HeadAPIV2SchedulesIDParams {
	return &HeadAPIV2SchedulesIDParams{
		Context: ctx,
	}
}

// NewHeadAPIV2SchedulesIDParamsWithHTTPClient creates a new HeadAPIV2SchedulesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadAPIV2SchedulesIDParamsWithHTTPClient(client *http.Client) *HeadAPIV2SchedulesIDParams {
	return &HeadAPIV2SchedulesIDParams{
		HTTPClient: client,
	}
}

/*
HeadAPIV2SchedulesIDParams contains all the parameters to send to the API endpoint

	for the head API v2 schedules ID operation.

	Typically these are written to a http.Request.
*/
type HeadAPIV2SchedulesIDParams struct {

	// Exclude.
	Exclude []string

	/* ID.

	   UUID of the Schedule to be retrieved.
	*/
	ID string

	// Include.
	Include []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the head API v2 schedules ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2SchedulesIDParams) WithDefaults() *HeadAPIV2SchedulesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the head API v2 schedules ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2SchedulesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) WithTimeout(timeout time.Duration) *HeadAPIV2SchedulesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) WithContext(ctx context.Context) *HeadAPIV2SchedulesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) WithHTTPClient(client *http.Client) *HeadAPIV2SchedulesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) WithExclude(exclude []string) *HeadAPIV2SchedulesIDParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithID adds the id to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) WithID(id string) *HeadAPIV2SchedulesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) SetID(id string) {
	o.ID = id
}

// WithInclude adds the include to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) WithInclude(include []string) *HeadAPIV2SchedulesIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the head API v2 schedules ID params
func (o *HeadAPIV2SchedulesIDParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *HeadAPIV2SchedulesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamHeadAPIV2SchedulesID binds the parameter exclude
func (o *HeadAPIV2SchedulesIDParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamHeadAPIV2SchedulesID binds the parameter include
func (o *HeadAPIV2SchedulesIDParams) bindParamInclude(formats strfmt.Registry) []string {
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
