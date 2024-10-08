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

// checks if the CreateSubscriptionRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionRequestBody{}

// CreateSubscriptionRequestBody struct for CreateSubscriptionRequestBody
type CreateSubscriptionRequestBody struct {
	// The product tier ID
	ProductTierId string `json:"productTierId"`
	// The service ID
	ServiceId string `json:"serviceId"`
}

type _CreateSubscriptionRequestBody CreateSubscriptionRequestBody

// NewCreateSubscriptionRequestBody instantiates a new CreateSubscriptionRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubscriptionRequestBody(productTierId string, serviceId string) *CreateSubscriptionRequestBody {
	this := CreateSubscriptionRequestBody{}
	this.ProductTierId = productTierId
	this.ServiceId = serviceId
	return &this
}

// NewCreateSubscriptionRequestBodyWithDefaults instantiates a new CreateSubscriptionRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubscriptionRequestBodyWithDefaults() *CreateSubscriptionRequestBody {
	this := CreateSubscriptionRequestBody{}
	return &this
}

// GetProductTierId returns the ProductTierId field value
func (o *CreateSubscriptionRequestBody) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionRequestBody) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *CreateSubscriptionRequestBody) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetServiceId returns the ServiceId field value
func (o *CreateSubscriptionRequestBody) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionRequestBody) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *CreateSubscriptionRequestBody) SetServiceId(v string) {
	o.ServiceId = v
}

func (o CreateSubscriptionRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productTierId"] = o.ProductTierId
	toSerialize["serviceId"] = o.ServiceId
	return toSerialize, nil
}

func (o *CreateSubscriptionRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productTierId",
		"serviceId",
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

	varCreateSubscriptionRequestBody := _CreateSubscriptionRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSubscriptionRequestBody)

	if err != nil {
		return err
	}

	*o = CreateSubscriptionRequestBody(varCreateSubscriptionRequestBody)

	return err
}

type NullableCreateSubscriptionRequestBody struct {
	value *CreateSubscriptionRequestBody
	isSet bool
}

func (v NullableCreateSubscriptionRequestBody) Get() *CreateSubscriptionRequestBody {
	return v.value
}

func (v *NullableCreateSubscriptionRequestBody) Set(val *CreateSubscriptionRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubscriptionRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubscriptionRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubscriptionRequestBody(val *CreateSubscriptionRequestBody) *NullableCreateSubscriptionRequestBody {
	return &NullableCreateSubscriptionRequestBody{value: val, isSet: true}
}

func (v NullableCreateSubscriptionRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubscriptionRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


