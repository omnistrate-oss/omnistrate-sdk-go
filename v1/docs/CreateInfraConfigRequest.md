# CreateInfraConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeConfigId** | Pointer to **string** | ID of a Compute Config | [optional] 
**CustomTag** | Pointer to [**CustomTag**](CustomTag.md) |  | [optional] 
**Description** | **string** | The description for the infra config | 
**Name** | **string** | The name of the infra config | 
**NetworkConfigId** | Pointer to **string** | ID of a Network Config | [optional] 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceId** | **string** | ID of a Service | 
**StorageConfigId** | Pointer to **string** | ID of a Storage Config | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateInfraConfigRequest

`func NewCreateInfraConfigRequest(description string, name string, serviceEnvironmentId string, serviceId string, token string, ) *CreateInfraConfigRequest`

NewCreateInfraConfigRequest instantiates a new CreateInfraConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInfraConfigRequestWithDefaults

`func NewCreateInfraConfigRequestWithDefaults() *CreateInfraConfigRequest`

NewCreateInfraConfigRequestWithDefaults instantiates a new CreateInfraConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeConfigId

`func (o *CreateInfraConfigRequest) GetComputeConfigId() string`

GetComputeConfigId returns the ComputeConfigId field if non-nil, zero value otherwise.

### GetComputeConfigIdOk

`func (o *CreateInfraConfigRequest) GetComputeConfigIdOk() (*string, bool)`

GetComputeConfigIdOk returns a tuple with the ComputeConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeConfigId

`func (o *CreateInfraConfigRequest) SetComputeConfigId(v string)`

SetComputeConfigId sets ComputeConfigId field to given value.

### HasComputeConfigId

`func (o *CreateInfraConfigRequest) HasComputeConfigId() bool`

HasComputeConfigId returns a boolean if a field has been set.

### GetCustomTag

`func (o *CreateInfraConfigRequest) GetCustomTag() CustomTag`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *CreateInfraConfigRequest) GetCustomTagOk() (*CustomTag, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *CreateInfraConfigRequest) SetCustomTag(v CustomTag)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *CreateInfraConfigRequest) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.

### GetDescription

`func (o *CreateInfraConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInfraConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInfraConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateInfraConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInfraConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInfraConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkConfigId

`func (o *CreateInfraConfigRequest) GetNetworkConfigId() string`

GetNetworkConfigId returns the NetworkConfigId field if non-nil, zero value otherwise.

### GetNetworkConfigIdOk

`func (o *CreateInfraConfigRequest) GetNetworkConfigIdOk() (*string, bool)`

GetNetworkConfigIdOk returns a tuple with the NetworkConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfigId

`func (o *CreateInfraConfigRequest) SetNetworkConfigId(v string)`

SetNetworkConfigId sets NetworkConfigId field to given value.

### HasNetworkConfigId

`func (o *CreateInfraConfigRequest) HasNetworkConfigId() bool`

HasNetworkConfigId returns a boolean if a field has been set.

### GetServiceEnvironmentId

`func (o *CreateInfraConfigRequest) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *CreateInfraConfigRequest) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *CreateInfraConfigRequest) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceId

`func (o *CreateInfraConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateInfraConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateInfraConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStorageConfigId

`func (o *CreateInfraConfigRequest) GetStorageConfigId() string`

GetStorageConfigId returns the StorageConfigId field if non-nil, zero value otherwise.

### GetStorageConfigIdOk

`func (o *CreateInfraConfigRequest) GetStorageConfigIdOk() (*string, bool)`

GetStorageConfigIdOk returns a tuple with the StorageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigId

`func (o *CreateInfraConfigRequest) SetStorageConfigId(v string)`

SetStorageConfigId sets StorageConfigId field to given value.

### HasStorageConfigId

`func (o *CreateInfraConfigRequest) HasStorageConfigId() bool`

HasStorageConfigId returns a boolean if a field has been set.

### GetToken

`func (o *CreateInfraConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateInfraConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateInfraConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


