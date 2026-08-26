# MarketplaceContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttentionSince** | Pointer to **time.Time** | Set while the contract is billable and fulfillment has not reached READY. How long it has been set is what the handoff SLA is measured against | [optional] 
**BuyerName** | Pointer to **string** |  | [optional] 
**BuyerRef** | Pointer to **string** |  | [optional] 
**Channel** | **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | 
**CommitAmount** | Pointer to **string** | Decimal string, never a number, for the same reason as on the event payload | [optional] 
**ContractStatus** | **string** | What the marketplace says about the contract. Mirrored, never authoritative for whether the buyer may be served | 
**ContractSyncedAt** | Pointer to **time.Time** | When the mirror last read this contract back from the channel | [optional] 
**ContractVersion** | Pointer to **int64** | Monotonic per contract. The value an ISV uses to discard a stale delivery | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**Events** | Pointer to [**[]MarketplaceContractEvent**](MarketplaceContractEvent.md) | Absent on the list response, which is a summary | [optional] 
**ExternalRef** | **string** | The channel&#39;s own identifier, unique per channel and the idempotency key for every write path | 
**FulfillmentState** | **string** | What Omnistrate decided. Deployments and metering are allowed if and only if this is READY, plus exactly one final metering window during DEPROVISIONING | 
**Id** | **string** |  | 
**IsSimulated** | **bool** | True when the contract came from a simulated channel. Derived from the channel and never accepted from a request. Simulated contracts are excluded from revenue rollups, invoices and usage exports by default | 
**OrgId** | Pointer to **string** | Populated once the buyer org and synthetic root user exist | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**PlanRef** | Pointer to **string** |  | [optional] 
**Quadrant** | Pointer to **string** | Where this contract sits in the reconciliation 2x2. PAYING_NOT_PROVISIONED is the one that costs money: the buyer is being billed and cannot use the product. ONBOARDING is on neither axis: the channel is HOLDING the contract, so the buyer is deliberately not being charged and is not yet being served, which is the normal path on a channel that can hold. It is called out because the four quadrants have nowhere to put it, and filing a held contract under CLOSED said an onboarding buyer had ended | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**RawContract** | Pointer to **map[string]interface{}** | The channel&#39;s native payload, verbatim, for support and reconciliation | [optional] 
**RootUserEmail** | Pointer to **string** | The synthetic address. An identifier rather than a deliverable mailbox | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**SubscriptionId** | Pointer to **string** | Only ever populated after the ISV approved, because the approval is what creates the subscription | [optional] 

## Methods

### NewMarketplaceContract

`func NewMarketplaceContract(channel string, contractStatus string, externalRef string, fulfillmentState string, id string, isSimulated bool, ) *MarketplaceContract`

NewMarketplaceContract instantiates a new MarketplaceContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceContractWithDefaults

`func NewMarketplaceContractWithDefaults() *MarketplaceContract`

NewMarketplaceContractWithDefaults instantiates a new MarketplaceContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttentionSince

`func (o *MarketplaceContract) GetAttentionSince() time.Time`

GetAttentionSince returns the AttentionSince field if non-nil, zero value otherwise.

### GetAttentionSinceOk

`func (o *MarketplaceContract) GetAttentionSinceOk() (*time.Time, bool)`

GetAttentionSinceOk returns a tuple with the AttentionSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttentionSince

`func (o *MarketplaceContract) SetAttentionSince(v time.Time)`

SetAttentionSince sets AttentionSince field to given value.

### HasAttentionSince

`func (o *MarketplaceContract) HasAttentionSince() bool`

HasAttentionSince returns a boolean if a field has been set.

### GetBuyerName

`func (o *MarketplaceContract) GetBuyerName() string`

GetBuyerName returns the BuyerName field if non-nil, zero value otherwise.

### GetBuyerNameOk

`func (o *MarketplaceContract) GetBuyerNameOk() (*string, bool)`

GetBuyerNameOk returns a tuple with the BuyerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerName

`func (o *MarketplaceContract) SetBuyerName(v string)`

