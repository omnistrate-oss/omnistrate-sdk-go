# DescribeProductTierResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Allow creates when payment not configured | [optional] 
**ApiGroups** | Pointer to **map[string]interface{}** | The resources that belong to this service API bundle and their active versions | [optional] 
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this product tier is available on | [optional] 
**AzureRegions** | Pointer to **[]string** | The Azure regions that this product tier is available on | [optional] 
**BillingProductID** | Pointer to **string** | Optional billing product ID for tax purposes | [optional] 
**BillingProviders** | Pointer to **[]string** | List of billing providers to be used for the product tier | [optional] 
**CloudProvidersConfigReadiness** | Pointer to **map[string]interface{}** | The readiness of the cloud providers configurations | [optional] 
**DefaultBillingProvider** | Pointer to **string** | The billing provider type | [optional] 
**Description** | **string** | A brief description of the product tier | 
**Documentation** | **string** | Documentation | 
**EnabledFeatures** | Pointer to [**[]ProductTierFeatureDetail**](ProductTierFeatureDetail.md) | The features that are enabled for this product tier, including scope details and configuration | [optional] 
**ExportUsageMetering** | Pointer to **bool** | Export usage metering data | [optional] 
**ExportUsageMeteringConfig** | Pointer to **map[string]interface{}** | Export usage metering data configuration | [optional] 
**Features** | Pointer to **map[string]interface{}** | The features that are enabled / disabled for this product tier | [optional] 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this product tier is available on | [optional] 
**Id** | **string** | ID of a Product Tier | 
**IsDisabled** | **bool** | Flag to indicate if the product tier is disabled. | 
**Key** | **string** | Unique Key of the product tier | 
**MaxNumberOfInstances** | Pointer to **int64** | Maximum number of instances | [optional] 
**Name** | **string** | Name of the product tier | 
**PlanDescription** | **string** | A brief description for the end user of the product tier | 
**PricePerUnit** | Pointer to **map[string]interface{}** | Price per unit. | [optional] 
**Pricing** | **interface{}** | Pricing | 
**PrivateRegions** | Pointer to **[]string** | The Private regions that this product tier is available on | [optional] 
**ServiceId** | **string** | ID of a Service | 
**ServiceModelId** | **string** | ID of a Service Model | 
**Support** | **string** | Support | 
**TierType** | **string** | ProductTierType is the type of tier for a product | 

## Methods

### NewDescribeProductTierResult

`func NewDescribeProductTierResult(description string, documentation string, id string, isDisabled bool, key string, name string, planDescription string, pricing interface{}, serviceId string, serviceModelId string, support string, tierType string, ) *DescribeProductTierResult`

NewDescribeProductTierResult instantiates a new DescribeProductTierResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProductTierResultWithDefaults

`func NewDescribeProductTierResultWithDefaults() *DescribeProductTierResult`

NewDescribeProductTierResultWithDefaults instantiates a new DescribeProductTierResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *DescribeProductTierResult) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *DescribeProductTierResult) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *DescribeProductTierResult) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *DescribeProductTierResult) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

### GetApiGroups

`func (o *DescribeProductTierResult) GetApiGroups() map[string]interface{}`

GetApiGroups returns the ApiGroups field if non-nil, zero value otherwise.

### GetApiGroupsOk

`func (o *DescribeProductTierResult) GetApiGroupsOk() (*map[string]interface{}, bool)`

GetApiGroupsOk returns a tuple with the ApiGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroups

`func (o *DescribeProductTierResult) SetApiGroups(v map[string]interface{})`

SetApiGroups sets ApiGroups field to given value.

### HasApiGroups

`func (o *DescribeProductTierResult) HasApiGroups() bool`

HasApiGroups returns a boolean if a field has been set.

### GetAutoApproveSubscription

