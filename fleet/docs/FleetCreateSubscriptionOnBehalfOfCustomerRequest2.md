# FleetCreateSubscriptionOnBehalfOfCustomerRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Whether to allow creating instances when payment is not configured. | [optional] 
**BillingProvider** | Pointer to **string** | The billing provider for this subscription | [optional] 
**CustomPrice** | Pointer to **bool** | Whether to use a custom price for this subscription | [optional] 
**CustomPricePerUnit** | Pointer to **map[string]interface{}** | If customPrice is true, provide the price per unit for the subscription here. | [optional] 
**ExternalPayerId** | Pointer to **string** | The external payer ID to record which customer should pay for this subscription | [optional] 
**MaxNumberOfInstances** | Pointer to **int64** | The maximum number of instances that can be created for this subscription. Use -1 to unset this restriction. | [optional] 
**OnBehalfOfCustomerUserId** | **string** | The user ID of the customer that this subscription is on behalf of | 
**PriceEffectiveDate** | Pointer to **string** | The effective date of the price, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
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

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

### GetBillingProvider

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetBillingProvider() string`

GetBillingProvider returns the BillingProvider field if non-nil, zero value otherwise.

### GetBillingProviderOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetBillingProviderOk() (*string, bool)`

GetBillingProviderOk returns a tuple with the BillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProvider

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetBillingProvider(v string)`

SetBillingProvider sets BillingProvider field to given value.

### HasBillingProvider

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasBillingProvider() bool`

HasBillingProvider returns a boolean if a field has been set.

### GetCustomPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetCustomPrice() bool`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetCustomPriceOk() (*bool, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetCustomPrice(v bool)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### GetCustomPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetCustomPricePerUnit() map[string]interface{}`

GetCustomPricePerUnit returns the CustomPricePerUnit field if non-nil, zero value otherwise.

### GetCustomPricePerUnitOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetCustomPricePerUnitOk() (*map[string]interface{}, bool)`

GetCustomPricePerUnitOk returns a tuple with the CustomPricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetCustomPricePerUnit(v map[string]interface{})`

SetCustomPricePerUnit sets CustomPricePerUnit field to given value.

### HasCustomPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasCustomPricePerUnit() bool`

HasCustomPricePerUnit returns a boolean if a field has been set.

### GetExternalPayerId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetExternalPayerId() string`

GetExternalPayerId returns the ExternalPayerId field if non-nil, zero value otherwise.

### GetExternalPayerIdOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetExternalPayerIdOk() (*string, bool)`

GetExternalPayerIdOk returns a tuple with the ExternalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayerId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetExternalPayerId(v string)`

SetExternalPayerId sets ExternalPayerId field to given value.

### HasExternalPayerId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasExternalPayerId() bool`

HasExternalPayerId returns a boolean if a field has been set.

### GetMaxNumberOfInstances

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetMaxNumberOfInstances() int64`

GetMaxNumberOfInstances returns the MaxNumberOfInstances field if non-nil, zero value otherwise.

### GetMaxNumberOfInstancesOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) GetMaxNumberOfInstancesOk() (*int64, bool)`

GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfInstances

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) SetMaxNumberOfInstances(v int64)`

SetMaxNumberOfInstances sets MaxNumberOfInstances field to given value.

### HasMaxNumberOfInstances

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest2) HasMaxNumberOfInstances() bool`

HasMaxNumberOfInstances returns a boolean if a field has been set.

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


