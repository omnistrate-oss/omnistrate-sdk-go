# Describeresourcemetricsconfigresult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the resource | 
**MetricEndpoint** | **string** | The local host endpoint to supply prometheus metric | 
**ServiceId** | **string** | The service ID that this API bundle belongs to | 

## Methods

### NewDescriberesourcemetricsconfigresult

`func NewDescriberesourcemetricsconfigresult(id string, metricEndpoint string, serviceId string, ) *Describeresourcemetricsconfigresult`

NewDescriberesourcemetricsconfigresult instantiates a new Describeresourcemetricsconfigresult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescriberesourcemetricsconfigresultWithDefaults

`func NewDescriberesourcemetricsconfigresultWithDefaults() *Describeresourcemetricsconfigresult`

NewDescriberesourcemetricsconfigresultWithDefaults instantiates a new Describeresourcemetricsconfigresult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Describeresourcemetricsconfigresult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Describeresourcemetricsconfigresult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Describeresourcemetricsconfigresult) SetId(v string)`

SetId sets Id field to given value.


### GetMetricEndpoint

`func (o *Describeresourcemetricsconfigresult) GetMetricEndpoint() string`

GetMetricEndpoint returns the MetricEndpoint field if non-nil, zero value otherwise.

### GetMetricEndpointOk

`func (o *Describeresourcemetricsconfigresult) GetMetricEndpointOk() (*string, bool)`

GetMetricEndpointOk returns a tuple with the MetricEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricEndpoint

`func (o *Describeresourcemetricsconfigresult) SetMetricEndpoint(v string)`

SetMetricEndpoint sets MetricEndpoint field to given value.


### GetServiceId

`func (o *Describeresourcemetricsconfigresult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Describeresourcemetricsconfigresult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Describeresourcemetricsconfigresult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


