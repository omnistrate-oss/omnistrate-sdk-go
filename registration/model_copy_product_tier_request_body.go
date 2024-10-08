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

// checks if the CopyProductTierRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyProductTierRequestBody{}

// CopyProductTierRequestBody struct for CopyProductTierRequestBody
type CopyProductTierRequestBody struct {
	// The AWS regions that this product tier is available on
	AwsRegions []string `json:"awsRegions,omitempty"`
	// A brief description of the product tier
	Description string `json:"description"`
	// Documentation
	Documentation *string `json:"documentation,omitempty"`
	// The GCP regions that this product tier is available on
	GcpRegions []string `json:"gcpRegions,omitempty"`
	// Name of the product tier
	Name string `json:"name"`
	// A brief description for the end user of the product tier
	PlanDescription *string `json:"planDescription,omitempty"`
	// Pricing
	Pricing interface{} `json:"pricing,omitempty"`
	// Service model ID
	ServiceModelId string `json:"serviceModelId"`
	// Support
	Support *string `json:"support,omitempty"`
	// Tier type
	TargetTierType *string `json:"targetTierType,omitempty"`
}

type _CopyProductTierRequestBody CopyProductTierRequestBody

// NewCopyProductTierRequestBody instantiates a new CopyProductTierRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyProductTierRequestBody(description string, name string, serviceModelId string) *CopyProductTierRequestBody {
	this := CopyProductTierRequestBody{}
	this.Description = description
	this.Name = name
	this.ServiceModelId = serviceModelId
	return &this
}

// NewCopyProductTierRequestBodyWithDefaults instantiates a new CopyProductTierRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyProductTierRequestBodyWithDefaults() *CopyProductTierRequestBody {
	this := CopyProductTierRequestBody{}
	return &this
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *CopyProductTierRequestBody) GetAwsRegions() []string {
	if o == nil || IsNil(o.AwsRegions) {
		var ret []string
		return ret
	}
	return o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyProductTierRequestBody) GetAwsRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AwsRegions) {
		return nil, false
	}
	return o.AwsRegions, true
}

// HasAwsRegions returns a boolean if a field has been set.
func (o *CopyProductTierRequestBody) HasAwsRegions() bool {
	if o != nil && !IsNil(o.AwsRegions) {
		return true
	}

	return false
}

// SetAwsRegions gets a reference to the given []string and assigns it to the AwsRegions field.
func (o *CopyProductTierRequestBody) SetAwsRegions(v []string) {
	o.AwsRegions = v
}

// GetDescription returns the Description field value
func (o *CopyProductTierRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CopyProductTierRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CopyProductTierRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *CopyProductTierRequestBody) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation) {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyProductTierRequestBody) GetDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *CopyProductTierRequestBody) HasDocumentation() bool {
	if o != nil && !IsNil(o.Documentation) {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *CopyProductTierRequestBody) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetGcpRegions returns the GcpRegions field value if set, zero value otherwise.
func (o *CopyProductTierRequestBody) GetGcpRegions() []string {
	if o == nil || IsNil(o.GcpRegions) {
		var ret []string
		return ret
	}
	return o.GcpRegions
}

// GetGcpRegionsOk returns a tuple with the GcpRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyProductTierRequestBody) GetGcpRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.GcpRegions) {
		return nil, false
	}
	return o.GcpRegions, true
}

// HasGcpRegions returns a boolean if a field has been set.
func (o *CopyProductTierRequestBody) HasGcpRegions() bool {
	if o != nil && !IsNil(o.GcpRegions) {
		return true
	}

	return false
}

// SetGcpRegions gets a reference to the given []string and assigns it to the GcpRegions field.
func (o *CopyProductTierRequestBody) SetGcpRegions(v []string) {
	o.GcpRegions = v
}

// GetName returns the Name field value
func (o *CopyProductTierRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CopyProductTierRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CopyProductTierRequestBody) SetName(v string) {
	o.Name = v
}

// GetPlanDescription returns the PlanDescription field value if set, zero value otherwise.
func (o *CopyProductTierRequestBody) GetPlanDescription() string {
	if o == nil || IsNil(o.PlanDescription) {
		var ret string
		return ret
	}
	return *o.PlanDescription
}

// GetPlanDescriptionOk returns a tuple with the PlanDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyProductTierRequestBody) GetPlanDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PlanDescription) {
		return nil, false
	}
	return o.PlanDescription, true
}

