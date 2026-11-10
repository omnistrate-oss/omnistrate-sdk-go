# BatchRevokeAccessRequestsRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | Explicit request IDs. Backend deduplicates and authorizes each ID independently; partial failures do not roll back successful revocations. | 

## Methods

### NewBatchRevokeAccessRequestsRequest2

`func NewBatchRevokeAccessRequestsRequest2(ids []string, ) *BatchRevokeAccessRequestsRequest2`

NewBatchRevokeAccessRequestsRequest2 instantiates a new BatchRevokeAccessRequestsRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRevokeAccessRequestsRequest2WithDefaults

`func NewBatchRevokeAccessRequestsRequest2WithDefaults() *BatchRevokeAccessRequestsRequest2`

NewBatchRevokeAccessRequestsRequest2WithDefaults instantiates a new BatchRevokeAccessRequestsRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BatchRevokeAccessRequestsRequest2) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BatchRevokeAccessRequestsRequest2) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BatchRevokeAccessRequestsRequest2) SetIds(v []string)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


