# ListServiceModelsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceApiId** | **string** | ID of a Service API | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListServiceModelsRequest

`func NewListServiceModelsRequest(serviceApiId string, serviceId string, token string, ) *ListServiceModelsRequest`

NewListServiceModelsRequest instantiates a new ListServiceModelsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceModelsRequestWithDefaults

`func NewListServiceModelsRequestWithDefaults() *ListServiceModelsRequest`

NewListServiceModelsRequestWithDefaults instantiates a new ListServiceModelsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceApiId

`func (o *ListServiceModelsRequest) GetServiceApiId() string`

GetServiceApiId returns the ServiceApiId field if non-nil, zero value otherwise.

### GetServiceApiIdOk

`func (o *ListServiceModelsRequest) GetServiceApiIdOk() (*string, bool)`

GetServiceApiIdOk returns a tuple with the ServiceApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiId

`func (o *ListServiceModelsRequest) SetServiceApiId(v string)`

SetServiceApiId sets ServiceApiId field to given value.


### GetServiceId

`func (o *ListServiceModelsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListServiceModelsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListServiceModelsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ListServiceModelsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListServiceModelsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListServiceModelsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


