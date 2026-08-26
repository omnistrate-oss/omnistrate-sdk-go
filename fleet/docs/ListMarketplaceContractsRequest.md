# ListMarketplaceContractsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | [optional] 
**FulfillmentState** | Pointer to **string** | What Omnistrate decided. Deployments and metering are allowed if and only if this is READY, plus exactly one final metering window during DEPROVISIONING | [optional] 
**IncludeSimulated** | Pointer to **bool** | Sandbox contracts are excluded unless this is set, so a revenue view cannot pick them up by forgetting a filter | [optional] 
**NeedsAttention** | Pointer to **bool** | Only contracts past the handoff SLA while billable. The operator&#39;s working queue: every one of these is a customer paying for something they cannot use | [optional] 
**Quadrant** | Pointer to **string** | Where this contract sits in the reconciliation 2x2. PAYING_NOT_PROVISIONED is the one that costs money: the buyer is being billed and cannot use the product. ONBOARDING is on neither axis: the channel is HOLDING the contract, so the buyer is deliberately not being charged and is not yet being served, which is the normal path on a channel that can hold. It is called out because the four quadrants have nowhere to put it, and filing a held contract under CLOSED said an onboarding buyer had ended | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListMarketplaceContractsRequest

`func NewListMarketplaceContractsRequest(token string, ) *ListMarketplaceContractsRequest`

NewListMarketplaceContractsRequest instantiates a new ListMarketplaceContractsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMarketplaceContractsRequestWithDefaults

`func NewListMarketplaceContractsRequestWithDefaults() *ListMarketplaceContractsRequest`

NewListMarketplaceContractsRequestWithDefaults instantiates a new ListMarketplaceContractsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ListMarketplaceContractsRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ListMarketplaceContractsRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ListMarketplaceContractsRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ListMarketplaceContractsRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetFulfillmentState

`func (o *ListMarketplaceContractsRequest) GetFulfillmentState() string`

GetFulfillmentState returns the FulfillmentState field if non-nil, zero value otherwise.

### GetFulfillmentStateOk

`func (o *ListMarketplaceContractsRequest) GetFulfillmentStateOk() (*string, bool)`

GetFulfillmentStateOk returns a tuple with the FulfillmentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentState

`func (o *ListMarketplaceContractsRequest) SetFulfillmentState(v string)`

SetFulfillmentState sets FulfillmentState field to given value.

### HasFulfillmentState

`func (o *ListMarketplaceContractsRequest) HasFulfillmentState() bool`

HasFulfillmentState returns a boolean if a field has been set.

### GetIncludeSimulated

`func (o *ListMarketplaceContractsRequest) GetIncludeSimulated() bool`

GetIncludeSimulated returns the IncludeSimulated field if non-nil, zero value otherwise.

### GetIncludeSimulatedOk

`func (o *ListMarketplaceContractsRequest) GetIncludeSimulatedOk() (*bool, bool)`

GetIncludeSimulatedOk returns a tuple with the IncludeSimulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSimulated

`func (o *ListMarketplaceContractsRequest) SetIncludeSimulated(v bool)`

SetIncludeSimulated sets IncludeSimulated field to given value.

### HasIncludeSimulated

`func (o *ListMarketplaceContractsRequest) HasIncludeSimulated() bool`

HasIncludeSimulated returns a boolean if a field has been set.

### GetNeedsAttention

`func (o *ListMarketplaceContractsRequest) GetNeedsAttention() bool`

GetNeedsAttention returns the NeedsAttention field if non-nil, zero value otherwise.

### GetNeedsAttentionOk

`func (o *ListMarketplaceContractsRequest) GetNeedsAttentionOk() (*bool, bool)`

GetNeedsAttentionOk returns a tuple with the NeedsAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsAttention

`func (o *ListMarketplaceContractsRequest) SetNeedsAttention(v bool)`

SetNeedsAttention sets NeedsAttention field to given value.

### HasNeedsAttention

`func (o *ListMarketplaceContractsRequest) HasNeedsAttention() bool`

HasNeedsAttention returns a boolean if a field has been set.

### GetQuadrant

`func (o *ListMarketplaceContractsRequest) GetQuadrant() string`

GetQuadrant returns the Quadrant field if non-nil, zero value otherwise.

### GetQuadrantOk

`func (o *ListMarketplaceContractsRequest) GetQuadrantOk() (*string, bool)`

GetQuadrantOk returns a tuple with the Quadrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuadrant

`func (o *ListMarketplaceContractsRequest) SetQuadrant(v string)`

SetQuadrant sets Quadrant field to given value.

### HasQuadrant

`func (o *ListMarketplaceContractsRequest) HasQuadrant() bool`

HasQuadrant returns a boolean if a field has been set.

### GetToken

`func (o *ListMarketplaceContractsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListMarketplaceContractsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListMarketplaceContractsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


