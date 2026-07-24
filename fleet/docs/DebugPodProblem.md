# DebugPodProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Workflow error taxonomy code, e.g. ImagePullError, PodUnschedulable, CrashLoop. | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** | Verbatim kubelet/scheduler message. | [optional] 
**Pod** | **string** |  | 
**Reason** | **string** |  | 
**State** | Pointer to **string** | The container/pod state: Waiting|Terminated|Pending | [optional] 

## Methods

### NewDebugPodProblem

`func NewDebugPodProblem(pod string, reason string, ) *DebugPodProblem`

NewDebugPodProblem instantiates a new DebugPodProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebugPodProblemWithDefaults

`func NewDebugPodProblemWithDefaults() *DebugPodProblem`

NewDebugPodProblemWithDefaults instantiates a new DebugPodProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DebugPodProblem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DebugPodProblem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DebugPodProblem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DebugPodProblem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainer

`func (o *DebugPodProblem) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *DebugPodProblem) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *DebugPodProblem) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *DebugPodProblem) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetMessage

`func (o *DebugPodProblem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DebugPodProblem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DebugPodProblem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DebugPodProblem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPod

`func (o *DebugPodProblem) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *DebugPodProblem) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *DebugPodProblem) SetPod(v string)`

SetPod sets Pod field to given value.


### GetReason

`func (o *DebugPodProblem) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DebugPodProblem) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DebugPodProblem) SetReason(v string)`

SetReason sets Reason field to given value.


### GetState

`func (o *DebugPodProblem) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DebugPodProblem) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DebugPodProblem) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DebugPodProblem) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


