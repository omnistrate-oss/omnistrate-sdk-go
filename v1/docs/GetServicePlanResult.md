# GetServicePlanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | **bool** | Auto approve subscription or not | 
**AccountConfigIds** | Pointer to **[]string** | The infrastructure account configuration ID list | [optional] 
**ActiveAccountConfigIds** | Pointer to **map[string]string** | The active infrastructure account configuration IDs per cloud provider | [optional] 
**ApiVersion** | **string** | The external version of the API | 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this service plan is available on | [optional] 
**AzureRegions** | Pointer to **[]string** | The Azure regions that this service plan is available on | [optional] 
**DeploymentConfigId** | **string** | ID of a Deployment Config | 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this service plan is available on | [optional] 
**HasPendingChanges** | Pointer to **bool** | Whether there are any pending changes for the product tier configuration | [optional] 
**IsProductTierDisabled** | **bool** | Whether the product tier is disabled | 
**LatestMajorVersion** | **string** | The version number for the latest major version set. | 
**ModelType** | **string** | The model type encapsulating this service | 
**PrivateRegions** | Pointer to **[]string** | The Private regions that this service plan is available on | [optional] 
**ProductTierDescription** | **string** | A brief description of the product tier | 
**ProductTierDocumentation** | **string** | Documentation | 
**ProductTierFeatures** | Pointer to **map[string]bool** | The features that are enabled / disabled for this product tier | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProductTierKey** | **string** | Unique Key of the product tier | 
**ProductTierName** | **string** | Name of the product tier | 
**ProductTierPlanDescription** | **string** | A brief description for the end user of the product tier | 
**ProductTierPricing** | **interface{}** | Pricing | 
**ProductTierSupport** | **string** | Support | 
**ServiceApiDescription** | **string** | A brief description of the service API bundle | 
**ServiceApiId** | **string** | ID of a Service API | 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceModelDescription** | **string** | A brief description of the service model | 
**ServiceModelFeatures** | Pointer to [**[]ServiceModelFeatureDetail**](ServiceModelFeatureDetail.md) | Enabled service model features | [optional] 
**ServiceModelId** | **string** | ID of a Service Model | 
**ServiceModelName** | **string** | Name of the Service Model | 
**TierType** | **string** | ProductTierType is the type of tier for a product | 
**VersionSetStatus** | **string** | The tier version set status. | 

## Methods

### NewGetServicePlanResult

`func NewGetServicePlanResult(autoApproveSubscription bool, apiVersion string, deploymentConfigId string, isProductTierDisabled bool, latestMajorVersion string, modelType string, productTierDescription string, productTierDocumentation string, productTierId string, productTierKey string, productTierName string, productTierPlanDescription string, productTierPricing interface{}, productTierSupport string, serviceApiDescription string, serviceApiId string, serviceEnvironmentId string, serviceModelDescription string, serviceModelId string, serviceModelName string, tierType string, versionSetStatus string, ) *GetServicePlanResult`

NewGetServicePlanResult instantiates a new GetServicePlanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServicePlanResultWithDefaults

`func NewGetServicePlanResultWithDefaults() *GetServicePlanResult`

NewGetServicePlanResultWithDefaults instantiates a new GetServicePlanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *GetServicePlanResult) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *GetServicePlanResult) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *GetServicePlanResult) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.


### GetAccountConfigIds

`func (o *GetServicePlanResult) GetAccountConfigIds() []string`

GetAccountConfigIds returns the AccountConfigIds field if non-nil, zero value otherwise.

### GetAccountConfigIdsOk

`func (o *GetServicePlanResult) GetAccountConfigIdsOk() (*[]string, bool)`

GetAccountConfigIdsOk returns a tuple with the AccountConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigIds

`func (o *GetServicePlanResult) SetAccountConfigIds(v []string)`

SetAccountConfigIds sets AccountConfigIds field to given value.

### HasAccountConfigIds

`func (o *GetServicePlanResult) HasAccountConfigIds() bool`

HasAccountConfigIds returns a boolean if a field has been set.

### GetActiveAccountConfigIds

`func (o *GetServicePlanResult) GetActiveAccountConfigIds() map[string]string`

GetActiveAccountConfigIds returns the ActiveAccountConfigIds field if non-nil, zero value otherwise.

### GetActiveAccountConfigIdsOk

`func (o *GetServicePlanResult) GetActiveAccountConfigIdsOk() (*map[string]string, bool)`

GetActiveAccountConfigIdsOk returns a tuple with the ActiveAccountConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAccountConfigIds

`func (o *GetServicePlanResult) SetActiveAccountConfigIds(v map[string]string)`

SetActiveAccountConfigIds sets ActiveAccountConfigIds field to given value.

### HasActiveAccountConfigIds

