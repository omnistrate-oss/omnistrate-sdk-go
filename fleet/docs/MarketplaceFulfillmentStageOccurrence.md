# MarketplaceFulfillmentStageOccurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **string** | Server computed. Absent while the occurrence is still running | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **string** | Why this occurrence failed | [optional] 
**Inputs** | Pointer to [**[]MarketplaceStageParameter**](MarketplaceStageParameter.md) |  | [optional] 
**IsStalled** | Pointer to **bool** | Open but no longer moving. Kept apart from a failure because nothing rejected it: it is waiting, past its SLA, and the fix is different | [optional] 
**Legs** | Pointer to [**[]MarketplaceStageLeg**](MarketplaceStageLeg.md) | The calls that took this occurrence from pending to complete, oldest first | [optional] 
**OccurrenceId** | **string** | Stable identifier for this occurrence, unique within the contract. Two suspensions of one contract have different values here and the same stage | 
**Outputs** | Pointer to [**[]MarketplaceStageParameter**](MarketplaceStageParameter.md) |  | [optional] 
**RetryCount** | Pointer to **int64** | Attempts after the first, within this occurrence | [optional] 
**Stage** | **string** | One stage of the fulfillment flow, in the order the workflow runs them. Creating the subscription request and committing the contract are separate stages because they are separate calls with separate failure modes, and an operator debugging a stuck contract has to know which did not happen. CONTRACT_HELD precedes SUBSCRIPTION_REQUEST_CREATED because the hold is taken before the handoff: holding is what stops the buyer being charged while the vendor onboards them, and channels that cannot hold a contract skip it | 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewMarketplaceFulfillmentStageOccurrence

`func NewMarketplaceFulfillmentStageOccurrence(occurrenceId string, stage string, status string, ) *MarketplaceFulfillmentStageOccurrence`

NewMarketplaceFulfillmentStageOccurrence instantiates a new MarketplaceFulfillmentStageOccurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceFulfillmentStageOccurrenceWithDefaults

`func NewMarketplaceFulfillmentStageOccurrenceWithDefaults() *MarketplaceFulfillmentStageOccurrence`

NewMarketplaceFulfillmentStageOccurrenceWithDefaults instantiates a new MarketplaceFulfillmentStageOccurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *MarketplaceFulfillmentStageOccurrence) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MarketplaceFulfillmentStageOccurrence) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MarketplaceFulfillmentStageOccurrence) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndedAt

`func (o *MarketplaceFulfillmentStageOccurrence) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *MarketplaceFulfillmentStageOccurrence) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *MarketplaceFulfillmentStageOccurrence) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetError

`func (o *MarketplaceFulfillmentStageOccurrence) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MarketplaceFulfillmentStageOccurrence) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MarketplaceFulfillmentStageOccurrence) HasError() bool`

HasError returns a boolean if a field has been set.

### GetInputs

`func (o *MarketplaceFulfillmentStageOccurrence) GetInputs() []MarketplaceStageParameter`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetInputsOk() (*[]MarketplaceStageParameter, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *MarketplaceFulfillmentStageOccurrence) SetInputs(v []MarketplaceStageParameter)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *MarketplaceFulfillmentStageOccurrence) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetIsStalled

`func (o *MarketplaceFulfillmentStageOccurrence) GetIsStalled() bool`

GetIsStalled returns the IsStalled field if non-nil, zero value otherwise.

### GetIsStalledOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetIsStalledOk() (*bool, bool)`

GetIsStalledOk returns a tuple with the IsStalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStalled

`func (o *MarketplaceFulfillmentStageOccurrence) SetIsStalled(v bool)`

SetIsStalled sets IsStalled field to given value.

### HasIsStalled

`func (o *MarketplaceFulfillmentStageOccurrence) HasIsStalled() bool`

HasIsStalled returns a boolean if a field has been set.

### GetLegs

`func (o *MarketplaceFulfillmentStageOccurrence) GetLegs() []MarketplaceStageLeg`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetLegsOk() (*[]MarketplaceStageLeg, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *MarketplaceFulfillmentStageOccurrence) SetLegs(v []MarketplaceStageLeg)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *MarketplaceFulfillmentStageOccurrence) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetOccurrenceId

`func (o *MarketplaceFulfillmentStageOccurrence) GetOccurrenceId() string`

GetOccurrenceId returns the OccurrenceId field if non-nil, zero value otherwise.

### GetOccurrenceIdOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetOccurrenceIdOk() (*string, bool)`

GetOccurrenceIdOk returns a tuple with the OccurrenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceId

`func (o *MarketplaceFulfillmentStageOccurrence) SetOccurrenceId(v string)`

SetOccurrenceId sets OccurrenceId field to given value.


### GetOutputs

`func (o *MarketplaceFulfillmentStageOccurrence) GetOutputs() []MarketplaceStageParameter`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetOutputsOk() (*[]MarketplaceStageParameter, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *MarketplaceFulfillmentStageOccurrence) SetOutputs(v []MarketplaceStageParameter)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *MarketplaceFulfillmentStageOccurrence) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetRetryCount

`func (o *MarketplaceFulfillmentStageOccurrence) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *MarketplaceFulfillmentStageOccurrence) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *MarketplaceFulfillmentStageOccurrence) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetStage

`func (o *MarketplaceFulfillmentStageOccurrence) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *MarketplaceFulfillmentStageOccurrence) SetStage(v string)`

SetStage sets Stage field to given value.


### GetStartedAt

`func (o *MarketplaceFulfillmentStageOccurrence) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MarketplaceFulfillmentStageOccurrence) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MarketplaceFulfillmentStageOccurrence) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *MarketplaceFulfillmentStageOccurrence) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceFulfillmentStageOccurrence) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceFulfillmentStageOccurrence) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


