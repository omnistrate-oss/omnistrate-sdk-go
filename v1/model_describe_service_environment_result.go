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

// checks if the DescribeServiceEnvironmentResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeServiceEnvironmentResult{}

// DescribeServiceEnvironmentResult struct for DescribeServiceEnvironmentResult
type DescribeServiceEnvironmentResult struct {
	// Auto approve subscription or not
	AutoApproveSubscription *bool `json:"autoApproveSubscription,omitempty"`
	// ID of a Deployment Config
	DeploymentConfigId string `json:"deploymentConfigId"`
	// A brief description of the service environment
	Description string `json:"description"`
	// ID of a Service Environment
	Id string `json:"id"`
	// Unique Key of the Service Environment
	Key string `json:"key"`
	// Name of the Service Environment
	Name string `json:"name"`
	// Type of the role
	RoleType *string `json:"roleType,omitempty"`
	// The status of an operation
	SaasPortalStatus *string `json:"saasPortalStatus,omitempty"`
	// The URL of the SaaS portal for this environment type
	SaasPortalUrl *string `json:"saasPortalUrl,omitempty"`
	// PEM-encoded Public key part of the key used to sign the JWT tokens for the service control plane APIs
	ServiceAuthPublicKey *string `json:"serviceAuthPublicKey,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// ID of a Service Environment
	SourceEnvironmentId *string `json:"sourceEnvironmentId,omitempty"`
	// The type of service environment
	Type string `json:"type"`
	// This parameter is used to configure the visibility of the service control-plane APIs
	Visibility string `json:"visibility"`
	AdditionalProperties map[string]interface{}
}

type _DescribeServiceEnvironmentResult DescribeServiceEnvironmentResult

// NewDescribeServiceEnvironmentResult instantiates a new DescribeServiceEnvironmentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeServiceEnvironmentResult(deploymentConfigId string, description string, id string, key string, name string, serviceId string, type_ string, visibility string) *DescribeServiceEnvironmentResult {
	this := DescribeServiceEnvironmentResult{}
	this.DeploymentConfigId = deploymentConfigId
	this.Description = description
	this.Id = id
	this.Key = key
	this.Name = name
	this.ServiceId = serviceId
	this.Type = type_
	this.Visibility = visibility
	return &this
}

// NewDescribeServiceEnvironmentResultWithDefaults instantiates a new DescribeServiceEnvironmentResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeServiceEnvironmentResultWithDefaults() *DescribeServiceEnvironmentResult {
	this := DescribeServiceEnvironmentResult{}
	return &this
}

// GetAutoApproveSubscription returns the AutoApproveSubscription field value if set, zero value otherwise.
func (o *DescribeServiceEnvironmentResult) GetAutoApproveSubscription() bool {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		var ret bool
		return ret
	}
	return *o.AutoApproveSubscription
}

// GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetAutoApproveSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoApproveSubscription) {
		return nil, false
	}
	return o.AutoApproveSubscription, true
}

// SetAutoApproveSubscription gets a reference to the given bool and assigns it to the AutoApproveSubscription field.
func (o *DescribeServiceEnvironmentResult) SetAutoApproveSubscription(v bool) {
	o.AutoApproveSubscription = &v
}

// GetDeploymentConfigId returns the DeploymentConfigId field value
func (o *DescribeServiceEnvironmentResult) GetDeploymentConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentConfigId
}

// GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetDeploymentConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentConfigId, true
}

// SetDeploymentConfigId sets field value
func (o *DescribeServiceEnvironmentResult) SetDeploymentConfigId(v string) {
	o.DeploymentConfigId = v
}

// GetDescription returns the Description field value
func (o *DescribeServiceEnvironmentResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeServiceEnvironmentResult) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *DescribeServiceEnvironmentResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeServiceEnvironmentResult) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *DescribeServiceEnvironmentResult) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DescribeServiceEnvironmentResult) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *DescribeServiceEnvironmentResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeServiceEnvironmentResult) SetName(v string) {
	o.Name = v
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *DescribeServiceEnvironmentResult) GetRoleType() string {
	if o == nil || IsNil(o.RoleType) {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetRoleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoleType) {
		return nil, false
	}
	return o.RoleType, true
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *DescribeServiceEnvironmentResult) SetRoleType(v string) {
	o.RoleType = &v
}

// GetSaasPortalStatus returns the SaasPortalStatus field value if set, zero value otherwise.
func (o *DescribeServiceEnvironmentResult) GetSaasPortalStatus() string {
	if o == nil || IsNil(o.SaasPortalStatus) {
		var ret string
		return ret
	}
	return *o.SaasPortalStatus
}

// GetSaasPortalStatusOk returns a tuple with the SaasPortalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetSaasPortalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SaasPortalStatus) {
		return nil, false
	}
	return o.SaasPortalStatus, true
}

// SetSaasPortalStatus gets a reference to the given string and assigns it to the SaasPortalStatus field.
func (o *DescribeServiceEnvironmentResult) SetSaasPortalStatus(v string) {
	o.SaasPortalStatus = &v
}

// GetSaasPortalUrl returns the SaasPortalUrl field value if set, zero value otherwise.
func (o *DescribeServiceEnvironmentResult) GetSaasPortalUrl() string {
	if o == nil || IsNil(o.SaasPortalUrl) {
		var ret string
		return ret
	}
	return *o.SaasPortalUrl
}

// GetSaasPortalUrlOk returns a tuple with the SaasPortalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetSaasPortalUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SaasPortalUrl) {
		return nil, false
	}
	return o.SaasPortalUrl, true
}

// SetSaasPortalUrl gets a reference to the given string and assigns it to the SaasPortalUrl field.
func (o *DescribeServiceEnvironmentResult) SetSaasPortalUrl(v string) {
	o.SaasPortalUrl = &v
}

// GetServiceAuthPublicKey returns the ServiceAuthPublicKey field value if set, zero value otherwise.
func (o *DescribeServiceEnvironmentResult) GetServiceAuthPublicKey() string {
	if o == nil || IsNil(o.ServiceAuthPublicKey) {
		var ret string
		return ret
	}
	return *o.ServiceAuthPublicKey
}

// GetServiceAuthPublicKeyOk returns a tuple with the ServiceAuthPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetServiceAuthPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAuthPublicKey) {
		return nil, false
	}
	return o.ServiceAuthPublicKey, true
}

// SetServiceAuthPublicKey gets a reference to the given string and assigns it to the ServiceAuthPublicKey field.
func (o *DescribeServiceEnvironmentResult) SetServiceAuthPublicKey(v string) {
	o.ServiceAuthPublicKey = &v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeServiceEnvironmentResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeServiceEnvironmentResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetSourceEnvironmentId returns the SourceEnvironmentId field value if set, zero value otherwise.
func (o *DescribeServiceEnvironmentResult) GetSourceEnvironmentId() string {
	if o == nil || IsNil(o.SourceEnvironmentId) {
		var ret string
		return ret
	}
	return *o.SourceEnvironmentId
}

// GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetSourceEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceEnvironmentId) {
		return nil, false
	}
	return o.SourceEnvironmentId, true
}

// SetSourceEnvironmentId gets a reference to the given string and assigns it to the SourceEnvironmentId field.
func (o *DescribeServiceEnvironmentResult) SetSourceEnvironmentId(v string) {
	o.SourceEnvironmentId = &v
}

// GetType returns the Type field value
func (o *DescribeServiceEnvironmentResult) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DescribeServiceEnvironmentResult) SetType(v string) {
	o.Type = v
}

// GetVisibility returns the Visibility field value
func (o *DescribeServiceEnvironmentResult) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *DescribeServiceEnvironmentResult) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *DescribeServiceEnvironmentResult) SetVisibility(v string) {
	o.Visibility = v
}

func (o DescribeServiceEnvironmentResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeServiceEnvironmentResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoApproveSubscription) {
		toSerialize["autoApproveSubscription"] = o.AutoApproveSubscription
	}
	toSerialize["deploymentConfigId"] = o.DeploymentConfigId
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	if !IsNil(o.RoleType) {
		toSerialize["roleType"] = o.RoleType
	}
	if !IsNil(o.SaasPortalStatus) {
		toSerialize["saasPortalStatus"] = o.SaasPortalStatus
	}
	if !IsNil(o.SaasPortalUrl) {
		toSerialize["saasPortalUrl"] = o.SaasPortalUrl
	}
	if !IsNil(o.ServiceAuthPublicKey) {
		toSerialize["serviceAuthPublicKey"] = o.ServiceAuthPublicKey
	}
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.SourceEnvironmentId) {
		toSerialize["sourceEnvironmentId"] = o.SourceEnvironmentId
	}
	toSerialize["type"] = o.Type
	toSerialize["visibility"] = o.Visibility

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeServiceEnvironmentResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deploymentConfigId",
		"description",
		"id",
		"key",
		"name",
		"serviceId",
		"type",
		"visibility",
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

	varDescribeServiceEnvironmentResult := _DescribeServiceEnvironmentResult{}

	err = json.Unmarshal(data, &varDescribeServiceEnvironmentResult)

	if err != nil {
		return err
	}

	*o = DescribeServiceEnvironmentResult(varDescribeServiceEnvironmentResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autoApproveSubscription")
		delete(additionalProperties, "deploymentConfigId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "roleType")
		delete(additionalProperties, "saasPortalStatus")
		delete(additionalProperties, "saasPortalUrl")
		delete(additionalProperties, "serviceAuthPublicKey")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "sourceEnvironmentId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeServiceEnvironmentResult struct {
	value *DescribeServiceEnvironmentResult
	isSet bool
}

func (v NullableDescribeServiceEnvironmentResult) Get() *DescribeServiceEnvironmentResult {
	return v.value
}

func (v *NullableDescribeServiceEnvironmentResult) Set(val *DescribeServiceEnvironmentResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeServiceEnvironmentResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeServiceEnvironmentResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeServiceEnvironmentResult(val *DescribeServiceEnvironmentResult) *NullableDescribeServiceEnvironmentResult {
	return &NullableDescribeServiceEnvironmentResult{value: val, isSet: true}
}

func (v NullableDescribeServiceEnvironmentResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeServiceEnvironmentResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


