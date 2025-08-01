# ServiceOffering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | **bool** | Auto approve subscription or not | 
**AllowCreatesWhenPaymentNotConfigured** | Pointer to **bool** | Allow creates when payment not configured | [optional] 
**Assets** | Pointer to [**ServiceAssets**](ServiceAssets.md) |  | [optional] 
**AwsRegions** | Pointer to **[]string** | The AWS regions that this service offering is available on | [optional] 
**AzureRegions** | Pointer to **[]string** | The Azure regions that this service offering is available on | [optional] 
**CloudProviders** | Pointer to **[]string** | List of supported cloud providers for this product tier. | [optional] 
**GcpRegions** | Pointer to **[]string** | The GCP regions that this service offering is available on | [optional] 
**MaxNumberOfInstances** | Pointer to **int64** | Maximum number of instances | [optional] 
**PrivateRegions** | Pointer to **[]string** | The Private regions that this service offering is available on | [optional] 
**ProductTierDescription** | Pointer to **string** | A brief description of the product tier | [optional] 
**ProductTierDocumentation** | **string** | Documentation | 
**ProductTierID** | **string** | ID of a Product Tier | 
**ProductTierName** | **string** | The product tier name | 
**ProductTierPlanDescription** | Pointer to **string** | A brief description for the end user of the product tier | [optional] 
**ProductTierPricing** | **interface{}** | Pricing | 
**ProductTierSupport** | **string** | Support | 
**ProductTierType** | **string** | ProductTierType is the type of tier for a product | 
**ProductTierURLKey** | **string** | The product tier URL key | 
**ProductTierVersion** | **string** | The product tier version | 
**ResourceParameters** | [**[]ResourceEntity**](ResourceEntity.md) | The resource parameters | 
**ServiceAPIID** | **string** | ID of a Service API | 
**ServiceAPIVersion** | **string** | The service API version | 
**ServiceEnvironmentID** | **string** | ID of a Service Environment | 
**ServiceEnvironmentName** | **string** | The service environment name | 
**ServiceEnvironmentType** | **string** | The type of service environment | 
**ServiceEnvironmentURLKey** | **string** | The service environment URL key | 
**ServiceEnvironmentVisibility** | **string** | This parameter is used to configure the visibility of the service control-plane APIs | 
**ServiceLogoURL** | **string** | The logo for the service | 
**ServiceModelFeatures** | Pointer to [**[]ServiceModelFeatureDetail**](ServiceModelFeatureDetail.md) | Enabled service model features | [optional] 
**ServiceModelID** | **string** | ID of a Service Model | 
**ServiceModelName** | **string** | The service model name | 
**ServiceModelStatus** | **string** | The service model status | 
**ServiceModelType** | **string** | The model type encapsulating this service | 
**ServiceModelURLKey** | **string** | The service model URL key | 
**SupportsPublicNetwork** | Pointer to **bool** | Indicates whether any of the resources in the product tier support public network | [optional] 

## Methods

### NewServiceOffering

`func NewServiceOffering(autoApproveSubscription bool, productTierDocumentation string, productTierID string, productTierName string, productTierPricing interface{}, productTierSupport string, productTierType string, productTierURLKey string, productTierVersion string, resourceParameters []ResourceEntity, serviceAPIID string, serviceAPIVersion string, serviceEnvironmentID string, serviceEnvironmentName string, serviceEnvironmentType string, serviceEnvironmentURLKey string, serviceEnvironmentVisibility string, serviceLogoURL string, serviceModelID string, serviceModelName string, serviceModelStatus string, serviceModelType string, serviceModelURLKey string, ) *ServiceOffering`

NewServiceOffering instantiates a new ServiceOffering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOfferingWithDefaults

`func NewServiceOfferingWithDefaults() *ServiceOffering`

NewServiceOfferingWithDefaults instantiates a new ServiceOffering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *ServiceOffering) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *ServiceOffering) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *ServiceOffering) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.


### GetAllowCreatesWhenPaymentNotConfigured

`func (o *ServiceOffering) GetAllowCreatesWhenPaymentNotConfigured() bool`

GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field if non-nil, zero value otherwise.

### GetAllowCreatesWhenPaymentNotConfiguredOk

`func (o *ServiceOffering) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool)`

GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreatesWhenPaymentNotConfigured

`func (o *ServiceOffering) SetAllowCreatesWhenPaymentNotConfigured(v bool)`

