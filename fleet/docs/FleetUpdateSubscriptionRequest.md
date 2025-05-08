# FleetUpdateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**Id** | **string** | ID of a Subscription | 
**InheritServicePlanPrice** | Pointer to **bool** | Whether to inherit the service plan price | [optional] 
**PriceEffectiveDate** | Pointer to **string** | The effective date of the price, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
**PricePerUnit** | Pointer to **map[string]interface{}** | If inheritServicePlanPrice is false, provide the price per unit for the subscription here. | [optional] 
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


### GetInheritServicePlanPrice

`func (o *FleetUpdateSubscriptionRequest) GetInheritServicePlanPrice() bool`

GetInheritServicePlanPrice returns the InheritServicePlanPrice field if non-nil, zero value otherwise.

### GetInheritServicePlanPriceOk

`func (o *FleetUpdateSubscriptionRequest) GetInheritServicePlanPriceOk() (*bool, bool)`

GetInheritServicePlanPriceOk returns a tuple with the InheritServicePlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritServicePlanPrice

`func (o *FleetUpdateSubscriptionRequest) SetInheritServicePlanPrice(v bool)`

SetInheritServicePlanPrice sets InheritServicePlanPrice field to given value.

### HasInheritServicePlanPrice

`func (o *FleetUpdateSubscriptionRequest) HasInheritServicePlanPrice() bool`

HasInheritServicePlanPrice returns a boolean if a field has been set.

### GetPriceEffectiveDate

`func (o *FleetUpdateSubscriptionRequest) GetPriceEffectiveDate() string`

GetPriceEffectiveDate returns the PriceEffectiveDate field if non-nil, zero value otherwise.

### GetPriceEffectiveDateOk

`func (o *FleetUpdateSubscriptionRequest) GetPriceEffectiveDateOk() (*string, bool)`

GetPriceEffectiveDateOk returns a tuple with the PriceEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEffectiveDate

`func (o *FleetUpdateSubscriptionRequest) SetPriceEffectiveDate(v string)`

SetPriceEffectiveDate sets PriceEffectiveDate field to given value.

### HasPriceEffectiveDate

`func (o *FleetUpdateSubscriptionRequest) HasPriceEffectiveDate() bool`

HasPriceEffectiveDate returns a boolean if a field has been set.

### GetPricePerUnit

`func (o *FleetUpdateSubscriptionRequest) GetPricePerUnit() map[string]interface{}`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *FleetUpdateSubscriptionRequest) GetPricePerUnitOk() (*map[string]interface{}, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *FleetUpdateSubscriptionRequest) SetPricePerUnit(v map[string]interface{})`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *FleetUpdateSubscriptionRequest) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

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


