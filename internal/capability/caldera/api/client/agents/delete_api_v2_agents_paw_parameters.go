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
)

// NewDeleteAPIV2AgentsPawParams creates a new DeleteAPIV2AgentsPawParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPIV2AgentsPawParams() *DeleteAPIV2AgentsPawParams {
	return &DeleteAPIV2AgentsPawParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV2AgentsPawParamsWithTimeout creates a new DeleteAPIV2AgentsPawParams object
// with the ability to set a timeout on a request.
func NewDeleteAPIV2AgentsPawParamsWithTimeout(timeout time.Duration) *DeleteAPIV2AgentsPawParams {
	return &DeleteAPIV2AgentsPawParams{
		timeout: timeout,
	}
}

// NewDeleteAPIV2AgentsPawParamsWithContext creates a new DeleteAPIV2AgentsPawParams object
// with the ability to set a context for a request.
func NewDeleteAPIV2AgentsPawParamsWithContext(ctx context.Context) *DeleteAPIV2AgentsPawParams {
	return &DeleteAPIV2AgentsPawParams{
		Context: ctx,
	}
}

// NewDeleteAPIV2AgentsPawParamsWithHTTPClient creates a new DeleteAPIV2AgentsPawParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPIV2AgentsPawParamsWithHTTPClient(client *http.Client) *DeleteAPIV2AgentsPawParams {
	return &DeleteAPIV2AgentsPawParams{
		HTTPClient: client,
	}
}

/*
DeleteAPIV2AgentsPawParams contains all the parameters to send to the API endpoint

	for the delete API v2 agents paw operation.

	Typically these are written to a http.Request.
*/
type DeleteAPIV2AgentsPawParams struct {

	/* Paw.

	   paw of the Agent to be deleted
	*/
	Paw string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API v2 agents paw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIV2AgentsPawParams) WithDefaults() *DeleteAPIV2AgentsPawParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API v2 agents paw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIV2AgentsPawParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API v2 agents paw params
func (o *DeleteAPIV2AgentsPawParams) WithTimeout(timeout time.Duration) *DeleteAPIV2AgentsPawParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API v2 agents paw params
func (o *DeleteAPIV2AgentsPawParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API v2 agents paw params
func (o *DeleteAPIV2AgentsPawParams) WithContext(ctx context.Context) *DeleteAPIV2AgentsPawParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API v2 agents paw params
func (o *DeleteAPIV2AgentsPawParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API v2 agents paw params
func (o *DeleteAPIV2AgentsPawParams) WithHTTPClient(client *http.Client) *DeleteAPIV2AgentsPawParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API v2 agents paw params
func (o *DeleteAPIV2AgentsPawParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaw adds the paw to the delete API v2 agents paw params
func (o *DeleteAPIV2AgentsPawParams) WithPaw(paw string) *DeleteAPIV2AgentsPawParams {
	o.SetPaw(paw)
	return o
}

// SetPaw adds the paw to the delete API v2 agents paw params
func (o *DeleteAPIV2AgentsPawParams) SetPaw(paw string) {
	o.Paw = paw
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV2AgentsPawParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param paw
	if err := r.SetPathParam("paw", o.Paw); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
