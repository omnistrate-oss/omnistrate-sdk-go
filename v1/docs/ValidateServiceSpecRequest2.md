# ValidateServiceSpecRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**[]ValidationArtifactInput**](ValidationArtifactInput.md) | Local artifact content for this request. Omit on a discovery request to learn which local content the specification requires; the response then reports the missing content through requiredArtifacts and an INCOMPLETE status. Supplying an unknown, unreferenced, duplicate or internally inconsistent artifact is a bad request. | [optional] 
**Configs** | Pointer to **map[string]string** | Configs for the service. Key is the compose spec name of the config and value is base64 encoded config content. Rejected when nonempty for the service-plan spec type. | [optional] 
**Description** | Pointer to **string** | Proposed description of the service. Used only as candidate metadata; the current service is never updated. | [optional] 
**Environment** | Pointer to **string** | The environment the candidate targets. Resolved with the same default as the corresponding real build caller when omitted; clients that resolve it themselves send it explicitly. | [optional] 
**EnvironmentType** | Pointer to **string** | The type of the environment the candidate targets. | [optional] 
**FileContent** | **string** | Base64 encoded specification YAML, using the same encoding convention as the build endpoints. These are the exact prepared bytes that are validated. | 
**ForceCreateNewServicePlanVersion** | Pointer to **bool** | Proposed intent to force a new service plan version. Describes the candidate only; it never bypasses validation and no version is created. | [optional] 
**Name** | **string** | Name of the Service. Validated with the same name rules a real build applies. | 
**Release** | Pointer to **bool** | Proposed release intent. Describes the candidate only; nothing is released. Supported release-readiness checks run regardless of this value. | [optional] 
**ReleaseAsPreferred** | Pointer to **bool** | Proposed intent to release as preferred. Describes the candidate only; no version is released or marked preferred. | [optional] 
**ReleaseVersionName** | Pointer to **string** | Proposed release version name. Describes the candidate only; no release name is written. | [optional] 
**Secrets** | Pointer to **map[string]string** | Secrets for the service. Key is the compose spec name of the secret and value is base64 encoded secret content. Rejected when nonempty for the service-plan spec type. Never logged, echoed back or persisted. | [optional] 
**ServiceLogoURL** | Pointer to **string** | Proposed logo for the service. Used only as candidate metadata; the current service is never updated. | [optional] 
**SpecType** | **string** | The format of fileContent. | 

## Methods

### NewValidateServiceSpecRequest2

`func NewValidateServiceSpecRequest2(fileContent string, name string, specType string, ) *ValidateServiceSpecRequest2`

NewValidateServiceSpecRequest2 instantiates a new ValidateServiceSpecRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateServiceSpecRequest2WithDefaults

`func NewValidateServiceSpecRequest2WithDefaults() *ValidateServiceSpecRequest2`

NewValidateServiceSpecRequest2WithDefaults instantiates a new ValidateServiceSpecRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *ValidateServiceSpecRequest2) GetArtifacts() []ValidationArtifactInput`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ValidateServiceSpecRequest2) GetArtifactsOk() (*[]ValidationArtifactInput, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ValidateServiceSpecRequest2) SetArtifacts(v []ValidationArtifactInput)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ValidateServiceSpecRequest2) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetConfigs

`func (o *ValidateServiceSpecRequest2) GetConfigs() map[string]string`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *ValidateServiceSpecRequest2) GetConfigsOk() (*map[string]string, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *ValidateServiceSpecRequest2) SetConfigs(v map[string]string)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *ValidateServiceSpecRequest2) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetDescription

`func (o *ValidateServiceSpecRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ValidateServiceSpecRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ValidateServiceSpecRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ValidateServiceSpecRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *ValidateServiceSpecRequest2) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ValidateServiceSpecRequest2) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ValidateServiceSpecRequest2) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ValidateServiceSpecRequest2) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *ValidateServiceSpecRequest2) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *ValidateServiceSpecRequest2) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *ValidateServiceSpecRequest2) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *ValidateServiceSpecRequest2) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetFileContent

