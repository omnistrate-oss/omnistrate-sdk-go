# CreateTierVersionSetRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the product-tier version set. | [optional] 
**Name** | Pointer to **string** | The name of the product-tier version set. | [optional] 
**ParentVersion** | Pointer to **string** | The parent version of this version set. | [optional] 
**Type** | **string** | The version-set type of the product-tier. | 

## Methods

### NewCreateTierVersionSetRequestBody

`func NewCreateTierVersionSetRequestBody(type_ string, ) *CreateTierVersionSetRequestBody`

NewCreateTierVersionSetRequestBody instantiates a new CreateTierVersionSetRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTierVersionSetRequestBodyWithDefaults

`func NewCreateTierVersionSetRequestBodyWithDefaults() *CreateTierVersionSetRequestBody`

NewCreateTierVersionSetRequestBodyWithDefaults instantiates a new CreateTierVersionSetRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateTierVersionSetRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTierVersionSetRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTierVersionSetRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTierVersionSetRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateTierVersionSetRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTierVersionSetRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTierVersionSetRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTierVersionSetRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentVersion

`func (o *CreateTierVersionSetRequestBody) GetParentVersion() string`

GetParentVersion returns the ParentVersion field if non-nil, zero value otherwise.

### GetParentVersionOk

`func (o *CreateTierVersionSetRequestBody) GetParentVersionOk() (*string, bool)`

GetParentVersionOk returns a tuple with the ParentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVersion

`func (o *CreateTierVersionSetRequestBody) SetParentVersion(v string)`

SetParentVersion sets ParentVersion field to given value.

### HasParentVersion

`func (o *CreateTierVersionSetRequestBody) HasParentVersion() bool`

HasParentVersion returns a boolean if a field has been set.

### GetType

`func (o *CreateTierVersionSetRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTierVersionSetRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTierVersionSetRequestBody) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


