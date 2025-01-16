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

// checks if the UpdateResourceInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateResourceInstanceRequest{}

// UpdateResourceInstanceRequest struct for UpdateResourceInstanceRequest
type UpdateResourceInstanceRequest struct {
	// The instance ID
	Id string `json:"id"`
	// The product tier name
	ProductTierKey string `json:"productTierKey"`
	// The request parameters
	RequestParams interface{} `json:"requestParams,omitempty"`
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
	// The subscription ID
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _UpdateResourceInstanceRequest UpdateResourceInstanceRequest

// NewUpdateResourceInstanceRequest instantiates a new UpdateResourceInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateResourceInstanceRequest(id string, productTierKey string, resourceKey string, serviceAPIVersion string, serviceEnvironmentKey string, serviceKey string, serviceModelKey string, serviceProviderId string, token string) *UpdateResourceInstanceRequest {
	this := UpdateResourceInstanceRequest{}
	this.Id = id
	this.ProductTierKey = productTierKey
	this.ResourceKey = resourceKey
	this.ServiceAPIVersion = serviceAPIVersion
	this.ServiceEnvironmentKey = serviceEnvironmentKey
	this.ServiceKey = serviceKey
	this.ServiceModelKey = serviceModelKey
	this.ServiceProviderId = serviceProviderId
	this.Token = token
	return &this
}

// NewUpdateResourceInstanceRequestWithDefaults instantiates a new UpdateResourceInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateResourceInstanceRequestWithDefaults() *UpdateResourceInstanceRequest {
	this := UpdateResourceInstanceRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateResourceInstanceRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceInstanceRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateResourceInstanceRequest) SetId(v string) {
	o.Id = v
}

// GetProductTierKey returns the ProductTierKey field value
func (o *UpdateResourceInstanceRequest) GetProductTierKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierKey
}

// GetProductTierKeyOk returns a tuple with the ProductTierKey field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceInstanceRequest) GetProductTierKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierKey, true
}

// SetProductTierKey sets field value
func (o *UpdateResourceInstanceRequest) SetProductTierKey(v string) {
	o.ProductTierKey = v
}

// GetRequestParams returns the RequestParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateResourceInstanceRequest) GetRequestParams() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RequestParams
}

// GetRequestParamsOk returns a tuple with the RequestParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateResourceInstanceRequest) GetRequestParamsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RequestParams) {
		return nil, false
	}
	return &o.RequestParams, true
}

// SetRequestParams gets a reference to the given interface{} and assigns it to the RequestParams field.
func (o *UpdateResourceInstanceRequest) SetRequestParams(v interface{}) {
	o.RequestParams = v
}

// GetResourceKey returns the ResourceKey field value
func (o *UpdateResourceInstanceRequest) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceInstanceRequest) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *UpdateResourceInstanceRequest) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetServiceAPIVersion returns the ServiceAPIVersion field value
func (o *UpdateResourceInstanceRequest) GetServiceAPIVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAPIVersion
}

// GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceInstanceRequest) GetServiceAPIVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAPIVersion, true
}

// SetServiceAPIVersion sets field value
func (o *UpdateResourceInstanceRequest) SetServiceAPIVersion(v string) {
	o.ServiceAPIVersion = v
}

// GetServiceEnvironmentKey returns the ServiceEnvironmentKey field value
func (o *UpdateResourceInstanceRequest) GetServiceEnvironmentKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentKey
}

// GetServiceEnvironmentKeyOk returns a tuple with the ServiceEnvironmentKey field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceInstanceRequest) GetServiceEnvironmentKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentKey, true
}

// SetServiceEnvironmentKey sets field value
func (o *UpdateResourceInstanceRequest) SetServiceEnvironmentKey(v string) {
	o.ServiceEnvironmentKey = v
}

// GetServiceKey returns the ServiceKey field value
func (o *UpdateResourceInstanceRequest) GetServiceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceInstanceRequest) GetServiceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceKey, true
}

// SetServiceKey sets field value
func (o *UpdateResourceInstanceRequest) SetServiceKey(v string) {
	o.ServiceKey = v
}

// GetServiceModelKey returns the ServiceModelKey field value
func (o *UpdateResourceInstanceRequest) GetServiceModelKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelKey
}

// GetServiceModelKeyOk returns a tuple with the ServiceModelKey field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceInstanceRequest) GetServiceModelKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelKey, true
}

// SetServiceModelKey sets field value
func (o *UpdateResourceInstanceRequest) SetServiceModelKey(v string) {
	o.ServiceModelKey = v
}

// GetServiceProviderId returns the ServiceProviderId field value
func (o *UpdateResourceInstanceRequest) GetServiceProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceProviderId
}

// GetServiceProviderIdOk returns a tuple with the ServiceProviderId field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceInstanceRequest) GetServiceProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceProviderId, true
}

// SetServiceProviderId sets field value
func (o *UpdateResourceInstanceRequest) SetServiceProviderId(v string) {
	o.ServiceProviderId = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UpdateResourceInstanceRequest) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResourceInstanceRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UpdateResourceInstanceRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetToken returns the Token field value
func (o *UpdateResourceInstanceRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceInstanceRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateResourceInstanceRequest) SetToken(v string) {
	o.Token = v
}

func (o UpdateResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateResourceInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["productTierKey"] = o.ProductTierKey
	if o.RequestParams != nil {
		toSerialize["requestParams"] = o.RequestParams
	}
	toSerialize["resourceKey"] = o.ResourceKey
	toSerialize["serviceAPIVersion"] = o.ServiceAPIVersion
	toSerialize["serviceEnvironmentKey"] = o.ServiceEnvironmentKey
	toSerialize["serviceKey"] = o.ServiceKey
	toSerialize["serviceModelKey"] = o.ServiceModelKey
	toSerialize["serviceProviderId"] = o.ServiceProviderId
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateResourceInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"productTierKey",
		"resourceKey",
		"serviceAPIVersion",
		"serviceEnvironmentKey",
		"serviceKey",
		"serviceModelKey",
		"serviceProviderId",
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

	varUpdateResourceInstanceRequest := _UpdateResourceInstanceRequest{}

	err = json.Unmarshal(data, &varUpdateResourceInstanceRequest)

	if err != nil {
		return err
	}

	*o = UpdateResourceInstanceRequest(varUpdateResourceInstanceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "productTierKey")
		delete(additionalProperties, "requestParams")
		delete(additionalProperties, "resourceKey")
		delete(additionalProperties, "serviceAPIVersion")
		delete(additionalProperties, "serviceEnvironmentKey")
		delete(additionalProperties, "serviceKey")
		delete(additionalProperties, "serviceModelKey")
		delete(additionalProperties, "serviceProviderId")
		delete(additionalProperties, "subscriptionId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateResourceInstanceRequest struct {
	value *UpdateResourceInstanceRequest
	isSet bool
}

func (v NullableUpdateResourceInstanceRequest) Get() *UpdateResourceInstanceRequest {
	return v.value
}

func (v *NullableUpdateResourceInstanceRequest) Set(val *UpdateResourceInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateResourceInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateResourceInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateResourceInstanceRequest(val *UpdateResourceInstanceRequest) *NullableUpdateResourceInstanceRequest {
	return &NullableUpdateResourceInstanceRequest{value: val, isSet: true}
}

func (v NullableUpdateResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateResourceInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


