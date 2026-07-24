# PodProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Stable error code from the workflow error taxonomy. | [optional] 
**Container** | Pointer to **string** | The name of the container experiencing the problem, when applicable. | [optional] 
**Message** | Pointer to **string** | The detailed message associated with the pod or container state. | [optional] 
**Pod** | Pointer to **string** | The name of the pod experiencing the problem. | [optional] 
**Reason** | Pointer to **string** | The reason associated with the pod or container state. | [optional] 
**State** | Pointer to **string** | The observed pod or container state. | [optional] 

## Methods

### NewPodProblem

`func NewPodProblem() *PodProblem`

NewPodProblem instantiates a new PodProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodProblemWithDefaults

`func NewPodProblemWithDefaults() *PodProblem`

NewPodProblemWithDefaults instantiates a new PodProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PodProblem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PodProblem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PodProblem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PodProblem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainer

`func (o *PodProblem) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *PodProblem) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *PodProblem) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *PodProblem) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetMessage

`func (o *PodProblem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PodProblem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PodProblem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PodProblem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPod

`func (o *PodProblem) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *PodProblem) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *PodProblem) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *PodProblem) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetReason

`func (o *PodProblem) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PodProblem) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PodProblem) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PodProblem) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetState

`func (o *PodProblem) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PodProblem) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PodProblem) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PodProblem) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


