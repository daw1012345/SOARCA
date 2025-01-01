// Code generated by go-swagger; DO NOT EDIT.

package operationsops

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

	"soarca/pkg/core/capability/caldera/api/models"
)

// NewPostAPIV2OperationsIDEventLogsParams creates a new PostAPIV2OperationsIDEventLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2OperationsIDEventLogsParams() *PostAPIV2OperationsIDEventLogsParams {
	return &PostAPIV2OperationsIDEventLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2OperationsIDEventLogsParamsWithTimeout creates a new PostAPIV2OperationsIDEventLogsParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2OperationsIDEventLogsParamsWithTimeout(timeout time.Duration) *PostAPIV2OperationsIDEventLogsParams {
	return &PostAPIV2OperationsIDEventLogsParams{
		timeout: timeout,
	}
}

// NewPostAPIV2OperationsIDEventLogsParamsWithContext creates a new PostAPIV2OperationsIDEventLogsParams object
// with the ability to set a context for a request.
func NewPostAPIV2OperationsIDEventLogsParamsWithContext(ctx context.Context) *PostAPIV2OperationsIDEventLogsParams {
	return &PostAPIV2OperationsIDEventLogsParams{
		Context: ctx,
	}
}

// NewPostAPIV2OperationsIDEventLogsParamsWithHTTPClient creates a new PostAPIV2OperationsIDEventLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2OperationsIDEventLogsParamsWithHTTPClient(client *http.Client) *PostAPIV2OperationsIDEventLogsParams {
	return &PostAPIV2OperationsIDEventLogsParams{
		HTTPClient: client,
	}
}

/*
PostAPIV2OperationsIDEventLogsParams contains all the parameters to send to the API endpoint

	for the post API v2 operations ID event logs operation.

	Typically these are written to a http.Request.
*/
type PostAPIV2OperationsIDEventLogsParams struct {

	// Body.
	Body *models.OperationOutputRequest

	// Exclude.
	Exclude []string

	// ID.
	ID string

	// Include.
	Include []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 operations ID event logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2OperationsIDEventLogsParams) WithDefaults() *PostAPIV2OperationsIDEventLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 operations ID event logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2OperationsIDEventLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) WithTimeout(timeout time.Duration) *PostAPIV2OperationsIDEventLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) WithContext(ctx context.Context) *PostAPIV2OperationsIDEventLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) WithHTTPClient(client *http.Client) *PostAPIV2OperationsIDEventLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) WithBody(body *models.OperationOutputRequest) *PostAPIV2OperationsIDEventLogsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) SetBody(body *models.OperationOutputRequest) {
	o.Body = body
}

// WithExclude adds the exclude to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) WithExclude(exclude []string) *PostAPIV2OperationsIDEventLogsParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithID adds the id to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) WithID(id string) *PostAPIV2OperationsIDEventLogsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) SetID(id string) {
	o.ID = id
}

// WithInclude adds the include to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) WithInclude(include []string) *PostAPIV2OperationsIDEventLogsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the post API v2 operations ID event logs params
func (o *PostAPIV2OperationsIDEventLogsParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2OperationsIDEventLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

// bindParamPostAPIV2OperationsIDEventLogs binds the parameter exclude
func (o *PostAPIV2OperationsIDEventLogsParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamPostAPIV2OperationsIDEventLogs binds the parameter include
func (o *PostAPIV2OperationsIDEventLogsParams) bindParamInclude(formats strfmt.Registry) []string {
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
