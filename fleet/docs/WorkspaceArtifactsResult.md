# WorkspaceArtifactsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastFailedReason** | Pointer to **string** | The last failed reason for preparing the workspace artifact archive | [optional] 
**Status** | **string** | The status of an operation | 
**WorkspaceArtifactId** | **string** | ID of a workspace artifact archive | 

## Methods

### NewWorkspaceArtifactsResult

`func NewWorkspaceArtifactsResult(status string, workspaceArtifactId string, ) *WorkspaceArtifactsResult`

NewWorkspaceArtifactsResult instantiates a new WorkspaceArtifactsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceArtifactsResultWithDefaults

`func NewWorkspaceArtifactsResultWithDefaults() *WorkspaceArtifactsResult`

NewWorkspaceArtifactsResultWithDefaults instantiates a new WorkspaceArtifactsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastFailedReason

`func (o *WorkspaceArtifactsResult) GetLastFailedReason() string`

GetLastFailedReason returns the LastFailedReason field if non-nil, zero value otherwise.

### GetLastFailedReasonOk

`func (o *WorkspaceArtifactsResult) GetLastFailedReasonOk() (*string, bool)`

GetLastFailedReasonOk returns a tuple with the LastFailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailedReason

`func (o *WorkspaceArtifactsResult) SetLastFailedReason(v string)`

SetLastFailedReason sets LastFailedReason field to given value.

### HasLastFailedReason

`func (o *WorkspaceArtifactsResult) HasLastFailedReason() bool`

HasLastFailedReason returns a boolean if a field has been set.

### GetStatus

`func (o *WorkspaceArtifactsResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkspaceArtifactsResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkspaceArtifactsResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWorkspaceArtifactId

`func (o *WorkspaceArtifactsResult) GetWorkspaceArtifactId() string`

GetWorkspaceArtifactId returns the WorkspaceArtifactId field if non-nil, zero value otherwise.

### GetWorkspaceArtifactIdOk

`func (o *WorkspaceArtifactsResult) GetWorkspaceArtifactIdOk() (*string, bool)`

GetWorkspaceArtifactIdOk returns a tuple with the WorkspaceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceArtifactId

`func (o *WorkspaceArtifactsResult) SetWorkspaceArtifactId(v string)`

SetWorkspaceArtifactId sets WorkspaceArtifactId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