`func (o *ValidateServiceSpecRequest2) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *ValidateServiceSpecRequest2) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *ValidateServiceSpecRequest2) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetForceCreateNewServicePlanVersion

`func (o *ValidateServiceSpecRequest2) GetForceCreateNewServicePlanVersion() bool`

GetForceCreateNewServicePlanVersion returns the ForceCreateNewServicePlanVersion field if non-nil, zero value otherwise.

### GetForceCreateNewServicePlanVersionOk

`func (o *ValidateServiceSpecRequest2) GetForceCreateNewServicePlanVersionOk() (*bool, bool)`

GetForceCreateNewServicePlanVersionOk returns a tuple with the ForceCreateNewServicePlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCreateNewServicePlanVersion

`func (o *ValidateServiceSpecRequest2) SetForceCreateNewServicePlanVersion(v bool)`

SetForceCreateNewServicePlanVersion sets ForceCreateNewServicePlanVersion field to given value.

### HasForceCreateNewServicePlanVersion

`func (o *ValidateServiceSpecRequest2) HasForceCreateNewServicePlanVersion() bool`

HasForceCreateNewServicePlanVersion returns a boolean if a field has been set.

### GetName

`func (o *ValidateServiceSpecRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidateServiceSpecRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidateServiceSpecRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetRelease

`func (o *ValidateServiceSpecRequest2) GetRelease() bool`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *ValidateServiceSpecRequest2) GetReleaseOk() (*bool, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *ValidateServiceSpecRequest2) SetRelease(v bool)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *ValidateServiceSpecRequest2) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetReleaseAsPreferred

`func (o *ValidateServiceSpecRequest2) GetReleaseAsPreferred() bool`

GetReleaseAsPreferred returns the ReleaseAsPreferred field if non-nil, zero value otherwise.

### GetReleaseAsPreferredOk

`func (o *ValidateServiceSpecRequest2) GetReleaseAsPreferredOk() (*bool, bool)`

GetReleaseAsPreferredOk returns a tuple with the ReleaseAsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseAsPreferred

`func (o *ValidateServiceSpecRequest2) SetReleaseAsPreferred(v bool)`

SetReleaseAsPreferred sets ReleaseAsPreferred field to given value.

### HasReleaseAsPreferred

`func (o *ValidateServiceSpecRequest2) HasReleaseAsPreferred() bool`

HasReleaseAsPreferred returns a boolean if a field has been set.

### GetReleaseVersionName

`func (o *ValidateServiceSpecRequest2) GetReleaseVersionName() string`

GetReleaseVersionName returns the ReleaseVersionName field if non-nil, zero value otherwise.

### GetReleaseVersionNameOk

`func (o *ValidateServiceSpecRequest2) GetReleaseVersionNameOk() (*string, bool)`

GetReleaseVersionNameOk returns a tuple with the ReleaseVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersionName

`func (o *ValidateServiceSpecRequest2) SetReleaseVersionName(v string)`

SetReleaseVersionName sets ReleaseVersionName field to given value.

### HasReleaseVersionName

`func (o *ValidateServiceSpecRequest2) HasReleaseVersionName() bool`

HasReleaseVersionName returns a boolean if a field has been set.

### GetSecrets

`func (o *ValidateServiceSpecRequest2) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ValidateServiceSpecRequest2) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ValidateServiceSpecRequest2) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *ValidateServiceSpecRequest2) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetServiceLogoURL

`func (o *ValidateServiceSpecRequest2) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *ValidateServiceSpecRequest2) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *ValidateServiceSpecRequest2) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *ValidateServiceSpecRequest2) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.

### GetSpecType

`func (o *ValidateServiceSpecRequest2) GetSpecType() string`

GetSpecType returns the SpecType field if non-nil, zero value otherwise.

### GetSpecTypeOk

`func (o *ValidateServiceSpecRequest2) GetSpecTypeOk() (*string, bool)`

GetSpecTypeOk returns a tuple with the SpecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecType

`func (o *ValidateServiceSpecRequest2) SetSpecType(v string)`

SetSpecType sets SpecType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


