# FleetUpdateSubscriptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Whether to allow creating instances when payment is not configured. | [optional] 
**BillingProvider** | Pointer to **string** | The billing provider type | [optional] 
**CustomPrice** | Pointer to **bool** | Whether to use a custom price for this subscription | [optional] 
**CustomPricePerUnit** | Pointer to **map[string]interface{}** | If customPrice is true, provide the price per unit for the subscription here. | [optional] 
**EnvironmentId** | **string** | ID of a Service Environment | 
**ExternalPayerId** | Pointer to **string** | The external payer ID to record which customer should pay for this subscription. | [optional] 
**MaxNumberOfInstances** | Pointer to **int64** | The maximum number of instances that can be created for this subscription. Use -1 to unset this restriction. | [optional] 
**PriceEffectiveDate** | Pointer to **string** | The effective date of the price, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
**ServiceId** | **string** | ID of a Service | 
**SubscriptionIDs** | **[]string** | List of subscription IDs to update | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetUpdateSubscriptionsRequest

`func NewFleetUpdateSubscriptionsRequest(environmentId string, serviceId string, subscriptionIDs []string, token string, ) *FleetUpdateSubscriptionsRequest`

NewFleetUpdateSubscriptionsRequest instantiates a new FleetUpdateSubscriptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateSubscriptionsRequestWithDefaults

`func NewFleetUpdateSubscriptionsRequestWithDefaults() *FleetUpdateSubscriptionsRequest`

NewFleetUpdateSubscriptionsRequestWithDefaults instantiates a new FleetUpdateSubscriptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *FleetUpdateSubscriptionsRequest) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *FleetUpdateSubscriptionsRequest) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *FleetUpdateSubscriptionsRequest) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *FleetUpdateSubscriptionsRequest) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

### GetBillingProvider

`func (o *FleetUpdateSubscriptionsRequest) GetBillingProvider() string`

GetBillingProvider returns the BillingProvider field if non-nil, zero value otherwise.

### GetBillingProviderOk

`func (o *FleetUpdateSubscriptionsRequest) GetBillingProviderOk() (*string, bool)`

GetBillingProviderOk returns a tuple with the BillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProvider

`func (o *FleetUpdateSubscriptionsRequest) SetBillingProvider(v string)`

SetBillingProvider sets BillingProvider field to given value.

### HasBillingProvider

`func (o *FleetUpdateSubscriptionsRequest) HasBillingProvider() bool`

HasBillingProvider returns a boolean if a field has been set.

### GetCustomPrice

`func (o *FleetUpdateSubscriptionsRequest) GetCustomPrice() bool`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *FleetUpdateSubscriptionsRequest) GetCustomPriceOk() (*bool, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *FleetUpdateSubscriptionsRequest) SetCustomPrice(v bool)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *FleetUpdateSubscriptionsRequest) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### GetCustomPricePerUnit

`func (o *FleetUpdateSubscriptionsRequest) GetCustomPricePerUnit() map[string]interface{}`

GetCustomPricePerUnit returns the CustomPricePerUnit field if non-nil, zero value otherwise.

### GetCustomPricePerUnitOk

`func (o *FleetUpdateSubscriptionsRequest) GetCustomPricePerUnitOk() (*map[string]interface{}, bool)`

GetCustomPricePerUnitOk returns a tuple with the CustomPricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPricePerUnit

`func (o *FleetUpdateSubscriptionsRequest) SetCustomPricePerUnit(v map[string]interface{})`

SetCustomPricePerUnit sets CustomPricePerUnit field to given value.

### HasCustomPricePerUnit

`func (o *FleetUpdateSubscriptionsRequest) HasCustomPricePerUnit() bool`

HasCustomPricePerUnit returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *FleetUpdateSubscriptionsRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetUpdateSubscriptionsRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetUpdateSubscriptionsRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExternalPayerId

`func (o *FleetUpdateSubscriptionsRequest) GetExternalPayerId() string`

GetExternalPayerId returns the ExternalPayerId field if non-nil, zero value otherwise.

### GetExternalPayerIdOk

`func (o *FleetUpdateSubscriptionsRequest) GetExternalPayerIdOk() (*string, bool)`

GetExternalPayerIdOk returns a tuple with the ExternalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayerId

`func (o *FleetUpdateSubscriptionsRequest) SetExternalPayerId(v string)`

SetExternalPayerId sets ExternalPayerId field to given value.

### HasExternalPayerId

`func (o *FleetUpdateSubscriptionsRequest) HasExternalPayerId() bool`

HasExternalPayerId returns a boolean if a field has been set.

### GetMaxNumberOfInstances

`func (o *FleetUpdateSubscriptionsRequest) GetMaxNumberOfInstances() int64`

GetMaxNumberOfInstances returns the MaxNumberOfInstances field if non-nil, zero value otherwise.

### GetMaxNumberOfInstancesOk

`func (o *FleetUpdateSubscriptionsRequest) GetMaxNumberOfInstancesOk() (*int64, bool)`

GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfInstances

`func (o *FleetUpdateSubscriptionsRequest) SetMaxNumberOfInstances(v int64)`

SetMaxNumberOfInstances sets MaxNumberOfInstances field to given value.

### HasMaxNumberOfInstances

`func (o *FleetUpdateSubscriptionsRequest) HasMaxNumberOfInstances() bool`

HasMaxNumberOfInstances returns a boolean if a field has been set.

### GetPriceEffectiveDate

`func (o *FleetUpdateSubscriptionsRequest) GetPriceEffectiveDate() string`

GetPriceEffectiveDate returns the PriceEffectiveDate field if non-nil, zero value otherwise.

### GetPriceEffectiveDateOk

`func (o *FleetUpdateSubscriptionsRequest) GetPriceEffectiveDateOk() (*string, bool)`

GetPriceEffectiveDateOk returns a tuple with the PriceEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEffectiveDate

`func (o *FleetUpdateSubscriptionsRequest) SetPriceEffectiveDate(v string)`

SetPriceEffectiveDate sets PriceEffectiveDate field to given value.

### HasPriceEffectiveDate

`func (o *FleetUpdateSubscriptionsRequest) HasPriceEffectiveDate() bool`

HasPriceEffectiveDate returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetUpdateSubscriptionsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetUpdateSubscriptionsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetUpdateSubscriptionsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSubscriptionIDs

`func (o *FleetUpdateSubscriptionsRequest) GetSubscriptionIDs() []string`

GetSubscriptionIDs returns the SubscriptionIDs field if non-nil, zero value otherwise.

### GetSubscriptionIDsOk

`func (o *FleetUpdateSubscriptionsRequest) GetSubscriptionIDsOk() (*[]string, bool)`

GetSubscriptionIDsOk returns a tuple with the SubscriptionIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIDs

`func (o *FleetUpdateSubscriptionsRequest) SetSubscriptionIDs(v []string)`

SetSubscriptionIDs sets SubscriptionIDs field to given value.


### GetToken

`func (o *FleetUpdateSubscriptionsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetUpdateSubscriptionsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetUpdateSubscriptionsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