`func (o *GetServicePlanResult) HasActiveAccountConfigIds() bool`

HasActiveAccountConfigIds returns a boolean if a field has been set.

### GetApiVersion

`func (o *GetServicePlanResult) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetServicePlanResult) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetServicePlanResult) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetAwsRegions

`func (o *GetServicePlanResult) GetAwsRegions() []string`

GetAwsRegions returns the AwsRegions field if non-nil, zero value otherwise.

### GetAwsRegionsOk

`func (o *GetServicePlanResult) GetAwsRegionsOk() (*[]string, bool)`

GetAwsRegionsOk returns a tuple with the AwsRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegions

`func (o *GetServicePlanResult) SetAwsRegions(v []string)`

SetAwsRegions sets AwsRegions field to given value.

### HasAwsRegions

`func (o *GetServicePlanResult) HasAwsRegions() bool`

HasAwsRegions returns a boolean if a field has been set.

### GetAzureRegions

`func (o *GetServicePlanResult) GetAzureRegions() []string`

GetAzureRegions returns the AzureRegions field if non-nil, zero value otherwise.

### GetAzureRegionsOk

`func (o *GetServicePlanResult) GetAzureRegionsOk() (*[]string, bool)`

GetAzureRegionsOk returns a tuple with the AzureRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRegions

`func (o *GetServicePlanResult) SetAzureRegions(v []string)`

SetAzureRegions sets AzureRegions field to given value.

### HasAzureRegions

`func (o *GetServicePlanResult) HasAzureRegions() bool`

HasAzureRegions returns a boolean if a field has been set.

### GetDeploymentConfigId

`func (o *GetServicePlanResult) GetDeploymentConfigId() string`

GetDeploymentConfigId returns the DeploymentConfigId field if non-nil, zero value otherwise.

### GetDeploymentConfigIdOk

`func (o *GetServicePlanResult) GetDeploymentConfigIdOk() (*string, bool)`

GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfigId

`func (o *GetServicePlanResult) SetDeploymentConfigId(v string)`

SetDeploymentConfigId sets DeploymentConfigId field to given value.


### GetGcpRegions

`func (o *GetServicePlanResult) GetGcpRegions() []string`

GetGcpRegions returns the GcpRegions field if non-nil, zero value otherwise.

### GetGcpRegionsOk

`func (o *GetServicePlanResult) GetGcpRegionsOk() (*[]string, bool)`

GetGcpRegionsOk returns a tuple with the GcpRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRegions

`func (o *GetServicePlanResult) SetGcpRegions(v []string)`

SetGcpRegions sets GcpRegions field to given value.

### HasGcpRegions

`func (o *GetServicePlanResult) HasGcpRegions() bool`

HasGcpRegions returns a boolean if a field has been set.

### GetHasPendingChanges

`func (o *GetServicePlanResult) GetHasPendingChanges() bool`

GetHasPendingChanges returns the HasPendingChanges field if non-nil, zero value otherwise.

### GetHasPendingChangesOk

`func (o *GetServicePlanResult) GetHasPendingChangesOk() (*bool, bool)`

GetHasPendingChangesOk returns a tuple with the HasPendingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingChanges

`func (o *GetServicePlanResult) SetHasPendingChanges(v bool)`

SetHasPendingChanges sets HasPendingChanges field to given value.

### HasHasPendingChanges

`func (o *GetServicePlanResult) HasHasPendingChanges() bool`

HasHasPendingChanges returns a boolean if a field has been set.

### GetIsProductTierDisabled

`func (o *GetServicePlanResult) GetIsProductTierDisabled() bool`

GetIsProductTierDisabled returns the IsProductTierDisabled field if non-nil, zero value otherwise.

### GetIsProductTierDisabledOk

`func (o *GetServicePlanResult) GetIsProductTierDisabledOk() (*bool, bool)`

GetIsProductTierDisabledOk returns a tuple with the IsProductTierDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProductTierDisabled

`func (o *GetServicePlanResult) SetIsProductTierDisabled(v bool)`

SetIsProductTierDisabled sets IsProductTierDisabled field to given value.


### GetLatestMajorVersion

`func (o *GetServicePlanResult) GetLatestMajorVersion() string`

GetLatestMajorVersion returns the LatestMajorVersion field if non-nil, zero value otherwise.

### GetLatestMajorVersionOk

`func (o *GetServicePlanResult) GetLatestMajorVersionOk() (*string, bool)`

GetLatestMajorVersionOk returns a tuple with the LatestMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMajorVersion

`func (o *GetServicePlanResult) SetLatestMajorVersion(v string)`

SetLatestMajorVersion sets LatestMajorVersion field to given value.


### GetModelType

`func (o *GetServicePlanResult) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *GetServicePlanResult) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *GetServicePlanResult) SetModelType(v string)`

SetModelType sets ModelType field to given value.


### GetPrivateRegions

`func (o *GetServicePlanResult) GetPrivateRegions() []string`

GetPrivateRegions returns the PrivateRegions field if non-nil, zero value otherwise.

### GetPrivateRegionsOk

`func (o *GetServicePlanResult) GetPrivateRegionsOk() (*[]string, bool)`

GetPrivateRegionsOk returns a tuple with the PrivateRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateRegions

`func (o *GetServicePlanResult) SetPrivateRegions(v []string)`

SetPrivateRegions sets PrivateRegions field to given value.

### HasPrivateRegions

`func (o *GetServicePlanResult) HasPrivateRegions() bool`

HasPrivateRegions returns a boolean if a field has been set.

### GetProductTierDescription

`func (o *GetServicePlanResult) GetProductTierDescription() string`

GetProductTierDescription returns the ProductTierDescription field if non-nil, zero value otherwise.

### GetProductTierDescriptionOk

`func (o *GetServicePlanResult) GetProductTierDescriptionOk() (*string, bool)`

GetProductTierDescriptionOk returns a tuple with the ProductTierDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierDescription

`func (o *GetServicePlanResult) SetProductTierDescription(v string)`

SetProductTierDescription sets ProductTierDescription field to given value.


### GetProductTierDocumentation

`func (o *GetServicePlanResult) GetProductTierDocumentation() string`

GetProductTierDocumentation returns the ProductTierDocumentation field if non-nil, zero value otherwise.

### GetProductTierDocumentationOk

`func (o *GetServicePlanResult) GetProductTierDocumentationOk() (*string, bool)`

GetProductTierDocumentationOk returns a tuple with the ProductTierDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierDocumentation

`func (o *GetServicePlanResult) SetProductTierDocumentation(v string)`

SetProductTierDocumentation sets ProductTierDocumentation field to given value.


### GetProductTierFeatures

`func (o *GetServicePlanResult) GetProductTierFeatures() map[string]bool`

GetProductTierFeatures returns the ProductTierFeatures field if non-nil, zero value otherwise.

### GetProductTierFeaturesOk

`func (o *GetServicePlanResult) GetProductTierFeaturesOk() (*map[string]bool, bool)`

GetProductTierFeaturesOk returns a tuple with the ProductTierFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierFeatures

`func (o *GetServicePlanResult) SetProductTierFeatures(v map[string]bool)`

SetProductTierFeatures sets ProductTierFeatures field to given value.

### HasProductTierFeatures

`func (o *GetServicePlanResult) HasProductTierFeatures() bool`

HasProductTierFeatures returns a boolean if a field has been set.

### GetProductTierId

`func (o *GetServicePlanResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *GetServicePlanResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *GetServicePlanResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierKey

`func (o *GetServicePlanResult) GetProductTierKey() string`

GetProductTierKey returns the ProductTierKey field if non-nil, zero value otherwise.

### GetProductTierKeyOk

`func (o *GetServicePlanResult) GetProductTierKeyOk() (*string, bool)`

GetProductTierKeyOk returns a tuple with the ProductTierKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierKey

`func (o *GetServicePlanResult) SetProductTierKey(v string)`

SetProductTierKey sets ProductTierKey field to given value.


### GetProductTierName

`func (o *GetServicePlanResult) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *GetServicePlanResult) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *GetServicePlanResult) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetProductTierPlanDescription

`func (o *GetServicePlanResult) GetProductTierPlanDescription() string`

GetProductTierPlanDescription returns the ProductTierPlanDescription field if non-nil, zero value otherwise.

### GetProductTierPlanDescriptionOk

`func (o *GetServicePlanResult) GetProductTierPlanDescriptionOk() (*string, bool)`

GetProductTierPlanDescriptionOk returns a tuple with the ProductTierPlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierPlanDescription

`func (o *GetServicePlanResult) SetProductTierPlanDescription(v string)`

SetProductTierPlanDescription sets ProductTierPlanDescription field to given value.


### GetProductTierPricing

`func (o *GetServicePlanResult) GetProductTierPricing() interface{}`

GetProductTierPricing returns the ProductTierPricing field if non-nil, zero value otherwise.

### GetProductTierPricingOk

`func (o *GetServicePlanResult) GetProductTierPricingOk() (*interface{}, bool)`

GetProductTierPricingOk returns a tuple with the ProductTierPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierPricing

`func (o *GetServicePlanResult) SetProductTierPricing(v interface{})`

SetProductTierPricing sets ProductTierPricing field to given value.


### SetProductTierPricingNil

`func (o *GetServicePlanResult) SetProductTierPricingNil(b bool)`

 SetProductTierPricingNil sets the value for ProductTierPricing to be an explicit nil

### UnsetProductTierPricing
`func (o *GetServicePlanResult) UnsetProductTierPricing()`

UnsetProductTierPricing ensures that no value is present for ProductTierPricing, not even an explicit nil
### GetProductTierSupport

`func (o *GetServicePlanResult) GetProductTierSupport() string`

GetProductTierSupport returns the ProductTierSupport field if non-nil, zero value otherwise.

### GetProductTierSupportOk

`func (o *GetServicePlanResult) GetProductTierSupportOk() (*string, bool)`

GetProductTierSupportOk returns a tuple with the ProductTierSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierSupport

`func (o *GetServicePlanResult) SetProductTierSupport(v string)`

SetProductTierSupport sets ProductTierSupport field to given value.


### GetServiceApiDescription

`func (o *GetServicePlanResult) GetServiceApiDescription() string`

GetServiceApiDescription returns the ServiceApiDescription field if non-nil, zero value otherwise.

### GetServiceApiDescriptionOk

`func (o *GetServicePlanResult) GetServiceApiDescriptionOk() (*string, bool)`

GetServiceApiDescriptionOk returns a tuple with the ServiceApiDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiDescription

`func (o *GetServicePlanResult) SetServiceApiDescription(v string)`

SetServiceApiDescription sets ServiceApiDescription field to given value.


### GetServiceApiId

`func (o *GetServicePlanResult) GetServiceApiId() string`

GetServiceApiId returns the ServiceApiId field if non-nil, zero value otherwise.

### GetServiceApiIdOk

`func (o *GetServicePlanResult) GetServiceApiIdOk() (*string, bool)`

GetServiceApiIdOk returns a tuple with the ServiceApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiId

`func (o *GetServicePlanResult) SetServiceApiId(v string)`

SetServiceApiId sets ServiceApiId field to given value.


### GetServiceEnvironmentId

`func (o *GetServicePlanResult) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *GetServicePlanResult) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *GetServicePlanResult) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceModelDescription

