# InfrastructureEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]InfrastructurePort**](InfrastructurePort.md) |  | [optional] 
**PortsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Ready** | Pointer to **bool** |  | [optional] 
**Target** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 

## Methods

### NewInfrastructureEndpoint

`func NewInfrastructureEndpoint() *InfrastructureEndpoint`

NewInfrastructureEndpoint instantiates a new InfrastructureEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureEndpointWithDefaults

`func NewInfrastructureEndpointWithDefaults() *InfrastructureEndpoint`

NewInfrastructureEndpointWithDefaults instantiates a new InfrastructureEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InfrastructureEndpoint) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InfrastructureEndpoint) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InfrastructureEndpoint) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InfrastructureEndpoint) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPorts

`func (o *InfrastructureEndpoint) GetPorts() []InfrastructurePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InfrastructureEndpoint) GetPortsOk() (*[]InfrastructurePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InfrastructureEndpoint) SetPorts(v []InfrastructurePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InfrastructureEndpoint) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPortsTruncation

`func (o *InfrastructureEndpoint) GetPortsTruncation() CheckpointTruncation`

GetPortsTruncation returns the PortsTruncation field if non-nil, zero value otherwise.

### GetPortsTruncationOk

`func (o *InfrastructureEndpoint) GetPortsTruncationOk() (*CheckpointTruncation, bool)`

GetPortsTruncationOk returns a tuple with the PortsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsTruncation

`func (o *InfrastructureEndpoint) SetPortsTruncation(v CheckpointTruncation)`

SetPortsTruncation sets PortsTruncation field to given value.

### HasPortsTruncation

`func (o *InfrastructureEndpoint) HasPortsTruncation() bool`

HasPortsTruncation returns a boolean if a field has been set.

### GetReady

`func (o *InfrastructureEndpoint) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *InfrastructureEndpoint) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *InfrastructureEndpoint) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *InfrastructureEndpoint) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetTarget

`func (o *InfrastructureEndpoint) GetTarget() InfrastructureObjectReference`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InfrastructureEndpoint) GetTargetOk() (*InfrastructureObjectReference, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InfrastructureEndpoint) SetTarget(v InfrastructureObjectReference)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *InfrastructureEndpoint) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


