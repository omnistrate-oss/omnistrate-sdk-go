# AutoScalingMetricSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricEndpoint** | **string** | The local host endpoint to supply prometheus metric | 
**MetricLabelName** | **string** | The prometheus metric label name for scaling | 
**MetricLabelValue** | **string** | The prometheus metric label value for scaling | 
**MetricName** | **string** | The prometheus metric name for scaling | 

## Methods

### NewAutoScalingMetricSpec

`func NewAutoScalingMetricSpec(metricEndpoint string, metricLabelName string, metricLabelValue string, metricName string, ) *AutoScalingMetricSpec`

NewAutoScalingMetricSpec instantiates a new AutoScalingMetricSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingMetricSpecWithDefaults

`func NewAutoScalingMetricSpecWithDefaults() *AutoScalingMetricSpec`

NewAutoScalingMetricSpecWithDefaults instantiates a new AutoScalingMetricSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricEndpoint

`func (o *AutoScalingMetricSpec) GetMetricEndpoint() string`

GetMetricEndpoint returns the MetricEndpoint field if non-nil, zero value otherwise.

### GetMetricEndpointOk

`func (o *AutoScalingMetricSpec) GetMetricEndpointOk() (*string, bool)`

GetMetricEndpointOk returns a tuple with the MetricEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricEndpoint

`func (o *AutoScalingMetricSpec) SetMetricEndpoint(v string)`

SetMetricEndpoint sets MetricEndpoint field to given value.


### GetMetricLabelName

`func (o *AutoScalingMetricSpec) GetMetricLabelName() string`

GetMetricLabelName returns the MetricLabelName field if non-nil, zero value otherwise.

### GetMetricLabelNameOk

`func (o *AutoScalingMetricSpec) GetMetricLabelNameOk() (*string, bool)`

GetMetricLabelNameOk returns a tuple with the MetricLabelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLabelName

`func (o *AutoScalingMetricSpec) SetMetricLabelName(v string)`

SetMetricLabelName sets MetricLabelName field to given value.


### GetMetricLabelValue

`func (o *AutoScalingMetricSpec) GetMetricLabelValue() string`

GetMetricLabelValue returns the MetricLabelValue field if non-nil, zero value otherwise.

### GetMetricLabelValueOk

`func (o *AutoScalingMetricSpec) GetMetricLabelValueOk() (*string, bool)`

GetMetricLabelValueOk returns a tuple with the MetricLabelValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLabelValue

`func (o *AutoScalingMetricSpec) SetMetricLabelValue(v string)`

SetMetricLabelValue sets MetricLabelValue field to given value.


### GetMetricName

`func (o *AutoScalingMetricSpec) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *AutoScalingMetricSpec) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *AutoScalingMetricSpec) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


