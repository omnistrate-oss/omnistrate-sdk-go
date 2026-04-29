# UpdateAPIKeyMetadataRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | New description. Optional; nil leaves the field unchanged. | [optional] 
**Name** | Pointer to **string** | New display name. Optional; nil leaves the field unchanged. Re-validated for uniqueness within the org on change (409 on conflict). | [optional] 

## Methods

### NewUpdateAPIKeyMetadataRequest2

`func NewUpdateAPIKeyMetadataRequest2() *UpdateAPIKeyMetadataRequest2`

NewUpdateAPIKeyMetadataRequest2 instantiates a new UpdateAPIKeyMetadataRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAPIKeyMetadataRequest2WithDefaults

`func NewUpdateAPIKeyMetadataRequest2WithDefaults() *UpdateAPIKeyMetadataRequest2`

NewUpdateAPIKeyMetadataRequest2WithDefaults instantiates a new UpdateAPIKeyMetadataRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateAPIKeyMetadataRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAPIKeyMetadataRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAPIKeyMetadataRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAPIKeyMetadataRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateAPIKeyMetadataRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAPIKeyMetadataRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAPIKeyMetadataRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAPIKeyMetadataRequest2) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


