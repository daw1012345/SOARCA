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

// NewHeadAPIV2DeployCommandsParams creates a new HeadAPIV2DeployCommandsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadAPIV2DeployCommandsParams() *HeadAPIV2DeployCommandsParams {
	return &HeadAPIV2DeployCommandsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadAPIV2DeployCommandsParamsWithTimeout creates a new HeadAPIV2DeployCommandsParams object
// with the ability to set a timeout on a request.
func NewHeadAPIV2DeployCommandsParamsWithTimeout(timeout time.Duration) *HeadAPIV2DeployCommandsParams {
	return &HeadAPIV2DeployCommandsParams{
		timeout: timeout,
	}
}

// NewHeadAPIV2DeployCommandsParamsWithContext creates a new HeadAPIV2DeployCommandsParams object
// with the ability to set a context for a request.
func NewHeadAPIV2DeployCommandsParamsWithContext(ctx context.Context) *HeadAPIV2DeployCommandsParams {
	return &HeadAPIV2DeployCommandsParams{
		Context: ctx,
	}
}

// NewHeadAPIV2DeployCommandsParamsWithHTTPClient creates a new HeadAPIV2DeployCommandsParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadAPIV2DeployCommandsParamsWithHTTPClient(client *http.Client) *HeadAPIV2DeployCommandsParams {
	return &HeadAPIV2DeployCommandsParams{
		HTTPClient: client,
	}
}

/*
HeadAPIV2DeployCommandsParams contains all the parameters to send to the API endpoint

	for the head API v2 deploy commands operation.

	Typically these are written to a http.Request.
*/
type HeadAPIV2DeployCommandsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the head API v2 deploy commands params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2DeployCommandsParams) WithDefaults() *HeadAPIV2DeployCommandsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the head API v2 deploy commands params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadAPIV2DeployCommandsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the head API v2 deploy commands params
func (o *HeadAPIV2DeployCommandsParams) WithTimeout(timeout time.Duration) *HeadAPIV2DeployCommandsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head API v2 deploy commands params
func (o *HeadAPIV2DeployCommandsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head API v2 deploy commands params
func (o *HeadAPIV2DeployCommandsParams) WithContext(ctx context.Context) *HeadAPIV2DeployCommandsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head API v2 deploy commands params
func (o *HeadAPIV2DeployCommandsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head API v2 deploy commands params
func (o *HeadAPIV2DeployCommandsParams) WithHTTPClient(client *http.Client) *HeadAPIV2DeployCommandsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head API v2 deploy commands params
func (o *HeadAPIV2DeployCommandsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HeadAPIV2DeployCommandsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
