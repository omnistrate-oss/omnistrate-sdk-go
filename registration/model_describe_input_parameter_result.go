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

// checks if the DescribeInputParameterResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeInputParameterResult{}

// DescribeInputParameterResult struct for DescribeInputParameterResult
type DescribeInputParameterResult struct {
	// Default value to use for an optional input parameter represented as a string
	DefaultValue *string `json:"defaultValue,omitempty"`
	// The ID of the resource whose instance this input parameter depends on
	DependentResourceId *string `json:"dependentResourceId,omitempty"`
	// A brief description of the input parameter
	Description string `json:"description"`
	// Marks the input parameter to be selectable from a list of values
	HasOptions *bool `json:"hasOptions,omitempty"`
	// ID of the input parameter
	Id string `json:"id"`
	// Marks the input parameter as a list of values
	IsList bool `json:"isList"`
	// Key of the input parameter
	Key string `json:"key"`
	// A map for labeled options. The key is the label and the value is the option. When the option is selected, the label will be displayed to the end customer.
	LabeledOptions *map[string]string `json:"labeledOptions,omitempty"`
	Limits *Limits `json:"limits,omitempty"`
	// Marks the input parameter as immutable
	Modifiable bool `json:"modifiable"`
	// External name for the input parameter
	Name string `json:"name"`
	// A list of options to restrict the value of the input parameter to (represented as a string)
	Options []string `json:"options,omitempty"`
	// Enforces the input parameter as a required parameter
	Required bool `json:"required"`
	// The ID of the resource that this input parameter belongs to
	ResourceId string `json:"resourceId"`
	// The ID of the service that this output parameter belongs to
	ServiceId string `json:"serviceId"`
	Type string `json:"type"`
}

type _DescribeInputParameterResult DescribeInputParameterResult

// NewDescribeInputParameterResult instantiates a new DescribeInputParameterResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeInputParameterResult(description string, id string, isList bool, key string, modifiable bool, name string, required bool, resourceId string, serviceId string, type_ string) *DescribeInputParameterResult {
	this := DescribeInputParameterResult{}
	this.Description = description
	var hasOptions bool = false
	this.HasOptions = &hasOptions
	this.Id = id
	this.IsList = isList
	this.Key = key
	this.Modifiable = modifiable
	this.Name = name
	this.Required = required
	this.ResourceId = resourceId
	this.ServiceId = serviceId
	this.Type = type_
	return &this
}

// NewDescribeInputParameterResultWithDefaults instantiates a new DescribeInputParameterResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeInputParameterResultWithDefaults() *DescribeInputParameterResult {
	this := DescribeInputParameterResult{}
	var hasOptions bool = false
	this.HasOptions = &hasOptions
	var isList bool = false
	this.IsList = isList
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *DescribeInputParameterResult) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *DescribeInputParameterResult) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *DescribeInputParameterResult) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDependentResourceId returns the DependentResourceId field value if set, zero value otherwise.
func (o *DescribeInputParameterResult) GetDependentResourceId() string {
	if o == nil || IsNil(o.DependentResourceId) {
		var ret string
		return ret
	}
	return *o.DependentResourceId
}

// GetDependentResourceIdOk returns a tuple with the DependentResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetDependentResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DependentResourceId) {
		return nil, false
	}
	return o.DependentResourceId, true
}

// HasDependentResourceId returns a boolean if a field has been set.
func (o *DescribeInputParameterResult) HasDependentResourceId() bool {
	if o != nil && !IsNil(o.DependentResourceId) {
		return true
	}

	return false
}

// SetDependentResourceId gets a reference to the given string and assigns it to the DependentResourceId field.
func (o *DescribeInputParameterResult) SetDependentResourceId(v string) {
	o.DependentResourceId = &v
}

// GetDescription returns the Description field value
func (o *DescribeInputParameterResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeInputParameterResult) SetDescription(v string) {
	o.Description = v
}

// GetHasOptions returns the HasOptions field value if set, zero value otherwise.
func (o *DescribeInputParameterResult) GetHasOptions() bool {
	if o == nil || IsNil(o.HasOptions) {
		var ret bool
		return ret
	}
	return *o.HasOptions
}

// GetHasOptionsOk returns a tuple with the HasOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetHasOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasOptions) {
		return nil, false
	}
	return o.HasOptions, true
}

// HasHasOptions returns a boolean if a field has been set.
func (o *DescribeInputParameterResult) HasHasOptions() bool {
	if o != nil && !IsNil(o.HasOptions) {
		return true
	}

	return false
}

// SetHasOptions gets a reference to the given bool and assigns it to the HasOptions field.
func (o *DescribeInputParameterResult) SetHasOptions(v bool) {
	o.HasOptions = &v
}

