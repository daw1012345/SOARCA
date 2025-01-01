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

// NewPatchAPIV2FactsParams creates a new PatchAPIV2FactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIV2FactsParams() *PatchAPIV2FactsParams {
	return &PatchAPIV2FactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIV2FactsParamsWithTimeout creates a new PatchAPIV2FactsParams object
// with the ability to set a timeout on a request.
func NewPatchAPIV2FactsParamsWithTimeout(timeout time.Duration) *PatchAPIV2FactsParams {
	return &PatchAPIV2FactsParams{
		timeout: timeout,
	}
}

// NewPatchAPIV2FactsParamsWithContext creates a new PatchAPIV2FactsParams object
// with the ability to set a context for a request.
func NewPatchAPIV2FactsParamsWithContext(ctx context.Context) *PatchAPIV2FactsParams {
	return &PatchAPIV2FactsParams{
		Context: ctx,
	}
}

// NewPatchAPIV2FactsParamsWithHTTPClient creates a new PatchAPIV2FactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIV2FactsParamsWithHTTPClient(client *http.Client) *PatchAPIV2FactsParams {
	return &PatchAPIV2FactsParams{
		HTTPClient: client,
	}
}

/*
PatchAPIV2FactsParams contains all the parameters to send to the API endpoint

	for the patch API v2 facts operation.

	Typically these are written to a http.Request.
*/
type PatchAPIV2FactsParams struct {

	// Body.
	Body *models.PartialFactUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API v2 facts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2FactsParams) WithDefaults() *PatchAPIV2FactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API v2 facts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2FactsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API v2 facts params
func (o *PatchAPIV2FactsParams) WithTimeout(timeout time.Duration) *PatchAPIV2FactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API v2 facts params
func (o *PatchAPIV2FactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API v2 facts params
func (o *PatchAPIV2FactsParams) WithContext(ctx context.Context) *PatchAPIV2FactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API v2 facts params
func (o *PatchAPIV2FactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API v2 facts params
func (o *PatchAPIV2FactsParams) WithHTTPClient(client *http.Client) *PatchAPIV2FactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API v2 facts params
func (o *PatchAPIV2FactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch API v2 facts params
func (o *PatchAPIV2FactsParams) WithBody(body *models.PartialFactUpdateRequest) *PatchAPIV2FactsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch API v2 facts params
func (o *PatchAPIV2FactsParams) SetBody(body *models.PartialFactUpdateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIV2FactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
