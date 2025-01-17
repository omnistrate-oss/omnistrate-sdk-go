# ListProductTiersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | ID of a Service | 
**ServiceModelId** | **string** | ID of a Service Model | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListProductTiersRequest

`func NewListProductTiersRequest(serviceId string, serviceModelId string, token string, ) *ListProductTiersRequest`

NewListProductTiersRequest instantiates a new ListProductTiersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProductTiersRequestWithDefaults

`func NewListProductTiersRequestWithDefaults() *ListProductTiersRequest`

NewListProductTiersRequestWithDefaults instantiates a new ListProductTiersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ListProductTiersRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListProductTiersRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListProductTiersRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceModelId

`func (o *ListProductTiersRequest) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *ListProductTiersRequest) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *ListProductTiersRequest) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetToken

`func (o *ListProductTiersRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListProductTiersRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListProductTiersRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


