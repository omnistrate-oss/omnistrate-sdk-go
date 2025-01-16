# CreatePipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the pipeline | 
**Name** | **string** | Name of the pipeline | 
**ServiceEnvironmentSequence** | **[]string** | Sequence of service environments to be deployed in the pipeline | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreatePipelineRequest

`func NewCreatePipelineRequest(description string, name string, serviceEnvironmentSequence []string, serviceId string, token string, ) *CreatePipelineRequest`

NewCreatePipelineRequest instantiates a new CreatePipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePipelineRequestWithDefaults

`func NewCreatePipelineRequestWithDefaults() *CreatePipelineRequest`

NewCreatePipelineRequestWithDefaults instantiates a new CreatePipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreatePipelineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePipelineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePipelineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreatePipelineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePipelineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePipelineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServiceEnvironmentSequence

`func (o *CreatePipelineRequest) GetServiceEnvironmentSequence() []string`

GetServiceEnvironmentSequence returns the ServiceEnvironmentSequence field if non-nil, zero value otherwise.

### GetServiceEnvironmentSequenceOk

`func (o *CreatePipelineRequest) GetServiceEnvironmentSequenceOk() (*[]string, bool)`

GetServiceEnvironmentSequenceOk returns a tuple with the ServiceEnvironmentSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentSequence

`func (o *CreatePipelineRequest) SetServiceEnvironmentSequence(v []string)`

SetServiceEnvironmentSequence sets ServiceEnvironmentSequence field to given value.


### GetServiceId

`func (o *CreatePipelineRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreatePipelineRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreatePipelineRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *CreatePipelineRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreatePipelineRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreatePipelineRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


