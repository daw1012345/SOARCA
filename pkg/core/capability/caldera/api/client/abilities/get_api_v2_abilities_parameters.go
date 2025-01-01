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

// NewGetAPIV2AbilitiesParams creates a new GetAPIV2AbilitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2AbilitiesParams() *GetAPIV2AbilitiesParams {
	return &GetAPIV2AbilitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2AbilitiesParamsWithTimeout creates a new GetAPIV2AbilitiesParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2AbilitiesParamsWithTimeout(timeout time.Duration) *GetAPIV2AbilitiesParams {
	return &GetAPIV2AbilitiesParams{
		timeout: timeout,
	}
}

// NewGetAPIV2AbilitiesParamsWithContext creates a new GetAPIV2AbilitiesParams object
// with the ability to set a context for a request.
func NewGetAPIV2AbilitiesParamsWithContext(ctx context.Context) *GetAPIV2AbilitiesParams {
	return &GetAPIV2AbilitiesParams{
		Context: ctx,
	}
}

// NewGetAPIV2AbilitiesParamsWithHTTPClient creates a new GetAPIV2AbilitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2AbilitiesParamsWithHTTPClient(client *http.Client) *GetAPIV2AbilitiesParams {
	return &GetAPIV2AbilitiesParams{
		HTTPClient: client,
	}
}

/*
GetAPIV2AbilitiesParams contains all the parameters to send to the API endpoint

	for the get API v2 abilities operation.

	Typically these are written to a http.Request.
*/
type GetAPIV2AbilitiesParams struct {

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

// WithDefaults hydrates default values in the get API v2 abilities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2AbilitiesParams) WithDefaults() *GetAPIV2AbilitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 abilities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2AbilitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) WithTimeout(timeout time.Duration) *GetAPIV2AbilitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) WithContext(ctx context.Context) *GetAPIV2AbilitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) WithHTTPClient(client *http.Client) *GetAPIV2AbilitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) WithExclude(exclude []string) *GetAPIV2AbilitiesParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) WithInclude(include []string) *GetAPIV2AbilitiesParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) SetInclude(include []string) {
	o.Include = include
}

// WithSort adds the sort to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) WithSort(sort *string) *GetAPIV2AbilitiesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get API v2 abilities params
func (o *GetAPIV2AbilitiesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2AbilitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetAPIV2Abilities binds the parameter exclude
func (o *GetAPIV2AbilitiesParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamGetAPIV2Abilities binds the parameter include
func (o *GetAPIV2AbilitiesParams) bindParamInclude(formats strfmt.Registry) []string {
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
