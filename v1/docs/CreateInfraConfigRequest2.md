# CreateInfraConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeConfigId** | Pointer to **string** | The compute config ID | [optional] 
**CustomTag** | Pointer to [**CustomTag**](CustomTag.md) |  | [optional] 
**Description** | **string** | The description for the infra config | 
**Name** | **string** | The name of the infra config | 
**NetworkConfigId** | Pointer to **string** | The network config ID | [optional] 
**ServiceEnvironmentId** | **string** | The service environment ID | 
**StorageConfigId** | Pointer to **string** | The storage config ID | [optional] 

## Methods

### NewCreateInfraConfigRequest2

`func NewCreateInfraConfigRequest2(description string, name string, serviceEnvironmentId string, ) *CreateInfraConfigRequest2`

NewCreateInfraConfigRequest2 instantiates a new CreateInfraConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInfraConfigRequest2WithDefaults

`func NewCreateInfraConfigRequest2WithDefaults() *CreateInfraConfigRequest2`

NewCreateInfraConfigRequest2WithDefaults instantiates a new CreateInfraConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeConfigId

`func (o *CreateInfraConfigRequest2) GetComputeConfigId() string`

GetComputeConfigId returns the ComputeConfigId field if non-nil, zero value otherwise.

### GetComputeConfigIdOk

`func (o *CreateInfraConfigRequest2) GetComputeConfigIdOk() (*string, bool)`

GetComputeConfigIdOk returns a tuple with the ComputeConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeConfigId

`func (o *CreateInfraConfigRequest2) SetComputeConfigId(v string)`

SetComputeConfigId sets ComputeConfigId field to given value.

### HasComputeConfigId

`func (o *CreateInfraConfigRequest2) HasComputeConfigId() bool`

HasComputeConfigId returns a boolean if a field has been set.

### GetCustomTag

`func (o *CreateInfraConfigRequest2) GetCustomTag() CustomTag`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *CreateInfraConfigRequest2) GetCustomTagOk() (*CustomTag, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *CreateInfraConfigRequest2) SetCustomTag(v CustomTag)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *CreateInfraConfigRequest2) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.

### GetDescription

`func (o *CreateInfraConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInfraConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInfraConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateInfraConfigRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInfraConfigRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInfraConfigRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkConfigId

`func (o *CreateInfraConfigRequest2) GetNetworkConfigId() string`

GetNetworkConfigId returns the NetworkConfigId field if non-nil, zero value otherwise.

### GetNetworkConfigIdOk

`func (o *CreateInfraConfigRequest2) GetNetworkConfigIdOk() (*string, bool)`

GetNetworkConfigIdOk returns a tuple with the NetworkConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfigId

`func (o *CreateInfraConfigRequest2) SetNetworkConfigId(v string)`

SetNetworkConfigId sets NetworkConfigId field to given value.

### HasNetworkConfigId

`func (o *CreateInfraConfigRequest2) HasNetworkConfigId() bool`

HasNetworkConfigId returns a boolean if a field has been set.

### GetServiceEnvironmentId

`func (o *CreateInfraConfigRequest2) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *CreateInfraConfigRequest2) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *CreateInfraConfigRequest2) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetStorageConfigId

`func (o *CreateInfraConfigRequest2) GetStorageConfigId() string`

GetStorageConfigId returns the StorageConfigId field if non-nil, zero value otherwise.

### GetStorageConfigIdOk

`func (o *CreateInfraConfigRequest2) GetStorageConfigIdOk() (*string, bool)`

GetStorageConfigIdOk returns a tuple with the StorageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigId

`func (o *CreateInfraConfigRequest2) SetStorageConfigId(v string)`

SetStorageConfigId sets StorageConfigId field to given value.

### HasStorageConfigId

`func (o *CreateInfraConfigRequest2) HasStorageConfigId() bool`

HasStorageConfigId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