SetAllowCreatesWhenPaymentNotConfigured sets AllowCreatesWhenPaymentNotConfigured field to given value.

### HasAllowCreatesWhenPaymentNotConfigured

`func (o *ServiceOffering) HasAllowCreatesWhenPaymentNotConfigured() bool`

HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.

### GetAssets

`func (o *ServiceOffering) GetAssets() ServiceAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *ServiceOffering) GetAssetsOk() (*ServiceAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *ServiceOffering) SetAssets(v ServiceAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *ServiceOffering) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetAwsRegions

`func (o *ServiceOffering) GetAwsRegions() []string`

GetAwsRegions returns the AwsRegions field if non-nil, zero value otherwise.

### GetAwsRegionsOk

`func (o *ServiceOffering) GetAwsRegionsOk() (*[]string, bool)`

GetAwsRegionsOk returns a tuple with the AwsRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegions

`func (o *ServiceOffering) SetAwsRegions(v []string)`

SetAwsRegions sets AwsRegions field to given value.

### HasAwsRegions

`func (o *ServiceOffering) HasAwsRegions() bool`

HasAwsRegions returns a boolean if a field has been set.

### GetAzureRegions

`func (o *ServiceOffering) GetAzureRegions() []string`

GetAzureRegions returns the AzureRegions field if non-nil, zero value otherwise.

### GetAzureRegionsOk

`func (o *ServiceOffering) GetAzureRegionsOk() (*[]string, bool)`

GetAzureRegionsOk returns a tuple with the AzureRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRegions

`func (o *ServiceOffering) SetAzureRegions(v []string)`

SetAzureRegions sets AzureRegions field to given value.

### HasAzureRegions

`func (o *ServiceOffering) HasAzureRegions() bool`

HasAzureRegions returns a boolean if a field has been set.

### GetCloudProviders

`func (o *ServiceOffering) GetCloudProviders() []string`

GetCloudProviders returns the CloudProviders field if non-nil, zero value otherwise.

### GetCloudProvidersOk

`func (o *ServiceOffering) GetCloudProvidersOk() (*[]string, bool)`

GetCloudProvidersOk returns a tuple with the CloudProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviders

`func (o *ServiceOffering) SetCloudProviders(v []string)`

SetCloudProviders sets CloudProviders field to given value.

### HasCloudProviders

`func (o *ServiceOffering) HasCloudProviders() bool`

HasCloudProviders returns a boolean if a field has been set.

### GetGcpRegions

`func (o *ServiceOffering) GetGcpRegions() []string`

GetGcpRegions returns the GcpRegions field if non-nil, zero value otherwise.

### GetGcpRegionsOk

`func (o *ServiceOffering) GetGcpRegionsOk() (*[]string, bool)`

GetGcpRegionsOk returns a tuple with the GcpRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRegions

`func (o *ServiceOffering) SetGcpRegions(v []string)`

SetGcpRegions sets GcpRegions field to given value.

### HasGcpRegions

`func (o *ServiceOffering) HasGcpRegions() bool`

HasGcpRegions returns a boolean if a field has been set.

### GetMaxNumberOfInstances

`func (o *ServiceOffering) GetMaxNumberOfInstances() int64`

GetMaxNumberOfInstances returns the MaxNumberOfInstances field if non-nil, zero value otherwise.

### GetMaxNumberOfInstancesOk

`func (o *ServiceOffering) GetMaxNumberOfInstancesOk() (*int64, bool)`

GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfInstances

`func (o *ServiceOffering) SetMaxNumberOfInstances(v int64)`

SetMaxNumberOfInstances sets MaxNumberOfInstances field to given value.

### HasMaxNumberOfInstances

`func (o *ServiceOffering) HasMaxNumberOfInstances() bool`

HasMaxNumberOfInstances returns a boolean if a field has been set.

### GetPrivateRegions

`func (o *ServiceOffering) GetPrivateRegions() []string`

GetPrivateRegions returns the PrivateRegions field if non-nil, zero value otherwise.

### GetPrivateRegionsOk

`func (o *ServiceOffering) GetPrivateRegionsOk() (*[]string, bool)`

GetPrivateRegionsOk returns a tuple with the PrivateRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateRegions

`func (o *ServiceOffering) SetPrivateRegions(v []string)`

SetPrivateRegions sets PrivateRegions field to given value.

### HasPrivateRegions

`func (o *ServiceOffering) HasPrivateRegions() bool`

HasPrivateRegions returns a boolean if a field has been set.

### GetProductTierDescription

`func (o *ServiceOffering) GetProductTierDescription() string`

GetProductTierDescription returns the ProductTierDescription field if non-nil, zero value otherwise.

### GetProductTierDescriptionOk

`func (o *ServiceOffering) GetProductTierDescriptionOk() (*string, bool)`

GetProductTierDescriptionOk returns a tuple with the ProductTierDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierDescription

`func (o *ServiceOffering) SetProductTierDescription(v string)`

SetProductTierDescription sets ProductTierDescription field to given value.

### HasProductTierDescription

`func (o *ServiceOffering) HasProductTierDescription() bool`

HasProductTierDescription returns a boolean if a field has been set.

### GetProductTierDocumentation

`func (o *ServiceOffering) GetProductTierDocumentation() string`

GetProductTierDocumentation returns the ProductTierDocumentation field if non-nil, zero value otherwise.

### GetProductTierDocumentationOk

`func (o *ServiceOffering) GetProductTierDocumentationOk() (*string, bool)`

GetProductTierDocumentationOk returns a tuple with the ProductTierDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierDocumentation

`func (o *ServiceOffering) SetProductTierDocumentation(v string)`

SetProductTierDocumentation sets ProductTierDocumentation field to given value.


### GetProductTierID

`func (o *ServiceOffering) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *ServiceOffering) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *ServiceOffering) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.


