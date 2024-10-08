# BuildServiceFromServicePlanSpecRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the service | [optional] 
**Environment** | Pointer to **string** | The environment to build the service in | [optional] 
**EnvironmentType** | Pointer to **string** | The type of the environment | [optional] 
**FileContent** | **string** | Base64 encoded Compose Spec YAML in service plan configuration format | 
**Name** | **string** | Name of the Service | 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 

## Methods

### NewBuildServiceFromServicePlanSpecRequestBody

`func NewBuildServiceFromServicePlanSpecRequestBody(fileContent string, name string, ) *BuildServiceFromServicePlanSpecRequestBody`

NewBuildServiceFromServicePlanSpecRequestBody instantiates a new BuildServiceFromServicePlanSpecRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildServiceFromServicePlanSpecRequestBodyWithDefaults

`func NewBuildServiceFromServicePlanSpecRequestBodyWithDefaults() *BuildServiceFromServicePlanSpecRequestBody`

NewBuildServiceFromServicePlanSpecRequestBodyWithDefaults instantiates a new BuildServiceFromServicePlanSpecRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuildServiceFromServicePlanSpecRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BuildServiceFromServicePlanSpecRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BuildServiceFromServicePlanSpecRequestBody) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BuildServiceFromServicePlanSpecRequestBody) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *BuildServiceFromServicePlanSpecRequestBody) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *BuildServiceFromServicePlanSpecRequestBody) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetFileContent

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *BuildServiceFromServicePlanSpecRequestBody) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetName

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuildServiceFromServicePlanSpecRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetServiceLogoURL

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *BuildServiceFromServicePlanSpecRequestBody) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *BuildServiceFromServicePlanSpecRequestBody) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *BuildServiceFromServicePlanSpecRequestBody) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


