# CheckpointCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **int64** | The number of units observed to have completed. Absent when not reported. | [optional] 
**Failed** | Pointer to **int64** | The number of units observed to have failed. Absent when not reported; a reported zero means &#39;none failed&#39;, not &#39;unknown&#39;. | [optional] 
**InFlight** | Pointer to **int64** | The number of units observed to be in flight. Absent when not reported. | [optional] 
**Total** | Pointer to **int64** | The denominator, present only when a trustworthy one was observed. Always absent when totalUnknown is true, and never emitted as zero to stand in for an unknown total. | [optional] 
**TotalUnknown** | Pointer to **bool** | True when no trustworthy denominator was observed, for example when no plan was captured for the operation. Consumers must render an indeterminate state and must not compute a ratio or a percentage. | [optional] 

## Methods

### NewCheckpointCounts

`func NewCheckpointCounts() *CheckpointCounts`

NewCheckpointCounts instantiates a new CheckpointCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointCountsWithDefaults

`func NewCheckpointCountsWithDefaults() *CheckpointCounts`

NewCheckpointCountsWithDefaults instantiates a new CheckpointCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *CheckpointCounts) GetCompleted() int64`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *CheckpointCounts) GetCompletedOk() (*int64, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *CheckpointCounts) SetCompleted(v int64)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *CheckpointCounts) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetFailed

`func (o *CheckpointCounts) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *CheckpointCounts) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *CheckpointCounts) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *CheckpointCounts) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetInFlight

`func (o *CheckpointCounts) GetInFlight() int64`

GetInFlight returns the InFlight field if non-nil, zero value otherwise.

### GetInFlightOk

`func (o *CheckpointCounts) GetInFlightOk() (*int64, bool)`

GetInFlightOk returns a tuple with the InFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInFlight

`func (o *CheckpointCounts) SetInFlight(v int64)`

SetInFlight sets InFlight field to given value.

### HasInFlight

`func (o *CheckpointCounts) HasInFlight() bool`

HasInFlight returns a boolean if a field has been set.

### GetTotal

`func (o *CheckpointCounts) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CheckpointCounts) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CheckpointCounts) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CheckpointCounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalUnknown

`func (o *CheckpointCounts) GetTotalUnknown() bool`

GetTotalUnknown returns the TotalUnknown field if non-nil, zero value otherwise.

### GetTotalUnknownOk

`func (o *CheckpointCounts) GetTotalUnknownOk() (*bool, bool)`

GetTotalUnknownOk returns a tuple with the TotalUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnknown

`func (o *CheckpointCounts) SetTotalUnknown(v bool)`

SetTotalUnknown sets TotalUnknown field to given value.

### HasTotalUnknown

`func (o *CheckpointCounts) HasTotalUnknown() bool`

HasTotalUnknown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


