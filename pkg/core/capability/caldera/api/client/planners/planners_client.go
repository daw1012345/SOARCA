// Code generated by go-swagger; DO NOT EDIT.

package planners

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new planners API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new planners API client with basic auth credentials.
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

// New creates a new planners API client with a bearer token for authentication.
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
Client for planners API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIV2Planners(params *GetAPIV2PlannersParams, opts ...ClientOption) (*GetAPIV2PlannersOK, error)

	GetAPIV2PlannersPlannerID(params *GetAPIV2PlannersPlannerIDParams, opts ...ClientOption) (*GetAPIV2PlannersPlannerIDOK, error)

	HeadAPIV2Planners(params *HeadAPIV2PlannersParams, opts ...ClientOption) (*HeadAPIV2PlannersOK, error)

	HeadAPIV2PlannersPlannerID(params *HeadAPIV2PlannersPlannerIDParams, opts ...ClientOption) (*HeadAPIV2PlannersPlannerIDOK, error)

	PatchAPIV2PlannersPlannerID(params *PatchAPIV2PlannersPlannerIDParams, opts ...ClientOption) (*PatchAPIV2PlannersPlannerIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAPIV2Planners retrieves planners

Retrieve Caldera planners by criteria. Supply fields from the `PlannerSchema` to the `include` and `exclude` fields of the `BaseGetAllQuerySchema` in the request body to filter retrieved planners.
*/
func (a *Client) GetAPIV2Planners(params *GetAPIV2PlannersParams, opts ...ClientOption) (*GetAPIV2PlannersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2PlannersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2Planners",
		Method:             "GET",
		PathPattern:        "/api/v2/planners",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV2PlannersReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2PlannersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2Planners: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIV2PlannersPlannerID retrieves a planner by planner id

Retrieve one Caldera planner based on the planner id (String `UUID`). Supply fields from the `PlannerSchema` to the `include` and `exclude` fields of the `BaseGetOneQuerySchema` in the request body to filter retrieved planners.
*/
func (a *Client) GetAPIV2PlannersPlannerID(params *GetAPIV2PlannersPlannerIDParams, opts ...ClientOption) (*GetAPIV2PlannersPlannerIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIV2PlannersPlannerIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIV2PlannersPlannerID",
		Method:             "GET",
		PathPattern:        "/api/v2/planners/{planner_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIV2PlannersPlannerIDReader{formats: a.formats},
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
	success, ok := result.(*GetAPIV2PlannersPlannerIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIV2PlannersPlannerID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HeadAPIV2Planners retrieves planners

Retrieve Caldera planners by criteria. Supply fields from the `PlannerSchema` to the `include` and `exclude` fields of the `BaseGetAllQuerySchema` in the request body to filter retrieved planners.
*/
func (a *Client) HeadAPIV2Planners(params *HeadAPIV2PlannersParams, opts ...ClientOption) (*HeadAPIV2PlannersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadAPIV2PlannersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeadAPIV2Planners",
		Method:             "HEAD",
		PathPattern:        "/api/v2/planners",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadAPIV2PlannersReader{formats: a.formats},
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
	success, ok := result.(*HeadAPIV2PlannersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HeadAPIV2Planners: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HeadAPIV2PlannersPlannerID retrieves a planner by planner id

Retrieve one Caldera planner based on the planner id (String `UUID`). Supply fields from the `PlannerSchema` to the `include` and `exclude` fields of the `BaseGetOneQuerySchema` in the request body to filter retrieved planners.
*/
func (a *Client) HeadAPIV2PlannersPlannerID(params *HeadAPIV2PlannersPlannerIDParams, opts ...ClientOption) (*HeadAPIV2PlannersPlannerIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeadAPIV2PlannersPlannerIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HeadAPIV2PlannersPlannerID",
		Method:             "HEAD",
		PathPattern:        "/api/v2/planners/{planner_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeadAPIV2PlannersPlannerIDReader{formats: a.formats},
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
	success, ok := result.(*HeadAPIV2PlannersPlannerIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HeadAPIV2PlannersPlannerID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPIV2PlannersPlannerID updates an existing planner

Updates a planner based on the `PlannerSchema` value provided in the message body.
*/
func (a *Client) PatchAPIV2PlannersPlannerID(params *PatchAPIV2PlannersPlannerIDParams, opts ...ClientOption) (*PatchAPIV2PlannersPlannerIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIV2PlannersPlannerIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIV2PlannersPlannerID",
		Method:             "PATCH",
		PathPattern:        "/api/v2/planners/{planner_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPIV2PlannersPlannerIDReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIV2PlannersPlannerIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIV2PlannersPlannerID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
