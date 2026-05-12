# CloudNativeOnboardingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeYAML** | Pointer to **string** | The cloud-native YAML content. | [optional] 
**DeploymentConfig** | Pointer to [**OnboardingDeploymentConfig**](OnboardingDeploymentConfig.md) |  | [optional] 
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | Environment variables for the service. | [optional] 
**ResourceConfig** | Pointer to [**OnboardingResourceConfig**](OnboardingResourceConfig.md) |  | [optional] 
**Secrets** | Pointer to [**[]OnboardingSecret**](OnboardingSecret.md) | Secrets for the service. | [optional] 
**ServiceConfig** | Pointer to [**OnboardingServiceConfig**](OnboardingServiceConfig.md) |  | [optional] 

## Methods

### NewCloudNativeOnboardingMetadata

`func NewCloudNativeOnboardingMetadata() *CloudNativeOnboardingMetadata`

NewCloudNativeOnboardingMetadata instantiates a new CloudNativeOnboardingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNativeOnboardingMetadataWithDefaults

`func NewCloudNativeOnboardingMetadataWithDefaults() *CloudNativeOnboardingMetadata`

NewCloudNativeOnboardingMetadataWithDefaults instantiates a new CloudNativeOnboardingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeYAML

`func (o *CloudNativeOnboardingMetadata) GetCloudNativeYAML() string`

GetCloudNativeYAML returns the CloudNativeYAML field if non-nil, zero value otherwise.

### GetCloudNativeYAMLOk

`func (o *CloudNativeOnboardingMetadata) GetCloudNativeYAMLOk() (*string, bool)`

GetCloudNativeYAMLOk returns a tuple with the CloudNativeYAML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeYAML

`func (o *CloudNativeOnboardingMetadata) SetCloudNativeYAML(v string)`

SetCloudNativeYAML sets CloudNativeYAML field to given value.

### HasCloudNativeYAML

`func (o *CloudNativeOnboardingMetadata) HasCloudNativeYAML() bool`

HasCloudNativeYAML returns a boolean if a field has been set.

### GetDeploymentConfig

`func (o *CloudNativeOnboardingMetadata) GetDeploymentConfig() OnboardingDeploymentConfig`

GetDeploymentConfig returns the DeploymentConfig field if non-nil, zero value otherwise.

### GetDeploymentConfigOk

`func (o *CloudNativeOnboardingMetadata) GetDeploymentConfigOk() (*OnboardingDeploymentConfig, bool)`

GetDeploymentConfigOk returns a tuple with the DeploymentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfig

`func (o *CloudNativeOnboardingMetadata) SetDeploymentConfig(v OnboardingDeploymentConfig)`

SetDeploymentConfig sets DeploymentConfig field to given value.

### HasDeploymentConfig

`func (o *CloudNativeOnboardingMetadata) HasDeploymentConfig() bool`

HasDeploymentConfig returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *CloudNativeOnboardingMetadata) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *CloudNativeOnboardingMetadata) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *CloudNativeOnboardingMetadata) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *CloudNativeOnboardingMetadata) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetResourceConfig

`func (o *CloudNativeOnboardingMetadata) GetResourceConfig() OnboardingResourceConfig`

GetResourceConfig returns the ResourceConfig field if non-nil, zero value otherwise.

### GetResourceConfigOk

`func (o *CloudNativeOnboardingMetadata) GetResourceConfigOk() (*OnboardingResourceConfig, bool)`

GetResourceConfigOk returns a tuple with the ResourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceConfig

`func (o *CloudNativeOnboardingMetadata) SetResourceConfig(v OnboardingResourceConfig)`

SetResourceConfig sets ResourceConfig field to given value.

### HasResourceConfig

`func (o *CloudNativeOnboardingMetadata) HasResourceConfig() bool`

HasResourceConfig returns a boolean if a field has been set.

### GetSecrets

`func (o *CloudNativeOnboardingMetadata) GetSecrets() []OnboardingSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *CloudNativeOnboardingMetadata) GetSecretsOk() (*[]OnboardingSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *CloudNativeOnboardingMetadata) SetSecrets(v []OnboardingSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *CloudNativeOnboardingMetadata) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetServiceConfig

`func (o *CloudNativeOnboardingMetadata) GetServiceConfig() OnboardingServiceConfig`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *CloudNativeOnboardingMetadata) GetServiceConfigOk() (*OnboardingServiceConfig, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *CloudNativeOnboardingMetadata) SetServiceConfig(v OnboardingServiceConfig)`

SetServiceConfig sets ServiceConfig field to given value.

### HasServiceConfig

`func (o *CloudNativeOnboardingMetadata) HasServiceConfig() bool`

HasServiceConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


