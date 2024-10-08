/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetServicePlanResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServicePlanResult{}

// GetServicePlanResult struct for GetServicePlanResult
type GetServicePlanResult struct {
	// Auto approve subscription or not
	AutoApproveSubscription bool `json:"AutoApproveSubscription"`
	// The infrastructure account configuration ID list
	AccountConfigIds []string `json:"accountConfigIds,omitempty"`
	// The active infrastructure account configuration IDs per cloud provider
	ActiveAccountConfigIds *map[string]string `json:"activeAccountConfigIds,omitempty"`
	// The external version of the API
	ApiVersion string `json:"apiVersion"`
	// The AWS regions that this service plan is available on
	AwsRegions []string `json:"awsRegions,omitempty"`
	// The deployment configuration ID
	DeploymentConfigId string `json:"deploymentConfigId"`
	// The GCP regions that this service plan is available on
	GcpRegions []string `json:"gcpRegions,omitempty"`
	// Whether there are any pending changes for the product tier configuration
	HasPendingChanges bool `json:"hasPendingChanges"`
	// Whether the product tier is disabled
	IsProductTierDisabled bool `json:"isProductTierDisabled"`
	// The version number for the latest major version set.
	LatestMajorVersion string `json:"latestMajorVersion"`
	// The model type encapsulating this service
	ModelType string `json:"modelType"`
	// A brief description of the product tier
	ProductTierDescription string `json:"productTierDescription"`
	// Documentation
	ProductTierDocumentation string `json:"productTierDocumentation"`
	// The features that are enabled / disabled for this product tier
	ProductTierFeatures *map[string]bool `json:"productTierFeatures,omitempty"`
	// Product tier ID
	ProductTierId string `json:"productTierId"`
	// Unique Key of the product tier
	ProductTierKey string `json:"productTierKey"`
	// Name of the product tier
	ProductTierName string `json:"productTierName"`
	// A brief description for the end user of the product tier
	ProductTierPlanDescription string `json:"productTierPlanDescription"`
	// Pricing
	ProductTierPricing interface{} `json:"productTierPricing"`
	// Support
	ProductTierSupport string `json:"productTierSupport"`
	// A brief description of the service API bundle
	ServiceApiDescription string `json:"serviceApiDescription"`
	// The service API ID
	ServiceApiId string `json:"serviceApiId"`
	// The service environment ID
	ServiceEnvironmentId string `json:"serviceEnvironmentId"`
	// A brief description of the service model
	ServiceModelDescription string `json:"serviceModelDescription"`
	// Enabled service model features
	ServiceModelFeatures []ServiceModelFeatureDetail `json:"serviceModelFeatures,omitempty"`
	// The service model ID
	ServiceModelId string `json:"serviceModelId"`
	// Name of the Service Model
	ServiceModelName string `json:"serviceModelName"`
	// Tier type
	TierType string `json:"tierType"`
	// The tier version set status.
	VersionSetStatus string `json:"versionSetStatus"`
}

type _GetServicePlanResult GetServicePlanResult

// NewGetServicePlanResult instantiates a new GetServicePlanResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServicePlanResult(autoApproveSubscription bool, apiVersion string, deploymentConfigId string, hasPendingChanges bool, isProductTierDisabled bool, latestMajorVersion string, modelType string, productTierDescription string, productTierDocumentation string, productTierId string, productTierKey string, productTierName string, productTierPlanDescription string, productTierPricing interface{}, productTierSupport string, serviceApiDescription string, serviceApiId string, serviceEnvironmentId string, serviceModelDescription string, serviceModelId string, serviceModelName string, tierType string, versionSetStatus string) *GetServicePlanResult {
	this := GetServicePlanResult{}
	this.AutoApproveSubscription = autoApproveSubscription
	this.ApiVersion = apiVersion
	this.DeploymentConfigId = deploymentConfigId
	this.HasPendingChanges = hasPendingChanges
	this.IsProductTierDisabled = isProductTierDisabled
	this.LatestMajorVersion = latestMajorVersion
	this.ModelType = modelType
	this.ProductTierDescription = productTierDescription
	this.ProductTierDocumentation = productTierDocumentation
	this.ProductTierId = productTierId
	this.ProductTierKey = productTierKey
	this.ProductTierName = productTierName
	this.ProductTierPlanDescription = productTierPlanDescription
	this.ProductTierPricing = productTierPricing
	this.ProductTierSupport = productTierSupport
	this.ServiceApiDescription = serviceApiDescription
	this.ServiceApiId = serviceApiId
	this.ServiceEnvironmentId = serviceEnvironmentId
	this.ServiceModelDescription = serviceModelDescription
	this.ServiceModelId = serviceModelId
	this.ServiceModelName = serviceModelName
	this.TierType = tierType
	this.VersionSetStatus = versionSetStatus
	return &this
}

