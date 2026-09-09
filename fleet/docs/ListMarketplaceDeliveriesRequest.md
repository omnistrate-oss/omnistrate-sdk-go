# ListMarketplaceDeliveriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | [optional] 
**Direction** | Pointer to **string** | Who called whom. OUTBOUND is a webhook Omnistrate sent to the ISV&#39;s receiver and INBOUND is a call the ISV made to the marketplace API, both relative to the ISV. CHANNEL_INBOUND is the marketplace calling Omnistrate, and CHANNEL_OUTBOUND is Omnistrate calling the marketplace | [optional] 
**EventType** | Pointer to **string** | The type of a marketplace fulfillment event delivered to an ISV receiver | [optional] 
**FailuresOnly** | Pointer to **bool** | A shortcut for the only filter combination anybody types twice | [optional] 
**MarketplaceContractId** | Pointer to **string** | Filter to one contract, which is how a contract detail view scopes its own trail | [optional] 
**Since** | Pointer to **time.Time** | Deliveries sent or received at or after this instant | [optional] 
**Status** | Pointer to **string** | PENDING has not been answered yet. DELIVERED was accepted. FAILED exhausted its retries. BLOCKED was never sent, because the receiver address was refused | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListMarketplaceDeliveriesRequest

`func NewListMarketplaceDeliveriesRequest(token string, ) *ListMarketplaceDeliveriesRequest`

NewListMarketplaceDeliveriesRequest instantiates a new ListMarketplaceDeliveriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMarketplaceDeliveriesRequestWithDefaults

`func NewListMarketplaceDeliveriesRequestWithDefaults() *ListMarketplaceDeliveriesRequest`

NewListMarketplaceDeliveriesRequestWithDefaults instantiates a new ListMarketplaceDeliveriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ListMarketplaceDeliveriesRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ListMarketplaceDeliveriesRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ListMarketplaceDeliveriesRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ListMarketplaceDeliveriesRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDirection

`func (o *ListMarketplaceDeliveriesRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ListMarketplaceDeliveriesRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ListMarketplaceDeliveriesRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ListMarketplaceDeliveriesRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEventType

`func (o *ListMarketplaceDeliveriesRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ListMarketplaceDeliveriesRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ListMarketplaceDeliveriesRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ListMarketplaceDeliveriesRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetFailuresOnly

`func (o *ListMarketplaceDeliveriesRequest) GetFailuresOnly() bool`

GetFailuresOnly returns the FailuresOnly field if non-nil, zero value otherwise.

### GetFailuresOnlyOk

`func (o *ListMarketplaceDeliveriesRequest) GetFailuresOnlyOk() (*bool, bool)`

GetFailuresOnlyOk returns a tuple with the FailuresOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailuresOnly

`func (o *ListMarketplaceDeliveriesRequest) SetFailuresOnly(v bool)`

SetFailuresOnly sets FailuresOnly field to given value.

### HasFailuresOnly

`func (o *ListMarketplaceDeliveriesRequest) HasFailuresOnly() bool`

HasFailuresOnly returns a boolean if a field has been set.

### GetMarketplaceContractId

`func (o *ListMarketplaceDeliveriesRequest) GetMarketplaceContractId() string`

GetMarketplaceContractId returns the MarketplaceContractId field if non-nil, zero value otherwise.

### GetMarketplaceContractIdOk

`func (o *ListMarketplaceDeliveriesRequest) GetMarketplaceContractIdOk() (*string, bool)`

GetMarketplaceContractIdOk returns a tuple with the MarketplaceContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceContractId

`func (o *ListMarketplaceDeliveriesRequest) SetMarketplaceContractId(v string)`

SetMarketplaceContractId sets MarketplaceContractId field to given value.

### HasMarketplaceContractId

`func (o *ListMarketplaceDeliveriesRequest) HasMarketplaceContractId() bool`

HasMarketplaceContractId returns a boolean if a field has been set.

### GetSince

`func (o *ListMarketplaceDeliveriesRequest) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *ListMarketplaceDeliveriesRequest) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *ListMarketplaceDeliveriesRequest) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *ListMarketplaceDeliveriesRequest) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetStatus

`func (o *ListMarketplaceDeliveriesRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListMarketplaceDeliveriesRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListMarketplaceDeliveriesRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListMarketplaceDeliveriesRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ListMarketplaceDeliveriesRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListMarketplaceDeliveriesRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListMarketplaceDeliveriesRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


