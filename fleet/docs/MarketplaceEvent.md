# MarketplaceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerRef** | Pointer to **string** | The channel&#39;s identifier for the buyer | [optional] 
**Capabilities** | Pointer to [**MarketplaceEventCapabilities**](MarketplaceEventCapabilities.md) |  | [optional] 
**Channel** | Pointer to **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | [optional] 
**ContractStatus** | Pointer to **string** | What the marketplace says about the contract. This is MIRRORED state and is never authoritative for whether the buyer may be served | [optional] 
**ContractVersion** | **int64** | Monotonically increasing per contract. Delivery order is NOT guaranteed, so if this is lower than the last version applied for this contract, drop the event. This field is the entire ordering story | 
**DetectedBy** | Pointer to **string** | How the change was noticed. RECONCILIATION means a periodic readback found it rather than the channel announcing it, which is normal on channels that collapse or drop events | [optional] 
**EventId** | **string** | Unique per event and STABLE ACROSS RETRIES. Deduplicate on it: the same event will arrive more than once, and that is normal rather than an incident. A redelivery requested from the console also reuses this id, deliberately, so it exercises your idempotency rather than bypassing it | 
**EventType** | **string** | The type of a marketplace fulfillment event delivered to an ISV receiver | 
**ExternalRef** | Pointer to **string** | The channel&#39;s own identifier for the contract, unique per channel | [optional] 
**FailureReason** | Pointer to **string** | Present on fulfillment.failed. Names what failed in words. A contract stuck past the handoff SLA also carries the orphaned condition here, which does not change fulfillmentState | [optional] 
**FulfillmentState** | Pointer to **string** | What Omnistrate decided. This is the one that governs: the buyer may deploy and is metered if and only if it is READY. A contract can be ACTIVE on the marketplace while fulfillment is still AWAITING_ISV, which means the buyer is paying and cannot yet be served | [optional] 
**HandoffExpiresAt** | Pointer to **time.Time** | When the handoff token stops being redeemable. Deliberately longer than the handoff SLA, so an ISV that breaches the SLA can still recover without an operator re-issuing anything | [optional] 
**HandoffToken** | Pointer to **string** | Present on contract.discovered only. Exchange it at POST /fleet/marketplace/handoff/redeem for the contract detail and the ids needed to approve. Redeeming is idempotent and non-consuming, so the same token returns the same payload until it expires, which is what lets an ISV retry a failed tenant setup without an operator in the loop. It is a bearer value for exactly one contract, so treat it as a credential: do not log it and do not put it in a URL | [optional] 
**MarketplaceContractId** | **string** | The Omnistrate contract identifier, stable for the life of the contract | 
**OccurredAt** | **time.Time** | When the event occurred. This is NOT the send instant: retries carry the original occurredAt while X-Omnistrate-Timestamp moves. Use the header for the replay window and this field for ordering by time | 
**Org** | Pointer to [**MarketplaceEventOrg**](MarketplaceEventOrg.md) |  | [optional] 
**Plan** | Pointer to [**MarketplaceEventPlan**](MarketplaceEventPlan.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | The Omnistrate subscription. Absent before the confirm, because no subscription exists until then; present on every event after it. The confirm response is where an ISV first receives this value | [optional] 
**SubscriptionRequest** | Pointer to [**MarketplaceEventSubscriptionRequest**](MarketplaceEventSubscriptionRequest.md) |  | [optional] 

## Methods

### NewMarketplaceEvent

`func NewMarketplaceEvent(contractVersion int64, eventId string, eventType string, marketplaceContractId string, occurredAt time.Time, ) *MarketplaceEvent`

NewMarketplaceEvent instantiates a new MarketplaceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceEventWithDefaults

`func NewMarketplaceEventWithDefaults() *MarketplaceEvent`

NewMarketplaceEventWithDefaults instantiates a new MarketplaceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerRef

`func (o *MarketplaceEvent) GetBuyerRef() string`

GetBuyerRef returns the BuyerRef field if non-nil, zero value otherwise.

### GetBuyerRefOk

`func (o *MarketplaceEvent) GetBuyerRefOk() (*string, bool)`

GetBuyerRefOk returns a tuple with the BuyerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerRef

`func (o *MarketplaceEvent) SetBuyerRef(v string)`

SetBuyerRef sets BuyerRef field to given value.

### HasBuyerRef

`func (o *MarketplaceEvent) HasBuyerRef() bool`

HasBuyerRef returns a boolean if a field has been set.

### GetCapabilities

`func (o *MarketplaceEvent) GetCapabilities() MarketplaceEventCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *MarketplaceEvent) GetCapabilitiesOk() (*MarketplaceEventCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *MarketplaceEvent) SetCapabilities(v MarketplaceEventCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *MarketplaceEvent) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetChannel

`func (o *MarketplaceEvent) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MarketplaceEvent) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MarketplaceEvent) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *MarketplaceEvent) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetContractStatus

