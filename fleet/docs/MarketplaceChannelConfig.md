# MarketplaceChannelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoConfirmIsv** | Pointer to **bool** | Confirms on the ISV&#39;s behalf as soon as the request exists. For an ISV with no onboarding of their own to do, and off by default because it removes the gate | [optional] 
**BillingBinding** | [**MarketplaceBillingBinding**](MarketplaceBillingBinding.md) |  | 
**Capabilities** | [**MarketplaceChannelCapabilities**](MarketplaceChannelCapabilities.md) |  | 
**Channel** | **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | 
**ContractCount** | Pointer to **int64** | Real contracts. Simulated ones are counted separately and are excluded from every revenue rollup | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CredentialsSet** | **bool** | Whether a credential is stored, without saying what it is. The only thing a read is entitled to know about a secret | 
**DefaultProductTierId** | Pointer to **string** | Where a purchase lands when the channel reports a plan that planMap does not name | [optional] 
**DefaultServiceEnvironmentId** | Pointer to **string** |  | [optional] 
**DefaultServiceId** | Pointer to **string** |  | [optional] 
**DimensionMap** | Pointer to **map[string]string** | Omnistrate metering dimension to the channel&#39;s billable dimension key | [optional] 
**Enabled** | **bool** | Whether contracts on this channel are fulfilled. A config can exist and be disabled, which is the state a partially configured connection sits in | 
**EventReceivers** | Pointer to [**[]MarketplaceEventReceiver**](MarketplaceEventReceiver.md) | Where each event type is delivered. An event with no entry here is NOT delivered; there is no inheritance and no default. Channels that predate per-event routing have an entry per event carrying whatever their single receiver was, so nothing goes quiet without somebody removing it | [optional] 
**HandoffTokenValiditySeconds** | Pointer to **int64** | How long a handoff credential this channel mints stays redeemable. Omit it for the platform default of seven days.  Separate from isvConfirmTimeoutSeconds on purpose. That one is the SLA, and the default validity is deliberately longer than it, so an ISV who breaches the SLA can still recover on their own rather than needing an operator to reissue. One is how long we wait before raising an alarm; the other is how long the credential works.  Between 3600 (an hour) and 2592000 (thirty days). The floor is because the chain from checkout to your callback is several hops and a browser is free to be slow at any of them, so a shorter credential is one a buyer can be handed already dead. The ceiling is the window in which a cloud marketplace can void a purchase; a credential outliving that keeps working after the thing it selects has stopped being real | [optional] 
**Id** | **string** |  | 
**IsSimulated** | **bool** | Derived from the channel, never accepted from a request. Every no-real-money affordance is gated on this rather than on the channel name | 
**IsvCallbackUrl** | Pointer to **string** | Where the buyer&#39;s BROWSER is sent after onboarding, with one appended query parameter. Not the webhook receiver: see portBReceiverUrl. Required before the channel can be enabled, because without it there is nowhere to send them | [optional] 
**IsvConfirmTimeoutSeconds** | Pointer to **int64** | The handoff SLA. On expiry the contract is flagged for an operator and the ISV is told; the fulfillment state does not move, because a stuck contract is stuck exactly where it was | [optional] 
**LandingUrl** | **string** | Where the buyer arrives from the marketplace. Generated per organization and channel and never editable, because it is the address the marketplace was told about | 
**LastSyncedAt** | Pointer to **time.Time** |  | [optional] 
**PlanMap** | Pointer to [**map[string]MarketplacePlanMapping**](MarketplacePlanMapping.md) | The channel&#39;s listing identifier to the Omnistrate service and plan it is sold as. Mapping a listing here also disables self-serve subscription on that plan, because a marketplace plan a customer can subscribe to directly is a route around the ISV confirm | [optional] 
**PortBReceiverUrl** | Pointer to **string** | DEPRECATED, and read only. The single receiver every event used to be delivered to, now migrated into an eventReceivers entry per event type. Still returned so a client built before per-event routing keeps working; write eventReceivers instead | [optional] 
**PortBSigningSecretId** | Pointer to **string** | Which secret is active, so a rotation can be confirmed by watching this change without anything sensitive being displayed. Not the secret and not derived from it | [optional] 
**PortBSigningSecretSet** | Pointer to **bool** | Whether a signing secret is configured. Deliveries are unsigned without one, and a conforming receiver rejects every unsigned delivery, so this is the field that explains a channel whose webhooks are all being refused | [optional] 
**SimulatedContractCount** | Pointer to **int64** |  | [optional] 
**Status** | **string** | The health of one channel connection. NOT_CONNECTED has no credentials yet. DEGRADED answers but is behind, so contracts may be stale. ERROR failed outright | 
**StatusMessage** | Pointer to **string** | Why the status is what it is, in words. The field an operator reads first | [optional] 
**SyntheticEmailDomain** | Pointer to **string** | The domain the buyer&#39;s derived address is minted under. Several marketplaces return no buyer email at all, so the address is derived rather than collected | [optional] 
**WebhookUrl** | Pointer to **string** | Where a channel that posts lifecycle changes to a webhook should post them. Generated, not editable. Absent on a channel that does not use one | [optional] 

