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

// checks if the UpdateInfraConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInfraConfigRequest{}

// UpdateInfraConfigRequest Update the infra required to host a resource
type UpdateInfraConfigRequest struct {
	// ID of a Compute Config
	ComputeConfigId *string `json:"computeConfigId,omitempty"`
	CustomTag *CustomTag `json:"customTag,omitempty"`
	// The description for the infra config
	Description *string `json:"description,omitempty"`
	// ID of an Infra Config
	Id string `json:"id"`
	// The name of the infra config
	Name *string `json:"name,omitempty"`
	// ID of a Network Config
	NetworkConfigId *string `json:"networkConfigId,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// ID of a Storage Config
	StorageConfigId *string `json:"storageConfigId,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _UpdateInfraConfigRequest UpdateInfraConfigRequest

// NewUpdateInfraConfigRequest instantiates a new UpdateInfraConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInfraConfigRequest(id string, serviceId string, token string) *UpdateInfraConfigRequest {
	this := UpdateInfraConfigRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewUpdateInfraConfigRequestWithDefaults instantiates a new UpdateInfraConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInfraConfigRequestWithDefaults() *UpdateInfraConfigRequest {
	this := UpdateInfraConfigRequest{}
	return &this
}

// GetComputeConfigId returns the ComputeConfigId field value if set, zero value otherwise.
func (o *UpdateInfraConfigRequest) GetComputeConfigId() string {
	if o == nil || IsNil(o.ComputeConfigId) {
		var ret string
		return ret
	}
	return *o.ComputeConfigId
}

// GetComputeConfigIdOk returns a tuple with the ComputeConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInfraConfigRequest) GetComputeConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComputeConfigId) {
		return nil, false
	}
	return o.ComputeConfigId, true
}

// SetComputeConfigId gets a reference to the given string and assigns it to the ComputeConfigId field.
func (o *UpdateInfraConfigRequest) SetComputeConfigId(v string) {
	o.ComputeConfigId = &v
}

// GetCustomTag returns the CustomTag field value if set, zero value otherwise.
func (o *UpdateInfraConfigRequest) GetCustomTag() CustomTag {
	if o == nil || IsNil(o.CustomTag) {
		var ret CustomTag
		return ret
	}
	return *o.CustomTag
}

// GetCustomTagOk returns a tuple with the CustomTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInfraConfigRequest) GetCustomTagOk() (*CustomTag, bool) {
	if o == nil || IsNil(o.CustomTag) {
		return nil, false
	}
	return o.CustomTag, true
}

// SetCustomTag gets a reference to the given CustomTag and assigns it to the CustomTag field.
func (o *UpdateInfraConfigRequest) SetCustomTag(v CustomTag) {
	o.CustomTag = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateInfraConfigRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInfraConfigRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateInfraConfigRequest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *UpdateInfraConfigRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateInfraConfigRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateInfraConfigRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateInfraConfigRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInfraConfigRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateInfraConfigRequest) SetName(v string) {
	o.Name = &v
}

// GetNetworkConfigId returns the NetworkConfigId field value if set, zero value otherwise.
func (o *UpdateInfraConfigRequest) GetNetworkConfigId() string {
	if o == nil || IsNil(o.NetworkConfigId) {
		var ret string
		return ret
	}
	return *o.NetworkConfigId
}

// GetNetworkConfigIdOk returns a tuple with the NetworkConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInfraConfigRequest) GetNetworkConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkConfigId) {
		return nil, false
	}
	return o.NetworkConfigId, true
}

// SetNetworkConfigId gets a reference to the given string and assigns it to the NetworkConfigId field.
func (o *UpdateInfraConfigRequest) SetNetworkConfigId(v string) {
	o.NetworkConfigId = &v
}

// GetServiceId returns the ServiceId field value
func (o *UpdateInfraConfigRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpdateInfraConfigRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpdateInfraConfigRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetStorageConfigId returns the StorageConfigId field value if set, zero value otherwise.
func (o *UpdateInfraConfigRequest) GetStorageConfigId() string {
	if o == nil || IsNil(o.StorageConfigId) {
		var ret string
		return ret
	}
	return *o.StorageConfigId
}

// GetStorageConfigIdOk returns a tuple with the StorageConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInfraConfigRequest) GetStorageConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.StorageConfigId) {
		return nil, false
	}
	return o.StorageConfigId, true
}

// SetStorageConfigId gets a reference to the given string and assigns it to the StorageConfigId field.
func (o *UpdateInfraConfigRequest) SetStorageConfigId(v string) {
	o.StorageConfigId = &v
}

// GetToken returns the Token field value
func (o *UpdateInfraConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateInfraConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateInfraConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o UpdateInfraConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInfraConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComputeConfigId) {
		toSerialize["computeConfigId"] = o.ComputeConfigId
	}
	if !IsNil(o.CustomTag) {
		toSerialize["customTag"] = o.CustomTag
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetworkConfigId) {
		toSerialize["networkConfigId"] = o.NetworkConfigId
	}
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.StorageConfigId) {
		toSerialize["storageConfigId"] = o.StorageConfigId
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateInfraConfigRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"serviceId",
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

	varUpdateInfraConfigRequest := _UpdateInfraConfigRequest{}

	err = json.Unmarshal(data, &varUpdateInfraConfigRequest)

	if err != nil {
		return err
	}

	*o = UpdateInfraConfigRequest(varUpdateInfraConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "computeConfigId")
		delete(additionalProperties, "customTag")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "networkConfigId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "storageConfigId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateInfraConfigRequest struct {
	value *UpdateInfraConfigRequest
	isSet bool
}

func (v NullableUpdateInfraConfigRequest) Get() *UpdateInfraConfigRequest {
	return v.value
}

func (v *NullableUpdateInfraConfigRequest) Set(val *UpdateInfraConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInfraConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInfraConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInfraConfigRequest(val *UpdateInfraConfigRequest) *NullableUpdateInfraConfigRequest {
	return &NullableUpdateInfraConfigRequest{value: val, isSet: true}
}

func (v NullableUpdateInfraConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInfraConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


