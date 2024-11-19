// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new config API client with basic auth credentials.
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

// New creates a new config API client with a bearer token for authentication.
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
Client for config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIV2ConfigName(params *GetAPIV2ConfigNameParams, opts ...ClientOption) (*GetAPIV2ConfigNameOK, error)

	HeadAPIV2ConfigName(params *HeadAPIV2ConfigNameParams, opts ...ClientOption) (*HeadAPIV2ConfigNameOK, error)

	PatchAPIV2ConfigAgents(params *PatchAPIV2ConfigAgentsParams, opts ...ClientOption) (*PatchAPIV2ConfigAgentsOK, error)

	PatchAPIV2ConfigMain(params *PatchAPIV2ConfigMainParams, opts ...ClientOption) (*PatchAPIV2ConfigMainOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAPIV2ConfigName retrieves config

Retrieves configuration by name, as specified by {name} in the request url.
*/
func (a *Client) GetAPIV2ConfigName(params *GetAPIV2ConfigNameParams, opts ...ClientOption) (*GetAPIV2ConfigNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2ConfigNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2ConfigName",
		Method:             "GET",
		PathPattern:        "/api/v2/config/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV2ConfigNameReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2ConfigNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2ConfigName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HeadAPIV2ConfigName retrieves config

Retrieves configuration by name, as specified by {name} in the request url.
*/
func (a *Client) HeadAPIV2ConfigName(params *HeadAPIV2ConfigNameParams, opts ...ClientOption) (*HeadAPIV2ConfigNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadAPIV2ConfigNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeadAPIV2ConfigName",
		Method:             "HEAD",
		PathPattern:        "/api/v2/config/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadAPIV2ConfigNameReader{formats: a.formats},
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
	success, ok := result.(*HeadAPIV2ConfigNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HeadAPIV2ConfigName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPIV2ConfigAgents updates agent config

Use fields from the AgentConfigUpdateSchema in the request body to update the Agent Configuration file.
*/
func (a *Client) PatchAPIV2ConfigAgents(params *PatchAPIV2ConfigAgentsParams, opts ...ClientOption) (*PatchAPIV2ConfigAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV2ConfigAgentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIV2ConfigAgents",
		Method:             "PATCH",
		PathPattern:        "/api/v2/config/agents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPIV2ConfigAgentsReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIV2ConfigAgentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIV2ConfigAgents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPIV2ConfigMain updates main config

Use fields from the ConfigUpdateSchema in the request body to update the main configuration file.
*/
func (a *Client) PatchAPIV2ConfigMain(params *PatchAPIV2ConfigMainParams, opts ...ClientOption) (*PatchAPIV2ConfigMainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV2ConfigMainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIV2ConfigMain",
		Method:             "PATCH",
		PathPattern:        "/api/v2/config/main",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPIV2ConfigMainReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIV2ConfigMainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIV2ConfigMain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
