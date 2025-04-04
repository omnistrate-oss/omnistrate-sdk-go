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

// checks if the CreateServicesOrchestrationResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServicesOrchestrationResponseBody{}

// CreateServicesOrchestrationResponseBody struct for CreateServicesOrchestrationResponseBody
type CreateServicesOrchestrationResponseBody struct {
	// Services Orchestration Id
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateServicesOrchestrationResponseBody CreateServicesOrchestrationResponseBody

// NewCreateServicesOrchestrationResponseBody instantiates a new CreateServicesOrchestrationResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServicesOrchestrationResponseBody() *CreateServicesOrchestrationResponseBody {
	this := CreateServicesOrchestrationResponseBody{}
	return &this
}

// NewCreateServicesOrchestrationResponseBodyWithDefaults instantiates a new CreateServicesOrchestrationResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServicesOrchestrationResponseBodyWithDefaults() *CreateServicesOrchestrationResponseBody {
	this := CreateServicesOrchestrationResponseBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateServicesOrchestrationResponseBody) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServicesOrchestrationResponseBody) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateServicesOrchestrationResponseBody) SetId(v string) {
	o.Id = &v
}

func (o CreateServicesOrchestrationResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServicesOrchestrationResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateServicesOrchestrationResponseBody) UnmarshalJSON(data []byte) (err error) {
	varCreateServicesOrchestrationResponseBody := _CreateServicesOrchestrationResponseBody{}

	err = json.Unmarshal(data, &varCreateServicesOrchestrationResponseBody)

	if err != nil {
		return err
	}

	*o = CreateServicesOrchestrationResponseBody(varCreateServicesOrchestrationResponseBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateServicesOrchestrationResponseBody struct {
	value *CreateServicesOrchestrationResponseBody
	isSet bool
}

func (v NullableCreateServicesOrchestrationResponseBody) Get() *CreateServicesOrchestrationResponseBody {
	return v.value
}

func (v *NullableCreateServicesOrchestrationResponseBody) Set(val *CreateServicesOrchestrationResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServicesOrchestrationResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServicesOrchestrationResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServicesOrchestrationResponseBody(val *CreateServicesOrchestrationResponseBody) *NullableCreateServicesOrchestrationResponseBody {
	return &NullableCreateServicesOrchestrationResponseBody{value: val, isSet: true}
}

func (v NullableCreateServicesOrchestrationResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServicesOrchestrationResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


