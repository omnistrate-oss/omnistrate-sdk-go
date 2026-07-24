# TaskAttemptError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Stable error code from the workflow error taxonomy. | 
**DisplayMessage** | Pointer to **string** | A concise human-readable summary derived from the code. | [optional] 
**FirstSeenAt** | Pointer to **string** | The time this error signature was first observed for the current task attempt, in RFC3339 format. | [optional] 
**Message** | **string** | Detailed error message for diagnosis. | 

## Methods

### NewTaskAttemptError

`func NewTaskAttemptError(code string, message string, ) *TaskAttemptError`

NewTaskAttemptError instantiates a new TaskAttemptError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskAttemptErrorWithDefaults

`func NewTaskAttemptErrorWithDefaults() *TaskAttemptError`

NewTaskAttemptErrorWithDefaults instantiates a new TaskAttemptError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TaskAttemptError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaskAttemptError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaskAttemptError) SetCode(v string)`

SetCode sets Code field to given value.


### GetDisplayMessage

`func (o *TaskAttemptError) GetDisplayMessage() string`

GetDisplayMessage returns the DisplayMessage field if non-nil, zero value otherwise.

### GetDisplayMessageOk

`func (o *TaskAttemptError) GetDisplayMessageOk() (*string, bool)`

GetDisplayMessageOk returns a tuple with the DisplayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMessage

`func (o *TaskAttemptError) SetDisplayMessage(v string)`

SetDisplayMessage sets DisplayMessage field to given value.

### HasDisplayMessage

`func (o *TaskAttemptError) HasDisplayMessage() bool`

HasDisplayMessage returns a boolean if a field has been set.

### GetFirstSeenAt

`func (o *TaskAttemptError) GetFirstSeenAt() string`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *TaskAttemptError) GetFirstSeenAtOk() (*string, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *TaskAttemptError) SetFirstSeenAt(v string)`

SetFirstSeenAt sets FirstSeenAt field to given value.

### HasFirstSeenAt

`func (o *TaskAttemptError) HasFirstSeenAt() bool`

HasFirstSeenAt returns a boolean if a field has been set.

### GetMessage

`func (o *TaskAttemptError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TaskAttemptError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TaskAttemptError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


