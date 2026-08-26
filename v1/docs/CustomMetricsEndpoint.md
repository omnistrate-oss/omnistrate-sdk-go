# CustomMetricsEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | Complete URL that accepts custom metric usage. | 
**ExpiresAt** | **time.Time** | RFC3339 expiration timestamp for the returned access token. | 
**Token** | **string** | Access token scoped to the authenticated organization. | 

## Methods

### NewCustomMetricsEndpoint

`func NewCustomMetricsEndpoint(endpoint string, expiresAt time.Time, token string, ) *CustomMetricsEndpoint`

NewCustomMetricsEndpoint instantiates a new CustomMetricsEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomMetricsEndpointWithDefaults

`func NewCustomMetricsEndpointWithDefaults() *CustomMetricsEndpoint`

NewCustomMetricsEndpointWithDefaults instantiates a new CustomMetricsEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *CustomMetricsEndpoint) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CustomMetricsEndpoint) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CustomMetricsEndpoint) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetExpiresAt

`func (o *CustomMetricsEndpoint) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CustomMetricsEndpoint) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CustomMetricsEndpoint) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetToken

`func (o *CustomMetricsEndpoint) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CustomMetricsEndpoint) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CustomMetricsEndpoint) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