// GetId returns the Id field value
func (o *DescribeInputParameterResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeInputParameterResult) SetId(v string) {
	o.Id = v
}

// GetIsList returns the IsList field value
func (o *DescribeInputParameterResult) GetIsList() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsList
}

// GetIsListOk returns a tuple with the IsList field value
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetIsListOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsList, true
}

// SetIsList sets field value
func (o *DescribeInputParameterResult) SetIsList(v bool) {
	o.IsList = v
}

// GetKey returns the Key field value
func (o *DescribeInputParameterResult) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DescribeInputParameterResult) SetKey(v string) {
	o.Key = v
}

// GetLabeledOptions returns the LabeledOptions field value if set, zero value otherwise.
func (o *DescribeInputParameterResult) GetLabeledOptions() map[string]string {
	if o == nil || IsNil(o.LabeledOptions) {
		var ret map[string]string
		return ret
	}
	return *o.LabeledOptions
}

// GetLabeledOptionsOk returns a tuple with the LabeledOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetLabeledOptionsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.LabeledOptions) {
		return nil, false
	}
	return o.LabeledOptions, true
}

// HasLabeledOptions returns a boolean if a field has been set.
func (o *DescribeInputParameterResult) HasLabeledOptions() bool {
	if o != nil && !IsNil(o.LabeledOptions) {
		return true
	}

	return false
}

// SetLabeledOptions gets a reference to the given map[string]string and assigns it to the LabeledOptions field.
func (o *DescribeInputParameterResult) SetLabeledOptions(v map[string]string) {
	o.LabeledOptions = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *DescribeInputParameterResult) GetLimits() Limits {
	if o == nil || IsNil(o.Limits) {
		var ret Limits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetLimitsOk() (*Limits, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *DescribeInputParameterResult) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given Limits and assigns it to the Limits field.
func (o *DescribeInputParameterResult) SetLimits(v Limits) {
	o.Limits = &v
}

// GetModifiable returns the Modifiable field value
func (o *DescribeInputParameterResult) GetModifiable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Modifiable
}

// GetModifiableOk returns a tuple with the Modifiable field value
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetModifiableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modifiable, true
}

// SetModifiable sets field value
func (o *DescribeInputParameterResult) SetModifiable(v bool) {
	o.Modifiable = v
}

// GetName returns the Name field value
func (o *DescribeInputParameterResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeInputParameterResult) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DescribeInputParameterResult) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DescribeInputParameterResult) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *DescribeInputParameterResult) SetOptions(v []string) {
	o.Options = v
}

// GetRequired returns the Required field value
func (o *DescribeInputParameterResult) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *DescribeInputParameterResult) SetRequired(v bool) {
	o.Required = v
}

// GetResourceId returns the ResourceId field value
func (o *DescribeInputParameterResult) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *DescribeInputParameterResult) SetResourceId(v string) {
	o.ResourceId = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeInputParameterResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeInputParameterResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetType returns the Type field value
func (o *DescribeInputParameterResult) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DescribeInputParameterResult) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DescribeInputParameterResult) SetType(v string) {
	o.Type = v
}

func (o DescribeInputParameterResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeInputParameterResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.DependentResourceId) {
		toSerialize["dependentResourceId"] = o.DependentResourceId
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.HasOptions) {
		toSerialize["hasOptions"] = o.HasOptions
	}
	toSerialize["id"] = o.Id
	toSerialize["isList"] = o.IsList
	toSerialize["key"] = o.Key
	if !IsNil(o.LabeledOptions) {
		toSerialize["labeledOptions"] = o.LabeledOptions
	}
	if !IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	toSerialize["modifiable"] = o.Modifiable
	toSerialize["name"] = o.Name
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["required"] = o.Required
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *DescribeInputParameterResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"isList",
		"key",
		"modifiable",
		"name",
		"required",
		"resourceId",
		"serviceId",
		"type",
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

	varDescribeInputParameterResult := _DescribeInputParameterResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeInputParameterResult)

	if err != nil {
		return err
	}

	*o = DescribeInputParameterResult(varDescribeInputParameterResult)

	return err
}

type NullableDescribeInputParameterResult struct {
	value *DescribeInputParameterResult
	isSet bool
}

func (v NullableDescribeInputParameterResult) Get() *DescribeInputParameterResult {
	return v.value
}

func (v *NullableDescribeInputParameterResult) Set(val *DescribeInputParameterResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeInputParameterResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeInputParameterResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeInputParameterResult(val *DescribeInputParameterResult) *NullableDescribeInputParameterResult {
	return &NullableDescribeInputParameterResult{value: val, isSet: true}
}

func (v NullableDescribeInputParameterResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeInputParameterResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


