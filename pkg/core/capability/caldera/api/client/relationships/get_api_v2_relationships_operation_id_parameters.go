// Code generated by go-swagger; DO NOT EDIT.

package relationships

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

// NewGetAPIV2RelationshipsOperationIDParams creates a new GetAPIV2RelationshipsOperationIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2RelationshipsOperationIDParams() *GetAPIV2RelationshipsOperationIDParams {
	return &GetAPIV2RelationshipsOperationIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2RelationshipsOperationIDParamsWithTimeout creates a new GetAPIV2RelationshipsOperationIDParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2RelationshipsOperationIDParamsWithTimeout(timeout time.Duration) *GetAPIV2RelationshipsOperationIDParams {
	return &GetAPIV2RelationshipsOperationIDParams{
		timeout: timeout,
	}
}

// NewGetAPIV2RelationshipsOperationIDParamsWithContext creates a new GetAPIV2RelationshipsOperationIDParams object
// with the ability to set a context for a request.
func NewGetAPIV2RelationshipsOperationIDParamsWithContext(ctx context.Context) *GetAPIV2RelationshipsOperationIDParams {
	return &GetAPIV2RelationshipsOperationIDParams{
		Context: ctx,
	}
}

// NewGetAPIV2RelationshipsOperationIDParamsWithHTTPClient creates a new GetAPIV2RelationshipsOperationIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2RelationshipsOperationIDParamsWithHTTPClient(client *http.Client) *GetAPIV2RelationshipsOperationIDParams {
	return &GetAPIV2RelationshipsOperationIDParams{
		HTTPClient: client,
	}
}

/*
GetAPIV2RelationshipsOperationIDParams contains all the parameters to send to the API endpoint

	for the get API v2 relationships operation ID operation.

	Typically these are written to a http.Request.
*/
type GetAPIV2RelationshipsOperationIDParams struct {

	// Exclude.
	Exclude []string

	// Include.
	Include []string

	// OperationID.
	OperationID string

	// Sort.
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 relationships operation ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2RelationshipsOperationIDParams) WithDefaults() *GetAPIV2RelationshipsOperationIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 relationships operation ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2RelationshipsOperationIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) WithTimeout(timeout time.Duration) *GetAPIV2RelationshipsOperationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) WithContext(ctx context.Context) *GetAPIV2RelationshipsOperationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) WithHTTPClient(client *http.Client) *GetAPIV2RelationshipsOperationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) WithExclude(exclude []string) *GetAPIV2RelationshipsOperationIDParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) WithInclude(include []string) *GetAPIV2RelationshipsOperationIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) SetInclude(include []string) {
	o.Include = include
}

// WithOperationID adds the operationID to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) WithOperationID(operationID string) *GetAPIV2RelationshipsOperationIDParams {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) SetOperationID(operationID string) {
	o.OperationID = operationID
}

// WithSort adds the sort to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) WithSort(sort *string) *GetAPIV2RelationshipsOperationIDParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get API v2 relationships operation ID params
func (o *GetAPIV2RelationshipsOperationIDParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2RelationshipsOperationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param operation_id
	if err := r.SetPathParam("operation_id", o.OperationID); err != nil {
		return err
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

// bindParamGetAPIV2RelationshipsOperationID binds the parameter exclude
func (o *GetAPIV2RelationshipsOperationIDParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamGetAPIV2RelationshipsOperationID binds the parameter include
func (o *GetAPIV2RelationshipsOperationIDParams) bindParamInclude(formats strfmt.Registry) []string {
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
