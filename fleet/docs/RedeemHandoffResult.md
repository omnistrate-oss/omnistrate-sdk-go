# RedeemHandoffResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerRef** | Pointer to **string** | The channel&#39;s stable identifier for the BUYER, which is not the entitlement. A buyer accumulates entitlements across renewals and upsells, so this is the value to key a tenant on | [optional] 
**Capabilities** | Pointer to [**MarketplaceEventCapabilities**](MarketplaceEventCapabilities.md) |  | [optional] 
**Channel** | **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | 
**ContractStatus** | **string** | What the marketplace says about the contract. Mirrored, never authoritative for whether the buyer may be served | 
**ContractVersion** | **int64** | Monotonic per contract. The same value carried by webhook bodies, so a redeem and a delivery can be ordered against each other | 
**ExternalRef** | **string** | The channel&#39;s own identifier for this entitlement. Unique per channel and the key both sides can correlate on | 
**FulfillmentState** | **string** | What Omnistrate decided. Deployments and metering are allowed if and only if this is READY, plus exactly one final metering window during DEPROVISIONING | 
**HandoffSlaExpiresAt** | Pointer to **time.Time** | When the handoff SLA is breached and the contract is reported as orphaned. Distinct from the token expiry, which is deliberately later so a breached SLA can still be recovered from without an operator | [optional] 
**HandoffTokenExpiresAt** | Pointer to **time.Time** | When this token stops being redeemable. After it, this endpoint answers 410 rather than 404, so an expired token can be told apart from one that never existed | [optional] 
**MarketplaceContractId** | **string** |  | 
**Org** | Pointer to [**MarketplaceEventOrg**](MarketplaceEventOrg.md) |  | [optional] 
**Plan** | Pointer to [**MarketplaceEventPlan**](MarketplaceEventPlan.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | The Omnistrate subscription, once the confirm has created it. Absent while fulfillment is still awaiting the confirm, which is how a redeem tells you whether your own confirm has already landed | [optional] 
**SubscriptionRequest** | Pointer to [**MarketplaceEventSubscriptionRequest**](MarketplaceEventSubscriptionRequest.md) |  | [optional] 

## Methods

### NewRedeemHandoffResult

`func NewRedeemHandoffResult(channel string, contractStatus string, contractVersion int64, externalRef string, fulfillmentState string, marketplaceContractId string, ) *RedeemHandoffResult`

NewRedeemHandoffResult instantiates a new RedeemHandoffResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemHandoffResultWithDefaults

`func NewRedeemHandoffResultWithDefaults() *RedeemHandoffResult`

NewRedeemHandoffResultWithDefaults instantiates a new RedeemHandoffResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerRef

`func (o *RedeemHandoffResult) GetBuyerRef() string`

GetBuyerRef returns the BuyerRef field if non-nil, zero value otherwise.

### GetBuyerRefOk

`func (o *RedeemHandoffResult) GetBuyerRefOk() (*string, bool)`

GetBuyerRefOk returns a tuple with the BuyerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerRef

`func (o *RedeemHandoffResult) SetBuyerRef(v string)`

SetBuyerRef sets BuyerRef field to given value.

### HasBuyerRef

`func (o *RedeemHandoffResult) HasBuyerRef() bool`

HasBuyerRef returns a boolean if a field has been set.

### GetCapabilities

`func (o *RedeemHandoffResult) GetCapabilities() MarketplaceEventCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *RedeemHandoffResult) GetCapabilitiesOk() (*MarketplaceEventCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *RedeemHandoffResult) SetCapabilities(v MarketplaceEventCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *RedeemHandoffResult) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetChannel

`func (o *RedeemHandoffResult) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *RedeemHandoffResult) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *RedeemHandoffResult) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetContractStatus

`func (o *RedeemHandoffResult) GetContractStatus() string`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *RedeemHandoffResult) GetContractStatusOk() (*string, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *RedeemHandoffResult) SetContractStatus(v string)`

SetContractStatus sets ContractStatus field to given value.


### GetContractVersion

`func (o *RedeemHandoffResult) GetContractVersion() int64`

GetContractVersion returns the ContractVersion field if non-nil, zero value otherwise.

### GetContractVersionOk

`func (o *RedeemHandoffResult) GetContractVersionOk() (*int64, bool)`

GetContractVersionOk returns a tuple with the ContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractVersion

`func (o *RedeemHandoffResult) SetContractVersion(v int64)`

SetContractVersion sets ContractVersion field to given value.


### GetExternalRef

`func (o *RedeemHandoffResult) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *RedeemHandoffResult) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *RedeemHandoffResult) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.


