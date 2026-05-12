# OnboardingDeploymentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigIds** | Pointer to **[]string** | The account config IDs to use for deployment. | [optional] 
**ByoaCloudProviders** | Pointer to **[]string** | Cloud providers enabled for Bring-Your-Own-Account deployment. | [optional] 
**DedicatedKubernetesEnabled** | Pointer to **bool** | Whether dedicated Kubernetes is enabled. | [optional] 
**DedicatedNetworkEnabled** | Pointer to **bool** | Whether dedicated network is enabled. | [optional] 
**IntermediaryAccountConfigId** | Pointer to **string** | The account config ID for an intermediary account. | [optional] 
**ModelType** | Pointer to **string** | The model type encapsulating this service | [optional] 
**TierType** | Pointer to **string** | ProductTierType is the type of tier for a product | [optional] 

## Methods

### NewOnboardingDeploymentConfig

`func NewOnboardingDeploymentConfig() *OnboardingDeploymentConfig`

NewOnboardingDeploymentConfig instantiates a new OnboardingDeploymentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingDeploymentConfigWithDefaults

`func NewOnboardingDeploymentConfigWithDefaults() *OnboardingDeploymentConfig`

NewOnboardingDeploymentConfigWithDefaults instantiates a new OnboardingDeploymentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigIds

`func (o *OnboardingDeploymentConfig) GetAccountConfigIds() []string`

GetAccountConfigIds returns the AccountConfigIds field if non-nil, zero value otherwise.

### GetAccountConfigIdsOk

`func (o *OnboardingDeploymentConfig) GetAccountConfigIdsOk() (*[]string, bool)`

GetAccountConfigIdsOk returns a tuple with the AccountConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigIds

`func (o *OnboardingDeploymentConfig) SetAccountConfigIds(v []string)`

SetAccountConfigIds sets AccountConfigIds field to given value.

### HasAccountConfigIds

`func (o *OnboardingDeploymentConfig) HasAccountConfigIds() bool`

HasAccountConfigIds returns a boolean if a field has been set.

### GetByoaCloudProviders

`func (o *OnboardingDeploymentConfig) GetByoaCloudProviders() []string`

GetByoaCloudProviders returns the ByoaCloudProviders field if non-nil, zero value otherwise.

### GetByoaCloudProvidersOk

`func (o *OnboardingDeploymentConfig) GetByoaCloudProvidersOk() (*[]string, bool)`

GetByoaCloudProvidersOk returns a tuple with the ByoaCloudProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaCloudProviders

`func (o *OnboardingDeploymentConfig) SetByoaCloudProviders(v []string)`

SetByoaCloudProviders sets ByoaCloudProviders field to given value.

### HasByoaCloudProviders

`func (o *OnboardingDeploymentConfig) HasByoaCloudProviders() bool`

HasByoaCloudProviders returns a boolean if a field has been set.

### GetDedicatedKubernetesEnabled

`func (o *OnboardingDeploymentConfig) GetDedicatedKubernetesEnabled() bool`

GetDedicatedKubernetesEnabled returns the DedicatedKubernetesEnabled field if non-nil, zero value otherwise.

### GetDedicatedKubernetesEnabledOk

`func (o *OnboardingDeploymentConfig) GetDedicatedKubernetesEnabledOk() (*bool, bool)`

GetDedicatedKubernetesEnabledOk returns a tuple with the DedicatedKubernetesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedKubernetesEnabled

`func (o *OnboardingDeploymentConfig) SetDedicatedKubernetesEnabled(v bool)`

SetDedicatedKubernetesEnabled sets DedicatedKubernetesEnabled field to given value.

### HasDedicatedKubernetesEnabled

`func (o *OnboardingDeploymentConfig) HasDedicatedKubernetesEnabled() bool`

HasDedicatedKubernetesEnabled returns a boolean if a field has been set.

### GetDedicatedNetworkEnabled

`func (o *OnboardingDeploymentConfig) GetDedicatedNetworkEnabled() bool`

GetDedicatedNetworkEnabled returns the DedicatedNetworkEnabled field if non-nil, zero value otherwise.

### GetDedicatedNetworkEnabledOk

`func (o *OnboardingDeploymentConfig) GetDedicatedNetworkEnabledOk() (*bool, bool)`

GetDedicatedNetworkEnabledOk returns a tuple with the DedicatedNetworkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedNetworkEnabled

`func (o *OnboardingDeploymentConfig) SetDedicatedNetworkEnabled(v bool)`

SetDedicatedNetworkEnabled sets DedicatedNetworkEnabled field to given value.

### HasDedicatedNetworkEnabled

`func (o *OnboardingDeploymentConfig) HasDedicatedNetworkEnabled() bool`

HasDedicatedNetworkEnabled returns a boolean if a field has been set.

### GetIntermediaryAccountConfigId

`func (o *OnboardingDeploymentConfig) GetIntermediaryAccountConfigId() string`

GetIntermediaryAccountConfigId returns the IntermediaryAccountConfigId field if non-nil, zero value otherwise.

### GetIntermediaryAccountConfigIdOk

`func (o *OnboardingDeploymentConfig) GetIntermediaryAccountConfigIdOk() (*string, bool)`

GetIntermediaryAccountConfigIdOk returns a tuple with the IntermediaryAccountConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryAccountConfigId

`func (o *OnboardingDeploymentConfig) SetIntermediaryAccountConfigId(v string)`

SetIntermediaryAccountConfigId sets IntermediaryAccountConfigId field to given value.

### HasIntermediaryAccountConfigId

`func (o *OnboardingDeploymentConfig) HasIntermediaryAccountConfigId() bool`

HasIntermediaryAccountConfigId returns a boolean if a field has been set.

### GetModelType

`func (o *OnboardingDeploymentConfig) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *OnboardingDeploymentConfig) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *OnboardingDeploymentConfig) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *OnboardingDeploymentConfig) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetTierType

`func (o *OnboardingDeploymentConfig) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *OnboardingDeploymentConfig) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *OnboardingDeploymentConfig) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *OnboardingDeploymentConfig) HasTierType() bool`

HasTierType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


