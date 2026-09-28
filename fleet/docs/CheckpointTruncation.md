# CheckpointTruncation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Returned** | Pointer to **int64** | The number of items actually returned in this response. | [optional] 
**Total** | Pointer to **int64** | The number of items the source observed, present only when known. Always absent when totalUnknown is true. | [optional] 
**TotalUnknown** | Pointer to **bool** | True when the source could not determine how many items exist. Consumers must not present the returned count as a complete total. | [optional] 
**Truncated** | Pointer to **bool** | True when the returned items are a subset of what the source observed. | [optional] 

## Methods

### NewCheckpointTruncation

`func NewCheckpointTruncation() *CheckpointTruncation`

NewCheckpointTruncation instantiates a new CheckpointTruncation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointTruncationWithDefaults

`func NewCheckpointTruncationWithDefaults() *CheckpointTruncation`

NewCheckpointTruncationWithDefaults instantiates a new CheckpointTruncation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturned

`func (o *CheckpointTruncation) GetReturned() int64`

GetReturned returns the Returned field if non-nil, zero value otherwise.

### GetReturnedOk

`func (o *CheckpointTruncation) GetReturnedOk() (*int64, bool)`

GetReturnedOk returns a tuple with the Returned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturned

`func (o *CheckpointTruncation) SetReturned(v int64)`

SetReturned sets Returned field to given value.

### HasReturned

`func (o *CheckpointTruncation) HasReturned() bool`

HasReturned returns a boolean if a field has been set.

### GetTotal

`func (o *CheckpointTruncation) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CheckpointTruncation) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CheckpointTruncation) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CheckpointTruncation) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalUnknown

`func (o *CheckpointTruncation) GetTotalUnknown() bool`

GetTotalUnknown returns the TotalUnknown field if non-nil, zero value otherwise.

### GetTotalUnknownOk

`func (o *CheckpointTruncation) GetTotalUnknownOk() (*bool, bool)`

GetTotalUnknownOk returns a tuple with the TotalUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnknown

`func (o *CheckpointTruncation) SetTotalUnknown(v bool)`

SetTotalUnknown sets TotalUnknown field to given value.

### HasTotalUnknown

`func (o *CheckpointTruncation) HasTotalUnknown() bool`

HasTotalUnknown returns a boolean if a field has been set.

### GetTruncated

`func (o *CheckpointTruncation) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *CheckpointTruncation) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *CheckpointTruncation) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *CheckpointTruncation) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