## Methods

### NewMarketplaceChannelConfig

`func NewMarketplaceChannelConfig(billingBinding MarketplaceBillingBinding, capabilities MarketplaceChannelCapabilities, channel string, credentialsSet bool, enabled bool, id string, isSimulated bool, landingUrl string, status string, ) *MarketplaceChannelConfig`

NewMarketplaceChannelConfig instantiates a new MarketplaceChannelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceChannelConfigWithDefaults

`func NewMarketplaceChannelConfigWithDefaults() *MarketplaceChannelConfig`

NewMarketplaceChannelConfigWithDefaults instantiates a new MarketplaceChannelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoConfirmIsv

`func (o *MarketplaceChannelConfig) GetAutoConfirmIsv() bool`

GetAutoConfirmIsv returns the AutoConfirmIsv field if non-nil, zero value otherwise.

### GetAutoConfirmIsvOk

`func (o *MarketplaceChannelConfig) GetAutoConfirmIsvOk() (*bool, bool)`

GetAutoConfirmIsvOk returns a tuple with the AutoConfirmIsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfirmIsv

`func (o *MarketplaceChannelConfig) SetAutoConfirmIsv(v bool)`

SetAutoConfirmIsv sets AutoConfirmIsv field to given value.

### HasAutoConfirmIsv

`func (o *MarketplaceChannelConfig) HasAutoConfirmIsv() bool`

HasAutoConfirmIsv returns a boolean if a field has been set.

### GetBillingBinding

`func (o *MarketplaceChannelConfig) GetBillingBinding() MarketplaceBillingBinding`

GetBillingBinding returns the BillingBinding field if non-nil, zero value otherwise.

### GetBillingBindingOk

`func (o *MarketplaceChannelConfig) GetBillingBindingOk() (*MarketplaceBillingBinding, bool)`

GetBillingBindingOk returns a tuple with the BillingBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingBinding

`func (o *MarketplaceChannelConfig) SetBillingBinding(v MarketplaceBillingBinding)`

SetBillingBinding sets BillingBinding field to given value.


### GetCapabilities

