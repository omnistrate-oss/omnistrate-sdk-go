/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateProductTierRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductTierRequest2{}

// CreateProductTierRequest2 struct for CreateProductTierRequest2
type CreateProductTierRequest2 struct {
	// Allow creates when payment not configured
	AllowCreatesWhenPaymentNotConfigured *bool `json:"allowCreatesWhenPaymentNotConfigured,omitempty"`
	// Auto approve subscription or not
	AutoApproveSubscription *bool `json:"autoApproveSubscription,omitempty"`
	// The AWS regions that this product tier is available on
	AwsRegions []string `json:"awsRegions,omitempty"`
	// The Azure regions that this product tier is available on
	AzureRegions []string `json:"azureRegions,omitempty"`
	// Optional billing product ID for tax purposes
	BillingProductID *string `json:"billingProductID,omitempty"`
	// List of billing providers to be used for the product tier
	BillingProviders []string `json:"billingProviders,omitempty"`
	// The default billing provider to be used for the product tier
	DefaultBillingProvider *string `json:"defaultBillingProvider,omitempty"`
	// A brief description of the product tier
	Description string `json:"description"`
	// Documentation
	Documentation *string `json:"documentation,omitempty"`
	// Export usage metering data
	ExportUsageMetering *bool `json:"exportUsageMetering,omitempty"`
	// Export usage metering data configuration
	ExportUsageMeteringConfig map[string]interface{} `json:"exportUsageMeteringConfig,omitempty"`
	// The GCP regions that this product tier is available on
	GcpRegions []string `json:"gcpRegions,omitempty"`
	// Create the product tier in a disabled state. Enabling the product tier will let end-customers subscribe and use the service plan.
	IsDisabled *bool `json:"isDisabled,omitempty"`
	// Maximum number of instances
	MaxNumberOfInstances *int64 `json:"maxNumberOfInstances,omitempty"`
	// Name of the product tier
	Name string `json:"name"`
	// A brief description for the end user of the product tier
	PlanDescription string `json:"planDescription"`
	// Price per unit.
	PricePerUnit map[string]interface{} `json:"pricePerUnit,omitempty"`
	// Pricing
	Pricing interface{} `json:"pricing,omitempty"`
	// The private regions that this product tier is available on
	PrivateRegions []string `json:"privateRegions,omitempty"`
	// Service model ID
	ServiceModelId string `json:"serviceModelId"`
	// Support
	Support *string `json:"support,omitempty"`
	// Tier type
	TierType string `json:"tierType"`
	AdditionalProperties map[string]interface{}
}

type _CreateProductTierRequest2 CreateProductTierRequest2

// NewCreateProductTierRequest2 instantiates a new CreateProductTierRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductTierRequest2(description string, name string, planDescription string, serviceModelId string, tierType string) *CreateProductTierRequest2 {
	this := CreateProductTierRequest2{}
	this.Description = description
	this.Name = name
	this.PlanDescription = planDescription
	this.ServiceModelId = serviceModelId
	this.TierType = tierType
	return &this
}

// NewCreateProductTierRequest2WithDefaults instantiates a new CreateProductTierRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductTierRequest2WithDefaults() *CreateProductTierRequest2 {
	this := CreateProductTierRequest2{}
	return &this
}

// GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetAllowCreatesWhenPaymentNotConfigured() bool {
	if o == nil || IsNil(o.AllowCreatesWhenPaymentNotConfigured) {
		var ret bool
		return ret
	}
	return *o.AllowCreatesWhenPaymentNotConfigured
}

// GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCreatesWhenPaymentNotConfigured) {
		return nil, false
	}
	return o.AllowCreatesWhenPaymentNotConfigured, true
}

// SetAllowCreatesWhenPaymentNotConfigured gets a reference to the given bool and assigns it to the AllowCreatesWhenPaymentNotConfigured field.
func (o *CreateProductTierRequest2) SetAllowCreatesWhenPaymentNotConfigured(v bool) {
	o.AllowCreatesWhenPaymentNotConfigured = &v
}

// GetAutoApproveSubscription returns the AutoApproveSubscription field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetAutoApproveSubscription() bool {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		var ret bool
		return ret
	}
	return *o.AutoApproveSubscription
}

// GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetAutoApproveSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		return nil, false
	}
	return o.AutoApproveSubscription, true
}

