# FleetUpdateSubscriptionRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InheritServicePlanPricing** | Pointer to **bool** | Whether to inherit the service plan pricing | [optional] 
**PricingEffectiveDate** | Pointer to **string** | The effective date of the pricing, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
**PricingPerUnit** | Pointer to **map[string]interface{}** | If inheritServicePlanPricing is false, provide the pricing per unit for the subscription here. | [optional] 

## Methods

### NewFleetUpdateSubscriptionRequest2

`func NewFleetUpdateSubscriptionRequest2() *FleetUpdateSubscriptionRequest2`

NewFleetUpdateSubscriptionRequest2 instantiates a new FleetUpdateSubscriptionRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateSubscriptionRequest2WithDefaults

`func NewFleetUpdateSubscriptionRequest2WithDefaults() *FleetUpdateSubscriptionRequest2`

NewFleetUpdateSubscriptionRequest2WithDefaults instantiates a new FleetUpdateSubscriptionRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInheritServicePlanPricing

`func (o *FleetUpdateSubscriptionRequest2) GetInheritServicePlanPricing() bool`

GetInheritServicePlanPricing returns the InheritServicePlanPricing field if non-nil, zero value otherwise.

### GetInheritServicePlanPricingOk

`func (o *FleetUpdateSubscriptionRequest2) GetInheritServicePlanPricingOk() (*bool, bool)`

GetInheritServicePlanPricingOk returns a tuple with the InheritServicePlanPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritServicePlanPricing

`func (o *FleetUpdateSubscriptionRequest2) SetInheritServicePlanPricing(v bool)`

SetInheritServicePlanPricing sets InheritServicePlanPricing field to given value.

### HasInheritServicePlanPricing

`func (o *FleetUpdateSubscriptionRequest2) HasInheritServicePlanPricing() bool`

HasInheritServicePlanPricing returns a boolean if a field has been set.

### GetPricingEffectiveDate

`func (o *FleetUpdateSubscriptionRequest2) GetPricingEffectiveDate() string`

GetPricingEffectiveDate returns the PricingEffectiveDate field if non-nil, zero value otherwise.

### GetPricingEffectiveDateOk

`func (o *FleetUpdateSubscriptionRequest2) GetPricingEffectiveDateOk() (*string, bool)`

GetPricingEffectiveDateOk returns a tuple with the PricingEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingEffectiveDate

`func (o *FleetUpdateSubscriptionRequest2) SetPricingEffectiveDate(v string)`

SetPricingEffectiveDate sets PricingEffectiveDate field to given value.

### HasPricingEffectiveDate

`func (o *FleetUpdateSubscriptionRequest2) HasPricingEffectiveDate() bool`

HasPricingEffectiveDate returns a boolean if a field has been set.

### GetPricingPerUnit

`func (o *FleetUpdateSubscriptionRequest2) GetPricingPerUnit() map[string]interface{}`

GetPricingPerUnit returns the PricingPerUnit field if non-nil, zero value otherwise.

### GetPricingPerUnitOk

`func (o *FleetUpdateSubscriptionRequest2) GetPricingPerUnitOk() (*map[string]interface{}, bool)`

GetPricingPerUnitOk returns a tuple with the PricingPerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPerUnit

`func (o *FleetUpdateSubscriptionRequest2) SetPricingPerUnit(v map[string]interface{})`

SetPricingPerUnit sets PricingPerUnit field to given value.

### HasPricingPerUnit

`func (o *FleetUpdateSubscriptionRequest2) HasPricingPerUnit() bool`

HasPricingPerUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


