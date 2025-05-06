# FleetUpdateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**Id** | **string** | ID of a Subscription | 
**InheritServicePlanPricing** | Pointer to **bool** | Whether to inherit the service plan pricing | [optional] 
**PricingEffectiveDate** | Pointer to **string** | The effective date of the pricing, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
**PricingPerUnit** | Pointer to **map[string]interface{}** | If inheritServicePlanPricing is false, provide the pricing per unit for the subscription here. | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetUpdateSubscriptionRequest

`func NewFleetUpdateSubscriptionRequest(environmentId string, id string, serviceId string, token string, ) *FleetUpdateSubscriptionRequest`

NewFleetUpdateSubscriptionRequest instantiates a new FleetUpdateSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateSubscriptionRequestWithDefaults

`func NewFleetUpdateSubscriptionRequestWithDefaults() *FleetUpdateSubscriptionRequest`

NewFleetUpdateSubscriptionRequestWithDefaults instantiates a new FleetUpdateSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetUpdateSubscriptionRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetUpdateSubscriptionRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetUpdateSubscriptionRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *FleetUpdateSubscriptionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetUpdateSubscriptionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetUpdateSubscriptionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetInheritServicePlanPricing

`func (o *FleetUpdateSubscriptionRequest) GetInheritServicePlanPricing() bool`

GetInheritServicePlanPricing returns the InheritServicePlanPricing field if non-nil, zero value otherwise.

### GetInheritServicePlanPricingOk

`func (o *FleetUpdateSubscriptionRequest) GetInheritServicePlanPricingOk() (*bool, bool)`

GetInheritServicePlanPricingOk returns a tuple with the InheritServicePlanPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritServicePlanPricing

`func (o *FleetUpdateSubscriptionRequest) SetInheritServicePlanPricing(v bool)`

SetInheritServicePlanPricing sets InheritServicePlanPricing field to given value.

### HasInheritServicePlanPricing

`func (o *FleetUpdateSubscriptionRequest) HasInheritServicePlanPricing() bool`

HasInheritServicePlanPricing returns a boolean if a field has been set.

### GetPricingEffectiveDate

`func (o *FleetUpdateSubscriptionRequest) GetPricingEffectiveDate() string`

GetPricingEffectiveDate returns the PricingEffectiveDate field if non-nil, zero value otherwise.

### GetPricingEffectiveDateOk

`func (o *FleetUpdateSubscriptionRequest) GetPricingEffectiveDateOk() (*string, bool)`

GetPricingEffectiveDateOk returns a tuple with the PricingEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingEffectiveDate

`func (o *FleetUpdateSubscriptionRequest) SetPricingEffectiveDate(v string)`

SetPricingEffectiveDate sets PricingEffectiveDate field to given value.

### HasPricingEffectiveDate

`func (o *FleetUpdateSubscriptionRequest) HasPricingEffectiveDate() bool`

HasPricingEffectiveDate returns a boolean if a field has been set.

### GetPricingPerUnit

`func (o *FleetUpdateSubscriptionRequest) GetPricingPerUnit() map[string]interface{}`

GetPricingPerUnit returns the PricingPerUnit field if non-nil, zero value otherwise.

### GetPricingPerUnitOk

`func (o *FleetUpdateSubscriptionRequest) GetPricingPerUnitOk() (*map[string]interface{}, bool)`

GetPricingPerUnitOk returns a tuple with the PricingPerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPerUnit

`func (o *FleetUpdateSubscriptionRequest) SetPricingPerUnit(v map[string]interface{})`

SetPricingPerUnit sets PricingPerUnit field to given value.

### HasPricingPerUnit

`func (o *FleetUpdateSubscriptionRequest) HasPricingPerUnit() bool`

HasPricingPerUnit returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetUpdateSubscriptionRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetUpdateSubscriptionRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetUpdateSubscriptionRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetUpdateSubscriptionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetUpdateSubscriptionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetUpdateSubscriptionRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