`func (o *DescribeProductTierResult) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *DescribeProductTierResult) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *DescribeProductTierResult) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *DescribeProductTierResult) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetAwsRegions

`func (o *DescribeProductTierResult) GetAwsRegions() []string`

GetAwsRegions returns the AwsRegions field if non-nil, zero value otherwise.

### GetAwsRegionsOk

`func (o *DescribeProductTierResult) GetAwsRegionsOk() (*[]string, bool)`

GetAwsRegionsOk returns a tuple with the AwsRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegions

`func (o *DescribeProductTierResult) SetAwsRegions(v []string)`

SetAwsRegions sets AwsRegions field to given value.

### HasAwsRegions

`func (o *DescribeProductTierResult) HasAwsRegions() bool`

HasAwsRegions returns a boolean if a field has been set.

### GetAzureRegions

`func (o *DescribeProductTierResult) GetAzureRegions() []string`

GetAzureRegions returns the AzureRegions field if non-nil, zero value otherwise.

### GetAzureRegionsOk

`func (o *DescribeProductTierResult) GetAzureRegionsOk() (*[]string, bool)`

GetAzureRegionsOk returns a tuple with the AzureRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRegions

`func (o *DescribeProductTierResult) SetAzureRegions(v []string)`

SetAzureRegions sets AzureRegions field to given value.

### HasAzureRegions

`func (o *DescribeProductTierResult) HasAzureRegions() bool`

HasAzureRegions returns a boolean if a field has been set.

### GetBillingProductID

`func (o *DescribeProductTierResult) GetBillingProductID() string`

GetBillingProductID returns the BillingProductID field if non-nil, zero value otherwise.

### GetBillingProductIDOk

`func (o *DescribeProductTierResult) GetBillingProductIDOk() (*string, bool)`

GetBillingProductIDOk returns a tuple with the BillingProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProductID

`func (o *DescribeProductTierResult) SetBillingProductID(v string)`

SetBillingProductID sets BillingProductID field to given value.

### HasBillingProductID

`func (o *DescribeProductTierResult) HasBillingProductID() bool`

HasBillingProductID returns a boolean if a field has been set.

### GetBillingProviders

`func (o *DescribeProductTierResult) GetBillingProviders() []string`

GetBillingProviders returns the BillingProviders field if non-nil, zero value otherwise.

### GetBillingProvidersOk

`func (o *DescribeProductTierResult) GetBillingProvidersOk() (*[]string, bool)`

GetBillingProvidersOk returns a tuple with the BillingProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProviders

`func (o *DescribeProductTierResult) SetBillingProviders(v []string)`

SetBillingProviders sets BillingProviders field to given value.

### HasBillingProviders

`func (o *DescribeProductTierResult) HasBillingProviders() bool`

HasBillingProviders returns a boolean if a field has been set.

### GetCloudProvidersConfigReadiness

`func (o *DescribeProductTierResult) GetCloudProvidersConfigReadiness() map[string]interface{}`

GetCloudProvidersConfigReadiness returns the CloudProvidersConfigReadiness field if non-nil, zero value otherwise.

### GetCloudProvidersConfigReadinessOk

`func (o *DescribeProductTierResult) GetCloudProvidersConfigReadinessOk() (*map[string]interface{}, bool)`

GetCloudProvidersConfigReadinessOk returns a tuple with the CloudProvidersConfigReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvidersConfigReadiness

`func (o *DescribeProductTierResult) SetCloudProvidersConfigReadiness(v map[string]interface{})`

SetCloudProvidersConfigReadiness sets CloudProvidersConfigReadiness field to given value.

### HasCloudProvidersConfigReadiness

`func (o *DescribeProductTierResult) HasCloudProvidersConfigReadiness() bool`

HasCloudProvidersConfigReadiness returns a boolean if a field has been set.

### GetDefaultBillingProvider

`func (o *DescribeProductTierResult) GetDefaultBillingProvider() string`

GetDefaultBillingProvider returns the DefaultBillingProvider field if non-nil, zero value otherwise.

### GetDefaultBillingProviderOk

`func (o *DescribeProductTierResult) GetDefaultBillingProviderOk() (*string, bool)`

GetDefaultBillingProviderOk returns a tuple with the DefaultBillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBillingProvider

`func (o *DescribeProductTierResult) SetDefaultBillingProvider(v string)`

SetDefaultBillingProvider sets DefaultBillingProvider field to given value.

### HasDefaultBillingProvider

`func (o *DescribeProductTierResult) HasDefaultBillingProvider() bool`

HasDefaultBillingProvider returns a boolean if a field has been set.

### GetDescription

`func (o *DescribeProductTierResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeProductTierResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeProductTierResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDocumentation

`func (o *DescribeProductTierResult) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *DescribeProductTierResult) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *DescribeProductTierResult) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.


### GetEnabledFeatures

`func (o *DescribeProductTierResult) GetEnabledFeatures() []ProductTierFeatureDetail`

GetEnabledFeatures returns the EnabledFeatures field if non-nil, zero value otherwise.

### GetEnabledFeaturesOk

`func (o *DescribeProductTierResult) GetEnabledFeaturesOk() (*[]ProductTierFeatureDetail, bool)`

GetEnabledFeaturesOk returns a tuple with the EnabledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFeatures

`func (o *DescribeProductTierResult) SetEnabledFeatures(v []ProductTierFeatureDetail)`

SetEnabledFeatures sets EnabledFeatures field to given value.

### HasEnabledFeatures

`func (o *DescribeProductTierResult) HasEnabledFeatures() bool`

HasEnabledFeatures returns a boolean if a field has been set.

### GetExportUsageMetering

`func (o *DescribeProductTierResult) GetExportUsageMetering() bool`

GetExportUsageMetering returns the ExportUsageMetering field if non-nil, zero value otherwise.

### GetExportUsageMeteringOk

`func (o *DescribeProductTierResult) GetExportUsageMeteringOk() (*bool, bool)`

GetExportUsageMeteringOk returns a tuple with the ExportUsageMetering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUsageMetering

`func (o *DescribeProductTierResult) SetExportUsageMetering(v bool)`

SetExportUsageMetering sets ExportUsageMetering field to given value.

### HasExportUsageMetering

`func (o *DescribeProductTierResult) HasExportUsageMetering() bool`

HasExportUsageMetering returns a boolean if a field has been set.

### GetExportUsageMeteringConfig

`func (o *DescribeProductTierResult) GetExportUsageMeteringConfig() map[string]interface{}`

GetExportUsageMeteringConfig returns the ExportUsageMeteringConfig field if non-nil, zero value otherwise.

### GetExportUsageMeteringConfigOk

`func (o *DescribeProductTierResult) GetExportUsageMeteringConfigOk() (*map[string]interface{}, bool)`

GetExportUsageMeteringConfigOk returns a tuple with the ExportUsageMeteringConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUsageMeteringConfig

`func (o *DescribeProductTierResult) SetExportUsageMeteringConfig(v map[string]interface{})`

SetExportUsageMeteringConfig sets ExportUsageMeteringConfig field to given value.

### HasExportUsageMeteringConfig

`func (o *DescribeProductTierResult) HasExportUsageMeteringConfig() bool`

HasExportUsageMeteringConfig returns a boolean if a field has been set.

### GetFeatures

`func (o *DescribeProductTierResult) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DescribeProductTierResult) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DescribeProductTierResult) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DescribeProductTierResult) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetGcpRegions

