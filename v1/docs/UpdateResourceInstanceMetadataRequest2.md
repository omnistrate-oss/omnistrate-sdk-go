# UpdateResourceInstanceMetadataRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTags** | Pointer to [**[]CustomTag**](CustomTag.md) | The custom tag for the resource instance | [optional] 
**DeletionProtection** | Pointer to **bool** | Set to true to enable deletion protection or false to disable it | [optional] 

## Methods

### NewUpdateResourceInstanceMetadataRequest2

`func NewUpdateResourceInstanceMetadataRequest2() *UpdateResourceInstanceMetadataRequest2`

NewUpdateResourceInstanceMetadataRequest2 instantiates a new UpdateResourceInstanceMetadataRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceInstanceMetadataRequest2WithDefaults

`func NewUpdateResourceInstanceMetadataRequest2WithDefaults() *UpdateResourceInstanceMetadataRequest2`

NewUpdateResourceInstanceMetadataRequest2WithDefaults instantiates a new UpdateResourceInstanceMetadataRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTags

`func (o *UpdateResourceInstanceMetadataRequest2) GetCustomTags() []CustomTag`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *UpdateResourceInstanceMetadataRequest2) GetCustomTagsOk() (*[]CustomTag, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *UpdateResourceInstanceMetadataRequest2) SetCustomTags(v []CustomTag)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *UpdateResourceInstanceMetadataRequest2) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### GetDeletionProtection

`func (o *UpdateResourceInstanceMetadataRequest2) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *UpdateResourceInstanceMetadataRequest2) GetDeletionProtectionOk() (*bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionProtection

`func (o *UpdateResourceInstanceMetadataRequest2) SetDeletionProtection(v bool)`

SetDeletionProtection sets DeletionProtection field to given value.

### HasDeletionProtection

`func (o *UpdateResourceInstanceMetadataRequest2) HasDeletionProtection() bool`

HasDeletionProtection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


