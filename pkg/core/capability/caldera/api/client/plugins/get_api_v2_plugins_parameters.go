// Code generated by go-swagger; DO NOT EDIT.

package plugins

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

// NewGetAPIV2PluginsParams creates a new GetAPIV2PluginsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2PluginsParams() *GetAPIV2PluginsParams {
	return &GetAPIV2PluginsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2PluginsParamsWithTimeout creates a new GetAPIV2PluginsParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2PluginsParamsWithTimeout(timeout time.Duration) *GetAPIV2PluginsParams {
	return &GetAPIV2PluginsParams{
		timeout: timeout,
	}
}

// NewGetAPIV2PluginsParamsWithContext creates a new GetAPIV2PluginsParams object
// with the ability to set a context for a request.
func NewGetAPIV2PluginsParamsWithContext(ctx context.Context) *GetAPIV2PluginsParams {
	return &GetAPIV2PluginsParams{
		Context: ctx,
	}
}

// NewGetAPIV2PluginsParamsWithHTTPClient creates a new GetAPIV2PluginsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2PluginsParamsWithHTTPClient(client *http.Client) *GetAPIV2PluginsParams {
	return &GetAPIV2PluginsParams{
		HTTPClient: client,
	}
}

/*
GetAPIV2PluginsParams contains all the parameters to send to the API endpoint

	for the get API v2 plugins operation.

	Typically these are written to a http.Request.
*/
type GetAPIV2PluginsParams struct {

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

// WithDefaults hydrates default values in the get API v2 plugins params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2PluginsParams) WithDefaults() *GetAPIV2PluginsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 plugins params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2PluginsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) WithTimeout(timeout time.Duration) *GetAPIV2PluginsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) WithContext(ctx context.Context) *GetAPIV2PluginsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) WithHTTPClient(client *http.Client) *GetAPIV2PluginsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) WithExclude(exclude []string) *GetAPIV2PluginsParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) WithInclude(include []string) *GetAPIV2PluginsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) SetInclude(include []string) {
	o.Include = include
}

// WithSort adds the sort to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) WithSort(sort *string) *GetAPIV2PluginsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get API v2 plugins params
func (o *GetAPIV2PluginsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2PluginsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetAPIV2Plugins binds the parameter exclude
func (o *GetAPIV2PluginsParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamGetAPIV2Plugins binds the parameter include
func (o *GetAPIV2PluginsParams) bindParamInclude(formats strfmt.Registry) []string {
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
