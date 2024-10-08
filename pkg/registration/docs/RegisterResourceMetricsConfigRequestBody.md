# RegisterResourceMetricsConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricEndpoint** | **string** | The local host endpoint to supply prometheus metric | 

## Methods

### NewRegisterResourceMetricsConfigRequestBody

`func NewRegisterResourceMetricsConfigRequestBody(metricEndpoint string, ) *RegisterResourceMetricsConfigRequestBody`

NewRegisterResourceMetricsConfigRequestBody instantiates a new RegisterResourceMetricsConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterResourceMetricsConfigRequestBodyWithDefaults

`func NewRegisterResourceMetricsConfigRequestBodyWithDefaults() *RegisterResourceMetricsConfigRequestBody`

NewRegisterResourceMetricsConfigRequestBodyWithDefaults instantiates a new RegisterResourceMetricsConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricEndpoint

`func (o *RegisterResourceMetricsConfigRequestBody) GetMetricEndpoint() string`

GetMetricEndpoint returns the MetricEndpoint field if non-nil, zero value otherwise.

### GetMetricEndpointOk

`func (o *RegisterResourceMetricsConfigRequestBody) GetMetricEndpointOk() (*string, bool)`

GetMetricEndpointOk returns a tuple with the MetricEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricEndpoint

`func (o *RegisterResourceMetricsConfigRequestBody) SetMetricEndpoint(v string)`

SetMetricEndpoint sets MetricEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


