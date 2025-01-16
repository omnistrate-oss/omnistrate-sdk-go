# CreateUpgradePathRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the upgrade path | 
**Name** | **string** | Name of the upgrade path | 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**SourceVersion** | **string** | Version of the Entity to operate on | 
**TargetVersion** | **string** | Version of the Entity to operate on | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateUpgradePathRequest

`func NewCreateUpgradePathRequest(description string, name string, serviceEnvironmentId string, sourceVersion string, targetVersion string, token string, ) *CreateUpgradePathRequest`

NewCreateUpgradePathRequest instantiates a new CreateUpgradePathRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpgradePathRequestWithDefaults

`func NewCreateUpgradePathRequestWithDefaults() *CreateUpgradePathRequest`

NewCreateUpgradePathRequestWithDefaults instantiates a new CreateUpgradePathRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateUpgradePathRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUpgradePathRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUpgradePathRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateUpgradePathRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUpgradePathRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUpgradePathRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServiceEnvironmentId

`func (o *CreateUpgradePathRequest) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *CreateUpgradePathRequest) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *CreateUpgradePathRequest) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceId

`func (o *CreateUpgradePathRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateUpgradePathRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateUpgradePathRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *CreateUpgradePathRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSourceVersion

`func (o *CreateUpgradePathRequest) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *CreateUpgradePathRequest) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *CreateUpgradePathRequest) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.


### GetTargetVersion

`func (o *CreateUpgradePathRequest) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *CreateUpgradePathRequest) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *CreateUpgradePathRequest) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.


### GetToken

`func (o *CreateUpgradePathRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateUpgradePathRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateUpgradePathRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


