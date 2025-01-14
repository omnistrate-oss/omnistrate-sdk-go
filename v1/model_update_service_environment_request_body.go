/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the UpdateServiceEnvironmentRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServiceEnvironmentRequestBody{}

// UpdateServiceEnvironmentRequestBody struct for UpdateServiceEnvironmentRequestBody
type UpdateServiceEnvironmentRequestBody struct {
	// Auto approve subscription or not
	AutoApproveSubscription *bool `json:"autoApproveSubscription,omitempty"`
	// The deployment configuration ID
	DeploymentConfigId *string `json:"deploymentConfigId,omitempty"`
	// A brief description of the service environment
	Description *string `json:"description,omitempty"`
	// Name of the Service Environment
	Name *string `json:"name,omitempty"`
	// PEM-encoded Public key part of the key used to sign the JWT tokens for the service control plane APIs
	ServiceAuthPublicKey *string `json:"serviceAuthPublicKey,omitempty"`
	// The ID of the service environment to use for promoting changes to this environment
	SourceEnvironmentId *string `json:"sourceEnvironmentId,omitempty"`
	// This parameter is used to configure the visibility of the service control-plane APIs
	Visibility *string `json:"visibility,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateServiceEnvironmentRequestBody UpdateServiceEnvironmentRequestBody

// NewUpdateServiceEnvironmentRequestBody instantiates a new UpdateServiceEnvironmentRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServiceEnvironmentRequestBody() *UpdateServiceEnvironmentRequestBody {
	this := UpdateServiceEnvironmentRequestBody{}
	return &this
}

// NewUpdateServiceEnvironmentRequestBodyWithDefaults instantiates a new UpdateServiceEnvironmentRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceEnvironmentRequestBodyWithDefaults() *UpdateServiceEnvironmentRequestBody {
	this := UpdateServiceEnvironmentRequestBody{}
	return &this
}

// GetAutoApproveSubscription returns the AutoApproveSubscription field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequestBody) GetAutoApproveSubscription() bool {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		var ret bool
		return ret
	}
	return *o.AutoApproveSubscription
}

// GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequestBody) GetAutoApproveSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		return nil, false
	}
	return o.AutoApproveSubscription, true
}

// SetAutoApproveSubscription gets a reference to the given bool and assigns it to the AutoApproveSubscription field.
func (o *UpdateServiceEnvironmentRequestBody) SetAutoApproveSubscription(v bool) {
	o.AutoApproveSubscription = &v
}

// GetDeploymentConfigId returns the DeploymentConfigId field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequestBody) GetDeploymentConfigId() string {
	if o == nil || IsNil(o.DeploymentConfigId) {
		var ret string
		return ret
	}
	return *o.DeploymentConfigId
}

// GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequestBody) GetDeploymentConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentConfigId) {
		return nil, false
	}
	return o.DeploymentConfigId, true
}

// SetDeploymentConfigId gets a reference to the given string and assigns it to the DeploymentConfigId field.
func (o *UpdateServiceEnvironmentRequestBody) SetDeploymentConfigId(v string) {
	o.DeploymentConfigId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateServiceEnvironmentRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequestBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequestBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateServiceEnvironmentRequestBody) SetName(v string) {
	o.Name = &v
}

// GetServiceAuthPublicKey returns the ServiceAuthPublicKey field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequestBody) GetServiceAuthPublicKey() string {
	if o == nil || IsNil(o.ServiceAuthPublicKey) {
		var ret string
		return ret
	}
	return *o.ServiceAuthPublicKey
}

// GetServiceAuthPublicKeyOk returns a tuple with the ServiceAuthPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequestBody) GetServiceAuthPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAuthPublicKey) {
		return nil, false
	}
	return o.ServiceAuthPublicKey, true
}

// SetServiceAuthPublicKey gets a reference to the given string and assigns it to the ServiceAuthPublicKey field.
func (o *UpdateServiceEnvironmentRequestBody) SetServiceAuthPublicKey(v string) {
	o.ServiceAuthPublicKey = &v
}

// GetSourceEnvironmentId returns the SourceEnvironmentId field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequestBody) GetSourceEnvironmentId() string {
	if o == nil || IsNil(o.SourceEnvironmentId) {
		var ret string
		return ret
	}
	return *o.SourceEnvironmentId
}

// GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequestBody) GetSourceEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceEnvironmentId) {
		return nil, false
	}
	return o.SourceEnvironmentId, true
}

// SetSourceEnvironmentId gets a reference to the given string and assigns it to the SourceEnvironmentId field.
func (o *UpdateServiceEnvironmentRequestBody) SetSourceEnvironmentId(v string) {
	o.SourceEnvironmentId = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *UpdateServiceEnvironmentRequestBody) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceEnvironmentRequestBody) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *UpdateServiceEnvironmentRequestBody) SetVisibility(v string) {
	o.Visibility = &v
}

func (o UpdateServiceEnvironmentRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServiceEnvironmentRequestBody) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServiceAuthPublicKey) {
		toSerialize["serviceAuthPublicKey"] = o.ServiceAuthPublicKey
	}
	if !IsNil(o.SourceEnvironmentId) {
		toSerialize["sourceEnvironmentId"] = o.SourceEnvironmentId
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateServiceEnvironmentRequestBody) UnmarshalJSON(data []byte) (err error) {
	varUpdateServiceEnvironmentRequestBody := _UpdateServiceEnvironmentRequestBody{}

	err = json.Unmarshal(data, &varUpdateServiceEnvironmentRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateServiceEnvironmentRequestBody(varUpdateServiceEnvironmentRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autoApproveSubscription")
		delete(additionalProperties, "deploymentConfigId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "serviceAuthPublicKey")
		delete(additionalProperties, "sourceEnvironmentId")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateServiceEnvironmentRequestBody struct {
	value *UpdateServiceEnvironmentRequestBody
	isSet bool
}

func (v NullableUpdateServiceEnvironmentRequestBody) Get() *UpdateServiceEnvironmentRequestBody {
	return v.value
}

func (v *NullableUpdateServiceEnvironmentRequestBody) Set(val *UpdateServiceEnvironmentRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServiceEnvironmentRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServiceEnvironmentRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServiceEnvironmentRequestBody(val *UpdateServiceEnvironmentRequestBody) *NullableUpdateServiceEnvironmentRequestBody {
	return &NullableUpdateServiceEnvironmentRequestBody{value: val, isSet: true}
}

func (v NullableUpdateServiceEnvironmentRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServiceEnvironmentRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


