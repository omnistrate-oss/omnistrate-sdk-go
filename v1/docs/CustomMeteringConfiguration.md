# CustomMeteringConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**[]CustomMeteringMetric**](CustomMeteringMetric.md) | Custom metrics used by the product tier billing configuration | 

## Methods

### NewCustomMeteringConfiguration

`func NewCustomMeteringConfiguration(metrics []CustomMeteringMetric, ) *CustomMeteringConfiguration`

NewCustomMeteringConfiguration instantiates a new CustomMeteringConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomMeteringConfigurationWithDefaults

`func NewCustomMeteringConfigurationWithDefaults() *CustomMeteringConfiguration`

NewCustomMeteringConfigurationWithDefaults instantiates a new CustomMeteringConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *CustomMeteringConfiguration) GetMetrics() []CustomMeteringMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CustomMeteringConfiguration) GetMetricsOk() (*[]CustomMeteringMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CustomMeteringConfiguration) SetMetrics(v []CustomMeteringMetric)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


