# FleetCreateSubscriptionOnBehalfOfCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InheritServicePlanPrice** | Pointer to **bool** | Whether to inherit the service plan price | [optional] 
**OnBehalfOfCustomerUserId** | **string** | ID of a User | 
**PriceEffectiveDate** | Pointer to **string** | The effective date of the price, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
**PricePerUnit** | Pointer to **map[string]interface{}** | If inheritServicePlanPrice is false, provide the price per unit for the subscription here. | [optional] 
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


### GetInheritServicePlanPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetInheritServicePlanPrice() bool`

GetInheritServicePlanPrice returns the InheritServicePlanPrice field if non-nil, zero value otherwise.

### GetInheritServicePlanPriceOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetInheritServicePlanPriceOk() (*bool, bool)`

GetInheritServicePlanPriceOk returns a tuple with the InheritServicePlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritServicePlanPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetInheritServicePlanPrice(v bool)`

SetInheritServicePlanPrice sets InheritServicePlanPrice field to given value.

### HasInheritServicePlanPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasInheritServicePlanPrice() bool`

HasInheritServicePlanPrice returns a boolean if a field has been set.

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


### GetPriceEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetPriceEffectiveDate() string`

GetPriceEffectiveDate returns the PriceEffectiveDate field if non-nil, zero value otherwise.

### GetPriceEffectiveDateOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetPriceEffectiveDateOk() (*string, bool)`

GetPriceEffectiveDateOk returns a tuple with the PriceEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetPriceEffectiveDate(v string)`

SetPriceEffectiveDate sets PriceEffectiveDate field to given value.

### HasPriceEffectiveDate

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasPriceEffectiveDate() bool`

HasPriceEffectiveDate returns a boolean if a field has been set.

### GetPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetPricePerUnit() map[string]interface{}`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetPricePerUnitOk() (*map[string]interface{}, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetPricePerUnit(v map[string]interface{})`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

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


