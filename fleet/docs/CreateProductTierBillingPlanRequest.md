# CreateProductTierBillingPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | **bool** | Allow creates when payment not configured | 
**Id** | **string** | ID of a Product Tier | 
**MaxNumberofInstances** | **int64** | Maximum number of instances | 
**PlanName** | **string** | Plan name | 
**Pricing** | **interface{}** | Pricing in dollars. | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateProductTierBillingPlanRequest

`func NewCreateProductTierBillingPlanRequest(allowCreatesWhenPaymentNotConfigured bool, id string, maxNumberofInstances int64, planName string, pricing interface{}, serviceId string, token string, ) *CreateProductTierBillingPlanRequest`

NewCreateProductTierBillingPlanRequest instantiates a new CreateProductTierBillingPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductTierBillingPlanRequestWithDefaults

`func NewCreateProductTierBillingPlanRequestWithDefaults() *CreateProductTierBillingPlanRequest`

NewCreateProductTierBillingPlanRequestWithDefaults instantiates a new CreateProductTierBillingPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *CreateProductTierBillingPlanRequest) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *CreateProductTierBillingPlanRequest) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *CreateProductTierBillingPlanRequest) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.


### GetId

`func (o *CreateProductTierBillingPlanRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateProductTierBillingPlanRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateProductTierBillingPlanRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMaxNumberofInstances

`func (o *CreateProductTierBillingPlanRequest) GetMaxNumberofInstances() int64`

GetMaxNumberofInstances returns the MaxNumberofInstances field if non-nil, zero value otherwise.

### GetMaxNumberofInstancesOk

`func (o *CreateProductTierBillingPlanRequest) GetMaxNumberofInstancesOk() (*int64, bool)`

GetMaxNumberofInstancesOk returns a tuple with the MaxNumberofInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofInstances

`func (o *CreateProductTierBillingPlanRequest) SetMaxNumberofInstances(v int64)`

SetMaxNumberofInstances sets MaxNumberofInstances field to given value.


### GetPlanName

`func (o *CreateProductTierBillingPlanRequest) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *CreateProductTierBillingPlanRequest) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *CreateProductTierBillingPlanRequest) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetPricing

`func (o *CreateProductTierBillingPlanRequest) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *CreateProductTierBillingPlanRequest) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *CreateProductTierBillingPlanRequest) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.


### SetPricingNil

`func (o *CreateProductTierBillingPlanRequest) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *CreateProductTierBillingPlanRequest) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetServiceId

`func (o *CreateProductTierBillingPlanRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateProductTierBillingPlanRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateProductTierBillingPlanRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *CreateProductTierBillingPlanRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateProductTierBillingPlanRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateProductTierBillingPlanRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


