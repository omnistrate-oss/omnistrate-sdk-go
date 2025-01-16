# UpdateInfraConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeConfigId** | Pointer to **string** | ID of a Compute Config | [optional] 
**CustomTag** | Pointer to [**CustomTag**](CustomTag.md) |  | [optional] 
**Description** | Pointer to **string** | The description for the infra config | [optional] 
**Id** | **string** | ID of an Infra Config | 
**Name** | Pointer to **string** | The name of the infra config | [optional] 
**NetworkConfigId** | Pointer to **string** | ID of a Network Config | [optional] 
**ServiceId** | **string** | ID of a Service | 
**StorageConfigId** | Pointer to **string** | ID of a Storage Config | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateInfraConfigRequest

`func NewUpdateInfraConfigRequest(id string, serviceId string, token string, ) *UpdateInfraConfigRequest`

NewUpdateInfraConfigRequest instantiates a new UpdateInfraConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInfraConfigRequestWithDefaults

`func NewUpdateInfraConfigRequestWithDefaults() *UpdateInfraConfigRequest`

NewUpdateInfraConfigRequestWithDefaults instantiates a new UpdateInfraConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeConfigId

`func (o *UpdateInfraConfigRequest) GetComputeConfigId() string`

GetComputeConfigId returns the ComputeConfigId field if non-nil, zero value otherwise.

### GetComputeConfigIdOk

`func (o *UpdateInfraConfigRequest) GetComputeConfigIdOk() (*string, bool)`

GetComputeConfigIdOk returns a tuple with the ComputeConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeConfigId

`func (o *UpdateInfraConfigRequest) SetComputeConfigId(v string)`

SetComputeConfigId sets ComputeConfigId field to given value.

### HasComputeConfigId

`func (o *UpdateInfraConfigRequest) HasComputeConfigId() bool`

HasComputeConfigId returns a boolean if a field has been set.

### GetCustomTag

`func (o *UpdateInfraConfigRequest) GetCustomTag() CustomTag`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *UpdateInfraConfigRequest) GetCustomTagOk() (*CustomTag, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *UpdateInfraConfigRequest) SetCustomTag(v CustomTag)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *UpdateInfraConfigRequest) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInfraConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInfraConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInfraConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInfraConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UpdateInfraConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInfraConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInfraConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateInfraConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInfraConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInfraConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInfraConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkConfigId

`func (o *UpdateInfraConfigRequest) GetNetworkConfigId() string`

GetNetworkConfigId returns the NetworkConfigId field if non-nil, zero value otherwise.

### GetNetworkConfigIdOk

`func (o *UpdateInfraConfigRequest) GetNetworkConfigIdOk() (*string, bool)`

GetNetworkConfigIdOk returns a tuple with the NetworkConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfigId

`func (o *UpdateInfraConfigRequest) SetNetworkConfigId(v string)`

SetNetworkConfigId sets NetworkConfigId field to given value.

### HasNetworkConfigId

`func (o *UpdateInfraConfigRequest) HasNetworkConfigId() bool`

HasNetworkConfigId returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateInfraConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateInfraConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateInfraConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStorageConfigId

`func (o *UpdateInfraConfigRequest) GetStorageConfigId() string`

GetStorageConfigId returns the StorageConfigId field if non-nil, zero value otherwise.

### GetStorageConfigIdOk

`func (o *UpdateInfraConfigRequest) GetStorageConfigIdOk() (*string, bool)`

GetStorageConfigIdOk returns a tuple with the StorageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigId

`func (o *UpdateInfraConfigRequest) SetStorageConfigId(v string)`

SetStorageConfigId sets StorageConfigId field to given value.

### HasStorageConfigId

`func (o *UpdateInfraConfigRequest) HasStorageConfigId() bool`

HasStorageConfigId returns a boolean if a field has been set.

### GetToken

`func (o *UpdateInfraConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateInfraConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateInfraConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


