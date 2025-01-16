# CreateProductTierRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this product tier is available on | [optional] 
**Description** | **string** | A brief description of the product tier | 
**Documentation** | Pointer to **string** | Documentation | [optional] 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this product tier is available on | [optional] 
**IsDisabled** | Pointer to **bool** | Create the product tier in a disabled state. Enabling the product tier will let end-customers subscribe and use the service plan. | [optional] 
**Name** | **string** | Name of the product tier | 
**PlanDescription** | **string** | A brief description for the end user of the product tier | 
**Pricing** | Pointer to **interface{}** | Pricing | [optional] 
**ServiceModelId** | **string** | Service model ID | 
**Support** | Pointer to **string** | Support | [optional] 
**TierType** | **string** | Tier type | 

## Methods

### NewCreateProductTierRequest2

`func NewCreateProductTierRequest2(description string, name string, planDescription string, serviceModelId string, tierType string, ) *CreateProductTierRequest2`

NewCreateProductTierRequest2 instantiates a new CreateProductTierRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductTierRequest2WithDefaults

`func NewCreateProductTierRequest2WithDefaults() *CreateProductTierRequest2`

NewCreateProductTierRequest2WithDefaults instantiates a new CreateProductTierRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *CreateProductTierRequest2) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *CreateProductTierRequest2) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *CreateProductTierRequest2) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *CreateProductTierRequest2) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetAwsRegions

`func (o *CreateProductTierRequest2) GetAwsRegions() []string`

GetAwsRegions returns the AwsRegions field if non-nil, zero value otherwise.

### GetAwsRegionsOk

`func (o *CreateProductTierRequest2) GetAwsRegionsOk() (*[]string, bool)`

GetAwsRegionsOk returns a tuple with the AwsRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegions

`func (o *CreateProductTierRequest2) SetAwsRegions(v []string)`

SetAwsRegions sets AwsRegions field to given value.

### HasAwsRegions

`func (o *CreateProductTierRequest2) HasAwsRegions() bool`

HasAwsRegions returns a boolean if a field has been set.

### GetDescription

`func (o *CreateProductTierRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductTierRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductTierRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDocumentation

`func (o *CreateProductTierRequest2) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *CreateProductTierRequest2) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *CreateProductTierRequest2) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *CreateProductTierRequest2) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetGcpRegions

`func (o *CreateProductTierRequest2) GetGcpRegions() []string`

GetGcpRegions returns the GcpRegions field if non-nil, zero value otherwise.

### GetGcpRegionsOk

`func (o *CreateProductTierRequest2) GetGcpRegionsOk() (*[]string, bool)`

GetGcpRegionsOk returns a tuple with the GcpRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRegions

`func (o *CreateProductTierRequest2) SetGcpRegions(v []string)`

SetGcpRegions sets GcpRegions field to given value.

### HasGcpRegions

`func (o *CreateProductTierRequest2) HasGcpRegions() bool`

HasGcpRegions returns a boolean if a field has been set.

### GetIsDisabled

`func (o *CreateProductTierRequest2) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *CreateProductTierRequest2) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *CreateProductTierRequest2) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *CreateProductTierRequest2) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetName

`func (o *CreateProductTierRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProductTierRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProductTierRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetPlanDescription

`func (o *CreateProductTierRequest2) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *CreateProductTierRequest2) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *CreateProductTierRequest2) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.


### GetPricing

`func (o *CreateProductTierRequest2) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *CreateProductTierRequest2) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *CreateProductTierRequest2) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *CreateProductTierRequest2) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *CreateProductTierRequest2) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *CreateProductTierRequest2) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetServiceModelId

`func (o *CreateProductTierRequest2) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *CreateProductTierRequest2) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *CreateProductTierRequest2) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetSupport

`func (o *CreateProductTierRequest2) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *CreateProductTierRequest2) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *CreateProductTierRequest2) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *CreateProductTierRequest2) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetTierType

`func (o *CreateProductTierRequest2) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *CreateProductTierRequest2) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *CreateProductTierRequest2) SetTierType(v string)`

SetTierType sets TierType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