### GetProductTierName

`func (o *ServiceOffering) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *ServiceOffering) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *ServiceOffering) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetProductTierPlanDescription

`func (o *ServiceOffering) GetProductTierPlanDescription() string`

GetProductTierPlanDescription returns the ProductTierPlanDescription field if non-nil, zero value otherwise.

### GetProductTierPlanDescriptionOk

`func (o *ServiceOffering) GetProductTierPlanDescriptionOk() (*string, bool)`

GetProductTierPlanDescriptionOk returns a tuple with the ProductTierPlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierPlanDescription

`func (o *ServiceOffering) SetProductTierPlanDescription(v string)`

SetProductTierPlanDescription sets ProductTierPlanDescription field to given value.

### HasProductTierPlanDescription

`func (o *ServiceOffering) HasProductTierPlanDescription() bool`

HasProductTierPlanDescription returns a boolean if a field has been set.

### GetProductTierPricing

`func (o *ServiceOffering) GetProductTierPricing() interface{}`

GetProductTierPricing returns the ProductTierPricing field if non-nil, zero value otherwise.

### GetProductTierPricingOk

`func (o *ServiceOffering) GetProductTierPricingOk() (*interface{}, bool)`

GetProductTierPricingOk returns a tuple with the ProductTierPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierPricing

`func (o *ServiceOffering) SetProductTierPricing(v interface{})`

SetProductTierPricing sets ProductTierPricing field to given value.


### SetProductTierPricingNil

`func (o *ServiceOffering) SetProductTierPricingNil(b bool)`

 SetProductTierPricingNil sets the value for ProductTierPricing to be an explicit nil

### UnsetProductTierPricing
`func (o *ServiceOffering) UnsetProductTierPricing()`

UnsetProductTierPricing ensures that no value is present for ProductTierPricing, not even an explicit nil
### GetProductTierSupport

`func (o *ServiceOffering) GetProductTierSupport() string`

GetProductTierSupport returns the ProductTierSupport field if non-nil, zero value otherwise.

### GetProductTierSupportOk

`func (o *ServiceOffering) GetProductTierSupportOk() (*string, bool)`

GetProductTierSupportOk returns a tuple with the ProductTierSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierSupport

`func (o *ServiceOffering) SetProductTierSupport(v string)`

SetProductTierSupport sets ProductTierSupport field to given value.


### GetProductTierType

`func (o *ServiceOffering) GetProductTierType() string`

GetProductTierType returns the ProductTierType field if non-nil, zero value otherwise.

### GetProductTierTypeOk

`func (o *ServiceOffering) GetProductTierTypeOk() (*string, bool)`

GetProductTierTypeOk returns a tuple with the ProductTierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierType

`func (o *ServiceOffering) SetProductTierType(v string)`

SetProductTierType sets ProductTierType field to given value.


### GetProductTierURLKey

`func (o *ServiceOffering) GetProductTierURLKey() string`

GetProductTierURLKey returns the ProductTierURLKey field if non-nil, zero value otherwise.

### GetProductTierURLKeyOk

`func (o *ServiceOffering) GetProductTierURLKeyOk() (*string, bool)`

GetProductTierURLKeyOk returns a tuple with the ProductTierURLKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierURLKey

`func (o *ServiceOffering) SetProductTierURLKey(v string)`

SetProductTierURLKey sets ProductTierURLKey field to given value.


### GetProductTierVersion

`func (o *ServiceOffering) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *ServiceOffering) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *ServiceOffering) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.


