# PromoteServiceEnvironmentStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Service Environment | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewPromoteServiceEnvironmentStatusRequest

`func NewPromoteServiceEnvironmentStatusRequest(id string, serviceId string, token string, ) *PromoteServiceEnvironmentStatusRequest`

NewPromoteServiceEnvironmentStatusRequest instantiates a new PromoteServiceEnvironmentStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoteServiceEnvironmentStatusRequestWithDefaults

`func NewPromoteServiceEnvironmentStatusRequestWithDefaults() *PromoteServiceEnvironmentStatusRequest`

NewPromoteServiceEnvironmentStatusRequestWithDefaults instantiates a new PromoteServiceEnvironmentStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromoteServiceEnvironmentStatusRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromoteServiceEnvironmentStatusRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromoteServiceEnvironmentStatusRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *PromoteServiceEnvironmentStatusRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PromoteServiceEnvironmentStatusRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PromoteServiceEnvironmentStatusRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *PromoteServiceEnvironmentStatusRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PromoteServiceEnvironmentStatusRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PromoteServiceEnvironmentStatusRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


