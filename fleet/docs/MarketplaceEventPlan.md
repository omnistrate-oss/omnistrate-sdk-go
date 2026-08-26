# MarketplaceEventPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitAmount** | Pointer to **string** | Committed amount as a decimal STRING, never a number. A float cannot hold a currency amount exactly, and this value is compared against what the channel invoiced | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**PlanRef** | **string** | The channel&#39;s identifier for the purchased plan | 
**Quantity** | Pointer to **int64** | Seat or unit count where the channel reports one | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMarketplaceEventPlan

`func NewMarketplaceEventPlan(planRef string, ) *MarketplaceEventPlan`

NewMarketplaceEventPlan instantiates a new MarketplaceEventPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceEventPlanWithDefaults

`func NewMarketplaceEventPlanWithDefaults() *MarketplaceEventPlan`

NewMarketplaceEventPlanWithDefaults instantiates a new MarketplaceEventPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitAmount

`func (o *MarketplaceEventPlan) GetCommitAmount() string`

GetCommitAmount returns the CommitAmount field if non-nil, zero value otherwise.

### GetCommitAmountOk

`func (o *MarketplaceEventPlan) GetCommitAmountOk() (*string, bool)`

GetCommitAmountOk returns a tuple with the CommitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitAmount

`func (o *MarketplaceEventPlan) SetCommitAmount(v string)`

SetCommitAmount sets CommitAmount field to given value.

### HasCommitAmount

`func (o *MarketplaceEventPlan) HasCommitAmount() bool`

HasCommitAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *MarketplaceEventPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MarketplaceEventPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MarketplaceEventPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MarketplaceEventPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEndsAt

`func (o *MarketplaceEventPlan) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *MarketplaceEventPlan) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *MarketplaceEventPlan) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *MarketplaceEventPlan) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetPlanRef

`func (o *MarketplaceEventPlan) GetPlanRef() string`

GetPlanRef returns the PlanRef field if non-nil, zero value otherwise.

### GetPlanRefOk

`func (o *MarketplaceEventPlan) GetPlanRefOk() (*string, bool)`

GetPlanRefOk returns a tuple with the PlanRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRef

`func (o *MarketplaceEventPlan) SetPlanRef(v string)`

SetPlanRef sets PlanRef field to given value.


### GetQuantity

`func (o *MarketplaceEventPlan) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MarketplaceEventPlan) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MarketplaceEventPlan) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MarketplaceEventPlan) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartsAt

`func (o *MarketplaceEventPlan) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *MarketplaceEventPlan) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *MarketplaceEventPlan) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *MarketplaceEventPlan) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


