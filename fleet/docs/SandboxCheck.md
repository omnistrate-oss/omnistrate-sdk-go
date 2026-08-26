# SandboxCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckId** | **string** | Which conformance check a result belongs to | 
**Description** | Pointer to **string** | What the check sends and why it matters | [optional] 
**ExpectedStatus** | **string** | The status the receiver is expected to return, as a family or an exact code | 
**LatencyMs** | Pointer to **int64** |  | [optional] 
**ObservedStatus** | Pointer to **int64** | What the receiver actually returned. Absent when there was no response at all, which is distinct from a 5xx and is reported through transportError instead | [optional] 
**Passed** | **bool** |  | 
**Remediation** | Pointer to **string** | What to change when the check failed. Empty when it passed | [optional] 
**Title** | **string** | Human readable name of the check | 
**TransportError** | Pointer to **string** | Set when no response was received: timeout, connection refused, TLS failure | [optional] 

## Methods

### NewSandboxCheck

`func NewSandboxCheck(checkId string, expectedStatus string, passed bool, title string, ) *SandboxCheck`

NewSandboxCheck instantiates a new SandboxCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCheckWithDefaults

`func NewSandboxCheckWithDefaults() *SandboxCheck`

NewSandboxCheckWithDefaults instantiates a new SandboxCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckId

`func (o *SandboxCheck) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *SandboxCheck) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *SandboxCheck) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.


### GetDescription

`func (o *SandboxCheck) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SandboxCheck) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SandboxCheck) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SandboxCheck) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpectedStatus

`func (o *SandboxCheck) GetExpectedStatus() string`

GetExpectedStatus returns the ExpectedStatus field if non-nil, zero value otherwise.

### GetExpectedStatusOk

`func (o *SandboxCheck) GetExpectedStatusOk() (*string, bool)`

GetExpectedStatusOk returns a tuple with the ExpectedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStatus

`func (o *SandboxCheck) SetExpectedStatus(v string)`

SetExpectedStatus sets ExpectedStatus field to given value.


### GetLatencyMs

`func (o *SandboxCheck) GetLatencyMs() int64`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *SandboxCheck) GetLatencyMsOk() (*int64, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *SandboxCheck) SetLatencyMs(v int64)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *SandboxCheck) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetObservedStatus

`func (o *SandboxCheck) GetObservedStatus() int64`

GetObservedStatus returns the ObservedStatus field if non-nil, zero value otherwise.

### GetObservedStatusOk

`func (o *SandboxCheck) GetObservedStatusOk() (*int64, bool)`

GetObservedStatusOk returns a tuple with the ObservedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedStatus

`func (o *SandboxCheck) SetObservedStatus(v int64)`

SetObservedStatus sets ObservedStatus field to given value.

### HasObservedStatus

`func (o *SandboxCheck) HasObservedStatus() bool`

HasObservedStatus returns a boolean if a field has been set.

### GetPassed

`func (o *SandboxCheck) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *SandboxCheck) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *SandboxCheck) SetPassed(v bool)`

SetPassed sets Passed field to given value.


### GetRemediation

`func (o *SandboxCheck) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *SandboxCheck) GetRemediationOk() (*string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *SandboxCheck) SetRemediation(v string)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *SandboxCheck) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### GetTitle

`func (o *SandboxCheck) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SandboxCheck) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SandboxCheck) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTransportError

`func (o *SandboxCheck) GetTransportError() string`

GetTransportError returns the TransportError field if non-nil, zero value otherwise.

### GetTransportErrorOk

`func (o *SandboxCheck) GetTransportErrorOk() (*string, bool)`

GetTransportErrorOk returns a tuple with the TransportError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportError

`func (o *SandboxCheck) SetTransportError(v string)`

SetTransportError sets TransportError field to given value.

### HasTransportError

`func (o *SandboxCheck) HasTransportError() bool`

HasTransportError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


