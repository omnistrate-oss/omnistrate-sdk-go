# BuildServiceFromComposeSpecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configs** | Pointer to **map[string]string** | Configs for the service. Key is the compose spec name of the config and value is base64 encoded config content | [optional] 
**Description** | Pointer to **string** | A brief description of the service | [optional] 
**Environment** | Pointer to **string** | The environment to build the service in | [optional] 
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**FileContent** | **string** | Base64 encoded Compose Spec YAML in docker compose format | 
**Name** | **string** | Name of the Service | 
**Release** | Pointer to **bool** | Release the service after building | [optional] 
**ReleaseAsPreferred** | Pointer to **bool** | Release the service as preferred | [optional] 
**ReleaseVersionName** | Pointer to **string** | Release version name | [optional] 
**Secrets** | Pointer to **map[string]string** | Secrets for the service. Key is the compose spec name of the secret and value is base64 encoded secret content | [optional] 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewBuildServiceFromComposeSpecRequest

`func NewBuildServiceFromComposeSpecRequest(fileContent string, name string, token string, ) *BuildServiceFromComposeSpecRequest`

NewBuildServiceFromComposeSpecRequest instantiates a new BuildServiceFromComposeSpecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildServiceFromComposeSpecRequestWithDefaults

`func NewBuildServiceFromComposeSpecRequestWithDefaults() *BuildServiceFromComposeSpecRequest`

NewBuildServiceFromComposeSpecRequestWithDefaults instantiates a new BuildServiceFromComposeSpecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigs

`func (o *BuildServiceFromComposeSpecRequest) GetConfigs() map[string]string`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *BuildServiceFromComposeSpecRequest) GetConfigsOk() (*map[string]string, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *BuildServiceFromComposeSpecRequest) SetConfigs(v map[string]string)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *BuildServiceFromComposeSpecRequest) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetDescription

`func (o *BuildServiceFromComposeSpecRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuildServiceFromComposeSpecRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuildServiceFromComposeSpecRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BuildServiceFromComposeSpecRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *BuildServiceFromComposeSpecRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BuildServiceFromComposeSpecRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BuildServiceFromComposeSpecRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BuildServiceFromComposeSpecRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *BuildServiceFromComposeSpecRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *BuildServiceFromComposeSpecRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *BuildServiceFromComposeSpecRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *BuildServiceFromComposeSpecRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetFileContent

`func (o *BuildServiceFromComposeSpecRequest) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *BuildServiceFromComposeSpecRequest) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *BuildServiceFromComposeSpecRequest) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetName

`func (o *BuildServiceFromComposeSpecRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuildServiceFromComposeSpecRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuildServiceFromComposeSpecRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRelease

`func (o *BuildServiceFromComposeSpecRequest) GetRelease() bool`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *BuildServiceFromComposeSpecRequest) GetReleaseOk() (*bool, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *BuildServiceFromComposeSpecRequest) SetRelease(v bool)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *BuildServiceFromComposeSpecRequest) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetReleaseAsPreferred

`func (o *BuildServiceFromComposeSpecRequest) GetReleaseAsPreferred() bool`

GetReleaseAsPreferred returns the ReleaseAsPreferred field if non-nil, zero value otherwise.

### GetReleaseAsPreferredOk

`func (o *BuildServiceFromComposeSpecRequest) GetReleaseAsPreferredOk() (*bool, bool)`

GetReleaseAsPreferredOk returns a tuple with the ReleaseAsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseAsPreferred

`func (o *BuildServiceFromComposeSpecRequest) SetReleaseAsPreferred(v bool)`

SetReleaseAsPreferred sets ReleaseAsPreferred field to given value.

### HasReleaseAsPreferred

`func (o *BuildServiceFromComposeSpecRequest) HasReleaseAsPreferred() bool`

HasReleaseAsPreferred returns a boolean if a field has been set.

### GetReleaseVersionName

`func (o *BuildServiceFromComposeSpecRequest) GetReleaseVersionName() string`

GetReleaseVersionName returns the ReleaseVersionName field if non-nil, zero value otherwise.

### GetReleaseVersionNameOk

`func (o *BuildServiceFromComposeSpecRequest) GetReleaseVersionNameOk() (*string, bool)`

GetReleaseVersionNameOk returns a tuple with the ReleaseVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersionName

`func (o *BuildServiceFromComposeSpecRequest) SetReleaseVersionName(v string)`

SetReleaseVersionName sets ReleaseVersionName field to given value.

### HasReleaseVersionName

`func (o *BuildServiceFromComposeSpecRequest) HasReleaseVersionName() bool`

HasReleaseVersionName returns a boolean if a field has been set.

### GetSecrets

`func (o *BuildServiceFromComposeSpecRequest) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *BuildServiceFromComposeSpecRequest) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *BuildServiceFromComposeSpecRequest) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *BuildServiceFromComposeSpecRequest) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetServiceLogoURL

`func (o *BuildServiceFromComposeSpecRequest) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *BuildServiceFromComposeSpecRequest) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *BuildServiceFromComposeSpecRequest) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *BuildServiceFromComposeSpecRequest) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.

### GetToken

`func (o *BuildServiceFromComposeSpecRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BuildServiceFromComposeSpecRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BuildServiceFromComposeSpecRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


