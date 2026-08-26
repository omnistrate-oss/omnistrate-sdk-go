# MarketplaceFulfillmentStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **string** | Server computed so the UI never reasons about clock skew. Absent while the stage is still running | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **string** | Why the stage failed | [optional] 
**EventIds** | Pointer to **[]string** | Referenced by id rather than duplicated, so the stage view and the contract event log can never disagree | [optional] 
**Inputs** | Pointer to [**[]MarketplaceStageParameter**](MarketplaceStageParameter.md) |  | [optional] 
**IsStalled** | Pointer to **bool** | The stage is open but has stopped moving, for example a handoff accepted and never answered. Kept apart from a failure because nothing rejected it: it is waiting, past its SLA, and the fix is different | [optional] 
**Outputs** | Pointer to [**[]MarketplaceStageParameter**](MarketplaceStageParameter.md) |  | [optional] 
**RetryCount** | Pointer to **int64** | Attempts after the first. Absent and zero both mean never retried | [optional] 
**Stage** | **string** | One stage of the fulfillment flow, in the order the workflow runs them. Creating the subscription request and committing the contract are separate stages because they are separate calls with separate failure modes, and an operator debugging a stuck contract has to know which did not happen. CONTRACT_HELD precedes SUBSCRIPTION_REQUEST_CREATED because the hold is taken before the handoff: holding is what stops the buyer being charged while the vendor onboards them, and channels that cannot hold a contract skip it | 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewMarketplaceFulfillmentStage

`func NewMarketplaceFulfillmentStage(stage string, status string, ) *MarketplaceFulfillmentStage`

NewMarketplaceFulfillmentStage instantiates a new MarketplaceFulfillmentStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceFulfillmentStageWithDefaults

`func NewMarketplaceFulfillmentStageWithDefaults() *MarketplaceFulfillmentStage`

NewMarketplaceFulfillmentStageWithDefaults instantiates a new MarketplaceFulfillmentStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *MarketplaceFulfillmentStage) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MarketplaceFulfillmentStage) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MarketplaceFulfillmentStage) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MarketplaceFulfillmentStage) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndedAt

`func (o *MarketplaceFulfillmentStage) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *MarketplaceFulfillmentStage) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *MarketplaceFulfillmentStage) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *MarketplaceFulfillmentStage) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetError

`func (o *MarketplaceFulfillmentStage) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MarketplaceFulfillmentStage) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MarketplaceFulfillmentStage) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MarketplaceFulfillmentStage) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEventIds

`func (o *MarketplaceFulfillmentStage) GetEventIds() []string`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *MarketplaceFulfillmentStage) GetEventIdsOk() (*[]string, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *MarketplaceFulfillmentStage) SetEventIds(v []string)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *MarketplaceFulfillmentStage) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.

### GetInputs

`func (o *MarketplaceFulfillmentStage) GetInputs() []MarketplaceStageParameter`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *MarketplaceFulfillmentStage) GetInputsOk() (*[]MarketplaceStageParameter, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *MarketplaceFulfillmentStage) SetInputs(v []MarketplaceStageParameter)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *MarketplaceFulfillmentStage) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetIsStalled

`func (o *MarketplaceFulfillmentStage) GetIsStalled() bool`

GetIsStalled returns the IsStalled field if non-nil, zero value otherwise.

### GetIsStalledOk

`func (o *MarketplaceFulfillmentStage) GetIsStalledOk() (*bool, bool)`

GetIsStalledOk returns a tuple with the IsStalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStalled

`func (o *MarketplaceFulfillmentStage) SetIsStalled(v bool)`

SetIsStalled sets IsStalled field to given value.

### HasIsStalled

`func (o *MarketplaceFulfillmentStage) HasIsStalled() bool`

HasIsStalled returns a boolean if a field has been set.

### GetOutputs

`func (o *MarketplaceFulfillmentStage) GetOutputs() []MarketplaceStageParameter`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *MarketplaceFulfillmentStage) GetOutputsOk() (*[]MarketplaceStageParameter, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *MarketplaceFulfillmentStage) SetOutputs(v []MarketplaceStageParameter)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *MarketplaceFulfillmentStage) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetRetryCount

`func (o *MarketplaceFulfillmentStage) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *MarketplaceFulfillmentStage) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *MarketplaceFulfillmentStage) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *MarketplaceFulfillmentStage) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetStage

`func (o *MarketplaceFulfillmentStage) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *MarketplaceFulfillmentStage) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *MarketplaceFulfillmentStage) SetStage(v string)`

SetStage sets Stage field to given value.


### GetStartedAt

`func (o *MarketplaceFulfillmentStage) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MarketplaceFulfillmentStage) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MarketplaceFulfillmentStage) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MarketplaceFulfillmentStage) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *MarketplaceFulfillmentStage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceFulfillmentStage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceFulfillmentStage) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


