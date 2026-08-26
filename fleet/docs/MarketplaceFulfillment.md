# MarketplaceFulfillment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedOn** | Pointer to **string** | What the workflow is waiting for, in words. This is the field that answers the only question an operator actually has about a stuck contract | [optional] 
**CurrentStage** | Pointer to **string** | One stage of the fulfillment flow, in the order the workflow runs them. Creating the subscription request and committing the contract are separate stages because they are separate calls with separate failure modes, and an operator debugging a stuck contract has to know which did not happen. CONTRACT_HELD precedes SUBSCRIPTION_REQUEST_CREATED because the hold is taken before the handoff: holding is what stops the buyer being charged while the vendor onboards them, and channels that cannot hold a contract skip it | [optional] 
**FulfillmentState** | **string** | What Omnistrate decided. Deployments and metering are allowed if and only if this is READY, plus exactly one final metering window during DEPROVISIONING | 
**LastRetryAt** | Pointer to **time.Time** |  | [optional] 
**MarketplaceContractId** | **string** |  | 
**RetryCount** | Pointer to **int64** | Operator requested retries of the whole fulfillment, distinct from the per-stage retryCount and from Port B delivery attempts | [optional] 
**RunId** | Pointer to **string** | The current Temporal run. Changes on ContinueAsNew, which a long-lived fulfillment does periodically to bound history growth | [optional] 
**Stages** | [**[]MarketplaceFulfillmentStage**](MarketplaceFulfillmentStage.md) | Every stage, in order, whether or not it has been reached | 
**SubscriptionId** | Pointer to **string** | The Omnistrate subscription, once it exists. The confirm waits for the fulfillment run to produce it and returns it, so there is no second call and no webhook to wait for. A repeat confirm returns the same id rather than approving twice | [optional] 
**WorkflowId** | Pointer to **string** | The Temporal workflow id, so an operator can open the same run in the Temporal UI rather than reasoning from this summary | [optional] 
**WorkflowStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewMarketplaceFulfillment

`func NewMarketplaceFulfillment(fulfillmentState string, marketplaceContractId string, stages []MarketplaceFulfillmentStage, ) *MarketplaceFulfillment`

NewMarketplaceFulfillment instantiates a new MarketplaceFulfillment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceFulfillmentWithDefaults

`func NewMarketplaceFulfillmentWithDefaults() *MarketplaceFulfillment`

NewMarketplaceFulfillmentWithDefaults instantiates a new MarketplaceFulfillment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedOn

`func (o *MarketplaceFulfillment) GetBlockedOn() string`

GetBlockedOn returns the BlockedOn field if non-nil, zero value otherwise.

### GetBlockedOnOk

`func (o *MarketplaceFulfillment) GetBlockedOnOk() (*string, bool)`

GetBlockedOnOk returns a tuple with the BlockedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedOn

`func (o *MarketplaceFulfillment) SetBlockedOn(v string)`

SetBlockedOn sets BlockedOn field to given value.

### HasBlockedOn

`func (o *MarketplaceFulfillment) HasBlockedOn() bool`

HasBlockedOn returns a boolean if a field has been set.

### GetCurrentStage

`func (o *MarketplaceFulfillment) GetCurrentStage() string`

GetCurrentStage returns the CurrentStage field if non-nil, zero value otherwise.

### GetCurrentStageOk

`func (o *MarketplaceFulfillment) GetCurrentStageOk() (*string, bool)`

GetCurrentStageOk returns a tuple with the CurrentStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStage

`func (o *MarketplaceFulfillment) SetCurrentStage(v string)`

SetCurrentStage sets CurrentStage field to given value.

### HasCurrentStage

`func (o *MarketplaceFulfillment) HasCurrentStage() bool`

HasCurrentStage returns a boolean if a field has been set.

### GetFulfillmentState

`func (o *MarketplaceFulfillment) GetFulfillmentState() string`

GetFulfillmentState returns the FulfillmentState field if non-nil, zero value otherwise.

### GetFulfillmentStateOk

`func (o *MarketplaceFulfillment) GetFulfillmentStateOk() (*string, bool)`

GetFulfillmentStateOk returns a tuple with the FulfillmentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentState

`func (o *MarketplaceFulfillment) SetFulfillmentState(v string)`

SetFulfillmentState sets FulfillmentState field to given value.


### GetLastRetryAt

`func (o *MarketplaceFulfillment) GetLastRetryAt() time.Time`

GetLastRetryAt returns the LastRetryAt field if non-nil, zero value otherwise.

### GetLastRetryAtOk

`func (o *MarketplaceFulfillment) GetLastRetryAtOk() (*time.Time, bool)`

GetLastRetryAtOk returns a tuple with the LastRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRetryAt

`func (o *MarketplaceFulfillment) SetLastRetryAt(v time.Time)`

SetLastRetryAt sets LastRetryAt field to given value.

### HasLastRetryAt

`func (o *MarketplaceFulfillment) HasLastRetryAt() bool`

HasLastRetryAt returns a boolean if a field has been set.

### GetMarketplaceContractId

`func (o *MarketplaceFulfillment) GetMarketplaceContractId() string`

GetMarketplaceContractId returns the MarketplaceContractId field if non-nil, zero value otherwise.

### GetMarketplaceContractIdOk

`func (o *MarketplaceFulfillment) GetMarketplaceContractIdOk() (*string, bool)`

GetMarketplaceContractIdOk returns a tuple with the MarketplaceContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceContractId

`func (o *MarketplaceFulfillment) SetMarketplaceContractId(v string)`

SetMarketplaceContractId sets MarketplaceContractId field to given value.


### GetRetryCount

`func (o *MarketplaceFulfillment) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *MarketplaceFulfillment) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *MarketplaceFulfillment) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *MarketplaceFulfillment) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRunId

`func (o *MarketplaceFulfillment) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *MarketplaceFulfillment) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *MarketplaceFulfillment) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *MarketplaceFulfillment) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetStages

`func (o *MarketplaceFulfillment) GetStages() []MarketplaceFulfillmentStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *MarketplaceFulfillment) GetStagesOk() (*[]MarketplaceFulfillmentStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *MarketplaceFulfillment) SetStages(v []MarketplaceFulfillmentStage)`

SetStages sets Stages field to given value.


### GetSubscriptionId

`func (o *MarketplaceFulfillment) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *MarketplaceFulfillment) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *MarketplaceFulfillment) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *MarketplaceFulfillment) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *MarketplaceFulfillment) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *MarketplaceFulfillment) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *MarketplaceFulfillment) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *MarketplaceFulfillment) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *MarketplaceFulfillment) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *MarketplaceFulfillment) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *MarketplaceFulfillment) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *MarketplaceFulfillment) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


