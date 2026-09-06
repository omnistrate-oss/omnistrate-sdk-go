# BuildServiceFromComposeSpecRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configs** | Pointer to **map[string]string** | Configs for the service. Key is the compose spec name of the config and value is base64 encoded config content | [optional] 
**Description** | Pointer to **string** | A brief description of the service | [optional] 
**Dryrun** | Pointer to **bool** | The dryrun flag is deprecated for validation use; send ValidateServiceSpec to /service/spec/validate instead. When dryrun is true, the build request is answered by a read-only compatibility adapter: nothing is created, updated, released, deprecated, promoted, published or discarded, no import workflow is started, and existing pending changes are left untouched. The adapter can only answer in this legacy result shape, so its reporting is narrower than the validation endpoint&#39;s. A candidate that validates cleanly and already exists returns this result populated with the real service, service environment and product tier IDs, the read-only list of resources present in the plan but absent from the spec, and isNewServicePlanVersionCreated&#x3D;false. A specification with validation errors returns 400 with bounded, redacted diagnostics. A specification that could not be fully validated returns 400 with a message beginning DRY_RUN_VALIDATION_INCOMPLETE and directs the caller to the validation endpoint. A candidate whose service, environment or product tier does not exist yet cannot be expressed in this result shape and returns 400 with a message beginning DRY_RUN_USE_VALIDATION_ENDPOINT; it is never created in order to produce identifiers, and no empty or fabricated identifier is returned. This request envelope carries no local artifact bytes, so a specification that needs local Terraform, Helm, Kustomize or operator content is reported as incomplete rather than validated against previously uploaded artifacts; supply that content through the artifacts field of ValidateServiceSpec. Explicit repository-backed sources are still read through audited read-only readers. | [optional] 
**Environment** | Pointer to **string** | The environment to build the service in | [optional] 
**EnvironmentType** | Pointer to **string** | The type of the environment | [optional] 
**FileContent** | **string** | Base64 encoded Compose Spec YAML in docker compose format | 
**ForceCreateNewServicePlanVersion** | Pointer to **bool** | Force create a new service plan version when the service is released | [optional] 
**Name** | **string** | Name of the Service | 
**Release** | Pointer to **bool** | Release the service after building | [optional] 
**ReleaseAsPreferred** | Pointer to **bool** | Release the service as preferred | [optional] 
**ReleaseVersionName** | Pointer to **string** | Release version name | [optional] 
**Secrets** | Pointer to **map[string]string** | Secrets for the service. Key is the compose spec name of the secret and value is base64 encoded secret content | [optional] 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 

## Methods

### NewBuildServiceFromComposeSpecRequest2

`func NewBuildServiceFromComposeSpecRequest2(fileContent string, name string, ) *BuildServiceFromComposeSpecRequest2`

NewBuildServiceFromComposeSpecRequest2 instantiates a new BuildServiceFromComposeSpecRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildServiceFromComposeSpecRequest2WithDefaults

`func NewBuildServiceFromComposeSpecRequest2WithDefaults() *BuildServiceFromComposeSpecRequest2`

NewBuildServiceFromComposeSpecRequest2WithDefaults instantiates a new BuildServiceFromComposeSpecRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigs

`func (o *BuildServiceFromComposeSpecRequest2) GetConfigs() map[string]string`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *BuildServiceFromComposeSpecRequest2) GetConfigsOk() (*map[string]string, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *BuildServiceFromComposeSpecRequest2) SetConfigs(v map[string]string)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *BuildServiceFromComposeSpecRequest2) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetDescription

`func (o *BuildServiceFromComposeSpecRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuildServiceFromComposeSpecRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuildServiceFromComposeSpecRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BuildServiceFromComposeSpecRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryrun

`func (o *BuildServiceFromComposeSpecRequest2) GetDryrun() bool`

GetDryrun returns the Dryrun field if non-nil, zero value otherwise.

### GetDryrunOk

`func (o *BuildServiceFromComposeSpecRequest2) GetDryrunOk() (*bool, bool)`

GetDryrunOk returns a tuple with the Dryrun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryrun

`func (o *BuildServiceFromComposeSpecRequest2) SetDryrun(v bool)`

SetDryrun sets Dryrun field to given value.

### HasDryrun

`func (o *BuildServiceFromComposeSpecRequest2) HasDryrun() bool`

HasDryrun returns a boolean if a field has been set.

### GetEnvironment

`func (o *BuildServiceFromComposeSpecRequest2) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BuildServiceFromComposeSpecRequest2) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BuildServiceFromComposeSpecRequest2) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BuildServiceFromComposeSpecRequest2) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *BuildServiceFromComposeSpecRequest2) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *BuildServiceFromComposeSpecRequest2) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *BuildServiceFromComposeSpecRequest2) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *BuildServiceFromComposeSpecRequest2) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetFileContent

