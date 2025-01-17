# ListSubscriptionRequestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the subscription request | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListSubscriptionRequestsRequest

`func NewListSubscriptionRequestsRequest(token string, ) *ListSubscriptionRequestsRequest`

NewListSubscriptionRequestsRequest instantiates a new ListSubscriptionRequestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptionRequestsRequestWithDefaults

`func NewListSubscriptionRequestsRequestWithDefaults() *ListSubscriptionRequestsRequest`

NewListSubscriptionRequestsRequestWithDefaults instantiates a new ListSubscriptionRequestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListSubscriptionRequestsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSubscriptionRequestsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSubscriptionRequestsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSubscriptionRequestsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ListSubscriptionRequestsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListSubscriptionRequestsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListSubscriptionRequestsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


