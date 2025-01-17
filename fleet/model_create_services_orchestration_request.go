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

// checks if the CreateServicesOrchestrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServicesOrchestrationRequest{}

// CreateServicesOrchestrationRequest struct for CreateServicesOrchestrationRequest
type CreateServicesOrchestrationRequest struct {
	// base64 encoded content of service orchestration create DSL
	OrchestrationCreateDSL string `json:"orchestrationCreateDSL"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _CreateServicesOrchestrationRequest CreateServicesOrchestrationRequest

// NewCreateServicesOrchestrationRequest instantiates a new CreateServicesOrchestrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServicesOrchestrationRequest(orchestrationCreateDSL string, token string) *CreateServicesOrchestrationRequest {
	this := CreateServicesOrchestrationRequest{}
	this.OrchestrationCreateDSL = orchestrationCreateDSL
	this.Token = token
	return &this
}

// NewCreateServicesOrchestrationRequestWithDefaults instantiates a new CreateServicesOrchestrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServicesOrchestrationRequestWithDefaults() *CreateServicesOrchestrationRequest {
	this := CreateServicesOrchestrationRequest{}
	return &this
}

// GetOrchestrationCreateDSL returns the OrchestrationCreateDSL field value
func (o *CreateServicesOrchestrationRequest) GetOrchestrationCreateDSL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrchestrationCreateDSL
}

// GetOrchestrationCreateDSLOk returns a tuple with the OrchestrationCreateDSL field value
// and a boolean to check if the value has been set.
func (o *CreateServicesOrchestrationRequest) GetOrchestrationCreateDSLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrchestrationCreateDSL, true
}

// SetOrchestrationCreateDSL sets field value
func (o *CreateServicesOrchestrationRequest) SetOrchestrationCreateDSL(v string) {
	o.OrchestrationCreateDSL = v
}

// GetToken returns the Token field value
func (o *CreateServicesOrchestrationRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CreateServicesOrchestrationRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CreateServicesOrchestrationRequest) SetToken(v string) {
	o.Token = v
}

func (o CreateServicesOrchestrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServicesOrchestrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orchestrationCreateDSL"] = o.OrchestrationCreateDSL
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateServicesOrchestrationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orchestrationCreateDSL",
		"token",
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

	varCreateServicesOrchestrationRequest := _CreateServicesOrchestrationRequest{}

	err = json.Unmarshal(data, &varCreateServicesOrchestrationRequest)

	if err != nil {
		return err
	}

	*o = CreateServicesOrchestrationRequest(varCreateServicesOrchestrationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orchestrationCreateDSL")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateServicesOrchestrationRequest struct {
	value *CreateServicesOrchestrationRequest
	isSet bool
}

func (v NullableCreateServicesOrchestrationRequest) Get() *CreateServicesOrchestrationRequest {
	return v.value
}

func (v *NullableCreateServicesOrchestrationRequest) Set(val *CreateServicesOrchestrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServicesOrchestrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServicesOrchestrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServicesOrchestrationRequest(val *CreateServicesOrchestrationRequest) *NullableCreateServicesOrchestrationRequest {
	return &NullableCreateServicesOrchestrationRequest{value: val, isSet: true}
}

func (v NullableCreateServicesOrchestrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServicesOrchestrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