`func (o *MarketplaceChannelConfig) GetCapabilities() MarketplaceChannelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *MarketplaceChannelConfig) GetCapabilitiesOk() (*MarketplaceChannelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *MarketplaceChannelConfig) SetCapabilities(v MarketplaceChannelCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetChannel

`func (o *MarketplaceChannelConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MarketplaceChannelConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MarketplaceChannelConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetContractCount

`func (o *MarketplaceChannelConfig) GetContractCount() int64`

GetContractCount returns the ContractCount field if non-nil, zero value otherwise.

### GetContractCountOk

`func (o *MarketplaceChannelConfig) GetContractCountOk() (*int64, bool)`

GetContractCountOk returns a tuple with the ContractCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCount

`func (o *MarketplaceChannelConfig) SetContractCount(v int64)`

SetContractCount sets ContractCount field to given value.

### HasContractCount

`func (o *MarketplaceChannelConfig) HasContractCount() bool`

HasContractCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketplaceChannelConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketplaceChannelConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketplaceChannelConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketplaceChannelConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialsSet

`func (o *MarketplaceChannelConfig) GetCredentialsSet() bool`

GetCredentialsSet returns the CredentialsSet field if non-nil, zero value otherwise.

### GetCredentialsSetOk

`func (o *MarketplaceChannelConfig) GetCredentialsSetOk() (*bool, bool)`

GetCredentialsSetOk returns a tuple with the CredentialsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsSet

`func (o *MarketplaceChannelConfig) SetCredentialsSet(v bool)`

SetCredentialsSet sets CredentialsSet field to given value.


### GetDefaultProductTierId

`func (o *MarketplaceChannelConfig) GetDefaultProductTierId() string`

GetDefaultProductTierId returns the DefaultProductTierId field if non-nil, zero value otherwise.

### GetDefaultProductTierIdOk

`func (o *MarketplaceChannelConfig) GetDefaultProductTierIdOk() (*string, bool)`

GetDefaultProductTierIdOk returns a tuple with the DefaultProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProductTierId

`func (o *MarketplaceChannelConfig) SetDefaultProductTierId(v string)`

SetDefaultProductTierId sets DefaultProductTierId field to given value.

### HasDefaultProductTierId

`func (o *MarketplaceChannelConfig) HasDefaultProductTierId() bool`

HasDefaultProductTierId returns a boolean if a field has been set.

### GetDefaultServiceEnvironmentId

`func (o *MarketplaceChannelConfig) GetDefaultServiceEnvironmentId() string`

GetDefaultServiceEnvironmentId returns the DefaultServiceEnvironmentId field if non-nil, zero value otherwise.

### GetDefaultServiceEnvironmentIdOk

`func (o *MarketplaceChannelConfig) GetDefaultServiceEnvironmentIdOk() (*string, bool)`

GetDefaultServiceEnvironmentIdOk returns a tuple with the DefaultServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceEnvironmentId

`func (o *MarketplaceChannelConfig) SetDefaultServiceEnvironmentId(v string)`

SetDefaultServiceEnvironmentId sets DefaultServiceEnvironmentId field to given value.

### HasDefaultServiceEnvironmentId

`func (o *MarketplaceChannelConfig) HasDefaultServiceEnvironmentId() bool`

HasDefaultServiceEnvironmentId returns a boolean if a field has been set.

### GetDefaultServiceId

`func (o *MarketplaceChannelConfig) GetDefaultServiceId() string`

GetDefaultServiceId returns the DefaultServiceId field if non-nil, zero value otherwise.

### GetDefaultServiceIdOk

`func (o *MarketplaceChannelConfig) GetDefaultServiceIdOk() (*string, bool)`

GetDefaultServiceIdOk returns a tuple with the DefaultServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceId

`func (o *MarketplaceChannelConfig) SetDefaultServiceId(v string)`

SetDefaultServiceId sets DefaultServiceId field to given value.

### HasDefaultServiceId

`func (o *MarketplaceChannelConfig) HasDefaultServiceId() bool`

HasDefaultServiceId returns a boolean if a field has been set.

### GetDimensionMap

`func (o *MarketplaceChannelConfig) GetDimensionMap() map[string]string`

GetDimensionMap returns the DimensionMap field if non-nil, zero value otherwise.

### GetDimensionMapOk

`func (o *MarketplaceChannelConfig) GetDimensionMapOk() (*map[string]string, bool)`

GetDimensionMapOk returns a tuple with the DimensionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionMap

`func (o *MarketplaceChannelConfig) SetDimensionMap(v map[string]string)`

SetDimensionMap sets DimensionMap field to given value.

### HasDimensionMap

`func (o *MarketplaceChannelConfig) HasDimensionMap() bool`

HasDimensionMap returns a boolean if a field has been set.

### GetEnabled

`func (o *MarketplaceChannelConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MarketplaceChannelConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MarketplaceChannelConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEventReceivers

`func (o *MarketplaceChannelConfig) GetEventReceivers() []MarketplaceEventReceiver`

GetEventReceivers returns the EventReceivers field if non-nil, zero value otherwise.

### GetEventReceiversOk

`func (o *MarketplaceChannelConfig) GetEventReceiversOk() (*[]MarketplaceEventReceiver, bool)`

GetEventReceiversOk returns a tuple with the EventReceivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReceivers

`func (o *MarketplaceChannelConfig) SetEventReceivers(v []MarketplaceEventReceiver)`

SetEventReceivers sets EventReceivers field to given value.

### HasEventReceivers

`func (o *MarketplaceChannelConfig) HasEventReceivers() bool`

HasEventReceivers returns a boolean if a field has been set.

### GetHandoffTokenValiditySeconds

`func (o *MarketplaceChannelConfig) GetHandoffTokenValiditySeconds() int64`

GetHandoffTokenValiditySeconds returns the HandoffTokenValiditySeconds field if non-nil, zero value otherwise.

### GetHandoffTokenValiditySecondsOk

`func (o *MarketplaceChannelConfig) GetHandoffTokenValiditySecondsOk() (*int64, bool)`

GetHandoffTokenValiditySecondsOk returns a tuple with the HandoffTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffTokenValiditySeconds

`func (o *MarketplaceChannelConfig) SetHandoffTokenValiditySeconds(v int64)`

SetHandoffTokenValiditySeconds sets HandoffTokenValiditySeconds field to given value.

### HasHandoffTokenValiditySeconds

`func (o *MarketplaceChannelConfig) HasHandoffTokenValiditySeconds() bool`

HasHandoffTokenValiditySeconds returns a boolean if a field has been set.

### GetId

`func (o *MarketplaceChannelConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceChannelConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceChannelConfig) SetId(v string)`

SetId sets Id field to given value.


### GetIsSimulated

`func (o *MarketplaceChannelConfig) GetIsSimulated() bool`

GetIsSimulated returns the IsSimulated field if non-nil, zero value otherwise.

### GetIsSimulatedOk

`func (o *MarketplaceChannelConfig) GetIsSimulatedOk() (*bool, bool)`

GetIsSimulatedOk returns a tuple with the IsSimulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSimulated

`func (o *MarketplaceChannelConfig) SetIsSimulated(v bool)`

SetIsSimulated sets IsSimulated field to given value.


### GetIsvCallbackUrl

`func (o *MarketplaceChannelConfig) GetIsvCallbackUrl() string`

GetIsvCallbackUrl returns the IsvCallbackUrl field if non-nil, zero value otherwise.

### GetIsvCallbackUrlOk

`func (o *MarketplaceChannelConfig) GetIsvCallbackUrlOk() (*string, bool)`

GetIsvCallbackUrlOk returns a tuple with the IsvCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvCallbackUrl

`func (o *MarketplaceChannelConfig) SetIsvCallbackUrl(v string)`

SetIsvCallbackUrl sets IsvCallbackUrl field to given value.

### HasIsvCallbackUrl

`func (o *MarketplaceChannelConfig) HasIsvCallbackUrl() bool`

HasIsvCallbackUrl returns a boolean if a field has been set.

### GetIsvConfirmTimeoutSeconds

`func (o *MarketplaceChannelConfig) GetIsvConfirmTimeoutSeconds() int64`

GetIsvConfirmTimeoutSeconds returns the IsvConfirmTimeoutSeconds field if non-nil, zero value otherwise.

### GetIsvConfirmTimeoutSecondsOk

`func (o *MarketplaceChannelConfig) GetIsvConfirmTimeoutSecondsOk() (*int64, bool)`

GetIsvConfirmTimeoutSecondsOk returns a tuple with the IsvConfirmTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvConfirmTimeoutSeconds

`func (o *MarketplaceChannelConfig) SetIsvConfirmTimeoutSeconds(v int64)`

SetIsvConfirmTimeoutSeconds sets IsvConfirmTimeoutSeconds field to given value.

### HasIsvConfirmTimeoutSeconds

`func (o *MarketplaceChannelConfig) HasIsvConfirmTimeoutSeconds() bool`

HasIsvConfirmTimeoutSeconds returns a boolean if a field has been set.

### GetLandingUrl

`func (o *MarketplaceChannelConfig) GetLandingUrl() string`

GetLandingUrl returns the LandingUrl field if non-nil, zero value otherwise.

### GetLandingUrlOk

`func (o *MarketplaceChannelConfig) GetLandingUrlOk() (*string, bool)`

GetLandingUrlOk returns a tuple with the LandingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrl

`func (o *MarketplaceChannelConfig) SetLandingUrl(v string)`

SetLandingUrl sets LandingUrl field to given value.


### GetLastSyncedAt

`func (o *MarketplaceChannelConfig) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *MarketplaceChannelConfig) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *MarketplaceChannelConfig) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *MarketplaceChannelConfig) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetPlanMap

