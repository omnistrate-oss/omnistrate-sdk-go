# UpdateProductTierRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Allow creates when payment not configured | [optional] 
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this product tier is available on | [optional] 
**AzureRegions** | Pointer to **[]string** | The Azure regions that this product tier is available on | [optional] 
**BillingProductID** | Pointer to **string** | Optional billing product ID for tax purposes | [optional] 
**BillingProviders** | Pointer to **[]string** | List of billing providers to be used for the product tier | [optional] 
**DefaultBillingProvider** | Pointer to **string** | The billing provider type | [optional] 
**DeploymentConfiguration** | Pointer to [**ProductTierDeploymentConfiguration**](ProductTierDeploymentConfiguration.md) |  | [optional] 
**Description** | Pointer to **string** | A brief description of the product tier | [optional] 
**Documentation** | Pointer to **string** | Documentation | [optional] 
**ExportUsageMetering** | Pointer to **bool** | Export usage metering data | [optional] 
**ExportUsageMeteringConfig** | Pointer to **map[string]interface{}** | Export usage metering data configuration | [optional] 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this product tier is available on | [optional] 
**Id** | **string** | ID of a Product Tier | 
**IsDisabled** | Pointer to **bool** | Update the product tier&#39;s state as enabled/disabled. Enabling the product tier will let end-customers subscribe and use the service plan. | [optional] 
**MaxNumberOfInstances** | Pointer to **int64** | Maximum number of instances. Set to -1 for unlimited. | [optional] 
**Name** | Pointer to **string** | Name of the product tier | [optional] 
**OciRegions** | Pointer to **[]string** | The OCI regions that this product tier is available on | [optional] 
**OnPremPlatforms** | Pointer to **[]string** | The on prem platforms that this product tier is available on | [optional] 
**PlanDescription** | Pointer to **string** | A brief description for the end user of the product tier | [optional] 
**PricePerUnit** | Pointer to **map[string]interface{}** | Price per unit. | [optional] 
**Pricing** | Pointer to **interface{}** | Pricing | [optional] 
**PrivateRegions** | Pointer to **[]string** | The private regions that this product tier is available on | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Support** | Pointer to **string** | Support | [optional] 
**TierType** | Pointer to **string** | ProductTierType is the type of tier for a product | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateProductTierRequest

`func NewUpdateProductTierRequest(id string, serviceId string, token string, ) *UpdateProductTierRequest`

NewUpdateProductTierRequest instantiates a new UpdateProductTierRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductTierRequestWithDefaults

`func NewUpdateProductTierRequestWithDefaults() *UpdateProductTierRequest`

NewUpdateProductTierRequestWithDefaults instantiates a new UpdateProductTierRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCreatesWhenPaymentNotConfigured

`func (o *UpdateProductTierRequest) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *UpdateProductTierRequest) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *UpdateProductTierRequest) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *UpdateProductTierRequest) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

### GetAutoApproveSubscription

`func (o *UpdateProductTierRequest) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *UpdateProductTierRequest) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *UpdateProductTierRequest) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *UpdateProductTierRequest) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetAwsRegions

`func (o *UpdateProductTierRequest) GetAwsRegions() []string`

GetAwsRegions returns the AwsRegions field if non-nil, zero value otherwise.

### GetAwsRegionsOk

`func (o *UpdateProductTierRequest) GetAwsRegionsOk() (*[]string, bool)`

GetAwsRegionsOk returns a tuple with the AwsRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegions

`func (o *UpdateProductTierRequest) SetAwsRegions(v []string)`

SetAwsRegions sets AwsRegions field to given value.

### HasAwsRegions

`func (o *UpdateProductTierRequest) HasAwsRegions() bool`

HasAwsRegions returns a boolean if a field has been set.

### GetAzureRegions

`func (o *UpdateProductTierRequest) GetAzureRegions() []string`

GetAzureRegions returns the AzureRegions field if non-nil, zero value otherwise.

### GetAzureRegionsOk

`func (o *UpdateProductTierRequest) GetAzureRegionsOk() (*[]string, bool)`

GetAzureRegionsOk returns a tuple with the AzureRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRegions

`func (o *UpdateProductTierRequest) SetAzureRegions(v []string)`

SetAzureRegions sets AzureRegions field to given value.

### HasAzureRegions

`func (o *UpdateProductTierRequest) HasAzureRegions() bool`

HasAzureRegions returns a boolean if a field has been set.

### GetBillingProductID

`func (o *UpdateProductTierRequest) GetBillingProductID() string`

GetBillingProductID returns the BillingProductID field if non-nil, zero value otherwise.

### GetBillingProductIDOk

`func (o *UpdateProductTierRequest) GetBillingProductIDOk() (*string, bool)`

GetBillingProductIDOk returns a tuple with the BillingProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProductID

`func (o *UpdateProductTierRequest) SetBillingProductID(v string)`

SetBillingProductID sets BillingProductID field to given value.

### HasBillingProductID

