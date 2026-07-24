# WarningEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of times this warning event has occurred. | [optional] 
**LastSeen** | Pointer to **string** | The time the warning event was last observed, in RFC3339 format. | [optional] 
**Message** | Pointer to **string** | The detailed message of the warning event. | [optional] 
**Reason** | Pointer to **string** | The reason for the warning event. | [optional] 

## Methods

### NewWarningEvent

`func NewWarningEvent() *WarningEvent`

NewWarningEvent instantiates a new WarningEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningEventWithDefaults

`func NewWarningEventWithDefaults() *WarningEvent`

NewWarningEventWithDefaults instantiates a new WarningEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *WarningEvent) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WarningEvent) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WarningEvent) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *WarningEvent) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastSeen

`func (o *WarningEvent) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *WarningEvent) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *WarningEvent) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *WarningEvent) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMessage

`func (o *WarningEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WarningEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WarningEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WarningEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *WarningEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WarningEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WarningEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WarningEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


