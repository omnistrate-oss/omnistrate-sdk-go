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

// checks if the Route53ConfigurationDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Route53ConfigurationDescription{}

// Route53ConfigurationDescription struct for Route53ConfigurationDescription
type Route53ConfigurationDescription struct {
	// The AWS account hosting the the hosted zone for the domain
	AwsAccountID string `json:"awsAccountID"`
	// The URL to the CloudFormation template
	AwsCloudFormationTemplateURL string `json:"awsCloudFormationTemplateURL"`
}

type _Route53ConfigurationDescription Route53ConfigurationDescription

// NewRoute53ConfigurationDescription instantiates a new Route53ConfigurationDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute53ConfigurationDescription(awsAccountID string, awsCloudFormationTemplateURL string) *Route53ConfigurationDescription {
	this := Route53ConfigurationDescription{}
	this.AwsAccountID = awsAccountID
	this.AwsCloudFormationTemplateURL = awsCloudFormationTemplateURL
	return &this
}

// NewRoute53ConfigurationDescriptionWithDefaults instantiates a new Route53ConfigurationDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoute53ConfigurationDescriptionWithDefaults() *Route53ConfigurationDescription {
	this := Route53ConfigurationDescription{}
	return &this
}

// GetAwsAccountID returns the AwsAccountID field value
func (o *Route53ConfigurationDescription) GetAwsAccountID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccountID
}

// GetAwsAccountIDOk returns a tuple with the AwsAccountID field value
// and a boolean to check if the value has been set.
func (o *Route53ConfigurationDescription) GetAwsAccountIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountID, true
}

// SetAwsAccountID sets field value
func (o *Route53ConfigurationDescription) SetAwsAccountID(v string) {
	o.AwsAccountID = v
}

// GetAwsCloudFormationTemplateURL returns the AwsCloudFormationTemplateURL field value
func (o *Route53ConfigurationDescription) GetAwsCloudFormationTemplateURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsCloudFormationTemplateURL
}

// GetAwsCloudFormationTemplateURLOk returns a tuple with the AwsCloudFormationTemplateURL field value
// and a boolean to check if the value has been set.
func (o *Route53ConfigurationDescription) GetAwsCloudFormationTemplateURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsCloudFormationTemplateURL, true
}

// SetAwsCloudFormationTemplateURL sets field value
func (o *Route53ConfigurationDescription) SetAwsCloudFormationTemplateURL(v string) {
	o.AwsCloudFormationTemplateURL = v
}

func (o Route53ConfigurationDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Route53ConfigurationDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["awsAccountID"] = o.AwsAccountID
	toSerialize["awsCloudFormationTemplateURL"] = o.AwsCloudFormationTemplateURL
	return toSerialize, nil
}

func (o *Route53ConfigurationDescription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"awsAccountID",
		"awsCloudFormationTemplateURL",
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

	varRoute53ConfigurationDescription := _Route53ConfigurationDescription{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoute53ConfigurationDescription)

	if err != nil {
		return err
	}

	*o = Route53ConfigurationDescription(varRoute53ConfigurationDescription)

	return err
}

type NullableRoute53ConfigurationDescription struct {
	value *Route53ConfigurationDescription
	isSet bool
}

func (v NullableRoute53ConfigurationDescription) Get() *Route53ConfigurationDescription {
	return v.value
}

func (v *NullableRoute53ConfigurationDescription) Set(val *Route53ConfigurationDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute53ConfigurationDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute53ConfigurationDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute53ConfigurationDescription(val *Route53ConfigurationDescription) *NullableRoute53ConfigurationDescription {
	return &NullableRoute53ConfigurationDescription{value: val, isSet: true}
}

func (v NullableRoute53ConfigurationDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute53ConfigurationDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