`func (o *DescribeProductTierResult) GetGcpRegions() []string`

GetGcpRegions returns the GcpRegions field if non-nil, zero value otherwise.

### GetGcpRegionsOk

`func (o *DescribeProductTierResult) GetGcpRegionsOk() (*[]string, bool)`

GetGcpRegionsOk returns a tuple with the GcpRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRegions

`func (o *DescribeProductTierResult) SetGcpRegions(v []string)`

SetGcpRegions sets GcpRegions field to given value.

### HasGcpRegions

`func (o *DescribeProductTierResult) HasGcpRegions() bool`

HasGcpRegions returns a boolean if a field has been set.

### GetId

`func (o *DescribeProductTierResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeProductTierResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeProductTierResult) SetId(v string)`

SetId sets Id field to given value.


### GetIsDisabled

`func (o *DescribeProductTierResult) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *DescribeProductTierResult) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *DescribeProductTierResult) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.


### GetKey

`func (o *DescribeProductTierResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DescribeProductTierResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DescribeProductTierResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetMaxNumberOfInstances

`func (o *DescribeProductTierResult) GetMaxNumberOfInstances() int64`

GetMaxNumberOfInstances returns the MaxNumberOfInstances field if non-nil, zero value otherwise.

### GetMaxNumberOfInstancesOk

`func (o *DescribeProductTierResult) GetMaxNumberOfInstancesOk() (*int64, bool)`

GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfInstances

`func (o *DescribeProductTierResult) SetMaxNumberOfInstances(v int64)`

SetMaxNumberOfInstances sets MaxNumberOfInstances field to given value.

### HasMaxNumberOfInstances

`func (o *DescribeProductTierResult) HasMaxNumberOfInstances() bool`

HasMaxNumberOfInstances returns a boolean if a field has been set.

### GetName

`func (o *DescribeProductTierResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeProductTierResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeProductTierResult) SetName(v string)`

SetName sets Name field to given value.


### GetPlanDescription

`func (o *DescribeProductTierResult) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *DescribeProductTierResult) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *DescribeProductTierResult) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.


### GetPricePerUnit

`func (o *DescribeProductTierResult) GetPricePerUnit() map[string]interface{}`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *DescribeProductTierResult) GetPricePerUnitOk() (*map[string]interface{}, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *DescribeProductTierResult) SetPricePerUnit(v map[string]interface{})`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *DescribeProductTierResult) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

### GetPricing

`func (o *DescribeProductTierResult) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *DescribeProductTierResult) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *DescribeProductTierResult) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.


### SetPricingNil

`func (o *DescribeProductTierResult) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *DescribeProductTierResult) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetPrivateRegions

`func (o *DescribeProductTierResult) GetPrivateRegions() []string`

GetPrivateRegions returns the PrivateRegions field if non-nil, zero value otherwise.

### GetPrivateRegionsOk

`func (o *DescribeProductTierResult) GetPrivateRegionsOk() (*[]string, bool)`

GetPrivateRegionsOk returns a tuple with the PrivateRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateRegions

`func (o *DescribeProductTierResult) SetPrivateRegions(v []string)`

SetPrivateRegions sets PrivateRegions field to given value.

### HasPrivateRegions

`func (o *DescribeProductTierResult) HasPrivateRegions() bool`

HasPrivateRegions returns a boolean if a field has been set.

### GetServiceId

`func (o *DescribeProductTierResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeProductTierResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeProductTierResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceModelId

`func (o *DescribeProductTierResult) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *DescribeProductTierResult) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *DescribeProductTierResult) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetSupport

`func (o *DescribeProductTierResult) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *DescribeProductTierResult) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *DescribeProductTierResult) SetSupport(v string)`

SetSupport sets Support field to given value.


### GetTierType

`func (o *DescribeProductTierResult) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *DescribeProductTierResult) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *DescribeProductTierResult) SetTierType(v string)`

SetTierType sets TierType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


