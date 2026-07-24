# DebugWarningEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**LastSeen** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**Reason** | **string** |  | 

## Methods

### NewDebugWarningEvent

`func NewDebugWarningEvent(message string, reason string, ) *DebugWarningEvent`

NewDebugWarningEvent instantiates a new DebugWarningEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebugWarningEventWithDefaults

`func NewDebugWarningEventWithDefaults() *DebugWarningEvent`

NewDebugWarningEventWithDefaults instantiates a new DebugWarningEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DebugWarningEvent) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DebugWarningEvent) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DebugWarningEvent) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *DebugWarningEvent) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastSeen

`func (o *DebugWarningEvent) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *DebugWarningEvent) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *DebugWarningEvent) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *DebugWarningEvent) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMessage

`func (o *DebugWarningEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DebugWarningEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DebugWarningEvent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *DebugWarningEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DebugWarningEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DebugWarningEvent) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


