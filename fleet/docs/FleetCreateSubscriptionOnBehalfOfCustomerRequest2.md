# FleetCreateSubscriptionOnBehalfOfCustomerRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InheritServicePlanPricing** | Pointer to **bool** | Whether to inherit the service plan pricing | [optional] 
**OnBehalfOfCustomerUserId** | **string** | The user ID of the customer that this subscription is on behalf of | 
**PricingEffectiveDate** | Pointer to **string** | The effective date of the pricing, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
**PricingPerUnit** | Pointer to **map[string]interface{}** | If inheritServicePlanPricing is false, provide the pricing per unit for the subscription here. | [optional] 
**ProductTierId** | **string** | The product tier ID | 

## Methods

### NewFleetCreateSubscriptionOnBehalfOfCustomerRequest2

`func NewFleetCreateSubscriptionOnBehalfOfCustomerRequest2(onBehalfOfCustomerUserId string, productTierId string, ) *FleetCreateSubscriptionOnBehalfOfCustomerRequest2`

NewFleetCreateSubscriptionOnBehalfOfCustomerRequest2 instantiates a new FleetCreateSubscriptionOnBehalfOfCustomerRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCreateSubscriptionOnBehalfOfCustomerRequest2WithDefaults

`func NewFleetCreateSubscriptionOnBehalfOfCustomerRequest2WithDefaults() *FleetCreateSubscriptionOnBehalfOfCustomerRequest2`

NewFleetCreateSubscriptionOnBehalfOfCustomerRequest2WithDefaults instantiates a new FleetCreateSubscriptionOnBehalfOfCustomerRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInheritServicePlanPricing

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetInheritServicePlanPricing() bool`

GetInheritServicePlanPricing returns the InheritServicePlanPricing field if non-nil, zero value otherwise.

### GetInheritServicePlanPricingOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetInheritServicePlanPricingOk() (*bool, bool)`

GetInheritServicePlanPricingOk returns a tuple with the InheritServicePlanPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritServicePlanPricing

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetInheritServicePlanPricing(v bool)`

SetInheritServicePlanPricing sets InheritServicePlanPricing field to given value.

### HasInheritServicePlanPricing

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasInheritServicePlanPricing() bool`

HasInheritServicePlanPricing returns a boolean if a field has been set.

### GetOnBehalfOfCustomerUserId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetOnBehalfOfCustomerUserId() string`

GetOnBehalfOfCustomerUserId returns the OnBehalfOfCustomerUserId field if non-nil, zero value otherwise.

### GetOnBehalfOfCustomerUserIdOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetOnBehalfOfCustomerUserIdOk() (*string, bool)`

GetOnBehalfOfCustomerUserIdOk returns a tuple with the OnBehalfOfCustomerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBehalfOfCustomerUserId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetOnBehalfOfCustomerUserId(v string)`

SetOnBehalfOfCustomerUserId sets OnBehalfOfCustomerUserId field to given value.


### GetPricingEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetPricingEffectiveDate() string`

GetPricingEffectiveDate returns the PricingEffectiveDate field if non-nil, zero value otherwise.

### GetPricingEffectiveDateOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetPricingEffectiveDateOk() (*string, bool)`

GetPricingEffectiveDateOk returns a tuple with the PricingEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetPricingEffectiveDate(v string)`

SetPricingEffectiveDate sets PricingEffectiveDate field to given value.

### HasPricingEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasPricingEffectiveDate() bool`

HasPricingEffectiveDate returns a boolean if a field has been set.

### GetPricingPerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetPricingPerUnit() map[string]interface{}`

GetPricingPerUnit returns the PricingPerUnit field if non-nil, zero value otherwise.

### GetPricingPerUnitOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetPricingPerUnitOk() (*map[string]interface{}, bool)`

GetPricingPerUnitOk returns a tuple with the PricingPerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetPricingPerUnit(v map[string]interface{})`

SetPricingPerUnit sets PricingPerUnit field to given value.

### HasPricingPerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasPricingPerUnit() bool`

HasPricingPerUnit returns a boolean if a field has been set.

### GetProductTierId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


