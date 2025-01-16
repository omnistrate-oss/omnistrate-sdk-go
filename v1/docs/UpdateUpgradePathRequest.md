# UpdateUpgradePathRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the upgrade path | [optional] 
**Id** | Pointer to **string** | ID of an Upgrade Path | [optional] 
**Name** | Pointer to **string** | Name of the upgrade path | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**SourceVersion** | Pointer to **string** | Version of the Entity to operate on | [optional] 
**TargetVersion** | Pointer to **string** | Version of the Entity to operate on | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateUpgradePathRequest

`func NewUpdateUpgradePathRequest(token string, ) *UpdateUpgradePathRequest`

NewUpdateUpgradePathRequest instantiates a new UpdateUpgradePathRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUpgradePathRequestWithDefaults

`func NewUpdateUpgradePathRequestWithDefaults() *UpdateUpgradePathRequest`

NewUpdateUpgradePathRequestWithDefaults instantiates a new UpdateUpgradePathRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateUpgradePathRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateUpgradePathRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateUpgradePathRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateUpgradePathRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UpdateUpgradePathRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateUpgradePathRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateUpgradePathRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateUpgradePathRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateUpgradePathRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUpgradePathRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUpgradePathRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUpgradePathRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateUpgradePathRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateUpgradePathRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateUpgradePathRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *UpdateUpgradePathRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSourceVersion

`func (o *UpdateUpgradePathRequest) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *UpdateUpgradePathRequest) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *UpdateUpgradePathRequest) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.

### HasSourceVersion

`func (o *UpdateUpgradePathRequest) HasSourceVersion() bool`

HasSourceVersion returns a boolean if a field has been set.

### GetTargetVersion

`func (o *UpdateUpgradePathRequest) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *UpdateUpgradePathRequest) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *UpdateUpgradePathRequest) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *UpdateUpgradePathRequest) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetToken

`func (o *UpdateUpgradePathRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateUpgradePathRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateUpgradePathRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