// SetAutoApproveSubscription gets a reference to the given bool and assigns it to the AutoApproveSubscription field.
func (o *CreateProductTierRequest2) SetAutoApproveSubscription(v bool) {
	o.AutoApproveSubscription = &v
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetAwsRegions() []string {
	if o == nil || IsNil(o.AwsRegions) {
		var ret []string
		return ret
	}
	return o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetAwsRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AwsRegions) {
		return nil, false
	}
	return o.AwsRegions, true
}

// SetAwsRegions gets a reference to the given []string and assigns it to the AwsRegions field.
func (o *CreateProductTierRequest2) SetAwsRegions(v []string) {
	o.AwsRegions = v
}

// GetAzureRegions returns the AzureRegions field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetAzureRegions() []string {
	if o == nil || IsNil(o.AzureRegions) {
		var ret []string
		return ret
	}
	return o.AzureRegions
}

// GetAzureRegionsOk returns a tuple with the AzureRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetAzureRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AzureRegions) {
		return nil, false
	}
	return o.AzureRegions, true
}

// SetAzureRegions gets a reference to the given []string and assigns it to the AzureRegions field.
func (o *CreateProductTierRequest2) SetAzureRegions(v []string) {
	o.AzureRegions = v
}

// GetBillingProductID returns the BillingProductID field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetBillingProductID() string {
	if o == nil || IsNil(o.BillingProductID) {
		var ret string
		return ret
	}
	return *o.BillingProductID
}

// GetBillingProductIDOk returns a tuple with the BillingProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetBillingProductIDOk() (*string, bool) {
	if o == nil || IsNil(o.BillingProductID) {
		return nil, false
	}
	return o.BillingProductID, true
}

// SetBillingProductID gets a reference to the given string and assigns it to the BillingProductID field.
func (o *CreateProductTierRequest2) SetBillingProductID(v string) {
	o.BillingProductID = &v
}

// GetBillingProviders returns the BillingProviders field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetBillingProviders() []string {
	if o == nil || IsNil(o.BillingProviders) {
		var ret []string
		return ret
	}
	return o.BillingProviders
}

// GetBillingProvidersOk returns a tuple with the BillingProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetBillingProvidersOk() ([]string, bool) {
	if o == nil || IsNil(o.BillingProviders) {
		return nil, false
	}
	return o.BillingProviders, true
}

// SetBillingProviders gets a reference to the given []string and assigns it to the BillingProviders field.
func (o *CreateProductTierRequest2) SetBillingProviders(v []string) {
	o.BillingProviders = v
}

// GetDefaultBillingProvider returns the DefaultBillingProvider field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetDefaultBillingProvider() string {
	if o == nil || IsNil(o.DefaultBillingProvider) {
		var ret string
		return ret
	}
	return *o.DefaultBillingProvider
}

// GetDefaultBillingProviderOk returns a tuple with the DefaultBillingProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetDefaultBillingProviderOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBillingProvider) {
		return nil, false
	}
	return o.DefaultBillingProvider, true
}

// SetDefaultBillingProvider gets a reference to the given string and assigns it to the DefaultBillingProvider field.
func (o *CreateProductTierRequest2) SetDefaultBillingProvider(v string) {
	o.DefaultBillingProvider = &v
}

// GetDescription returns the Description field value
func (o *CreateProductTierRequest2) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateProductTierRequest2) SetDescription(v string) {
	o.Description = v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation) {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *CreateProductTierRequest2) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetExportUsageMetering returns the ExportUsageMetering field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetExportUsageMetering() bool {
	if o == nil || IsNil(o.ExportUsageMetering) {
		var ret bool
		return ret
	}
	return *o.ExportUsageMetering
}

// GetExportUsageMeteringOk returns a tuple with the ExportUsageMetering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetExportUsageMeteringOk() (*bool, bool) {
	if o == nil || IsNil(o.ExportUsageMetering) {
		return nil, false
	}
	return o.ExportUsageMetering, true
}

// SetExportUsageMetering gets a reference to the given bool and assigns it to the ExportUsageMetering field.
func (o *CreateProductTierRequest2) SetExportUsageMetering(v bool) {
	o.ExportUsageMetering = &v
}

// GetExportUsageMeteringConfig returns the ExportUsageMeteringConfig field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetExportUsageMeteringConfig() map[string]interface{} {
	if o == nil || IsNil(o.ExportUsageMeteringConfig) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExportUsageMeteringConfig
}

// GetExportUsageMeteringConfigOk returns a tuple with the ExportUsageMeteringConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetExportUsageMeteringConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExportUsageMeteringConfig) {
		return map[string]interface{}{}, false
	}
	return o.ExportUsageMeteringConfig, true
}

