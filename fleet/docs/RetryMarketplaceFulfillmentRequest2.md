# RetryMarketplaceFulfillmentRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Recorded on the audit event. Worth filling in, because a retry is a manual intervention on a contract that is usually already billing | [optional] 
**Stage** | Pointer to **string** | Which stage to re-enter. Defaults to the stage the workflow is stalled or failed on, which is what an operator wants in almost every case | [optional] 

## Methods

### NewRetryMarketplaceFulfillmentRequest2

`func NewRetryMarketplaceFulfillmentRequest2() *RetryMarketplaceFulfillmentRequest2`

NewRetryMarketplaceFulfillmentRequest2 instantiates a new RetryMarketplaceFulfillmentRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryMarketplaceFulfillmentRequest2WithDefaults

`func NewRetryMarketplaceFulfillmentRequest2WithDefaults() *RetryMarketplaceFulfillmentRequest2`

NewRetryMarketplaceFulfillmentRequest2WithDefaults instantiates a new RetryMarketplaceFulfillmentRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *RetryMarketplaceFulfillmentRequest2) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RetryMarketplaceFulfillmentRequest2) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RetryMarketplaceFulfillmentRequest2) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RetryMarketplaceFulfillmentRequest2) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStage

`func (o *RetryMarketplaceFulfillmentRequest2) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *RetryMarketplaceFulfillmentRequest2) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *RetryMarketplaceFulfillmentRequest2) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *RetryMarketplaceFulfillmentRequest2) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