`func (o *UpdateProductTierRequest) HasBillingProductID() bool`

HasBillingProductID returns a boolean if a field has been set.

### GetBillingProviders

`func (o *UpdateProductTierRequest) GetBillingProviders() []string`

GetBillingProviders returns the BillingProviders field if non-nil, zero value otherwise.

### GetBillingProvidersOk

`func (o *UpdateProductTierRequest) GetBillingProvidersOk() (*[]string, bool)`

GetBillingProvidersOk returns a tuple with the BillingProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProviders

`func (o *UpdateProductTierRequest) SetBillingProviders(v []string)`

SetBillingProviders sets BillingProviders field to given value.

### HasBillingProviders

`func (o *UpdateProductTierRequest) HasBillingProviders() bool`

HasBillingProviders returns a boolean if a field has been set.

### GetDefaultBillingProvider

`func (o *UpdateProductTierRequest) GetDefaultBillingProvider() string`

GetDefaultBillingProvider returns the DefaultBillingProvider field if non-nil, zero value otherwise.

### GetDefaultBillingProviderOk

`func (o *UpdateProductTierRequest) GetDefaultBillingProviderOk() (*string, bool)`

GetDefaultBillingProviderOk returns a tuple with the DefaultBillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBillingProvider

`func (o *UpdateProductTierRequest) SetDefaultBillingProvider(v string)`

SetDefaultBillingProvider sets DefaultBillingProvider field to given value.

### HasDefaultBillingProvider

`func (o *UpdateProductTierRequest) HasDefaultBillingProvider() bool`

HasDefaultBillingProvider returns a boolean if a field has been set.

### GetDeploymentConfiguration

`func (o *UpdateProductTierRequest) GetDeploymentConfiguration() ProductTierDeploymentConfiguration`

GetDeploymentConfiguration returns the DeploymentConfiguration field if non-nil, zero value otherwise.

### GetDeploymentConfigurationOk

`func (o *UpdateProductTierRequest) GetDeploymentConfigurationOk() (*ProductTierDeploymentConfiguration, bool)`

GetDeploymentConfigurationOk returns a tuple with the DeploymentConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfiguration

`func (o *UpdateProductTierRequest) SetDeploymentConfiguration(v ProductTierDeploymentConfiguration)`

SetDeploymentConfiguration sets DeploymentConfiguration field to given value.

### HasDeploymentConfiguration

`func (o *UpdateProductTierRequest) HasDeploymentConfiguration() bool`

HasDeploymentConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProductTierRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProductTierRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProductTierRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProductTierRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentation

`func (o *UpdateProductTierRequest) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *UpdateProductTierRequest) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *UpdateProductTierRequest) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *UpdateProductTierRequest) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetExportUsageMetering

`func (o *UpdateProductTierRequest) GetExportUsageMetering() bool`

GetExportUsageMetering returns the ExportUsageMetering field if non-nil, zero value otherwise.

### GetExportUsageMeteringOk

`func (o *UpdateProductTierRequest) GetExportUsageMeteringOk() (*bool, bool)`

GetExportUsageMeteringOk returns a tuple with the ExportUsageMetering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUsageMetering

`func (o *UpdateProductTierRequest) SetExportUsageMetering(v bool)`

SetExportUsageMetering sets ExportUsageMetering field to given value.

### HasExportUsageMetering

`func (o *UpdateProductTierRequest) HasExportUsageMetering() bool`

HasExportUsageMetering returns a boolean if a field has been set.

### GetExportUsageMeteringConfig

`func (o *UpdateProductTierRequest) GetExportUsageMeteringConfig() map[string]interface{}`

GetExportUsageMeteringConfig returns the ExportUsageMeteringConfig field if non-nil, zero value otherwise.

### GetExportUsageMeteringConfigOk

`func (o *UpdateProductTierRequest) GetExportUsageMeteringConfigOk() (*map[string]interface{}, bool)`

GetExportUsageMeteringConfigOk returns a tuple with the ExportUsageMeteringConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUsageMeteringConfig

`func (o *UpdateProductTierRequest) SetExportUsageMeteringConfig(v map[string]interface{})`

SetExportUsageMeteringConfig sets ExportUsageMeteringConfig field to given value.

### HasExportUsageMeteringConfig

`func (o *UpdateProductTierRequest) HasExportUsageMeteringConfig() bool`

HasExportUsageMeteringConfig returns a boolean if a field has been set.

### GetGcpRegions

`func (o *UpdateProductTierRequest) GetGcpRegions() []string`

GetGcpRegions returns the GcpRegions field if non-nil, zero value otherwise.

### GetGcpRegionsOk

`func (o *UpdateProductTierRequest) GetGcpRegionsOk() (*[]string, bool)`

GetGcpRegionsOk returns a tuple with the GcpRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRegions

`func (o *UpdateProductTierRequest) SetGcpRegions(v []string)`

SetGcpRegions sets GcpRegions field to given value.

### HasGcpRegions

`func (o *UpdateProductTierRequest) HasGcpRegions() bool`

HasGcpRegions returns a boolean if a field has been set.

### GetId

`func (o *UpdateProductTierRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateProductTierRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateProductTierRequest) SetId(v string)`

SetId sets Id field to given value.


### GetIsDisabled

`func (o *UpdateProductTierRequest) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UpdateProductTierRequest) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UpdateProductTierRequest) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UpdateProductTierRequest) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetMaxNumberOfInstances

