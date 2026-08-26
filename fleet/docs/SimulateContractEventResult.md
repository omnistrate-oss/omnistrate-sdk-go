# SimulateContractEventResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractStatus** | Pointer to **string** | What the marketplace says about the contract. Mirrored, never authoritative for whether the buyer may be served | [optional] 
**ContractVersion** | Pointer to **int64** |  | [optional] 
**Control** | **string** | Which lifecycle event to simulate. Each drives the real fulfillment path | 
**ExternalRef** | Pointer to **string** | The armed purchase&#39;s reference on the simulated channel, returned by &#x60;purchase&#x60;. This is what identifies it before any Omnistrate contract exists, and what the checkout link resolves | [optional] 
**FulfillmentState** | Pointer to **string** | What Omnistrate decided. Deployments and metering are allowed if and only if this is READY, plus exactly one final metering window during DEPROVISIONING | [optional] 
**HandoffStalledUntil** | Pointer to **time.Time** | Set by stall_handoff. The contract sits in AWAITING_ISV past its SLA so the orphan alarm and the paying-and-cannot-use-it quadrant can be seen firing | [optional] 
**MarketplaceCheckoutUrl** | Pointer to **string** | Where to go to be the buyer, returned by &#x60;purchase&#x60; and by nothing else. Arming a purchase is the marketplace&#39;s side of a checkout, so this stands in for the page a buyer completes it on: following it redirects onward through the real landing route, exactly as a marketplace redirects a real browser. Until it is followed there is no Omnistrate contract at all, which is also true of a real purchase and is the whole reason fulfillment starts at the landing route rather than earlier | [optional] 
**MarketplaceContractId** | Pointer to **string** | The Omnistrate contract. ABSENT after &#x60;purchase&#x60;, because arming one creates no Omnistrate contract: the purchase exists on the marketplace and nothing exists here until the buyer arrives, which is the same thing that is true of a real purchase | [optional] 
**MeteringUnlockedAt** | Pointer to **time.Time** | Set by release_usage_gate, through the same activity the real AWS License Updated event uses. That shared path is what makes rehearsing the hard-gated case a test of the production gate rather than of the console | [optional] 
**SubscriptionRequestId** | Pointer to **string** | Present once a purchase has reached AWAITING_ISV. This is the id to approve | [optional] 

## Methods

### NewSimulateContractEventResult

`func NewSimulateContractEventResult(control string, ) *SimulateContractEventResult`

NewSimulateContractEventResult instantiates a new SimulateContractEventResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateContractEventResultWithDefaults

`func NewSimulateContractEventResultWithDefaults() *SimulateContractEventResult`

NewSimulateContractEventResultWithDefaults instantiates a new SimulateContractEventResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractStatus

`func (o *SimulateContractEventResult) GetContractStatus() string`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *SimulateContractEventResult) GetContractStatusOk() (*string, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *SimulateContractEventResult) SetContractStatus(v string)`

SetContractStatus sets ContractStatus field to given value.

### HasContractStatus

`func (o *SimulateContractEventResult) HasContractStatus() bool`

HasContractStatus returns a boolean if a field has been set.

### GetContractVersion

`func (o *SimulateContractEventResult) GetContractVersion() int64`

GetContractVersion returns the ContractVersion field if non-nil, zero value otherwise.

### GetContractVersionOk

`func (o *SimulateContractEventResult) GetContractVersionOk() (*int64, bool)`

GetContractVersionOk returns a tuple with the ContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractVersion

`func (o *SimulateContractEventResult) SetContractVersion(v int64)`

SetContractVersion sets ContractVersion field to given value.

### HasContractVersion

`func (o *SimulateContractEventResult) HasContractVersion() bool`

HasContractVersion returns a boolean if a field has been set.

### GetControl

`func (o *SimulateContractEventResult) GetControl() string`

GetControl returns the Control field if non-nil, zero value otherwise.

### GetControlOk

`func (o *SimulateContractEventResult) GetControlOk() (*string, bool)`

GetControlOk returns a tuple with the Control field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControl

`func (o *SimulateContractEventResult) SetControl(v string)`

SetControl sets Control field to given value.


### GetExternalRef

`func (o *SimulateContractEventResult) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *SimulateContractEventResult) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *SimulateContractEventResult) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.

### HasExternalRef

`func (o *SimulateContractEventResult) HasExternalRef() bool`

HasExternalRef returns a boolean if a field has been set.

### GetFulfillmentState

`func (o *SimulateContractEventResult) GetFulfillmentState() string`

GetFulfillmentState returns the FulfillmentState field if non-nil, zero value otherwise.

### GetFulfillmentStateOk

