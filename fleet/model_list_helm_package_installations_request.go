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

// checks if the ListHelmPackageInstallationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHelmPackageInstallationsRequest{}

// ListHelmPackageInstallationsRequest struct for ListHelmPackageInstallationsRequest
type ListHelmPackageInstallationsRequest struct {
	// ID of a Host Cluster
	HostClusterID *string `json:"hostClusterID,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListHelmPackageInstallationsRequest ListHelmPackageInstallationsRequest

// NewListHelmPackageInstallationsRequest instantiates a new ListHelmPackageInstallationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHelmPackageInstallationsRequest(token string) *ListHelmPackageInstallationsRequest {
	this := ListHelmPackageInstallationsRequest{}
	this.Token = token
	return &this
}

// NewListHelmPackageInstallationsRequestWithDefaults instantiates a new ListHelmPackageInstallationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHelmPackageInstallationsRequestWithDefaults() *ListHelmPackageInstallationsRequest {
	this := ListHelmPackageInstallationsRequest{}
	return &this
}

// GetHostClusterID returns the HostClusterID field value if set, zero value otherwise.
func (o *ListHelmPackageInstallationsRequest) GetHostClusterID() string {
	if o == nil || IsNil(o.HostClusterID) {
		var ret string
		return ret
	}
	return *o.HostClusterID
}

// GetHostClusterIDOk returns a tuple with the HostClusterID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHelmPackageInstallationsRequest) GetHostClusterIDOk() (*string, bool) {
	if o == nil || IsNil(o.HostClusterID) {
		return nil, false
	}
	return o.HostClusterID, true
}

// HasHostClusterID returns a boolean if a field has been set.
func (o *ListHelmPackageInstallationsRequest) HasHostClusterID() bool {
	if o != nil && !IsNil(o.HostClusterID) {
		return true
	}

	return false
}

// SetHostClusterID gets a reference to the given string and assigns it to the HostClusterID field.
func (o *ListHelmPackageInstallationsRequest) SetHostClusterID(v string) {
	o.HostClusterID = &v
}

// GetToken returns the Token field value
func (o *ListHelmPackageInstallationsRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListHelmPackageInstallationsRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListHelmPackageInstallationsRequest) SetToken(v string) {
	o.Token = v
}

func (o ListHelmPackageInstallationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHelmPackageInstallationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostClusterID) {
		toSerialize["hostClusterID"] = o.HostClusterID
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListHelmPackageInstallationsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varListHelmPackageInstallationsRequest := _ListHelmPackageInstallationsRequest{}

	err = json.Unmarshal(data, &varListHelmPackageInstallationsRequest)

	if err != nil {
		return err
	}

	*o = ListHelmPackageInstallationsRequest(varListHelmPackageInstallationsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hostClusterID")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListHelmPackageInstallationsRequest struct {
	value *ListHelmPackageInstallationsRequest
	isSet bool
}

func (v NullableListHelmPackageInstallationsRequest) Get() *ListHelmPackageInstallationsRequest {
	return v.value
}

func (v *NullableListHelmPackageInstallationsRequest) Set(val *ListHelmPackageInstallationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListHelmPackageInstallationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListHelmPackageInstallationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHelmPackageInstallationsRequest(val *ListHelmPackageInstallationsRequest) *NullableListHelmPackageInstallationsRequest {
	return &NullableListHelmPackageInstallationsRequest{value: val, isSet: true}
}

func (v NullableListHelmPackageInstallationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHelmPackageInstallationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


