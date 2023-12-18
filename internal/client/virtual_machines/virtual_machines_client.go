// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new virtual machines API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for virtual machines API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ConnectVM(params *ConnectVMParams, opts ...ClientOption) (*ConnectVMOK, error)

	CountVMs(params *CountVMsParams, opts ...ClientOption) (*CountVMsOK, error)

	CreatePrivateVMImage(params *CreatePrivateVMImageParams, opts ...ClientOption) (*CreatePrivateVMImageOK, error)

	CreateVM(params *CreateVMParams, opts ...ClientOption) (*CreateVMOK, error)

	DeletePrivateVMImage(params *DeletePrivateVMImageParams, opts ...ClientOption) (*DeletePrivateVMImageOK, error)

	GetPrivateVMImage(params *GetPrivateVMImageParams, opts ...ClientOption) (*GetPrivateVMImageOK, error)

	GetVM(params *GetVMParams, opts ...ClientOption) (*GetVMOK, error)

	ListPrivateVMImages(params *ListPrivateVMImagesParams, opts ...ClientOption) (*ListPrivateVMImagesOK, error)

	ListPublicVMImages(params *ListPublicVMImagesParams, opts ...ClientOption) (*ListPublicVMImagesOK, error)

	ListVMDataCenters(params *ListVMDataCentersParams, opts ...ClientOption) (*ListVMDataCentersOK, error)

	ListVMDisks(params *ListVMDisksParams, opts ...ClientOption) (*ListVMDisksOK, error)

	ListVMMachineTypes(params *ListVMMachineTypesParams, opts ...ClientOption) (*ListVMMachineTypesOK, error)

	ListVMs(params *ListVMsParams, opts ...ClientOption) (*ListVMsOK, error)

	MonitorVM(params *MonitorVMParams, opts ...ClientOption) (*MonitorVMOK, error)

	RebootVM(params *RebootVMParams, opts ...ClientOption) (*RebootVMOK, error)

	ResizeVM(params *ResizeVMParams, opts ...ClientOption) (*ResizeVMOK, error)

	ResizeVMDisk(params *ResizeVMDiskParams, opts ...ClientOption) (*ResizeVMDiskOK, error)

	StartVM(params *StartVMParams, opts ...ClientOption) (*StartVMOK, error)

	StopVM(params *StopVMParams, opts ...ClientOption) (*StopVMOK, error)

	TerminateVM(params *TerminateVMParams, opts ...ClientOption) (*TerminateVMOK, error)

	UpdatePrivateVMImage(params *UpdatePrivateVMImageParams, opts ...ClientOption) (*UpdatePrivateVMImageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ConnectVM connects via v n c
*/
func (a *Client) ConnectVM(params *ConnectVMParams, opts ...ClientOption) (*ConnectVMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectVMParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConnectVM",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/vms/{id}/connect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConnectVMReader{formats: a.formats},
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
	success, ok := result.(*ConnectVMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConnectVMDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CountVMs counts
*/
func (a *Client) CountVMs(params *CountVMsParams, opts ...ClientOption) (*CountVMsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountVMsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CountVMs",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/count-vms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CountVMsReader{formats: a.formats},
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
	success, ok := result.(*CountVMsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CountVMsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreatePrivateVMImage creates private VM image
*/
func (a *Client) CreatePrivateVMImage(params *CreatePrivateVMImageParams, opts ...ClientOption) (*CreatePrivateVMImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePrivateVMImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreatePrivateVMImage",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePrivateVMImageReader{formats: a.formats},
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
	success, ok := result.(*CreatePrivateVMImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreatePrivateVMImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateVM creates virtual machine
*/
func (a *Client) CreateVM(params *CreateVMParams, opts ...ClientOption) (*CreateVMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVMParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateVM",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/vm",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateVMReader{formats: a.formats},
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
	success, ok := result.(*CreateVMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateVMDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeletePrivateVMImage deletes private VM image
*/
func (a *Client) DeletePrivateVMImage(params *DeletePrivateVMImageParams, opts ...ClientOption) (*DeletePrivateVMImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivateVMImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeletePrivateVMImage",
		Method:             "DELETE",
		PathPattern:        "/v1/projects/{projectId}/images/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePrivateVMImageReader{formats: a.formats},
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
	success, ok := result.(*DeletePrivateVMImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePrivateVMImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetPrivateVMImage gets private VM image
*/
func (a *Client) GetPrivateVMImage(params *GetPrivateVMImageParams, opts ...ClientOption) (*GetPrivateVMImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateVMImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPrivateVMImage",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/images/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPrivateVMImageReader{formats: a.formats},
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
	success, ok := result.(*GetPrivateVMImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPrivateVMImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVM gets
*/
func (a *Client) GetVM(params *GetVMParams, opts ...ClientOption) (*GetVMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVM",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/vms/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVMReader{formats: a.formats},
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
	success, ok := result.(*GetVMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVMDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListPrivateVMImages lists private VM images
*/
func (a *Client) ListPrivateVMImages(params *ListPrivateVMImagesParams, opts ...ClientOption) (*ListPrivateVMImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPrivateVMImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListPrivateVMImages",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPrivateVMImagesReader{formats: a.formats},
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
	success, ok := result.(*ListPrivateVMImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListPrivateVMImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListPublicVMImages lists public VM images
*/
func (a *Client) ListPublicVMImages(params *ListPublicVMImagesParams, opts ...ClientOption) (*ListPublicVMImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPublicVMImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListPublicVMImages",
		Method:             "GET",
		PathPattern:        "/v1/vms/public-images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPublicVMImagesReader{formats: a.formats},
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
	success, ok := result.(*ListPublicVMImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListPublicVMImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVMDataCenters lists data centers
*/
func (a *Client) ListVMDataCenters(params *ListVMDataCentersParams, opts ...ClientOption) (*ListVMDataCentersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVMDataCentersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListVMDataCenters",
		Method:             "GET",
		PathPattern:        "/v1/vms/data-centers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVMDataCentersReader{formats: a.formats},
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
	success, ok := result.(*ListVMDataCentersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVMDataCentersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVMDisks lists disks attached to VM
*/
func (a *Client) ListVMDisks(params *ListVMDisksParams, opts ...ClientOption) (*ListVMDisksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVMDisksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListVMDisks",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/vms/{id}/disks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVMDisksReader{formats: a.formats},
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
	success, ok := result.(*ListVMDisksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVMDisksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVMMachineTypes lists machine types
*/
func (a *Client) ListVMMachineTypes(params *ListVMMachineTypesParams, opts ...ClientOption) (*ListVMMachineTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVMMachineTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListVMMachineTypes",
		Method:             "GET",
		PathPattern:        "/v1/vms/machine-types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVMMachineTypesReader{formats: a.formats},
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
	success, ok := result.(*ListVMMachineTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVMMachineTypesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVMs lists
*/
func (a *Client) ListVMs(params *ListVMsParams, opts ...ClientOption) (*ListVMsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVMsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListVMs",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/vms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVMsReader{formats: a.formats},
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
	success, ok := result.(*ListVMsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVMsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MonitorVM monitors
*/
func (a *Client) MonitorVM(params *MonitorVMParams, opts ...ClientOption) (*MonitorVMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMonitorVMParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MonitorVM",
		Method:             "GET",
		PathPattern:        "/v1/projects/{projectId}/vms/{id}/monitor",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MonitorVMReader{formats: a.formats},
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
	success, ok := result.(*MonitorVMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MonitorVMDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RebootVM reboots
*/
func (a *Client) RebootVM(params *RebootVMParams, opts ...ClientOption) (*RebootVMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebootVMParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RebootVM",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/vms/{id}/reboot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RebootVMReader{formats: a.formats},
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
	success, ok := result.(*RebootVMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RebootVMDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ResizeVM resizes v CPU and memory of VM
*/
func (a *Client) ResizeVM(params *ResizeVMParams, opts ...ClientOption) (*ResizeVMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResizeVMParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResizeVM",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/vms/{id}/resize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResizeVMReader{formats: a.formats},
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
	success, ok := result.(*ResizeVMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResizeVMDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ResizeVMDisk resizes a VM s disk
*/
func (a *Client) ResizeVMDisk(params *ResizeVMDiskParams, opts ...ClientOption) (*ResizeVMDiskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResizeVMDiskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResizeVMDisk",
		Method:             "PATCH",
		PathPattern:        "/v1/projects/{projectId}/vms/{id}/disks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResizeVMDiskReader{formats: a.formats},
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
	success, ok := result.(*ResizeVMDiskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResizeVMDiskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StartVM starts
*/
func (a *Client) StartVM(params *StartVMParams, opts ...ClientOption) (*StartVMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartVMParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartVM",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/vms/{id}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartVMReader{formats: a.formats},
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
	success, ok := result.(*StartVMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartVMDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StopVM stops
*/
func (a *Client) StopVM(params *StopVMParams, opts ...ClientOption) (*StopVMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopVMParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StopVM",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/vms/{id}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StopVMReader{formats: a.formats},
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
	success, ok := result.(*StopVMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StopVMDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
TerminateVM terminates
*/
func (a *Client) TerminateVM(params *TerminateVMParams, opts ...ClientOption) (*TerminateVMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTerminateVMParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TerminateVM",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/vms/{id}/terminate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TerminateVMReader{formats: a.formats},
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
	success, ok := result.(*TerminateVMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TerminateVMDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdatePrivateVMImage updates private VM image
*/
func (a *Client) UpdatePrivateVMImage(params *UpdatePrivateVMImageParams, opts ...ClientOption) (*UpdatePrivateVMImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePrivateVMImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdatePrivateVMImage",
		Method:             "POST",
		PathPattern:        "/v1/projects/{projectId}/images/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePrivateVMImageReader{formats: a.formats},
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
	success, ok := result.(*UpdatePrivateVMImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdatePrivateVMImageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
