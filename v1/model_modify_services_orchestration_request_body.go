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

// checks if the ModifyServicesOrchestrationRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyServicesOrchestrationRequestBody{}

// ModifyServicesOrchestrationRequestBody struct for ModifyServicesOrchestrationRequestBody
type ModifyServicesOrchestrationRequestBody struct {
	// base64 encoded content of services orchestration modify DSL
	OrchestrationModifyDSL string `json:"orchestrationModifyDSL"`
	AdditionalProperties map[string]interface{}
}

type _ModifyServicesOrchestrationRequestBody ModifyServicesOrchestrationRequestBody

// NewModifyServicesOrchestrationRequestBody instantiates a new ModifyServicesOrchestrationRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyServicesOrchestrationRequestBody(orchestrationModifyDSL string) *ModifyServicesOrchestrationRequestBody {
	this := ModifyServicesOrchestrationRequestBody{}
	this.OrchestrationModifyDSL = orchestrationModifyDSL
	return &this
}

// NewModifyServicesOrchestrationRequestBodyWithDefaults instantiates a new ModifyServicesOrchestrationRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyServicesOrchestrationRequestBodyWithDefaults() *ModifyServicesOrchestrationRequestBody {
	this := ModifyServicesOrchestrationRequestBody{}
	return &this
}

// GetOrchestrationModifyDSL returns the OrchestrationModifyDSL field value
func (o *ModifyServicesOrchestrationRequestBody) GetOrchestrationModifyDSL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrchestrationModifyDSL
}

// GetOrchestrationModifyDSLOk returns a tuple with the OrchestrationModifyDSL field value
// and a boolean to check if the value has been set.
func (o *ModifyServicesOrchestrationRequestBody) GetOrchestrationModifyDSLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrchestrationModifyDSL, true
}

// SetOrchestrationModifyDSL sets field value
func (o *ModifyServicesOrchestrationRequestBody) SetOrchestrationModifyDSL(v string) {
	o.OrchestrationModifyDSL = v
}

func (o ModifyServicesOrchestrationRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyServicesOrchestrationRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orchestrationModifyDSL"] = o.OrchestrationModifyDSL

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModifyServicesOrchestrationRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orchestrationModifyDSL",
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

	varModifyServicesOrchestrationRequestBody := _ModifyServicesOrchestrationRequestBody{}

	err = json.Unmarshal(data, &varModifyServicesOrchestrationRequestBody)

	if err != nil {
		return err
	}

	*o = ModifyServicesOrchestrationRequestBody(varModifyServicesOrchestrationRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orchestrationModifyDSL")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModifyServicesOrchestrationRequestBody struct {
	value *ModifyServicesOrchestrationRequestBody
	isSet bool
}

func (v NullableModifyServicesOrchestrationRequestBody) Get() *ModifyServicesOrchestrationRequestBody {
	return v.value
}

func (v *NullableModifyServicesOrchestrationRequestBody) Set(val *ModifyServicesOrchestrationRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyServicesOrchestrationRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyServicesOrchestrationRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyServicesOrchestrationRequestBody(val *ModifyServicesOrchestrationRequestBody) *NullableModifyServicesOrchestrationRequestBody {
	return &NullableModifyServicesOrchestrationRequestBody{value: val, isSet: true}
}

func (v NullableModifyServicesOrchestrationRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyServicesOrchestrationRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

