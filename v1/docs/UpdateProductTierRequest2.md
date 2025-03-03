# UpdateProductTierRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this product tier is available on | [optional] 
**AzureRegions** | Pointer to **[]string** | The Azure regions that this product tier is available on | [optional] 
**Description** | Pointer to **string** | A brief description of the product tier | [optional] 
**Documentation** | Pointer to **string** | Documentation | [optional] 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this product tier is available on | [optional] 
**IsDisabled** | Pointer to **bool** | Update the product tier&#39;s state as enabled/disabled. Enabling the product tier will let end-customers subscribe and use the service plan. | [optional] 
**Name** | Pointer to **string** | Name of the product tier | [optional] 
**PlanDescription** | Pointer to **string** | A brief description for the end user of the product tier | [optional] 
**Pricing** | Pointer to **interface{}** | Pricing | [optional] 
**Support** | Pointer to **string** | Support | [optional] 
**TierType** | Pointer to **string** | Tier type | [optional] 

## Methods

### NewUpdateProductTierRequest2

`func NewUpdateProductTierRequest2() *UpdateProductTierRequest2`

NewUpdateProductTierRequest2 instantiates a new UpdateProductTierRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductTierRequest2WithDefaults

`func NewUpdateProductTierRequest2WithDefaults() *UpdateProductTierRequest2`

NewUpdateProductTierRequest2WithDefaults instantiates a new UpdateProductTierRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *UpdateProductTierRequest2) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *UpdateProductTierRequest2) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *UpdateProductTierRequest2) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *UpdateProductTierRequest2) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetAwsRegions

`func (o *UpdateProductTierRequest2) GetAwsRegions() []string`

GetAwsRegions returns the AwsRegions field if non-nil, zero value otherwise.

### GetAwsRegionsOk

`func (o *UpdateProductTierRequest2) GetAwsRegionsOk() (*[]string, bool)`

GetAwsRegionsOk returns a tuple with the AwsRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegions

`func (o *UpdateProductTierRequest2) SetAwsRegions(v []string)`

SetAwsRegions sets AwsRegions field to given value.

### HasAwsRegions

`func (o *UpdateProductTierRequest2) HasAwsRegions() bool`

HasAwsRegions returns a boolean if a field has been set.

### GetAzureRegions

`func (o *UpdateProductTierRequest2) GetAzureRegions() []string`

GetAzureRegions returns the AzureRegions field if non-nil, zero value otherwise.

### GetAzureRegionsOk

`func (o *UpdateProductTierRequest2) GetAzureRegionsOk() (*[]string, bool)`

GetAzureRegionsOk returns a tuple with the AzureRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRegions

`func (o *UpdateProductTierRequest2) SetAzureRegions(v []string)`

SetAzureRegions sets AzureRegions field to given value.

### HasAzureRegions

`func (o *UpdateProductTierRequest2) HasAzureRegions() bool`

HasAzureRegions returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProductTierRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProductTierRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProductTierRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProductTierRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentation

`func (o *UpdateProductTierRequest2) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *UpdateProductTierRequest2) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *UpdateProductTierRequest2) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *UpdateProductTierRequest2) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetGcpRegions

`func (o *UpdateProductTierRequest2) GetGcpRegions() []string`

GetGcpRegions returns the GcpRegions field if non-nil, zero value otherwise.

### GetGcpRegionsOk

`func (o *UpdateProductTierRequest2) GetGcpRegionsOk() (*[]string, bool)`

GetGcpRegionsOk returns a tuple with the GcpRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRegions

`func (o *UpdateProductTierRequest2) SetGcpRegions(v []string)`

SetGcpRegions sets GcpRegions field to given value.

### HasGcpRegions

`func (o *UpdateProductTierRequest2) HasGcpRegions() bool`

HasGcpRegions returns a boolean if a field has been set.

### GetIsDisabled

`func (o *UpdateProductTierRequest2) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UpdateProductTierRequest2) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UpdateProductTierRequest2) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UpdateProductTierRequest2) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetName

`func (o *UpdateProductTierRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProductTierRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProductTierRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProductTierRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanDescription

`func (o *UpdateProductTierRequest2) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *UpdateProductTierRequest2) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *UpdateProductTierRequest2) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *UpdateProductTierRequest2) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### GetPricing

`func (o *UpdateProductTierRequest2) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *UpdateProductTierRequest2) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *UpdateProductTierRequest2) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *UpdateProductTierRequest2) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *UpdateProductTierRequest2) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *UpdateProductTierRequest2) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetSupport

`func (o *UpdateProductTierRequest2) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *UpdateProductTierRequest2) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *UpdateProductTierRequest2) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *UpdateProductTierRequest2) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetTierType

`func (o *UpdateProductTierRequest2) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *UpdateProductTierRequest2) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *UpdateProductTierRequest2) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *UpdateProductTierRequest2) HasTierType() bool`

HasTierType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


