// Code generated by go-swagger; DO NOT EDIT.

package disks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new disks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for disks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AttachStorageDisk(params *AttachStorageDiskParams, opts ...ClientOption) (*AttachStorageDiskOK, error)

	CreateDiskSnapshot(params *CreateDiskSnapshotParams, opts ...ClientOption) (*CreateDiskSnapshotOK, error)

	CreateStorageDisk(params *CreateStorageDiskParams, opts ...ClientOption) (*CreateStorageDiskOK, error)

	DeleteDiskSnapshot(params *DeleteDiskSnapshotParams, opts ...ClientOption) (*DeleteDiskSnapshotOK, error)

	DeleteStorageDisk(params *DeleteStorageDiskParams, opts ...ClientOption) (*DeleteStorageDiskOK, error)

	DetachStorageDisk(params *DetachStorageDiskParams, opts ...ClientOption) (*DetachStorageDiskOK, error)

	GetDisk(params *GetDiskParams, opts ...ClientOption) (*GetDiskOK, error)

	ListDiskSnapshots(params *ListDiskSnapshotsParams, opts ...ClientOption) (*ListDiskSnapshotsOK, error)

	ListDisks(params *ListDisksParams, opts ...ClientOption) (*ListDisksOK, error)

	RevertDisk(params *RevertDiskParams, opts ...ClientOption) (*RevertDiskOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AttachStorageDisk attaches storage disk to VM
*/
func (a *Client) AttachStorageDisk(params *AttachStorageDiskParams, opts ...ClientOption) (*AttachStorageDiskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachStorageDiskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachStorageDisk",
		Method:             "PATCH",
		PathPattern:        "/v1/projects/{projectId}/disk/{id}/attach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AttachStorageDiskReader{formats: a.formats},
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
	success, ok := result.(*AttachStorageDiskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachStorageDiskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateDiskSnapshot creates disk snapshot
*/
func (a *Client) CreateDiskSnapshot(params *CreateDiskSnapshotParams, opts ...ClientOption) (*CreateDiskSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDiskSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateDiskSnapshot",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/disks/{id}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDiskSnapshotReader{formats: a.formats},
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
	success, ok := result.(*CreateDiskSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDiskSnapshotDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateStorageDisk creates storage disk
*/
func (a *Client) CreateStorageDisk(params *CreateStorageDiskParams, opts ...ClientOption) (*CreateStorageDiskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStorageDiskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateStorageDisk",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/disks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateStorageDiskReader{formats: a.formats},
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
	success, ok := result.(*CreateStorageDiskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateStorageDiskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteDiskSnapshot deletes disk snapshots
*/
func (a *Client) DeleteDiskSnapshot(params *DeleteDiskSnapshotParams, opts ...ClientOption) (*DeleteDiskSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDiskSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDiskSnapshot",
		Method:             "DELETE",
		PathPattern:        "/v1/projects/{projectId}/disks/{id}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDiskSnapshotReader{formats: a.formats},
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
	success, ok := result.(*DeleteDiskSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDiskSnapshotDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteStorageDisk deletes storage disk
*/
func (a *Client) DeleteStorageDisk(params *DeleteStorageDiskParams, opts ...ClientOption) (*DeleteStorageDiskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStorageDiskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteStorageDisk",
		Method:             "DELETE",
		PathPattern:        "/v1/projects/{projectId}/disks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteStorageDiskReader{formats: a.formats},
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
	success, ok := result.(*DeleteStorageDiskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteStorageDiskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DetachStorageDisk detaches storage disk from VM
*/
func (a *Client) DetachStorageDisk(params *DetachStorageDiskParams, opts ...ClientOption) (*DetachStorageDiskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachStorageDiskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DetachStorageDisk",
		Method:             "PUT",
		PathPattern:        "/v1/projects/{projectId}/disk/{id}/detach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetachStorageDiskReader{formats: a.formats},
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
	success, ok := result.(*DetachStorageDiskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DetachStorageDiskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDisk lists disks
*/
func (a *Client) GetDisk(params *GetDiskParams, opts ...ClientOption) (*GetDiskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDiskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDisk",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/disks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDiskReader{formats: a.formats},
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
	success, ok := result.(*GetDiskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDiskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListDiskSnapshots lists disk snapshots
*/
func (a *Client) ListDiskSnapshots(params *ListDiskSnapshotsParams, opts ...ClientOption) (*ListDiskSnapshotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDiskSnapshotsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListDiskSnapshots",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/disks/{id}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDiskSnapshotsReader{formats: a.formats},
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
	success, ok := result.(*ListDiskSnapshotsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDiskSnapshotsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListDisks lists disks
*/
func (a *Client) ListDisks(params *ListDisksParams, opts ...ClientOption) (*ListDisksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDisksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListDisks",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/disks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDisksReader{formats: a.formats},
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
	success, ok := result.(*ListDisksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDisksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RevertDisk reverts disk to snapshot
*/
func (a *Client) RevertDisk(params *RevertDiskParams, opts ...ClientOption) (*RevertDiskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevertDiskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RevertDisk",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/disks/{id}/revert",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RevertDiskReader{formats: a.formats},
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
	success, ok := result.(*RevertDiskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RevertDiskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
