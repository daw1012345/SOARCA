// Code generated by go-swagger; DO NOT EDIT.

package sources

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

// NewHeadAPIV2SourcesParams creates a new HeadAPIV2SourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadAPIV2SourcesParams() *HeadAPIV2SourcesParams {
	return &HeadAPIV2SourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadAPIV2SourcesParamsWithTimeout creates a new HeadAPIV2SourcesParams object
// with the ability to set a timeout on a request.
func NewHeadAPIV2SourcesParamsWithTimeout(timeout time.Duration) *HeadAPIV2SourcesParams {
	return &HeadAPIV2SourcesParams{
		timeout: timeout,
	}
}

// NewHeadAPIV2SourcesParamsWithContext creates a new HeadAPIV2SourcesParams object
// with the ability to set a context for a request.
func NewHeadAPIV2SourcesParamsWithContext(ctx context.Context) *HeadAPIV2SourcesParams {
	return &HeadAPIV2SourcesParams{
		Context: ctx,
	}
}

// NewHeadAPIV2SourcesParamsWithHTTPClient creates a new HeadAPIV2SourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadAPIV2SourcesParamsWithHTTPClient(client *http.Client) *HeadAPIV2SourcesParams {
	return &HeadAPIV2SourcesParams{
		HTTPClient: client,
	}
}

/*
HeadAPIV2SourcesParams contains all the parameters to send to the API endpoint

	for the head API v2 sources operation.

	Typically these are written to a http.Request.
*/
type HeadAPIV2SourcesParams struct {

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

// WithDefaults hydrates default values in the head API v2 sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2SourcesParams) WithDefaults() *HeadAPIV2SourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the head API v2 sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2SourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) WithTimeout(timeout time.Duration) *HeadAPIV2SourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) WithContext(ctx context.Context) *HeadAPIV2SourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) WithHTTPClient(client *http.Client) *HeadAPIV2SourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) WithExclude(exclude []string) *HeadAPIV2SourcesParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) WithInclude(include []string) *HeadAPIV2SourcesParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) SetInclude(include []string) {
	o.Include = include
}

// WithSort adds the sort to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) WithSort(sort *string) *HeadAPIV2SourcesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the head API v2 sources params
func (o *HeadAPIV2SourcesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *HeadAPIV2SourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamHeadAPIV2Sources binds the parameter exclude
func (o *HeadAPIV2SourcesParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamHeadAPIV2Sources binds the parameter include
func (o *HeadAPIV2SourcesParams) bindParamInclude(formats strfmt.Registry) []string {
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
