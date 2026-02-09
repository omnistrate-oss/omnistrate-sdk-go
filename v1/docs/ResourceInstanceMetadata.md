# ResourceInstanceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTags** | Pointer to [**[]CustomTag**](CustomTag.md) | The custom tags for the resource instance. | [optional] 
**DeletionProtection** | Pointer to **bool** | Indicates if deletion protection is enabled for the instance | [optional] 
**EnableDebugMode** | Pointer to **bool** | Enable debug mode | [optional] 

## Methods

### NewResourceInstanceMetadata

`func NewResourceInstanceMetadata() *ResourceInstanceMetadata`

NewResourceInstanceMetadata instantiates a new ResourceInstanceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInstanceMetadataWithDefaults

`func NewResourceInstanceMetadataWithDefaults() *ResourceInstanceMetadata`

NewResourceInstanceMetadataWithDefaults instantiates a new ResourceInstanceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTags

`func (o *ResourceInstanceMetadata) GetCustomTags() []CustomTag`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *ResourceInstanceMetadata) GetCustomTagsOk() (*[]CustomTag, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *ResourceInstanceMetadata) SetCustomTags(v []CustomTag)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *ResourceInstanceMetadata) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### GetDeletionProtection

`func (o *ResourceInstanceMetadata) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *ResourceInstanceMetadata) GetDeletionProtectionOk() (*bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionProtection

`func (o *ResourceInstanceMetadata) SetDeletionProtection(v bool)`

SetDeletionProtection sets DeletionProtection field to given value.

### HasDeletionProtection

`func (o *ResourceInstanceMetadata) HasDeletionProtection() bool`

HasDeletionProtection returns a boolean if a field has been set.

### GetEnableDebugMode

`func (o *ResourceInstanceMetadata) GetEnableDebugMode() bool`

GetEnableDebugMode returns the EnableDebugMode field if non-nil, zero value otherwise.

### GetEnableDebugModeOk

`func (o *ResourceInstanceMetadata) GetEnableDebugModeOk() (*bool, bool)`

GetEnableDebugModeOk returns a tuple with the EnableDebugMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebugMode

`func (o *ResourceInstanceMetadata) SetEnableDebugMode(v bool)`

SetEnableDebugMode sets EnableDebugMode field to given value.

### HasEnableDebugMode

`func (o *ResourceInstanceMetadata) HasEnableDebugMode() bool`

HasEnableDebugMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