`func (o *GetServicePlanResult) GetServiceModelDescription() string`

GetServiceModelDescription returns the ServiceModelDescription field if non-nil, zero value otherwise.

### GetServiceModelDescriptionOk

`func (o *GetServicePlanResult) GetServiceModelDescriptionOk() (*string, bool)`

GetServiceModelDescriptionOk returns a tuple with the ServiceModelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelDescription

`func (o *GetServicePlanResult) SetServiceModelDescription(v string)`

SetServiceModelDescription sets ServiceModelDescription field to given value.


### GetServiceModelFeatures

`func (o *GetServicePlanResult) GetServiceModelFeatures() []ServiceModelFeatureDetail`

GetServiceModelFeatures returns the ServiceModelFeatures field if non-nil, zero value otherwise.

### GetServiceModelFeaturesOk

`func (o *GetServicePlanResult) GetServiceModelFeaturesOk() (*[]ServiceModelFeatureDetail, bool)`

GetServiceModelFeaturesOk returns a tuple with the ServiceModelFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelFeatures

`func (o *GetServicePlanResult) SetServiceModelFeatures(v []ServiceModelFeatureDetail)`

SetServiceModelFeatures sets ServiceModelFeatures field to given value.

### HasServiceModelFeatures

`func (o *GetServicePlanResult) HasServiceModelFeatures() bool`

