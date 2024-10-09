/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InputParameterEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputParameterEntity{}

// InputParameterEntity struct for InputParameterEntity
type InputParameterEntity struct {
	// Whether the parameter is custom
	Custom bool `json:"custom"`
	// The parameter default value
	DefaultValue *string `json:"defaultValue,omitempty"`
	// The parameter dependent resource ID
	DependentResourceID *string `json:"dependentResourceID,omitempty"`
	// The parameter description
	Description string `json:"description"`
	// The parameter display name
	DisplayName string `json:"displayName"`
	// Whether the parameter is a list
	IsList bool `json:"isList"`
	// The parameter key
	Key string `json:"key"`
	// Whether the parameter is modifiable
	Modifiable bool `json:"modifiable"`
	// The parameter options
	Options []string `json:"options,omitempty"`
	// Whether the parameter is required
	Required bool `json:"required"`
	// The parameter type
	Type string `json:"type"`
}

type _InputParameterEntity InputParameterEntity

// NewInputParameterEntity instantiates a new InputParameterEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputParameterEntity(custom bool, description string, displayName string, isList bool, key string, modifiable bool, required bool, type_ string) *InputParameterEntity {
	this := InputParameterEntity{}
	this.Custom = custom
	this.Description = description
	this.DisplayName = displayName
	this.IsList = isList
	this.Key = key
	this.Modifiable = modifiable
	this.Required = required
	this.Type = type_
	return &this
}

// NewInputParameterEntityWithDefaults instantiates a new InputParameterEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputParameterEntityWithDefaults() *InputParameterEntity {
	this := InputParameterEntity{}
	return &this
}

// GetCustom returns the Custom field value
func (o *InputParameterEntity) GetCustom() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetCustomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Custom, true
}

// SetCustom sets field value
func (o *InputParameterEntity) SetCustom(v bool) {
	o.Custom = v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *InputParameterEntity) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *InputParameterEntity) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *InputParameterEntity) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDependentResourceID returns the DependentResourceID field value if set, zero value otherwise.
func (o *InputParameterEntity) GetDependentResourceID() string {
	if o == nil || IsNil(o.DependentResourceID) {
		var ret string
		return ret
	}
	return *o.DependentResourceID
}

// GetDependentResourceIDOk returns a tuple with the DependentResourceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetDependentResourceIDOk() (*string, bool) {
	if o == nil || IsNil(o.DependentResourceID) {
		return nil, false
	}
	return o.DependentResourceID, true
}

// HasDependentResourceID returns a boolean if a field has been set.
func (o *InputParameterEntity) HasDependentResourceID() bool {
	if o != nil && !IsNil(o.DependentResourceID) {
		return true
	}

	return false
}

// SetDependentResourceID gets a reference to the given string and assigns it to the DependentResourceID field.
func (o *InputParameterEntity) SetDependentResourceID(v string) {
	o.DependentResourceID = &v
}

// GetDescription returns the Description field value
func (o *InputParameterEntity) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InputParameterEntity) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value
func (o *InputParameterEntity) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *InputParameterEntity) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetIsList returns the IsList field value
func (o *InputParameterEntity) GetIsList() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsList
}

// GetIsListOk returns a tuple with the IsList field value
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetIsListOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsList, true
}

// SetIsList sets field value
func (o *InputParameterEntity) SetIsList(v bool) {
	o.IsList = v
}

// GetKey returns the Key field value
func (o *InputParameterEntity) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *InputParameterEntity) SetKey(v string) {
	o.Key = v
}

// GetModifiable returns the Modifiable field value
func (o *InputParameterEntity) GetModifiable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Modifiable
}

// GetModifiableOk returns a tuple with the Modifiable field value
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetModifiableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modifiable, true
}

// SetModifiable sets field value
func (o *InputParameterEntity) SetModifiable(v bool) {
	o.Modifiable = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InputParameterEntity) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InputParameterEntity) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *InputParameterEntity) SetOptions(v []string) {
	o.Options = v
}

// GetRequired returns the Required field value
func (o *InputParameterEntity) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *InputParameterEntity) SetRequired(v bool) {
	o.Required = v
}

// GetType returns the Type field value
func (o *InputParameterEntity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputParameterEntity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputParameterEntity) SetType(v string) {
	o.Type = v
}

func (o InputParameterEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputParameterEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["custom"] = o.Custom
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.DependentResourceID) {
		toSerialize["dependentResourceID"] = o.DependentResourceID
	}
	toSerialize["description"] = o.Description
	toSerialize["displayName"] = o.DisplayName
	toSerialize["isList"] = o.IsList
	toSerialize["key"] = o.Key
	toSerialize["modifiable"] = o.Modifiable
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["required"] = o.Required
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *InputParameterEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"custom",
		"description",
		"displayName",
		"isList",
		"key",
		"modifiable",
		"required",
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

	varInputParameterEntity := _InputParameterEntity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputParameterEntity)

	if err != nil {
		return err
	}

	*o = InputParameterEntity(varInputParameterEntity)

	return err
}

type NullableInputParameterEntity struct {
	value *InputParameterEntity
	isSet bool
}

func (v NullableInputParameterEntity) Get() *InputParameterEntity {
	return v.value
}

func (v *NullableInputParameterEntity) Set(val *InputParameterEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableInputParameterEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableInputParameterEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputParameterEntity(val *InputParameterEntity) *NullableInputParameterEntity {
	return &NullableInputParameterEntity{value: val, isSet: true}
}

func (v NullableInputParameterEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputParameterEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


