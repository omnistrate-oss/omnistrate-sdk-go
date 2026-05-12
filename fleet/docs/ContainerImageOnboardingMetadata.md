# ContainerImageOnboardingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerImageURL** | Pointer to **string** | The container image URL. | [optional] 
**DeploymentConfig** | Pointer to [**OnboardingDeploymentConfig**](OnboardingDeploymentConfig.md) |  | [optional] 
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | Environment variables for the container. | [optional] 
**Secrets** | Pointer to [**[]OnboardingSecret**](OnboardingSecret.md) | Secrets for the service. | [optional] 
**ServiceConfig** | Pointer to [**OnboardingServiceConfig**](OnboardingServiceConfig.md) |  | [optional] 

## Methods

### NewContainerImageOnboardingMetadata

`func NewContainerImageOnboardingMetadata() *ContainerImageOnboardingMetadata`

NewContainerImageOnboardingMetadata instantiates a new ContainerImageOnboardingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageOnboardingMetadataWithDefaults

`func NewContainerImageOnboardingMetadataWithDefaults() *ContainerImageOnboardingMetadata`

NewContainerImageOnboardingMetadataWithDefaults instantiates a new ContainerImageOnboardingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerImageURL

`func (o *ContainerImageOnboardingMetadata) GetContainerImageURL() string`

GetContainerImageURL returns the ContainerImageURL field if non-nil, zero value otherwise.

### GetContainerImageURLOk

`func (o *ContainerImageOnboardingMetadata) GetContainerImageURLOk() (*string, bool)`

GetContainerImageURLOk returns a tuple with the ContainerImageURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImageURL

`func (o *ContainerImageOnboardingMetadata) SetContainerImageURL(v string)`

SetContainerImageURL sets ContainerImageURL field to given value.

### HasContainerImageURL

`func (o *ContainerImageOnboardingMetadata) HasContainerImageURL() bool`

HasContainerImageURL returns a boolean if a field has been set.

### GetDeploymentConfig

`func (o *ContainerImageOnboardingMetadata) GetDeploymentConfig() OnboardingDeploymentConfig`

GetDeploymentConfig returns the DeploymentConfig field if non-nil, zero value otherwise.

### GetDeploymentConfigOk

`func (o *ContainerImageOnboardingMetadata) GetDeploymentConfigOk() (*OnboardingDeploymentConfig, bool)`

GetDeploymentConfigOk returns a tuple with the DeploymentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfig

`func (o *ContainerImageOnboardingMetadata) SetDeploymentConfig(v OnboardingDeploymentConfig)`

SetDeploymentConfig sets DeploymentConfig field to given value.

### HasDeploymentConfig

`func (o *ContainerImageOnboardingMetadata) HasDeploymentConfig() bool`

HasDeploymentConfig returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *ContainerImageOnboardingMetadata) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ContainerImageOnboardingMetadata) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ContainerImageOnboardingMetadata) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ContainerImageOnboardingMetadata) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetSecrets

`func (o *ContainerImageOnboardingMetadata) GetSecrets() []OnboardingSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ContainerImageOnboardingMetadata) GetSecretsOk() (*[]OnboardingSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ContainerImageOnboardingMetadata) SetSecrets(v []OnboardingSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *ContainerImageOnboardingMetadata) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetServiceConfig

`func (o *ContainerImageOnboardingMetadata) GetServiceConfig() OnboardingServiceConfig`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *ContainerImageOnboardingMetadata) GetServiceConfigOk() (*OnboardingServiceConfig, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *ContainerImageOnboardingMetadata) SetServiceConfig(v OnboardingServiceConfig)`

SetServiceConfig sets ServiceConfig field to given value.

### HasServiceConfig

`func (o *ContainerImageOnboardingMetadata) HasServiceConfig() bool`

HasServiceConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


