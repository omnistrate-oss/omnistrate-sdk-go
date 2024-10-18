# CopyProductTierRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this product tier is available on | [optional] 
**Description** | **string** | A brief description of the product tier | 
**Documentation** | Pointer to **string** | Documentation | [optional] 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this product tier is available on | [optional] 
**Name** | **string** | Name of the product tier | 
**PlanDescription** | Pointer to **string** | A brief description for the end user of the product tier | [optional] 
**Pricing** | Pointer to **interface{}** | Pricing | [optional] 
**ServiceModelId** | **string** | Service model ID | 
**Support** | Pointer to **string** | Support | [optional] 
**TargetTierType** | Pointer to **string** | Tier type | [optional] 

## Methods

### NewCopyProductTierRequestBody

`func NewCopyProductTierRequestBody(description string, name string, serviceModelId string, ) *CopyProductTierRequestBody`

NewCopyProductTierRequestBody instantiates a new CopyProductTierRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyProductTierRequestBodyWithDefaults

`func NewCopyProductTierRequestBodyWithDefaults() *CopyProductTierRequestBody`

NewCopyProductTierRequestBodyWithDefaults instantiates a new CopyProductTierRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *CopyProductTierRequestBody) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *CopyProductTierRequestBody) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *CopyProductTierRequestBody) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *CopyProductTierRequestBody) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetAwsRegions

`func (o *CopyProductTierRequestBody) GetAwsRegions() []string`

GetAwsRegions returns the AwsRegions field if non-nil, zero value otherwise.

### GetAwsRegionsOk

`func (o *CopyProductTierRequestBody) GetAwsRegionsOk() (*[]string, bool)`

GetAwsRegionsOk returns a tuple with the AwsRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegions

`func (o *CopyProductTierRequestBody) SetAwsRegions(v []string)`

SetAwsRegions sets AwsRegions field to given value.

### HasAwsRegions

`func (o *CopyProductTierRequestBody) HasAwsRegions() bool`

HasAwsRegions returns a boolean if a field has been set.

### GetDescription

`func (o *CopyProductTierRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CopyProductTierRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CopyProductTierRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDocumentation

`func (o *CopyProductTierRequestBody) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *CopyProductTierRequestBody) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *CopyProductTierRequestBody) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *CopyProductTierRequestBody) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetGcpRegions

`func (o *CopyProductTierRequestBody) GetGcpRegions() []string`

GetGcpRegions returns the GcpRegions field if non-nil, zero value otherwise.

### GetGcpRegionsOk

`func (o *CopyProductTierRequestBody) GetGcpRegionsOk() (*[]string, bool)`

GetGcpRegionsOk returns a tuple with the GcpRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRegions

`func (o *CopyProductTierRequestBody) SetGcpRegions(v []string)`

SetGcpRegions sets GcpRegions field to given value.

### HasGcpRegions

`func (o *CopyProductTierRequestBody) HasGcpRegions() bool`

HasGcpRegions returns a boolean if a field has been set.

### GetName

`func (o *CopyProductTierRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CopyProductTierRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CopyProductTierRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetPlanDescription

`func (o *CopyProductTierRequestBody) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *CopyProductTierRequestBody) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *CopyProductTierRequestBody) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *CopyProductTierRequestBody) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### GetPricing

`func (o *CopyProductTierRequestBody) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *CopyProductTierRequestBody) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *CopyProductTierRequestBody) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *CopyProductTierRequestBody) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *CopyProductTierRequestBody) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *CopyProductTierRequestBody) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetServiceModelId

`func (o *CopyProductTierRequestBody) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *CopyProductTierRequestBody) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *CopyProductTierRequestBody) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetSupport

`func (o *CopyProductTierRequestBody) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *CopyProductTierRequestBody) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *CopyProductTierRequestBody) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *CopyProductTierRequestBody) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetTargetTierType

`func (o *CopyProductTierRequestBody) GetTargetTierType() string`

GetTargetTierType returns the TargetTierType field if non-nil, zero value otherwise.

### GetTargetTierTypeOk

`func (o *CopyProductTierRequestBody) GetTargetTierTypeOk() (*string, bool)`

GetTargetTierTypeOk returns a tuple with the TargetTierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTierType

`func (o *CopyProductTierRequestBody) SetTargetTierType(v string)`

SetTargetTierType sets TargetTierType field to given value.

### HasTargetTierType

`func (o *CopyProductTierRequestBody) HasTargetTierType() bool`

HasTargetTierType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


