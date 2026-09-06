# ValidateServiceSpecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**[]ValidationArtifactInput**](ValidationArtifactInput.md) | Local artifact content for this request. Omit on a discovery request to learn which local content the specification requires; the response then reports the missing content through requiredArtifacts and an INCOMPLETE status. Supplying an unknown, unreferenced, duplicate or internally inconsistent artifact is a bad request. | [optional] 
**Configs** | Pointer to **map[string]string** | Configs for the service. Key is the compose spec name of the config and value is base64 encoded config content. Rejected when nonempty for the service-plan spec type. | [optional] 
**Description** | Pointer to **string** | Proposed description of the service. Used only as candidate metadata; the current service is never updated. | [optional] 
**Environment** | Pointer to **string** | The environment the candidate targets. Resolved with the same default as the corresponding real build caller when omitted; clients that resolve it themselves send it explicitly. | [optional] 
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**FileContent** | **string** | Base64 encoded specification YAML, using the same encoding convention as the build endpoints. These are the exact prepared bytes that are validated. | 
**ForceCreateNewServicePlanVersion** | Pointer to **bool** | Proposed intent to force a new service plan version. Describes the candidate only; it never bypasses validation and no version is created. | [optional] 
**Name** | **string** | Name of the Service. Validated with the same name rules a real build applies. | 
**Release** | Pointer to **bool** | Proposed release intent. Describes the candidate only; nothing is released. Supported release-readiness checks run regardless of this value. | [optional] 
**ReleaseAsPreferred** | Pointer to **bool** | Proposed intent to release as preferred. Describes the candidate only; no version is released or marked preferred. | [optional] 
**ReleaseVersionName** | Pointer to **string** | Proposed release version name. Describes the candidate only; no release name is written. | [optional] 
**Secrets** | Pointer to **map[string]string** | Secrets for the service. Key is the compose spec name of the secret and value is base64 encoded secret content. Rejected when nonempty for the service-plan spec type. Never logged, echoed back or persisted. | [optional] 
**ServiceLogoURL** | Pointer to **string** | Proposed logo for the service. Used only as candidate metadata; the current service is never updated. | [optional] 
**SpecType** | **string** | The format of the submitted specification. The caller always states the format explicitly; the server never infers it by attempting to parse one format and recovering from the error. | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewValidateServiceSpecRequest

`func NewValidateServiceSpecRequest(fileContent string, name string, specType string, token string, ) *ValidateServiceSpecRequest`

NewValidateServiceSpecRequest instantiates a new ValidateServiceSpecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateServiceSpecRequestWithDefaults

`func NewValidateServiceSpecRequestWithDefaults() *ValidateServiceSpecRequest`

