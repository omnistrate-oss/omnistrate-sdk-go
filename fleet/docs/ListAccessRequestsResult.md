# ListAccessRequestsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**[]AccessRequest**](AccessRequest.md) | The access requests returned in descending creation order. | [optional] 

## Methods

### NewListAccessRequestsResult

`func NewListAccessRequestsResult() *ListAccessRequestsResult`

NewListAccessRequestsResult instantiates a new ListAccessRequestsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccessRequestsResultWithDefaults

`func NewListAccessRequestsResultWithDefaults() *ListAccessRequestsResult`

NewListAccessRequestsResultWithDefaults instantiates a new ListAccessRequestsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *ListAccessRequestsResult) GetRequests() []AccessRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ListAccessRequestsResult) GetRequestsOk() (*[]AccessRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ListAccessRequestsResult) SetRequests(v []AccessRequest)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ListAccessRequestsResult) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


