# BuildServiceFromServicePlanSpecRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the service | [optional] 
**Dryrun** | Pointer to **bool** | If set to true, performs a dry run of the build operation without making any changes | [optional] 
**Environment** | Pointer to **string** | The environment to build the service in | [optional] 
**EnvironmentType** | Pointer to **string** | The type of the environment | [optional] 
**FileContent** | **string** | Base64 encoded Compose Spec YAML in service plan configuration format | 
**Name** | **string** | Name of the Service | 
**Release** | Pointer to **bool** | Release the service after building | [optional] 
**ReleaseAsPreferred** | Pointer to **bool** | Release the service as preferred | [optional] 
**ReleaseVersionName** | Pointer to **string** | Release version name | [optional] 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 

## Methods

### NewBuildServiceFromServicePlanSpecRequest2

`func NewBuildServiceFromServicePlanSpecRequest2(fileContent string, name string, ) *BuildServiceFromServicePlanSpecRequest2`

NewBuildServiceFromServicePlanSpecRequest2 instantiates a new BuildServiceFromServicePlanSpecRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildServiceFromServicePlanSpecRequest2WithDefaults

`func NewBuildServiceFromServicePlanSpecRequest2WithDefaults() *BuildServiceFromServicePlanSpecRequest2`

NewBuildServiceFromServicePlanSpecRequest2WithDefaults instantiates a new BuildServiceFromServicePlanSpecRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BuildServiceFromServicePlanSpecRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuildServiceFromServicePlanSpecRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuildServiceFromServicePlanSpecRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BuildServiceFromServicePlanSpecRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDryrun

`func (o *BuildServiceFromServicePlanSpecRequest2) GetDryrun() bool`

GetDryrun returns the Dryrun field if non-nil, zero value otherwise.

### GetDryrunOk

`func (o *BuildServiceFromServicePlanSpecRequest2) GetDryrunOk() (*bool, bool)`

GetDryrunOk returns a tuple with the Dryrun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryrun

`func (o *BuildServiceFromServicePlanSpecRequest2) SetDryrun(v bool)`

SetDryrun sets Dryrun field to given value.

### HasDryrun

`func (o *BuildServiceFromServicePlanSpecRequest2) HasDryrun() bool`

HasDryrun returns a boolean if a field has been set.

### GetEnvironment

`func (o *BuildServiceFromServicePlanSpecRequest2) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BuildServiceFromServicePlanSpecRequest2) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BuildServiceFromServicePlanSpecRequest2) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BuildServiceFromServicePlanSpecRequest2) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *BuildServiceFromServicePlanSpecRequest2) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *BuildServiceFromServicePlanSpecRequest2) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *BuildServiceFromServicePlanSpecRequest2) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *BuildServiceFromServicePlanSpecRequest2) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetFileContent

`func (o *BuildServiceFromServicePlanSpecRequest2) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *BuildServiceFromServicePlanSpecRequest2) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *BuildServiceFromServicePlanSpecRequest2) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetName

`func (o *BuildServiceFromServicePlanSpecRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuildServiceFromServicePlanSpecRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuildServiceFromServicePlanSpecRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetRelease

`func (o *BuildServiceFromServicePlanSpecRequest2) GetRelease() bool`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *BuildServiceFromServicePlanSpecRequest2) GetReleaseOk() (*bool, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *BuildServiceFromServicePlanSpecRequest2) SetRelease(v bool)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *BuildServiceFromServicePlanSpecRequest2) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetReleaseAsPreferred

`func (o *BuildServiceFromServicePlanSpecRequest2) GetReleaseAsPreferred() bool`

GetReleaseAsPreferred returns the ReleaseAsPreferred field if non-nil, zero value otherwise.

### GetReleaseAsPreferredOk

`func (o *BuildServiceFromServicePlanSpecRequest2) GetReleaseAsPreferredOk() (*bool, bool)`

GetReleaseAsPreferredOk returns a tuple with the ReleaseAsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseAsPreferred

`func (o *BuildServiceFromServicePlanSpecRequest2) SetReleaseAsPreferred(v bool)`

SetReleaseAsPreferred sets ReleaseAsPreferred field to given value.

### HasReleaseAsPreferred

`func (o *BuildServiceFromServicePlanSpecRequest2) HasReleaseAsPreferred() bool`

HasReleaseAsPreferred returns a boolean if a field has been set.

### GetReleaseVersionName

`func (o *BuildServiceFromServicePlanSpecRequest2) GetReleaseVersionName() string`

GetReleaseVersionName returns the ReleaseVersionName field if non-nil, zero value otherwise.

### GetReleaseVersionNameOk

`func (o *BuildServiceFromServicePlanSpecRequest2) GetReleaseVersionNameOk() (*string, bool)`

GetReleaseVersionNameOk returns a tuple with the ReleaseVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersionName

`func (o *BuildServiceFromServicePlanSpecRequest2) SetReleaseVersionName(v string)`

SetReleaseVersionName sets ReleaseVersionName field to given value.

### HasReleaseVersionName

`func (o *BuildServiceFromServicePlanSpecRequest2) HasReleaseVersionName() bool`

HasReleaseVersionName returns a boolean if a field has been set.

### GetServiceLogoURL

`func (o *BuildServiceFromServicePlanSpecRequest2) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *BuildServiceFromServicePlanSpecRequest2) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *BuildServiceFromServicePlanSpecRequest2) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *BuildServiceFromServicePlanSpecRequest2) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