NewValidateServiceSpecRequestWithDefaults instantiates a new ValidateServiceSpecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *ValidateServiceSpecRequest) GetArtifacts() []ValidationArtifactInput`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ValidateServiceSpecRequest) GetArtifactsOk() (*[]ValidationArtifactInput, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ValidateServiceSpecRequest) SetArtifacts(v []ValidationArtifactInput)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ValidateServiceSpecRequest) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetConfigs

`func (o *ValidateServiceSpecRequest) GetConfigs() map[string]string`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *ValidateServiceSpecRequest) GetConfigsOk() (*map[string]string, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *ValidateServiceSpecRequest) SetConfigs(v map[string]string)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *ValidateServiceSpecRequest) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetDescription

`func (o *ValidateServiceSpecRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ValidateServiceSpecRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ValidateServiceSpecRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ValidateServiceSpecRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *ValidateServiceSpecRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ValidateServiceSpecRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ValidateServiceSpecRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ValidateServiceSpecRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *ValidateServiceSpecRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *ValidateServiceSpecRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *ValidateServiceSpecRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *ValidateServiceSpecRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetFileContent

`func (o *ValidateServiceSpecRequest) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *ValidateServiceSpecRequest) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *ValidateServiceSpecRequest) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetForceCreateNewServicePlanVersion

`func (o *ValidateServiceSpecRequest) GetForceCreateNewServicePlanVersion() bool`

GetForceCreateNewServicePlanVersion returns the ForceCreateNewServicePlanVersion field if non-nil, zero value otherwise.

### GetForceCreateNewServicePlanVersionOk

`func (o *ValidateServiceSpecRequest) GetForceCreateNewServicePlanVersionOk() (*bool, bool)`

GetForceCreateNewServicePlanVersionOk returns a tuple with the ForceCreateNewServicePlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCreateNewServicePlanVersion

`func (o *ValidateServiceSpecRequest) SetForceCreateNewServicePlanVersion(v bool)`

SetForceCreateNewServicePlanVersion sets ForceCreateNewServicePlanVersion field to given value.

### HasForceCreateNewServicePlanVersion

`func (o *ValidateServiceSpecRequest) HasForceCreateNewServicePlanVersion() bool`

HasForceCreateNewServicePlanVersion returns a boolean if a field has been set.

### GetName

`func (o *ValidateServiceSpecRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidateServiceSpecRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidateServiceSpecRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRelease

`func (o *ValidateServiceSpecRequest) GetRelease() bool`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *ValidateServiceSpecRequest) GetReleaseOk() (*bool, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *ValidateServiceSpecRequest) SetRelease(v bool)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *ValidateServiceSpecRequest) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetReleaseAsPreferred

`func (o *ValidateServiceSpecRequest) GetReleaseAsPreferred() bool`

GetReleaseAsPreferred returns the ReleaseAsPreferred field if non-nil, zero value otherwise.

### GetReleaseAsPreferredOk

`func (o *ValidateServiceSpecRequest) GetReleaseAsPreferredOk() (*bool, bool)`

GetReleaseAsPreferredOk returns a tuple with the ReleaseAsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseAsPreferred

`func (o *ValidateServiceSpecRequest) SetReleaseAsPreferred(v bool)`

SetReleaseAsPreferred sets ReleaseAsPreferred field to given value.

### HasReleaseAsPreferred

`func (o *ValidateServiceSpecRequest) HasReleaseAsPreferred() bool`

HasReleaseAsPreferred returns a boolean if a field has been set.

### GetReleaseVersionName

`func (o *ValidateServiceSpecRequest) GetReleaseVersionName() string`

GetReleaseVersionName returns the ReleaseVersionName field if non-nil, zero value otherwise.

### GetReleaseVersionNameOk

`func (o *ValidateServiceSpecRequest) GetReleaseVersionNameOk() (*string, bool)`

GetReleaseVersionNameOk returns a tuple with the ReleaseVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersionName

`func (o *ValidateServiceSpecRequest) SetReleaseVersionName(v string)`

SetReleaseVersionName sets ReleaseVersionName field to given value.

### HasReleaseVersionName

`func (o *ValidateServiceSpecRequest) HasReleaseVersionName() bool`

HasReleaseVersionName returns a boolean if a field has been set.

### GetSecrets

`func (o *ValidateServiceSpecRequest) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ValidateServiceSpecRequest) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ValidateServiceSpecRequest) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *ValidateServiceSpecRequest) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetServiceLogoURL

`func (o *ValidateServiceSpecRequest) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *ValidateServiceSpecRequest) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *ValidateServiceSpecRequest) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *ValidateServiceSpecRequest) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.

### GetSpecType

`func (o *ValidateServiceSpecRequest) GetSpecType() string`

GetSpecType returns the SpecType field if non-nil, zero value otherwise.

### GetSpecTypeOk

`func (o *ValidateServiceSpecRequest) GetSpecTypeOk() (*string, bool)`

GetSpecTypeOk returns a tuple with the SpecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecType

`func (o *ValidateServiceSpecRequest) SetSpecType(v string)`

SetSpecType sets SpecType field to given value.


### GetToken

`func (o *ValidateServiceSpecRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ValidateServiceSpecRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ValidateServiceSpecRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