// NewGetServicePlanResultWithDefaults instantiates a new GetServicePlanResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServicePlanResultWithDefaults() *GetServicePlanResult {
	this := GetServicePlanResult{}
	return &this
}

// GetAutoApproveSubscription returns the AutoApproveSubscription field value
func (o *GetServicePlanResult) GetAutoApproveSubscription() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoApproveSubscription
}

// GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetAutoApproveSubscriptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoApproveSubscription, true
}

// SetAutoApproveSubscription sets field value
func (o *GetServicePlanResult) SetAutoApproveSubscription(v bool) {
	o.AutoApproveSubscription = v
}

// GetAccountConfigIds returns the AccountConfigIds field value if set, zero value otherwise.
func (o *GetServicePlanResult) GetAccountConfigIds() []string {
	if o == nil || IsNil(o.AccountConfigIds) {
		var ret []string
		return ret
	}
	return o.AccountConfigIds
}

// GetAccountConfigIdsOk returns a tuple with the AccountConfigIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetAccountConfigIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AccountConfigIds) {
		return nil, false
	}
	return o.AccountConfigIds, true
}

// SetAccountConfigIds gets a reference to the given []string and assigns it to the AccountConfigIds field.
func (o *GetServicePlanResult) SetAccountConfigIds(v []string) {
	o.AccountConfigIds = v
}

// GetActiveAccountConfigIds returns the ActiveAccountConfigIds field value if set, zero value otherwise.
func (o *GetServicePlanResult) GetActiveAccountConfigIds() map[string]string {
	if o == nil || IsNil(o.ActiveAccountConfigIds) {
		var ret map[string]string
		return ret
	}
	return *o.ActiveAccountConfigIds
}

// GetActiveAccountConfigIdsOk returns a tuple with the ActiveAccountConfigIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetActiveAccountConfigIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ActiveAccountConfigIds) {
		return nil, false
	}
	return o.ActiveAccountConfigIds, true
}

// SetActiveAccountConfigIds gets a reference to the given map[string]string and assigns it to the ActiveAccountConfigIds field.
func (o *GetServicePlanResult) SetActiveAccountConfigIds(v map[string]string) {
	o.ActiveAccountConfigIds = &v
}

// GetApiVersion returns the ApiVersion field value
func (o *GetServicePlanResult) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *GetServicePlanResult) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *GetServicePlanResult) GetAwsRegions() []string {
	if o == nil || IsNil(o.AwsRegions) {
		var ret []string
		return ret
	}
	return o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetAwsRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AwsRegions) {
		return nil, false
	}
	return o.AwsRegions, true
}

// SetAwsRegions gets a reference to the given []string and assigns it to the AwsRegions field.
func (o *GetServicePlanResult) SetAwsRegions(v []string) {
	o.AwsRegions = v
}

// GetDeploymentConfigId returns the DeploymentConfigId field value
func (o *GetServicePlanResult) GetDeploymentConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentConfigId
}

// GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetDeploymentConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentConfigId, true
}

// SetDeploymentConfigId sets field value
func (o *GetServicePlanResult) SetDeploymentConfigId(v string) {
	o.DeploymentConfigId = v
}

// GetGcpRegions returns the GcpRegions field value if set, zero value otherwise.
func (o *GetServicePlanResult) GetGcpRegions() []string {
	if o == nil || IsNil(o.GcpRegions) {
		var ret []string
		return ret
	}
	return o.GcpRegions
}

// GetGcpRegionsOk returns a tuple with the GcpRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetGcpRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.GcpRegions) {
		return nil, false
	}
	return o.GcpRegions, true
}

// SetGcpRegions gets a reference to the given []string and assigns it to the GcpRegions field.
func (o *GetServicePlanResult) SetGcpRegions(v []string) {
	o.GcpRegions = v
}

// GetHasPendingChanges returns the HasPendingChanges field value
func (o *GetServicePlanResult) GetHasPendingChanges() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPendingChanges
}

// GetHasPendingChangesOk returns a tuple with the HasPendingChanges field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetHasPendingChangesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPendingChanges, true
}