// HasPlanDescription returns a boolean if a field has been set.
func (o *CopyProductTierRequestBody) HasPlanDescription() bool {
	if o != nil && !IsNil(o.PlanDescription) {
		return true
	}

	return false
}

// SetPlanDescription gets a reference to the given string and assigns it to the PlanDescription field.
func (o *CopyProductTierRequestBody) SetPlanDescription(v string) {
	o.PlanDescription = &v
}

// GetPricing returns the Pricing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CopyProductTierRequestBody) GetPricing() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CopyProductTierRequestBody) GetPricingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Pricing) {
		return nil, false
	}
	return &o.Pricing, true
}

// HasPricing returns a boolean if a field has been set.
func (o *CopyProductTierRequestBody) HasPricing() bool {
	if o != nil && !IsNil(o.Pricing) {
		return true
	}

	return false
}

// SetPricing gets a reference to the given interface{} and assigns it to the Pricing field.
func (o *CopyProductTierRequestBody) SetPricing(v interface{}) {
	o.Pricing = v
}

// GetServiceModelId returns the ServiceModelId field value
func (o *CopyProductTierRequestBody) GetServiceModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelId
}

// GetServiceModelIdOk returns a tuple with the ServiceModelId field value
// and a boolean to check if the value has been set.
func (o *CopyProductTierRequestBody) GetServiceModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelId, true
}

// SetServiceModelId sets field value
func (o *CopyProductTierRequestBody) SetServiceModelId(v string) {
	o.ServiceModelId = v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *CopyProductTierRequestBody) GetSupport() string {
	if o == nil || IsNil(o.Support) {
		var ret string
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyProductTierRequestBody) GetSupportOk() (*string, bool) {
	if o == nil || IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *CopyProductTierRequestBody) HasSupport() bool {
	if o != nil && !IsNil(o.Support) {
		return true
	}

	return false
}

// SetSupport gets a reference to the given string and assigns it to the Support field.
func (o *CopyProductTierRequestBody) SetSupport(v string) {
	o.Support = &v
}

// GetTargetTierType returns the TargetTierType field value if set, zero value otherwise.
func (o *CopyProductTierRequestBody) GetTargetTierType() string {
	if o == nil || IsNil(o.TargetTierType) {
		var ret string
		return ret
	}
	return *o.TargetTierType
}

// GetTargetTierTypeOk returns a tuple with the TargetTierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyProductTierRequestBody) GetTargetTierTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetTierType) {
		return nil, false
	}
	return o.TargetTierType, true
}

// HasTargetTierType returns a boolean if a field has been set.
func (o *CopyProductTierRequestBody) HasTargetTierType() bool {
	if o != nil && !IsNil(o.TargetTierType) {
		return true
	}

	return false
}

// SetTargetTierType gets a reference to the given string and assigns it to the TargetTierType field.
func (o *CopyProductTierRequestBody) SetTargetTierType(v string) {
	o.TargetTierType = &v
}

func (o CopyProductTierRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyProductTierRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsRegions) {
		toSerialize["awsRegions"] = o.AwsRegions
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.Documentation) {
		toSerialize["documentation"] = o.Documentation
	}
	if !IsNil(o.GcpRegions) {
		toSerialize["gcpRegions"] = o.GcpRegions
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PlanDescription) {
		toSerialize["planDescription"] = o.PlanDescription
	}
	if o.Pricing != nil {
		toSerialize["pricing"] = o.Pricing
	}
	toSerialize["serviceModelId"] = o.ServiceModelId
	if !IsNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	if !IsNil(o.TargetTierType) {
		toSerialize["targetTierType"] = o.TargetTierType
	}
	return toSerialize, nil
}

func (o *CopyProductTierRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"name",
		"serviceModelId",
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

	varCopyProductTierRequestBody := _CopyProductTierRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCopyProductTierRequestBody)

	if err != nil {
		return err
	}

	*o = CopyProductTierRequestBody(varCopyProductTierRequestBody)

	return err
}

type NullableCopyProductTierRequestBody struct {
	value *CopyProductTierRequestBody
	isSet bool
}

func (v NullableCopyProductTierRequestBody) Get() *CopyProductTierRequestBody {
	return v.value
}

func (v *NullableCopyProductTierRequestBody) Set(val *CopyProductTierRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyProductTierRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyProductTierRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyProductTierRequestBody(val *CopyProductTierRequestBody) *NullableCopyProductTierRequestBody {
	return &NullableCopyProductTierRequestBody{value: val, isSet: true}
}

func (v NullableCopyProductTierRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyProductTierRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


