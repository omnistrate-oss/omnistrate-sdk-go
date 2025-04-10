/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the UpdateProductTierRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductTierRequest2{}

// UpdateProductTierRequest2 struct for UpdateProductTierRequest2
type UpdateProductTierRequest2 struct {
	// Allow creates when payment not configured
	AllowCreatesWhenPaymentNotConfigured *bool `json:"allowCreatesWhenPaymentNotConfigured,omitempty"`
	// Auto approve subscription or not
	AutoApproveSubscription *bool `json:"autoApproveSubscription,omitempty"`
	// The AWS regions that this product tier is available on
	AwsRegions []string `json:"awsRegions,omitempty"`
	// The Azure regions that this product tier is available on
	AzureRegions []string `json:"azureRegions,omitempty"`
	// A brief description of the product tier
	Description *string `json:"description,omitempty"`
	// Documentation
	Documentation *string `json:"documentation,omitempty"`
	// Export usage metering data
	ExportUsageMetering *bool `json:"exportUsageMetering,omitempty"`
	// Export usage metering data configuration
	ExportUsageMeteringConfig map[string]interface{} `json:"exportUsageMeteringConfig,omitempty"`
	// The GCP regions that this product tier is available on
	GcpRegions []string `json:"gcpRegions,omitempty"`
	// Update the product tier's state as enabled/disabled. Enabling the product tier will let end-customers subscribe and use the service plan.
	IsDisabled *bool `json:"isDisabled,omitempty"`
	// Maximum number of instances. Set to -1 for unlimited.
	MaxNumberOfInstances *int64 `json:"maxNumberOfInstances,omitempty"`
	// Name of the product tier
	Name *string `json:"name,omitempty"`
	// A brief description for the end user of the product tier
	PlanDescription *string `json:"planDescription,omitempty"`
	// Price per unit.
	PricePerUnit map[string]interface{} `json:"pricePerUnit,omitempty"`
	// Pricing
	Pricing interface{} `json:"pricing,omitempty"`
	// Support
	Support *string `json:"support,omitempty"`
	// Tier type
	TierType *string `json:"tierType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateProductTierRequest2 UpdateProductTierRequest2

// NewUpdateProductTierRequest2 instantiates a new UpdateProductTierRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductTierRequest2() *UpdateProductTierRequest2 {
	this := UpdateProductTierRequest2{}
	return &this
}

// NewUpdateProductTierRequest2WithDefaults instantiates a new UpdateProductTierRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductTierRequest2WithDefaults() *UpdateProductTierRequest2 {
	this := UpdateProductTierRequest2{}
	return &this
}

// GetAllowCreatesWhenPaymentNotConfigured returns the AllowCreatesWhenPaymentNotConfigured field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetAllowCreatesWhenPaymentNotConfigured() bool {
	if o == nil || IsNil(o.AllowCreatesWhenPaymentNotConfigured) {
		var ret bool
		return ret
	}
	return *o.AllowCreatesWhenPaymentNotConfigured
}

// GetAllowCreatesWhenPaymentNotConfiguredOk returns a tuple with the AllowCreatesWhenPaymentNotConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetAllowCreatesWhenPaymentNotConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCreatesWhenPaymentNotConfigured) {
		return nil, false
	}
	return o.AllowCreatesWhenPaymentNotConfigured, true
}

// SetAllowCreatesWhenPaymentNotConfigured gets a reference to the given bool and assigns it to the AllowCreatesWhenPaymentNotConfigured field.
func (o *UpdateProductTierRequest2) SetAllowCreatesWhenPaymentNotConfigured(v bool) {
	o.AllowCreatesWhenPaymentNotConfigured = &v
}

// GetAutoApproveSubscription returns the AutoApproveSubscription field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetAutoApproveSubscription() bool {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		var ret bool
		return ret
	}
	return *o.AutoApproveSubscription
}

// GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetAutoApproveSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		return nil, false
	}
	return o.AutoApproveSubscription, true
}

// SetAutoApproveSubscription gets a reference to the given bool and assigns it to the AutoApproveSubscription field.
func (o *UpdateProductTierRequest2) SetAutoApproveSubscription(v bool) {
	o.AutoApproveSubscription = &v
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetAwsRegions() []string {
	if o == nil || IsNil(o.AwsRegions) {
		var ret []string
		return ret
	}
	return o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetAwsRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AwsRegions) {
		return nil, false
	}
	return o.AwsRegions, true
}

// SetAwsRegions gets a reference to the given []string and assigns it to the AwsRegions field.
func (o *UpdateProductTierRequest2) SetAwsRegions(v []string) {
	o.AwsRegions = v
}

// GetAzureRegions returns the AzureRegions field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetAzureRegions() []string {
	if o == nil || IsNil(o.AzureRegions) {
		var ret []string
		return ret
	}
	return o.AzureRegions
}

// GetAzureRegionsOk returns a tuple with the AzureRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetAzureRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AzureRegions) {
		return nil, false
	}
	return o.AzureRegions, true
}

// SetAzureRegions gets a reference to the given []string and assigns it to the AzureRegions field.
func (o *UpdateProductTierRequest2) SetAzureRegions(v []string) {
	o.AzureRegions = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateProductTierRequest2) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation) {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *UpdateProductTierRequest2) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetExportUsageMetering returns the ExportUsageMetering field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetExportUsageMetering() bool {
	if o == nil || IsNil(o.ExportUsageMetering) {
		var ret bool
		return ret
	}
	return *o.ExportUsageMetering
}

// GetExportUsageMeteringOk returns a tuple with the ExportUsageMetering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetExportUsageMeteringOk() (*bool, bool) {
	if o == nil || IsNil(o.ExportUsageMetering) {
		return nil, false
	}
	return o.ExportUsageMetering, true
}

// SetExportUsageMetering gets a reference to the given bool and assigns it to the ExportUsageMetering field.
func (o *UpdateProductTierRequest2) SetExportUsageMetering(v bool) {
	o.ExportUsageMetering = &v
}

// GetExportUsageMeteringConfig returns the ExportUsageMeteringConfig field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetExportUsageMeteringConfig() map[string]interface{} {
	if o == nil || IsNil(o.ExportUsageMeteringConfig) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExportUsageMeteringConfig
}

// GetExportUsageMeteringConfigOk returns a tuple with the ExportUsageMeteringConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetExportUsageMeteringConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExportUsageMeteringConfig) {
		return map[string]interface{}{}, false
	}
	return o.ExportUsageMeteringConfig, true
}

// SetExportUsageMeteringConfig gets a reference to the given map[string]interface{} and assigns it to the ExportUsageMeteringConfig field.
func (o *UpdateProductTierRequest2) SetExportUsageMeteringConfig(v map[string]interface{}) {
	o.ExportUsageMeteringConfig = v
}

// GetGcpRegions returns the GcpRegions field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetGcpRegions() []string {
	if o == nil || IsNil(o.GcpRegions) {
		var ret []string
		return ret
	}
	return o.GcpRegions
}

// GetGcpRegionsOk returns a tuple with the GcpRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetGcpRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.GcpRegions) {
		return nil, false
	}
	return o.GcpRegions, true
}

// SetGcpRegions gets a reference to the given []string and assigns it to the GcpRegions field.
func (o *UpdateProductTierRequest2) SetGcpRegions(v []string) {
	o.GcpRegions = v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *UpdateProductTierRequest2) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetMaxNumberOfInstances returns the MaxNumberOfInstances field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetMaxNumberOfInstances() int64 {
	if o == nil || IsNil(o.MaxNumberOfInstances) {
		var ret int64
		return ret
	}
	return *o.MaxNumberOfInstances
}

// GetMaxNumberOfInstancesOk returns a tuple with the MaxNumberOfInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetMaxNumberOfInstancesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxNumberOfInstances) {
		return nil, false
	}
	return o.MaxNumberOfInstances, true
}

// SetMaxNumberOfInstances gets a reference to the given int64 and assigns it to the MaxNumberOfInstances field.
func (o *UpdateProductTierRequest2) SetMaxNumberOfInstances(v int64) {
	o.MaxNumberOfInstances = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateProductTierRequest2) SetName(v string) {
	o.Name = &v
}

// GetPlanDescription returns the PlanDescription field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetPlanDescription() string {
	if o == nil || IsNil(o.PlanDescription) {
		var ret string
		return ret
	}
	return *o.PlanDescription
}

// GetPlanDescriptionOk returns a tuple with the PlanDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetPlanDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PlanDescription) {
		return nil, false
	}
	return o.PlanDescription, true
}

// SetPlanDescription gets a reference to the given string and assigns it to the PlanDescription field.
func (o *UpdateProductTierRequest2) SetPlanDescription(v string) {
	o.PlanDescription = &v
}

// GetPricePerUnit returns the PricePerUnit field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetPricePerUnit() map[string]interface{} {
	if o == nil || IsNil(o.PricePerUnit) {
		var ret map[string]interface{}
		return ret
	}
	return o.PricePerUnit
}

// GetPricePerUnitOk returns a tuple with the PricePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetPricePerUnitOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PricePerUnit) {
		return map[string]interface{}{}, false
	}
	return o.PricePerUnit, true
}

// SetPricePerUnit gets a reference to the given map[string]interface{} and assigns it to the PricePerUnit field.
func (o *UpdateProductTierRequest2) SetPricePerUnit(v map[string]interface{}) {
	o.PricePerUnit = v
}

// GetPricing returns the Pricing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductTierRequest2) GetPricing() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductTierRequest2) GetPricingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Pricing) {
		return nil, false
	}
	return &o.Pricing, true
}

// SetPricing gets a reference to the given interface{} and assigns it to the Pricing field.
func (o *UpdateProductTierRequest2) SetPricing(v interface{}) {
	o.Pricing = v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetSupport() string {
	if o == nil || IsNil(o.Support) {
		var ret string
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetSupportOk() (*string, bool) {
	if o == nil || IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// SetSupport gets a reference to the given string and assigns it to the Support field.
func (o *UpdateProductTierRequest2) SetSupport(v string) {
	o.Support = &v
}

// GetTierType returns the TierType field value if set, zero value otherwise.
func (o *UpdateProductTierRequest2) GetTierType() string {
	if o == nil || IsNil(o.TierType) {
		var ret string
		return ret
	}
	return *o.TierType
}

// GetTierTypeOk returns a tuple with the TierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequest2) GetTierTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TierType) {
		return nil, false
	}
	return o.TierType, true
}

// SetTierType gets a reference to the given string and assigns it to the TierType field.
func (o *UpdateProductTierRequest2) SetTierType(v string) {
	o.TierType = &v
}

func (o UpdateProductTierRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductTierRequest2) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlanDescription) {
		toSerialize["planDescription"] = o.PlanDescription
	}
	if !IsNil(o.PricePerUnit) {
		toSerialize["pricePerUnit"] = o.PricePerUnit
	}
	if o.Pricing != nil {
		toSerialize["pricing"] = o.Pricing
	}
	if !IsNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	if !IsNil(o.TierType) {
		toSerialize["tierType"] = o.TierType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateProductTierRequest2) UnmarshalJSON(data []byte) (err error) {
	varUpdateProductTierRequest2 := _UpdateProductTierRequest2{}

	err = json.Unmarshal(data, &varUpdateProductTierRequest2)

	if err != nil {
		return err
	}

	*o = UpdateProductTierRequest2(varUpdateProductTierRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowCreatesWhenPaymentNotConfigured")
		delete(additionalProperties, "autoApproveSubscription")
		delete(additionalProperties, "awsRegions")
		delete(additionalProperties, "azureRegions")
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
		delete(additionalProperties, "support")
		delete(additionalProperties, "tierType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateProductTierRequest2 struct {
	value *UpdateProductTierRequest2
	isSet bool
}

func (v NullableUpdateProductTierRequest2) Get() *UpdateProductTierRequest2 {
	return v.value
}

func (v *NullableUpdateProductTierRequest2) Set(val *UpdateProductTierRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductTierRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductTierRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductTierRequest2(val *UpdateProductTierRequest2) *NullableUpdateProductTierRequest2 {
	return &NullableUpdateProductTierRequest2{value: val, isSet: true}
}

func (v NullableUpdateProductTierRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductTierRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


