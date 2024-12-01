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

// checks if the ServiceDeploymentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDeploymentDetails{}

// ServiceDeploymentDetails struct for ServiceDeploymentDetails
type ServiceDeploymentDetails struct {
	// The deployment keys that this deployment depends on
	DependsOnDeployment []string `json:"dependsOnDeployment,omitempty"`
	// The deployment key
	DeploymentKey string `json:"deploymentKey"`
	// The instance ID
	InstanceId string `json:"instanceId"`
	// The service ID
	ServiceId string `json:"serviceId"`
	// The instance subscription ID
	SubscriptionId string `json:"subscriptionId"`
	AdditionalProperties map[string]interface{}
}

type _ServiceDeploymentDetails ServiceDeploymentDetails

// NewServiceDeploymentDetails instantiates a new ServiceDeploymentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDeploymentDetails(deploymentKey string, instanceId string, serviceId string, subscriptionId string) *ServiceDeploymentDetails {
	this := ServiceDeploymentDetails{}
	this.DeploymentKey = deploymentKey
	this.InstanceId = instanceId
	this.ServiceId = serviceId
	this.SubscriptionId = subscriptionId
	return &this
}

// NewServiceDeploymentDetailsWithDefaults instantiates a new ServiceDeploymentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDeploymentDetailsWithDefaults() *ServiceDeploymentDetails {
	this := ServiceDeploymentDetails{}
	return &this
}

// GetDependsOnDeployment returns the DependsOnDeployment field value if set, zero value otherwise.
func (o *ServiceDeploymentDetails) GetDependsOnDeployment() []string {
	if o == nil || IsNil(o.DependsOnDeployment) {
		var ret []string
		return ret
	}
	return o.DependsOnDeployment
}

// GetDependsOnDeploymentOk returns a tuple with the DependsOnDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDeploymentDetails) GetDependsOnDeploymentOk() ([]string, bool) {
	if o == nil || IsNil(o.DependsOnDeployment) {
		return nil, false
	}
	return o.DependsOnDeployment, true
}

// SetDependsOnDeployment gets a reference to the given []string and assigns it to the DependsOnDeployment field.
func (o *ServiceDeploymentDetails) SetDependsOnDeployment(v []string) {
	o.DependsOnDeployment = v
}

// GetDeploymentKey returns the DeploymentKey field value
func (o *ServiceDeploymentDetails) GetDeploymentKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentKey
}

// GetDeploymentKeyOk returns a tuple with the DeploymentKey field value
// and a boolean to check if the value has been set.
func (o *ServiceDeploymentDetails) GetDeploymentKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentKey, true
}

// SetDeploymentKey sets field value
func (o *ServiceDeploymentDetails) SetDeploymentKey(v string) {
	o.DeploymentKey = v
}

// GetInstanceId returns the InstanceId field value
func (o *ServiceDeploymentDetails) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *ServiceDeploymentDetails) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *ServiceDeploymentDetails) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetServiceId returns the ServiceId field value
func (o *ServiceDeploymentDetails) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ServiceDeploymentDetails) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ServiceDeploymentDetails) SetServiceId(v string) {
	o.ServiceId = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *ServiceDeploymentDetails) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *ServiceDeploymentDetails) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *ServiceDeploymentDetails) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o ServiceDeploymentDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDeploymentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DependsOnDeployment) {
		toSerialize["dependsOnDeployment"] = o.DependsOnDeployment
	}
	toSerialize["deploymentKey"] = o.DeploymentKey
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["subscriptionId"] = o.SubscriptionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceDeploymentDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deploymentKey",
		"instanceId",
		"serviceId",
		"subscriptionId",
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

	varServiceDeploymentDetails := _ServiceDeploymentDetails{}

	err = json.Unmarshal(data, &varServiceDeploymentDetails)

	if err != nil {
		return err
	}

	*o = ServiceDeploymentDetails(varServiceDeploymentDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dependsOnDeployment")
		delete(additionalProperties, "deploymentKey")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "subscriptionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceDeploymentDetails struct {
	value *ServiceDeploymentDetails
	isSet bool
}

func (v NullableServiceDeploymentDetails) Get() *ServiceDeploymentDetails {
	return v.value
}

func (v *NullableServiceDeploymentDetails) Set(val *ServiceDeploymentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDeploymentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDeploymentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDeploymentDetails(val *ServiceDeploymentDetails) *NullableServiceDeploymentDetails {
	return &NullableServiceDeploymentDetails{value: val, isSet: true}
}

func (v NullableServiceDeploymentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDeploymentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


