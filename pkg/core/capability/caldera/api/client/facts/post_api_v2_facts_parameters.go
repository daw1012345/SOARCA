// Code generated by go-swagger; DO NOT EDIT.

package facts

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

// NewPostAPIV2FactsParams creates a new PostAPIV2FactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2FactsParams() *PostAPIV2FactsParams {
	return &PostAPIV2FactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2FactsParamsWithTimeout creates a new PostAPIV2FactsParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2FactsParamsWithTimeout(timeout time.Duration) *PostAPIV2FactsParams {
	return &PostAPIV2FactsParams{
		timeout: timeout,
	}
}

// NewPostAPIV2FactsParamsWithContext creates a new PostAPIV2FactsParams object
// with the ability to set a context for a request.
func NewPostAPIV2FactsParamsWithContext(ctx context.Context) *PostAPIV2FactsParams {
	return &PostAPIV2FactsParams{
		Context: ctx,
	}
}

// NewPostAPIV2FactsParamsWithHTTPClient creates a new PostAPIV2FactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2FactsParamsWithHTTPClient(client *http.Client) *PostAPIV2FactsParams {
	return &PostAPIV2FactsParams{
		HTTPClient: client,
	}
}

/*
PostAPIV2FactsParams contains all the parameters to send to the API endpoint

	for the post API v2 facts operation.

	Typically these are written to a http.Request.
*/
type PostAPIV2FactsParams struct {

	// Body.
	Body *models.Fact

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 facts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2FactsParams) WithDefaults() *PostAPIV2FactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 facts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2FactsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 facts params
func (o *PostAPIV2FactsParams) WithTimeout(timeout time.Duration) *PostAPIV2FactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 facts params
func (o *PostAPIV2FactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 facts params
func (o *PostAPIV2FactsParams) WithContext(ctx context.Context) *PostAPIV2FactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 facts params
func (o *PostAPIV2FactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 facts params
func (o *PostAPIV2FactsParams) WithHTTPClient(client *http.Client) *PostAPIV2FactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 facts params
func (o *PostAPIV2FactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API v2 facts params
func (o *PostAPIV2FactsParams) WithBody(body *models.Fact) *PostAPIV2FactsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API v2 facts params
func (o *PostAPIV2FactsParams) SetBody(body *models.Fact) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2FactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
