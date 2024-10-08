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

// checks if the FailoverResourceInstanceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailoverResourceInstanceRequestBody{}

// FailoverResourceInstanceRequestBody struct for FailoverResourceInstanceRequestBody
type FailoverResourceInstanceRequestBody struct {
	// The failed replica action
	FailedReplicaAction *string `json:"failedReplicaAction,omitempty"`
	// The failed replica ID
	FailedReplicaID string `json:"failedReplicaID"`
}

type _FailoverResourceInstanceRequestBody FailoverResourceInstanceRequestBody

// NewFailoverResourceInstanceRequestBody instantiates a new FailoverResourceInstanceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailoverResourceInstanceRequestBody(failedReplicaID string) *FailoverResourceInstanceRequestBody {
	this := FailoverResourceInstanceRequestBody{}
	this.FailedReplicaID = failedReplicaID
	return &this
}

// NewFailoverResourceInstanceRequestBodyWithDefaults instantiates a new FailoverResourceInstanceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailoverResourceInstanceRequestBodyWithDefaults() *FailoverResourceInstanceRequestBody {
	this := FailoverResourceInstanceRequestBody{}
	return &this
}

// GetFailedReplicaAction returns the FailedReplicaAction field value if set, zero value otherwise.
func (o *FailoverResourceInstanceRequestBody) GetFailedReplicaAction() string {
	if o == nil || IsNil(o.FailedReplicaAction) {
		var ret string
		return ret
	}
	return *o.FailedReplicaAction
}

// GetFailedReplicaActionOk returns a tuple with the FailedReplicaAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailoverResourceInstanceRequestBody) GetFailedReplicaActionOk() (*string, bool) {
	if o == nil || IsNil(o.FailedReplicaAction) {
		return nil, false
	}
	return o.FailedReplicaAction, true
}

// HasFailedReplicaAction returns a boolean if a field has been set.
func (o *FailoverResourceInstanceRequestBody) HasFailedReplicaAction() bool {
	if o != nil && !IsNil(o.FailedReplicaAction) {
		return true
	}

	return false
}

// SetFailedReplicaAction gets a reference to the given string and assigns it to the FailedReplicaAction field.
func (o *FailoverResourceInstanceRequestBody) SetFailedReplicaAction(v string) {
	o.FailedReplicaAction = &v
}

// GetFailedReplicaID returns the FailedReplicaID field value
func (o *FailoverResourceInstanceRequestBody) GetFailedReplicaID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FailedReplicaID
}

// GetFailedReplicaIDOk returns a tuple with the FailedReplicaID field value
// and a boolean to check if the value has been set.
func (o *FailoverResourceInstanceRequestBody) GetFailedReplicaIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedReplicaID, true
}

// SetFailedReplicaID sets field value
func (o *FailoverResourceInstanceRequestBody) SetFailedReplicaID(v string) {
	o.FailedReplicaID = v
}

func (o FailoverResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailoverResourceInstanceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailedReplicaAction) {
		toSerialize["failedReplicaAction"] = o.FailedReplicaAction
	}
	toSerialize["failedReplicaID"] = o.FailedReplicaID
	return toSerialize, nil
}

func (o *FailoverResourceInstanceRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"failedReplicaID",
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

	varFailoverResourceInstanceRequestBody := _FailoverResourceInstanceRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFailoverResourceInstanceRequestBody)

	if err != nil {
		return err
	}

	*o = FailoverResourceInstanceRequestBody(varFailoverResourceInstanceRequestBody)

	return err
}

type NullableFailoverResourceInstanceRequestBody struct {
	value *FailoverResourceInstanceRequestBody
	isSet bool
}

func (v NullableFailoverResourceInstanceRequestBody) Get() *FailoverResourceInstanceRequestBody {
	return v.value
}

func (v *NullableFailoverResourceInstanceRequestBody) Set(val *FailoverResourceInstanceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableFailoverResourceInstanceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableFailoverResourceInstanceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailoverResourceInstanceRequestBody(val *FailoverResourceInstanceRequestBody) *NullableFailoverResourceInstanceRequestBody {
	return &NullableFailoverResourceInstanceRequestBody{value: val, isSet: true}
}

func (v NullableFailoverResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailoverResourceInstanceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


