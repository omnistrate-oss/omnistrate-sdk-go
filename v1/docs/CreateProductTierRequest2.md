# CreateProductTierRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Allow creates when payment not configured | [optional] 
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this product tier is available on | [optional] 
**AzureRegions** | Pointer to **[]string** | The Azure regions that this product tier is available on | [optional] 
**BillingProductID** | Pointer to **string** | Optional billing product ID for tax purposes | [optional] 
**BillingProviders** | Pointer to **[]string** | List of billing providers to be used for the product tier | [optional] 
**DefaultBillingProvider** | Pointer to **string** | The default billing provider to be used for the product tier | [optional] 
**DeploymentConfiguration** | Pointer to [**ProductTierDeploymentConfiguration**](ProductTierDeploymentConfiguration.md) |  | [optional] 
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
**PrivateRegions** | Pointer to **[]string** | The private regions that this product tier is available on | [optional] 
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

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *CreateProductTierRequest2) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *CreateProductTierRequest2) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *CreateProductTierRequest2) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *CreateProductTierRequest2) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

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

### GetAzureRegions

`func (o *CreateProductTierRequest2) GetAzureRegions() []string`

GetAzureRegions returns the AzureRegions field if non-nil, zero value otherwise.

### GetAzureRegionsOk

`func (o *CreateProductTierRequest2) GetAzureRegionsOk() (*[]string, bool)`

GetAzureRegionsOk returns a tuple with the AzureRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRegions

`func (o *CreateProductTierRequest2) SetAzureRegions(v []string)`

SetAzureRegions sets AzureRegions field to given value.

### HasAzureRegions

`func (o *CreateProductTierRequest2) HasAzureRegions() bool`

HasAzureRegions returns a boolean if a field has been set.

### GetBillingProductID

`func (o *CreateProductTierRequest2) GetBillingProductID() string`

GetBillingProductID returns the BillingProductID field if non-nil, zero value otherwise.

### GetBillingProductIDOk

`func (o *CreateProductTierRequest2) GetBillingProductIDOk() (*string, bool)`

GetBillingProductIDOk returns a tuple with the BillingProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProductID

`func (o *CreateProductTierRequest2) SetBillingProductID(v string)`

SetBillingProductID sets BillingProductID field to given value.

### HasBillingProductID

`func (o *CreateProductTierRequest2) HasBillingProductID() bool`

HasBillingProductID returns a boolean if a field has been set.

### GetBillingProviders

`func (o *CreateProductTierRequest2) GetBillingProviders() []string`

GetBillingProviders returns the BillingProviders field if non-nil, zero value otherwise.

### GetBillingProvidersOk

`func (o *CreateProductTierRequest2) GetBillingProvidersOk() (*[]string, bool)`

GetBillingProvidersOk returns a tuple with the BillingProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProviders

`func (o *CreateProductTierRequest2) SetBillingProviders(v []string)`

SetBillingProviders sets BillingProviders field to given value.

### HasBillingProviders

`func (o *CreateProductTierRequest2) HasBillingProviders() bool`

HasBillingProviders returns a boolean if a field has been set.

### GetDefaultBillingProvider

`func (o *CreateProductTierRequest2) GetDefaultBillingProvider() string`

GetDefaultBillingProvider returns the DefaultBillingProvider field if non-nil, zero value otherwise.

### GetDefaultBillingProviderOk

`func (o *CreateProductTierRequest2) GetDefaultBillingProviderOk() (*string, bool)`

GetDefaultBillingProviderOk returns a tuple with the DefaultBillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBillingProvider

`func (o *CreateProductTierRequest2) SetDefaultBillingProvider(v string)`

SetDefaultBillingProvider sets DefaultBillingProvider field to given value.

### HasDefaultBillingProvider

`func (o *CreateProductTierRequest2) HasDefaultBillingProvider() bool`

HasDefaultBillingProvider returns a boolean if a field has been set.

### GetDeploymentConfiguration

`func (o *CreateProductTierRequest2) GetDeploymentConfiguration() ProductTierDeploymentConfiguration`

GetDeploymentConfiguration returns the DeploymentConfiguration field if non-nil, zero value otherwise.

### GetDeploymentConfigurationOk

