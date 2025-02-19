// Code generated by go-swagger; DO NOT EDIT.

package payloads

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

// NewPostAPIV2PayloadsParams creates a new PostAPIV2PayloadsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV2PayloadsParams() *PostAPIV2PayloadsParams {
	return &PostAPIV2PayloadsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV2PayloadsParamsWithTimeout creates a new PostAPIV2PayloadsParams object
// with the ability to set a timeout on a request.
func NewPostAPIV2PayloadsParamsWithTimeout(timeout time.Duration) *PostAPIV2PayloadsParams {
	return &PostAPIV2PayloadsParams{
		timeout: timeout,
	}
}

// NewPostAPIV2PayloadsParamsWithContext creates a new PostAPIV2PayloadsParams object
// with the ability to set a context for a request.
func NewPostAPIV2PayloadsParamsWithContext(ctx context.Context) *PostAPIV2PayloadsParams {
	return &PostAPIV2PayloadsParams{
		Context: ctx,
	}
}

// NewPostAPIV2PayloadsParamsWithHTTPClient creates a new PostAPIV2PayloadsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV2PayloadsParamsWithHTTPClient(client *http.Client) *PostAPIV2PayloadsParams {
	return &PostAPIV2PayloadsParams{
		HTTPClient: client,
	}
}

/*
PostAPIV2PayloadsParams contains all the parameters to send to the API endpoint

	for the post API v2 payloads operation.

	Typically these are written to a http.Request.
*/
type PostAPIV2PayloadsParams struct {

	// File.
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API v2 payloads params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2PayloadsParams) WithDefaults() *PostAPIV2PayloadsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v2 payloads params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV2PayloadsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v2 payloads params
func (o *PostAPIV2PayloadsParams) WithTimeout(timeout time.Duration) *PostAPIV2PayloadsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v2 payloads params
func (o *PostAPIV2PayloadsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v2 payloads params
func (o *PostAPIV2PayloadsParams) WithContext(ctx context.Context) *PostAPIV2PayloadsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v2 payloads params
func (o *PostAPIV2PayloadsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v2 payloads params
func (o *PostAPIV2PayloadsParams) WithHTTPClient(client *http.Client) *PostAPIV2PayloadsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v2 payloads params
func (o *PostAPIV2PayloadsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the post API v2 payloads params
func (o *PostAPIV2PayloadsParams) WithFile(file runtime.NamedReadCloser) *PostAPIV2PayloadsParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the post API v2 payloads params
func (o *PostAPIV2PayloadsParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV2PayloadsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