HasServiceModelFeatures returns a boolean if a field has been set.

### GetServiceModelId

`func (o *GetServicePlanResult) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *GetServicePlanResult) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *GetServicePlanResult) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetServiceModelName

`func (o *GetServicePlanResult) GetServiceModelName() string`

GetServiceModelName returns the ServiceModelName field if non-nil, zero value otherwise.

### GetServiceModelNameOk

`func (o *GetServicePlanResult) GetServiceModelNameOk() (*string, bool)`

GetServiceModelNameOk returns a tuple with the ServiceModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelName

`func (o *GetServicePlanResult) SetServiceModelName(v string)`

SetServiceModelName sets ServiceModelName field to given value.


### GetTierType

`func (o *GetServicePlanResult) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *GetServicePlanResult) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *GetServicePlanResult) SetTierType(v string)`

SetTierType sets TierType field to given value.


### GetVersionSetStatus

`func (o *GetServicePlanResult) GetVersionSetStatus() string`

GetVersionSetStatus returns the VersionSetStatus field if non-nil, zero value otherwise.

### GetVersionSetStatusOk

`func (o *GetServicePlanResult) GetVersionSetStatusOk() (*string, bool)`

GetVersionSetStatusOk returns a tuple with the VersionSetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSetStatus

`func (o *GetServicePlanResult) SetVersionSetStatus(v string)`

SetVersionSetStatus sets VersionSetStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


