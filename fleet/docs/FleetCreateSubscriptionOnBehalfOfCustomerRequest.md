# FleetCreateSubscriptionOnBehalfOfCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Whether to allow creating instances when payment is not configured. | [optional] 
**BillingProvider** | Pointer to **string** | The billing provider type | [optional] 
**CustomPrice** | Pointer to **bool** | Whether to use a custom price for this subscription | [optional] 
**CustomPricePerUnit** | Pointer to **map[string]interface{}** | If customPrice is true, provide the price per unit for the subscription here. | [optional] 
**EnvironmentId** | **string** | ID of a Service Environment | 
**ExternalPayerId** | Pointer to **string** | The external payer ID to record which customer should pay for this subscription | [optional] 
**MaxNumberOfInstances** | Pointer to **int64** | The maximum number of instances that can be created for this subscription. Use -1 to unset this restriction. | [optional] 
**OnBehalfOfCustomerUserId** | **string** | ID of a User | 
**PriceEffectiveDate** | Pointer to **string** | The effective date of the price, truncated to the first day of the month. Only the current or future months may be specified. | [optional] 
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

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

### GetBillingProvider

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetBillingProvider() string`

GetBillingProvider returns the BillingProvider field if non-nil, zero value otherwise.

### GetBillingProviderOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetBillingProviderOk() (*string, bool)`

GetBillingProviderOk returns a tuple with the BillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProvider

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetBillingProvider(v string)`

SetBillingProvider sets BillingProvider field to given value.

### HasBillingProvider

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasBillingProvider() bool`

HasBillingProvider returns a boolean if a field has been set.

### GetCustomPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetCustomPrice() bool`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetCustomPriceOk() (*bool, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetCustomPrice(v bool)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### GetCustomPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetCustomPricePerUnit() map[string]interface{}`

GetCustomPricePerUnit returns the CustomPricePerUnit field if non-nil, zero value otherwise.

### GetCustomPricePerUnitOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetCustomPricePerUnitOk() (*map[string]interface{}, bool)`

GetCustomPricePerUnitOk returns a tuple with the CustomPricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetCustomPricePerUnit(v map[string]interface{})`

SetCustomPricePerUnit sets CustomPricePerUnit field to given value.

### HasCustomPricePerUnit

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasCustomPricePerUnit() bool`

HasCustomPricePerUnit returns a boolean if a field has been set.

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


### GetExternalPayerId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetExternalPayerId() string`

GetExternalPayerId returns the ExternalPayerId field if non-nil, zero value otherwise.

### GetExternalPayerIdOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetExternalPayerIdOk() (*string, bool)`

GetExternalPayerIdOk returns a tuple with the ExternalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayerId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetExternalPayerId(v string)`

SetExternalPayerId sets ExternalPayerId field to given value.

### HasExternalPayerId

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasExternalPayerId() bool`

HasExternalPayerId returns a boolean if a field has been set.

### GetMaxNumberOfInstances

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetMaxNumberOfInstances() int64`

GetMaxNumberOfInstances returns the MaxNumberOfInstances field if non-nil, zero value otherwise.

### GetMaxNumberOfInstancesOk

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) GetMaxNumberOfInstancesOk() (*int64, bool)`

GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfInstances

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) SetMaxNumberOfInstances(v int64)`

SetMaxNumberOfInstances sets MaxNumberOfInstances field to given value.

### HasMaxNumberOfInstances

`func (o *FleetCreateSubscriptionOnBehalfOfCustomerRequest) HasMaxNumberOfInstances() bool`

HasMaxNumberOfInstances returns a boolean if a field has been set.

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


