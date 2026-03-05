# PrepareServiceFromServicePlanSpecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** | The environment to build the service in | 
**EnvironmentType** | **string** | The type of service environment | 
**FileContent** | **string** | Base64 encoded Compose Spec YAML in service plan configuration format | 
**Name** | **string** | Name of the Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewPrepareServiceFromServicePlanSpecRequest

`func NewPrepareServiceFromServicePlanSpecRequest(environment string, environmentType string, fileContent string, name string, token string, ) *PrepareServiceFromServicePlanSpecRequest`

NewPrepareServiceFromServicePlanSpecRequest instantiates a new PrepareServiceFromServicePlanSpecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareServiceFromServicePlanSpecRequestWithDefaults

`func NewPrepareServiceFromServicePlanSpecRequestWithDefaults() *PrepareServiceFromServicePlanSpecRequest`

NewPrepareServiceFromServicePlanSpecRequestWithDefaults instantiates a new PrepareServiceFromServicePlanSpecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *PrepareServiceFromServicePlanSpecRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PrepareServiceFromServicePlanSpecRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PrepareServiceFromServicePlanSpecRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetEnvironmentType

`func (o *PrepareServiceFromServicePlanSpecRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *PrepareServiceFromServicePlanSpecRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *PrepareServiceFromServicePlanSpecRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetFileContent

`func (o *PrepareServiceFromServicePlanSpecRequest) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *PrepareServiceFromServicePlanSpecRequest) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *PrepareServiceFromServicePlanSpecRequest) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetName

`func (o *PrepareServiceFromServicePlanSpecRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrepareServiceFromServicePlanSpecRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrepareServiceFromServicePlanSpecRequest) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *PrepareServiceFromServicePlanSpecRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PrepareServiceFromServicePlanSpecRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PrepareServiceFromServicePlanSpecRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


