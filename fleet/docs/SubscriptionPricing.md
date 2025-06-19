# SubscriptionPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The time that this pricing was created | [optional] 
**CreatedByUserId** | Pointer to **string** | ID of a User | [optional] 
**CreatedByUserName** | Pointer to **string** | The name of the user that created the price | [optional] 
**CustomPrice** | Pointer to **bool** | Whether this price is a custom price | [optional] 
**EndDate** | Pointer to **string** | The end date of the price | [optional] 
**IsActiveBillingPrice** | Pointer to **bool** | Whether this price should be used for billing | [optional] 
**PricePerUnit** | Pointer to **map[string]interface{}** | The price per unit for the subscription | [optional] 
**StartDate** | Pointer to **string** | The start date of the price | [optional] 

## Methods

### NewSubscriptionPricing

`func NewSubscriptionPricing() *SubscriptionPricing`

NewSubscriptionPricing instantiates a new SubscriptionPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPricingWithDefaults

`func NewSubscriptionPricingWithDefaults() *SubscriptionPricing`

NewSubscriptionPricingWithDefaults instantiates a new SubscriptionPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SubscriptionPricing) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionPricing) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionPricing) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionPricing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *SubscriptionPricing) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *SubscriptionPricing) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *SubscriptionPricing) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *SubscriptionPricing) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedByUserName

`func (o *SubscriptionPricing) GetCreatedByUserName() string`

GetCreatedByUserName returns the CreatedByUserName field if non-nil, zero value otherwise.

### GetCreatedByUserNameOk

`func (o *SubscriptionPricing) GetCreatedByUserNameOk() (*string, bool)`

GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserName

`func (o *SubscriptionPricing) SetCreatedByUserName(v string)`

SetCreatedByUserName sets CreatedByUserName field to given value.

### HasCreatedByUserName

`func (o *SubscriptionPricing) HasCreatedByUserName() bool`

HasCreatedByUserName returns a boolean if a field has been set.

### GetCustomPrice

`func (o *SubscriptionPricing) GetCustomPrice() bool`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *SubscriptionPricing) GetCustomPriceOk() (*bool, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *SubscriptionPricing) SetCustomPrice(v bool)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *SubscriptionPricing) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### GetEndDate

`func (o *SubscriptionPricing) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SubscriptionPricing) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SubscriptionPricing) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SubscriptionPricing) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIsActiveBillingPrice

`func (o *SubscriptionPricing) GetIsActiveBillingPrice() bool`

GetIsActiveBillingPrice returns the IsActiveBillingPrice field if non-nil, zero value otherwise.

### GetIsActiveBillingPriceOk

`func (o *SubscriptionPricing) GetIsActiveBillingPriceOk() (*bool, bool)`

GetIsActiveBillingPriceOk returns a tuple with the IsActiveBillingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActiveBillingPrice

`func (o *SubscriptionPricing) SetIsActiveBillingPrice(v bool)`

SetIsActiveBillingPrice sets IsActiveBillingPrice field to given value.

### HasIsActiveBillingPrice

`func (o *SubscriptionPricing) HasIsActiveBillingPrice() bool`

HasIsActiveBillingPrice returns a boolean if a field has been set.

### GetPricePerUnit

`func (o *SubscriptionPricing) GetPricePerUnit() map[string]interface{}`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *SubscriptionPricing) GetPricePerUnitOk() (*map[string]interface{}, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *SubscriptionPricing) SetPricePerUnit(v map[string]interface{})`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *SubscriptionPricing) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

### GetStartDate

`func (o *SubscriptionPricing) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SubscriptionPricing) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SubscriptionPricing) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SubscriptionPricing) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


