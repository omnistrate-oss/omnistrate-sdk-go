# UpdateProductTierRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Allow creates when payment not configured | [optional] 
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this product tier is available on | [optional] 
**AzureRegions** | Pointer to **[]string** | The Azure regions that this product tier is available on | [optional] 
**Description** | Pointer to **string** | A brief description of the product tier | [optional] 
**Documentation** | Pointer to **string** | Documentation | [optional] 
**ExportUsageMetering** | Pointer to **bool** | Export usage metering data | [optional] 
**ExportUsageMeteringConfig** | Pointer to **map[string]interface{}** | Export usage metering data configuration | [optional] 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this product tier is available on | [optional] 
**Id** | **string** | ID of a Product Tier | 
**IsDisabled** | Pointer to **bool** | Update the product tier&#39;s state as enabled/disabled. Enabling the product tier will let end-customers subscribe and use the service plan. | [optional] 
**MaxNumberOfInstances** | Pointer to **int64** | Maximum number of instances. Set to -1 for unlimited. | [optional] 
**Name** | Pointer to **string** | Name of the product tier | [optional] 
**PlanDescription** | Pointer to **string** | A brief description for the end user of the product tier | [optional] 
**PricePerUnit** | Pointer to **map[string]interface{}** | Price per unit. | [optional] 
**Pricing** | Pointer to **interface{}** | Pricing | [optional] 
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


