# BillingPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | **bool** | Allow creates when payment not configured | 
**Id** | **string** | Product tier billing plan ID | 
**MaxNumberofInstances** | **int64** | Maximum number of instances | 
**PlanName** | **string** | Plan name | 
**Pricing** | **interface{}** | Pricing in dollars. | 
**ProductTierId** | **string** | Product tier ID | 
**ServiceId** | **string** | Service ID | 

## Methods

### NewBillingPlan

`func NewBillingPlan(allowCreatesWhenPaymentNotConfigured bool, id string, maxNumberofInstances int64, planName string, pricing interface{}, productTierId string, serviceId string, ) *BillingPlan`

NewBillingPlan instantiates a new BillingPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPlanWithDefaults

`func NewBillingPlanWithDefaults() *BillingPlan`

NewBillingPlanWithDefaults instantiates a new BillingPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *BillingPlan) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *BillingPlan) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *BillingPlan) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.


### GetId

`func (o *BillingPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingPlan) SetId(v string)`

SetId sets Id field to given value.


### GetMaxNumberofInstances

`func (o *BillingPlan) GetMaxNumberofInstances() int64`

GetMaxNumberofInstances returns the MaxNumberofInstances field if non-nil, zero value otherwise.

### GetMaxNumberofInstancesOk

`func (o *BillingPlan) GetMaxNumberofInstancesOk() (*int64, bool)`

GetMaxNumberofInstancesOk returns a tuple with the MaxNumberofInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofInstances

`func (o *BillingPlan) SetMaxNumberofInstances(v int64)`

SetMaxNumberofInstances sets MaxNumberofInstances field to given value.


### GetPlanName

`func (o *BillingPlan) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *BillingPlan) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *BillingPlan) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetPricing

`func (o *BillingPlan) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *BillingPlan) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *BillingPlan) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.


### SetPricingNil

`func (o *BillingPlan) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *BillingPlan) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetProductTierId

`func (o *BillingPlan) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *BillingPlan) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *BillingPlan) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *BillingPlan) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *BillingPlan) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *BillingPlan) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


