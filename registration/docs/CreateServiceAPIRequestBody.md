# CreateServiceAPIRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the service API bundle | 
**ServiceEnvironmentId** | **string** | The service environment ID | 

## Methods

### NewCreateServiceAPIRequestBody

`func NewCreateServiceAPIRequestBody(description string, serviceEnvironmentId string, ) *CreateServiceAPIRequestBody`

NewCreateServiceAPIRequestBody instantiates a new CreateServiceAPIRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAPIRequestBodyWithDefaults

`func NewCreateServiceAPIRequestBodyWithDefaults() *CreateServiceAPIRequestBody`

NewCreateServiceAPIRequestBodyWithDefaults instantiates a new CreateServiceAPIRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateServiceAPIRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceAPIRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceAPIRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetServiceEnvironmentId

`func (o *CreateServiceAPIRequestBody) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *CreateServiceAPIRequestBody) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *CreateServiceAPIRequestBody) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


