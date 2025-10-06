# FleetGetUsageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | Pointer to [**[]UsagePerSubscription**](UsagePerSubscription.md) | Usage broken down per subscription | [optional] 

## Methods

### NewFleetGetUsageResult

`func NewFleetGetUsageResult() *FleetGetUsageResult`

NewFleetGetUsageResult instantiates a new FleetGetUsageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetGetUsageResultWithDefaults

`func NewFleetGetUsageResultWithDefaults() *FleetGetUsageResult`

NewFleetGetUsageResultWithDefaults instantiates a new FleetGetUsageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *FleetGetUsageResult) GetSubscriptions() []UsagePerSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *FleetGetUsageResult) GetSubscriptionsOk() (*[]UsagePerSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *FleetGetUsageResult) SetSubscriptions(v []UsagePerSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *FleetGetUsageResult) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


