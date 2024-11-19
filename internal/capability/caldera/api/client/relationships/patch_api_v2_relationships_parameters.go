// Code generated by go-swagger; DO NOT EDIT.

package relationships

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

	"soarca/internal/capability/caldera/api/models"
)

// NewPatchAPIV2RelationshipsParams creates a new PatchAPIV2RelationshipsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIV2RelationshipsParams() *PatchAPIV2RelationshipsParams {
	return &PatchAPIV2RelationshipsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIV2RelationshipsParamsWithTimeout creates a new PatchAPIV2RelationshipsParams object
// with the ability to set a timeout on a request.
func NewPatchAPIV2RelationshipsParamsWithTimeout(timeout time.Duration) *PatchAPIV2RelationshipsParams {
	return &PatchAPIV2RelationshipsParams{
		timeout: timeout,
	}
}

// NewPatchAPIV2RelationshipsParamsWithContext creates a new PatchAPIV2RelationshipsParams object
// with the ability to set a context for a request.
func NewPatchAPIV2RelationshipsParamsWithContext(ctx context.Context) *PatchAPIV2RelationshipsParams {
	return &PatchAPIV2RelationshipsParams{
		Context: ctx,
	}
}

// NewPatchAPIV2RelationshipsParamsWithHTTPClient creates a new PatchAPIV2RelationshipsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIV2RelationshipsParamsWithHTTPClient(client *http.Client) *PatchAPIV2RelationshipsParams {
	return &PatchAPIV2RelationshipsParams{
		HTTPClient: client,
	}
}

/*
PatchAPIV2RelationshipsParams contains all the parameters to send to the API endpoint

	for the patch API v2 relationships operation.

	Typically these are written to a http.Request.
*/
type PatchAPIV2RelationshipsParams struct {

	// Body.
	Body *models.PartialRelationshipUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API v2 relationships params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2RelationshipsParams) WithDefaults() *PatchAPIV2RelationshipsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API v2 relationships params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIV2RelationshipsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API v2 relationships params
func (o *PatchAPIV2RelationshipsParams) WithTimeout(timeout time.Duration) *PatchAPIV2RelationshipsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API v2 relationships params
func (o *PatchAPIV2RelationshipsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API v2 relationships params
func (o *PatchAPIV2RelationshipsParams) WithContext(ctx context.Context) *PatchAPIV2RelationshipsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API v2 relationships params
func (o *PatchAPIV2RelationshipsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API v2 relationships params
func (o *PatchAPIV2RelationshipsParams) WithHTTPClient(client *http.Client) *PatchAPIV2RelationshipsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API v2 relationships params
func (o *PatchAPIV2RelationshipsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch API v2 relationships params
func (o *PatchAPIV2RelationshipsParams) WithBody(body *models.PartialRelationshipUpdate) *PatchAPIV2RelationshipsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch API v2 relationships params
func (o *PatchAPIV2RelationshipsParams) SetBody(body *models.PartialRelationshipUpdate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIV2RelationshipsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
