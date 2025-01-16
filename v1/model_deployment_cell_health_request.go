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

// checks if the DeploymentCellHealthRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentCellHealthRequest{}

// DeploymentCellHealthRequest struct for DeploymentCellHealthRequest
type DeploymentCellHealthRequest struct {
	// ID of a Host Cluster
	HostClusterID *string `json:"hostClusterID,omitempty"`
	// ID of a Service Environment
	ServiceEnvironmentID *string `json:"serviceEnvironmentID,omitempty"`
	// ID of a Service
	ServiceID *string `json:"serviceID,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentCellHealthRequest DeploymentCellHealthRequest

// NewDeploymentCellHealthRequest instantiates a new DeploymentCellHealthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentCellHealthRequest(token string) *DeploymentCellHealthRequest {
	this := DeploymentCellHealthRequest{}
	this.Token = token
	return &this
}

// NewDeploymentCellHealthRequestWithDefaults instantiates a new DeploymentCellHealthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentCellHealthRequestWithDefaults() *DeploymentCellHealthRequest {
	this := DeploymentCellHealthRequest{}
	return &this
}

// GetHostClusterID returns the HostClusterID field value if set, zero value otherwise.
func (o *DeploymentCellHealthRequest) GetHostClusterID() string {
	if o == nil || IsNil(o.HostClusterID) {
		var ret string
		return ret
	}
	return *o.HostClusterID
}

// GetHostClusterIDOk returns a tuple with the HostClusterID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentCellHealthRequest) GetHostClusterIDOk() (*string, bool) {
	if o == nil || IsNil(o.HostClusterID) {
		return nil, false
	}
	return o.HostClusterID, true
}

// SetHostClusterID gets a reference to the given string and assigns it to the HostClusterID field.
func (o *DeploymentCellHealthRequest) SetHostClusterID(v string) {
	o.HostClusterID = &v
}

// GetServiceEnvironmentID returns the ServiceEnvironmentID field value if set, zero value otherwise.
func (o *DeploymentCellHealthRequest) GetServiceEnvironmentID() string {
	if o == nil || IsNil(o.ServiceEnvironmentID) {
		var ret string
		return ret
	}
	return *o.ServiceEnvironmentID
}

// GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentCellHealthRequest) GetServiceEnvironmentIDOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceEnvironmentID) {
		return nil, false
	}
	return o.ServiceEnvironmentID, true
}

// SetServiceEnvironmentID gets a reference to the given string and assigns it to the ServiceEnvironmentID field.
func (o *DeploymentCellHealthRequest) SetServiceEnvironmentID(v string) {
	o.ServiceEnvironmentID = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *DeploymentCellHealthRequest) GetServiceID() string {
	if o == nil || IsNil(o.ServiceID) {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentCellHealthRequest) GetServiceIDOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceID) {
		return nil, false
	}
	return o.ServiceID, true
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *DeploymentCellHealthRequest) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetToken returns the Token field value
func (o *DeploymentCellHealthRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DeploymentCellHealthRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DeploymentCellHealthRequest) SetToken(v string) {
	o.Token = v
}

func (o DeploymentCellHealthRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentCellHealthRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostClusterID) {
		toSerialize["hostClusterID"] = o.HostClusterID
	}
	if !IsNil(o.ServiceEnvironmentID) {
		toSerialize["serviceEnvironmentID"] = o.ServiceEnvironmentID
	}
	if !IsNil(o.ServiceID) {
		toSerialize["serviceID"] = o.ServiceID
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentCellHealthRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDeploymentCellHealthRequest := _DeploymentCellHealthRequest{}

	err = json.Unmarshal(data, &varDeploymentCellHealthRequest)

	if err != nil {
		return err
	}

	*o = DeploymentCellHealthRequest(varDeploymentCellHealthRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hostClusterID")
		delete(additionalProperties, "serviceEnvironmentID")
		delete(additionalProperties, "serviceID")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentCellHealthRequest struct {
	value *DeploymentCellHealthRequest
	isSet bool
}

func (v NullableDeploymentCellHealthRequest) Get() *DeploymentCellHealthRequest {
	return v.value
}

func (v *NullableDeploymentCellHealthRequest) Set(val *DeploymentCellHealthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentCellHealthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentCellHealthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentCellHealthRequest(val *DeploymentCellHealthRequest) *NullableDeploymentCellHealthRequest {
	return &NullableDeploymentCellHealthRequest{value: val, isSet: true}
}

func (v NullableDeploymentCellHealthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentCellHealthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


