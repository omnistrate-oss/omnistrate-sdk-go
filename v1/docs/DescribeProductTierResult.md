# DescribeProductTierResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroups** | Pointer to **map[string]interface{}** | The resources that belong to this service API bundle and their active versions | [optional] 
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this product tier is available on | [optional] 
**CloudProvidersConfigReadiness** | Pointer to **map[string]interface{}** | The readiness of the cloud providers configurations | [optional] 
**Description** | **string** | A brief description of the product tier | 
**Documentation** | **string** | Documentation | 
**EnabledFeatures** | Pointer to [**[]ProductTierFeatureDetail**](ProductTierFeatureDetail.md) | The features that are enabled for this product tier, including scope details and configuration | [optional] 
**Features** | Pointer to **map[string]interface{}** | The features that are enabled / disabled for this product tier | [optional] 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this product tier is available on | [optional] 
**Id** | **string** | ID of a Product Tier | 
**IsDisabled** | **bool** | Flag to indicate if the product tier is disabled. | 
**Key** | **string** | Unique Key of the product tier | 
**Name** | **string** | Name of the product tier | 
**PlanDescription** | **string** | A brief description for the end user of the product tier | 
**Pricing** | **interface{}** | Pricing | 
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