### GetFulfillmentState

`func (o *RedeemHandoffResult) GetFulfillmentState() string`

GetFulfillmentState returns the FulfillmentState field if non-nil, zero value otherwise.

### GetFulfillmentStateOk

`func (o *RedeemHandoffResult) GetFulfillmentStateOk() (*string, bool)`

GetFulfillmentStateOk returns a tuple with the FulfillmentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentState

`func (o *RedeemHandoffResult) SetFulfillmentState(v string)`

SetFulfillmentState sets FulfillmentState field to given value.


### GetHandoffSlaExpiresAt

`func (o *RedeemHandoffResult) GetHandoffSlaExpiresAt() time.Time`

GetHandoffSlaExpiresAt returns the HandoffSlaExpiresAt field if non-nil, zero value otherwise.

### GetHandoffSlaExpiresAtOk

`func (o *RedeemHandoffResult) GetHandoffSlaExpiresAtOk() (*time.Time, bool)`

GetHandoffSlaExpiresAtOk returns a tuple with the HandoffSlaExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffSlaExpiresAt

`func (o *RedeemHandoffResult) SetHandoffSlaExpiresAt(v time.Time)`

SetHandoffSlaExpiresAt sets HandoffSlaExpiresAt field to given value.

### HasHandoffSlaExpiresAt

`func (o *RedeemHandoffResult) HasHandoffSlaExpiresAt() bool`

HasHandoffSlaExpiresAt returns a boolean if a field has been set.

### GetHandoffTokenExpiresAt

`func (o *RedeemHandoffResult) GetHandoffTokenExpiresAt() time.Time`

GetHandoffTokenExpiresAt returns the HandoffTokenExpiresAt field if non-nil, zero value otherwise.

### GetHandoffTokenExpiresAtOk

`func (o *RedeemHandoffResult) GetHandoffTokenExpiresAtOk() (*time.Time, bool)`

GetHandoffTokenExpiresAtOk returns a tuple with the HandoffTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffTokenExpiresAt

`func (o *RedeemHandoffResult) SetHandoffTokenExpiresAt(v time.Time)`

SetHandoffTokenExpiresAt sets HandoffTokenExpiresAt field to given value.

### HasHandoffTokenExpiresAt

`func (o *RedeemHandoffResult) HasHandoffTokenExpiresAt() bool`

HasHandoffTokenExpiresAt returns a boolean if a field has been set.

### GetMarketplaceContractId

`func (o *RedeemHandoffResult) GetMarketplaceContractId() string`

GetMarketplaceContractId returns the MarketplaceContractId field if non-nil, zero value otherwise.

### GetMarketplaceContractIdOk

`func (o *RedeemHandoffResult) GetMarketplaceContractIdOk() (*string, bool)`

GetMarketplaceContractIdOk returns a tuple with the MarketplaceContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceContractId

`func (o *RedeemHandoffResult) SetMarketplaceContractId(v string)`

SetMarketplaceContractId sets MarketplaceContractId field to given value.


### GetOrg

`func (o *RedeemHandoffResult) GetOrg() MarketplaceEventOrg`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *RedeemHandoffResult) GetOrgOk() (*MarketplaceEventOrg, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *RedeemHandoffResult) SetOrg(v MarketplaceEventOrg)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *RedeemHandoffResult) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPlan

`func (o *RedeemHandoffResult) GetPlan() MarketplaceEventPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *RedeemHandoffResult) GetPlanOk() (*MarketplaceEventPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *RedeemHandoffResult) SetPlan(v MarketplaceEventPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *RedeemHandoffResult) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *RedeemHandoffResult) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RedeemHandoffResult) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RedeemHandoffResult) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RedeemHandoffResult) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionRequest

`func (o *RedeemHandoffResult) GetSubscriptionRequest() MarketplaceEventSubscriptionRequest`

GetSubscriptionRequest returns the SubscriptionRequest field if non-nil, zero value otherwise.

### GetSubscriptionRequestOk

`func (o *RedeemHandoffResult) GetSubscriptionRequestOk() (*MarketplaceEventSubscriptionRequest, bool)`

GetSubscriptionRequestOk returns a tuple with the SubscriptionRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRequest

`func (o *RedeemHandoffResult) SetSubscriptionRequest(v MarketplaceEventSubscriptionRequest)`

SetSubscriptionRequest sets SubscriptionRequest field to given value.

### HasSubscriptionRequest

`func (o *RedeemHandoffResult) HasSubscriptionRequest() bool`

HasSubscriptionRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


