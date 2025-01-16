# CreateServiceAPIRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the service API bundle | 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateServiceAPIRequest

`func NewCreateServiceAPIRequest(description string, serviceEnvironmentId string, serviceId string, token string, ) *CreateServiceAPIRequest`

NewCreateServiceAPIRequest instantiates a new CreateServiceAPIRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAPIRequestWithDefaults

`func NewCreateServiceAPIRequestWithDefaults() *CreateServiceAPIRequest`

NewCreateServiceAPIRequestWithDefaults instantiates a new CreateServiceAPIRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateServiceAPIRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceAPIRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceAPIRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetServiceEnvironmentId

`func (o *CreateServiceAPIRequest) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *CreateServiceAPIRequest) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *CreateServiceAPIRequest) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceId

`func (o *CreateServiceAPIRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateServiceAPIRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateServiceAPIRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *CreateServiceAPIRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateServiceAPIRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateServiceAPIRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


