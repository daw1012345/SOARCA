// Code generated by go-swagger; DO NOT EDIT.

package operationsops

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

// NewPostAPIV2OperationsIDPotentialLinksParams creates a new PostAPIV2OperationsIDPotentialLinksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2OperationsIDPotentialLinksParams() *PostAPIV2OperationsIDPotentialLinksParams {
	return &PostAPIV2OperationsIDPotentialLinksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2OperationsIDPotentialLinksParamsWithTimeout creates a new PostAPIV2OperationsIDPotentialLinksParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2OperationsIDPotentialLinksParamsWithTimeout(timeout time.Duration) *PostAPIV2OperationsIDPotentialLinksParams {
	return &PostAPIV2OperationsIDPotentialLinksParams{
		timeout: timeout,
	}
}

// NewPostAPIV2OperationsIDPotentialLinksParamsWithContext creates a new PostAPIV2OperationsIDPotentialLinksParams object
// with the ability to set a context for a request.
func NewPostAPIV2OperationsIDPotentialLinksParamsWithContext(ctx context.Context) *PostAPIV2OperationsIDPotentialLinksParams {
	return &PostAPIV2OperationsIDPotentialLinksParams{
		Context: ctx,
	}
}

// NewPostAPIV2OperationsIDPotentialLinksParamsWithHTTPClient creates a new PostAPIV2OperationsIDPotentialLinksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2OperationsIDPotentialLinksParamsWithHTTPClient(client *http.Client) *PostAPIV2OperationsIDPotentialLinksParams {
	return &PostAPIV2OperationsIDPotentialLinksParams{
		HTTPClient: client,
	}
}

/*
PostAPIV2OperationsIDPotentialLinksParams contains all the parameters to send to the API endpoint

	for the post API v2 operations ID potential links operation.

	Typically these are written to a http.Request.
*/
type PostAPIV2OperationsIDPotentialLinksParams struct {

	// Body.
	Body *models.Link

	/* ID.

	   UUID of the operation object for the link to be created on.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 operations ID potential links params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2OperationsIDPotentialLinksParams) WithDefaults() *PostAPIV2OperationsIDPotentialLinksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 operations ID potential links params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2OperationsIDPotentialLinksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 operations ID potential links params
func (o *PostAPIV2OperationsIDPotentialLinksParams) WithTimeout(timeout time.Duration) *PostAPIV2OperationsIDPotentialLinksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 operations ID potential links params
func (o *PostAPIV2OperationsIDPotentialLinksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 operations ID potential links params
func (o *PostAPIV2OperationsIDPotentialLinksParams) WithContext(ctx context.Context) *PostAPIV2OperationsIDPotentialLinksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 operations ID potential links params
func (o *PostAPIV2OperationsIDPotentialLinksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 operations ID potential links params
func (o *PostAPIV2OperationsIDPotentialLinksParams) WithHTTPClient(client *http.Client) *PostAPIV2OperationsIDPotentialLinksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 operations ID potential links params
func (o *PostAPIV2OperationsIDPotentialLinksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API v2 operations ID potential links params
func (o *PostAPIV2OperationsIDPotentialLinksParams) WithBody(body *models.Link) *PostAPIV2OperationsIDPotentialLinksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API v2 operations ID potential links params
func (o *PostAPIV2OperationsIDPotentialLinksParams) SetBody(body *models.Link) {
	o.Body = body
}

// WithID adds the id to the post API v2 operations ID potential links params
func (o *PostAPIV2OperationsIDPotentialLinksParams) WithID(id string) *PostAPIV2OperationsIDPotentialLinksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post API v2 operations ID potential links params
func (o *PostAPIV2OperationsIDPotentialLinksParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2OperationsIDPotentialLinksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
