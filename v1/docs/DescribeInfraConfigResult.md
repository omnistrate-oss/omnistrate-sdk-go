# DescribeInfraConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeConfigId** | Pointer to **string** | The compute config ID | [optional] 
**CustomTag** | Pointer to [**EnvironmentVariable**](EnvironmentVariable.md) |  | [optional] 
**Description** | **string** | The description for the infra config | 
**Id** | **string** | Infra Config ID to operate on | 
**Name** | **string** | The name of the infra config | 
**NetworkConfigId** | Pointer to **string** | The network config ID | [optional] 
**ServiceEnvironmentId** | **string** | The service environment ID | 
**ServiceId** | **string** | The service ID | 
**StorageConfigId** | Pointer to **string** | The storage config ID | [optional] 

## Methods

### NewDescribeInfraConfigResult

`func NewDescribeInfraConfigResult(description string, id string, name string, serviceEnvironmentId string, serviceId string, ) *DescribeInfraConfigResult`

NewDescribeInfraConfigResult instantiates a new DescribeInfraConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeInfraConfigResultWithDefaults

`func NewDescribeInfraConfigResultWithDefaults() *DescribeInfraConfigResult`

NewDescribeInfraConfigResultWithDefaults instantiates a new DescribeInfraConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeConfigId

`func (o *DescribeInfraConfigResult) GetComputeConfigId() string`

GetComputeConfigId returns the ComputeConfigId field if non-nil, zero value otherwise.

### GetComputeConfigIdOk

`func (o *DescribeInfraConfigResult) GetComputeConfigIdOk() (*string, bool)`

GetComputeConfigIdOk returns a tuple with the ComputeConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeConfigId

`func (o *DescribeInfraConfigResult) SetComputeConfigId(v string)`

SetComputeConfigId sets ComputeConfigId field to given value.

### HasComputeConfigId

`func (o *DescribeInfraConfigResult) HasComputeConfigId() bool`

HasComputeConfigId returns a boolean if a field has been set.

### GetCustomTag

`func (o *DescribeInfraConfigResult) GetCustomTag() EnvironmentVariable`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *DescribeInfraConfigResult) GetCustomTagOk() (*EnvironmentVariable, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *DescribeInfraConfigResult) SetCustomTag(v EnvironmentVariable)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *DescribeInfraConfigResult) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.

### GetDescription

`func (o *DescribeInfraConfigResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeInfraConfigResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeInfraConfigResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeInfraConfigResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeInfraConfigResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeInfraConfigResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DescribeInfraConfigResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeInfraConfigResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeInfraConfigResult) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkConfigId

`func (o *DescribeInfraConfigResult) GetNetworkConfigId() string`

GetNetworkConfigId returns the NetworkConfigId field if non-nil, zero value otherwise.

### GetNetworkConfigIdOk

`func (o *DescribeInfraConfigResult) GetNetworkConfigIdOk() (*string, bool)`

GetNetworkConfigIdOk returns a tuple with the NetworkConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfigId

`func (o *DescribeInfraConfigResult) SetNetworkConfigId(v string)`

SetNetworkConfigId sets NetworkConfigId field to given value.

### HasNetworkConfigId

`func (o *DescribeInfraConfigResult) HasNetworkConfigId() bool`

HasNetworkConfigId returns a boolean if a field has been set.

### GetServiceEnvironmentId

`func (o *DescribeInfraConfigResult) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *DescribeInfraConfigResult) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *DescribeInfraConfigResult) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceId

`func (o *DescribeInfraConfigResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeInfraConfigResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeInfraConfigResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStorageConfigId

`func (o *DescribeInfraConfigResult) GetStorageConfigId() string`

GetStorageConfigId returns the StorageConfigId field if non-nil, zero value otherwise.

### GetStorageConfigIdOk

`func (o *DescribeInfraConfigResult) GetStorageConfigIdOk() (*string, bool)`

GetStorageConfigIdOk returns a tuple with the StorageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigId

`func (o *DescribeInfraConfigResult) SetStorageConfigId(v string)`

SetStorageConfigId sets StorageConfigId field to given value.

### HasStorageConfigId

`func (o *DescribeInfraConfigResult) HasStorageConfigId() bool`

HasStorageConfigId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


