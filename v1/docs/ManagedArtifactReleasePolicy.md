# ManagedArtifactReleasePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpgrade** | **bool** | Whether the template automatically follows the newest published managed artifact release. The effective default is true when no policy has been stored. | 
**CloudProvider** | **string** | Name of the Infra Provider | 
**EffectiveBundleVersion** | Pointer to **string** | Release currently selected by the template. It is omitted only when no published release is available. | [optional] 
**EnvironmentType** | **string** | The type of service environment | 
**PreferredBundleVersion** | Pointer to **string** | Pinned release selected while auto-upgrade is disabled. Omitted while auto-upgrade is enabled. | [optional] 

## Methods

### NewManagedArtifactReleasePolicy

`func NewManagedArtifactReleasePolicy(autoUpgrade bool, cloudProvider string, environmentType string, ) *ManagedArtifactReleasePolicy`

NewManagedArtifactReleasePolicy instantiates a new ManagedArtifactReleasePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedArtifactReleasePolicyWithDefaults

`func NewManagedArtifactReleasePolicyWithDefaults() *ManagedArtifactReleasePolicy`

NewManagedArtifactReleasePolicyWithDefaults instantiates a new ManagedArtifactReleasePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpgrade

`func (o *ManagedArtifactReleasePolicy) GetAutoUpgrade() bool`

GetAutoUpgrade returns the AutoUpgrade field if non-nil, zero value otherwise.

### GetAutoUpgradeOk

`func (o *ManagedArtifactReleasePolicy) GetAutoUpgradeOk() (*bool, bool)`

GetAutoUpgradeOk returns a tuple with the AutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgrade

`func (o *ManagedArtifactReleasePolicy) SetAutoUpgrade(v bool)`

SetAutoUpgrade sets AutoUpgrade field to given value.


### GetCloudProvider

`func (o *ManagedArtifactReleasePolicy) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ManagedArtifactReleasePolicy) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ManagedArtifactReleasePolicy) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetEffectiveBundleVersion

`func (o *ManagedArtifactReleasePolicy) GetEffectiveBundleVersion() string`

GetEffectiveBundleVersion returns the EffectiveBundleVersion field if non-nil, zero value otherwise.

### GetEffectiveBundleVersionOk

`func (o *ManagedArtifactReleasePolicy) GetEffectiveBundleVersionOk() (*string, bool)`

GetEffectiveBundleVersionOk returns a tuple with the EffectiveBundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveBundleVersion

`func (o *ManagedArtifactReleasePolicy) SetEffectiveBundleVersion(v string)`

SetEffectiveBundleVersion sets EffectiveBundleVersion field to given value.

### HasEffectiveBundleVersion

`func (o *ManagedArtifactReleasePolicy) HasEffectiveBundleVersion() bool`

HasEffectiveBundleVersion returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *ManagedArtifactReleasePolicy) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *ManagedArtifactReleasePolicy) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *ManagedArtifactReleasePolicy) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetPreferredBundleVersion

`func (o *ManagedArtifactReleasePolicy) GetPreferredBundleVersion() string`

GetPreferredBundleVersion returns the PreferredBundleVersion field if non-nil, zero value otherwise.

### GetPreferredBundleVersionOk

`func (o *ManagedArtifactReleasePolicy) GetPreferredBundleVersionOk() (*string, bool)`

GetPreferredBundleVersionOk returns a tuple with the PreferredBundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredBundleVersion

`func (o *ManagedArtifactReleasePolicy) SetPreferredBundleVersion(v string)`

SetPreferredBundleVersion sets PreferredBundleVersion field to given value.

### HasPreferredBundleVersion

`func (o *ManagedArtifactReleasePolicy) HasPreferredBundleVersion() bool`

HasPreferredBundleVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