// SetExportUsageMeteringConfig gets a reference to the given map[string]interface{} and assigns it to the ExportUsageMeteringConfig field.
func (o *CreateProductTierRequest2) SetExportUsageMeteringConfig(v map[string]interface{}) {
	o.ExportUsageMeteringConfig = v
}

// GetGcpRegions returns the GcpRegions field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetGcpRegions() []string {
	if o == nil || IsNil(o.GcpRegions) {
		var ret []string
		return ret
	}
	return o.GcpRegions
}

// GetGcpRegionsOk returns a tuple with the GcpRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetGcpRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.GcpRegions) {
		return nil, false
	}
	return o.GcpRegions, true
}

// SetGcpRegions gets a reference to the given []string and assigns it to the GcpRegions field.
func (o *CreateProductTierRequest2) SetGcpRegions(v []string) {
	o.GcpRegions = v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *CreateProductTierRequest2) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetMaxNumberOfInstances returns the MaxNumberOfInstances field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetMaxNumberOfInstances() int64 {
	if o == nil || IsNil(o.MaxNumberOfInstances) {
		var ret int64
		return ret
	}
	return *o.MaxNumberOfInstances
}

// GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetMaxNumberOfInstancesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxNumberOfInstances) {
		return nil, false
	}
	return o.MaxNumberOfInstances, true
}

// SetMaxNumberOfInstances gets a reference to the given int64 and assigns it to the MaxNumberOfInstances field.
func (o *CreateProductTierRequest2) SetMaxNumberOfInstances(v int64) {
	o.MaxNumberOfInstances = &v
}

// GetName returns the Name field value
func (o *CreateProductTierRequest2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProductTierRequest2) SetName(v string) {
	o.Name = v
}

// GetPlanDescription returns the PlanDescription field value
func (o *CreateProductTierRequest2) GetPlanDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanDescription
}

// GetPlanDescriptionOk returns a tuple with the PlanDescription field value
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetPlanDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanDescription, true
}

// SetPlanDescription sets field value
func (o *CreateProductTierRequest2) SetPlanDescription(v string) {
	o.PlanDescription = v
}

// GetPricePerUnit returns the PricePerUnit field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetPricePerUnit() map[string]interface{} {
	if o == nil || IsNil(o.PricePerUnit) {
		var ret map[string]interface{}
		return ret
	}
	return o.PricePerUnit
}

// GetPricePerUnitOk returns a tuple with the PricePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetPricePerUnitOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PricePerUnit) {
		return map[string]interface{}{}, false
	}
	return o.PricePerUnit, true
}

// SetPricePerUnit gets a reference to the given map[string]interface{} and assigns it to the PricePerUnit field.
func (o *CreateProductTierRequest2) SetPricePerUnit(v map[string]interface{}) {
	o.PricePerUnit = v
}

// GetPricing returns the Pricing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProductTierRequest2) GetPricing() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProductTierRequest2) GetPricingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Pricing) {
		return nil, false
	}
	return &o.Pricing, true
}

// SetPricing gets a reference to the given interface{} and assigns it to the Pricing field.
func (o *CreateProductTierRequest2) SetPricing(v interface{}) {
	o.Pricing = v
}

// GetPrivateRegions returns the PrivateRegions field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetPrivateRegions() []string {
	if o == nil || IsNil(o.PrivateRegions) {
		var ret []string
		return ret
	}
	return o.PrivateRegions
}

// GetPrivateRegionsOk returns a tuple with the PrivateRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetPrivateRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.PrivateRegions) {
		return nil, false
	}
	return o.PrivateRegions, true
}

// SetPrivateRegions gets a reference to the given []string and assigns it to the PrivateRegions field.
func (o *CreateProductTierRequest2) SetPrivateRegions(v []string) {
	o.PrivateRegions = v
}

// GetServiceModelId returns the ServiceModelId field value
func (o *CreateProductTierRequest2) GetServiceModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelId
}

// GetServiceModelIdOk returns a tuple with the ServiceModelId field value
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetServiceModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelId, true
}

