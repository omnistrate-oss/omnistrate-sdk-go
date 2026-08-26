# RetryMarketplaceFulfillmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Omnistrate contract identifier | 
**Reason** | Pointer to **string** | Recorded on the audit event. Worth filling in, because a retry is a manual intervention on a contract that is usually already billing | [optional] 
**Stage** | Pointer to **string** | One stage of the fulfillment flow, in the order the workflow runs them. Creating the subscription request and committing the contract are separate stages because they are separate calls with separate failure modes, and an operator debugging a stuck contract has to know which did not happen. CONTRACT_HELD precedes SUBSCRIPTION_REQUEST_CREATED because the hold is taken before the handoff: holding is what stops the buyer being charged while the vendor onboards them, and channels that cannot hold a contract skip it | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRetryMarketplaceFulfillmentRequest

`func NewRetryMarketplaceFulfillmentRequest(id string, token string, ) *RetryMarketplaceFulfillmentRequest`

NewRetryMarketplaceFulfillmentRequest instantiates a new RetryMarketplaceFulfillmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryMarketplaceFulfillmentRequestWithDefaults

`func NewRetryMarketplaceFulfillmentRequestWithDefaults() *RetryMarketplaceFulfillmentRequest`

NewRetryMarketplaceFulfillmentRequestWithDefaults instantiates a new RetryMarketplaceFulfillmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RetryMarketplaceFulfillmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RetryMarketplaceFulfillmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RetryMarketplaceFulfillmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *RetryMarketplaceFulfillmentRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RetryMarketplaceFulfillmentRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RetryMarketplaceFulfillmentRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RetryMarketplaceFulfillmentRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStage

`func (o *RetryMarketplaceFulfillmentRequest) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *RetryMarketplaceFulfillmentRequest) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *RetryMarketplaceFulfillmentRequest) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *RetryMarketplaceFulfillmentRequest) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetToken

`func (o *RetryMarketplaceFulfillmentRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RetryMarketplaceFulfillmentRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RetryMarketplaceFulfillmentRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


