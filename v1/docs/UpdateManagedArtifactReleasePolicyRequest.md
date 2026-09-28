# UpdateManagedArtifactReleasePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpgrade** | **bool** | Enable automatic adoption of the newest published release. When true, preferredBundleVersion must be omitted and any stored pin is cleared. | 
**CloudProvider** | **string** | Name of the Infra Provider | 
**EnvironmentType** | **string** | The type of service environment | 
**PreferredBundleVersion** | Pointer to **string** | Published release to prefer while auto-upgrade is disabled. When disabling auto-upgrade, omission preserves an existing preferred release or pins the current effective release if no preference exists. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateManagedArtifactReleasePolicyRequest

`func NewUpdateManagedArtifactReleasePolicyRequest(autoUpgrade bool, cloudProvider string, environmentType string, token string, ) *UpdateManagedArtifactReleasePolicyRequest`

NewUpdateManagedArtifactReleasePolicyRequest instantiates a new UpdateManagedArtifactReleasePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateManagedArtifactReleasePolicyRequestWithDefaults

`func NewUpdateManagedArtifactReleasePolicyRequestWithDefaults() *UpdateManagedArtifactReleasePolicyRequest`

NewUpdateManagedArtifactReleasePolicyRequestWithDefaults instantiates a new UpdateManagedArtifactReleasePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpgrade

`func (o *UpdateManagedArtifactReleasePolicyRequest) GetAutoUpgrade() bool`

GetAutoUpgrade returns the AutoUpgrade field if non-nil, zero value otherwise.

### GetAutoUpgradeOk

`func (o *UpdateManagedArtifactReleasePolicyRequest) GetAutoUpgradeOk() (*bool, bool)`

GetAutoUpgradeOk returns a tuple with the AutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgrade

`func (o *UpdateManagedArtifactReleasePolicyRequest) SetAutoUpgrade(v bool)`

SetAutoUpgrade sets AutoUpgrade field to given value.


### GetCloudProvider

`func (o *UpdateManagedArtifactReleasePolicyRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *UpdateManagedArtifactReleasePolicyRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *UpdateManagedArtifactReleasePolicyRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetEnvironmentType

`func (o *UpdateManagedArtifactReleasePolicyRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *UpdateManagedArtifactReleasePolicyRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *UpdateManagedArtifactReleasePolicyRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetPreferredBundleVersion

`func (o *UpdateManagedArtifactReleasePolicyRequest) GetPreferredBundleVersion() string`

GetPreferredBundleVersion returns the PreferredBundleVersion field if non-nil, zero value otherwise.

### GetPreferredBundleVersionOk

`func (o *UpdateManagedArtifactReleasePolicyRequest) GetPreferredBundleVersionOk() (*string, bool)`

GetPreferredBundleVersionOk returns a tuple with the PreferredBundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredBundleVersion

`func (o *UpdateManagedArtifactReleasePolicyRequest) SetPreferredBundleVersion(v string)`

SetPreferredBundleVersion sets PreferredBundleVersion field to given value.

### HasPreferredBundleVersion

`func (o *UpdateManagedArtifactReleasePolicyRequest) HasPreferredBundleVersion() bool`

HasPreferredBundleVersion returns a boolean if a field has been set.

### GetToken

`func (o *UpdateManagedArtifactReleasePolicyRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateManagedArtifactReleasePolicyRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateManagedArtifactReleasePolicyRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


