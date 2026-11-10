# BatchRevokeAccessRequestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | Explicit request IDs. Backend deduplicates and authorizes each ID independently; partial failures do not roll back successful revocations. | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewBatchRevokeAccessRequestsRequest

`func NewBatchRevokeAccessRequestsRequest(ids []string, token string, ) *BatchRevokeAccessRequestsRequest`

NewBatchRevokeAccessRequestsRequest instantiates a new BatchRevokeAccessRequestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRevokeAccessRequestsRequestWithDefaults

`func NewBatchRevokeAccessRequestsRequestWithDefaults() *BatchRevokeAccessRequestsRequest`

NewBatchRevokeAccessRequestsRequestWithDefaults instantiates a new BatchRevokeAccessRequestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BatchRevokeAccessRequestsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BatchRevokeAccessRequestsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BatchRevokeAccessRequestsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetToken

`func (o *BatchRevokeAccessRequestsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BatchRevokeAccessRequestsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BatchRevokeAccessRequestsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


