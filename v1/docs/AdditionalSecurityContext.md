# AdditionalSecurityContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddCapabilities** | Pointer to **[]string** | Capabilities to add | [optional] 
**DropCapabilities** | Pointer to **[]string** | Capabilities to drop | [optional] 
**SupplementalGroupIDs** | Pointer to **[]int64** | Supplemental group IDs | [optional] 

## Methods

### NewAdditionalSecurityContext

`func NewAdditionalSecurityContext() *AdditionalSecurityContext`

NewAdditionalSecurityContext instantiates a new AdditionalSecurityContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalSecurityContextWithDefaults

`func NewAdditionalSecurityContextWithDefaults() *AdditionalSecurityContext`

NewAdditionalSecurityContextWithDefaults instantiates a new AdditionalSecurityContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddCapabilities

`func (o *AdditionalSecurityContext) GetAddCapabilities() []string`

GetAddCapabilities returns the AddCapabilities field if non-nil, zero value otherwise.

### GetAddCapabilitiesOk

`func (o *AdditionalSecurityContext) GetAddCapabilitiesOk() (*[]string, bool)`

GetAddCapabilitiesOk returns a tuple with the AddCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCapabilities

`func (o *AdditionalSecurityContext) SetAddCapabilities(v []string)`

SetAddCapabilities sets AddCapabilities field to given value.

### HasAddCapabilities

`func (o *AdditionalSecurityContext) HasAddCapabilities() bool`

HasAddCapabilities returns a boolean if a field has been set.

### GetDropCapabilities

`func (o *AdditionalSecurityContext) GetDropCapabilities() []string`

GetDropCapabilities returns the DropCapabilities field if non-nil, zero value otherwise.

### GetDropCapabilitiesOk

`func (o *AdditionalSecurityContext) GetDropCapabilitiesOk() (*[]string, bool)`

GetDropCapabilitiesOk returns a tuple with the DropCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropCapabilities

`func (o *AdditionalSecurityContext) SetDropCapabilities(v []string)`

SetDropCapabilities sets DropCapabilities field to given value.

### HasDropCapabilities

`func (o *AdditionalSecurityContext) HasDropCapabilities() bool`

HasDropCapabilities returns a boolean if a field has been set.

### GetSupplementalGroupIDs

`func (o *AdditionalSecurityContext) GetSupplementalGroupIDs() []int64`

GetSupplementalGroupIDs returns the SupplementalGroupIDs field if non-nil, zero value otherwise.

### GetSupplementalGroupIDsOk

`func (o *AdditionalSecurityContext) GetSupplementalGroupIDsOk() (*[]int64, bool)`

GetSupplementalGroupIDsOk returns a tuple with the SupplementalGroupIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalGroupIDs

`func (o *AdditionalSecurityContext) SetSupplementalGroupIDs(v []int64)`

SetSupplementalGroupIDs sets SupplementalGroupIDs field to given value.

### HasSupplementalGroupIDs

`func (o *AdditionalSecurityContext) HasSupplementalGroupIDs() bool`

HasSupplementalGroupIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


