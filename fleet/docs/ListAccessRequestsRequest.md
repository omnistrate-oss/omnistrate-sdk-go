# ListAccessRequestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeDeleted** | Pointer to **bool** | Include all retained soft-deleted requests in the organization&#39;s authorized environment scope instead of only terminal requests from the previous 24 hours. Denied, revoked, and expired requests are soft-deleted while retaining their lifecycle and approval facts. | [optional] 
**Limit** | Pointer to **int64** | Maximum number of requests returned when includeDeleted is true. Omission defaults to 1000. This field is invalid when includeDeleted is false or omitted. | [optional] 
**Statuses** | Pointer to **[]string** | Filter by any of the supplied lifecycle statuses, such as PENDING and APPROVED. Omission or an empty list returns all statuses. Supported values are validated by the backend. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListAccessRequestsRequest

`func NewListAccessRequestsRequest(token string, ) *ListAccessRequestsRequest`

NewListAccessRequestsRequest instantiates a new ListAccessRequestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccessRequestsRequestWithDefaults

`func NewListAccessRequestsRequestWithDefaults() *ListAccessRequestsRequest`

NewListAccessRequestsRequestWithDefaults instantiates a new ListAccessRequestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeDeleted

`func (o *ListAccessRequestsRequest) GetIncludeDeleted() bool`

GetIncludeDeleted returns the IncludeDeleted field if non-nil, zero value otherwise.

### GetIncludeDeletedOk

`func (o *ListAccessRequestsRequest) GetIncludeDeletedOk() (*bool, bool)`

GetIncludeDeletedOk returns a tuple with the IncludeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDeleted

`func (o *ListAccessRequestsRequest) SetIncludeDeleted(v bool)`

SetIncludeDeleted sets IncludeDeleted field to given value.

### HasIncludeDeleted

`func (o *ListAccessRequestsRequest) HasIncludeDeleted() bool`

HasIncludeDeleted returns a boolean if a field has been set.

### GetLimit

`func (o *ListAccessRequestsRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListAccessRequestsRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListAccessRequestsRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListAccessRequestsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetStatuses

`func (o *ListAccessRequestsRequest) GetStatuses() []string`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ListAccessRequestsRequest) GetStatusesOk() (*[]string, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ListAccessRequestsRequest) SetStatuses(v []string)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ListAccessRequestsRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetToken

`func (o *ListAccessRequestsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAccessRequestsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAccessRequestsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


