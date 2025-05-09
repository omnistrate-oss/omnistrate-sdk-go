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

// checks if the UpdateServiceEnvironmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServiceEnvironmentRequest{}

// UpdateServiceEnvironmentRequest struct for UpdateServiceEnvironmentRequest
type UpdateServiceEnvironmentRequest struct {
	// Auto approve subscription or not
	AutoApproveSubscription *bool `json:"autoApproveSubscription,omitempty"`
	// ID of a Deployment Config
	DeploymentConfigId *string `json:"deploymentConfigId,omitempty"`
	// A brief description of the service environment
	Description *string `json:"description,omitempty"`
	// ID of a Service Environment
	Id string `json:"id"`
	// Name of the Service Environment
	Name *string `json:"name,omitempty"`
	// PEM-encoded Public key part of the key used to sign the JWT tokens for the service control plane APIs
	ServiceAuthPublicKey *string `json:"serviceAuthPublicKey,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// The ID of the service environment to use for promoting changes to this environment
	SourceEnvironmentId *string `json:"sourceEnvironmentId,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	// This parameter is used to configure the visibility of the service control-plane APIs
	Visibility *string `json:"visibility,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateServiceEnvironmentRequest UpdateServiceEnvironmentRequest

// NewUpdateServiceEnvironmentRequest instantiates a new UpdateServiceEnvironmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServiceEnvironmentRequest(id string, serviceId string, token string) *UpdateServiceEnvironmentRequest {
	this := UpdateServiceEnvironmentRequest{}
	this.Id = id
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewUpdateServiceEnvironmentRequestWithDefaults instantiates a new UpdateServiceEnvironmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceEnvironmentRequestWithDefaults() *UpdateServiceEnvironmentRequest {
	this := UpdateServiceEnvironmentRequest{}
	return &this
}

// GetAutoApproveSubscription returns the AutoApproveSubscription field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequest) GetAutoApproveSubscription() bool {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		var ret bool
		return ret
	}
	return *o.AutoApproveSubscription
}

// GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequest) GetAutoApproveSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		return nil, false
	}
	return o.AutoApproveSubscription, true
}

// HasAutoApproveSubscription returns a boolean if a field has been set.
func (o *UpdateServiceEnvironmentRequest) HasAutoApproveSubscription() bool {
	if o != nil && !IsNil(o.AutoApproveSubscription) {
		return true
	}

	return false
}

// SetAutoApproveSubscription gets a reference to the given bool and assigns it to the AutoApproveSubscription field.
func (o *UpdateServiceEnvironmentRequest) SetAutoApproveSubscription(v bool) {
	o.AutoApproveSubscription = &v
}

// GetDeploymentConfigId returns the DeploymentConfigId field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequest) GetDeploymentConfigId() string {
	if o == nil || IsNil(o.DeploymentConfigId) {
		var ret string
		return ret
	}
	return *o.DeploymentConfigId
}

// GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequest) GetDeploymentConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentConfigId) {
		return nil, false
	}
	return o.DeploymentConfigId, true
}

// HasDeploymentConfigId returns a boolean if a field has been set.
func (o *UpdateServiceEnvironmentRequest) HasDeploymentConfigId() bool {
	if o != nil && !IsNil(o.DeploymentConfigId) {
		return true
	}

	return false
}

// SetDeploymentConfigId gets a reference to the given string and assigns it to the DeploymentConfigId field.
func (o *UpdateServiceEnvironmentRequest) SetDeploymentConfigId(v string) {
	o.DeploymentConfigId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateServiceEnvironmentRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateServiceEnvironmentRequest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *UpdateServiceEnvironmentRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateServiceEnvironmentRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateServiceEnvironmentRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateServiceEnvironmentRequest) SetName(v string) {
	o.Name = &v
}

// GetServiceAuthPublicKey returns the ServiceAuthPublicKey field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequest) GetServiceAuthPublicKey() string {
	if o == nil || IsNil(o.ServiceAuthPublicKey) {
		var ret string
		return ret
	}
	return *o.ServiceAuthPublicKey
}

// GetServiceAuthPublicKeyOk returns a tuple with the ServiceAuthPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequest) GetServiceAuthPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAuthPublicKey) {
		return nil, false
	}
	return o.ServiceAuthPublicKey, true
}

// HasServiceAuthPublicKey returns a boolean if a field has been set.
func (o *UpdateServiceEnvironmentRequest) HasServiceAuthPublicKey() bool {
	if o != nil && !IsNil(o.ServiceAuthPublicKey) {
		return true
	}

	return false
}

// SetServiceAuthPublicKey gets a reference to the given string and assigns it to the ServiceAuthPublicKey field.
func (o *UpdateServiceEnvironmentRequest) SetServiceAuthPublicKey(v string) {
	o.ServiceAuthPublicKey = &v
}

// GetServiceId returns the ServiceId field value
func (o *UpdateServiceEnvironmentRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpdateServiceEnvironmentRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetSourceEnvironmentId returns the SourceEnvironmentId field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequest) GetSourceEnvironmentId() string {
	if o == nil || IsNil(o.SourceEnvironmentId) {
		var ret string
		return ret
	}
	return *o.SourceEnvironmentId
}

// GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequest) GetSourceEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceEnvironmentId) {
		return nil, false
	}
	return o.SourceEnvironmentId, true
}

// HasSourceEnvironmentId returns a boolean if a field has been set.
func (o *UpdateServiceEnvironmentRequest) HasSourceEnvironmentId() bool {
	if o != nil && !IsNil(o.SourceEnvironmentId) {
		return true
	}

	return false
}

// SetSourceEnvironmentId gets a reference to the given string and assigns it to the SourceEnvironmentId field.
func (o *UpdateServiceEnvironmentRequest) SetSourceEnvironmentId(v string) {
	o.SourceEnvironmentId = &v
}

// GetToken returns the Token field value
func (o *UpdateServiceEnvironmentRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateServiceEnvironmentRequest) SetToken(v string) {
	o.Token = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequest) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequest) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *UpdateServiceEnvironmentRequest) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *UpdateServiceEnvironmentRequest) SetVisibility(v string) {
	o.Visibility = &v
}

func (o UpdateServiceEnvironmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServiceEnvironmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoApproveSubscription) {
		toSerialize["autoApproveSubscription"] = o.AutoApproveSubscription
	}
	if !IsNil(o.DeploymentConfigId) {
		toSerialize["deploymentConfigId"] = o.DeploymentConfigId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServiceAuthPublicKey) {
		toSerialize["serviceAuthPublicKey"] = o.ServiceAuthPublicKey
	}
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.SourceEnvironmentId) {
		toSerialize["sourceEnvironmentId"] = o.SourceEnvironmentId
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateServiceEnvironmentRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateServiceEnvironmentRequest := _UpdateServiceEnvironmentRequest{}

	err = json.Unmarshal(data, &varUpdateServiceEnvironmentRequest)

	if err != nil {
		return err
	}

	*o = UpdateServiceEnvironmentRequest(varUpdateServiceEnvironmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autoApproveSubscription")
		delete(additionalProperties, "deploymentConfigId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "serviceAuthPublicKey")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "sourceEnvironmentId")
		delete(additionalProperties, "token")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateServiceEnvironmentRequest struct {
	value *UpdateServiceEnvironmentRequest
	isSet bool
}

func (v NullableUpdateServiceEnvironmentRequest) Get() *UpdateServiceEnvironmentRequest {
	return v.value
}

func (v *NullableUpdateServiceEnvironmentRequest) Set(val *UpdateServiceEnvironmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServiceEnvironmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServiceEnvironmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServiceEnvironmentRequest(val *UpdateServiceEnvironmentRequest) *NullableUpdateServiceEnvironmentRequest {
	return &NullableUpdateServiceEnvironmentRequest{value: val, isSet: true}
}

func (v NullableUpdateServiceEnvironmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServiceEnvironmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


