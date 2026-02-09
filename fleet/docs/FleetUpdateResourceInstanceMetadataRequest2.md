# FleetUpdateResourceInstanceMetadataRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTags** | Pointer to [**[]CustomTag**](CustomTag.md) | The custom tags for the resource instance. | [optional] 
**DeletionProtection** | Pointer to **bool** | Set to true to enable deletion protection or false to disable it | [optional] 
**EnableDebugMode** | Pointer to **bool** | Enable debug mode | [optional] 

## Methods

### NewFleetUpdateResourceInstanceMetadataRequest2

`func NewFleetUpdateResourceInstanceMetadataRequest2() *FleetUpdateResourceInstanceMetadataRequest2`

NewFleetUpdateResourceInstanceMetadataRequest2 instantiates a new FleetUpdateResourceInstanceMetadataRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateResourceInstanceMetadataRequest2WithDefaults

`func NewFleetUpdateResourceInstanceMetadataRequest2WithDefaults() *FleetUpdateResourceInstanceMetadataRequest2`

NewFleetUpdateResourceInstanceMetadataRequest2WithDefaults instantiates a new FleetUpdateResourceInstanceMetadataRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTags

`func (o *FleetUpdateResourceInstanceMetadataRequest2) GetCustomTags() []CustomTag`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *FleetUpdateResourceInstanceMetadataRequest2) GetCustomTagsOk() (*[]CustomTag, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *FleetUpdateResourceInstanceMetadataRequest2) SetCustomTags(v []CustomTag)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *FleetUpdateResourceInstanceMetadataRequest2) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### GetDeletionProtection

`func (o *FleetUpdateResourceInstanceMetadataRequest2) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *FleetUpdateResourceInstanceMetadataRequest2) GetDeletionProtectionOk() (*bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionProtection

`func (o *FleetUpdateResourceInstanceMetadataRequest2) SetDeletionProtection(v bool)`

SetDeletionProtection sets DeletionProtection field to given value.

### HasDeletionProtection

`func (o *FleetUpdateResourceInstanceMetadataRequest2) HasDeletionProtection() bool`

HasDeletionProtection returns a boolean if a field has been set.

### GetEnableDebugMode

`func (o *FleetUpdateResourceInstanceMetadataRequest2) GetEnableDebugMode() bool`

GetEnableDebugMode returns the EnableDebugMode field if non-nil, zero value otherwise.

### GetEnableDebugModeOk

`func (o *FleetUpdateResourceInstanceMetadataRequest2) GetEnableDebugModeOk() (*bool, bool)`

GetEnableDebugModeOk returns a tuple with the EnableDebugMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebugMode

`func (o *FleetUpdateResourceInstanceMetadataRequest2) SetEnableDebugMode(v bool)`

SetEnableDebugMode sets EnableDebugMode field to given value.

### HasEnableDebugMode

`func (o *FleetUpdateResourceInstanceMetadataRequest2) HasEnableDebugMode() bool`

HasEnableDebugMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


