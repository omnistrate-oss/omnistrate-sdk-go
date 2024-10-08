/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateInputParameterRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInputParameterRequestBody{}

// CreateInputParameterRequestBody struct for CreateInputParameterRequestBody
type CreateInputParameterRequestBody struct {
	// Default value to use for an optional input parameter represented as a string
	DefaultValue *string `json:"defaultValue,omitempty"`
	// The ID of the resource whose instance this input parameter depends on
	DependentResourceId *string `json:"dependentResourceId,omitempty"`
	// A brief description of the input parameter
	Description string `json:"description"`
	// Marks the input parameter to be selectable from a list of values
	HasOptions *bool `json:"hasOptions,omitempty"`
	// Marks the input parameter as a list of values
	IsList *bool `json:"isList,omitempty"`
	// Key of the input parameter
	Key string `json:"key" validate:"regexp=^[a-zA-Z][a-zA-Z0-9_]*$"`
	// A map for labeled options. The key is the label and the value is the option. When the option is selected, the label will be displayed to the end customer. Specify either options or labeledOptions when defining the input parameter.
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
	Type string `json:"type"`
}

type _CreateInputParameterRequestBody CreateInputParameterRequestBody

// NewCreateInputParameterRequestBody instantiates a new CreateInputParameterRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInputParameterRequestBody(description string, key string, modifiable bool, name string, required bool, resourceId string, type_ string) *CreateInputParameterRequestBody {
	this := CreateInputParameterRequestBody{}
	this.Description = description
	var hasOptions bool = false
	this.HasOptions = &hasOptions
	var isList bool = false
	this.IsList = &isList
	this.Key = key
	this.Modifiable = modifiable
	this.Name = name
	this.Required = required
	this.ResourceId = resourceId
	this.Type = type_
	return &this
}

// NewCreateInputParameterRequestBodyWithDefaults instantiates a new CreateInputParameterRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInputParameterRequestBodyWithDefaults() *CreateInputParameterRequestBody {
	this := CreateInputParameterRequestBody{}
	var hasOptions bool = false
	this.HasOptions = &hasOptions
	var isList bool = false
	this.IsList = &isList
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *CreateInputParameterRequestBody) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *CreateInputParameterRequestBody) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *CreateInputParameterRequestBody) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDependentResourceId returns the DependentResourceId field value if set, zero value otherwise.
func (o *CreateInputParameterRequestBody) GetDependentResourceId() string {
	if o == nil || IsNil(o.DependentResourceId) {
		var ret string
		return ret
	}
	return *o.DependentResourceId
}

// GetDependentResourceIdOk returns a tuple with the DependentResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetDependentResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DependentResourceId) {
		return nil, false
	}
	return o.DependentResourceId, true
}

// HasDependentResourceId returns a boolean if a field has been set.
func (o *CreateInputParameterRequestBody) HasDependentResourceId() bool {
	if o != nil && !IsNil(o.DependentResourceId) {
		return true
	}

	return false
}

// SetDependentResourceId gets a reference to the given string and assigns it to the DependentResourceId field.
func (o *CreateInputParameterRequestBody) SetDependentResourceId(v string) {
	o.DependentResourceId = &v
}

// GetDescription returns the Description field value
func (o *CreateInputParameterRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateInputParameterRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetHasOptions returns the HasOptions field value if set, zero value otherwise.
func (o *CreateInputParameterRequestBody) GetHasOptions() bool {
	if o == nil || IsNil(o.HasOptions) {
		var ret bool
		return ret
	}
	return *o.HasOptions
}

// GetHasOptionsOk returns a tuple with the HasOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetHasOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasOptions) {
		return nil, false
	}
	return o.HasOptions, true
}

// HasHasOptions returns a boolean if a field has been set.
func (o *CreateInputParameterRequestBody) HasHasOptions() bool {
	if o != nil && !IsNil(o.HasOptions) {
		return true
	}

	return false
}

// SetHasOptions gets a reference to the given bool and assigns it to the HasOptions field.
func (o *CreateInputParameterRequestBody) SetHasOptions(v bool) {
	o.HasOptions = &v
}

// GetIsList returns the IsList field value if set, zero value otherwise.
func (o *CreateInputParameterRequestBody) GetIsList() bool {
	if o == nil || IsNil(o.IsList) {
		var ret bool
		return ret
	}
	return *o.IsList
}

