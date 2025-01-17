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

// checks if the RestoreResourceInstanceFromSnapshotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreResourceInstanceFromSnapshotRequest{}

// RestoreResourceInstanceFromSnapshotRequest struct for RestoreResourceInstanceFromSnapshotRequest
type RestoreResourceInstanceFromSnapshotRequest struct {
	// The network type
	NetworkType *string `json:"network_type,omitempty"`
	// The product tier name
	ProductTierKey string `json:"productTierKey"`
	// The resource key
	ResourceKey string `json:"resourceKey"`
	// The service API version
	ServiceAPIVersion string `json:"serviceAPIVersion"`
	// The service environment name
	ServiceEnvironmentKey string `json:"serviceEnvironmentKey"`
	// The service name
	ServiceKey string `json:"serviceKey"`
	// The service model name
	ServiceModelKey string `json:"serviceModelKey"`
	// ID of a Service Provider
	ServiceProviderId string `json:"serviceProviderId"`
	// ID of a Resource Instance Snapshot
	SnapshotId string `json:"snapshotId"`
	// The subscription ID
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _RestoreResourceInstanceFromSnapshotRequest RestoreResourceInstanceFromSnapshotRequest

// NewRestoreResourceInstanceFromSnapshotRequest instantiates a new RestoreResourceInstanceFromSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreResourceInstanceFromSnapshotRequest(productTierKey string, resourceKey string, serviceAPIVersion string, serviceEnvironmentKey string, serviceKey string, serviceModelKey string, serviceProviderId string, snapshotId string, token string) *RestoreResourceInstanceFromSnapshotRequest {
	this := RestoreResourceInstanceFromSnapshotRequest{}
	this.ProductTierKey = productTierKey
	this.ResourceKey = resourceKey
	this.ServiceAPIVersion = serviceAPIVersion
	this.ServiceEnvironmentKey = serviceEnvironmentKey
	this.ServiceKey = serviceKey
	this.ServiceModelKey = serviceModelKey
	this.ServiceProviderId = serviceProviderId
	this.SnapshotId = snapshotId
	this.Token = token
	return &this
}

// NewRestoreResourceInstanceFromSnapshotRequestWithDefaults instantiates a new RestoreResourceInstanceFromSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreResourceInstanceFromSnapshotRequestWithDefaults() *RestoreResourceInstanceFromSnapshotRequest {
	this := RestoreResourceInstanceFromSnapshotRequest{}
	return &this
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetNetworkType() string {
	if o == nil || IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) HasNetworkType() bool {
	if o != nil && !IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *RestoreResourceInstanceFromSnapshotRequest) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetProductTierKey returns the ProductTierKey field value
func (o *RestoreResourceInstanceFromSnapshotRequest) GetProductTierKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierKey
}

// GetProductTierKeyOk returns a tuple with the ProductTierKey field value
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetProductTierKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierKey, true
}

// SetProductTierKey sets field value
func (o *RestoreResourceInstanceFromSnapshotRequest) SetProductTierKey(v string) {
	o.ProductTierKey = v
}

// GetResourceKey returns the ResourceKey field value
func (o *RestoreResourceInstanceFromSnapshotRequest) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *RestoreResourceInstanceFromSnapshotRequest) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetServiceAPIVersion returns the ServiceAPIVersion field value
func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceAPIVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAPIVersion
}

// GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field value
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceAPIVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAPIVersion, true
}

// SetServiceAPIVersion sets field value
func (o *RestoreResourceInstanceFromSnapshotRequest) SetServiceAPIVersion(v string) {
	o.ServiceAPIVersion = v
}

// GetServiceEnvironmentKey returns the ServiceEnvironmentKey field value
func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceEnvironmentKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentKey
}

// GetServiceEnvironmentKeyOk returns a tuple with the ServiceEnvironmentKey field value
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceEnvironmentKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentKey, true
}

// SetServiceEnvironmentKey sets field value
func (o *RestoreResourceInstanceFromSnapshotRequest) SetServiceEnvironmentKey(v string) {
	o.ServiceEnvironmentKey = v
}

// GetServiceKey returns the ServiceKey field value
func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceKey, true
}

