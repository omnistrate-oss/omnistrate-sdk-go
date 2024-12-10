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

// checks if the ReceiveWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceiveWebhookRequest{}

// ReceiveWebhookRequest struct for ReceiveWebhookRequest
type ReceiveWebhookRequest struct {
	// The event data
	Data *string `json:"data,omitempty"`
	// The unique id per producer.
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _ReceiveWebhookRequest ReceiveWebhookRequest

// NewReceiveWebhookRequest instantiates a new ReceiveWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiveWebhookRequest(id string) *ReceiveWebhookRequest {
	this := ReceiveWebhookRequest{}
	this.Id = id
	return &this
}

// NewReceiveWebhookRequestWithDefaults instantiates a new ReceiveWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiveWebhookRequestWithDefaults() *ReceiveWebhookRequest {
	this := ReceiveWebhookRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReceiveWebhookRequest) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveWebhookRequest) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReceiveWebhookRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *ReceiveWebhookRequest) SetData(v string) {
	o.Data = &v
}

// GetId returns the Id field value
func (o *ReceiveWebhookRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReceiveWebhookRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReceiveWebhookRequest) SetId(v string) {
	o.Id = v
}

func (o ReceiveWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiveWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReceiveWebhookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varReceiveWebhookRequest := _ReceiveWebhookRequest{}

	err = json.Unmarshal(data, &varReceiveWebhookRequest)

	if err != nil {
		return err
	}

	*o = ReceiveWebhookRequest(varReceiveWebhookRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReceiveWebhookRequest struct {
	value *ReceiveWebhookRequest
	isSet bool
}

func (v NullableReceiveWebhookRequest) Get() *ReceiveWebhookRequest {
	return v.value
}

func (v *NullableReceiveWebhookRequest) Set(val *ReceiveWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiveWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiveWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiveWebhookRequest(val *ReceiveWebhookRequest) *NullableReceiveWebhookRequest {
	return &NullableReceiveWebhookRequest{value: val, isSet: true}
}

func (v NullableReceiveWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiveWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


