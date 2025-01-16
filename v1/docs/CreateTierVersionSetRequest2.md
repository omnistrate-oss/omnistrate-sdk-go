# CreateTierVersionSetRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the product-tier version set. | [optional] 
**Name** | Pointer to **string** | The name of the product-tier version set. | [optional] 
**ParentVersion** | Pointer to **string** | The parent version of this version set. | [optional] 
**Type** | **string** | The version-set type of the product-tier. | 

## Methods

### NewCreateTierVersionSetRequest2

`func NewCreateTierVersionSetRequest2(type_ string, ) *CreateTierVersionSetRequest2`

NewCreateTierVersionSetRequest2 instantiates a new CreateTierVersionSetRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTierVersionSetRequest2WithDefaults

`func NewCreateTierVersionSetRequest2WithDefaults() *CreateTierVersionSetRequest2`

NewCreateTierVersionSetRequest2WithDefaults instantiates a new CreateTierVersionSetRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateTierVersionSetRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTierVersionSetRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTierVersionSetRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTierVersionSetRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateTierVersionSetRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTierVersionSetRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTierVersionSetRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTierVersionSetRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentVersion

`func (o *CreateTierVersionSetRequest2) GetParentVersion() string`

GetParentVersion returns the ParentVersion field if non-nil, zero value otherwise.

### GetParentVersionOk

`func (o *CreateTierVersionSetRequest2) GetParentVersionOk() (*string, bool)`

GetParentVersionOk returns a tuple with the ParentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVersion

`func (o *CreateTierVersionSetRequest2) SetParentVersion(v string)`

SetParentVersion sets ParentVersion field to given value.

### HasParentVersion

`func (o *CreateTierVersionSetRequest2) HasParentVersion() bool`

HasParentVersion returns a boolean if a field has been set.

### GetType

`func (o *CreateTierVersionSetRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTierVersionSetRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTierVersionSetRequest2) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


