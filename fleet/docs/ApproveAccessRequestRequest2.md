# ApproveAccessRequestRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMinutes** | Pointer to **int64** | MANUAL approval duration. Omission uses the finite backend default within 1–720; no client default, unlimited choice, active extension, or credential renewal. | [optional] 

## Methods

### NewApproveAccessRequestRequest2

`func NewApproveAccessRequestRequest2() *ApproveAccessRequestRequest2`

NewApproveAccessRequestRequest2 instantiates a new ApproveAccessRequestRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveAccessRequestRequest2WithDefaults

`func NewApproveAccessRequestRequest2WithDefaults() *ApproveAccessRequestRequest2`

NewApproveAccessRequestRequest2WithDefaults instantiates a new ApproveAccessRequestRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMinutes

`func (o *ApproveAccessRequestRequest2) GetDurationMinutes() int64`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *ApproveAccessRequestRequest2) GetDurationMinutesOk() (*int64, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *ApproveAccessRequestRequest2) SetDurationMinutes(v int64)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *ApproveAccessRequestRequest2) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


