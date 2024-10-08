# UpdateInfraConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeConfigId** | Pointer to **string** | The compute config ID | [optional] 
**CustomTag** | Pointer to [**EnvironmentVariable**](EnvironmentVariable.md) |  | [optional] 
**Description** | Pointer to **string** | The description for the infra config | [optional] 
**Name** | Pointer to **string** | The name of the infra config | [optional] 
**NetworkConfigId** | Pointer to **string** | The network config ID | [optional] 
**StorageConfigId** | Pointer to **string** | The storage config ID per compute node | [optional] 

## Methods

### NewUpdateInfraConfigRequestBody

`func NewUpdateInfraConfigRequestBody() *UpdateInfraConfigRequestBody`

NewUpdateInfraConfigRequestBody instantiates a new UpdateInfraConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInfraConfigRequestBodyWithDefaults

`func NewUpdateInfraConfigRequestBodyWithDefaults() *UpdateInfraConfigRequestBody`

NewUpdateInfraConfigRequestBodyWithDefaults instantiates a new UpdateInfraConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeConfigId

`func (o *UpdateInfraConfigRequestBody) GetComputeConfigId() string`

GetComputeConfigId returns the ComputeConfigId field if non-nil, zero value otherwise.

### GetComputeConfigIdOk

`func (o *UpdateInfraConfigRequestBody) GetComputeConfigIdOk() (*string, bool)`

GetComputeConfigIdOk returns a tuple with the ComputeConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeConfigId

`func (o *UpdateInfraConfigRequestBody) SetComputeConfigId(v string)`

SetComputeConfigId sets ComputeConfigId field to given value.

### HasComputeConfigId

`func (o *UpdateInfraConfigRequestBody) HasComputeConfigId() bool`

HasComputeConfigId returns a boolean if a field has been set.

### GetCustomTag

`func (o *UpdateInfraConfigRequestBody) GetCustomTag() EnvironmentVariable`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *UpdateInfraConfigRequestBody) GetCustomTagOk() (*EnvironmentVariable, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *UpdateInfraConfigRequestBody) SetCustomTag(v EnvironmentVariable)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *UpdateInfraConfigRequestBody) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInfraConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInfraConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInfraConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInfraConfigRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateInfraConfigRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInfraConfigRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInfraConfigRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInfraConfigRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkConfigId

`func (o *UpdateInfraConfigRequestBody) GetNetworkConfigId() string`

GetNetworkConfigId returns the NetworkConfigId field if non-nil, zero value otherwise.

### GetNetworkConfigIdOk

`func (o *UpdateInfraConfigRequestBody) GetNetworkConfigIdOk() (*string, bool)`

GetNetworkConfigIdOk returns a tuple with the NetworkConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfigId

`func (o *UpdateInfraConfigRequestBody) SetNetworkConfigId(v string)`

SetNetworkConfigId sets NetworkConfigId field to given value.

### HasNetworkConfigId

`func (o *UpdateInfraConfigRequestBody) HasNetworkConfigId() bool`

HasNetworkConfigId returns a boolean if a field has been set.

### GetStorageConfigId

`func (o *UpdateInfraConfigRequestBody) GetStorageConfigId() string`

GetStorageConfigId returns the StorageConfigId field if non-nil, zero value otherwise.

### GetStorageConfigIdOk

`func (o *UpdateInfraConfigRequestBody) GetStorageConfigIdOk() (*string, bool)`

GetStorageConfigIdOk returns a tuple with the StorageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigId

`func (o *UpdateInfraConfigRequestBody) SetStorageConfigId(v string)`

SetStorageConfigId sets StorageConfigId field to given value.

### HasStorageConfigId

`func (o *UpdateInfraConfigRequestBody) HasStorageConfigId() bool`

HasStorageConfigId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


