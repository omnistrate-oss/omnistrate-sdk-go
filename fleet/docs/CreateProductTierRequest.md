# CreateProductTierRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Allow creates when payment not configured | [optional] 
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this product tier is available on | [optional] 
**AzureRegions** | Pointer to **[]string** | The Azure regions that this product tier is available on | [optional] 
**Description** | **string** | A brief description of the product tier | 
**Documentation** | Pointer to **string** | Documentation | [optional] 
**ExportUsageMetering** | Pointer to **bool** | Export usage metering data | [optional] 
**ExportUsageMeteringConfig** | Pointer to **map[string]interface{}** | Export usage metering data configuration | [optional] 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this product tier is available on | [optional] 
**IsDisabled** | Pointer to **bool** | Create the product tier in a disabled state. Enabling the product tier will let end-customers subscribe and use the service plan. | [optional] 
**MaxNumberOfInstances** | Pointer to **int64** | Maximum number of instances | [optional] 
**Name** | **string** | Name of the product tier | 
**PlanDescription** | **string** | A brief description for the end user of the product tier | 
**PricePerUnit** | Pointer to **map[string]interface{}** | Price per unit. | [optional] 
**Pricing** | Pointer to **interface{}** | Pricing | [optional] 
**ServiceId** | **string** | ID of a Service | 
**ServiceModelId** | **string** | ID of a Service Model | 
**Support** | Pointer to **string** | Support | [optional] 
**TierType** | **string** | ProductTierType is the type of tier for a product | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateProductTierRequest

`func NewCreateProductTierRequest(description string, name string, planDescription string, serviceId string, serviceModelId string, tierType string, token string, ) *CreateProductTierRequest`

NewCreateProductTierRequest instantiates a new CreateProductTierRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductTierRequestWithDefaults

`func NewCreateProductTierRequestWithDefaults() *CreateProductTierRequest`

NewCreateProductTierRequestWithDefaults instantiates a new CreateProductTierRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *CreateProductTierRequest) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *CreateProductTierRequest) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *CreateProductTierRequest) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *CreateProductTierRequest) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

### GetAutoApproveSubscription

`func (o *CreateProductTierRequest) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *CreateProductTierRequest) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *CreateProductTierRequest) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *CreateProductTierRequest) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetAwsRegions

`func (o *CreateProductTierRequest) GetAwsRegions() []string`

GetAwsRegions returns the AwsRegions field if non-nil, zero value otherwise.

### GetAwsRegionsOk

`func (o *CreateProductTierRequest) GetAwsRegionsOk() (*[]string, bool)`

GetAwsRegionsOk returns a tuple with the AwsRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegions

`func (o *CreateProductTierRequest) SetAwsRegions(v []string)`

SetAwsRegions sets AwsRegions field to given value.

### HasAwsRegions

`func (o *CreateProductTierRequest) HasAwsRegions() bool`

HasAwsRegions returns a boolean if a field has been set.

### GetAzureRegions

`func (o *CreateProductTierRequest) GetAzureRegions() []string`

GetAzureRegions returns the AzureRegions field if non-nil, zero value otherwise.

### GetAzureRegionsOk

`func (o *CreateProductTierRequest) GetAzureRegionsOk() (*[]string, bool)`

GetAzureRegionsOk returns a tuple with the AzureRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRegions

`func (o *CreateProductTierRequest) SetAzureRegions(v []string)`

SetAzureRegions sets AzureRegions field to given value.

### HasAzureRegions

`func (o *CreateProductTierRequest) HasAzureRegions() bool`

HasAzureRegions returns a boolean if a field has been set.

### GetDescription

`func (o *CreateProductTierRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductTierRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductTierRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDocumentation

`func (o *CreateProductTierRequest) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *CreateProductTierRequest) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *CreateProductTierRequest) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *CreateProductTierRequest) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetExportUsageMetering

`func (o *CreateProductTierRequest) GetExportUsageMetering() bool`

GetExportUsageMetering returns the ExportUsageMetering field if non-nil, zero value otherwise.

### GetExportUsageMeteringOk

`func (o *CreateProductTierRequest) GetExportUsageMeteringOk() (*bool, bool)`

GetExportUsageMeteringOk returns a tuple with the ExportUsageMetering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUsageMetering

`func (o *CreateProductTierRequest) SetExportUsageMetering(v bool)`

SetExportUsageMetering sets ExportUsageMetering field to given value.

### HasExportUsageMetering

`func (o *CreateProductTierRequest) HasExportUsageMetering() bool`

HasExportUsageMetering returns a boolean if a field has been set.

### GetExportUsageMeteringConfig

`func (o *CreateProductTierRequest) GetExportUsageMeteringConfig() map[string]interface{}`

GetExportUsageMeteringConfig returns the ExportUsageMeteringConfig field if non-nil, zero value otherwise.

### GetExportUsageMeteringConfigOk

`func (o *CreateProductTierRequest) GetExportUsageMeteringConfigOk() (*map[string]interface{}, bool)`

GetExportUsageMeteringConfigOk returns a tuple with the ExportUsageMeteringConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUsageMeteringConfig

`func (o *CreateProductTierRequest) SetExportUsageMeteringConfig(v map[string]interface{})`

SetExportUsageMeteringConfig sets ExportUsageMeteringConfig field to given value.

### HasExportUsageMeteringConfig

`func (o *CreateProductTierRequest) HasExportUsageMeteringConfig() bool`

HasExportUsageMeteringConfig returns a boolean if a field has been set.

### GetGcpRegions

`func (o *CreateProductTierRequest) GetGcpRegions() []string`

GetGcpRegions returns the GcpRegions field if non-nil, zero value otherwise.

### GetGcpRegionsOk

`func (o *CreateProductTierRequest) GetGcpRegionsOk() (*[]string, bool)`

GetGcpRegionsOk returns a tuple with the GcpRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRegions

`func (o *CreateProductTierRequest) SetGcpRegions(v []string)`

SetGcpRegions sets GcpRegions field to given value.

### HasGcpRegions

`func (o *CreateProductTierRequest) HasGcpRegions() bool`

HasGcpRegions returns a boolean if a field has been set.

### GetIsDisabled

`func (o *CreateProductTierRequest) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *CreateProductTierRequest) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *CreateProductTierRequest) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *CreateProductTierRequest) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetMaxNumberOfInstances

`func (o *CreateProductTierRequest) GetMaxNumberOfInstances() int64`

GetMaxNumberOfInstances returns the MaxNumberOfInstances field if non-nil, zero value otherwise.

### GetMaxNumberOfInstancesOk

`func (o *CreateProductTierRequest) GetMaxNumberOfInstancesOk() (*int64, bool)`

GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfInstances

`func (o *CreateProductTierRequest) SetMaxNumberOfInstances(v int64)`

SetMaxNumberOfInstances sets MaxNumberOfInstances field to given value.

### HasMaxNumberOfInstances

`func (o *CreateProductTierRequest) HasMaxNumberOfInstances() bool`

HasMaxNumberOfInstances returns a boolean if a field has been set.

### GetName

`func (o *CreateProductTierRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProductTierRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProductTierRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPlanDescription

`func (o *CreateProductTierRequest) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *CreateProductTierRequest) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *CreateProductTierRequest) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.


### GetPricePerUnit

`func (o *CreateProductTierRequest) GetPricePerUnit() map[string]interface{}`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *CreateProductTierRequest) GetPricePerUnitOk() (*map[string]interface{}, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *CreateProductTierRequest) SetPricePerUnit(v map[string]interface{})`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *CreateProductTierRequest) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

### GetPricing

`func (o *CreateProductTierRequest) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *CreateProductTierRequest) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *CreateProductTierRequest) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *CreateProductTierRequest) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *CreateProductTierRequest) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *CreateProductTierRequest) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetServiceId

`func (o *CreateProductTierRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateProductTierRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateProductTierRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceModelId

`func (o *CreateProductTierRequest) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *CreateProductTierRequest) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *CreateProductTierRequest) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetSupport

`func (o *CreateProductTierRequest) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *CreateProductTierRequest) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *CreateProductTierRequest) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *CreateProductTierRequest) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetTierType

`func (o *CreateProductTierRequest) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *CreateProductTierRequest) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *CreateProductTierRequest) SetTierType(v string)`

SetTierType sets TierType field to given value.


### GetToken

`func (o *CreateProductTierRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateProductTierRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateProductTierRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