`func (o *MarketplaceChannelConfig) GetPlanMap() map[string]MarketplacePlanMapping`

GetPlanMap returns the PlanMap field if non-nil, zero value otherwise.

### GetPlanMapOk

`func (o *MarketplaceChannelConfig) GetPlanMapOk() (*map[string]MarketplacePlanMapping, bool)`

GetPlanMapOk returns a tuple with the PlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMap

`func (o *MarketplaceChannelConfig) SetPlanMap(v map[string]MarketplacePlanMapping)`

SetPlanMap sets PlanMap field to given value.

### HasPlanMap

`func (o *MarketplaceChannelConfig) HasPlanMap() bool`

HasPlanMap returns a boolean if a field has been set.

### GetPortBReceiverUrl

`func (o *MarketplaceChannelConfig) GetPortBReceiverUrl() string`

GetPortBReceiverUrl returns the PortBReceiverUrl field if non-nil, zero value otherwise.

### GetPortBReceiverUrlOk

`func (o *MarketplaceChannelConfig) GetPortBReceiverUrlOk() (*string, bool)`

GetPortBReceiverUrlOk returns a tuple with the PortBReceiverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortBReceiverUrl

`func (o *MarketplaceChannelConfig) SetPortBReceiverUrl(v string)`

SetPortBReceiverUrl sets PortBReceiverUrl field to given value.

