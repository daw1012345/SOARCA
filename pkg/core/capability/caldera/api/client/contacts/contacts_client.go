// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new contacts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new contacts API client with basic auth credentials.
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

// New creates a new contacts API client with a bearer token for authentication.
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
Client for contacts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIV2Contacts(params *GetAPIV2ContactsParams, opts ...ClientOption) (*GetAPIV2ContactsOK, error)

	GetAPIV2ContactsName(params *GetAPIV2ContactsNameParams, opts ...ClientOption) (*GetAPIV2ContactsNameOK, error)

	HeadAPIV2Contacts(params *HeadAPIV2ContactsParams, opts ...ClientOption) (*HeadAPIV2ContactsOK, error)

	HeadAPIV2ContactsName(params *HeadAPIV2ContactsNameParams, opts ...ClientOption) (*HeadAPIV2ContactsNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAPIV2Contacts retrieves a list of all available contact reports

Returns a list of contacts that at least one agent has beaconed to. As soon as any agent beacons over a given contact, the contact will be returned here.
*/
func (a *Client) GetAPIV2Contacts(params *GetAPIV2ContactsParams, opts ...ClientOption) (*GetAPIV2ContactsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2ContactsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2Contacts",
		Method:             "GET",
		PathPattern:        "/api/v2/contacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV2ContactsReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2ContactsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2Contacts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV2ContactsName retrieves a list of beacons made by agents to the specified contact

Returns a list of beacons made by agents to the specified contact. The response is formatted as a list of dictionaries. The dictionaries have the keys `paw`, `instructions`, and `date`. `paw` being the paw of the agent that made the beacon. `instructions` being a list of strings (commands) executed by the agent since its last beacon. `date` being a UTC date/time string.
*/
func (a *Client) GetAPIV2ContactsName(params *GetAPIV2ContactsNameParams, opts ...ClientOption) (*GetAPIV2ContactsNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2ContactsNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2ContactsName",
		Method:             "GET",
		PathPattern:        "/api/v2/contacts/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV2ContactsNameReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2ContactsNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2ContactsName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HeadAPIV2Contacts retrieves a list of all available contact reports

Returns a list of contacts that at least one agent has beaconed to. As soon as any agent beacons over a given contact, the contact will be returned here.
*/
func (a *Client) HeadAPIV2Contacts(params *HeadAPIV2ContactsParams, opts ...ClientOption) (*HeadAPIV2ContactsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadAPIV2ContactsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeadAPIV2Contacts",
		Method:             "HEAD",
		PathPattern:        "/api/v2/contacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadAPIV2ContactsReader{formats: a.formats},
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
	success, ok := result.(*HeadAPIV2ContactsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HeadAPIV2Contacts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HeadAPIV2ContactsName retrieves a list of beacons made by agents to the specified contact

Returns a list of beacons made by agents to the specified contact. The response is formatted as a list of dictionaries. The dictionaries have the keys `paw`, `instructions`, and `date`. `paw` being the paw of the agent that made the beacon. `instructions` being a list of strings (commands) executed by the agent since its last beacon. `date` being a UTC date/time string.
*/
func (a *Client) HeadAPIV2ContactsName(params *HeadAPIV2ContactsNameParams, opts ...ClientOption) (*HeadAPIV2ContactsNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadAPIV2ContactsNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeadAPIV2ContactsName",
		Method:             "HEAD",
		PathPattern:        "/api/v2/contacts/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadAPIV2ContactsNameReader{formats: a.formats},
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
	success, ok := result.(*HeadAPIV2ContactsNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HeadAPIV2ContactsName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
