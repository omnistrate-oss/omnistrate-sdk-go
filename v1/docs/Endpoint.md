# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The host of the endpoint | [optional] [default to ""]
**NetworkingType** | Pointer to **string** | The networking type of the endpoint | [optional] [default to "PUBLIC"]
**Ports** | Pointer to **[]int64** | The ports of the endpoint | [optional] 
**Primary** | Pointer to **bool** | Whether this is the primary endpoint to highlight | [optional] [default to false]

## Methods

### NewEndpoint

`func NewEndpoint() *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *Endpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Endpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Endpoint) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Endpoint) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNetworkingType

`func (o *Endpoint) GetNetworkingType() string`

GetNetworkingType returns the NetworkingType field if non-nil, zero value otherwise.

### GetNetworkingTypeOk

`func (o *Endpoint) GetNetworkingTypeOk() (*string, bool)`

GetNetworkingTypeOk returns a tuple with the NetworkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingType

`func (o *Endpoint) SetNetworkingType(v string)`

SetNetworkingType sets NetworkingType field to given value.

### HasNetworkingType

`func (o *Endpoint) HasNetworkingType() bool`

HasNetworkingType returns a boolean if a field has been set.

### GetPorts

`func (o *Endpoint) GetPorts() []int64`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Endpoint) GetPortsOk() (*[]int64, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Endpoint) SetPorts(v []int64)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *Endpoint) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPrimary

`func (o *Endpoint) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Endpoint) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Endpoint) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *Endpoint) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