### HasPortBReceiverUrl

`func (o *MarketplaceChannelConfig) HasPortBReceiverUrl() bool`

HasPortBReceiverUrl returns a boolean if a field has been set.

### GetPortBSigningSecretId

`func (o *MarketplaceChannelConfig) GetPortBSigningSecretId() string`

GetPortBSigningSecretId returns the PortBSigningSecretId field if non-nil, zero value otherwise.

### GetPortBSigningSecretIdOk

`func (o *MarketplaceChannelConfig) GetPortBSigningSecretIdOk() (*string, bool)`

GetPortBSigningSecretIdOk returns a tuple with the PortBSigningSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortBSigningSecretId

`func (o *MarketplaceChannelConfig) SetPortBSigningSecretId(v string)`

SetPortBSigningSecretId sets PortBSigningSecretId field to given value.

### HasPortBSigningSecretId

`func (o *MarketplaceChannelConfig) HasPortBSigningSecretId() bool`

HasPortBSigningSecretId returns a boolean if a field has been set.

### GetPortBSigningSecretSet

`func (o *MarketplaceChannelConfig) GetPortBSigningSecretSet() bool`

GetPortBSigningSecretSet returns the PortBSigningSecretSet field if non-nil, zero value otherwise.

### GetPortBSigningSecretSetOk

`func (o *MarketplaceChannelConfig) GetPortBSigningSecretSetOk() (*bool, bool)`

GetPortBSigningSecretSetOk returns a tuple with the PortBSigningSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortBSigningSecretSet

`func (o *MarketplaceChannelConfig) SetPortBSigningSecretSet(v bool)`

SetPortBSigningSecretSet sets PortBSigningSecretSet field to given value.

### HasPortBSigningSecretSet

`func (o *MarketplaceChannelConfig) HasPortBSigningSecretSet() bool`

HasPortBSigningSecretSet returns a boolean if a field has been set.

### GetSimulatedContractCount

`func (o *MarketplaceChannelConfig) GetSimulatedContractCount() int64`

GetSimulatedContractCount returns the SimulatedContractCount field if non-nil, zero value otherwise.

### GetSimulatedContractCountOk

`func (o *MarketplaceChannelConfig) GetSimulatedContractCountOk() (*int64, bool)`

GetSimulatedContractCountOk returns a tuple with the SimulatedContractCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulatedContractCount

`func (o *MarketplaceChannelConfig) SetSimulatedContractCount(v int64)`

SetSimulatedContractCount sets SimulatedContractCount field to given value.

### HasSimulatedContractCount

`func (o *MarketplaceChannelConfig) HasSimulatedContractCount() bool`

HasSimulatedContractCount returns a boolean if a field has been set.

### GetStatus

`func (o *MarketplaceChannelConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceChannelConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceChannelConfig) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *MarketplaceChannelConfig) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *MarketplaceChannelConfig) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *MarketplaceChannelConfig) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *MarketplaceChannelConfig) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSyntheticEmailDomain

`func (o *MarketplaceChannelConfig) GetSyntheticEmailDomain() string`

GetSyntheticEmailDomain returns the SyntheticEmailDomain field if non-nil, zero value otherwise.

### GetSyntheticEmailDomainOk

`func (o *MarketplaceChannelConfig) GetSyntheticEmailDomainOk() (*string, bool)`

GetSyntheticEmailDomainOk returns a tuple with the SyntheticEmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEmailDomain

`func (o *MarketplaceChannelConfig) SetSyntheticEmailDomain(v string)`

SetSyntheticEmailDomain sets SyntheticEmailDomain field to given value.

### HasSyntheticEmailDomain

`func (o *MarketplaceChannelConfig) HasSyntheticEmailDomain() bool`

HasSyntheticEmailDomain returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *MarketplaceChannelConfig) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *MarketplaceChannelConfig) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *MarketplaceChannelConfig) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *MarketplaceChannelConfig) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


