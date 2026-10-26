# InstanceHealthHistoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | [**InstanceHealthSummary**](InstanceHealthSummary.md) |  | 
**Timestamp** | **time.Time** | UTC timestamp of the recorded change, preserving nanosecond precision | 

## Methods

### NewInstanceHealthHistoryEntry

`func NewInstanceHealthHistoryEntry(health InstanceHealthSummary, timestamp time.Time, ) *InstanceHealthHistoryEntry`

NewInstanceHealthHistoryEntry instantiates a new InstanceHealthHistoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceHealthHistoryEntryWithDefaults

`func NewInstanceHealthHistoryEntryWithDefaults() *InstanceHealthHistoryEntry`

NewInstanceHealthHistoryEntryWithDefaults instantiates a new InstanceHealthHistoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *InstanceHealthHistoryEntry) GetHealth() InstanceHealthSummary`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *InstanceHealthHistoryEntry) GetHealthOk() (*InstanceHealthSummary, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *InstanceHealthHistoryEntry) SetHealth(v InstanceHealthSummary)`

SetHealth sets Health field to given value.


### GetTimestamp

`func (o *InstanceHealthHistoryEntry) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InstanceHealthHistoryEntry) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InstanceHealthHistoryEntry) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


