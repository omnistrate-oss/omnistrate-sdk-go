# UpdateManagedArtifactReleasePolicyRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpgrade** | **bool** | Enable automatic adoption of the newest published release. When true, preferredBundleVersion must be omitted and any stored pin is cleared. | 
**PreferredBundleVersion** | Pointer to **string** | Published release to prefer while auto-upgrade is disabled. When disabling auto-upgrade, omission preserves an existing preferred release or pins the current effective release if no preference exists. | [optional] 

## Methods

### NewUpdateManagedArtifactReleasePolicyRequest2

`func NewUpdateManagedArtifactReleasePolicyRequest2(autoUpgrade bool, ) *UpdateManagedArtifactReleasePolicyRequest2`

NewUpdateManagedArtifactReleasePolicyRequest2 instantiates a new UpdateManagedArtifactReleasePolicyRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateManagedArtifactReleasePolicyRequest2WithDefaults

`func NewUpdateManagedArtifactReleasePolicyRequest2WithDefaults() *UpdateManagedArtifactReleasePolicyRequest2`

NewUpdateManagedArtifactReleasePolicyRequest2WithDefaults instantiates a new UpdateManagedArtifactReleasePolicyRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpgrade

`func (o *UpdateManagedArtifactReleasePolicyRequest2) GetAutoUpgrade() bool`

GetAutoUpgrade returns the AutoUpgrade field if non-nil, zero value otherwise.

### GetAutoUpgradeOk

`func (o *UpdateManagedArtifactReleasePolicyRequest2) GetAutoUpgradeOk() (*bool, bool)`

GetAutoUpgradeOk returns a tuple with the AutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgrade

`func (o *UpdateManagedArtifactReleasePolicyRequest2) SetAutoUpgrade(v bool)`

SetAutoUpgrade sets AutoUpgrade field to given value.


### GetPreferredBundleVersion

`func (o *UpdateManagedArtifactReleasePolicyRequest2) GetPreferredBundleVersion() string`

GetPreferredBundleVersion returns the PreferredBundleVersion field if non-nil, zero value otherwise.

### GetPreferredBundleVersionOk

`func (o *UpdateManagedArtifactReleasePolicyRequest2) GetPreferredBundleVersionOk() (*string, bool)`

GetPreferredBundleVersionOk returns a tuple with the PreferredBundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredBundleVersion

`func (o *UpdateManagedArtifactReleasePolicyRequest2) SetPreferredBundleVersion(v string)`

SetPreferredBundleVersion sets PreferredBundleVersion field to given value.

### HasPreferredBundleVersion

`func (o *UpdateManagedArtifactReleasePolicyRequest2) HasPreferredBundleVersion() bool`

HasPreferredBundleVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