`func (o *SimulateContractEventResult) GetFulfillmentStateOk() (*string, bool)`

GetFulfillmentStateOk returns a tuple with the FulfillmentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentState

`func (o *SimulateContractEventResult) SetFulfillmentState(v string)`

SetFulfillmentState sets FulfillmentState field to given value.

### HasFulfillmentState

`func (o *SimulateContractEventResult) HasFulfillmentState() bool`

HasFulfillmentState returns a boolean if a field has been set.

### GetHandoffStalledUntil

`func (o *SimulateContractEventResult) GetHandoffStalledUntil() time.Time`

GetHandoffStalledUntil returns the HandoffStalledUntil field if non-nil, zero value otherwise.

### GetHandoffStalledUntilOk

`func (o *SimulateContractEventResult) GetHandoffStalledUntilOk() (*time.Time, bool)`

GetHandoffStalledUntilOk returns a tuple with the HandoffStalledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffStalledUntil

`func (o *SimulateContractEventResult) SetHandoffStalledUntil(v time.Time)`

SetHandoffStalledUntil sets HandoffStalledUntil field to given value.

### HasHandoffStalledUntil

`func (o *SimulateContractEventResult) HasHandoffStalledUntil() bool`

HasHandoffStalledUntil returns a boolean if a field has been set.

### GetMarketplaceCheckoutUrl

`func (o *SimulateContractEventResult) GetMarketplaceCheckoutUrl() string`

GetMarketplaceCheckoutUrl returns the MarketplaceCheckoutUrl field if non-nil, zero value otherwise.

### GetMarketplaceCheckoutUrlOk

`func (o *SimulateContractEventResult) GetMarketplaceCheckoutUrlOk() (*string, bool)`

GetMarketplaceCheckoutUrlOk returns a tuple with the MarketplaceCheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceCheckoutUrl

`func (o *SimulateContractEventResult) SetMarketplaceCheckoutUrl(v string)`

SetMarketplaceCheckoutUrl sets MarketplaceCheckoutUrl field to given value.

### HasMarketplaceCheckoutUrl

`func (o *SimulateContractEventResult) HasMarketplaceCheckoutUrl() bool`

HasMarketplaceCheckoutUrl returns a boolean if a field has been set.

### GetMarketplaceContractId

`func (o *SimulateContractEventResult) GetMarketplaceContractId() string`

GetMarketplaceContractId returns the MarketplaceContractId field if non-nil, zero value otherwise.

### GetMarketplaceContractIdOk

`func (o *SimulateContractEventResult) GetMarketplaceContractIdOk() (*string, bool)`

GetMarketplaceContractIdOk returns a tuple with the MarketplaceContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceContractId

`func (o *SimulateContractEventResult) SetMarketplaceContractId(v string)`

SetMarketplaceContractId sets MarketplaceContractId field to given value.

### HasMarketplaceContractId

`func (o *SimulateContractEventResult) HasMarketplaceContractId() bool`

HasMarketplaceContractId returns a boolean if a field has been set.

### GetMeteringUnlockedAt

`func (o *SimulateContractEventResult) GetMeteringUnlockedAt() time.Time`

GetMeteringUnlockedAt returns the MeteringUnlockedAt field if non-nil, zero value otherwise.

### GetMeteringUnlockedAtOk

`func (o *SimulateContractEventResult) GetMeteringUnlockedAtOk() (*time.Time, bool)`

GetMeteringUnlockedAtOk returns a tuple with the MeteringUnlockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteringUnlockedAt

`func (o *SimulateContractEventResult) SetMeteringUnlockedAt(v time.Time)`

SetMeteringUnlockedAt sets MeteringUnlockedAt field to given value.

### HasMeteringUnlockedAt

`func (o *SimulateContractEventResult) HasMeteringUnlockedAt() bool`

HasMeteringUnlockedAt returns a boolean if a field has been set.

### GetSubscriptionRequestId

`func (o *SimulateContractEventResult) GetSubscriptionRequestId() string`

GetSubscriptionRequestId returns the SubscriptionRequestId field if non-nil, zero value otherwise.

### GetSubscriptionRequestIdOk

`func (o *SimulateContractEventResult) GetSubscriptionRequestIdOk() (*string, bool)`

GetSubscriptionRequestIdOk returns a tuple with the SubscriptionRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionRequestId

`func (o *SimulateContractEventResult) SetSubscriptionRequestId(v string)`

SetSubscriptionRequestId sets SubscriptionRequestId field to given value.

### HasSubscriptionRequestId

`func (o *SimulateContractEventResult) HasSubscriptionRequestId() bool`

HasSubscriptionRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


