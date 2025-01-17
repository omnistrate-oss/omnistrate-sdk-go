# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | ID of a Service Environment | [optional] 
**InstanceCount** | Pointer to **int64** | The number of instances for the specific resource. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**PerVersionInstanceCount** | Pointer to **map[string]int64** | The per-version instance count for the resource. | [optional] 
**ProductTierType** | Pointer to **string** | The product tier type | [optional] 
**ProxyType** | Pointer to **string** | The proxy type of instance | [optional] 
**ResourceId** | Pointer to **string** | ID of a resource | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**ServiceModelType** | Pointer to **string** | The service model type | [optional] 
**Version** | Pointer to **string** | The latest version of the resource. | [optional] 
**VersionHistory** | Pointer to **[]string** | The version history of the resource. | [optional] 

## Methods

### NewResource

`func NewResource() *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *Resource) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Resource) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Resource) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *Resource) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetInstanceCount

`func (o *Resource) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *Resource) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *Resource) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *Resource) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Resource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPerVersionInstanceCount

`func (o *Resource) GetPerVersionInstanceCount() map[string]int64`

GetPerVersionInstanceCount returns the PerVersionInstanceCount field if non-nil, zero value otherwise.

### GetPerVersionInstanceCountOk

`func (o *Resource) GetPerVersionInstanceCountOk() (*map[string]int64, bool)`

GetPerVersionInstanceCountOk returns a tuple with the PerVersionInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerVersionInstanceCount

`func (o *Resource) SetPerVersionInstanceCount(v map[string]int64)`

SetPerVersionInstanceCount sets PerVersionInstanceCount field to given value.

### HasPerVersionInstanceCount

`func (o *Resource) HasPerVersionInstanceCount() bool`

HasPerVersionInstanceCount returns a boolean if a field has been set.

### GetProductTierType

`func (o *Resource) GetProductTierType() string`

GetProductTierType returns the ProductTierType field if non-nil, zero value otherwise.

### GetProductTierTypeOk

`func (o *Resource) GetProductTierTypeOk() (*string, bool)`

GetProductTierTypeOk returns a tuple with the ProductTierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierType

`func (o *Resource) SetProductTierType(v string)`

SetProductTierType sets ProductTierType field to given value.

### HasProductTierType

`func (o *Resource) HasProductTierType() bool`

HasProductTierType returns a boolean if a field has been set.

### GetProxyType

`func (o *Resource) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *Resource) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *Resource) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *Resource) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetResourceId

`func (o *Resource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Resource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Resource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Resource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetServiceId

`func (o *Resource) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Resource) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Resource) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Resource) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceModelType

`func (o *Resource) GetServiceModelType() string`

GetServiceModelType returns the ServiceModelType field if non-nil, zero value otherwise.

### GetServiceModelTypeOk

`func (o *Resource) GetServiceModelTypeOk() (*string, bool)`

GetServiceModelTypeOk returns a tuple with the ServiceModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelType

`func (o *Resource) SetServiceModelType(v string)`

SetServiceModelType sets ServiceModelType field to given value.

### HasServiceModelType

`func (o *Resource) HasServiceModelType() bool`

HasServiceModelType returns a boolean if a field has been set.

### GetVersion

`func (o *Resource) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Resource) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Resource) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Resource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionHistory

`func (o *Resource) GetVersionHistory() []string`

GetVersionHistory returns the VersionHistory field if non-nil, zero value otherwise.

### GetVersionHistoryOk

`func (o *Resource) GetVersionHistoryOk() (*[]string, bool)`

GetVersionHistoryOk returns a tuple with the VersionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHistory

`func (o *Resource) SetVersionHistory(v []string)`

SetVersionHistory sets VersionHistory field to given value.

### HasVersionHistory

`func (o *Resource) HasVersionHistory() bool`

HasVersionHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


