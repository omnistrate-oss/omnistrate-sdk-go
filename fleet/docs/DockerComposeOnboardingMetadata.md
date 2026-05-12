# DockerComposeOnboardingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentConfig** | Pointer to [**OnboardingDeploymentConfig**](OnboardingDeploymentConfig.md) |  | [optional] 
**DockerComposeYAML** | Pointer to **string** | The Docker Compose YAML content. | [optional] 
**Secrets** | Pointer to [**[]OnboardingSecret**](OnboardingSecret.md) | Secrets for the service. | [optional] 
**ServiceConfig** | Pointer to [**OnboardingServiceConfig**](OnboardingServiceConfig.md) |  | [optional] 

## Methods

### NewDockerComposeOnboardingMetadata

`func NewDockerComposeOnboardingMetadata() *DockerComposeOnboardingMetadata`

NewDockerComposeOnboardingMetadata instantiates a new DockerComposeOnboardingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerComposeOnboardingMetadataWithDefaults

`func NewDockerComposeOnboardingMetadataWithDefaults() *DockerComposeOnboardingMetadata`

NewDockerComposeOnboardingMetadataWithDefaults instantiates a new DockerComposeOnboardingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentConfig

`func (o *DockerComposeOnboardingMetadata) GetDeploymentConfig() OnboardingDeploymentConfig`

GetDeploymentConfig returns the DeploymentConfig field if non-nil, zero value otherwise.

### GetDeploymentConfigOk

`func (o *DockerComposeOnboardingMetadata) GetDeploymentConfigOk() (*OnboardingDeploymentConfig, bool)`

GetDeploymentConfigOk returns a tuple with the DeploymentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfig

`func (o *DockerComposeOnboardingMetadata) SetDeploymentConfig(v OnboardingDeploymentConfig)`

SetDeploymentConfig sets DeploymentConfig field to given value.

### HasDeploymentConfig

`func (o *DockerComposeOnboardingMetadata) HasDeploymentConfig() bool`

HasDeploymentConfig returns a boolean if a field has been set.

### GetDockerComposeYAML

`func (o *DockerComposeOnboardingMetadata) GetDockerComposeYAML() string`

GetDockerComposeYAML returns the DockerComposeYAML field if non-nil, zero value otherwise.

### GetDockerComposeYAMLOk

`func (o *DockerComposeOnboardingMetadata) GetDockerComposeYAMLOk() (*string, bool)`

GetDockerComposeYAMLOk returns a tuple with the DockerComposeYAML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerComposeYAML

`func (o *DockerComposeOnboardingMetadata) SetDockerComposeYAML(v string)`

SetDockerComposeYAML sets DockerComposeYAML field to given value.

### HasDockerComposeYAML

`func (o *DockerComposeOnboardingMetadata) HasDockerComposeYAML() bool`

HasDockerComposeYAML returns a boolean if a field has been set.

### GetSecrets

`func (o *DockerComposeOnboardingMetadata) GetSecrets() []OnboardingSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *DockerComposeOnboardingMetadata) GetSecretsOk() (*[]OnboardingSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *DockerComposeOnboardingMetadata) SetSecrets(v []OnboardingSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *DockerComposeOnboardingMetadata) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetServiceConfig

`func (o *DockerComposeOnboardingMetadata) GetServiceConfig() OnboardingServiceConfig`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *DockerComposeOnboardingMetadata) GetServiceConfigOk() (*OnboardingServiceConfig, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *DockerComposeOnboardingMetadata) SetServiceConfig(v OnboardingServiceConfig)`

SetServiceConfig sets ServiceConfig field to given value.

### HasServiceConfig

`func (o *DockerComposeOnboardingMetadata) HasServiceConfig() bool`

HasServiceConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


