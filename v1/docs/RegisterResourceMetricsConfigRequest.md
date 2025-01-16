# RegisterResourceMetricsConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a resource | 
**MetricEndpoint** | **string** | The local host endpoint to supply prometheus metric | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRegisterResourceMetricsConfigRequest

`func NewRegisterResourceMetricsConfigRequest(id string, metricEndpoint string, serviceId string, token string, ) *RegisterResourceMetricsConfigRequest`

NewRegisterResourceMetricsConfigRequest instantiates a new RegisterResourceMetricsConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterResourceMetricsConfigRequestWithDefaults

`func NewRegisterResourceMetricsConfigRequestWithDefaults() *RegisterResourceMetricsConfigRequest`

NewRegisterResourceMetricsConfigRequestWithDefaults instantiates a new RegisterResourceMetricsConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegisterResourceMetricsConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterResourceMetricsConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterResourceMetricsConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMetricEndpoint

`func (o *RegisterResourceMetricsConfigRequest) GetMetricEndpoint() string`

GetMetricEndpoint returns the MetricEndpoint field if non-nil, zero value otherwise.

### GetMetricEndpointOk

`func (o *RegisterResourceMetricsConfigRequest) GetMetricEndpointOk() (*string, bool)`

GetMetricEndpointOk returns a tuple with the MetricEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricEndpoint

`func (o *RegisterResourceMetricsConfigRequest) SetMetricEndpoint(v string)`

SetMetricEndpoint sets MetricEndpoint field to given value.


### GetServiceId

`func (o *RegisterResourceMetricsConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RegisterResourceMetricsConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RegisterResourceMetricsConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *RegisterResourceMetricsConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RegisterResourceMetricsConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RegisterResourceMetricsConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


