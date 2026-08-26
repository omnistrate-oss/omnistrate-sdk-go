# GetCustomMetricsEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresInSeconds** | Pointer to **int64** | Optional lifetime in seconds for the access token. It cannot exceed the configured maximum validity. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewGetCustomMetricsEndpointRequest

`func NewGetCustomMetricsEndpointRequest(token string, ) *GetCustomMetricsEndpointRequest`

NewGetCustomMetricsEndpointRequest instantiates a new GetCustomMetricsEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomMetricsEndpointRequestWithDefaults

`func NewGetCustomMetricsEndpointRequestWithDefaults() *GetCustomMetricsEndpointRequest`

NewGetCustomMetricsEndpointRequestWithDefaults instantiates a new GetCustomMetricsEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresInSeconds

`func (o *GetCustomMetricsEndpointRequest) GetExpiresInSeconds() int64`

GetExpiresInSeconds returns the ExpiresInSeconds field if non-nil, zero value otherwise.

### GetExpiresInSecondsOk

`func (o *GetCustomMetricsEndpointRequest) GetExpiresInSecondsOk() (*int64, bool)`

GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInSeconds

`func (o *GetCustomMetricsEndpointRequest) SetExpiresInSeconds(v int64)`

SetExpiresInSeconds sets ExpiresInSeconds field to given value.

### HasExpiresInSeconds

`func (o *GetCustomMetricsEndpointRequest) HasExpiresInSeconds() bool`

HasExpiresInSeconds returns a boolean if a field has been set.

### GetToken

`func (o *GetCustomMetricsEndpointRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetCustomMetricsEndpointRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetCustomMetricsEndpointRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


