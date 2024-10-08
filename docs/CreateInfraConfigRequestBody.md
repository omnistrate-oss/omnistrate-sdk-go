# CreateInfraConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeConfigId** | Pointer to **string** | The compute config ID | [optional] 
**CustomTag** | Pointer to [**EnvironmentVariable**](EnvironmentVariable.md) |  | [optional] 
**Description** | **string** | The description for the infra config | 
**Name** | **string** | The name of the infra config | 
**NetworkConfigId** | Pointer to **string** | The network config ID | [optional] 
**ServiceEnvironmentId** | **string** | The service environment ID | 
**StorageConfigId** | Pointer to **string** | The storage config ID | [optional] 

## Methods

### NewCreateInfraConfigRequestBody

`func NewCreateInfraConfigRequestBody(description string, name string, serviceEnvironmentId string, ) *CreateInfraConfigRequestBody`

NewCreateInfraConfigRequestBody instantiates a new CreateInfraConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInfraConfigRequestBodyWithDefaults

`func NewCreateInfraConfigRequestBodyWithDefaults() *CreateInfraConfigRequestBody`

NewCreateInfraConfigRequestBodyWithDefaults instantiates a new CreateInfraConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeConfigId

`func (o *CreateInfraConfigRequestBody) GetComputeConfigId() string`

GetComputeConfigId returns the ComputeConfigId field if non-nil, zero value otherwise.

### GetComputeConfigIdOk

`func (o *CreateInfraConfigRequestBody) GetComputeConfigIdOk() (*string, bool)`

GetComputeConfigIdOk returns a tuple with the ComputeConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeConfigId

`func (o *CreateInfraConfigRequestBody) SetComputeConfigId(v string)`

SetComputeConfigId sets ComputeConfigId field to given value.

### HasComputeConfigId

`func (o *CreateInfraConfigRequestBody) HasComputeConfigId() bool`

HasComputeConfigId returns a boolean if a field has been set.

### GetCustomTag

`func (o *CreateInfraConfigRequestBody) GetCustomTag() EnvironmentVariable`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *CreateInfraConfigRequestBody) GetCustomTagOk() (*EnvironmentVariable, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *CreateInfraConfigRequestBody) SetCustomTag(v EnvironmentVariable)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *CreateInfraConfigRequestBody) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.

### GetDescription

`func (o *CreateInfraConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInfraConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInfraConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateInfraConfigRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInfraConfigRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInfraConfigRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkConfigId

`func (o *CreateInfraConfigRequestBody) GetNetworkConfigId() string`

GetNetworkConfigId returns the NetworkConfigId field if non-nil, zero value otherwise.

### GetNetworkConfigIdOk

`func (o *CreateInfraConfigRequestBody) GetNetworkConfigIdOk() (*string, bool)`

GetNetworkConfigIdOk returns a tuple with the NetworkConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfigId

`func (o *CreateInfraConfigRequestBody) SetNetworkConfigId(v string)`

SetNetworkConfigId sets NetworkConfigId field to given value.

### HasNetworkConfigId

`func (o *CreateInfraConfigRequestBody) HasNetworkConfigId() bool`

HasNetworkConfigId returns a boolean if a field has been set.

### GetServiceEnvironmentId

`func (o *CreateInfraConfigRequestBody) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *CreateInfraConfigRequestBody) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *CreateInfraConfigRequestBody) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetStorageConfigId

`func (o *CreateInfraConfigRequestBody) GetStorageConfigId() string`

GetStorageConfigId returns the StorageConfigId field if non-nil, zero value otherwise.

### GetStorageConfigIdOk

`func (o *CreateInfraConfigRequestBody) GetStorageConfigIdOk() (*string, bool)`

GetStorageConfigIdOk returns a tuple with the StorageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigId

`func (o *CreateInfraConfigRequestBody) SetStorageConfigId(v string)`

SetStorageConfigId sets StorageConfigId field to given value.

### HasStorageConfigId

`func (o *CreateInfraConfigRequestBody) HasStorageConfigId() bool`

HasStorageConfigId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


