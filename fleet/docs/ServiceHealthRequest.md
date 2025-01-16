# ServiceHealthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceEnvironmentID** | **string** | ID of a Service Environment | 
**ServiceID** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewServiceHealthRequest

`func NewServiceHealthRequest(serviceEnvironmentID string, serviceID string, token string, ) *ServiceHealthRequest`

NewServiceHealthRequest instantiates a new ServiceHealthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHealthRequestWithDefaults

`func NewServiceHealthRequestWithDefaults() *ServiceHealthRequest`

NewServiceHealthRequestWithDefaults instantiates a new ServiceHealthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceEnvironmentID

`func (o *ServiceHealthRequest) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *ServiceHealthRequest) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *ServiceHealthRequest) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.


### GetServiceID

`func (o *ServiceHealthRequest) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *ServiceHealthRequest) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *ServiceHealthRequest) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.


### GetToken

`func (o *ServiceHealthRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ServiceHealthRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ServiceHealthRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


