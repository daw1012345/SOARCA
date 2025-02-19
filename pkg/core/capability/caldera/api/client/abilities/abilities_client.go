// Code generated by go-swagger; DO NOT EDIT.

package abilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new abilities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new abilities API client with basic auth credentials.
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

// New creates a new abilities API client with a bearer token for authentication.
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
Client for abilities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIV2AbilitiesAbilityID(params *DeleteAPIV2AbilitiesAbilityIDParams, opts ...ClientOption) (*DeleteAPIV2AbilitiesAbilityIDNoContent, error)

	GetAPIV2Abilities(params *GetAPIV2AbilitiesParams, opts ...ClientOption) (*GetAPIV2AbilitiesOK, error)

	GetAPIV2AbilitiesAbilityID(params *GetAPIV2AbilitiesAbilityIDParams, opts ...ClientOption) (*GetAPIV2AbilitiesAbilityIDOK, error)

	HeadAPIV2Abilities(params *HeadAPIV2AbilitiesParams, opts ...ClientOption) (*HeadAPIV2AbilitiesOK, error)

	HeadAPIV2AbilitiesAbilityID(params *HeadAPIV2AbilitiesAbilityIDParams, opts ...ClientOption) (*HeadAPIV2AbilitiesAbilityIDOK, error)

	PatchAPIV2AbilitiesAbilityID(params *PatchAPIV2AbilitiesAbilityIDParams, opts ...ClientOption) (*PatchAPIV2AbilitiesAbilityIDOK, error)

	PostAPIV2Abilities(params *PostAPIV2AbilitiesParams, opts ...ClientOption) (*PostAPIV2AbilitiesOK, error)

	PutAPIV2AbilitiesAbilityID(params *PutAPIV2AbilitiesAbilityIDParams, opts ...ClientOption) (*PutAPIV2AbilitiesAbilityIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAPIV2AbilitiesAbilityID deletes an ability

Deletes an existing ability.
*/
func (a *Client) DeleteAPIV2AbilitiesAbilityID(params *DeleteAPIV2AbilitiesAbilityIDParams, opts ...ClientOption) (*DeleteAPIV2AbilitiesAbilityIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIV2AbilitiesAbilityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPIV2AbilitiesAbilityID",
		Method:             "DELETE",
		PathPattern:        "/api/v2/abilities/{ability_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAPIV2AbilitiesAbilityIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIV2AbilitiesAbilityIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPIV2AbilitiesAbilityID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV2Abilities gets all abilities

Provides a list of all available abilities.
*/
func (a *Client) GetAPIV2Abilities(params *GetAPIV2AbilitiesParams, opts ...ClientOption) (*GetAPIV2AbilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2AbilitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2Abilities",
		Method:             "GET",
		PathPattern:        "/api/v2/abilities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV2AbilitiesReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2AbilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2Abilities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV2AbilitiesAbilityID gets an ability

Provides one ability based on its ability id.
*/
func (a *Client) GetAPIV2AbilitiesAbilityID(params *GetAPIV2AbilitiesAbilityIDParams, opts ...ClientOption) (*GetAPIV2AbilitiesAbilityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2AbilitiesAbilityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2AbilitiesAbilityID",
		Method:             "GET",
		PathPattern:        "/api/v2/abilities/{ability_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV2AbilitiesAbilityIDReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2AbilitiesAbilityIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2AbilitiesAbilityID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HeadAPIV2Abilities gets all abilities

Provides a list of all available abilities.
*/
func (a *Client) HeadAPIV2Abilities(params *HeadAPIV2AbilitiesParams, opts ...ClientOption) (*HeadAPIV2AbilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadAPIV2AbilitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeadAPIV2Abilities",
		Method:             "HEAD",
		PathPattern:        "/api/v2/abilities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadAPIV2AbilitiesReader{formats: a.formats},
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
	success, ok := result.(*HeadAPIV2AbilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HeadAPIV2Abilities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HeadAPIV2AbilitiesAbilityID gets an ability

Provides one ability based on its ability id.
*/
func (a *Client) HeadAPIV2AbilitiesAbilityID(params *HeadAPIV2AbilitiesAbilityIDParams, opts ...ClientOption) (*HeadAPIV2AbilitiesAbilityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadAPIV2AbilitiesAbilityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeadAPIV2AbilitiesAbilityID",
		Method:             "HEAD",
		PathPattern:        "/api/v2/abilities/{ability_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadAPIV2AbilitiesAbilityIDReader{formats: a.formats},
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
	success, ok := result.(*HeadAPIV2AbilitiesAbilityIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HeadAPIV2AbilitiesAbilityID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPIV2AbilitiesAbilityID updates an existing ability

Updates an ability based on the `AbilitySchema` values provided in the message body.
*/
func (a *Client) PatchAPIV2AbilitiesAbilityID(params *PatchAPIV2AbilitiesAbilityIDParams, opts ...ClientOption) (*PatchAPIV2AbilitiesAbilityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV2AbilitiesAbilityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIV2AbilitiesAbilityID",
		Method:             "PATCH",
		PathPattern:        "/api/v2/abilities/{ability_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPIV2AbilitiesAbilityIDReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIV2AbilitiesAbilityIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIV2AbilitiesAbilityID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPIV2Abilities creates a new ability

Creates a new ability based on the `AbilitySchema`. "name", "tactic", "technique_name", "technique_id" and "executors" are all required fields.
*/
func (a *Client) PostAPIV2Abilities(params *PostAPIV2AbilitiesParams, opts ...ClientOption) (*PostAPIV2AbilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIV2AbilitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIV2Abilities",
		Method:             "POST",
		PathPattern:        "/api/v2/abilities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPIV2AbilitiesReader{formats: a.formats},
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
	success, ok := result.(*PostAPIV2AbilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIV2Abilities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutAPIV2AbilitiesAbilityID replaces an existing ability

Replaces an ability based on the `AbilitySchema` values provided in the message body. "name", "tactic", and "executors" are all required fields.
*/
func (a *Client) PutAPIV2AbilitiesAbilityID(params *PutAPIV2AbilitiesAbilityIDParams, opts ...ClientOption) (*PutAPIV2AbilitiesAbilityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAPIV2AbilitiesAbilityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutAPIV2AbilitiesAbilityID",
		Method:             "PUT",
		PathPattern:        "/api/v2/abilities/{ability_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutAPIV2AbilitiesAbilityIDReader{formats: a.formats},
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
	success, ok := result.(*PutAPIV2AbilitiesAbilityIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutAPIV2AbilitiesAbilityID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
