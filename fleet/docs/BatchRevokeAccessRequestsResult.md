# BatchRevokeAccessRequestsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]BatchRevokeAccessRequestResult**](BatchRevokeAccessRequestResult.md) | One result per distinct ID with its lifecycle status or a safe per-item error. | [optional] 

## Methods

### NewBatchRevokeAccessRequestsResult

`func NewBatchRevokeAccessRequestsResult() *BatchRevokeAccessRequestsResult`

NewBatchRevokeAccessRequestsResult instantiates a new BatchRevokeAccessRequestsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRevokeAccessRequestsResultWithDefaults

`func NewBatchRevokeAccessRequestsResultWithDefaults() *BatchRevokeAccessRequestsResult`

NewBatchRevokeAccessRequestsResultWithDefaults instantiates a new BatchRevokeAccessRequestsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BatchRevokeAccessRequestsResult) GetResults() []BatchRevokeAccessRequestResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchRevokeAccessRequestsResult) GetResultsOk() (*[]BatchRevokeAccessRequestResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchRevokeAccessRequestsResult) SetResults(v []BatchRevokeAccessRequestResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *BatchRevokeAccessRequestsResult) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


