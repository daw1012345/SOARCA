// Code generated by go-swagger; DO NOT EDIT.

package sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new sources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new sources API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new sources API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for sources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIV2SourcesID(params *DeleteAPIV2SourcesIDParams, opts ...ClientOption) (*DeleteAPIV2SourcesIDOK, error)

	GetAPIV2Sources(params *GetAPIV2SourcesParams, opts ...ClientOption) (*GetAPIV2SourcesOK, error)

	GetAPIV2SourcesID(params *GetAPIV2SourcesIDParams, opts ...ClientOption) (*GetAPIV2SourcesIDOK, error)

	HeadAPIV2Sources(params *HeadAPIV2SourcesParams, opts ...ClientOption) (*HeadAPIV2SourcesOK, error)

	HeadAPIV2SourcesID(params *HeadAPIV2SourcesIDParams, opts ...ClientOption) (*HeadAPIV2SourcesIDOK, error)

	PatchAPIV2SourcesID(params *PatchAPIV2SourcesIDParams, opts ...ClientOption) (*PatchAPIV2SourcesIDOK, error)

	PostAPIV2Sources(params *PostAPIV2SourcesParams, opts ...ClientOption) (*PostAPIV2SourcesOK, error)

	PutAPIV2SourcesID(params *PutAPIV2SourcesIDParams, opts ...ClientOption) (*PutAPIV2SourcesIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAPIV2SourcesID deletes an existing fact source

Delete a Fact Source, given its id.
*/
func (a *Client) DeleteAPIV2SourcesID(params *DeleteAPIV2SourcesIDParams, opts ...ClientOption) (*DeleteAPIV2SourcesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIV2SourcesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIV2SourcesID",
		Method:             "DELETE",
		PathPattern:        "/api/v2/sources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAPIV2SourcesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAPIV2SourcesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIV2SourcesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV2Sources retrieves all fact sources

Returns a list of all Fact Sources, including custom-created ones.
*/
func (a *Client) GetAPIV2Sources(params *GetAPIV2SourcesParams, opts ...ClientOption) (*GetAPIV2SourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2SourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2Sources",
		Method:             "GET",
		PathPattern:        "/api/v2/sources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV2SourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIV2SourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2Sources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV2SourcesID retrieves a fact source by its id

Returns a Fact Source, given a source id.
*/
func (a *Client) GetAPIV2SourcesID(params *GetAPIV2SourcesIDParams, opts ...ClientOption) (*GetAPIV2SourcesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2SourcesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2SourcesID",
		Method:             "GET",
		PathPattern:        "/api/v2/sources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV2SourcesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIV2SourcesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2SourcesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HeadAPIV2Sources retrieves all fact sources

Returns a list of all Fact Sources, including custom-created ones.
*/
func (a *Client) HeadAPIV2Sources(params *HeadAPIV2SourcesParams, opts ...ClientOption) (*HeadAPIV2SourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadAPIV2SourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeadAPIV2Sources",
		Method:             "HEAD",
		PathPattern:        "/api/v2/sources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadAPIV2SourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HeadAPIV2SourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HeadAPIV2Sources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HeadAPIV2SourcesID retrieves a fact source by its id

Returns a Fact Source, given a source id.
*/
func (a *Client) HeadAPIV2SourcesID(params *HeadAPIV2SourcesIDParams, opts ...ClientOption) (*HeadAPIV2SourcesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadAPIV2SourcesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeadAPIV2SourcesID",
		Method:             "HEAD",
		PathPattern:        "/api/v2/sources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadAPIV2SourcesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HeadAPIV2SourcesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HeadAPIV2SourcesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPIV2SourcesID updates an existing fact source

Returns an updated Fact Source. All fields in a Fact Source can be updated, except for "id" and "adjustments".
*/
func (a *Client) PatchAPIV2SourcesID(params *PatchAPIV2SourcesIDParams, opts ...ClientOption) (*PatchAPIV2SourcesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV2SourcesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIV2SourcesID",
		Method:             "PATCH",
		PathPattern:        "/api/v2/sources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPIV2SourcesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAPIV2SourcesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIV2SourcesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPIV2Sources creates a fact source

Create a new Fact Source using the format provided in the SourceSchema.
*/
func (a *Client) PostAPIV2Sources(params *PostAPIV2SourcesParams, opts ...ClientOption) (*PostAPIV2SourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV2SourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV2Sources",
		Method:             "POST",
		PathPattern:        "/api/v2/sources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPIV2SourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAPIV2SourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIV2Sources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutAPIV2SourcesID updates an existing or create a new fact source

Use fields from the SourceSchema in the request body to replace an existing Fact Source or create a new Fact Source.
*/
func (a *Client) PutAPIV2SourcesID(params *PutAPIV2SourcesIDParams, opts ...ClientOption) (*PutAPIV2SourcesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAPIV2SourcesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutAPIV2SourcesID",
		Method:             "PUT",
		PathPattern:        "/api/v2/sources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutAPIV2SourcesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutAPIV2SourcesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutAPIV2SourcesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
