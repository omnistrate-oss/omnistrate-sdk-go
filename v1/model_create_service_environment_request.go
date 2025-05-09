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

// checks if the CreateServiceEnvironmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServiceEnvironmentRequest{}

// CreateServiceEnvironmentRequest struct for CreateServiceEnvironmentRequest
type CreateServiceEnvironmentRequest struct {
	// Auto approve subscription or not
	AutoApproveSubscription *bool `json:"autoApproveSubscription,omitempty"`
	// ID of a Deployment Config
	DeploymentConfigId string `json:"deploymentConfigId"`
	// A brief description of the service environment
	Description string `json:"description"`
	// Name of the Service Environment
	Name string `json:"name"`
	// PEM-encoded Public key part of the key used to sign the JWT tokens for the service control plane APIs
	ServiceAuthPublicKey *string `json:"serviceAuthPublicKey,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// ID of a Service Environment
	SourceEnvironmentId *string `json:"sourceEnvironmentId,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	// The type of service environment
	Type *string `json:"type,omitempty"`
	// This parameter is used to configure the visibility of the service control-plane APIs
	Visibility *string `json:"visibility,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateServiceEnvironmentRequest CreateServiceEnvironmentRequest

// NewCreateServiceEnvironmentRequest instantiates a new CreateServiceEnvironmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceEnvironmentRequest(deploymentConfigId string, description string, name string, serviceId string, token string) *CreateServiceEnvironmentRequest {
	this := CreateServiceEnvironmentRequest{}
	this.DeploymentConfigId = deploymentConfigId
	this.Description = description
	this.Name = name
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewCreateServiceEnvironmentRequestWithDefaults instantiates a new CreateServiceEnvironmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceEnvironmentRequestWithDefaults() *CreateServiceEnvironmentRequest {
	this := CreateServiceEnvironmentRequest{}
	return &this
}

// GetAutoApproveSubscription returns the AutoApproveSubscription field value if set, zero value otherwise.
func (o *CreateServiceEnvironmentRequest) GetAutoApproveSubscription() bool {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		var ret bool
		return ret
	}
	return *o.AutoApproveSubscription
}

// GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceEnvironmentRequest) GetAutoApproveSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		return nil, false
	}
	return o.AutoApproveSubscription, true
}

// SetAutoApproveSubscription gets a reference to the given bool and assigns it to the AutoApproveSubscription field.
func (o *CreateServiceEnvironmentRequest) SetAutoApproveSubscription(v bool) {
	o.AutoApproveSubscription = &v
}

// GetDeploymentConfigId returns the DeploymentConfigId field value
func (o *CreateServiceEnvironmentRequest) GetDeploymentConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentConfigId
}

// GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field value
// and a boolean to check if the value has been set.
func (o *CreateServiceEnvironmentRequest) GetDeploymentConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentConfigId, true
}

// SetDeploymentConfigId sets field value
func (o *CreateServiceEnvironmentRequest) SetDeploymentConfigId(v string) {
	o.DeploymentConfigId = v
}

// GetDescription returns the Description field value
func (o *CreateServiceEnvironmentRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateServiceEnvironmentRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateServiceEnvironmentRequest) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *CreateServiceEnvironmentRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateServiceEnvironmentRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateServiceEnvironmentRequest) SetName(v string) {
	o.Name = v
}

// GetServiceAuthPublicKey returns the ServiceAuthPublicKey field value if set, zero value otherwise.
func (o *CreateServiceEnvironmentRequest) GetServiceAuthPublicKey() string {
	if o == nil || IsNil(o.ServiceAuthPublicKey) {
		var ret string
		return ret
	}
	return *o.ServiceAuthPublicKey
}

// GetServiceAuthPublicKeyOk returns a tuple with the ServiceAuthPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceEnvironmentRequest) GetServiceAuthPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAuthPublicKey) {
		return nil, false
	}
	return o.ServiceAuthPublicKey, true
}

// SetServiceAuthPublicKey gets a reference to the given string and assigns it to the ServiceAuthPublicKey field.
func (o *CreateServiceEnvironmentRequest) SetServiceAuthPublicKey(v string) {
	o.ServiceAuthPublicKey = &v
}

// GetServiceId returns the ServiceId field value
func (o *CreateServiceEnvironmentRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *CreateServiceEnvironmentRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *CreateServiceEnvironmentRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetSourceEnvironmentId returns the SourceEnvironmentId field value if set, zero value otherwise.
func (o *CreateServiceEnvironmentRequest) GetSourceEnvironmentId() string {
	if o == nil || IsNil(o.SourceEnvironmentId) {
		var ret string
		return ret
	}
	return *o.SourceEnvironmentId
}

// GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceEnvironmentRequest) GetSourceEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceEnvironmentId) {
		return nil, false
	}
	return o.SourceEnvironmentId, true
}

// SetSourceEnvironmentId gets a reference to the given string and assigns it to the SourceEnvironmentId field.
func (o *CreateServiceEnvironmentRequest) SetSourceEnvironmentId(v string) {
	o.SourceEnvironmentId = &v
}

// GetToken returns the Token field value
func (o *CreateServiceEnvironmentRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CreateServiceEnvironmentRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CreateServiceEnvironmentRequest) SetToken(v string) {
	o.Token = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateServiceEnvironmentRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceEnvironmentRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateServiceEnvironmentRequest) SetType(v string) {
	o.Type = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *CreateServiceEnvironmentRequest) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceEnvironmentRequest) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *CreateServiceEnvironmentRequest) SetVisibility(v string) {
	o.Visibility = &v
}

func (o CreateServiceEnvironmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServiceEnvironmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoApproveSubscription) {
		toSerialize["autoApproveSubscription"] = o.AutoApproveSubscription
	}
	toSerialize["deploymentConfigId"] = o.DeploymentConfigId
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	if !IsNil(o.ServiceAuthPublicKey) {
		toSerialize["serviceAuthPublicKey"] = o.ServiceAuthPublicKey
	}
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.SourceEnvironmentId) {
		toSerialize["sourceEnvironmentId"] = o.SourceEnvironmentId
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateServiceEnvironmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deploymentConfigId",
		"description",
		"name",
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

	varCreateServiceEnvironmentRequest := _CreateServiceEnvironmentRequest{}

	err = json.Unmarshal(data, &varCreateServiceEnvironmentRequest)

	if err != nil {
		return err
	}

	*o = CreateServiceEnvironmentRequest(varCreateServiceEnvironmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autoApproveSubscription")
		delete(additionalProperties, "deploymentConfigId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "serviceAuthPublicKey")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "sourceEnvironmentId")
		delete(additionalProperties, "token")
		delete(additionalProperties, "type")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateServiceEnvironmentRequest struct {
	value *CreateServiceEnvironmentRequest
	isSet bool
}

func (v NullableCreateServiceEnvironmentRequest) Get() *CreateServiceEnvironmentRequest {
	return v.value
}

func (v *NullableCreateServiceEnvironmentRequest) Set(val *CreateServiceEnvironmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceEnvironmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceEnvironmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceEnvironmentRequest(val *CreateServiceEnvironmentRequest) *NullableCreateServiceEnvironmentRequest {
	return &NullableCreateServiceEnvironmentRequest{value: val, isSet: true}
}

func (v NullableCreateServiceEnvironmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceEnvironmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


