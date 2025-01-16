# CreateProductTierBillingPlanRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | **bool** | Allow creates when payment not configured | 
**MaxNumberofInstances** | **int64** | Maximum number of instances | 
**PlanName** | **string** | Plan name | 
**Pricing** | **interface{}** | Pricing in dollars. | 

## Methods

### NewCreateProductTierBillingPlanRequest2

`func NewCreateProductTierBillingPlanRequest2(allowCreatesWhenPaymentNotConfigured bool, maxNumberofInstances int64, planName string, pricing interface{}, ) *CreateProductTierBillingPlanRequest2`

NewCreateProductTierBillingPlanRequest2 instantiates a new CreateProductTierBillingPlanRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductTierBillingPlanRequest2WithDefaults

`func NewCreateProductTierBillingPlanRequest2WithDefaults() *CreateProductTierBillingPlanRequest2`

NewCreateProductTierBillingPlanRequest2WithDefaults instantiates a new CreateProductTierBillingPlanRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *CreateProductTierBillingPlanRequest2) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *CreateProductTierBillingPlanRequest2) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *CreateProductTierBillingPlanRequest2) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.


### GetMaxNumberofInstances

`func (o *CreateProductTierBillingPlanRequest2) GetMaxNumberofInstances() int64`

GetMaxNumberofInstances returns the MaxNumberofInstances field if non-nil, zero value otherwise.

### GetMaxNumberofInstancesOk

`func (o *CreateProductTierBillingPlanRequest2) GetMaxNumberofInstancesOk() (*int64, bool)`

GetMaxNumberofInstancesOk returns a tuple with the MaxNumberofInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofInstances

`func (o *CreateProductTierBillingPlanRequest2) SetMaxNumberofInstances(v int64)`

SetMaxNumberofInstances sets MaxNumberofInstances field to given value.


### GetPlanName

`func (o *CreateProductTierBillingPlanRequest2) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *CreateProductTierBillingPlanRequest2) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *CreateProductTierBillingPlanRequest2) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetPricing

`func (o *CreateProductTierBillingPlanRequest2) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *CreateProductTierBillingPlanRequest2) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *CreateProductTierBillingPlanRequest2) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.


### SetPricingNil

`func (o *CreateProductTierBillingPlanRequest2) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *CreateProductTierBillingPlanRequest2) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


