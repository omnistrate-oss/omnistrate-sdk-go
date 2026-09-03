# MarketplaceDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptCount** | Pointer to **int64** | Outbound only. Absent inbound rather than reported as 1 | [optional] 
**Attempts** | Pointer to [**[]SandboxDelivery**](SandboxDelivery.md) | Every attempt with its request and response bytes. The same shape the sandbox conformance run reports, deliberately: an ISV debugging production should be reading the view they already learned in the sandbox | [optional] 
**CallerCredentialId** | Pointer to **string** | Inbound only. Which of the ISV&#39;s credentials was used, by id. Never the credential itself | [optional] 
**Channel** | Pointer to **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | [optional] 
**ContractFulfillmentState** | Pointer to **string** | What Omnistrate decided. Deployments and metering are allowed if and only if this is READY, plus exactly one final metering window during DEPROVISIONING | [optional] 
**ContractRef** | Pointer to **string** | The channel&#39;s own reference for the contract, which is the identifier a buyer quotes and an operator recognises. Omitted from a contract-scoped list, where every row has the same one | [optional] 
**DeliveryId** | **string** |  | 
**DestinationUrl** | Pointer to **string** | Outbound only. Where this was sent, which is not necessarily where the channel is configured to send now | [optional] 
**Direction** | **string** | Who called whom. OUTBOUND is a webhook Omnistrate sent to the ISV&#39;s receiver and INBOUND is a call the ISV made to the marketplace API, both relative to the ISV. CHANNEL_INBOUND is the marketplace calling Omnistrate, and CHANNEL_OUTBOUND is Omnistrate calling the marketplace | 
**Endpoint** | Pointer to **string** | Inbound only. Which API the ISV called | [optional] 
**EventId** | Pointer to **string** | Outbound only. The value a conforming receiver deduplicates on, so a retry carries the same one | [optional] 
**EventType** | Pointer to **string** | The type of a marketplace fulfillment event delivered to an ISV receiver | [optional] 
**IsSimulated** | Pointer to **bool** | True when the interaction belongs to a simulated contract. Derived from the channel and never accepted from a request, so a rehearsal cannot be made to look real | [optional] 
**MarketplaceContractId** | Pointer to **string** |  | [optional] 
**MaxAttempts** | Pointer to **int64** | Outbound only. The retry budget this delivery was sent under, so \&quot;3 of 7\&quot; is readable without opening the row. A number rather than a constant the console repeats, because the budget is a property of the delivery and not of the client drawing it | [optional] 
**Method** | Pointer to **string** | Inbound only. The verb for this endpoint, rendered from it rather than recorded | [optional] 
**NextAttemptAt** | Pointer to **time.Time** | Outbound only, and set only while a retry is still scheduled. Absent means nothing further is coming, which is the difference between a delivery that is failing and one that has stopped | [optional] 
**Path** | Pointer to **string** | Inbound only. The route this endpoint is served at, including the API version, rendered from the endpoint rather than recorded. The logical endpoint above is what filters and health numbers group by; this is what a vendor comparing against their own client logs is looking for | [optional] 
**ReceivedAt** | Pointer to **time.Time** | Inbound only | [optional] 
**RefusalReason** | Pointer to **string** | Inbound only, and set when we refused the call. Named in words, because \&quot;why was my approve rejected\&quot; is otherwise a support ticket | [optional] 
**SentAt** | Pointer to **time.Time** |  | [optional] 
**StateChange** | Pointer to [**MarketplaceInteractionStateChange**](MarketplaceInteractionStateChange.md) |  | [optional] 
**Status** | **string** | PENDING has not been answered yet. DELIVERED was accepted. FAILED exhausted its retries. BLOCKED was never sent, because the receiver address was refused | 
**SubscriptionId** | Pointer to **string** | Set once a delivery or a call has produced a subscription, so a row can be followed to the tenant it created | [optional] 
**TerminalReason** | Pointer to **string** | Outbound only. Why we stopped before the budget ran out, for example a receiver that rejected the request itself. Absent on a delivery that simply exhausted its attempts | [optional] 

## Methods

### NewMarketplaceDelivery

`func NewMarketplaceDelivery(deliveryId string, direction string, status string, ) *MarketplaceDelivery`

NewMarketplaceDelivery instantiates a new MarketplaceDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceDeliveryWithDefaults

`func NewMarketplaceDeliveryWithDefaults() *MarketplaceDelivery`

NewMarketplaceDeliveryWithDefaults instantiates a new MarketplaceDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptCount

`func (o *MarketplaceDelivery) GetAttemptCount() int64`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *MarketplaceDelivery) GetAttemptCountOk() (*int64, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *MarketplaceDelivery) SetAttemptCount(v int64)`

SetAttemptCount sets AttemptCount field to given value.

### HasAttemptCount

`func (o *MarketplaceDelivery) HasAttemptCount() bool`

HasAttemptCount returns a boolean if a field has been set.

### GetAttempts

`func (o *MarketplaceDelivery) GetAttempts() []SandboxDelivery`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *MarketplaceDelivery) GetAttemptsOk() (*[]SandboxDelivery, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *MarketplaceDelivery) SetAttempts(v []SandboxDelivery)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *MarketplaceDelivery) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetCallerCredentialId

`func (o *MarketplaceDelivery) GetCallerCredentialId() string`

GetCallerCredentialId returns the CallerCredentialId field if non-nil, zero value otherwise.

### GetCallerCredentialIdOk

`func (o *MarketplaceDelivery) GetCallerCredentialIdOk() (*string, bool)`

GetCallerCredentialIdOk returns a tuple with the CallerCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerCredentialId

`func (o *MarketplaceDelivery) SetCallerCredentialId(v string)`

SetCallerCredentialId sets CallerCredentialId field to given value.

### HasCallerCredentialId

`func (o *MarketplaceDelivery) HasCallerCredentialId() bool`

HasCallerCredentialId returns a boolean if a field has been set.

### GetChannel

`func (o *MarketplaceDelivery) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MarketplaceDelivery) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MarketplaceDelivery) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *MarketplaceDelivery) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetContractFulfillmentState

`func (o *MarketplaceDelivery) GetContractFulfillmentState() string`

GetContractFulfillmentState returns the ContractFulfillmentState field if non-nil, zero value otherwise.

### GetContractFulfillmentStateOk

`func (o *MarketplaceDelivery) GetContractFulfillmentStateOk() (*string, bool)`

GetContractFulfillmentStateOk returns a tuple with the ContractFulfillmentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractFulfillmentState

`func (o *MarketplaceDelivery) SetContractFulfillmentState(v string)`

SetContractFulfillmentState sets ContractFulfillmentState field to given value.

### HasContractFulfillmentState

`func (o *MarketplaceDelivery) HasContractFulfillmentState() bool`

HasContractFulfillmentState returns a boolean if a field has been set.

### GetContractRef

`func (o *MarketplaceDelivery) GetContractRef() string`

GetContractRef returns the ContractRef field if non-nil, zero value otherwise.

### GetContractRefOk

`func (o *MarketplaceDelivery) GetContractRefOk() (*string, bool)`

GetContractRefOk returns a tuple with the ContractRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractRef

`func (o *MarketplaceDelivery) SetContractRef(v string)`

SetContractRef sets ContractRef field to given value.

### HasContractRef

`func (o *MarketplaceDelivery) HasContractRef() bool`

HasContractRef returns a boolean if a field has been set.

### GetDeliveryId

`func (o *MarketplaceDelivery) GetDeliveryId() string`

GetDeliveryId returns the DeliveryId field if non-nil, zero value otherwise.

### GetDeliveryIdOk

`func (o *MarketplaceDelivery) GetDeliveryIdOk() (*string, bool)`

GetDeliveryIdOk returns a tuple with the DeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryId

`func (o *MarketplaceDelivery) SetDeliveryId(v string)`

SetDeliveryId sets DeliveryId field to given value.


### GetDestinationUrl

`func (o *MarketplaceDelivery) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *MarketplaceDelivery) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *MarketplaceDelivery) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *MarketplaceDelivery) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### GetDirection

`func (o *MarketplaceDelivery) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MarketplaceDelivery) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MarketplaceDelivery) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetEndpoint

`func (o *MarketplaceDelivery) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MarketplaceDelivery) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MarketplaceDelivery) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *MarketplaceDelivery) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEventId

`func (o *MarketplaceDelivery) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *MarketplaceDelivery) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *MarketplaceDelivery) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *MarketplaceDelivery) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventType

