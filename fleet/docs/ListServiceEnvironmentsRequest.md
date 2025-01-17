# ListServiceEnvironmentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListServiceEnvironmentsRequest

`func NewListServiceEnvironmentsRequest(serviceId string, token string, ) *ListServiceEnvironmentsRequest`

NewListServiceEnvironmentsRequest instantiates a new ListServiceEnvironmentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceEnvironmentsRequestWithDefaults

`func NewListServiceEnvironmentsRequestWithDefaults() *ListServiceEnvironmentsRequest`

NewListServiceEnvironmentsRequestWithDefaults instantiates a new ListServiceEnvironmentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ListServiceEnvironmentsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListServiceEnvironmentsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListServiceEnvironmentsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ListServiceEnvironmentsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListServiceEnvironmentsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListServiceEnvironmentsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