// GetIsListOk returns a tuple with the IsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetIsListOk() (*bool, bool) {
	if o == nil || IsNil(o.IsList) {
		return nil, false
	}
	return o.IsList, true
}

// HasIsList returns a boolean if a field has been set.
func (o *CreateInputParameterRequestBody) HasIsList() bool {
	if o != nil && !IsNil(o.IsList) {
		return true
	}

	return false
}

// SetIsList gets a reference to the given bool and assigns it to the IsList field.
func (o *CreateInputParameterRequestBody) SetIsList(v bool) {
	o.IsList = &v
}

// GetKey returns the Key field value
func (o *CreateInputParameterRequestBody) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateInputParameterRequestBody) SetKey(v string) {
	o.Key = v
}

// GetLabeledOptions returns the LabeledOptions field value if set, zero value otherwise.
func (o *CreateInputParameterRequestBody) GetLabeledOptions() map[string]string {
	if o == nil || IsNil(o.LabeledOptions) {
		var ret map[string]string
		return ret
	}
	return *o.LabeledOptions
}

// GetLabeledOptionsOk returns a tuple with the LabeledOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetLabeledOptionsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.LabeledOptions) {
		return nil, false
	}
	return o.LabeledOptions, true
}

// HasLabeledOptions returns a boolean if a field has been set.
func (o *CreateInputParameterRequestBody) HasLabeledOptions() bool {
	if o != nil && !IsNil(o.LabeledOptions) {
		return true
	}

	return false
}

// SetLabeledOptions gets a reference to the given map[string]string and assigns it to the LabeledOptions field.
func (o *CreateInputParameterRequestBody) SetLabeledOptions(v map[string]string) {
	o.LabeledOptions = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *CreateInputParameterRequestBody) GetLimits() Limits {
	if o == nil || IsNil(o.Limits) {
		var ret Limits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetLimitsOk() (*Limits, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *CreateInputParameterRequestBody) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given Limits and assigns it to the Limits field.
func (o *CreateInputParameterRequestBody) SetLimits(v Limits) {
	o.Limits = &v
}

// GetModifiable returns the Modifiable field value
func (o *CreateInputParameterRequestBody) GetModifiable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Modifiable
}

// GetModifiableOk returns a tuple with the Modifiable field value
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetModifiableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modifiable, true
}

// SetModifiable sets field value
func (o *CreateInputParameterRequestBody) SetModifiable(v bool) {
	o.Modifiable = v
}

// GetName returns the Name field value
func (o *CreateInputParameterRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateInputParameterRequestBody) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreateInputParameterRequestBody) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreateInputParameterRequestBody) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *CreateInputParameterRequestBody) SetOptions(v []string) {
	o.Options = v
}

// GetRequired returns the Required field value
func (o *CreateInputParameterRequestBody) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *CreateInputParameterRequestBody) SetRequired(v bool) {
	o.Required = v
}

// GetResourceId returns the ResourceId field value
func (o *CreateInputParameterRequestBody) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *CreateInputParameterRequestBody) SetResourceId(v string) {
	o.ResourceId = v
}

// GetType returns the Type field value
func (o *CreateInputParameterRequestBody) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateInputParameterRequestBody) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateInputParameterRequestBody) SetType(v string) {
	o.Type = v
}

func (o CreateInputParameterRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInputParameterRequestBody) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsList) {
		toSerialize["isList"] = o.IsList
	}
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
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *CreateInputParameterRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"key",
		"modifiable",
		"name",
		"required",
		"resourceId",
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

	varCreateInputParameterRequestBody := _CreateInputParameterRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateInputParameterRequestBody)

	if err != nil {
		return err
	}

	*o = CreateInputParameterRequestBody(varCreateInputParameterRequestBody)

	return err
}

type NullableCreateInputParameterRequestBody struct {
	value *CreateInputParameterRequestBody
	isSet bool
}

func (v NullableCreateInputParameterRequestBody) Get() *CreateInputParameterRequestBody {
	return v.value
}

func (v *NullableCreateInputParameterRequestBody) Set(val *CreateInputParameterRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInputParameterRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInputParameterRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInputParameterRequestBody(val *CreateInputParameterRequestBody) *NullableCreateInputParameterRequestBody {
	return &NullableCreateInputParameterRequestBody{value: val, isSet: true}
}

func (v NullableCreateInputParameterRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInputParameterRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