SetBuyerName sets BuyerName field to given value.

### HasBuyerName

`func (o *MarketplaceContract) HasBuyerName() bool`

HasBuyerName returns a boolean if a field has been set.

### GetBuyerRef

`func (o *MarketplaceContract) GetBuyerRef() string`

GetBuyerRef returns the BuyerRef field if non-nil, zero value otherwise.

### GetBuyerRefOk

`func (o *MarketplaceContract) GetBuyerRefOk() (*string, bool)`

GetBuyerRefOk returns a tuple with the BuyerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerRef

`func (o *MarketplaceContract) SetBuyerRef(v string)`

SetBuyerRef sets BuyerRef field to given value.

### HasBuyerRef

`func (o *MarketplaceContract) HasBuyerRef() bool`

HasBuyerRef returns a boolean if a field has been set.

### GetChannel

`func (o *MarketplaceContract) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MarketplaceContract) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MarketplaceContract) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCommitAmount

`func (o *MarketplaceContract) GetCommitAmount() string`

GetCommitAmount returns the CommitAmount field if non-nil, zero value otherwise.

### GetCommitAmountOk

`func (o *MarketplaceContract) GetCommitAmountOk() (*string, bool)`

GetCommitAmountOk returns a tuple with the CommitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitAmount

`func (o *MarketplaceContract) SetCommitAmount(v string)`

SetCommitAmount sets CommitAmount field to given value.

### HasCommitAmount

`func (o *MarketplaceContract) HasCommitAmount() bool`

HasCommitAmount returns a boolean if a field has been set.

### GetContractStatus

`func (o *MarketplaceContract) GetContractStatus() string`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *MarketplaceContract) GetContractStatusOk() (*string, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *MarketplaceContract) SetContractStatus(v string)`

SetContractStatus sets ContractStatus field to given value.


### GetContractSyncedAt

`func (o *MarketplaceContract) GetContractSyncedAt() time.Time`

GetContractSyncedAt returns the ContractSyncedAt field if non-nil, zero value otherwise.

### GetContractSyncedAtOk

`func (o *MarketplaceContract) GetContractSyncedAtOk() (*time.Time, bool)`

GetContractSyncedAtOk returns a tuple with the ContractSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractSyncedAt

`func (o *MarketplaceContract) SetContractSyncedAt(v time.Time)`

SetContractSyncedAt sets ContractSyncedAt field to given value.

### HasContractSyncedAt

`func (o *MarketplaceContract) HasContractSyncedAt() bool`

HasContractSyncedAt returns a boolean if a field has been set.

### GetContractVersion

`func (o *MarketplaceContract) GetContractVersion() int64`

GetContractVersion returns the ContractVersion field if non-nil, zero value otherwise.

### GetContractVersionOk

`func (o *MarketplaceContract) GetContractVersionOk() (*int64, bool)`

GetContractVersionOk returns a tuple with the ContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractVersion

`func (o *MarketplaceContract) SetContractVersion(v int64)`

SetContractVersion sets ContractVersion field to given value.

### HasContractVersion

`func (o *MarketplaceContract) HasContractVersion() bool`

HasContractVersion returns a boolean if a field has been set.

### GetCurrency

`func (o *MarketplaceContract) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MarketplaceContract) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MarketplaceContract) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MarketplaceContract) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEndsAt

`func (o *MarketplaceContract) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *MarketplaceContract) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *MarketplaceContract) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *MarketplaceContract) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetEvents

`func (o *MarketplaceContract) GetEvents() []MarketplaceContractEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MarketplaceContract) GetEventsOk() (*[]MarketplaceContractEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MarketplaceContract) SetEvents(v []MarketplaceContractEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *MarketplaceContract) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExternalRef

`func (o *MarketplaceContract) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *MarketplaceContract) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *MarketplaceContract) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.


### GetFulfillmentState

`func (o *MarketplaceContract) GetFulfillmentState() string`

GetFulfillmentState returns the FulfillmentState field if non-nil, zero value otherwise.

### GetFulfillmentStateOk

