# WorkflowTaskObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | apply, patch, delete or get | 
**ApiVersion** | Pointer to **string** |  | [optional] 
**ConditionSummary** | Pointer to **string** | The condition evaluation recorded when the task finished | [optional] 
**CustomResource** | Pointer to **bool** | The kind is served by a CustomResourceDefinition | [optional] 
**FailureCondition** | Pointer to **string** | The task&#39;s failure condition, as rendered | [optional] 
**Kind** | **string** |  | 
**Live** | Pointer to [**WorkflowTaskObjectLive**](WorkflowTaskObjectLive.md) |  | [optional] 
**ManifestRecorded** | Pointer to **bool** | False for executions that ran before rendered manifests were recorded | [optional] 
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to **string** | succeeded or failed | [optional] 
**RecordedAt** | Pointer to **string** | When the task finished, in RFC3339 format | [optional] 
**RecordedUid** | Pointer to **string** | UID when the task finished | [optional] 
**RedactedValues** | Pointer to **int64** | How many values were redacted | [optional] 
**RedactionReasons** | Pointer to **[]string** | Why values were redacted: secret-input, secret-data or sensitive-key | [optional] 
**RenderedManifest** | Pointer to **string** | The manifest as rendered when the task ran, with every value that came from a secret input, every Secret data value and every value under a sensitive key redacted before storage | [optional] 
**SuccessCondition** | Pointer to **string** | The task&#39;s success condition, as rendered | [optional] 

## Methods

### NewWorkflowTaskObject

`func NewWorkflowTaskObject(action string, kind string, name string, ) *WorkflowTaskObject`

NewWorkflowTaskObject instantiates a new WorkflowTaskObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskObjectWithDefaults

`func NewWorkflowTaskObjectWithDefaults() *WorkflowTaskObject`

NewWorkflowTaskObjectWithDefaults instantiates a new WorkflowTaskObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WorkflowTaskObject) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowTaskObject) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowTaskObject) SetAction(v string)`

SetAction sets Action field to given value.


### GetApiVersion

`func (o *WorkflowTaskObject) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WorkflowTaskObject) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WorkflowTaskObject) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *WorkflowTaskObject) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetConditionSummary

`func (o *WorkflowTaskObject) GetConditionSummary() string`

GetConditionSummary returns the ConditionSummary field if non-nil, zero value otherwise.

### GetConditionSummaryOk

`func (o *WorkflowTaskObject) GetConditionSummaryOk() (*string, bool)`

GetConditionSummaryOk returns a tuple with the ConditionSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionSummary

`func (o *WorkflowTaskObject) SetConditionSummary(v string)`

SetConditionSummary sets ConditionSummary field to given value.

### HasConditionSummary

`func (o *WorkflowTaskObject) HasConditionSummary() bool`

HasConditionSummary returns a boolean if a field has been set.

### GetCustomResource

`func (o *WorkflowTaskObject) GetCustomResource() bool`

GetCustomResource returns the CustomResource field if non-nil, zero value otherwise.

### GetCustomResourceOk

`func (o *WorkflowTaskObject) GetCustomResourceOk() (*bool, bool)`

GetCustomResourceOk returns a tuple with the CustomResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResource

`func (o *WorkflowTaskObject) SetCustomResource(v bool)`

SetCustomResource sets CustomResource field to given value.

### HasCustomResource

`func (o *WorkflowTaskObject) HasCustomResource() bool`

HasCustomResource returns a boolean if a field has been set.

### GetFailureCondition

`func (o *WorkflowTaskObject) GetFailureCondition() string`

GetFailureCondition returns the FailureCondition field if non-nil, zero value otherwise.

### GetFailureConditionOk

`func (o *WorkflowTaskObject) GetFailureConditionOk() (*string, bool)`

GetFailureConditionOk returns a tuple with the FailureCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCondition

`func (o *WorkflowTaskObject) SetFailureCondition(v string)`

SetFailureCondition sets FailureCondition field to given value.

### HasFailureCondition

`func (o *WorkflowTaskObject) HasFailureCondition() bool`

HasFailureCondition returns a boolean if a field has been set.

### GetKind

`func (o *WorkflowTaskObject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkflowTaskObject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkflowTaskObject) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLive

`func (o *WorkflowTaskObject) GetLive() WorkflowTaskObjectLive`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *WorkflowTaskObject) GetLiveOk() (*WorkflowTaskObjectLive, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *WorkflowTaskObject) SetLive(v WorkflowTaskObjectLive)`

SetLive sets Live field to given value.

### HasLive

