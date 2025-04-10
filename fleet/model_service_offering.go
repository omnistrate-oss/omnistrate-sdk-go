/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"fmt"
)

// checks if the ServiceOffering type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceOffering{}

// ServiceOffering struct for ServiceOffering
type ServiceOffering struct {
	// Auto approve subscription or not
	AutoApproveSubscription bool `json:"AutoApproveSubscription"`
	// Allow creates when payment not configured
	AllowCreatesWhenPaymentNotConfigured *bool `json:"allowCreatesWhenPaymentNotConfigured,omitempty"`
	Assets *ServiceAssets `json:"assets,omitempty"`
	// The AWS regions that this service offering is available on
	AwsRegions []string `json:"awsRegions,omitempty"`
	// The Azure regions that this service offering is available on
	AzureRegions []string `json:"azureRegions,omitempty"`
	// List of supported cloud providers for this product tier.
	CloudProviders []string `json:"cloudProviders,omitempty"`
	// The GCP regions that this service offering is available on
	GcpRegions []string `json:"gcpRegions,omitempty"`
	// Maximum number of instances
	MaxNumberOfInstances *int64 `json:"maxNumberOfInstances,omitempty"`
	// A brief description of the product tier
	ProductTierDescription *string `json:"productTierDescription,omitempty"`
	// Documentation
	ProductTierDocumentation string `json:"productTierDocumentation"`
	// ID of a Product Tier
	ProductTierID string `json:"productTierID"`
	// The product tier name
	ProductTierName string `json:"productTierName"`
	// A brief description for the end user of the product tier
	ProductTierPlanDescription *string `json:"productTierPlanDescription,omitempty"`
	// Pricing
	ProductTierPricing interface{} `json:"productTierPricing"`
	// Support
	ProductTierSupport string `json:"productTierSupport"`
	// ProductTierType is the type of tier for a product
	ProductTierType string `json:"productTierType"`
	// The product tier URL key
	ProductTierURLKey string `json:"productTierURLKey"`
	// The product tier version
	ProductTierVersion string `json:"productTierVersion"`
	// The resource parameters
	ResourceParameters []ResourceEntity `json:"resourceParameters"`
	// ID of a Service API
	ServiceAPIID string `json:"serviceAPIID"`
	// The service API version
	ServiceAPIVersion string `json:"serviceAPIVersion"`
	// ID of a Service Environment
	ServiceEnvironmentID string `json:"serviceEnvironmentID"`
	// The service environment name
	ServiceEnvironmentName string `json:"serviceEnvironmentName"`
	// The type of service environment
	ServiceEnvironmentType string `json:"serviceEnvironmentType"`
	// The service environment URL key
	ServiceEnvironmentURLKey string `json:"serviceEnvironmentURLKey"`
	// This parameter is used to configure the visibility of the service control-plane APIs
	ServiceEnvironmentVisibility string `json:"serviceEnvironmentVisibility"`
	// The logo for the service
	ServiceLogoURL string `json:"serviceLogoURL"`
	// Enabled service model features
	ServiceModelFeatures []ServiceModelFeatureDetail `json:"serviceModelFeatures,omitempty"`
	// ID of a Service Model
	ServiceModelID string `json:"serviceModelID"`
	// The service model name
	ServiceModelName string `json:"serviceModelName"`
	// The service model status
	ServiceModelStatus string `json:"serviceModelStatus"`
	// The model type encapsulating this service
	ServiceModelType string `json:"serviceModelType"`
	// The service model URL key
	ServiceModelURLKey string `json:"serviceModelURLKey"`
	// Indicates whether any of the resources in the product tier support public network
	SupportsPublicNetwork *bool `json:"supportsPublicNetwork,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceOffering ServiceOffering

// NewServiceOffering instantiates a new ServiceOffering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceOffering(autoApproveSubscription bool, productTierDocumentation string, productTierID string, productTierName string, productTierPricing interface{}, productTierSupport string, productTierType string, productTierURLKey string, productTierVersion string, resourceParameters []ResourceEntity, serviceAPIID string, serviceAPIVersion string, serviceEnvironmentID string, serviceEnvironmentName string, serviceEnvironmentType string, serviceEnvironmentURLKey string, serviceEnvironmentVisibility string, serviceLogoURL string, serviceModelID string, serviceModelName string, serviceModelStatus string, serviceModelType string, serviceModelURLKey string) *ServiceOffering {
	this := ServiceOffering{}
	this.AutoApproveSubscription = autoApproveSubscription
	this.ProductTierDocumentation = productTierDocumentation
	this.ProductTierID = productTierID
	this.ProductTierName = productTierName
	this.ProductTierPricing = productTierPricing
	this.ProductTierSupport = productTierSupport
	this.ProductTierType = productTierType
	this.ProductTierURLKey = productTierURLKey
	this.ProductTierVersion = productTierVersion
	this.ResourceParameters = resourceParameters
	this.ServiceAPIID = serviceAPIID
	this.ServiceAPIVersion = serviceAPIVersion
	this.ServiceEnvironmentID = serviceEnvironmentID
	this.ServiceEnvironmentName = serviceEnvironmentName
	this.ServiceEnvironmentType = serviceEnvironmentType
	this.ServiceEnvironmentURLKey = serviceEnvironmentURLKey
	this.ServiceEnvironmentVisibility = serviceEnvironmentVisibility
	this.ServiceLogoURL = serviceLogoURL
	this.ServiceModelID = serviceModelID
	this.ServiceModelName = serviceModelName
	this.ServiceModelStatus = serviceModelStatus
	this.ServiceModelType = serviceModelType
	this.ServiceModelURLKey = serviceModelURLKey
	return &this
}

// NewServiceOfferingWithDefaults instantiates a new ServiceOffering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceOfferingWithDefaults() *ServiceOffering {
	this := ServiceOffering{}
	return &this
}

// GetAutoApproveSubscription returns the AutoApproveSubscription field value
func (o *ServiceOffering) GetAutoApproveSubscription() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoApproveSubscription
}

// GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetAutoApproveSubscriptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoApproveSubscription, true
}

// SetAutoApproveSubscription sets field value
func (o *ServiceOffering) SetAutoApproveSubscription(v bool) {
	o.AutoApproveSubscription = v
}

// GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field value if set, zero value otherwise.
func (o *ServiceOffering) GetAllowCreatesWhenPaymentNotConfigured() bool {
	if o == nil || IsNil(o.AllowCreatesWhenPaymentNotConfigured) {
		var ret bool
		return ret
	}
	return *o.AllowCreatesWhenPaymentNotConfigured
}

// GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCreatesWhenPaymentNotConfigured) {
		return nil, false
	}
	return o.AllowCreatesWhenPaymentNotConfigured, true
}

// HasAllowCreatesWhenPaymentNotConfigured returns a boolean if a field has been set.
func (o *ServiceOffering) HasAllowCreatesWhenPaymentNotConfigured() bool {
	if o != nil && !IsNil(o.AllowCreatesWhenPaymentNotConfigured) {
		return true
	}

	return false
}

// SetAllowCreatesWhenPaymentNotConfigured gets a reference to the given bool and assigns it to the AllowCreatesWhenPaymentNotConfigured field.
func (o *ServiceOffering) SetAllowCreatesWhenPaymentNotConfigured(v bool) {
	o.AllowCreatesWhenPaymentNotConfigured = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *ServiceOffering) GetAssets() ServiceAssets {
	if o == nil || IsNil(o.Assets) {
		var ret ServiceAssets
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetAssetsOk() (*ServiceAssets, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *ServiceOffering) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given ServiceAssets and assigns it to the Assets field.
func (o *ServiceOffering) SetAssets(v ServiceAssets) {
	o.Assets = &v
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *ServiceOffering) GetAwsRegions() []string {
	if o == nil || IsNil(o.AwsRegions) {
		var ret []string
		return ret
	}
	return o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetAwsRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AwsRegions) {
		return nil, false
	}
	return o.AwsRegions, true
}

// HasAwsRegions returns a boolean if a field has been set.
func (o *ServiceOffering) HasAwsRegions() bool {
	if o != nil && !IsNil(o.AwsRegions) {
		return true
	}

	return false
}

// SetAwsRegions gets a reference to the given []string and assigns it to the AwsRegions field.
func (o *ServiceOffering) SetAwsRegions(v []string) {
	o.AwsRegions = v
}

// GetAzureRegions returns the AzureRegions field value if set, zero value otherwise.
func (o *ServiceOffering) GetAzureRegions() []string {
	if o == nil || IsNil(o.AzureRegions) {
		var ret []string
		return ret
	}
	return o.AzureRegions
}

// GetAzureRegionsOk returns a tuple with the AzureRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetAzureRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AzureRegions) {
		return nil, false
	}
	return o.AzureRegions, true
}

// HasAzureRegions returns a boolean if a field has been set.
func (o *ServiceOffering) HasAzureRegions() bool {
	if o != nil && !IsNil(o.AzureRegions) {
		return true
	}

	return false
}

// SetAzureRegions gets a reference to the given []string and assigns it to the AzureRegions field.
func (o *ServiceOffering) SetAzureRegions(v []string) {
	o.AzureRegions = v
}

// GetCloudProviders returns the CloudProviders field value if set, zero value otherwise.
func (o *ServiceOffering) GetCloudProviders() []string {
	if o == nil || IsNil(o.CloudProviders) {
		var ret []string
		return ret
	}
	return o.CloudProviders
}

// GetCloudProvidersOk returns a tuple with the CloudProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetCloudProvidersOk() ([]string, bool) {
	if o == nil || IsNil(o.CloudProviders) {
		return nil, false
	}
	return o.CloudProviders, true
}

// HasCloudProviders returns a boolean if a field has been set.
func (o *ServiceOffering) HasCloudProviders() bool {
	if o != nil && !IsNil(o.CloudProviders) {
		return true
	}

	return false
}

// SetCloudProviders gets a reference to the given []string and assigns it to the CloudProviders field.
func (o *ServiceOffering) SetCloudProviders(v []string) {
	o.CloudProviders = v
}

// GetGcpRegions returns the GcpRegions field value if set, zero value otherwise.
func (o *ServiceOffering) GetGcpRegions() []string {
	if o == nil || IsNil(o.GcpRegions) {
		var ret []string
		return ret
	}
	return o.GcpRegions
}

// GetGcpRegionsOk returns a tuple with the GcpRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetGcpRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.GcpRegions) {
		return nil, false
	}
	return o.GcpRegions, true
}

// HasGcpRegions returns a boolean if a field has been set.
func (o *ServiceOffering) HasGcpRegions() bool {
	if o != nil && !IsNil(o.GcpRegions) {
		return true
	}

	return false
}

// SetGcpRegions gets a reference to the given []string and assigns it to the GcpRegions field.
func (o *ServiceOffering) SetGcpRegions(v []string) {
	o.GcpRegions = v
}

// GetMaxNumberOfInstances returns the MaxNumberOfInstances field value if set, zero value otherwise.
func (o *ServiceOffering) GetMaxNumberOfInstances() int64 {
	if o == nil || IsNil(o.MaxNumberOfInstances) {
		var ret int64
		return ret
	}
	return *o.MaxNumberOfInstances
}

// GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetMaxNumberOfInstancesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxNumberOfInstances) {
		return nil, false
	}
	return o.MaxNumberOfInstances, true
}

// HasMaxNumberOfInstances returns a boolean if a field has been set.
func (o *ServiceOffering) HasMaxNumberOfInstances() bool {
	if o != nil && !IsNil(o.MaxNumberOfInstances) {
		return true
	}

	return false
}

// SetMaxNumberOfInstances gets a reference to the given int64 and assigns it to the MaxNumberOfInstances field.
func (o *ServiceOffering) SetMaxNumberOfInstances(v int64) {
	o.MaxNumberOfInstances = &v
}

// GetProductTierDescription returns the ProductTierDescription field value if set, zero value otherwise.
func (o *ServiceOffering) GetProductTierDescription() string {
	if o == nil || IsNil(o.ProductTierDescription) {
		var ret string
		return ret
	}
	return *o.ProductTierDescription
}

// GetProductTierDescriptionOk returns a tuple with the ProductTierDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetProductTierDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierDescription) {
		return nil, false
	}
	return o.ProductTierDescription, true
}

// HasProductTierDescription returns a boolean if a field has been set.
func (o *ServiceOffering) HasProductTierDescription() bool {
	if o != nil && !IsNil(o.ProductTierDescription) {
		return true
	}

	return false
}

// SetProductTierDescription gets a reference to the given string and assigns it to the ProductTierDescription field.
func (o *ServiceOffering) SetProductTierDescription(v string) {
	o.ProductTierDescription = &v
}

// GetProductTierDocumentation returns the ProductTierDocumentation field value
func (o *ServiceOffering) GetProductTierDocumentation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierDocumentation
}

// GetProductTierDocumentationOk returns a tuple with the ProductTierDocumentation field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetProductTierDocumentationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierDocumentation, true
}

// SetProductTierDocumentation sets field value
func (o *ServiceOffering) SetProductTierDocumentation(v string) {
	o.ProductTierDocumentation = v
}

// GetProductTierID returns the ProductTierID field value
func (o *ServiceOffering) GetProductTierID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierID
}

// GetProductTierIDOk returns a tuple with the ProductTierID field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetProductTierIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierID, true
}

// SetProductTierID sets field value
func (o *ServiceOffering) SetProductTierID(v string) {
	o.ProductTierID = v
}

// GetProductTierName returns the ProductTierName field value
func (o *ServiceOffering) GetProductTierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierName
}

// GetProductTierNameOk returns a tuple with the ProductTierName field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetProductTierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierName, true
}

// SetProductTierName sets field value
func (o *ServiceOffering) SetProductTierName(v string) {
	o.ProductTierName = v
}

// GetProductTierPlanDescription returns the ProductTierPlanDescription field value if set, zero value otherwise.
func (o *ServiceOffering) GetProductTierPlanDescription() string {
	if o == nil || IsNil(o.ProductTierPlanDescription) {
		var ret string
		return ret
	}
	return *o.ProductTierPlanDescription
}

// GetProductTierPlanDescriptionOk returns a tuple with the ProductTierPlanDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetProductTierPlanDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierPlanDescription) {
		return nil, false
	}
	return o.ProductTierPlanDescription, true
}

// HasProductTierPlanDescription returns a boolean if a field has been set.
func (o *ServiceOffering) HasProductTierPlanDescription() bool {
	if o != nil && !IsNil(o.ProductTierPlanDescription) {
		return true
	}

	return false
}

// SetProductTierPlanDescription gets a reference to the given string and assigns it to the ProductTierPlanDescription field.
func (o *ServiceOffering) SetProductTierPlanDescription(v string) {
	o.ProductTierPlanDescription = &v
}

// GetProductTierPricing returns the ProductTierPricing field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ServiceOffering) GetProductTierPricing() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ProductTierPricing
}

// GetProductTierPricingOk returns a tuple with the ProductTierPricing field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceOffering) GetProductTierPricingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ProductTierPricing) {
		return nil, false
	}
	return &o.ProductTierPricing, true
}

// SetProductTierPricing sets field value
func (o *ServiceOffering) SetProductTierPricing(v interface{}) {
	o.ProductTierPricing = v
}

// GetProductTierSupport returns the ProductTierSupport field value
func (o *ServiceOffering) GetProductTierSupport() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierSupport
}

// GetProductTierSupportOk returns a tuple with the ProductTierSupport field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetProductTierSupportOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierSupport, true
}

// SetProductTierSupport sets field value
func (o *ServiceOffering) SetProductTierSupport(v string) {
	o.ProductTierSupport = v
}

// GetProductTierType returns the ProductTierType field value
func (o *ServiceOffering) GetProductTierType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierType
}

// GetProductTierTypeOk returns a tuple with the ProductTierType field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetProductTierTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierType, true
}

// SetProductTierType sets field value
func (o *ServiceOffering) SetProductTierType(v string) {
	o.ProductTierType = v
}

// GetProductTierURLKey returns the ProductTierURLKey field value
func (o *ServiceOffering) GetProductTierURLKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierURLKey
}

// GetProductTierURLKeyOk returns a tuple with the ProductTierURLKey field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetProductTierURLKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierURLKey, true
}

// SetProductTierURLKey sets field value
func (o *ServiceOffering) SetProductTierURLKey(v string) {
	o.ProductTierURLKey = v
}

// GetProductTierVersion returns the ProductTierVersion field value
func (o *ServiceOffering) GetProductTierVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierVersion
}

// GetProductTierVersionOk returns a tuple with the ProductTierVersion field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetProductTierVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierVersion, true
}

// SetProductTierVersion sets field value
func (o *ServiceOffering) SetProductTierVersion(v string) {
	o.ProductTierVersion = v
}

// GetResourceParameters returns the ResourceParameters field value
func (o *ServiceOffering) GetResourceParameters() []ResourceEntity {
	if o == nil {
		var ret []ResourceEntity
		return ret
	}

	return o.ResourceParameters
}

// GetResourceParametersOk returns a tuple with the ResourceParameters field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetResourceParametersOk() ([]ResourceEntity, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceParameters, true
}

// SetResourceParameters sets field value
func (o *ServiceOffering) SetResourceParameters(v []ResourceEntity) {
	o.ResourceParameters = v
}

// GetServiceAPIID returns the ServiceAPIID field value
func (o *ServiceOffering) GetServiceAPIID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAPIID
}

// GetServiceAPIIDOk returns a tuple with the ServiceAPIID field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceAPIIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAPIID, true
}

// SetServiceAPIID sets field value
func (o *ServiceOffering) SetServiceAPIID(v string) {
	o.ServiceAPIID = v
}

// GetServiceAPIVersion returns the ServiceAPIVersion field value
func (o *ServiceOffering) GetServiceAPIVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAPIVersion
}

// GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceAPIVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAPIVersion, true
}

// SetServiceAPIVersion sets field value
func (o *ServiceOffering) SetServiceAPIVersion(v string) {
	o.ServiceAPIVersion = v
}

// GetServiceEnvironmentID returns the ServiceEnvironmentID field value
func (o *ServiceOffering) GetServiceEnvironmentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentID
}

// GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceEnvironmentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentID, true
}

// SetServiceEnvironmentID sets field value
func (o *ServiceOffering) SetServiceEnvironmentID(v string) {
	o.ServiceEnvironmentID = v
}

// GetServiceEnvironmentName returns the ServiceEnvironmentName field value
func (o *ServiceOffering) GetServiceEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentName
}

// GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentName, true
}

// SetServiceEnvironmentName sets field value
func (o *ServiceOffering) SetServiceEnvironmentName(v string) {
	o.ServiceEnvironmentName = v
}

// GetServiceEnvironmentType returns the ServiceEnvironmentType field value
func (o *ServiceOffering) GetServiceEnvironmentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentType
}

// GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceEnvironmentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentType, true
}

// SetServiceEnvironmentType sets field value
func (o *ServiceOffering) SetServiceEnvironmentType(v string) {
	o.ServiceEnvironmentType = v
}

// GetServiceEnvironmentURLKey returns the ServiceEnvironmentURLKey field value
func (o *ServiceOffering) GetServiceEnvironmentURLKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentURLKey
}

// GetServiceEnvironmentURLKeyOk returns a tuple with the ServiceEnvironmentURLKey field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceEnvironmentURLKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentURLKey, true
}

// SetServiceEnvironmentURLKey sets field value
func (o *ServiceOffering) SetServiceEnvironmentURLKey(v string) {
	o.ServiceEnvironmentURLKey = v
}

// GetServiceEnvironmentVisibility returns the ServiceEnvironmentVisibility field value
func (o *ServiceOffering) GetServiceEnvironmentVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentVisibility
}

// GetServiceEnvironmentVisibilityOk returns a tuple with the ServiceEnvironmentVisibility field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceEnvironmentVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentVisibility, true
}

// SetServiceEnvironmentVisibility sets field value
func (o *ServiceOffering) SetServiceEnvironmentVisibility(v string) {
	o.ServiceEnvironmentVisibility = v
}

// GetServiceLogoURL returns the ServiceLogoURL field value
func (o *ServiceOffering) GetServiceLogoURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceLogoURL
}

// GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceLogoURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceLogoURL, true
}

// SetServiceLogoURL sets field value
func (o *ServiceOffering) SetServiceLogoURL(v string) {
	o.ServiceLogoURL = v
}

// GetServiceModelFeatures returns the ServiceModelFeatures field value if set, zero value otherwise.
func (o *ServiceOffering) GetServiceModelFeatures() []ServiceModelFeatureDetail {
	if o == nil || IsNil(o.ServiceModelFeatures) {
		var ret []ServiceModelFeatureDetail
		return ret
	}
	return o.ServiceModelFeatures
}

// GetServiceModelFeaturesOk returns a tuple with the ServiceModelFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceModelFeaturesOk() ([]ServiceModelFeatureDetail, bool) {
	if o == nil || IsNil(o.ServiceModelFeatures) {
		return nil, false
	}
	return o.ServiceModelFeatures, true
}

// HasServiceModelFeatures returns a boolean if a field has been set.
func (o *ServiceOffering) HasServiceModelFeatures() bool {
	if o != nil && !IsNil(o.ServiceModelFeatures) {
		return true
	}

	return false
}

// SetServiceModelFeatures gets a reference to the given []ServiceModelFeatureDetail and assigns it to the ServiceModelFeatures field.
func (o *ServiceOffering) SetServiceModelFeatures(v []ServiceModelFeatureDetail) {
	o.ServiceModelFeatures = v
}

// GetServiceModelID returns the ServiceModelID field value
func (o *ServiceOffering) GetServiceModelID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelID
}

// GetServiceModelIDOk returns a tuple with the ServiceModelID field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceModelIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelID, true
}

// SetServiceModelID sets field value
func (o *ServiceOffering) SetServiceModelID(v string) {
	o.ServiceModelID = v
}

// GetServiceModelName returns the ServiceModelName field value
func (o *ServiceOffering) GetServiceModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelName
}

// GetServiceModelNameOk returns a tuple with the ServiceModelName field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelName, true
}

// SetServiceModelName sets field value
func (o *ServiceOffering) SetServiceModelName(v string) {
	o.ServiceModelName = v
}

// GetServiceModelStatus returns the ServiceModelStatus field value
func (o *ServiceOffering) GetServiceModelStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelStatus
}

// GetServiceModelStatusOk returns a tuple with the ServiceModelStatus field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceModelStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelStatus, true
}

// SetServiceModelStatus sets field value
func (o *ServiceOffering) SetServiceModelStatus(v string) {
	o.ServiceModelStatus = v
}

// GetServiceModelType returns the ServiceModelType field value
func (o *ServiceOffering) GetServiceModelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelType
}

// GetServiceModelTypeOk returns a tuple with the ServiceModelType field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceModelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelType, true
}

// SetServiceModelType sets field value
func (o *ServiceOffering) SetServiceModelType(v string) {
	o.ServiceModelType = v
}

// GetServiceModelURLKey returns the ServiceModelURLKey field value
func (o *ServiceOffering) GetServiceModelURLKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelURLKey
}

// GetServiceModelURLKeyOk returns a tuple with the ServiceModelURLKey field value
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetServiceModelURLKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelURLKey, true
}

// SetServiceModelURLKey sets field value
func (o *ServiceOffering) SetServiceModelURLKey(v string) {
	o.ServiceModelURLKey = v
}

// GetSupportsPublicNetwork returns the SupportsPublicNetwork field value if set, zero value otherwise.
func (o *ServiceOffering) GetSupportsPublicNetwork() bool {
	if o == nil || IsNil(o.SupportsPublicNetwork) {
		var ret bool
		return ret
	}
	return *o.SupportsPublicNetwork
}

// GetSupportsPublicNetworkOk returns a tuple with the SupportsPublicNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOffering) GetSupportsPublicNetworkOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsPublicNetwork) {
		return nil, false
	}
	return o.SupportsPublicNetwork, true
}

// HasSupportsPublicNetwork returns a boolean if a field has been set.
func (o *ServiceOffering) HasSupportsPublicNetwork() bool {
	if o != nil && !IsNil(o.SupportsPublicNetwork) {
		return true
	}

	return false
}

// SetSupportsPublicNetwork gets a reference to the given bool and assigns it to the SupportsPublicNetwork field.
func (o *ServiceOffering) SetSupportsPublicNetwork(v bool) {
	o.SupportsPublicNetwork = &v
}

func (o ServiceOffering) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceOffering) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AutoApproveSubscription"] = o.AutoApproveSubscription
	if !IsNil(o.AllowCreatesWhenPaymentNotConfigured) {
		toSerialize["allowCreatesWhenPaymentNotConfigured"] = o.AllowCreatesWhenPaymentNotConfigured
	}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !IsNil(o.AwsRegions) {
		toSerialize["awsRegions"] = o.AwsRegions
	}
	if !IsNil(o.AzureRegions) {
		toSerialize["azureRegions"] = o.AzureRegions
	}
	if !IsNil(o.CloudProviders) {
		toSerialize["cloudProviders"] = o.CloudProviders
	}
	if !IsNil(o.GcpRegions) {
		toSerialize["gcpRegions"] = o.GcpRegions
	}
	if !IsNil(o.MaxNumberOfInstances) {
		toSerialize["maxNumberOfInstances"] = o.MaxNumberOfInstances
	}
	if !IsNil(o.ProductTierDescription) {
		toSerialize["productTierDescription"] = o.ProductTierDescription
	}
	toSerialize["productTierDocumentation"] = o.ProductTierDocumentation
	toSerialize["productTierID"] = o.ProductTierID
	toSerialize["productTierName"] = o.ProductTierName
	if !IsNil(o.ProductTierPlanDescription) {
		toSerialize["productTierPlanDescription"] = o.ProductTierPlanDescription
	}
	if o.ProductTierPricing != nil {
		toSerialize["productTierPricing"] = o.ProductTierPricing
	}
	toSerialize["productTierSupport"] = o.ProductTierSupport
	toSerialize["productTierType"] = o.ProductTierType
	toSerialize["productTierURLKey"] = o.ProductTierURLKey
	toSerialize["productTierVersion"] = o.ProductTierVersion
	toSerialize["resourceParameters"] = o.ResourceParameters
	toSerialize["serviceAPIID"] = o.ServiceAPIID
	toSerialize["serviceAPIVersion"] = o.ServiceAPIVersion
	toSerialize["serviceEnvironmentID"] = o.ServiceEnvironmentID
	toSerialize["serviceEnvironmentName"] = o.ServiceEnvironmentName
	toSerialize["serviceEnvironmentType"] = o.ServiceEnvironmentType
	toSerialize["serviceEnvironmentURLKey"] = o.ServiceEnvironmentURLKey
	toSerialize["serviceEnvironmentVisibility"] = o.ServiceEnvironmentVisibility
	toSerialize["serviceLogoURL"] = o.ServiceLogoURL
	if !IsNil(o.ServiceModelFeatures) {
		toSerialize["serviceModelFeatures"] = o.ServiceModelFeatures
	}
	toSerialize["serviceModelID"] = o.ServiceModelID
	toSerialize["serviceModelName"] = o.ServiceModelName
	toSerialize["serviceModelStatus"] = o.ServiceModelStatus
	toSerialize["serviceModelType"] = o.ServiceModelType
	toSerialize["serviceModelURLKey"] = o.ServiceModelURLKey
	if !IsNil(o.SupportsPublicNetwork) {
		toSerialize["supportsPublicNetwork"] = o.SupportsPublicNetwork
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceOffering) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AutoApproveSubscription",
		"productTierDocumentation",
		"productTierID",
		"productTierName",
		"productTierPricing",
		"productTierSupport",
		"productTierType",
		"productTierURLKey",
		"productTierVersion",
		"resourceParameters",
		"serviceAPIID",
		"serviceAPIVersion",
		"serviceEnvironmentID",
		"serviceEnvironmentName",
		"serviceEnvironmentType",
		"serviceEnvironmentURLKey",
		"serviceEnvironmentVisibility",
		"serviceLogoURL",
		"serviceModelID",
		"serviceModelName",
		"serviceModelStatus",
		"serviceModelType",
		"serviceModelURLKey",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServiceOffering := _ServiceOffering{}

	err = json.Unmarshal(data, &varServiceOffering)

	if err != nil {
		return err
	}

	*o = ServiceOffering(varServiceOffering)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "AutoApproveSubscription")
		delete(additionalProperties, "allowCreatesWhenPaymentNotConfigured")
		delete(additionalProperties, "assets")
		delete(additionalProperties, "awsRegions")
		delete(additionalProperties, "azureRegions")
		delete(additionalProperties, "cloudProviders")
		delete(additionalProperties, "gcpRegions")
		delete(additionalProperties, "maxNumberOfInstances")
		delete(additionalProperties, "productTierDescription")
		delete(additionalProperties, "productTierDocumentation")
		delete(additionalProperties, "productTierID")
		delete(additionalProperties, "productTierName")
		delete(additionalProperties, "productTierPlanDescription")
		delete(additionalProperties, "productTierPricing")
		delete(additionalProperties, "productTierSupport")
		delete(additionalProperties, "productTierType")
		delete(additionalProperties, "productTierURLKey")
		delete(additionalProperties, "productTierVersion")
		delete(additionalProperties, "resourceParameters")
		delete(additionalProperties, "serviceAPIID")
		delete(additionalProperties, "serviceAPIVersion")
		delete(additionalProperties, "serviceEnvironmentID")
		delete(additionalProperties, "serviceEnvironmentName")
		delete(additionalProperties, "serviceEnvironmentType")
		delete(additionalProperties, "serviceEnvironmentURLKey")
		delete(additionalProperties, "serviceEnvironmentVisibility")
		delete(additionalProperties, "serviceLogoURL")
		delete(additionalProperties, "serviceModelFeatures")
		delete(additionalProperties, "serviceModelID")
		delete(additionalProperties, "serviceModelName")
		delete(additionalProperties, "serviceModelStatus")
		delete(additionalProperties, "serviceModelType")
		delete(additionalProperties, "serviceModelURLKey")
		delete(additionalProperties, "supportsPublicNetwork")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceOffering struct {
	value *ServiceOffering
	isSet bool
}

func (v NullableServiceOffering) Get() *ServiceOffering {
	return v.value
}

func (v *NullableServiceOffering) Set(val *ServiceOffering) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceOffering) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceOffering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceOffering(val *ServiceOffering) *NullableServiceOffering {
	return &NullableServiceOffering{value: val, isSet: true}
}

func (v NullableServiceOffering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceOffering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


