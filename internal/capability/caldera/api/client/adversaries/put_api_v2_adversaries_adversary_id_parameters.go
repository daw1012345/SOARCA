// Code generated by go-swagger; DO NOT EDIT.

package adversaries

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

// NewPutAPIV2AdversariesAdversaryIDParams creates a new PutAPIV2AdversariesAdversaryIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAPIV2AdversariesAdversaryIDParams() *PutAPIV2AdversariesAdversaryIDParams {
	return &PutAPIV2AdversariesAdversaryIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV2AdversariesAdversaryIDParamsWithTimeout creates a new PutAPIV2AdversariesAdversaryIDParams object
// with the ability to set a timeout on a request.
func NewPutAPIV2AdversariesAdversaryIDParamsWithTimeout(timeout time.Duration) *PutAPIV2AdversariesAdversaryIDParams {
	return &PutAPIV2AdversariesAdversaryIDParams{
		timeout: timeout,
	}
}

// NewPutAPIV2AdversariesAdversaryIDParamsWithContext creates a new PutAPIV2AdversariesAdversaryIDParams object
// with the ability to set a context for a request.
func NewPutAPIV2AdversariesAdversaryIDParamsWithContext(ctx context.Context) *PutAPIV2AdversariesAdversaryIDParams {
	return &PutAPIV2AdversariesAdversaryIDParams{
		Context: ctx,
	}
}

// NewPutAPIV2AdversariesAdversaryIDParamsWithHTTPClient creates a new PutAPIV2AdversariesAdversaryIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutAPIV2AdversariesAdversaryIDParamsWithHTTPClient(client *http.Client) *PutAPIV2AdversariesAdversaryIDParams {
	return &PutAPIV2AdversariesAdversaryIDParams{
		HTTPClient: client,
	}
}

/*
PutAPIV2AdversariesAdversaryIDParams contains all the parameters to send to the API endpoint

	for the put API v2 adversaries adversary ID operation.

	Typically these are written to a http.Request.
*/
type PutAPIV2AdversariesAdversaryIDParams struct {

	/* AdversaryID.

	   UUID of the adversary to be created or updated
	*/
	AdversaryID string

	// Body.
	Body *models.PartialAdversary

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put API v2 adversaries adversary ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPIV2AdversariesAdversaryIDParams) WithDefaults() *PutAPIV2AdversariesAdversaryIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put API v2 adversaries adversary ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPIV2AdversariesAdversaryIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put API v2 adversaries adversary ID params
func (o *PutAPIV2AdversariesAdversaryIDParams) WithTimeout(timeout time.Duration) *PutAPIV2AdversariesAdversaryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v2 adversaries adversary ID params
func (o *PutAPIV2AdversariesAdversaryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v2 adversaries adversary ID params
func (o *PutAPIV2AdversariesAdversaryIDParams) WithContext(ctx context.Context) *PutAPIV2AdversariesAdversaryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v2 adversaries adversary ID params
func (o *PutAPIV2AdversariesAdversaryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v2 adversaries adversary ID params
func (o *PutAPIV2AdversariesAdversaryIDParams) WithHTTPClient(client *http.Client) *PutAPIV2AdversariesAdversaryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v2 adversaries adversary ID params
func (o *PutAPIV2AdversariesAdversaryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdversaryID adds the adversaryID to the put API v2 adversaries adversary ID params
func (o *PutAPIV2AdversariesAdversaryIDParams) WithAdversaryID(adversaryID string) *PutAPIV2AdversariesAdversaryIDParams {
	o.SetAdversaryID(adversaryID)
	return o
}

// SetAdversaryID adds the adversaryId to the put API v2 adversaries adversary ID params
func (o *PutAPIV2AdversariesAdversaryIDParams) SetAdversaryID(adversaryID string) {
	o.AdversaryID = adversaryID
}

// WithBody adds the body to the put API v2 adversaries adversary ID params
func (o *PutAPIV2AdversariesAdversaryIDParams) WithBody(body *models.PartialAdversary) *PutAPIV2AdversariesAdversaryIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put API v2 adversaries adversary ID params
func (o *PutAPIV2AdversariesAdversaryIDParams) SetBody(body *models.PartialAdversary) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV2AdversariesAdversaryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param adversary_id
	if err := r.SetPathParam("adversary_id", o.AdversaryID); err != nil {
		return err
	}
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
