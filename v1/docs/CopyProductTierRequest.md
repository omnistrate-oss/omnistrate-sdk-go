# CopyProductTierRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this product tier is available on | [optional] 
**AzureRegions** | Pointer to **[]string** | The Azure regions that this product tier is available on | [optional] 
**Description** | **string** | A brief description of the product tier | 
**Documentation** | Pointer to **string** | Documentation | [optional] 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this product tier is available on | [optional] 
**Name** | **string** | Name of the product tier | 
**PlanDescription** | Pointer to **string** | A brief description for the end user of the product tier | [optional] 
**Pricing** | Pointer to **interface{}** | Pricing | [optional] 
**ServiceId** | **string** | ID of a Service | 
**ServiceModelId** | **string** | ID of a Service Model | 
**SourceId** | **string** | ID of a Product Tier | 
**Support** | Pointer to **string** | Support | [optional] 
**TargetTierType** | Pointer to **string** | ProductTierType is the type of tier for a product | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCopyProductTierRequest

`func NewCopyProductTierRequest(description string, name string, serviceId string, serviceModelId string, sourceId string, token string, ) *CopyProductTierRequest`

NewCopyProductTierRequest instantiates a new CopyProductTierRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyProductTierRequestWithDefaults

`func NewCopyProductTierRequestWithDefaults() *CopyProductTierRequest`

NewCopyProductTierRequestWithDefaults instantiates a new CopyProductTierRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *CopyProductTierRequest) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *CopyProductTierRequest) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *CopyProductTierRequest) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *CopyProductTierRequest) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetAwsRegions

`func (o *CopyProductTierRequest) GetAwsRegions() []string`

GetAwsRegions returns the AwsRegions field if non-nil, zero value otherwise.

### GetAwsRegionsOk

`func (o *CopyProductTierRequest) GetAwsRegionsOk() (*[]string, bool)`

GetAwsRegionsOk returns a tuple with the AwsRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegions

`func (o *CopyProductTierRequest) SetAwsRegions(v []string)`

SetAwsRegions sets AwsRegions field to given value.

### HasAwsRegions

`func (o *CopyProductTierRequest) HasAwsRegions() bool`

HasAwsRegions returns a boolean if a field has been set.

### GetAzureRegions

`func (o *CopyProductTierRequest) GetAzureRegions() []string`

GetAzureRegions returns the AzureRegions field if non-nil, zero value otherwise.

### GetAzureRegionsOk

`func (o *CopyProductTierRequest) GetAzureRegionsOk() (*[]string, bool)`

GetAzureRegionsOk returns a tuple with the AzureRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRegions

`func (o *CopyProductTierRequest) SetAzureRegions(v []string)`

SetAzureRegions sets AzureRegions field to given value.

### HasAzureRegions

`func (o *CopyProductTierRequest) HasAzureRegions() bool`

HasAzureRegions returns a boolean if a field has been set.

### GetDescription

`func (o *CopyProductTierRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CopyProductTierRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CopyProductTierRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDocumentation

`func (o *CopyProductTierRequest) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *CopyProductTierRequest) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *CopyProductTierRequest) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *CopyProductTierRequest) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetGcpRegions

`func (o *CopyProductTierRequest) GetGcpRegions() []string`

GetGcpRegions returns the GcpRegions field if non-nil, zero value otherwise.

### GetGcpRegionsOk

`func (o *CopyProductTierRequest) GetGcpRegionsOk() (*[]string, bool)`

GetGcpRegionsOk returns a tuple with the GcpRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRegions

`func (o *CopyProductTierRequest) SetGcpRegions(v []string)`

SetGcpRegions sets GcpRegions field to given value.

### HasGcpRegions

`func (o *CopyProductTierRequest) HasGcpRegions() bool`

HasGcpRegions returns a boolean if a field has been set.

### GetName

`func (o *CopyProductTierRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CopyProductTierRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CopyProductTierRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPlanDescription

`func (o *CopyProductTierRequest) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *CopyProductTierRequest) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *CopyProductTierRequest) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *CopyProductTierRequest) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### GetPricing

`func (o *CopyProductTierRequest) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *CopyProductTierRequest) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *CopyProductTierRequest) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *CopyProductTierRequest) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *CopyProductTierRequest) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *CopyProductTierRequest) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetServiceId

`func (o *CopyProductTierRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CopyProductTierRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CopyProductTierRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceModelId

`func (o *CopyProductTierRequest) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *CopyProductTierRequest) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *CopyProductTierRequest) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetSourceId

`func (o *CopyProductTierRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CopyProductTierRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CopyProductTierRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSupport

`func (o *CopyProductTierRequest) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *CopyProductTierRequest) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *CopyProductTierRequest) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *CopyProductTierRequest) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetTargetTierType

`func (o *CopyProductTierRequest) GetTargetTierType() string`

GetTargetTierType returns the TargetTierType field if non-nil, zero value otherwise.

### GetTargetTierTypeOk

`func (o *CopyProductTierRequest) GetTargetTierTypeOk() (*string, bool)`

GetTargetTierTypeOk returns a tuple with the TargetTierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTierType

`func (o *CopyProductTierRequest) SetTargetTierType(v string)`

SetTargetTierType sets TargetTierType field to given value.

### HasTargetTierType

`func (o *CopyProductTierRequest) HasTargetTierType() bool`

HasTargetTierType returns a boolean if a field has been set.

### GetToken

`func (o *CopyProductTierRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CopyProductTierRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CopyProductTierRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


