# UpdateProductTierBillingPlanRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Allow creates when payment not configured | [optional] 
**MaxNumberofInstances** | Pointer to **int64** | Maximum number of instances | [optional] 
**PlanName** | Pointer to **string** | Plan name | [optional] 
**Pricing** | Pointer to **interface{}** | Pricing in dollars. | [optional] 

## Methods

### NewUpdateProductTierBillingPlanRequest2

`func NewUpdateProductTierBillingPlanRequest2() *UpdateProductTierBillingPlanRequest2`

NewUpdateProductTierBillingPlanRequest2 instantiates a new UpdateProductTierBillingPlanRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductTierBillingPlanRequest2WithDefaults

`func NewUpdateProductTierBillingPlanRequest2WithDefaults() *UpdateProductTierBillingPlanRequest2`

NewUpdateProductTierBillingPlanRequest2WithDefaults instantiates a new UpdateProductTierBillingPlanRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *UpdateProductTierBillingPlanRequest2) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *UpdateProductTierBillingPlanRequest2) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *UpdateProductTierBillingPlanRequest2) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *UpdateProductTierBillingPlanRequest2) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

### GetMaxNumberofInstances

`func (o *UpdateProductTierBillingPlanRequest2) GetMaxNumberofInstances() int64`

GetMaxNumberofInstances returns the MaxNumberofInstances field if non-nil, zero value otherwise.

### GetMaxNumberofInstancesOk

`func (o *UpdateProductTierBillingPlanRequest2) GetMaxNumberofInstancesOk() (*int64, bool)`

GetMaxNumberofInstancesOk returns a tuple with the MaxNumberofInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofInstances

`func (o *UpdateProductTierBillingPlanRequest2) SetMaxNumberofInstances(v int64)`

SetMaxNumberofInstances sets MaxNumberofInstances field to given value.

### HasMaxNumberofInstances

`func (o *UpdateProductTierBillingPlanRequest2) HasMaxNumberofInstances() bool`

HasMaxNumberofInstances returns a boolean if a field has been set.

### GetPlanName

`func (o *UpdateProductTierBillingPlanRequest2) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UpdateProductTierBillingPlanRequest2) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UpdateProductTierBillingPlanRequest2) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UpdateProductTierBillingPlanRequest2) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPricing

`func (o *UpdateProductTierBillingPlanRequest2) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *UpdateProductTierBillingPlanRequest2) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *UpdateProductTierBillingPlanRequest2) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *UpdateProductTierBillingPlanRequest2) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *UpdateProductTierBillingPlanRequest2) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *UpdateProductTierBillingPlanRequest2) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


