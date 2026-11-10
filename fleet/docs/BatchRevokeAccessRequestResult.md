# BatchRevokeAccessRequestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Safe per-item error without leaking hidden request metadata or credentials. Omitted on success. | [optional] 
**Id** | Pointer to **string** | Server-generated ID of a provider cluster access request. | [optional] 
**Status** | Pointer to **string** | Access request lifecycle status. Known values: PENDING, APPROVING, APPROVED, REVOKING, DENIED, REVOKED, EXPIRED. APPROVED covers both manual and automatic approval. Clients must preserve unknown future values. | [optional] 

## Methods

### NewBatchRevokeAccessRequestResult

`func NewBatchRevokeAccessRequestResult() *BatchRevokeAccessRequestResult`

NewBatchRevokeAccessRequestResult instantiates a new BatchRevokeAccessRequestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRevokeAccessRequestResultWithDefaults

`func NewBatchRevokeAccessRequestResultWithDefaults() *BatchRevokeAccessRequestResult`

NewBatchRevokeAccessRequestResultWithDefaults instantiates a new BatchRevokeAccessRequestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *BatchRevokeAccessRequestResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BatchRevokeAccessRequestResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BatchRevokeAccessRequestResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BatchRevokeAccessRequestResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *BatchRevokeAccessRequestResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchRevokeAccessRequestResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchRevokeAccessRequestResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BatchRevokeAccessRequestResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *BatchRevokeAccessRequestResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchRevokeAccessRequestResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchRevokeAccessRequestResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchRevokeAccessRequestResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


