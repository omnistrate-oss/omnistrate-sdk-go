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

// checks if the UpdateProductTierRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductTierRequestBody{}

// UpdateProductTierRequestBody struct for UpdateProductTierRequestBody
type UpdateProductTierRequestBody struct {
	// The AWS regions that this product tier is available on
	AwsRegions []string `json:"awsRegions,omitempty"`
	// A brief description of the product tier
	Description *string `json:"description,omitempty"`
	// Documentation
	Documentation *string `json:"documentation,omitempty"`
	// The GCP regions that this product tier is available on
	GcpRegions []string `json:"gcpRegions,omitempty"`
	// Update the product tier's state as enabled/disabled. Enabling the product tier will let end-customers subscribe and use the service plan.
	IsDisabled *bool `json:"isDisabled,omitempty"`
	// Name of the product tier
	Name *string `json:"name,omitempty"`
	// A brief description for the end user of the product tier
	PlanDescription *string `json:"planDescription,omitempty"`
	// Pricing
	Pricing interface{} `json:"pricing,omitempty"`
	// Support
	Support *string `json:"support,omitempty"`
	// Tier type
	TierType *string `json:"tierType,omitempty"`
}

// NewUpdateProductTierRequestBody instantiates a new UpdateProductTierRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductTierRequestBody() *UpdateProductTierRequestBody {
	this := UpdateProductTierRequestBody{}
	return &this
}

// NewUpdateProductTierRequestBodyWithDefaults instantiates a new UpdateProductTierRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductTierRequestBodyWithDefaults() *UpdateProductTierRequestBody {
	this := UpdateProductTierRequestBody{}
	return &this
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *UpdateProductTierRequestBody) GetAwsRegions() []string {
	if o == nil || IsNil(o.AwsRegions) {
		var ret []string
		return ret
	}
	return o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequestBody) GetAwsRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AwsRegions) {
		return nil, false
	}
	return o.AwsRegions, true
}

// SetAwsRegions gets a reference to the given []string and assigns it to the AwsRegions field.
func (o *UpdateProductTierRequestBody) SetAwsRegions(v []string) {
	o.AwsRegions = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateProductTierRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateProductTierRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *UpdateProductTierRequestBody) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation) {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequestBody) GetDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *UpdateProductTierRequestBody) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetGcpRegions returns the GcpRegions field value if set, zero value otherwise.
func (o *UpdateProductTierRequestBody) GetGcpRegions() []string {
	if o == nil || IsNil(o.GcpRegions) {
		var ret []string
		return ret
	}
	return o.GcpRegions
}

// GetGcpRegionsOk returns a tuple with the GcpRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequestBody) GetGcpRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.GcpRegions) {
		return nil, false
	}
	return o.GcpRegions, true
}

// SetGcpRegions gets a reference to the given []string and assigns it to the GcpRegions field.
func (o *UpdateProductTierRequestBody) SetGcpRegions(v []string) {
	o.GcpRegions = v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *UpdateProductTierRequestBody) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequestBody) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *UpdateProductTierRequestBody) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateProductTierRequestBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequestBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateProductTierRequestBody) SetName(v string) {
	o.Name = &v
}

// GetPlanDescription returns the PlanDescription field value if set, zero value otherwise.
func (o *UpdateProductTierRequestBody) GetPlanDescription() string {
	if o == nil || IsNil(o.PlanDescription) {
		var ret string
		return ret
	}
	return *o.PlanDescription
}

// GetPlanDescriptionOk returns a tuple with the PlanDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequestBody) GetPlanDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PlanDescription) {
		return nil, false
	}
	return o.PlanDescription, true
}

// SetPlanDescription gets a reference to the given string and assigns it to the PlanDescription field.
func (o *UpdateProductTierRequestBody) SetPlanDescription(v string) {
	o.PlanDescription = &v
}

// GetPricing returns the Pricing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductTierRequestBody) GetPricing() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductTierRequestBody) GetPricingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Pricing) {
		return nil, false
	}
	return &o.Pricing, true
}

// SetPricing gets a reference to the given interface{} and assigns it to the Pricing field.
func (o *UpdateProductTierRequestBody) SetPricing(v interface{}) {
	o.Pricing = v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *UpdateProductTierRequestBody) GetSupport() string {
	if o == nil || IsNil(o.Support) {
		var ret string
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequestBody) GetSupportOk() (*string, bool) {
	if o == nil || IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// SetSupport gets a reference to the given string and assigns it to the Support field.
func (o *UpdateProductTierRequestBody) SetSupport(v string) {
	o.Support = &v
}

// GetTierType returns the TierType field value if set, zero value otherwise.
func (o *UpdateProductTierRequestBody) GetTierType() string {
	if o == nil || IsNil(o.TierType) {
		var ret string
		return ret
	}
	return *o.TierType
}

// GetTierTypeOk returns a tuple with the TierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductTierRequestBody) GetTierTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TierType) {
		return nil, false
	}
	return o.TierType, true
}

// SetTierType gets a reference to the given string and assigns it to the TierType field.
func (o *UpdateProductTierRequestBody) SetTierType(v string) {
	o.TierType = &v
}

func (o UpdateProductTierRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductTierRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsRegions) {
		toSerialize["awsRegions"] = o.AwsRegions
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Documentation) {
		toSerialize["documentation"] = o.Documentation
	}
	if !IsNil(o.GcpRegions) {
		toSerialize["gcpRegions"] = o.GcpRegions
	}
	if !IsNil(o.IsDisabled) {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlanDescription) {
		toSerialize["planDescription"] = o.PlanDescription
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
	return toSerialize, nil
}

type NullableUpdateProductTierRequestBody struct {
	value *UpdateProductTierRequestBody
	isSet bool
}

func (v NullableUpdateProductTierRequestBody) Get() *UpdateProductTierRequestBody {
	return v.value
}

func (v *NullableUpdateProductTierRequestBody) Set(val *UpdateProductTierRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductTierRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductTierRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductTierRequestBody(val *UpdateProductTierRequestBody) *NullableUpdateProductTierRequestBody {
	return &NullableUpdateProductTierRequestBody{value: val, isSet: true}
}

func (v NullableUpdateProductTierRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductTierRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