// SetHasPendingChanges sets field value
func (o *GetServicePlanResult) SetHasPendingChanges(v bool) {
	o.HasPendingChanges = v
}

// GetIsProductTierDisabled returns the IsProductTierDisabled field value
func (o *GetServicePlanResult) GetIsProductTierDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsProductTierDisabled
}

// GetIsProductTierDisabledOk returns a tuple with the IsProductTierDisabled field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetIsProductTierDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsProductTierDisabled, true
}

// SetIsProductTierDisabled sets field value
func (o *GetServicePlanResult) SetIsProductTierDisabled(v bool) {
	o.IsProductTierDisabled = v
}

// GetLatestMajorVersion returns the LatestMajorVersion field value
func (o *GetServicePlanResult) GetLatestMajorVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LatestMajorVersion
}

// GetLatestMajorVersionOk returns a tuple with the LatestMajorVersion field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetLatestMajorVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LatestMajorVersion, true
}

// SetLatestMajorVersion sets field value
func (o *GetServicePlanResult) SetLatestMajorVersion(v string) {
	o.LatestMajorVersion = v
}

// GetModelType returns the ModelType field value
func (o *GetServicePlanResult) GetModelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetModelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelType, true
}

// SetModelType sets field value
func (o *GetServicePlanResult) SetModelType(v string) {
	o.ModelType = v
}

// GetProductTierDescription returns the ProductTierDescription field value
func (o *GetServicePlanResult) GetProductTierDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierDescription
}

// GetProductTierDescriptionOk returns a tuple with the ProductTierDescription field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetProductTierDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierDescription, true
}

// SetProductTierDescription sets field value
func (o *GetServicePlanResult) SetProductTierDescription(v string) {
	o.ProductTierDescription = v
}

// GetProductTierDocumentation returns the ProductTierDocumentation field value
func (o *GetServicePlanResult) GetProductTierDocumentation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierDocumentation
}

// GetProductTierDocumentationOk returns a tuple with the ProductTierDocumentation field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetProductTierDocumentationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierDocumentation, true
}

// SetProductTierDocumentation sets field value
func (o *GetServicePlanResult) SetProductTierDocumentation(v string) {
	o.ProductTierDocumentation = v
}

// GetProductTierFeatures returns the ProductTierFeatures field value if set, zero value otherwise.
func (o *GetServicePlanResult) GetProductTierFeatures() map[string]bool {
	if o == nil || IsNil(o.ProductTierFeatures) {
		var ret map[string]bool
		return ret
	}
	return *o.ProductTierFeatures
}

// GetProductTierFeaturesOk returns a tuple with the ProductTierFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetProductTierFeaturesOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.ProductTierFeatures) {
		return nil, false
	}
	return o.ProductTierFeatures, true
}

// SetProductTierFeatures gets a reference to the given map[string]bool and assigns it to the ProductTierFeatures field.
func (o *GetServicePlanResult) SetProductTierFeatures(v map[string]bool) {
	o.ProductTierFeatures = &v
}

// GetProductTierId returns the ProductTierId field value
func (o *GetServicePlanResult) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *GetServicePlanResult) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetProductTierKey returns the ProductTierKey field value
func (o *GetServicePlanResult) GetProductTierKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierKey
}

// GetProductTierKeyOk returns a tuple with the ProductTierKey field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetProductTierKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierKey, true
}

// SetProductTierKey sets field value
func (o *GetServicePlanResult) SetProductTierKey(v string) {
	o.ProductTierKey = v
}

// GetProductTierName returns the ProductTierName field value
func (o *GetServicePlanResult) GetProductTierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierName
}

// GetProductTierNameOk returns a tuple with the ProductTierName field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetProductTierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierName, true
}

// SetProductTierName sets field value
func (o *GetServicePlanResult) SetProductTierName(v string) {
	o.ProductTierName = v
}

// GetProductTierPlanDescription returns the ProductTierPlanDescription field value
func (o *GetServicePlanResult) GetProductTierPlanDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierPlanDescription
}

// GetProductTierPlanDescriptionOk returns a tuple with the ProductTierPlanDescription field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetProductTierPlanDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierPlanDescription, true
}

// SetProductTierPlanDescription sets field value
func (o *GetServicePlanResult) SetProductTierPlanDescription(v string) {
	o.ProductTierPlanDescription = v
}

// GetProductTierPricing returns the ProductTierPricing field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GetServicePlanResult) GetProductTierPricing() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ProductTierPricing
}

