# FleetUpdateSubscriptionsRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Whether to allow creating instances when payment is not configured. | [optional] 
**BillingProvider** | Pointer to **string** | The billing provider for this subscription | [optional] 
**CustomPrice** | Pointer to **bool** | Whether to use a custom price for this subscription | [optional] 
**CustomPricePerUnit** | Pointer to **map[string]interface{}** | If customPrice is true, provide the price per unit for the subscription here. | [optional] 
**ExternalPayerId** | Pointer to **string** | The external payer ID to record which customer should pay for this subscription. | [optional] 
**MaxNumberOfInstances** | Pointer to **int64** | The maximum number of instances that can be created for this subscription. Use -1 to unset this restriction. | [optional] 
**PriceEffectiveDate** | Pointer to **string** | The effective date of the price, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
**SubscriptionIDs** | **[]string** | List of subscription IDs to update | 

## Methods

### NewFleetUpdateSubscriptionsRequest2

`func NewFleetUpdateSubscriptionsRequest2(subscriptionIDs []string, ) *FleetUpdateSubscriptionsRequest2`

NewFleetUpdateSubscriptionsRequest2 instantiates a new FleetUpdateSubscriptionsRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateSubscriptionsRequest2WithDefaults

`func NewFleetUpdateSubscriptionsRequest2WithDefaults() *FleetUpdateSubscriptionsRequest2`

NewFleetUpdateSubscriptionsRequest2WithDefaults instantiates a new FleetUpdateSubscriptionsRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *FleetUpdateSubscriptionsRequest2) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *FleetUpdateSubscriptionsRequest2) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *FleetUpdateSubscriptionsRequest2) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *FleetUpdateSubscriptionsRequest2) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

### GetBillingProvider

`func (o *FleetUpdateSubscriptionsRequest2) GetBillingProvider() string`

GetBillingProvider returns the BillingProvider field if non-nil, zero value otherwise.

### GetBillingProviderOk

`func (o *FleetUpdateSubscriptionsRequest2) GetBillingProviderOk() (*string, bool)`

GetBillingProviderOk returns a tuple with the BillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProvider

`func (o *FleetUpdateSubscriptionsRequest2) SetBillingProvider(v string)`

SetBillingProvider sets BillingProvider field to given value.

### HasBillingProvider

`func (o *FleetUpdateSubscriptionsRequest2) HasBillingProvider() bool`

HasBillingProvider returns a boolean if a field has been set.

### GetCustomPrice

`func (o *FleetUpdateSubscriptionsRequest2) GetCustomPrice() bool`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *FleetUpdateSubscriptionsRequest2) GetCustomPriceOk() (*bool, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *FleetUpdateSubscriptionsRequest2) SetCustomPrice(v bool)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *FleetUpdateSubscriptionsRequest2) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### GetCustomPricePerUnit

`func (o *FleetUpdateSubscriptionsRequest2) GetCustomPricePerUnit() map[string]interface{}`

GetCustomPricePerUnit returns the CustomPricePerUnit field if non-nil, zero value otherwise.

### GetCustomPricePerUnitOk

`func (o *FleetUpdateSubscriptionsRequest2) GetCustomPricePerUnitOk() (*map[string]interface{}, bool)`

GetCustomPricePerUnitOk returns a tuple with the CustomPricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPricePerUnit

`func (o *FleetUpdateSubscriptionsRequest2) SetCustomPricePerUnit(v map[string]interface{})`

SetCustomPricePerUnit sets CustomPricePerUnit field to given value.

### HasCustomPricePerUnit

`func (o *FleetUpdateSubscriptionsRequest2) HasCustomPricePerUnit() bool`

HasCustomPricePerUnit returns a boolean if a field has been set.

### GetExternalPayerId

`func (o *FleetUpdateSubscriptionsRequest2) GetExternalPayerId() string`

GetExternalPayerId returns the ExternalPayerId field if non-nil, zero value otherwise.

### GetExternalPayerIdOk

`func (o *FleetUpdateSubscriptionsRequest2) GetExternalPayerIdOk() (*string, bool)`

GetExternalPayerIdOk returns a tuple with the ExternalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayerId

`func (o *FleetUpdateSubscriptionsRequest2) SetExternalPayerId(v string)`

SetExternalPayerId sets ExternalPayerId field to given value.

### HasExternalPayerId

`func (o *FleetUpdateSubscriptionsRequest2) HasExternalPayerId() bool`

HasExternalPayerId returns a boolean if a field has been set.

### GetMaxNumberOfInstances

`func (o *FleetUpdateSubscriptionsRequest2) GetMaxNumberOfInstances() int64`

GetMaxNumberOfInstances returns the MaxNumberOfInstances field if non-nil, zero value otherwise.

### GetMaxNumberOfInstancesOk

`func (o *FleetUpdateSubscriptionsRequest2) GetMaxNumberOfInstancesOk() (*int64, bool)`

GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfInstances

`func (o *FleetUpdateSubscriptionsRequest2) SetMaxNumberOfInstances(v int64)`

SetMaxNumberOfInstances sets MaxNumberOfInstances field to given value.

### HasMaxNumberOfInstances

`func (o *FleetUpdateSubscriptionsRequest2) HasMaxNumberOfInstances() bool`

HasMaxNumberOfInstances returns a boolean if a field has been set.

### GetPriceEffectiveDate

`func (o *FleetUpdateSubscriptionsRequest2) GetPriceEffectiveDate() string`

GetPriceEffectiveDate returns the PriceEffectiveDate field if non-nil, zero value otherwise.

### GetPriceEffectiveDateOk

`func (o *FleetUpdateSubscriptionsRequest2) GetPriceEffectiveDateOk() (*string, bool)`

GetPriceEffectiveDateOk returns a tuple with the PriceEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEffectiveDate

`func (o *FleetUpdateSubscriptionsRequest2) SetPriceEffectiveDate(v string)`

SetPriceEffectiveDate sets PriceEffectiveDate field to given value.

### HasPriceEffectiveDate

`func (o *FleetUpdateSubscriptionsRequest2) HasPriceEffectiveDate() bool`

HasPriceEffectiveDate returns a boolean if a field has been set.

### GetSubscriptionIDs

`func (o *FleetUpdateSubscriptionsRequest2) GetSubscriptionIDs() []string`

GetSubscriptionIDs returns the SubscriptionIDs field if non-nil, zero value otherwise.

### GetSubscriptionIDsOk

`func (o *FleetUpdateSubscriptionsRequest2) GetSubscriptionIDsOk() (*[]string, bool)`

GetSubscriptionIDsOk returns a tuple with the SubscriptionIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIDs

`func (o *FleetUpdateSubscriptionsRequest2) SetSubscriptionIDs(v []string)`

SetSubscriptionIDs sets SubscriptionIDs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


