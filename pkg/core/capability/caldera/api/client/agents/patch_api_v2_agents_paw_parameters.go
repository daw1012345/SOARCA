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

// NewPatchAPIV2AgentsPawParams creates a new PatchAPIV2AgentsPawParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIV2AgentsPawParams() *PatchAPIV2AgentsPawParams {
	return &PatchAPIV2AgentsPawParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIV2AgentsPawParamsWithTimeout creates a new PatchAPIV2AgentsPawParams object
// with the ability to set a timeout on a request.
func NewPatchAPIV2AgentsPawParamsWithTimeout(timeout time.Duration) *PatchAPIV2AgentsPawParams {
	return &PatchAPIV2AgentsPawParams{
		timeout: timeout,
	}
}

// NewPatchAPIV2AgentsPawParamsWithContext creates a new PatchAPIV2AgentsPawParams object
// with the ability to set a context for a request.
func NewPatchAPIV2AgentsPawParamsWithContext(ctx context.Context) *PatchAPIV2AgentsPawParams {
	return &PatchAPIV2AgentsPawParams{
		Context: ctx,
	}
}

// NewPatchAPIV2AgentsPawParamsWithHTTPClient creates a new PatchAPIV2AgentsPawParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIV2AgentsPawParamsWithHTTPClient(client *http.Client) *PatchAPIV2AgentsPawParams {
	return &PatchAPIV2AgentsPawParams{
		HTTPClient: client,
	}
}

/*
PatchAPIV2AgentsPawParams contains all the parameters to send to the API endpoint

	for the patch API v2 agents paw operation.

	Typically these are written to a http.Request.
*/
type PatchAPIV2AgentsPawParams struct {

	// Body.
	Body *models.PartialAgent1

	/* Paw.

	   ID of the Agent to update
	*/
	Paw string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API v2 agents paw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2AgentsPawParams) WithDefaults() *PatchAPIV2AgentsPawParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API v2 agents paw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2AgentsPawParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API v2 agents paw params
func (o *PatchAPIV2AgentsPawParams) WithTimeout(timeout time.Duration) *PatchAPIV2AgentsPawParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API v2 agents paw params
func (o *PatchAPIV2AgentsPawParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API v2 agents paw params
func (o *PatchAPIV2AgentsPawParams) WithContext(ctx context.Context) *PatchAPIV2AgentsPawParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API v2 agents paw params
func (o *PatchAPIV2AgentsPawParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API v2 agents paw params
func (o *PatchAPIV2AgentsPawParams) WithHTTPClient(client *http.Client) *PatchAPIV2AgentsPawParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API v2 agents paw params
func (o *PatchAPIV2AgentsPawParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch API v2 agents paw params
func (o *PatchAPIV2AgentsPawParams) WithBody(body *models.PartialAgent1) *PatchAPIV2AgentsPawParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch API v2 agents paw params
func (o *PatchAPIV2AgentsPawParams) SetBody(body *models.PartialAgent1) {
	o.Body = body
}

// WithPaw adds the paw to the patch API v2 agents paw params
func (o *PatchAPIV2AgentsPawParams) WithPaw(paw string) *PatchAPIV2AgentsPawParams {
	o.SetPaw(paw)
	return o
}

// SetPaw adds the paw to the patch API v2 agents paw params
func (o *PatchAPIV2AgentsPawParams) SetPaw(paw string) {
	o.Paw = paw
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIV2AgentsPawParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param paw
	if err := r.SetPathParam("paw", o.Paw); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