`func (o *BuildServiceFromComposeSpecRequest2) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *BuildServiceFromComposeSpecRequest2) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *BuildServiceFromComposeSpecRequest2) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetForceCreateNewServicePlanVersion

`func (o *BuildServiceFromComposeSpecRequest2) GetForceCreateNewServicePlanVersion() bool`

GetForceCreateNewServicePlanVersion returns the ForceCreateNewServicePlanVersion field if non-nil, zero value otherwise.

### GetForceCreateNewServicePlanVersionOk

`func (o *BuildServiceFromComposeSpecRequest2) GetForceCreateNewServicePlanVersionOk() (*bool, bool)`

GetForceCreateNewServicePlanVersionOk returns a tuple with the ForceCreateNewServicePlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCreateNewServicePlanVersion

`func (o *BuildServiceFromComposeSpecRequest2) SetForceCreateNewServicePlanVersion(v bool)`

SetForceCreateNewServicePlanVersion sets ForceCreateNewServicePlanVersion field to given value.

### HasForceCreateNewServicePlanVersion

`func (o *BuildServiceFromComposeSpecRequest2) HasForceCreateNewServicePlanVersion() bool`

HasForceCreateNewServicePlanVersion returns a boolean if a field has been set.

### GetName

`func (o *BuildServiceFromComposeSpecRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuildServiceFromComposeSpecRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuildServiceFromComposeSpecRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetRelease

`func (o *BuildServiceFromComposeSpecRequest2) GetRelease() bool`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *BuildServiceFromComposeSpecRequest2) GetReleaseOk() (*bool, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *BuildServiceFromComposeSpecRequest2) SetRelease(v bool)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *BuildServiceFromComposeSpecRequest2) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetReleaseAsPreferred

`func (o *BuildServiceFromComposeSpecRequest2) GetReleaseAsPreferred() bool`

GetReleaseAsPreferred returns the ReleaseAsPreferred field if non-nil, zero value otherwise.

### GetReleaseAsPreferredOk

`func (o *BuildServiceFromComposeSpecRequest2) GetReleaseAsPreferredOk() (*bool, bool)`

GetReleaseAsPreferredOk returns a tuple with the ReleaseAsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseAsPreferred

`func (o *BuildServiceFromComposeSpecRequest2) SetReleaseAsPreferred(v bool)`

SetReleaseAsPreferred sets ReleaseAsPreferred field to given value.

### HasReleaseAsPreferred

`func (o *BuildServiceFromComposeSpecRequest2) HasReleaseAsPreferred() bool`

HasReleaseAsPreferred returns a boolean if a field has been set.

### GetReleaseVersionName

`func (o *BuildServiceFromComposeSpecRequest2) GetReleaseVersionName() string`

GetReleaseVersionName returns the ReleaseVersionName field if non-nil, zero value otherwise.

### GetReleaseVersionNameOk

`func (o *BuildServiceFromComposeSpecRequest2) GetReleaseVersionNameOk() (*string, bool)`

GetReleaseVersionNameOk returns a tuple with the ReleaseVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersionName

`func (o *BuildServiceFromComposeSpecRequest2) SetReleaseVersionName(v string)`

SetReleaseVersionName sets ReleaseVersionName field to given value.

### HasReleaseVersionName

`func (o *BuildServiceFromComposeSpecRequest2) HasReleaseVersionName() bool`

HasReleaseVersionName returns a boolean if a field has been set.

### GetSecrets

`func (o *BuildServiceFromComposeSpecRequest2) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *BuildServiceFromComposeSpecRequest2) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *BuildServiceFromComposeSpecRequest2) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *BuildServiceFromComposeSpecRequest2) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetServiceLogoURL

`func (o *BuildServiceFromComposeSpecRequest2) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *BuildServiceFromComposeSpecRequest2) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *BuildServiceFromComposeSpecRequest2) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *BuildServiceFromComposeSpecRequest2) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


