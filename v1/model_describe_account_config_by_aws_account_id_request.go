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

// checks if the DescribeAccountConfigByAWSAccountIDRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeAccountConfigByAWSAccountIDRequest{}

// DescribeAccountConfigByAWSAccountIDRequest struct for DescribeAccountConfigByAWSAccountIDRequest
type DescribeAccountConfigByAWSAccountIDRequest struct {
	// The AWS account ID
	AwsAccountID string `json:"awsAccountID"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DescribeAccountConfigByAWSAccountIDRequest DescribeAccountConfigByAWSAccountIDRequest

// NewDescribeAccountConfigByAWSAccountIDRequest instantiates a new DescribeAccountConfigByAWSAccountIDRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeAccountConfigByAWSAccountIDRequest(awsAccountID string, token string) *DescribeAccountConfigByAWSAccountIDRequest {
	this := DescribeAccountConfigByAWSAccountIDRequest{}
	this.AwsAccountID = awsAccountID
	this.Token = token
	return &this
}

// NewDescribeAccountConfigByAWSAccountIDRequestWithDefaults instantiates a new DescribeAccountConfigByAWSAccountIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeAccountConfigByAWSAccountIDRequestWithDefaults() *DescribeAccountConfigByAWSAccountIDRequest {
	this := DescribeAccountConfigByAWSAccountIDRequest{}
	return &this
}

// GetAwsAccountID returns the AwsAccountID field value
func (o *DescribeAccountConfigByAWSAccountIDRequest) GetAwsAccountID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccountID
}

// GetAwsAccountIDOk returns a tuple with the AwsAccountID field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDRequest) GetAwsAccountIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountID, true
}

// SetAwsAccountID sets field value
func (o *DescribeAccountConfigByAWSAccountIDRequest) SetAwsAccountID(v string) {
	o.AwsAccountID = v
}

// GetToken returns the Token field value
func (o *DescribeAccountConfigByAWSAccountIDRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DescribeAccountConfigByAWSAccountIDRequest) SetToken(v string) {
	o.Token = v
}

func (o DescribeAccountConfigByAWSAccountIDRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeAccountConfigByAWSAccountIDRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["awsAccountID"] = o.AwsAccountID
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeAccountConfigByAWSAccountIDRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"awsAccountID",
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

	varDescribeAccountConfigByAWSAccountIDRequest := _DescribeAccountConfigByAWSAccountIDRequest{}

	err = json.Unmarshal(data, &varDescribeAccountConfigByAWSAccountIDRequest)

	if err != nil {
		return err
	}

	*o = DescribeAccountConfigByAWSAccountIDRequest(varDescribeAccountConfigByAWSAccountIDRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "awsAccountID")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeAccountConfigByAWSAccountIDRequest struct {
	value *DescribeAccountConfigByAWSAccountIDRequest
	isSet bool
}

func (v NullableDescribeAccountConfigByAWSAccountIDRequest) Get() *DescribeAccountConfigByAWSAccountIDRequest {
	return v.value
}

func (v *NullableDescribeAccountConfigByAWSAccountIDRequest) Set(val *DescribeAccountConfigByAWSAccountIDRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeAccountConfigByAWSAccountIDRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeAccountConfigByAWSAccountIDRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeAccountConfigByAWSAccountIDRequest(val *DescribeAccountConfigByAWSAccountIDRequest) *NullableDescribeAccountConfigByAWSAccountIDRequest {
	return &NullableDescribeAccountConfigByAWSAccountIDRequest{value: val, isSet: true}
}

func (v NullableDescribeAccountConfigByAWSAccountIDRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeAccountConfigByAWSAccountIDRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


