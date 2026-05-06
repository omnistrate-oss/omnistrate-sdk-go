# OnboardingTerraformConfigurations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationPerCloudProvider** | [**map[string]OnboardingTerraformConfiguration**](OnboardingTerraformConfiguration.md) | Terraform configuration per cloud provider. | 
**ConfigurationPerOnPremPlatform** | Pointer to [**map[string]OnboardingTerraformConfiguration**](OnboardingTerraformConfiguration.md) | Terraform configuration per on-prem platform. | [optional] 

## Methods

### NewOnboardingTerraformConfigurations

`func NewOnboardingTerraformConfigurations(configurationPerCloudProvider map[string]OnboardingTerraformConfiguration, ) *OnboardingTerraformConfigurations`

NewOnboardingTerraformConfigurations instantiates a new OnboardingTerraformConfigurations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingTerraformConfigurationsWithDefaults

`func NewOnboardingTerraformConfigurationsWithDefaults() *OnboardingTerraformConfigurations`

NewOnboardingTerraformConfigurationsWithDefaults instantiates a new OnboardingTerraformConfigurations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationPerCloudProvider

`func (o *OnboardingTerraformConfigurations) GetConfigurationPerCloudProvider() map[string]OnboardingTerraformConfiguration`

GetConfigurationPerCloudProvider returns the ConfigurationPerCloudProvider field if non-nil, zero value otherwise.

### GetConfigurationPerCloudProviderOk

`func (o *OnboardingTerraformConfigurations) GetConfigurationPerCloudProviderOk() (*map[string]OnboardingTerraformConfiguration, bool)`

GetConfigurationPerCloudProviderOk returns a tuple with the ConfigurationPerCloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPerCloudProvider

`func (o *OnboardingTerraformConfigurations) SetConfigurationPerCloudProvider(v map[string]OnboardingTerraformConfiguration)`

SetConfigurationPerCloudProvider sets ConfigurationPerCloudProvider field to given value.


### GetConfigurationPerOnPremPlatform

`func (o *OnboardingTerraformConfigurations) GetConfigurationPerOnPremPlatform() map[string]OnboardingTerraformConfiguration`

GetConfigurationPerOnPremPlatform returns the ConfigurationPerOnPremPlatform field if non-nil, zero value otherwise.

### GetConfigurationPerOnPremPlatformOk

`func (o *OnboardingTerraformConfigurations) GetConfigurationPerOnPremPlatformOk() (*map[string]OnboardingTerraformConfiguration, bool)`

GetConfigurationPerOnPremPlatformOk returns a tuple with the ConfigurationPerOnPremPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPerOnPremPlatform

`func (o *OnboardingTerraformConfigurations) SetConfigurationPerOnPremPlatform(v map[string]OnboardingTerraformConfiguration)`

SetConfigurationPerOnPremPlatform sets ConfigurationPerOnPremPlatform field to given value.

### HasConfigurationPerOnPremPlatform

`func (o *OnboardingTerraformConfigurations) HasConfigurationPerOnPremPlatform() bool`

HasConfigurationPerOnPremPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


