# PrepareServiceFromServicePlanSpecRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** | The environment to build the service in | 
**EnvironmentType** | **string** | The type of the environment | 
**FileContent** | **string** | Base64 encoded Compose Spec YAML in service plan configuration format | 
**Name** | **string** | Name of the Service | 

## Methods

### NewPrepareServiceFromServicePlanSpecRequest2

`func NewPrepareServiceFromServicePlanSpecRequest2(environment string, environmentType string, fileContent string, name string, ) *PrepareServiceFromServicePlanSpecRequest2`

NewPrepareServiceFromServicePlanSpecRequest2 instantiates a new PrepareServiceFromServicePlanSpecRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareServiceFromServicePlanSpecRequest2WithDefaults

`func NewPrepareServiceFromServicePlanSpecRequest2WithDefaults() *PrepareServiceFromServicePlanSpecRequest2`

NewPrepareServiceFromServicePlanSpecRequest2WithDefaults instantiates a new PrepareServiceFromServicePlanSpecRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *PrepareServiceFromServicePlanSpecRequest2) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PrepareServiceFromServicePlanSpecRequest2) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PrepareServiceFromServicePlanSpecRequest2) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetEnvironmentType

`func (o *PrepareServiceFromServicePlanSpecRequest2) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *PrepareServiceFromServicePlanSpecRequest2) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *PrepareServiceFromServicePlanSpecRequest2) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetFileContent

`func (o *PrepareServiceFromServicePlanSpecRequest2) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *PrepareServiceFromServicePlanSpecRequest2) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *PrepareServiceFromServicePlanSpecRequest2) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetName

`func (o *PrepareServiceFromServicePlanSpecRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrepareServiceFromServicePlanSpecRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrepareServiceFromServicePlanSpecRequest2) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


