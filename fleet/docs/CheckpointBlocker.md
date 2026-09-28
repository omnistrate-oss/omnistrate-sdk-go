# CheckpointBlocker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int64** | How many times the blocked unit has been tried and has hit this same condition. Absent when no attempt evidence was observed, which is not the same as zero attempts: a unit nobody counted retries for is not a unit that has never been tried, and only one of the two supports &#39;this has been failing repeatedly&#39;. It counts attempts at the blocked unit and never reconcile cycles of the observer, and it is evidence rather than identity - it must never be used as a component of a checkpoint id. | [optional] 
**Code** | Pointer to **string** | Stable blocker code from the workflow error taxonomy, e.g. ImagePullError, PodUnschedulable, CrashLoop, HookFailed, StateLocked, QuotaExceeded, InsufficientCapacity. | [optional] 
**FirstSeenAt** | Pointer to **string** | The time this blocker signature was first observed, in RFC3339 format. Absent when the source reported no first-seen evidence. | [optional] 
**Persistence** | Pointer to **string** | Whether the condition can clear on its own: transient|permanent. Transient is a state lock another operation holds, capacity that has not arrived, or a dependency still being created; permanent is a condition that waiting will not resolve, such as a hook Job whose backoff limit is exhausted or an apply blocked on a credential the workspace does not have. Absent when the producer has no evidence either way, which is the common case and must be rendered as unknown rather than as either answer - presenting an unknown as transient tells an operator to wait out a blocker that never clears. | [optional] 
**Reason** | Pointer to **string** | A short, sanitized, operator-facing reason for the blocker, capped by the producer at 512 characters. Producers must sanitize before setting this: provider and engine error text is untrusted and is known to echo variable values and secret material. | [optional] 
**ReasonTruncated** | Pointer to **bool** | True when reason was shortened to respect the producer cap. | [optional] 
**Subject** | Pointer to **string** | The name or address the blocker refers to: the Kubernetes object, Helm hook Job or Terraform resource address the condition is about, as opposed to the operation that observed it. It is a neutral identity - a kind and a name, or an engine resource address - and never an engine-internal identifier such as a Pulumi URN. Absent when the source reported no subject; consumers must not substitute the operation&#39;s own name, because &#39;the deployment is blocked&#39; and &#39;Job/queue-migration is blocked&#39; are different facts. | [optional] 

## Methods

### NewCheckpointBlocker

`func NewCheckpointBlocker() *CheckpointBlocker`

NewCheckpointBlocker instantiates a new CheckpointBlocker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointBlockerWithDefaults

`func NewCheckpointBlockerWithDefaults() *CheckpointBlocker`

NewCheckpointBlockerWithDefaults instantiates a new CheckpointBlocker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *CheckpointBlocker) GetAttempts() int64`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *CheckpointBlocker) GetAttemptsOk() (*int64, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *CheckpointBlocker) SetAttempts(v int64)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *CheckpointBlocker) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetCode

`func (o *CheckpointBlocker) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CheckpointBlocker) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CheckpointBlocker) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CheckpointBlocker) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFirstSeenAt

`func (o *CheckpointBlocker) GetFirstSeenAt() string`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *CheckpointBlocker) GetFirstSeenAtOk() (*string, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *CheckpointBlocker) SetFirstSeenAt(v string)`

SetFirstSeenAt sets FirstSeenAt field to given value.

### HasFirstSeenAt

`func (o *CheckpointBlocker) HasFirstSeenAt() bool`

HasFirstSeenAt returns a boolean if a field has been set.

### GetPersistence

`func (o *CheckpointBlocker) GetPersistence() string`

GetPersistence returns the Persistence field if non-nil, zero value otherwise.

### GetPersistenceOk

`func (o *CheckpointBlocker) GetPersistenceOk() (*string, bool)`

GetPersistenceOk returns a tuple with the Persistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistence

`func (o *CheckpointBlocker) SetPersistence(v string)`

SetPersistence sets Persistence field to given value.

### HasPersistence

`func (o *CheckpointBlocker) HasPersistence() bool`

HasPersistence returns a boolean if a field has been set.

### GetReason

`func (o *CheckpointBlocker) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CheckpointBlocker) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CheckpointBlocker) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CheckpointBlocker) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonTruncated

`func (o *CheckpointBlocker) GetReasonTruncated() bool`

GetReasonTruncated returns the ReasonTruncated field if non-nil, zero value otherwise.

### GetReasonTruncatedOk

`func (o *CheckpointBlocker) GetReasonTruncatedOk() (*bool, bool)`

GetReasonTruncatedOk returns a tuple with the ReasonTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonTruncated

`func (o *CheckpointBlocker) SetReasonTruncated(v bool)`

SetReasonTruncated sets ReasonTruncated field to given value.

### HasReasonTruncated

`func (o *CheckpointBlocker) HasReasonTruncated() bool`

HasReasonTruncated returns a boolean if a field has been set.

### GetSubject

`func (o *CheckpointBlocker) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CheckpointBlocker) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CheckpointBlocker) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CheckpointBlocker) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


