# FleetCreateSubscriptionOnBehalfOfCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InheritServicePlanPricing** | Pointer to **bool** | Whether to inherit the service plan pricing | [optional] 
**OnBehalfOfCustomerUserId** | **string** | ID of a User | 
**PricingEffectiveDate** | Pointer to **string** | The effective date of the pricing, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
**PricingPerUnit** | Pointer to **map[string]interface{}** | If inheritServicePlanPricing is false, provide the pricing per unit for the subscription here. | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetCreateSubscriptionOnBehalfOfCustomerRequest

`func NewFleetCreateSubscriptionOnBehalfOfCustomerRequest(environmentId string, onBehalfOfCustomerUserId string, productTierId string, serviceId string, token string, ) *FleetCreateSubscriptionOnBehalfOfCustomerRequest`

NewFleetCreateSubscriptionOnBehalfOfCustomerRequest instantiates a new FleetCreateSubscriptionOnBehalfOfCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCreateSubscriptionOnBehalfOfCustomerRequestWithDefaults

`func NewFleetCreateSubscriptionOnBehalfOfCustomerRequestWithDefaults() *FleetCreateSubscriptionOnBehalfOfCustomerRequest`

NewFleetCreateSubscriptionOnBehalfOfCustomerRequestWithDefaults instantiates a new FleetCreateSubscriptionOnBehalfOfCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInheritServicePlanPricing

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetInheritServicePlanPricing() bool`

GetInheritServicePlanPricing returns the InheritServicePlanPricing field if non-nil, zero value otherwise.

### GetInheritServicePlanPricingOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetInheritServicePlanPricingOk() (*bool, bool)`

GetInheritServicePlanPricingOk returns a tuple with the InheritServicePlanPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritServicePlanPricing

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetInheritServicePlanPricing(v bool)`

SetInheritServicePlanPricing sets InheritServicePlanPricing field to given value.

### HasInheritServicePlanPricing

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasInheritServicePlanPricing() bool`

HasInheritServicePlanPricing returns a boolean if a field has been set.

### GetOnBehalfOfCustomerUserId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetOnBehalfOfCustomerUserId() string`

GetOnBehalfOfCustomerUserId returns the OnBehalfOfCustomerUserId field if non-nil, zero value otherwise.

### GetOnBehalfOfCustomerUserIdOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetOnBehalfOfCustomerUserIdOk() (*string, bool)`

GetOnBehalfOfCustomerUserIdOk returns a tuple with the OnBehalfOfCustomerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBehalfOfCustomerUserId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetOnBehalfOfCustomerUserId(v string)`

SetOnBehalfOfCustomerUserId sets OnBehalfOfCustomerUserId field to given value.


### GetPricingEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetPricingEffectiveDate() string`

GetPricingEffectiveDate returns the PricingEffectiveDate field if non-nil, zero value otherwise.

### GetPricingEffectiveDateOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetPricingEffectiveDateOk() (*string, bool)`

GetPricingEffectiveDateOk returns a tuple with the PricingEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetPricingEffectiveDate(v string)`

SetPricingEffectiveDate sets PricingEffectiveDate field to given value.

### HasPricingEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasPricingEffectiveDate() bool`

HasPricingEffectiveDate returns a boolean if a field has been set.

### GetPricingPerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetPricingPerUnit() map[string]interface{}`

GetPricingPerUnit returns the PricingPerUnit field if non-nil, zero value otherwise.

### GetPricingPerUnitOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetPricingPerUnitOk() (*map[string]interface{}, bool)`

GetPricingPerUnitOk returns a tuple with the PricingPerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetPricingPerUnit(v map[string]interface{})`

SetPricingPerUnit sets PricingPerUnit field to given value.

### HasPricingPerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasPricingPerUnit() bool`

HasPricingPerUnit returns a boolean if a field has been set.

### GetProductTierId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