// SetServiceKey sets field value
func (o *RestoreResourceInstanceFromSnapshotRequest) SetServiceKey(v string) {
	o.ServiceKey = v
}

// GetServiceModelKey returns the ServiceModelKey field value
func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceModelKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelKey
}

// GetServiceModelKeyOk returns a tuple with the ServiceModelKey field value
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceModelKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelKey, true
}

// SetServiceModelKey sets field value
func (o *RestoreResourceInstanceFromSnapshotRequest) SetServiceModelKey(v string) {
	o.ServiceModelKey = v
}

// GetServiceProviderId returns the ServiceProviderId field value
func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceProviderId
}

// GetServiceProviderIdOk returns a tuple with the ServiceProviderId field value
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceProviderId, true
}

// SetServiceProviderId sets field value
func (o *RestoreResourceInstanceFromSnapshotRequest) SetServiceProviderId(v string) {
	o.ServiceProviderId = v
}

// GetSnapshotId returns the SnapshotId field value
func (o *RestoreResourceInstanceFromSnapshotRequest) GetSnapshotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotId, true
}

// SetSnapshotId sets field value
func (o *RestoreResourceInstanceFromSnapshotRequest) SetSnapshotId(v string) {
	o.SnapshotId = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *RestoreResourceInstanceFromSnapshotRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetToken returns the Token field value
func (o *RestoreResourceInstanceFromSnapshotRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RestoreResourceInstanceFromSnapshotRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RestoreResourceInstanceFromSnapshotRequest) SetToken(v string) {
	o.Token = v
}

func (o RestoreResourceInstanceFromSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreResourceInstanceFromSnapshotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkType) {
		toSerialize["network_type"] = o.NetworkType
	}
	toSerialize["productTierKey"] = o.ProductTierKey
	toSerialize["resourceKey"] = o.ResourceKey
	toSerialize["serviceAPIVersion"] = o.ServiceAPIVersion
	toSerialize["serviceEnvironmentKey"] = o.ServiceEnvironmentKey
	toSerialize["serviceKey"] = o.ServiceKey
	toSerialize["serviceModelKey"] = o.ServiceModelKey
	toSerialize["serviceProviderId"] = o.ServiceProviderId
	toSerialize["snapshotId"] = o.SnapshotId
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RestoreResourceInstanceFromSnapshotRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productTierKey",
		"resourceKey",
		"serviceAPIVersion",
		"serviceEnvironmentKey",
		"serviceKey",
		"serviceModelKey",
		"serviceProviderId",
		"snapshotId",
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

	varRestoreResourceInstanceFromSnapshotRequest := _RestoreResourceInstanceFromSnapshotRequest{}

	err = json.Unmarshal(data, &varRestoreResourceInstanceFromSnapshotRequest)

	if err != nil {
		return err
	}

	*o = RestoreResourceInstanceFromSnapshotRequest(varRestoreResourceInstanceFromSnapshotRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network_type")
		delete(additionalProperties, "productTierKey")
		delete(additionalProperties, "resourceKey")
		delete(additionalProperties, "serviceAPIVersion")
		delete(additionalProperties, "serviceEnvironmentKey")
		delete(additionalProperties, "serviceKey")
		delete(additionalProperties, "serviceModelKey")
		delete(additionalProperties, "serviceProviderId")
		delete(additionalProperties, "snapshotId")
		delete(additionalProperties, "subscriptionId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRestoreResourceInstanceFromSnapshotRequest struct {
	value *RestoreResourceInstanceFromSnapshotRequest
	isSet bool
}

func (v NullableRestoreResourceInstanceFromSnapshotRequest) Get() *RestoreResourceInstanceFromSnapshotRequest {
	return v.value
}

func (v *NullableRestoreResourceInstanceFromSnapshotRequest) Set(val *RestoreResourceInstanceFromSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreResourceInstanceFromSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreResourceInstanceFromSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreResourceInstanceFromSnapshotRequest(val *RestoreResourceInstanceFromSnapshotRequest) *NullableRestoreResourceInstanceFromSnapshotRequest {
	return &NullableRestoreResourceInstanceFromSnapshotRequest{value: val, isSet: true}
}

func (v NullableRestoreResourceInstanceFromSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreResourceInstanceFromSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