`func (o *WorkflowTaskObject) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetManifestRecorded

`func (o *WorkflowTaskObject) GetManifestRecorded() bool`

GetManifestRecorded returns the ManifestRecorded field if non-nil, zero value otherwise.

### GetManifestRecordedOk

`func (o *WorkflowTaskObject) GetManifestRecordedOk() (*bool, bool)`

GetManifestRecordedOk returns a tuple with the ManifestRecorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestRecorded

`func (o *WorkflowTaskObject) SetManifestRecorded(v bool)`

SetManifestRecorded sets ManifestRecorded field to given value.

### HasManifestRecorded

`func (o *WorkflowTaskObject) HasManifestRecorded() bool`

HasManifestRecorded returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskObject) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *WorkflowTaskObject) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *WorkflowTaskObject) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *WorkflowTaskObject) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *WorkflowTaskObject) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOutcome

`func (o *WorkflowTaskObject) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *WorkflowTaskObject) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *WorkflowTaskObject) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *WorkflowTaskObject) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetRecordedAt

`func (o *WorkflowTaskObject) GetRecordedAt() string`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *WorkflowTaskObject) GetRecordedAtOk() (*string, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *WorkflowTaskObject) SetRecordedAt(v string)`

SetRecordedAt sets RecordedAt field to given value.

### HasRecordedAt

`func (o *WorkflowTaskObject) HasRecordedAt() bool`

HasRecordedAt returns a boolean if a field has been set.

### GetRecordedUid

`func (o *WorkflowTaskObject) GetRecordedUid() string`

GetRecordedUid returns the RecordedUid field if non-nil, zero value otherwise.

### GetRecordedUidOk

`func (o *WorkflowTaskObject) GetRecordedUidOk() (*string, bool)`

GetRecordedUidOk returns a tuple with the RecordedUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedUid

`func (o *WorkflowTaskObject) SetRecordedUid(v string)`

SetRecordedUid sets RecordedUid field to given value.

### HasRecordedUid

`func (o *WorkflowTaskObject) HasRecordedUid() bool`

HasRecordedUid returns a boolean if a field has been set.

### GetRedactedValues

`func (o *WorkflowTaskObject) GetRedactedValues() int64`

GetRedactedValues returns the RedactedValues field if non-nil, zero value otherwise.

### GetRedactedValuesOk

`func (o *WorkflowTaskObject) GetRedactedValuesOk() (*int64, bool)`

GetRedactedValuesOk returns a tuple with the RedactedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactedValues

`func (o *WorkflowTaskObject) SetRedactedValues(v int64)`

SetRedactedValues sets RedactedValues field to given value.

### HasRedactedValues

`func (o *WorkflowTaskObject) HasRedactedValues() bool`

HasRedactedValues returns a boolean if a field has been set.

### GetRedactionReasons

`func (o *WorkflowTaskObject) GetRedactionReasons() []string`

GetRedactionReasons returns the RedactionReasons field if non-nil, zero value otherwise.

### GetRedactionReasonsOk

`func (o *WorkflowTaskObject) GetRedactionReasonsOk() (*[]string, bool)`

GetRedactionReasonsOk returns a tuple with the RedactionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactionReasons

`func (o *WorkflowTaskObject) SetRedactionReasons(v []string)`

SetRedactionReasons sets RedactionReasons field to given value.

### HasRedactionReasons

`func (o *WorkflowTaskObject) HasRedactionReasons() bool`

HasRedactionReasons returns a boolean if a field has been set.

### GetRenderedManifest

`func (o *WorkflowTaskObject) GetRenderedManifest() string`

GetRenderedManifest returns the RenderedManifest field if non-nil, zero value otherwise.

### GetRenderedManifestOk

`func (o *WorkflowTaskObject) GetRenderedManifestOk() (*string, bool)`

GetRenderedManifestOk returns a tuple with the RenderedManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedManifest

`func (o *WorkflowTaskObject) SetRenderedManifest(v string)`

SetRenderedManifest sets RenderedManifest field to given value.

### HasRenderedManifest

`func (o *WorkflowTaskObject) HasRenderedManifest() bool`

HasRenderedManifest returns a boolean if a field has been set.

### GetSuccessCondition

`func (o *WorkflowTaskObject) GetSuccessCondition() string`

GetSuccessCondition returns the SuccessCondition field if non-nil, zero value otherwise.

### GetSuccessConditionOk

`func (o *WorkflowTaskObject) GetSuccessConditionOk() (*string, bool)`

GetSuccessConditionOk returns a tuple with the SuccessCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCondition

`func (o *WorkflowTaskObject) SetSuccessCondition(v string)`

SetSuccessCondition sets SuccessCondition field to given value.

### HasSuccessCondition

`func (o *WorkflowTaskObject) HasSuccessCondition() bool`

HasSuccessCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


