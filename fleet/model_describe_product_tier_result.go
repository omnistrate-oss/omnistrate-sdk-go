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

// checks if the DescribeProductTierResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeProductTierResult{}

// DescribeProductTierResult struct for DescribeProductTierResult
type DescribeProductTierResult struct {
	// The resources that belong to this service API bundle and their active versions
	ApiGroups map[string]interface{} `json:"apiGroups,omitempty"`
	// Auto approve subscription or not
	AutoApproveSubscription *bool `json:"autoApproveSubscription,omitempty"`
	// The AWS regions that this product tier is available on
	AwsRegions []string `json:"awsRegions,omitempty"`
	// The readiness of the cloud providers configurations
	CloudProvidersConfigReadiness map[string]interface{} `json:"cloudProvidersConfigReadiness,omitempty"`
	// A brief description of the product tier
	Description string `json:"description"`
	// Documentation
	Documentation string `json:"documentation"`
	// The features that are enabled for this product tier, including scope details and configuration
	EnabledFeatures []ProductTierFeatureDetail `json:"enabledFeatures,omitempty"`
	// The features that are enabled / disabled for this product tier
	Features map[string]interface{} `json:"features,omitempty"`
	// The GCP regions that this product tier is available on
	GcpRegions []string `json:"gcpRegions,omitempty"`
	// ID of a Product Tier
	Id string `json:"id"`
	// Flag to indicate if the product tier is disabled.
	IsDisabled bool `json:"isDisabled"`
	// Unique Key of the product tier
	Key string `json:"key"`
	// Name of the product tier
	Name string `json:"name"`
	// A brief description for the end user of the product tier
	PlanDescription string `json:"planDescription"`
	// Pricing
	Pricing interface{} `json:"pricing"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// ID of a Service Model
	ServiceModelId string `json:"serviceModelId"`
	// Support
	Support string `json:"support"`
	// ProductTierType is the type of tier for a product
	TierType string `json:"tierType"`
	AdditionalProperties map[string]interface{}
}

type _DescribeProductTierResult DescribeProductTierResult

// NewDescribeProductTierResult instantiates a new DescribeProductTierResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProductTierResult(description string, documentation string, id string, isDisabled bool, key string, name string, planDescription string, pricing interface{}, serviceId string, serviceModelId string, support string, tierType string) *DescribeProductTierResult {
	this := DescribeProductTierResult{}
	this.Description = description
	this.Documentation = documentation
	this.Id = id
	this.IsDisabled = isDisabled
	this.Key = key
	this.Name = name
	this.PlanDescription = planDescription
	this.Pricing = pricing
	this.ServiceId = serviceId
	this.ServiceModelId = serviceModelId
	this.Support = support
	this.TierType = tierType
	return &this
}

// NewDescribeProductTierResultWithDefaults instantiates a new DescribeProductTierResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProductTierResultWithDefaults() *DescribeProductTierResult {
	this := DescribeProductTierResult{}
	return &this
}

// GetApiGroups returns the ApiGroups field value if set, zero value otherwise.
func (o *DescribeProductTierResult) GetApiGroups() map[string]interface{} {
	if o == nil || IsNil(o.ApiGroups) {
		var ret map[string]interface{}
		return ret
	}
	return o.ApiGroups
}

// GetApiGroupsOk returns a tuple with the ApiGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetApiGroupsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ApiGroups) {
		return map[string]interface{}{}, false
	}
	return o.ApiGroups, true
}

// HasApiGroups returns a boolean if a field has been set.
func (o *DescribeProductTierResult) HasApiGroups() bool {
	if o != nil && !IsNil(o.ApiGroups) {
		return true
	}

	return false
}

// SetApiGroups gets a reference to the given map[string]interface{} and assigns it to the ApiGroups field.
func (o *DescribeProductTierResult) SetApiGroups(v map[string]interface{}) {
	o.ApiGroups = v
}

// GetAutoApproveSubscription returns the AutoApproveSubscription field value if set, zero value otherwise.
func (o *DescribeProductTierResult) GetAutoApproveSubscription() bool {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		var ret bool
		return ret
	}
	return *o.AutoApproveSubscription
}

// GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetAutoApproveSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		return nil, false
	}
	return o.AutoApproveSubscription, true
}

// HasAutoApproveSubscription returns a boolean if a field has been set.
func (o *DescribeProductTierResult) HasAutoApproveSubscription() bool {
	if o != nil && !IsNil(o.AutoApproveSubscription) {
		return true
	}

	return false
}

// SetAutoApproveSubscription gets a reference to the given bool and assigns it to the AutoApproveSubscription field.
func (o *DescribeProductTierResult) SetAutoApproveSubscription(v bool) {
	o.AutoApproveSubscription = &v
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *DescribeProductTierResult) GetAwsRegions() []string {
	if o == nil || IsNil(o.AwsRegions) {
		var ret []string
		return ret
	}
	return o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetAwsRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AwsRegions) {
		return nil, false
	}
	return o.AwsRegions, true
}

// HasAwsRegions returns a boolean if a field has been set.
func (o *DescribeProductTierResult) HasAwsRegions() bool {
	if o != nil && !IsNil(o.AwsRegions) {
		return true
	}

	return false
}

// SetAwsRegions gets a reference to the given []string and assigns it to the AwsRegions field.
func (o *DescribeProductTierResult) SetAwsRegions(v []string) {
	o.AwsRegions = v
}

// GetCloudProvidersConfigReadiness returns the CloudProvidersConfigReadiness field value if set, zero value otherwise.
func (o *DescribeProductTierResult) GetCloudProvidersConfigReadiness() map[string]interface{} {
	if o == nil || IsNil(o.CloudProvidersConfigReadiness) {
		var ret map[string]interface{}
		return ret
	}
	return o.CloudProvidersConfigReadiness
}

// GetCloudProvidersConfigReadinessOk returns a tuple with the CloudProvidersConfigReadiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetCloudProvidersConfigReadinessOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CloudProvidersConfigReadiness) {
		return map[string]interface{}{}, false
	}
	return o.CloudProvidersConfigReadiness, true
}

// HasCloudProvidersConfigReadiness returns a boolean if a field has been set.
func (o *DescribeProductTierResult) HasCloudProvidersConfigReadiness() bool {
	if o != nil && !IsNil(o.CloudProvidersConfigReadiness) {
		return true
	}

	return false
}

// SetCloudProvidersConfigReadiness gets a reference to the given map[string]interface{} and assigns it to the CloudProvidersConfigReadiness field.
func (o *DescribeProductTierResult) SetCloudProvidersConfigReadiness(v map[string]interface{}) {
	o.CloudProvidersConfigReadiness = v
}

// GetDescription returns the Description field value
func (o *DescribeProductTierResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeProductTierResult) SetDescription(v string) {
	o.Description = v
}

// GetDocumentation returns the Documentation field value
func (o *DescribeProductTierResult) GetDocumentation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetDocumentationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Documentation, true
}

// SetDocumentation sets field value
func (o *DescribeProductTierResult) SetDocumentation(v string) {
	o.Documentation = v
}

// GetEnabledFeatures returns the EnabledFeatures field value if set, zero value otherwise.
func (o *DescribeProductTierResult) GetEnabledFeatures() []ProductTierFeatureDetail {
	if o == nil || IsNil(o.EnabledFeatures) {
		var ret []ProductTierFeatureDetail
		return ret
	}
	return o.EnabledFeatures
}

// GetEnabledFeaturesOk returns a tuple with the EnabledFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetEnabledFeaturesOk() ([]ProductTierFeatureDetail, bool) {
	if o == nil || IsNil(o.EnabledFeatures) {
		return nil, false
	}
	return o.EnabledFeatures, true
}

// HasEnabledFeatures returns a boolean if a field has been set.
func (o *DescribeProductTierResult) HasEnabledFeatures() bool {
	if o != nil && !IsNil(o.EnabledFeatures) {
		return true
	}

	return false
}

// SetEnabledFeatures gets a reference to the given []ProductTierFeatureDetail and assigns it to the EnabledFeatures field.
func (o *DescribeProductTierResult) SetEnabledFeatures(v []ProductTierFeatureDetail) {
	o.EnabledFeatures = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *DescribeProductTierResult) GetFeatures() map[string]interface{} {
	if o == nil || IsNil(o.Features) {
		var ret map[string]interface{}
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetFeaturesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Features) {
		return map[string]interface{}{}, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *DescribeProductTierResult) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given map[string]interface{} and assigns it to the Features field.
func (o *DescribeProductTierResult) SetFeatures(v map[string]interface{}) {
	o.Features = v
}

// GetGcpRegions returns the GcpRegions field value if set, zero value otherwise.
func (o *DescribeProductTierResult) GetGcpRegions() []string {
	if o == nil || IsNil(o.GcpRegions) {
		var ret []string
		return ret
	}
	return o.GcpRegions
}

// GetGcpRegionsOk returns a tuple with the GcpRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetGcpRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.GcpRegions) {
		return nil, false
	}
	return o.GcpRegions, true
}

// HasGcpRegions returns a boolean if a field has been set.
func (o *DescribeProductTierResult) HasGcpRegions() bool {
	if o != nil && !IsNil(o.GcpRegions) {
		return true
	}

	return false
}

// SetGcpRegions gets a reference to the given []string and assigns it to the GcpRegions field.
func (o *DescribeProductTierResult) SetGcpRegions(v []string) {
	o.GcpRegions = v
}

// GetId returns the Id field value
func (o *DescribeProductTierResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeProductTierResult) SetId(v string) {
	o.Id = v
}

// GetIsDisabled returns the IsDisabled field value
func (o *DescribeProductTierResult) GetIsDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetIsDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDisabled, true
}

// SetIsDisabled sets field value
func (o *DescribeProductTierResult) SetIsDisabled(v bool) {
	o.IsDisabled = v
}

// GetKey returns the Key field value
func (o *DescribeProductTierResult) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DescribeProductTierResult) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *DescribeProductTierResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeProductTierResult) SetName(v string) {
	o.Name = v
}

// GetPlanDescription returns the PlanDescription field value
func (o *DescribeProductTierResult) GetPlanDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanDescription
}

// GetPlanDescriptionOk returns a tuple with the PlanDescription field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetPlanDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanDescription, true
}

// SetPlanDescription sets field value
func (o *DescribeProductTierResult) SetPlanDescription(v string) {
	o.PlanDescription = v
}

// GetPricing returns the Pricing field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *DescribeProductTierResult) GetPricing() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribeProductTierResult) GetPricingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Pricing) {
		return nil, false
	}
	return &o.Pricing, true
}

// SetPricing sets field value
func (o *DescribeProductTierResult) SetPricing(v interface{}) {
	o.Pricing = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeProductTierResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeProductTierResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetServiceModelId returns the ServiceModelId field value
func (o *DescribeProductTierResult) GetServiceModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelId
}

// GetServiceModelIdOk returns a tuple with the ServiceModelId field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetServiceModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelId, true
}

// SetServiceModelId sets field value
func (o *DescribeProductTierResult) SetServiceModelId(v string) {
	o.ServiceModelId = v
}

// GetSupport returns the Support field value
func (o *DescribeProductTierResult) GetSupport() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Support
}

// GetSupportOk returns a tuple with the Support field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetSupportOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Support, true
}

// SetSupport sets field value
func (o *DescribeProductTierResult) SetSupport(v string) {
	o.Support = v
}

// GetTierType returns the TierType field value
func (o *DescribeProductTierResult) GetTierType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TierType
}

// GetTierTypeOk returns a tuple with the TierType field value
// and a boolean to check if the value has been set.
func (o *DescribeProductTierResult) GetTierTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TierType, true
}

// SetTierType sets field value
func (o *DescribeProductTierResult) SetTierType(v string) {
	o.TierType = v
}

func (o DescribeProductTierResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProductTierResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiGroups) {
		toSerialize["apiGroups"] = o.ApiGroups
	}
	if !IsNil(o.AutoApproveSubscription) {
		toSerialize["autoApproveSubscription"] = o.AutoApproveSubscription
	}
	if !IsNil(o.AwsRegions) {
		toSerialize["awsRegions"] = o.AwsRegions
	}
	if !IsNil(o.CloudProvidersConfigReadiness) {
		toSerialize["cloudProvidersConfigReadiness"] = o.CloudProvidersConfigReadiness
	}
	toSerialize["description"] = o.Description
	toSerialize["documentation"] = o.Documentation
	if !IsNil(o.EnabledFeatures) {
		toSerialize["enabledFeatures"] = o.EnabledFeatures
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.GcpRegions) {
		toSerialize["gcpRegions"] = o.GcpRegions
	}
	toSerialize["id"] = o.Id
	toSerialize["isDisabled"] = o.IsDisabled
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["planDescription"] = o.PlanDescription
	if o.Pricing != nil {
		toSerialize["pricing"] = o.Pricing
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["serviceModelId"] = o.ServiceModelId
	toSerialize["support"] = o.Support
	toSerialize["tierType"] = o.TierType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeProductTierResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"documentation",
		"id",
		"isDisabled",
		"key",
		"name",
		"planDescription",
		"pricing",
		"serviceId",
		"serviceModelId",
		"support",
		"tierType",
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

	varDescribeProductTierResult := _DescribeProductTierResult{}

	err = json.Unmarshal(data, &varDescribeProductTierResult)

	if err != nil {
		return err
	}

	*o = DescribeProductTierResult(varDescribeProductTierResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiGroups")
		delete(additionalProperties, "autoApproveSubscription")
		delete(additionalProperties, "awsRegions")
		delete(additionalProperties, "cloudProvidersConfigReadiness")
		delete(additionalProperties, "description")
		delete(additionalProperties, "documentation")
		delete(additionalProperties, "enabledFeatures")
		delete(additionalProperties, "features")
		delete(additionalProperties, "gcpRegions")
		delete(additionalProperties, "id")
		delete(additionalProperties, "isDisabled")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "planDescription")
		delete(additionalProperties, "pricing")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "serviceModelId")
		delete(additionalProperties, "support")
		delete(additionalProperties, "tierType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeProductTierResult struct {
	value *DescribeProductTierResult
	isSet bool
}

func (v NullableDescribeProductTierResult) Get() *DescribeProductTierResult {
	return v.value
}

func (v *NullableDescribeProductTierResult) Set(val *DescribeProductTierResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProductTierResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProductTierResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProductTierResult(val *DescribeProductTierResult) *NullableDescribeProductTierResult {
	return &NullableDescribeProductTierResult{value: val, isSet: true}
}

func (v NullableDescribeProductTierResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProductTierResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