// GetProductTierPricingOk returns a tuple with the ProductTierPricing field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetServicePlanResult) GetProductTierPricingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ProductTierPricing) {
		return nil, false
	}
	return &o.ProductTierPricing, true
}

// SetProductTierPricing sets field value
func (o *GetServicePlanResult) SetProductTierPricing(v interface{}) {
	o.ProductTierPricing = v
}

// GetProductTierSupport returns the ProductTierSupport field value
func (o *GetServicePlanResult) GetProductTierSupport() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierSupport
}

// GetProductTierSupportOk returns a tuple with the ProductTierSupport field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetProductTierSupportOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierSupport, true
}

// SetProductTierSupport sets field value
func (o *GetServicePlanResult) SetProductTierSupport(v string) {
	o.ProductTierSupport = v
}

// GetServiceApiDescription returns the ServiceApiDescription field value
func (o *GetServicePlanResult) GetServiceApiDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceApiDescription
}

// GetServiceApiDescriptionOk returns a tuple with the ServiceApiDescription field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetServiceApiDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceApiDescription, true
}

// SetServiceApiDescription sets field value
func (o *GetServicePlanResult) SetServiceApiDescription(v string) {
	o.ServiceApiDescription = v
}

// GetServiceApiId returns the ServiceApiId field value
func (o *GetServicePlanResult) GetServiceApiId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceApiId
}

// GetServiceApiIdOk returns a tuple with the ServiceApiId field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetServiceApiIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceApiId, true
}

// SetServiceApiId sets field value
func (o *GetServicePlanResult) SetServiceApiId(v string) {
	o.ServiceApiId = v
}

// GetServiceEnvironmentId returns the ServiceEnvironmentId field value
func (o *GetServicePlanResult) GetServiceEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentId
}

// GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetServiceEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentId, true
}

// SetServiceEnvironmentId sets field value
func (o *GetServicePlanResult) SetServiceEnvironmentId(v string) {
	o.ServiceEnvironmentId = v
}

// GetServiceModelDescription returns the ServiceModelDescription field value
func (o *GetServicePlanResult) GetServiceModelDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelDescription
}

// GetServiceModelDescriptionOk returns a tuple with the ServiceModelDescription field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetServiceModelDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelDescription, true
}

// SetServiceModelDescription sets field value
func (o *GetServicePlanResult) SetServiceModelDescription(v string) {
	o.ServiceModelDescription = v
}

// GetServiceModelFeatures returns the ServiceModelFeatures field value if set, zero value otherwise.
func (o *GetServicePlanResult) GetServiceModelFeatures() []ServiceModelFeatureDetail {
	if o == nil || IsNil(o.ServiceModelFeatures) {
		var ret []ServiceModelFeatureDetail
		return ret
	}
	return o.ServiceModelFeatures
}

// GetServiceModelFeaturesOk returns a tuple with the ServiceModelFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetServiceModelFeaturesOk() ([]ServiceModelFeatureDetail, bool) {
	if o == nil || IsNil(o.ServiceModelFeatures) {
		return nil, false
	}
	return o.ServiceModelFeatures, true
}

// SetServiceModelFeatures gets a reference to the given []ServiceModelFeatureDetail and assigns it to the ServiceModelFeatures field.
func (o *GetServicePlanResult) SetServiceModelFeatures(v []ServiceModelFeatureDetail) {
	o.ServiceModelFeatures = v
}

// GetServiceModelId returns the ServiceModelId field value
func (o *GetServicePlanResult) GetServiceModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelId
}

// GetServiceModelIdOk returns a tuple with the ServiceModelId field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetServiceModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelId, true
}

// SetServiceModelId sets field value
func (o *GetServicePlanResult) SetServiceModelId(v string) {
	o.ServiceModelId = v
}

// GetServiceModelName returns the ServiceModelName field value
func (o *GetServicePlanResult) GetServiceModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelName
}

// GetServiceModelNameOk returns a tuple with the ServiceModelName field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetServiceModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelName, true
}

// SetServiceModelName sets field value
func (o *GetServicePlanResult) SetServiceModelName(v string) {
	o.ServiceModelName = v
}

// GetTierType returns the TierType field value
func (o *GetServicePlanResult) GetTierType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TierType
}

// GetTierTypeOk returns a tuple with the TierType field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetTierTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TierType, true
}

// SetTierType sets field value
func (o *GetServicePlanResult) SetTierType(v string) {
	o.TierType = v
}

