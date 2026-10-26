# InstanceHealthHistoryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | [**[]InstanceHealthHistoryEntry**](InstanceHealthHistoryEntry.md) | Recorded changes in the requested period | 
**EndTime** | **time.Time** | Exclusive end of the requested period | 
**InitialStatus** | Pointer to [**InstanceHealthHistoryEntry**](InstanceHealthHistoryEntry.md) |  | [optional] 
**InstanceId** | **string** | ID of a Resource Instance | 
**NextPageToken** | Pointer to **string** | Continuation token for the next page; repeat the same instance and time range | [optional] 
**StartTime** | **time.Time** | Inclusive start of the requested period | 

## Methods

### NewInstanceHealthHistoryResult

`func NewInstanceHealthHistoryResult(changes []InstanceHealthHistoryEntry, endTime time.Time, instanceId string, startTime time.Time, ) *InstanceHealthHistoryResult`

NewInstanceHealthHistoryResult instantiates a new InstanceHealthHistoryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceHealthHistoryResultWithDefaults

`func NewInstanceHealthHistoryResultWithDefaults() *InstanceHealthHistoryResult`

NewInstanceHealthHistoryResultWithDefaults instantiates a new InstanceHealthHistoryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *InstanceHealthHistoryResult) GetChanges() []InstanceHealthHistoryEntry`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *InstanceHealthHistoryResult) GetChangesOk() (*[]InstanceHealthHistoryEntry, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *InstanceHealthHistoryResult) SetChanges(v []InstanceHealthHistoryEntry)`

SetChanges sets Changes field to given value.


### GetEndTime

`func (o *InstanceHealthHistoryResult) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InstanceHealthHistoryResult) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InstanceHealthHistoryResult) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetInitialStatus

`func (o *InstanceHealthHistoryResult) GetInitialStatus() InstanceHealthHistoryEntry`

GetInitialStatus returns the InitialStatus field if non-nil, zero value otherwise.

### GetInitialStatusOk

`func (o *InstanceHealthHistoryResult) GetInitialStatusOk() (*InstanceHealthHistoryEntry, bool)`

GetInitialStatusOk returns a tuple with the InitialStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialStatus

`func (o *InstanceHealthHistoryResult) SetInitialStatus(v InstanceHealthHistoryEntry)`

SetInitialStatus sets InitialStatus field to given value.

### HasInitialStatus

`func (o *InstanceHealthHistoryResult) HasInitialStatus() bool`

HasInitialStatus returns a boolean if a field has been set.

### GetInstanceId

`func (o *InstanceHealthHistoryResult) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceHealthHistoryResult) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceHealthHistoryResult) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetNextPageToken

`func (o *InstanceHealthHistoryResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *InstanceHealthHistoryResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *InstanceHealthHistoryResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *InstanceHealthHistoryResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetStartTime

`func (o *InstanceHealthHistoryResult) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InstanceHealthHistoryResult) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InstanceHealthHistoryResult) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


