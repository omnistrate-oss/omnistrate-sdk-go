# ClusterEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** | The endpoint | [optional] [default to ""]
**NetworkingType** | Pointer to **string** | The networking type for this resource | [optional] 
**OpenPorts** | Pointer to **[]int64** | The open ports | [optional] 
**Primary** | Pointer to **bool** | Whether this is the primary endpoint | [optional] [default to false]

## Methods

### NewClusterEndpoint

`func NewClusterEndpoint() *ClusterEndpoint`

NewClusterEndpoint instantiates a new ClusterEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEndpointWithDefaults

`func NewClusterEndpointWithDefaults() *ClusterEndpoint`

NewClusterEndpointWithDefaults instantiates a new ClusterEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *ClusterEndpoint) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ClusterEndpoint) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ClusterEndpoint) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ClusterEndpoint) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetNetworkingType

`func (o *ClusterEndpoint) GetNetworkingType() string`

GetNetworkingType returns the NetworkingType field if non-nil, zero value otherwise.

### GetNetworkingTypeOk

`func (o *ClusterEndpoint) GetNetworkingTypeOk() (*string, bool)`

GetNetworkingTypeOk returns a tuple with the NetworkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingType

`func (o *ClusterEndpoint) SetNetworkingType(v string)`

SetNetworkingType sets NetworkingType field to given value.

### HasNetworkingType

`func (o *ClusterEndpoint) HasNetworkingType() bool`

HasNetworkingType returns a boolean if a field has been set.

### GetOpenPorts

`func (o *ClusterEndpoint) GetOpenPorts() []int64`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *ClusterEndpoint) GetOpenPortsOk() (*[]int64, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *ClusterEndpoint) SetOpenPorts(v []int64)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *ClusterEndpoint) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetPrimary

`func (o *ClusterEndpoint) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *ClusterEndpoint) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *ClusterEndpoint) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *ClusterEndpoint) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


