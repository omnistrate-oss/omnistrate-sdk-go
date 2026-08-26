# CustomMeteringMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationFunction** | **string** | Aggregation function applied to a custom metering metric | 
**Name** | **string** | Case-sensitive custom metric name | 

## Methods

### NewCustomMeteringMetric

`func NewCustomMeteringMetric(aggregationFunction string, name string, ) *CustomMeteringMetric`

NewCustomMeteringMetric instantiates a new CustomMeteringMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomMeteringMetricWithDefaults

`func NewCustomMeteringMetricWithDefaults() *CustomMeteringMetric`

NewCustomMeteringMetricWithDefaults instantiates a new CustomMeteringMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationFunction

`func (o *CustomMeteringMetric) GetAggregationFunction() string`

GetAggregationFunction returns the AggregationFunction field if non-nil, zero value otherwise.

### GetAggregationFunctionOk

`func (o *CustomMeteringMetric) GetAggregationFunctionOk() (*string, bool)`

GetAggregationFunctionOk returns a tuple with the AggregationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFunction

`func (o *CustomMeteringMetric) SetAggregationFunction(v string)`

SetAggregationFunction sets AggregationFunction field to given value.


### GetName

`func (o *CustomMeteringMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomMeteringMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomMeteringMetric) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