`func (o *CreateProductTierRequest2) GetDeploymentConfigurationOk() (*ProductTierDeploymentConfiguration, bool)`

GetDeploymentConfigurationOk returns a tuple with the DeploymentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfiguration

`func (o *CreateProductTierRequest2) SetDeploymentConfiguration(v ProductTierDeploymentConfiguration)`

SetDeploymentConfiguration sets DeploymentConfiguration field to given value.

### HasDeploymentConfiguration

`func (o *CreateProductTierRequest2) HasDeploymentConfiguration() bool`

HasDeploymentConfiguration returns a boolean if a field has been set.

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

### GetExportUsageMetering

`func (o *CreateProductTierRequest2) GetExportUsageMetering() bool`

GetExportUsageMetering returns the ExportUsageMetering field if non-nil, zero value otherwise.

### GetExportUsageMeteringOk

`func (o *CreateProductTierRequest2) GetExportUsageMeteringOk() (*bool, bool)`

GetExportUsageMeteringOk returns a tuple with the ExportUsageMetering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUsageMetering

`func (o *CreateProductTierRequest2) SetExportUsageMetering(v bool)`

SetExportUsageMetering sets ExportUsageMetering field to given value.

### HasExportUsageMetering

`func (o *CreateProductTierRequest2) HasExportUsageMetering() bool`

HasExportUsageMetering returns a boolean if a field has been set.

### GetExportUsageMeteringConfig

`func (o *CreateProductTierRequest2) GetExportUsageMeteringConfig() map[string]interface{}`

GetExportUsageMeteringConfig returns the ExportUsageMeteringConfig field if non-nil, zero value otherwise.

### GetExportUsageMeteringConfigOk

`func (o *CreateProductTierRequest2) GetExportUsageMeteringConfigOk() (*map[string]interface{}, bool)`

GetExportUsageMeteringConfigOk returns a tuple with the ExportUsageMeteringConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUsageMeteringConfig

`func (o *CreateProductTierRequest2) SetExportUsageMeteringConfig(v map[string]interface{})`

SetExportUsageMeteringConfig sets ExportUsageMeteringConfig field to given value.

### HasExportUsageMeteringConfig

`func (o *CreateProductTierRequest2) HasExportUsageMeteringConfig() bool`

HasExportUsageMeteringConfig returns a boolean if a field has been set.

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

### GetMaxNumberOfInstances

`func (o *CreateProductTierRequest2) GetMaxNumberOfInstances() int64`

GetMaxNumberOfInstances returns the MaxNumberOfInstances field if non-nil, zero value otherwise.

### GetMaxNumberOfInstancesOk

`func (o *CreateProductTierRequest2) GetMaxNumberOfInstancesOk() (*int64, bool)`

GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfInstances

`func (o *CreateProductTierRequest2) SetMaxNumberOfInstances(v int64)`

SetMaxNumberOfInstances sets MaxNumberOfInstances field to given value.

### HasMaxNumberOfInstances

`func (o *CreateProductTierRequest2) HasMaxNumberOfInstances() bool`

HasMaxNumberOfInstances returns a boolean if a field has been set.

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


### GetPricePerUnit

`func (o *CreateProductTierRequest2) GetPricePerUnit() map[string]interface{}`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *CreateProductTierRequest2) GetPricePerUnitOk() (*map[string]interface{}, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *CreateProductTierRequest2) SetPricePerUnit(v map[string]interface{})`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *CreateProductTierRequest2) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

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
### GetPrivateRegions

`func (o *CreateProductTierRequest2) GetPrivateRegions() []string`

GetPrivateRegions returns the PrivateRegions field if non-nil, zero value otherwise.

### GetPrivateRegionsOk

`func (o *CreateProductTierRequest2) GetPrivateRegionsOk() (*[]string, bool)`

GetPrivateRegionsOk returns a tuple with the PrivateRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateRegions

`func (o *CreateProductTierRequest2) SetPrivateRegions(v []string)`

SetPrivateRegions sets PrivateRegions field to given value.

### HasPrivateRegions

`func (o *CreateProductTierRequest2) HasPrivateRegions() bool`

HasPrivateRegions returns a boolean if a field has been set.

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