`func (o *UpdateProductTierRequest) GetMaxNumberOfInstances() int64`

GetMaxNumberOfInstances returns the MaxNumberOfInstances field if non-nil, zero value otherwise.

### GetMaxNumberOfInstancesOk

`func (o *UpdateProductTierRequest) GetMaxNumberOfInstancesOk() (*int64, bool)`

GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfInstances

`func (o *UpdateProductTierRequest) SetMaxNumberOfInstances(v int64)`

SetMaxNumberOfInstances sets MaxNumberOfInstances field to given value.

### HasMaxNumberOfInstances

`func (o *UpdateProductTierRequest) HasMaxNumberOfInstances() bool`

HasMaxNumberOfInstances returns a boolean if a field has been set.

### GetName

`func (o *UpdateProductTierRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProductTierRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProductTierRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProductTierRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOciRegions

`func (o *UpdateProductTierRequest) GetOciRegions() []string`

GetOciRegions returns the OciRegions field if non-nil, zero value otherwise.

### GetOciRegionsOk

`func (o *UpdateProductTierRequest) GetOciRegionsOk() (*[]string, bool)`

GetOciRegionsOk returns a tuple with the OciRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciRegions

`func (o *UpdateProductTierRequest) SetOciRegions(v []string)`

SetOciRegions sets OciRegions field to given value.

### HasOciRegions

`func (o *UpdateProductTierRequest) HasOciRegions() bool`

HasOciRegions returns a boolean if a field has been set.

### GetOnPremPlatforms

`func (o *UpdateProductTierRequest) GetOnPremPlatforms() []string`

GetOnPremPlatforms returns the OnPremPlatforms field if non-nil, zero value otherwise.

### GetOnPremPlatformsOk

`func (o *UpdateProductTierRequest) GetOnPremPlatformsOk() (*[]string, bool)`

GetOnPremPlatformsOk returns a tuple with the OnPremPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremPlatforms

`func (o *UpdateProductTierRequest) SetOnPremPlatforms(v []string)`

SetOnPremPlatforms sets OnPremPlatforms field to given value.

### HasOnPremPlatforms

`func (o *UpdateProductTierRequest) HasOnPremPlatforms() bool`

HasOnPremPlatforms returns a boolean if a field has been set.

### GetPlanDescription

`func (o *UpdateProductTierRequest) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *UpdateProductTierRequest) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *UpdateProductTierRequest) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *UpdateProductTierRequest) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### GetPricePerUnit

`func (o *UpdateProductTierRequest) GetPricePerUnit() map[string]interface{}`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *UpdateProductTierRequest) GetPricePerUnitOk() (*map[string]interface{}, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *UpdateProductTierRequest) SetPricePerUnit(v map[string]interface{})`

SetPricePerUnit sets PricePerUnit field to given value.

### HasPricePerUnit

`func (o *UpdateProductTierRequest) HasPricePerUnit() bool`

HasPricePerUnit returns a boolean if a field has been set.

### GetPricing

`func (o *UpdateProductTierRequest) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *UpdateProductTierRequest) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *UpdateProductTierRequest) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *UpdateProductTierRequest) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *UpdateProductTierRequest) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *UpdateProductTierRequest) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetPrivateRegions

`func (o *UpdateProductTierRequest) GetPrivateRegions() []string`

GetPrivateRegions returns the PrivateRegions field if non-nil, zero value otherwise.

### GetPrivateRegionsOk

`func (o *UpdateProductTierRequest) GetPrivateRegionsOk() (*[]string, bool)`

GetPrivateRegionsOk returns a tuple with the PrivateRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateRegions

`func (o *UpdateProductTierRequest) SetPrivateRegions(v []string)`

SetPrivateRegions sets PrivateRegions field to given value.

### HasPrivateRegions

`func (o *UpdateProductTierRequest) HasPrivateRegions() bool`

HasPrivateRegions returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateProductTierRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateProductTierRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateProductTierRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSupport

`func (o *UpdateProductTierRequest) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *UpdateProductTierRequest) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *UpdateProductTierRequest) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *UpdateProductTierRequest) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetTierType

`func (o *UpdateProductTierRequest) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *UpdateProductTierRequest) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *UpdateProductTierRequest) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *UpdateProductTierRequest) HasTierType() bool`

HasTierType returns a boolean if a field has been set.

### GetToken

`func (o *UpdateProductTierRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateProductTierRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateProductTierRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


