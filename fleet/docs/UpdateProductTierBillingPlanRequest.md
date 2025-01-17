# UpdateProductTierBillingPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Allow creates when payment not configured | [optional] 
**Id** | **string** | ID of a Product Tier Billing Plan | 
**MaxNumberofInstances** | Pointer to **int64** | Maximum number of instances | [optional] 
**PlanName** | Pointer to **string** | Plan name | [optional] 
**Pricing** | Pointer to **interface{}** | Pricing in dollars. | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateProductTierBillingPlanRequest

`func NewUpdateProductTierBillingPlanRequest(id string, productTierId string, serviceId string, token string, ) *UpdateProductTierBillingPlanRequest`

NewUpdateProductTierBillingPlanRequest instantiates a new UpdateProductTierBillingPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductTierBillingPlanRequestWithDefaults

`func NewUpdateProductTierBillingPlanRequestWithDefaults() *UpdateProductTierBillingPlanRequest`

NewUpdateProductTierBillingPlanRequestWithDefaults instantiates a new UpdateProductTierBillingPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *UpdateProductTierBillingPlanRequest) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *UpdateProductTierBillingPlanRequest) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *UpdateProductTierBillingPlanRequest) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *UpdateProductTierBillingPlanRequest) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

### GetId

`func (o *UpdateProductTierBillingPlanRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateProductTierBillingPlanRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateProductTierBillingPlanRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMaxNumberofInstances

`func (o *UpdateProductTierBillingPlanRequest) GetMaxNumberofInstances() int64`

GetMaxNumberofInstances returns the MaxNumberofInstances field if non-nil, zero value otherwise.

### GetMaxNumberofInstancesOk

`func (o *UpdateProductTierBillingPlanRequest) GetMaxNumberofInstancesOk() (*int64, bool)`

GetMaxNumberofInstancesOk returns a tuple with the MaxNumberofInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofInstances

`func (o *UpdateProductTierBillingPlanRequest) SetMaxNumberofInstances(v int64)`

SetMaxNumberofInstances sets MaxNumberofInstances field to given value.

### HasMaxNumberofInstances

`func (o *UpdateProductTierBillingPlanRequest) HasMaxNumberofInstances() bool`

HasMaxNumberofInstances returns a boolean if a field has been set.

### GetPlanName

`func (o *UpdateProductTierBillingPlanRequest) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UpdateProductTierBillingPlanRequest) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UpdateProductTierBillingPlanRequest) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UpdateProductTierBillingPlanRequest) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPricing

`func (o *UpdateProductTierBillingPlanRequest) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *UpdateProductTierBillingPlanRequest) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *UpdateProductTierBillingPlanRequest) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *UpdateProductTierBillingPlanRequest) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *UpdateProductTierBillingPlanRequest) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *UpdateProductTierBillingPlanRequest) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetProductTierId

`func (o *UpdateProductTierBillingPlanRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *UpdateProductTierBillingPlanRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *UpdateProductTierBillingPlanRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *UpdateProductTierBillingPlanRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateProductTierBillingPlanRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateProductTierBillingPlanRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *UpdateProductTierBillingPlanRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateProductTierBillingPlanRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateProductTierBillingPlanRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


