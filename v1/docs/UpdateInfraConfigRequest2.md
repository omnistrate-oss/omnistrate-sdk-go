# UpdateInfraConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeConfigId** | Pointer to **string** | The compute config ID | [optional] 
**CustomTag** | Pointer to [**CustomTag**](CustomTag.md) |  | [optional] 
**Description** | Pointer to **string** | The description for the infra config | [optional] 
**Name** | Pointer to **string** | The name of the infra config | [optional] 
**NetworkConfigId** | Pointer to **string** | The network config ID | [optional] 
**StorageConfigId** | Pointer to **string** | The storage config ID per compute node | [optional] 

## Methods

### NewUpdateInfraConfigRequest2

`func NewUpdateInfraConfigRequest2() *UpdateInfraConfigRequest2`

NewUpdateInfraConfigRequest2 instantiates a new UpdateInfraConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInfraConfigRequest2WithDefaults

`func NewUpdateInfraConfigRequest2WithDefaults() *UpdateInfraConfigRequest2`

NewUpdateInfraConfigRequest2WithDefaults instantiates a new UpdateInfraConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeConfigId

`func (o *UpdateInfraConfigRequest2) GetComputeConfigId() string`

GetComputeConfigId returns the ComputeConfigId field if non-nil, zero value otherwise.

### GetComputeConfigIdOk

`func (o *UpdateInfraConfigRequest2) GetComputeConfigIdOk() (*string, bool)`

GetComputeConfigIdOk returns a tuple with the ComputeConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeConfigId

`func (o *UpdateInfraConfigRequest2) SetComputeConfigId(v string)`

SetComputeConfigId sets ComputeConfigId field to given value.

### HasComputeConfigId

`func (o *UpdateInfraConfigRequest2) HasComputeConfigId() bool`

HasComputeConfigId returns a boolean if a field has been set.

### GetCustomTag

`func (o *UpdateInfraConfigRequest2) GetCustomTag() CustomTag`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *UpdateInfraConfigRequest2) GetCustomTagOk() (*CustomTag, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *UpdateInfraConfigRequest2) SetCustomTag(v CustomTag)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *UpdateInfraConfigRequest2) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInfraConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInfraConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInfraConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInfraConfigRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateInfraConfigRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInfraConfigRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInfraConfigRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInfraConfigRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkConfigId

`func (o *UpdateInfraConfigRequest2) GetNetworkConfigId() string`

GetNetworkConfigId returns the NetworkConfigId field if non-nil, zero value otherwise.

### GetNetworkConfigIdOk

`func (o *UpdateInfraConfigRequest2) GetNetworkConfigIdOk() (*string, bool)`

GetNetworkConfigIdOk returns a tuple with the NetworkConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfigId

`func (o *UpdateInfraConfigRequest2) SetNetworkConfigId(v string)`

SetNetworkConfigId sets NetworkConfigId field to given value.

### HasNetworkConfigId

`func (o *UpdateInfraConfigRequest2) HasNetworkConfigId() bool`

HasNetworkConfigId returns a boolean if a field has been set.

### GetStorageConfigId

`func (o *UpdateInfraConfigRequest2) GetStorageConfigId() string`

GetStorageConfigId returns the StorageConfigId field if non-nil, zero value otherwise.

### GetStorageConfigIdOk

`func (o *UpdateInfraConfigRequest2) GetStorageConfigIdOk() (*string, bool)`

GetStorageConfigIdOk returns a tuple with the StorageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigId

`func (o *UpdateInfraConfigRequest2) SetStorageConfigId(v string)`

SetStorageConfigId sets StorageConfigId field to given value.

### HasStorageConfigId

`func (o *UpdateInfraConfigRequest2) HasStorageConfigId() bool`

HasStorageConfigId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


