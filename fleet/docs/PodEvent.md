# PodEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstTimestamp** | Pointer to **string** | The first timestamp when the pod event occurred | [optional] 
**LastTimestamp** | Pointer to **string** | The last timestamp when the pod event occurred | [optional] 
**Message** | Pointer to **string** | The message associated with the pod event | [optional] 
**NumberOfOccurrences** | Pointer to **int64** | The number of occurrences of the pod event | [optional] 
**Reason** | Pointer to **string** | The reason for the pod event | [optional] 

## Methods

### NewPodEvent

`func NewPodEvent() *PodEvent`

NewPodEvent instantiates a new PodEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodEventWithDefaults

`func NewPodEventWithDefaults() *PodEvent`

NewPodEventWithDefaults instantiates a new PodEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstTimestamp

`func (o *PodEvent) GetFirstTimestamp() string`

GetFirstTimestamp returns the FirstTimestamp field if non-nil, zero value otherwise.

### GetFirstTimestampOk

`func (o *PodEvent) GetFirstTimestampOk() (*string, bool)`

GetFirstTimestampOk returns a tuple with the FirstTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimestamp

`func (o *PodEvent) SetFirstTimestamp(v string)`

SetFirstTimestamp sets FirstTimestamp field to given value.

### HasFirstTimestamp

`func (o *PodEvent) HasFirstTimestamp() bool`

HasFirstTimestamp returns a boolean if a field has been set.

### GetLastTimestamp

`func (o *PodEvent) GetLastTimestamp() string`

GetLastTimestamp returns the LastTimestamp field if non-nil, zero value otherwise.

### GetLastTimestampOk

`func (o *PodEvent) GetLastTimestampOk() (*string, bool)`

GetLastTimestampOk returns a tuple with the LastTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimestamp

`func (o *PodEvent) SetLastTimestamp(v string)`

SetLastTimestamp sets LastTimestamp field to given value.

### HasLastTimestamp

`func (o *PodEvent) HasLastTimestamp() bool`

HasLastTimestamp returns a boolean if a field has been set.

### GetMessage

`func (o *PodEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PodEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PodEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PodEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNumberOfOccurrences

`func (o *PodEvent) GetNumberOfOccurrences() int64`

GetNumberOfOccurrences returns the NumberOfOccurrences field if non-nil, zero value otherwise.

### GetNumberOfOccurrencesOk

`func (o *PodEvent) GetNumberOfOccurrencesOk() (*int64, bool)`

GetNumberOfOccurrencesOk returns a tuple with the NumberOfOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOccurrences

`func (o *PodEvent) SetNumberOfOccurrences(v int64)`

SetNumberOfOccurrences sets NumberOfOccurrences field to given value.

### HasNumberOfOccurrences

`func (o *PodEvent) HasNumberOfOccurrences() bool`

HasNumberOfOccurrences returns a boolean if a field has been set.

### GetReason

`func (o *PodEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PodEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PodEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PodEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


