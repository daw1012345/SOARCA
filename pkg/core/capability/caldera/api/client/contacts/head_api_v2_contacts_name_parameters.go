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

// NewHeadAPIV2ContactsNameParams creates a new HeadAPIV2ContactsNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadAPIV2ContactsNameParams() *HeadAPIV2ContactsNameParams {
	return &HeadAPIV2ContactsNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadAPIV2ContactsNameParamsWithTimeout creates a new HeadAPIV2ContactsNameParams object
// with the ability to set a timeout on a request.
func NewHeadAPIV2ContactsNameParamsWithTimeout(timeout time.Duration) *HeadAPIV2ContactsNameParams {
	return &HeadAPIV2ContactsNameParams{
		timeout: timeout,
	}
}

// NewHeadAPIV2ContactsNameParamsWithContext creates a new HeadAPIV2ContactsNameParams object
// with the ability to set a context for a request.
func NewHeadAPIV2ContactsNameParamsWithContext(ctx context.Context) *HeadAPIV2ContactsNameParams {
	return &HeadAPIV2ContactsNameParams{
		Context: ctx,
	}
}

// NewHeadAPIV2ContactsNameParamsWithHTTPClient creates a new HeadAPIV2ContactsNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadAPIV2ContactsNameParamsWithHTTPClient(client *http.Client) *HeadAPIV2ContactsNameParams {
	return &HeadAPIV2ContactsNameParams{
		HTTPClient: client,
	}
}

/*
HeadAPIV2ContactsNameParams contains all the parameters to send to the API endpoint

	for the head API v2 contacts name operation.

	Typically these are written to a http.Request.
*/
type HeadAPIV2ContactsNameParams struct {

	/* Name.

	   Name of the contact to get beacons for, e.g. HTTP, TCP, et cetera.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the head API v2 contacts name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2ContactsNameParams) WithDefaults() *HeadAPIV2ContactsNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the head API v2 contacts name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2ContactsNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the head API v2 contacts name params
func (o *HeadAPIV2ContactsNameParams) WithTimeout(timeout time.Duration) *HeadAPIV2ContactsNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head API v2 contacts name params
func (o *HeadAPIV2ContactsNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head API v2 contacts name params
func (o *HeadAPIV2ContactsNameParams) WithContext(ctx context.Context) *HeadAPIV2ContactsNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head API v2 contacts name params
func (o *HeadAPIV2ContactsNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head API v2 contacts name params
func (o *HeadAPIV2ContactsNameParams) WithHTTPClient(client *http.Client) *HeadAPIV2ContactsNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head API v2 contacts name params
func (o *HeadAPIV2ContactsNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the head API v2 contacts name params
func (o *HeadAPIV2ContactsNameParams) WithName(name string) *HeadAPIV2ContactsNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the head API v2 contacts name params
func (o *HeadAPIV2ContactsNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *HeadAPIV2ContactsNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
