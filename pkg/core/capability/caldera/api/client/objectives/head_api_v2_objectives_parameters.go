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

// NewHeadAPIV2ObjectivesParams creates a new HeadAPIV2ObjectivesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadAPIV2ObjectivesParams() *HeadAPIV2ObjectivesParams {
	return &HeadAPIV2ObjectivesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadAPIV2ObjectivesParamsWithTimeout creates a new HeadAPIV2ObjectivesParams object
// with the ability to set a timeout on a request.
func NewHeadAPIV2ObjectivesParamsWithTimeout(timeout time.Duration) *HeadAPIV2ObjectivesParams {
	return &HeadAPIV2ObjectivesParams{
		timeout: timeout,
	}
}

// NewHeadAPIV2ObjectivesParamsWithContext creates a new HeadAPIV2ObjectivesParams object
// with the ability to set a context for a request.
func NewHeadAPIV2ObjectivesParamsWithContext(ctx context.Context) *HeadAPIV2ObjectivesParams {
	return &HeadAPIV2ObjectivesParams{
		Context: ctx,
	}
}

// NewHeadAPIV2ObjectivesParamsWithHTTPClient creates a new HeadAPIV2ObjectivesParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadAPIV2ObjectivesParamsWithHTTPClient(client *http.Client) *HeadAPIV2ObjectivesParams {
	return &HeadAPIV2ObjectivesParams{
		HTTPClient: client,
	}
}

/*
HeadAPIV2ObjectivesParams contains all the parameters to send to the API endpoint

	for the head API v2 objectives operation.

	Typically these are written to a http.Request.
*/
type HeadAPIV2ObjectivesParams struct {

	// Exclude.
	Exclude []string

	// Include.
	Include []string

	// Sort.
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the head API v2 objectives params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2ObjectivesParams) WithDefaults() *HeadAPIV2ObjectivesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the head API v2 objectives params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2ObjectivesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) WithTimeout(timeout time.Duration) *HeadAPIV2ObjectivesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) WithContext(ctx context.Context) *HeadAPIV2ObjectivesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) WithHTTPClient(client *http.Client) *HeadAPIV2ObjectivesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) WithExclude(exclude []string) *HeadAPIV2ObjectivesParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) WithInclude(include []string) *HeadAPIV2ObjectivesParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) SetInclude(include []string) {
	o.Include = include
}

// WithSort adds the sort to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) WithSort(sort *string) *HeadAPIV2ObjectivesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the head API v2 objectives params
func (o *HeadAPIV2ObjectivesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *HeadAPIV2ObjectivesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Include != nil {

		// binding items for include
		joinedInclude := o.bindParamInclude(reg)

		// query array param include
		if err := r.SetQueryParam("include", joinedInclude...); err != nil {
			return err
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamHeadAPIV2Objectives binds the parameter exclude
func (o *HeadAPIV2ObjectivesParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamHeadAPIV2Objectives binds the parameter include
func (o *HeadAPIV2ObjectivesParams) bindParamInclude(formats strfmt.Registry) []string {
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
