// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api keys API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api keys API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPIKey(params *DeleteAPIKeyParams, opts ...ClientOption) (*DeleteAPIKeyOK, error)

	GenerateAPIKey(params *GenerateAPIKeyParams, opts ...ClientOption) (*GenerateAPIKeyOK, error)

	ListAPIKeys(params *ListAPIKeysParams, opts ...ClientOption) (*ListAPIKeysOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAPIKey deletes

Deletes an API key, revoking all access for requests that use the key.
*/
func (a *Client) DeleteAPIKey(params *DeleteAPIKeyParams, opts ...ClientOption) (*DeleteAPIKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteApiKey",
		Method:             "DELETE",
		PathPattern:        "/v1/api-keys/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIKeyReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteAPIKeyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GenerateAPIKey generates

Creates a new API key for the requesting user. The API key is returned in the response, and this is the only time it can be viewed.
*/
func (a *Client) GenerateAPIKey(params *GenerateAPIKeyParams, opts ...ClientOption) (*GenerateAPIKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GenerateApiKey",
		Method:             "POST",
		PathPattern:        "/v1/api-keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GenerateAPIKeyReader{formats: a.formats},
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
	success, ok := result.(*GenerateAPIKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GenerateAPIKeyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAPIKeys lists

List the details of all API keys created by the requesting user. This does not include the API key itself which is only visible once when the API key is created.
*/
func (a *Client) ListAPIKeys(params *ListAPIKeysParams, opts ...ClientOption) (*ListAPIKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAPIKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListApiKeys",
		Method:             "GET",
		PathPattern:        "/v1/api-keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAPIKeysReader{formats: a.formats},
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
	success, ok := result.(*ListAPIKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAPIKeysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