`func (o *MarketplaceEvent) GetContractStatus() string`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *MarketplaceEvent) GetContractStatusOk() (*string, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *MarketplaceEvent) SetContractStatus(v string)`

SetContractStatus sets ContractStatus field to given value.

### HasContractStatus

`func (o *MarketplaceEvent) HasContractStatus() bool`

HasContractStatus returns a boolean if a field has been set.

### GetContractVersion

`func (o *MarketplaceEvent) GetContractVersion() int64`

GetContractVersion returns the ContractVersion field if non-nil, zero value otherwise.

### GetContractVersionOk

`func (o *MarketplaceEvent) GetContractVersionOk() (*int64, bool)`

GetContractVersionOk returns a tuple with the ContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractVersion

`func (o *MarketplaceEvent) SetContractVersion(v int64)`

SetContractVersion sets ContractVersion field to given value.


### GetDetectedBy

`func (o *MarketplaceEvent) GetDetectedBy() string`

GetDetectedBy returns the DetectedBy field if non-nil, zero value otherwise.

### GetDetectedByOk

`func (o *MarketplaceEvent) GetDetectedByOk() (*string, bool)`

GetDetectedByOk returns a tuple with the DetectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedBy

`func (o *MarketplaceEvent) SetDetectedBy(v string)`

SetDetectedBy sets DetectedBy field to given value.

### HasDetectedBy

`func (o *MarketplaceEvent) HasDetectedBy() bool`

HasDetectedBy returns a boolean if a field has been set.

### GetEventId

`func (o *MarketplaceEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *MarketplaceEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *MarketplaceEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventType

`func (o *MarketplaceEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MarketplaceEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MarketplaceEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetExternalRef

`func (o *MarketplaceEvent) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *MarketplaceEvent) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *MarketplaceEvent) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.

### HasExternalRef

`func (o *MarketplaceEvent) HasExternalRef() bool`

HasExternalRef returns a boolean if a field has been set.

### GetFailureReason

`func (o *MarketplaceEvent) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MarketplaceEvent) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *MarketplaceEvent) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *MarketplaceEvent) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetFulfillmentState

`func (o *MarketplaceEvent) GetFulfillmentState() string`

GetFulfillmentState returns the FulfillmentState field if non-nil, zero value otherwise.

### GetFulfillmentStateOk

`func (o *MarketplaceEvent) GetFulfillmentStateOk() (*string, bool)`

GetFulfillmentStateOk returns a tuple with the FulfillmentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentState

`func (o *MarketplaceEvent) SetFulfillmentState(v string)`

SetFulfillmentState sets FulfillmentState field to given value.

### HasFulfillmentState

`func (o *MarketplaceEvent) HasFulfillmentState() bool`

HasFulfillmentState returns a boolean if a field has been set.

### GetHandoffExpiresAt

`func (o *MarketplaceEvent) GetHandoffExpiresAt() time.Time`

GetHandoffExpiresAt returns the HandoffExpiresAt field if non-nil, zero value otherwise.

### GetHandoffExpiresAtOk

`func (o *MarketplaceEvent) GetHandoffExpiresAtOk() (*time.Time, bool)`

GetHandoffExpiresAtOk returns a tuple with the HandoffExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffExpiresAt

`func (o *MarketplaceEvent) SetHandoffExpiresAt(v time.Time)`

SetHandoffExpiresAt sets HandoffExpiresAt field to given value.

### HasHandoffExpiresAt

`func (o *MarketplaceEvent) HasHandoffExpiresAt() bool`

HasHandoffExpiresAt returns a boolean if a field has been set.

### GetHandoffToken

`func (o *MarketplaceEvent) GetHandoffToken() string`

GetHandoffToken returns the HandoffToken field if non-nil, zero value otherwise.

### GetHandoffTokenOk

`func (o *MarketplaceEvent) GetHandoffTokenOk() (*string, bool)`

GetHandoffTokenOk returns a tuple with the HandoffToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffToken

`func (o *MarketplaceEvent) SetHandoffToken(v string)`

SetHandoffToken sets HandoffToken field to given value.

### HasHandoffToken

`func (o *MarketplaceEvent) HasHandoffToken() bool`

HasHandoffToken returns a boolean if a field has been set.

### GetMarketplaceContractId

`func (o *MarketplaceEvent) GetMarketplaceContractId() string`

GetMarketplaceContractId returns the MarketplaceContractId field if non-nil, zero value otherwise.

### GetMarketplaceContractIdOk

`func (o *MarketplaceEvent) GetMarketplaceContractIdOk() (*string, bool)`

GetMarketplaceContractIdOk returns a tuple with the MarketplaceContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceContractId

`func (o *MarketplaceEvent) SetMarketplaceContractId(v string)`

SetMarketplaceContractId sets MarketplaceContractId field to given value.


### GetOccurredAt

`func (o *MarketplaceEvent) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *MarketplaceEvent) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *MarketplaceEvent) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetOrg

`func (o *MarketplaceEvent) GetOrg() MarketplaceEventOrg`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *MarketplaceEvent) GetOrgOk() (*MarketplaceEventOrg, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *MarketplaceEvent) SetOrg(v MarketplaceEventOrg)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *MarketplaceEvent) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPlan

`func (o *MarketplaceEvent) GetPlan() MarketplaceEventPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *MarketplaceEvent) GetPlanOk() (*MarketplaceEventPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *MarketplaceEvent) SetPlan(v MarketplaceEventPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *MarketplaceEvent) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *MarketplaceEvent) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *MarketplaceEvent) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *MarketplaceEvent) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *MarketplaceEvent) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionRequest

`func (o *MarketplaceEvent) GetSubscriptionRequest() MarketplaceEventSubscriptionRequest`

GetSubscriptionRequest returns the SubscriptionRequest field if non-nil, zero value otherwise.

### GetSubscriptionRequestOk

`func (o *MarketplaceEvent) GetSubscriptionRequestOk() (*MarketplaceEventSubscriptionRequest, bool)`

GetSubscriptionRequestOk returns a tuple with the SubscriptionRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRequest

`func (o *MarketplaceEvent) SetSubscriptionRequest(v MarketplaceEventSubscriptionRequest)`

SetSubscriptionRequest sets SubscriptionRequest field to given value.

### HasSubscriptionRequest

`func (o *MarketplaceEvent) HasSubscriptionRequest() bool`

HasSubscriptionRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


