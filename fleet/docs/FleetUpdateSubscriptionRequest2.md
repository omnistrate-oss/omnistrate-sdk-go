# FleetUpdateSubscriptionRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalPayerId** | Pointer to **string** | The external payer ID to record which customer should pay for this subscription. | [optional] 
**InheritServicePlanPrice** | Pointer to **bool** | Whether to inherit the service plan price | [optional] 
**PriceEffectiveDate** | Pointer to **string** | The effective date of the price, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
**PricePerUnit** | Pointer to **map[string]interface{}** | If inheritServicePlanPrice is false, provide the price per unit for the subscription here. | [optional] 

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

### GetExternalPayerId

`func (o *FleetUpdateSubscriptionRequest2) GetExternalPayerId() string`

GetExternalPayerId returns the ExternalPayerId field if non-nil, zero value otherwise.

### GetExternalPayerIdOk

`func (o *FleetUpdateSubscriptionRequest2) GetExternalPayerIdOk() (*string, bool)`

GetExternalPayerIdOk returns a tuple with the ExternalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayerId

`func (o *FleetUpdateSubscriptionRequest2) SetExternalPayerId(v string)`

SetExternalPayerId sets ExternalPayerId field to given value.

### HasExternalPayerId

`func (o *FleetUpdateSubscriptionRequest2) HasExternalPayerId() bool`

HasExternalPayerId returns a boolean if a field has been set.

### GetInheritServicePlanPrice

`func (o *FleetUpdateSubscriptionRequest2) GetInheritServicePlanPrice() bool`

GetInheritServicePlanPrice returns the InheritServicePlanPrice field if non-nil, zero value otherwise.

### GetInheritServicePlanPriceOk

`func (o *FleetUpdateSubscriptionRequest2) GetInheritServicePlanPriceOk() (*bool, bool)`

GetInheritServicePlanPriceOk returns a tuple with the InheritServicePlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritServicePlanPrice

`func (o *FleetUpdateSubscriptionRequest2) SetInheritServicePlanPrice(v bool)`

SetInheritServicePlanPrice sets InheritServicePlanPrice field to given value.

### HasInheritServicePlanPrice

`func (o *FleetUpdateSubscriptionRequest2) HasInheritServicePlanPrice() bool`

HasInheritServicePlanPrice returns a boolean if a field has been set.

### GetPriceEffectiveDate

`func (o *FleetUpdateSubscriptionRequest2) GetPriceEffectiveDate() string`

GetPriceEffectiveDate returns the PriceEffectiveDate field if non-nil, zero value otherwise.

### GetPriceEffectiveDateOk

`func (o *FleetUpdateSubscriptionRequest2) GetPriceEffectiveDateOk() (*string, bool)`

GetPriceEffectiveDateOk returns a tuple with the PriceEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEffectiveDate

`func (o *FleetUpdateSubscriptionRequest2) SetPriceEffectiveDate(v string)`

SetPriceEffectiveDate sets PriceEffectiveDate field to given value.

### HasPriceEffectiveDate

`func (o *FleetUpdateSubscriptionRequest2) HasPriceEffectiveDate() bool`

HasPriceEffectiveDate returns a boolean if a field has been set.

### GetPricePerUnit

`func (o *FleetUpdateSubscriptionRequest2) GetPricePerUnit() map[string]interface{}`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *FleetUpdateSubscriptionRequest2) GetPricePerUnitOk() (*map[string]interface{}, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *FleetUpdateSubscriptionRequest2) SetPricePerUnit(v map[string]interface{})`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *FleetUpdateSubscriptionRequest2) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