### GetResourceParameters

`func (o *ServiceOffering) GetResourceParameters() []ResourceEntity`

GetResourceParameters returns the ResourceParameters field if non-nil, zero value otherwise.

### GetResourceParametersOk

`func (o *ServiceOffering) GetResourceParametersOk() (*[]ResourceEntity, bool)`

GetResourceParametersOk returns a tuple with the ResourceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceParameters

`func (o *ServiceOffering) SetResourceParameters(v []ResourceEntity)`

SetResourceParameters sets ResourceParameters field to given value.


### GetServiceAPIID

`func (o *ServiceOffering) GetServiceAPIID() string`

GetServiceAPIID returns the ServiceAPIID field if non-nil, zero value otherwise.

### GetServiceAPIIDOk

`func (o *ServiceOffering) GetServiceAPIIDOk() (*string, bool)`

GetServiceAPIIDOk returns a tuple with the ServiceAPIID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPIID

`func (o *ServiceOffering) SetServiceAPIID(v string)`

SetServiceAPIID sets ServiceAPIID field to given value.


### GetServiceAPIVersion

`func (o *ServiceOffering) GetServiceAPIVersion() string`

GetServiceAPIVersion returns the ServiceAPIVersion field if non-nil, zero value otherwise.

### GetServiceAPIVersionOk

`func (o *ServiceOffering) GetServiceAPIVersionOk() (*string, bool)`

GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPIVersion

`func (o *ServiceOffering) SetServiceAPIVersion(v string)`

SetServiceAPIVersion sets ServiceAPIVersion field to given value.


### GetServiceEnvironmentID

`func (o *ServiceOffering) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *ServiceOffering) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *ServiceOffering) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.


### GetServiceEnvironmentName

`func (o *ServiceOffering) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *ServiceOffering) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *ServiceOffering) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceEnvironmentType

`func (o *ServiceOffering) GetServiceEnvironmentType() string`

GetServiceEnvironmentType returns the ServiceEnvironmentType field if non-nil, zero value otherwise.

### GetServiceEnvironmentTypeOk

`func (o *ServiceOffering) GetServiceEnvironmentTypeOk() (*string, bool)`

GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentType

`func (o *ServiceOffering) SetServiceEnvironmentType(v string)`

SetServiceEnvironmentType sets ServiceEnvironmentType field to given value.


### GetServiceEnvironmentURLKey

`func (o *ServiceOffering) GetServiceEnvironmentURLKey() string`

GetServiceEnvironmentURLKey returns the ServiceEnvironmentURLKey field if non-nil, zero value otherwise.

### GetServiceEnvironmentURLKeyOk

`func (o *ServiceOffering) GetServiceEnvironmentURLKeyOk() (*string, bool)`

GetServiceEnvironmentURLKeyOk returns a tuple with the ServiceEnvironmentURLKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentURLKey

`func (o *ServiceOffering) SetServiceEnvironmentURLKey(v string)`

SetServiceEnvironmentURLKey sets ServiceEnvironmentURLKey field to given value.


### GetServiceEnvironmentVisibility

`func (o *ServiceOffering) GetServiceEnvironmentVisibility() string`

GetServiceEnvironmentVisibility returns the ServiceEnvironmentVisibility field if non-nil, zero value otherwise.

### GetServiceEnvironmentVisibilityOk

`func (o *ServiceOffering) GetServiceEnvironmentVisibilityOk() (*string, bool)`

GetServiceEnvironmentVisibilityOk returns a tuple with the ServiceEnvironmentVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentVisibility

`func (o *ServiceOffering) SetServiceEnvironmentVisibility(v string)`

SetServiceEnvironmentVisibility sets ServiceEnvironmentVisibility field to given value.


