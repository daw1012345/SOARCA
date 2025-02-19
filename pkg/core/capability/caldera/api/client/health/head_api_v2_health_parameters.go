// Code generated by go-swagger; DO NOT EDIT.

package health

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

// NewHeadAPIV2HealthParams creates a new HeadAPIV2HealthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadAPIV2HealthParams() *HeadAPIV2HealthParams {
	return &HeadAPIV2HealthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadAPIV2HealthParamsWithTimeout creates a new HeadAPIV2HealthParams object
// with the ability to set a timeout on a request.
func NewHeadAPIV2HealthParamsWithTimeout(timeout time.Duration) *HeadAPIV2HealthParams {
	return &HeadAPIV2HealthParams{
		timeout: timeout,
	}
}

// NewHeadAPIV2HealthParamsWithContext creates a new HeadAPIV2HealthParams object
// with the ability to set a context for a request.
func NewHeadAPIV2HealthParamsWithContext(ctx context.Context) *HeadAPIV2HealthParams {
	return &HeadAPIV2HealthParams{
		Context: ctx,
	}
}

// NewHeadAPIV2HealthParamsWithHTTPClient creates a new HeadAPIV2HealthParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadAPIV2HealthParamsWithHTTPClient(client *http.Client) *HeadAPIV2HealthParams {
	return &HeadAPIV2HealthParams{
		HTTPClient: client,
	}
}

/*
HeadAPIV2HealthParams contains all the parameters to send to the API endpoint

	for the head API v2 health operation.

	Typically these are written to a http.Request.
*/
type HeadAPIV2HealthParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the head API v2 health params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2HealthParams) WithDefaults() *HeadAPIV2HealthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the head API v2 health params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2HealthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the head API v2 health params
func (o *HeadAPIV2HealthParams) WithTimeout(timeout time.Duration) *HeadAPIV2HealthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head API v2 health params
func (o *HeadAPIV2HealthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head API v2 health params
func (o *HeadAPIV2HealthParams) WithContext(ctx context.Context) *HeadAPIV2HealthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head API v2 health params
func (o *HeadAPIV2HealthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head API v2 health params
func (o *HeadAPIV2HealthParams) WithHTTPClient(client *http.Client) *HeadAPIV2HealthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head API v2 health params
func (o *HeadAPIV2HealthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HeadAPIV2HealthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
