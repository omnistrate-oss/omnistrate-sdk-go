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

// checks if the Route53Configuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Route53Configuration{}

// Route53Configuration struct for Route53Configuration
type Route53Configuration struct {
	// The AWS account hosting the the hosted zone for the domain
	AwsAccountID string `json:"awsAccountID"`
	AdditionalProperties map[string]interface{}
}

type _Route53Configuration Route53Configuration

// NewRoute53Configuration instantiates a new Route53Configuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute53Configuration(awsAccountID string) *Route53Configuration {
	this := Route53Configuration{}
	this.AwsAccountID = awsAccountID
	return &this
}

// NewRoute53ConfigurationWithDefaults instantiates a new Route53Configuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoute53ConfigurationWithDefaults() *Route53Configuration {
	this := Route53Configuration{}
	return &this
}

// GetAwsAccountID returns the AwsAccountID field value
func (o *Route53Configuration) GetAwsAccountID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccountID
}

// GetAwsAccountIDOk returns a tuple with the AwsAccountID field value
// and a boolean to check if the value has been set.
func (o *Route53Configuration) GetAwsAccountIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountID, true
}

// SetAwsAccountID sets field value
func (o *Route53Configuration) SetAwsAccountID(v string) {
	o.AwsAccountID = v
}

func (o Route53Configuration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Route53Configuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["awsAccountID"] = o.AwsAccountID

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Route53Configuration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"awsAccountID",
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

	varRoute53Configuration := _Route53Configuration{}

	err = json.Unmarshal(data, &varRoute53Configuration)

	if err != nil {
		return err
	}

	*o = Route53Configuration(varRoute53Configuration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "awsAccountID")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoute53Configuration struct {
	value *Route53Configuration
	isSet bool
}

func (v NullableRoute53Configuration) Get() *Route53Configuration {
	return v.value
}

func (v *NullableRoute53Configuration) Set(val *Route53Configuration) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute53Configuration) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute53Configuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute53Configuration(val *Route53Configuration) *NullableRoute53Configuration {
	return &NullableRoute53Configuration{value: val, isSet: true}
}

func (v NullableRoute53Configuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute53Configuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


