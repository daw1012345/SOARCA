// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewPostAPIV2AgentsParams creates a new PostAPIV2AgentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2AgentsParams() *PostAPIV2AgentsParams {
	return &PostAPIV2AgentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2AgentsParamsWithTimeout creates a new PostAPIV2AgentsParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2AgentsParamsWithTimeout(timeout time.Duration) *PostAPIV2AgentsParams {
	return &PostAPIV2AgentsParams{
		timeout: timeout,
	}
}

// NewPostAPIV2AgentsParamsWithContext creates a new PostAPIV2AgentsParams object
// with the ability to set a context for a request.
func NewPostAPIV2AgentsParamsWithContext(ctx context.Context) *PostAPIV2AgentsParams {
	return &PostAPIV2AgentsParams{
		Context: ctx,
	}
}

// NewPostAPIV2AgentsParamsWithHTTPClient creates a new PostAPIV2AgentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2AgentsParamsWithHTTPClient(client *http.Client) *PostAPIV2AgentsParams {
	return &PostAPIV2AgentsParams{
		HTTPClient: client,
	}
}

/*
PostAPIV2AgentsParams contains all the parameters to send to the API endpoint

	for the post API v2 agents operation.

	Typically these are written to a http.Request.
*/
type PostAPIV2AgentsParams struct {

	// Body.
	Body *models.Agent

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 agents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2AgentsParams) WithDefaults() *PostAPIV2AgentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 agents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2AgentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 agents params
func (o *PostAPIV2AgentsParams) WithTimeout(timeout time.Duration) *PostAPIV2AgentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 agents params
func (o *PostAPIV2AgentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 agents params
func (o *PostAPIV2AgentsParams) WithContext(ctx context.Context) *PostAPIV2AgentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 agents params
func (o *PostAPIV2AgentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 agents params
func (o *PostAPIV2AgentsParams) WithHTTPClient(client *http.Client) *PostAPIV2AgentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 agents params
func (o *PostAPIV2AgentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API v2 agents params
func (o *PostAPIV2AgentsParams) WithBody(body *models.Agent) *PostAPIV2AgentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API v2 agents params
func (o *PostAPIV2AgentsParams) SetBody(body *models.Agent) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2AgentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