### GetServiceLogoURL

`func (o *ServiceOffering) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *ServiceOffering) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *ServiceOffering) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.


### GetServiceModelFeatures

`func (o *ServiceOffering) GetServiceModelFeatures() []ServiceModelFeatureDetail`

GetServiceModelFeatures returns the ServiceModelFeatures field if non-nil, zero value otherwise.

### GetServiceModelFeaturesOk

`func (o *ServiceOffering) GetServiceModelFeaturesOk() (*[]ServiceModelFeatureDetail, bool)`

GetServiceModelFeaturesOk returns a tuple with the ServiceModelFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelFeatures

`func (o *ServiceOffering) SetServiceModelFeatures(v []ServiceModelFeatureDetail)`

SetServiceModelFeatures sets ServiceModelFeatures field to given value.

### HasServiceModelFeatures

`func (o *ServiceOffering) HasServiceModelFeatures() bool`

HasServiceModelFeatures returns a boolean if a field has been set.

### GetServiceModelID

`func (o *ServiceOffering) GetServiceModelID() string`

GetServiceModelID returns the ServiceModelID field if non-nil, zero value otherwise.

### GetServiceModelIDOk

`func (o *ServiceOffering) GetServiceModelIDOk() (*string, bool)`

GetServiceModelIDOk returns a tuple with the ServiceModelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelID

`func (o *ServiceOffering) SetServiceModelID(v string)`

SetServiceModelID sets ServiceModelID field to given value.


### GetServiceModelName

`func (o *ServiceOffering) GetServiceModelName() string`

GetServiceModelName returns the ServiceModelName field if non-nil, zero value otherwise.

### GetServiceModelNameOk

`func (o *ServiceOffering) GetServiceModelNameOk() (*string, bool)`

GetServiceModelNameOk returns a tuple with the ServiceModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelName

`func (o *ServiceOffering) SetServiceModelName(v string)`

SetServiceModelName sets ServiceModelName field to given value.


### GetServiceModelStatus

`func (o *ServiceOffering) GetServiceModelStatus() string`

GetServiceModelStatus returns the ServiceModelStatus field if non-nil, zero value otherwise.

### GetServiceModelStatusOk

`func (o *ServiceOffering) GetServiceModelStatusOk() (*string, bool)`

GetServiceModelStatusOk returns a tuple with the ServiceModelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelStatus

`func (o *ServiceOffering) SetServiceModelStatus(v string)`

SetServiceModelStatus sets ServiceModelStatus field to given value.


### GetServiceModelType

`func (o *ServiceOffering) GetServiceModelType() string`

GetServiceModelType returns the ServiceModelType field if non-nil, zero value otherwise.

### GetServiceModelTypeOk

`func (o *ServiceOffering) GetServiceModelTypeOk() (*string, bool)`

GetServiceModelTypeOk returns a tuple with the ServiceModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelType

`func (o *ServiceOffering) SetServiceModelType(v string)`

SetServiceModelType sets ServiceModelType field to given value.


### GetServiceModelURLKey

`func (o *ServiceOffering) GetServiceModelURLKey() string`

GetServiceModelURLKey returns the ServiceModelURLKey field if non-nil, zero value otherwise.

### GetServiceModelURLKeyOk

`func (o *ServiceOffering) GetServiceModelURLKeyOk() (*string, bool)`

GetServiceModelURLKeyOk returns a tuple with the ServiceModelURLKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelURLKey

`func (o *ServiceOffering) SetServiceModelURLKey(v string)`

SetServiceModelURLKey sets ServiceModelURLKey field to given value.


### GetSupportsPublicNetwork

`func (o *ServiceOffering) GetSupportsPublicNetwork() bool`

GetSupportsPublicNetwork returns the SupportsPublicNetwork field if non-nil, zero value otherwise.

### GetSupportsPublicNetworkOk

`func (o *ServiceOffering) GetSupportsPublicNetworkOk() (*bool, bool)`

GetSupportsPublicNetworkOk returns a tuple with the SupportsPublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPublicNetwork

`func (o *ServiceOffering) SetSupportsPublicNetwork(v bool)`

SetSupportsPublicNetwork sets SupportsPublicNetwork field to given value.

### HasSupportsPublicNetwork

`func (o *ServiceOffering) HasSupportsPublicNetwork() bool`

HasSupportsPublicNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


