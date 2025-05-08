# FleetCreateSubscriptionOnBehalfOfCustomerRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InheritServicePlanPrice** | Pointer to **bool** | Whether to inherit the service plan price | [optional] 
**OnBehalfOfCustomerUserId** | **string** | The user ID of the customer that this subscription is on behalf of | 
**PriceEffectiveDate** | Pointer to **string** | The effective date of the price, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
**PricePerUnit** | Pointer to **map[string]interface{}** | If inheritServicePlanPrice is false, provide the price per unit for the subscription here. | [optional] 
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

### GetInheritServicePlanPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetInheritServicePlanPrice() bool`

GetInheritServicePlanPrice returns the InheritServicePlanPrice field if non-nil, zero value otherwise.

### GetInheritServicePlanPriceOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetInheritServicePlanPriceOk() (*bool, bool)`

GetInheritServicePlanPriceOk returns a tuple with the InheritServicePlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritServicePlanPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetInheritServicePlanPrice(v bool)`

SetInheritServicePlanPrice sets InheritServicePlanPrice field to given value.

### HasInheritServicePlanPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasInheritServicePlanPrice() bool`

HasInheritServicePlanPrice returns a boolean if a field has been set.

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


### GetPriceEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetPriceEffectiveDate() string`

GetPriceEffectiveDate returns the PriceEffectiveDate field if non-nil, zero value otherwise.

### GetPriceEffectiveDateOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetPriceEffectiveDateOk() (*string, bool)`

GetPriceEffectiveDateOk returns a tuple with the PriceEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetPriceEffectiveDate(v string)`

SetPriceEffectiveDate sets PriceEffectiveDate field to given value.

### HasPriceEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasPriceEffectiveDate() bool`

HasPriceEffectiveDate returns a boolean if a field has been set.

### GetPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetPricePerUnit() map[string]interface{}`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetPricePerUnitOk() (*map[string]interface{}, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetPricePerUnit(v map[string]interface{})`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

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


