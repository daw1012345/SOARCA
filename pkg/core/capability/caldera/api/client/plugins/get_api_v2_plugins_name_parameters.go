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

// NewGetAPIV2PluginsNameParams creates a new GetAPIV2PluginsNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2PluginsNameParams() *GetAPIV2PluginsNameParams {
	return &GetAPIV2PluginsNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2PluginsNameParamsWithTimeout creates a new GetAPIV2PluginsNameParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2PluginsNameParamsWithTimeout(timeout time.Duration) *GetAPIV2PluginsNameParams {
	return &GetAPIV2PluginsNameParams{
		timeout: timeout,
	}
}

// NewGetAPIV2PluginsNameParamsWithContext creates a new GetAPIV2PluginsNameParams object
// with the ability to set a context for a request.
func NewGetAPIV2PluginsNameParamsWithContext(ctx context.Context) *GetAPIV2PluginsNameParams {
	return &GetAPIV2PluginsNameParams{
		Context: ctx,
	}
}

// NewGetAPIV2PluginsNameParamsWithHTTPClient creates a new GetAPIV2PluginsNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2PluginsNameParamsWithHTTPClient(client *http.Client) *GetAPIV2PluginsNameParams {
	return &GetAPIV2PluginsNameParams{
		HTTPClient: client,
	}
}

/*
GetAPIV2PluginsNameParams contains all the parameters to send to the API endpoint

	for the get API v2 plugins name operation.

	Typically these are written to a http.Request.
*/
type GetAPIV2PluginsNameParams struct {

	// Exclude.
	Exclude []string

	// Include.
	Include []string

	/* Name.

	   The name of the plugin
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 plugins name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2PluginsNameParams) WithDefaults() *GetAPIV2PluginsNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 plugins name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2PluginsNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) WithTimeout(timeout time.Duration) *GetAPIV2PluginsNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) WithContext(ctx context.Context) *GetAPIV2PluginsNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) WithHTTPClient(client *http.Client) *GetAPIV2PluginsNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) WithExclude(exclude []string) *GetAPIV2PluginsNameParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) WithInclude(include []string) *GetAPIV2PluginsNameParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) SetInclude(include []string) {
	o.Include = include
}

// WithName adds the name to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) WithName(name string) *GetAPIV2PluginsNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API v2 plugins name params
func (o *GetAPIV2PluginsNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2PluginsNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAPIV2PluginsName binds the parameter exclude
func (o *GetAPIV2PluginsNameParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamGetAPIV2PluginsName binds the parameter include
func (o *GetAPIV2PluginsNameParams) bindParamInclude(formats strfmt.Registry) []string {
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