`func (o *MarketplaceDelivery) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MarketplaceDelivery) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MarketplaceDelivery) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MarketplaceDelivery) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetIsSimulated

`func (o *MarketplaceDelivery) GetIsSimulated() bool`

GetIsSimulated returns the IsSimulated field if non-nil, zero value otherwise.

### GetIsSimulatedOk

`func (o *MarketplaceDelivery) GetIsSimulatedOk() (*bool, bool)`

GetIsSimulatedOk returns a tuple with the IsSimulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSimulated

`func (o *MarketplaceDelivery) SetIsSimulated(v bool)`

SetIsSimulated sets IsSimulated field to given value.

### HasIsSimulated

`func (o *MarketplaceDelivery) HasIsSimulated() bool`

HasIsSimulated returns a boolean if a field has been set.

### GetMarketplaceContractId

`func (o *MarketplaceDelivery) GetMarketplaceContractId() string`

GetMarketplaceContractId returns the MarketplaceContractId field if non-nil, zero value otherwise.

### GetMarketplaceContractIdOk

`func (o *MarketplaceDelivery) GetMarketplaceContractIdOk() (*string, bool)`

GetMarketplaceContractIdOk returns a tuple with the MarketplaceContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceContractId

`func (o *MarketplaceDelivery) SetMarketplaceContractId(v string)`

SetMarketplaceContractId sets MarketplaceContractId field to given value.

### HasMarketplaceContractId

`func (o *MarketplaceDelivery) HasMarketplaceContractId() bool`

HasMarketplaceContractId returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *MarketplaceDelivery) GetMaxAttempts() int64`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *MarketplaceDelivery) GetMaxAttemptsOk() (*int64, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *MarketplaceDelivery) SetMaxAttempts(v int64)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *MarketplaceDelivery) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetMethod

`func (o *MarketplaceDelivery) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MarketplaceDelivery) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MarketplaceDelivery) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *MarketplaceDelivery) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNextAttemptAt

`func (o *MarketplaceDelivery) GetNextAttemptAt() time.Time`

GetNextAttemptAt returns the NextAttemptAt field if non-nil, zero value otherwise.

### GetNextAttemptAtOk

`func (o *MarketplaceDelivery) GetNextAttemptAtOk() (*time.Time, bool)`

GetNextAttemptAtOk returns a tuple with the NextAttemptAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttemptAt

`func (o *MarketplaceDelivery) SetNextAttemptAt(v time.Time)`

SetNextAttemptAt sets NextAttemptAt field to given value.

### HasNextAttemptAt

`func (o *MarketplaceDelivery) HasNextAttemptAt() bool`

HasNextAttemptAt returns a boolean if a field has been set.

### GetPath

`func (o *MarketplaceDelivery) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MarketplaceDelivery) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MarketplaceDelivery) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MarketplaceDelivery) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReceivedAt

`func (o *MarketplaceDelivery) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *MarketplaceDelivery) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *MarketplaceDelivery) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *MarketplaceDelivery) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### GetRefusalReason

`func (o *MarketplaceDelivery) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *MarketplaceDelivery) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *MarketplaceDelivery) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *MarketplaceDelivery) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetSentAt

`func (o *MarketplaceDelivery) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *MarketplaceDelivery) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *MarketplaceDelivery) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *MarketplaceDelivery) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetStateChange

`func (o *MarketplaceDelivery) GetStateChange() MarketplaceInteractionStateChange`

GetStateChange returns the StateChange field if non-nil, zero value otherwise.

### GetStateChangeOk

`func (o *MarketplaceDelivery) GetStateChangeOk() (*MarketplaceInteractionStateChange, bool)`

GetStateChangeOk returns a tuple with the StateChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChange

`func (o *MarketplaceDelivery) SetStateChange(v MarketplaceInteractionStateChange)`

SetStateChange sets StateChange field to given value.

### HasStateChange

`func (o *MarketplaceDelivery) HasStateChange() bool`

HasStateChange returns a boolean if a field has been set.

### GetStatus

`func (o *MarketplaceDelivery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceDelivery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceDelivery) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubscriptionId

`func (o *MarketplaceDelivery) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *MarketplaceDelivery) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *MarketplaceDelivery) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *MarketplaceDelivery) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTerminalReason

`func (o *MarketplaceDelivery) GetTerminalReason() string`

GetTerminalReason returns the TerminalReason field if non-nil, zero value otherwise.

### GetTerminalReasonOk

`func (o *MarketplaceDelivery) GetTerminalReasonOk() (*string, bool)`

GetTerminalReasonOk returns a tuple with the TerminalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalReason

`func (o *MarketplaceDelivery) SetTerminalReason(v string)`

SetTerminalReason sets TerminalReason field to given value.

### HasTerminalReason

`func (o *MarketplaceDelivery) HasTerminalReason() bool`

HasTerminalReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