`func (o *MarketplaceContract) GetFulfillmentStateOk() (*string, bool)`

GetFulfillmentStateOk returns a tuple with the FulfillmentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentState

`func (o *MarketplaceContract) SetFulfillmentState(v string)`

SetFulfillmentState sets FulfillmentState field to given value.


### GetId

`func (o *MarketplaceContract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceContract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceContract) SetId(v string)`

SetId sets Id field to given value.


### GetIsSimulated

`func (o *MarketplaceContract) GetIsSimulated() bool`

GetIsSimulated returns the IsSimulated field if non-nil, zero value otherwise.

### GetIsSimulatedOk

`func (o *MarketplaceContract) GetIsSimulatedOk() (*bool, bool)`

GetIsSimulatedOk returns a tuple with the IsSimulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSimulated

`func (o *MarketplaceContract) SetIsSimulated(v bool)`

SetIsSimulated sets IsSimulated field to given value.


### GetOrgId

`func (o *MarketplaceContract) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MarketplaceContract) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MarketplaceContract) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MarketplaceContract) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *MarketplaceContract) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *MarketplaceContract) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *MarketplaceContract) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *MarketplaceContract) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPlanRef

`func (o *MarketplaceContract) GetPlanRef() string`

GetPlanRef returns the PlanRef field if non-nil, zero value otherwise.

### GetPlanRefOk

`func (o *MarketplaceContract) GetPlanRefOk() (*string, bool)`

GetPlanRefOk returns a tuple with the PlanRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRef

`func (o *MarketplaceContract) SetPlanRef(v string)`

SetPlanRef sets PlanRef field to given value.

### HasPlanRef

`func (o *MarketplaceContract) HasPlanRef() bool`

HasPlanRef returns a boolean if a field has been set.

### GetQuadrant

`func (o *MarketplaceContract) GetQuadrant() string`

GetQuadrant returns the Quadrant field if non-nil, zero value otherwise.

### GetQuadrantOk

`func (o *MarketplaceContract) GetQuadrantOk() (*string, bool)`

GetQuadrantOk returns a tuple with the Quadrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuadrant

`func (o *MarketplaceContract) SetQuadrant(v string)`

SetQuadrant sets Quadrant field to given value.

### HasQuadrant

`func (o *MarketplaceContract) HasQuadrant() bool`

HasQuadrant returns a boolean if a field has been set.

### GetQuantity

`func (o *MarketplaceContract) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MarketplaceContract) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MarketplaceContract) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MarketplaceContract) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRawContract

`func (o *MarketplaceContract) GetRawContract() map[string]interface{}`

GetRawContract returns the RawContract field if non-nil, zero value otherwise.

### GetRawContractOk

`func (o *MarketplaceContract) GetRawContractOk() (*map[string]interface{}, bool)`

GetRawContractOk returns a tuple with the RawContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawContract

`func (o *MarketplaceContract) SetRawContract(v map[string]interface{})`

SetRawContract sets RawContract field to given value.

### HasRawContract

`func (o *MarketplaceContract) HasRawContract() bool`

HasRawContract returns a boolean if a field has been set.

### GetRootUserEmail

`func (o *MarketplaceContract) GetRootUserEmail() string`

GetRootUserEmail returns the RootUserEmail field if non-nil, zero value otherwise.

### GetRootUserEmailOk

`func (o *MarketplaceContract) GetRootUserEmailOk() (*string, bool)`

GetRootUserEmailOk returns a tuple with the RootUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserEmail

`func (o *MarketplaceContract) SetRootUserEmail(v string)`

SetRootUserEmail sets RootUserEmail field to given value.

### HasRootUserEmail

`func (o *MarketplaceContract) HasRootUserEmail() bool`

HasRootUserEmail returns a boolean if a field has been set.

### GetStartsAt

`func (o *MarketplaceContract) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *MarketplaceContract) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *MarketplaceContract) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *MarketplaceContract) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *MarketplaceContract) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *MarketplaceContract) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *MarketplaceContract) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *MarketplaceContract) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


