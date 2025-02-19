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

	"soarca/pkg/core/capability/caldera/api/models"
)

// NewPostAPIV2ObjectivesParams creates a new PostAPIV2ObjectivesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2ObjectivesParams() *PostAPIV2ObjectivesParams {
	return &PostAPIV2ObjectivesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2ObjectivesParamsWithTimeout creates a new PostAPIV2ObjectivesParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2ObjectivesParamsWithTimeout(timeout time.Duration) *PostAPIV2ObjectivesParams {
	return &PostAPIV2ObjectivesParams{
		timeout: timeout,
	}
}

// NewPostAPIV2ObjectivesParamsWithContext creates a new PostAPIV2ObjectivesParams object
// with the ability to set a context for a request.
func NewPostAPIV2ObjectivesParamsWithContext(ctx context.Context) *PostAPIV2ObjectivesParams {
	return &PostAPIV2ObjectivesParams{
		Context: ctx,
	}
}

// NewPostAPIV2ObjectivesParamsWithHTTPClient creates a new PostAPIV2ObjectivesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2ObjectivesParamsWithHTTPClient(client *http.Client) *PostAPIV2ObjectivesParams {
	return &PostAPIV2ObjectivesParams{
		HTTPClient: client,
	}
}

/*
PostAPIV2ObjectivesParams contains all the parameters to send to the API endpoint

	for the post API v2 objectives operation.

	Typically these are written to a http.Request.
*/
type PostAPIV2ObjectivesParams struct {

	// Body.
	Body *models.Objective

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 objectives params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ObjectivesParams) WithDefaults() *PostAPIV2ObjectivesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 objectives params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2ObjectivesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 objectives params
func (o *PostAPIV2ObjectivesParams) WithTimeout(timeout time.Duration) *PostAPIV2ObjectivesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 objectives params
func (o *PostAPIV2ObjectivesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 objectives params
func (o *PostAPIV2ObjectivesParams) WithContext(ctx context.Context) *PostAPIV2ObjectivesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 objectives params
func (o *PostAPIV2ObjectivesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 objectives params
func (o *PostAPIV2ObjectivesParams) WithHTTPClient(client *http.Client) *PostAPIV2ObjectivesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 objectives params
func (o *PostAPIV2ObjectivesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API v2 objectives params
func (o *PostAPIV2ObjectivesParams) WithBody(body *models.Objective) *PostAPIV2ObjectivesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API v2 objectives params
func (o *PostAPIV2ObjectivesParams) SetBody(body *models.Objective) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2ObjectivesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