// GetVersionSetStatus returns the VersionSetStatus field value
func (o *GetServicePlanResult) GetVersionSetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionSetStatus
}

// GetVersionSetStatusOk returns a tuple with the VersionSetStatus field value
// and a boolean to check if the value has been set.
func (o *GetServicePlanResult) GetVersionSetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionSetStatus, true
}

// SetVersionSetStatus sets field value
func (o *GetServicePlanResult) SetVersionSetStatus(v string) {
	o.VersionSetStatus = v
}

func (o GetServicePlanResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServicePlanResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AutoApproveSubscription"] = o.AutoApproveSubscription
	if !IsNil(o.AccountConfigIds) {
		toSerialize["accountConfigIds"] = o.AccountConfigIds
	}
	if !IsNil(o.ActiveAccountConfigIds) {
		toSerialize["activeAccountConfigIds"] = o.ActiveAccountConfigIds
	}
	toSerialize["apiVersion"] = o.ApiVersion
	if !IsNil(o.AwsRegions) {
		toSerialize["awsRegions"] = o.AwsRegions
	}
	toSerialize["deploymentConfigId"] = o.DeploymentConfigId
	if !IsNil(o.GcpRegions) {
		toSerialize["gcpRegions"] = o.GcpRegions
	}
	toSerialize["hasPendingChanges"] = o.HasPendingChanges
	toSerialize["isProductTierDisabled"] = o.IsProductTierDisabled
	toSerialize["latestMajorVersion"] = o.LatestMajorVersion
	toSerialize["modelType"] = o.ModelType
	toSerialize["productTierDescription"] = o.ProductTierDescription
	toSerialize["productTierDocumentation"] = o.ProductTierDocumentation
	if !IsNil(o.ProductTierFeatures) {
		toSerialize["productTierFeatures"] = o.ProductTierFeatures
	}
	toSerialize["productTierId"] = o.ProductTierId
	toSerialize["productTierKey"] = o.ProductTierKey
	toSerialize["productTierName"] = o.ProductTierName
	toSerialize["productTierPlanDescription"] = o.ProductTierPlanDescription
	if o.ProductTierPricing != nil {
		toSerialize["productTierPricing"] = o.ProductTierPricing
	}
	toSerialize["productTierSupport"] = o.ProductTierSupport
	toSerialize["serviceApiDescription"] = o.ServiceApiDescription
	toSerialize["serviceApiId"] = o.ServiceApiId
	toSerialize["serviceEnvironmentId"] = o.ServiceEnvironmentId
	toSerialize["serviceModelDescription"] = o.ServiceModelDescription
	if !IsNil(o.ServiceModelFeatures) {
		toSerialize["serviceModelFeatures"] = o.ServiceModelFeatures
	}
	toSerialize["serviceModelId"] = o.ServiceModelId
	toSerialize["serviceModelName"] = o.ServiceModelName
	toSerialize["tierType"] = o.TierType
	toSerialize["versionSetStatus"] = o.VersionSetStatus
	return toSerialize, nil
}

func (o *GetServicePlanResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AutoApproveSubscription",
		"apiVersion",
		"deploymentConfigId",
		"hasPendingChanges",
		"isProductTierDisabled",
		"latestMajorVersion",
		"modelType",
		"productTierDescription",
		"productTierDocumentation",
		"productTierId",
		"productTierKey",
		"productTierName",
		"productTierPlanDescription",
		"productTierPricing",
		"productTierSupport",
		"serviceApiDescription",
		"serviceApiId",
		"serviceEnvironmentId",
		"serviceModelDescription",
		"serviceModelId",
		"serviceModelName",
		"tierType",
		"versionSetStatus",
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

	varGetServicePlanResult := _GetServicePlanResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetServicePlanResult)

	if err != nil {
		return err
	}

	*o = GetServicePlanResult(varGetServicePlanResult)

	return err
}

type NullableGetServicePlanResult struct {
	value *GetServicePlanResult
	isSet bool
}

func (v NullableGetServicePlanResult) Get() *GetServicePlanResult {
	return v.value
}

func (v *NullableGetServicePlanResult) Set(val *GetServicePlanResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServicePlanResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServicePlanResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServicePlanResult(val *GetServicePlanResult) *NullableGetServicePlanResult {
	return &NullableGetServicePlanResult{value: val, isSet: true}
}

func (v NullableGetServicePlanResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServicePlanResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