// SetServiceModelId sets field value
func (o *CreateProductTierRequest2) SetServiceModelId(v string) {
	o.ServiceModelId = v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *CreateProductTierRequest2) GetSupport() string {
	if o == nil || IsNil(o.Support) {
		var ret string
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetSupportOk() (*string, bool) {
	if o == nil || IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// SetSupport gets a reference to the given string and assigns it to the Support field.
func (o *CreateProductTierRequest2) SetSupport(v string) {
	o.Support = &v
}

// GetTierType returns the TierType field value
func (o *CreateProductTierRequest2) GetTierType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TierType
}

// GetTierTypeOk returns a tuple with the TierType field value
// and a boolean to check if the value has been set.
func (o *CreateProductTierRequest2) GetTierTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TierType, true
}

// SetTierType sets field value
func (o *CreateProductTierRequest2) SetTierType(v string) {
	o.TierType = v
}

func (o CreateProductTierRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProductTierRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowCreatesWhenPaymentNotConfigured) {
		toSerialize["allowCreatesWhenPaymentNotConfigured"] = o.AllowCreatesWhenPaymentNotConfigured
	}
	if !IsNil(o.AutoApproveSubscription) {
		toSerialize["autoApproveSubscription"] = o.AutoApproveSubscription
	}
	if !IsNil(o.AwsRegions) {
		toSerialize["awsRegions"] = o.AwsRegions
	}
	if !IsNil(o.AzureRegions) {
		toSerialize["azureRegions"] = o.AzureRegions
	}
	if !IsNil(o.BillingProductID) {
		toSerialize["billingProductID"] = o.BillingProductID
	}
	if !IsNil(o.BillingProviders) {
		toSerialize["billingProviders"] = o.BillingProviders
	}
	if !IsNil(o.DefaultBillingProvider) {
		toSerialize["defaultBillingProvider"] = o.DefaultBillingProvider
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.Documentation) {
		toSerialize["documentation"] = o.Documentation
	}
	if !IsNil(o.ExportUsageMetering) {
		toSerialize["exportUsageMetering"] = o.ExportUsageMetering
	}
	if !IsNil(o.ExportUsageMeteringConfig) {
		toSerialize["exportUsageMeteringConfig"] = o.ExportUsageMeteringConfig
	}
	if !IsNil(o.GcpRegions) {
		toSerialize["gcpRegions"] = o.GcpRegions
	}
	if !IsNil(o.IsDisabled) {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	if !IsNil(o.MaxNumberOfInstances) {
		toSerialize["maxNumberOfInstances"] = o.MaxNumberOfInstances
	}
	toSerialize["name"] = o.Name
	toSerialize["planDescription"] = o.PlanDescription
	if !IsNil(o.PricePerUnit) {
		toSerialize["pricePerUnit"] = o.PricePerUnit
	}
	if o.Pricing != nil {
		toSerialize["pricing"] = o.Pricing
	}
	if !IsNil(o.PrivateRegions) {
		toSerialize["privateRegions"] = o.PrivateRegions
	}
	toSerialize["serviceModelId"] = o.ServiceModelId
	if !IsNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	toSerialize["tierType"] = o.TierType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateProductTierRequest2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"name",
		"planDescription",
		"serviceModelId",
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

	varCreateProductTierRequest2 := _CreateProductTierRequest2{}

	err = json.Unmarshal(data, &varCreateProductTierRequest2)

	if err != nil {
		return err
	}

	*o = CreateProductTierRequest2(varCreateProductTierRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowCreatesWhenPaymentNotConfigured")
		delete(additionalProperties, "autoApproveSubscription")
		delete(additionalProperties, "awsRegions")
		delete(additionalProperties, "azureRegions")
		delete(additionalProperties, "billingProductID")
		delete(additionalProperties, "billingProviders")
		delete(additionalProperties, "defaultBillingProvider")
		delete(additionalProperties, "description")
		delete(additionalProperties, "documentation")
		delete(additionalProperties, "exportUsageMetering")
		delete(additionalProperties, "exportUsageMeteringConfig")
		delete(additionalProperties, "gcpRegions")
		delete(additionalProperties, "isDisabled")
		delete(additionalProperties, "maxNumberOfInstances")
		delete(additionalProperties, "name")
		delete(additionalProperties, "planDescription")
		delete(additionalProperties, "pricePerUnit")
		delete(additionalProperties, "pricing")
		delete(additionalProperties, "privateRegions")
		delete(additionalProperties, "serviceModelId")
		delete(additionalProperties, "support")
		delete(additionalProperties, "tierType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateProductTierRequest2 struct {
	value *CreateProductTierRequest2
	isSet bool
}

func (v NullableCreateProductTierRequest2) Get() *CreateProductTierRequest2 {
	return v.value
}

func (v *NullableCreateProductTierRequest2) Set(val *CreateProductTierRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductTierRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductTierRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductTierRequest2(val *CreateProductTierRequest2) *NullableCreateProductTierRequest2 {
	return &NullableCreateProductTierRequest2{value: val, isSet: true}
}

func (v NullableCreateProductTierRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProductTierRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


