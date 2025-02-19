// Code generated by go-swagger; DO NOT EDIT.

package contacts

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
)

// NewGetAPIV2ContactsParams creates a new GetAPIV2ContactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIV2ContactsParams() *GetAPIV2ContactsParams {
	return &GetAPIV2ContactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV2ContactsParamsWithTimeout creates a new GetAPIV2ContactsParams object
// with the ability to set a timeout on a request.
func NewGetAPIV2ContactsParamsWithTimeout(timeout time.Duration) *GetAPIV2ContactsParams {
	return &GetAPIV2ContactsParams{
		timeout: timeout,
	}
}

// NewGetAPIV2ContactsParamsWithContext creates a new GetAPIV2ContactsParams object
// with the ability to set a context for a request.
func NewGetAPIV2ContactsParamsWithContext(ctx context.Context) *GetAPIV2ContactsParams {
	return &GetAPIV2ContactsParams{
		Context: ctx,
	}
}

// NewGetAPIV2ContactsParamsWithHTTPClient creates a new GetAPIV2ContactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIV2ContactsParamsWithHTTPClient(client *http.Client) *GetAPIV2ContactsParams {
	return &GetAPIV2ContactsParams{
		HTTPClient: client,
	}
}

/*
GetAPIV2ContactsParams contains all the parameters to send to the API endpoint

	for the get API v2 contacts operation.

	Typically these are written to a http.Request.
*/
type GetAPIV2ContactsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API v2 contacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ContactsParams) WithDefaults() *GetAPIV2ContactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API v2 contacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIV2ContactsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API v2 contacts params
func (o *GetAPIV2ContactsParams) WithTimeout(timeout time.Duration) *GetAPIV2ContactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v2 contacts params
func (o *GetAPIV2ContactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v2 contacts params
func (o *GetAPIV2ContactsParams) WithContext(ctx context.Context) *GetAPIV2ContactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v2 contacts params
func (o *GetAPIV2ContactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v2 contacts params
func (o *GetAPIV2ContactsParams) WithHTTPClient(client *http.Client) *GetAPIV2ContactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v2 contacts params
func (o *GetAPIV2ContactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV2ContactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
