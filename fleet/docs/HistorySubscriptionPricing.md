# HistorySubscriptionPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The time that this pricing was created | [optional] 
**CreatedByUserId** | Pointer to **string** | ID of a User | [optional] 
**CreatedByUserName** | Pointer to **string** | The name of the user that created the price | [optional] 
**EndDate** | Pointer to **string** | The end date of the price | [optional] 
**InheritServicePlanPrice** | Pointer to **bool** | Whether this price inherits the service plan price | [optional] 
**PricePerUnit** | Pointer to **map[string]interface{}** | The price per unit for the subscription | [optional] 
**StartDate** | Pointer to **string** | The start date of the price | [optional] 

## Methods

### NewHistorySubscriptionPricing

`func NewHistorySubscriptionPricing() *HistorySubscriptionPricing`

NewHistorySubscriptionPricing instantiates a new HistorySubscriptionPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistorySubscriptionPricingWithDefaults

`func NewHistorySubscriptionPricingWithDefaults() *HistorySubscriptionPricing`

NewHistorySubscriptionPricingWithDefaults instantiates a new HistorySubscriptionPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HistorySubscriptionPricing) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HistorySubscriptionPricing) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HistorySubscriptionPricing) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HistorySubscriptionPricing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *HistorySubscriptionPricing) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *HistorySubscriptionPricing) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *HistorySubscriptionPricing) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *HistorySubscriptionPricing) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedByUserName

`func (o *HistorySubscriptionPricing) GetCreatedByUserName() string`

GetCreatedByUserName returns the CreatedByUserName field if non-nil, zero value otherwise.

### GetCreatedByUserNameOk

`func (o *HistorySubscriptionPricing) GetCreatedByUserNameOk() (*string, bool)`

GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserName

`func (o *HistorySubscriptionPricing) SetCreatedByUserName(v string)`

SetCreatedByUserName sets CreatedByUserName field to given value.

### HasCreatedByUserName

`func (o *HistorySubscriptionPricing) HasCreatedByUserName() bool`

HasCreatedByUserName returns a boolean if a field has been set.

### GetEndDate

`func (o *HistorySubscriptionPricing) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *HistorySubscriptionPricing) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *HistorySubscriptionPricing) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *HistorySubscriptionPricing) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInheritServicePlanPrice

`func (o *HistorySubscriptionPricing) GetInheritServicePlanPrice() bool`

GetInheritServicePlanPrice returns the InheritServicePlanPrice field if non-nil, zero value otherwise.

### GetInheritServicePlanPriceOk

`func (o *HistorySubscriptionPricing) GetInheritServicePlanPriceOk() (*bool, bool)`

GetInheritServicePlanPriceOk returns a tuple with the InheritServicePlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritServicePlanPrice

`func (o *HistorySubscriptionPricing) SetInheritServicePlanPrice(v bool)`

SetInheritServicePlanPrice sets InheritServicePlanPrice field to given value.

### HasInheritServicePlanPrice

`func (o *HistorySubscriptionPricing) HasInheritServicePlanPrice() bool`

HasInheritServicePlanPrice returns a boolean if a field has been set.

### GetPricePerUnit

`func (o *HistorySubscriptionPricing) GetPricePerUnit() map[string]interface{}`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *HistorySubscriptionPricing) GetPricePerUnitOk() (*map[string]interface{}, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *HistorySubscriptionPricing) SetPricePerUnit(v map[string]interface{})`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *HistorySubscriptionPricing) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

### GetStartDate

`func (o *HistorySubscriptionPricing) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *HistorySubscriptionPricing) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *HistorySubscriptionPricing) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *HistorySubscriptionPricing) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


