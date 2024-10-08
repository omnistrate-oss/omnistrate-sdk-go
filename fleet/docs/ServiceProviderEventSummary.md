# ServiceProviderEventSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventsSummary** | Pointer to **map[string]int64** | The number of outstanding events by type | [optional] 

## Methods

### NewServiceProviderEventSummary

`func NewServiceProviderEventSummary() *ServiceProviderEventSummary`

NewServiceProviderEventSummary instantiates a new ServiceProviderEventSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProviderEventSummaryWithDefaults

`func NewServiceProviderEventSummaryWithDefaults() *ServiceProviderEventSummary`

NewServiceProviderEventSummaryWithDefaults instantiates a new ServiceProviderEventSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventsSummary

`func (o *ServiceProviderEventSummary) GetEventsSummary() map[string]int64`

GetEventsSummary returns the EventsSummary field if non-nil, zero value otherwise.

### GetEventsSummaryOk

`func (o *ServiceProviderEventSummary) GetEventsSummaryOk() (*map[string]int64, bool)`

GetEventsSummaryOk returns a tuple with the EventsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsSummary

`func (o *ServiceProviderEventSummary) SetEventsSummary(v map[string]int64)`

SetEventsSummary sets EventsSummary field to given value.

### HasEventsSummary

`func (o *ServiceProviderEventSummary) HasEventsSummary() bool`

HasEventsSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


