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

// checks if the AdoptResourceInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdoptResourceInstanceRequest{}

// AdoptResourceInstanceRequest struct for AdoptResourceInstanceRequest
type AdoptResourceInstanceRequest struct {
	// The cloud provider name
	CloudProvider *string `json:"cloud_provider,omitempty"`
	// Custom network for resource
	CustomNetworkId *string `json:"custom_network_id,omitempty"`
	// The external payer id to record which customer should pay for this resource instance. This will override the subscription level external payer id if set.
	ExternalPayerId *string `json:"externalPayerId,omitempty"`
	// The product tier name
	ProductTierKey string `json:"productTierKey"`
	// The product tier version
	ProductTierVersion *string `json:"productTierVersion,omitempty"`
	// The region code
	Region *string `json:"region,omitempty"`
	// The resource configuration
	ResourceConfiguration map[string]interface{} `json:"resourceConfiguration,omitempty"`
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
	// ID of a Subscription
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _AdoptResourceInstanceRequest AdoptResourceInstanceRequest

// NewAdoptResourceInstanceRequest instantiates a new AdoptResourceInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdoptResourceInstanceRequest(productTierKey string, serviceAPIVersion string, serviceEnvironmentKey string, serviceKey string, serviceModelKey string, serviceProviderId string, token string) *AdoptResourceInstanceRequest {
	this := AdoptResourceInstanceRequest{}
	this.ProductTierKey = productTierKey
	this.ServiceAPIVersion = serviceAPIVersion
	this.ServiceEnvironmentKey = serviceEnvironmentKey
	this.ServiceKey = serviceKey
	this.ServiceModelKey = serviceModelKey
	this.ServiceProviderId = serviceProviderId
	this.Token = token
	return &this
}

// NewAdoptResourceInstanceRequestWithDefaults instantiates a new AdoptResourceInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdoptResourceInstanceRequestWithDefaults() *AdoptResourceInstanceRequest {
	this := AdoptResourceInstanceRequest{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *AdoptResourceInstanceRequest) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *AdoptResourceInstanceRequest) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *AdoptResourceInstanceRequest) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetCustomNetworkId returns the CustomNetworkId field value if set, zero value otherwise.
func (o *AdoptResourceInstanceRequest) GetCustomNetworkId() string {
	if o == nil || IsNil(o.CustomNetworkId) {
		var ret string
		return ret
	}
	return *o.CustomNetworkId
}

// GetCustomNetworkIdOk returns a tuple with the CustomNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetCustomNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomNetworkId) {
		return nil, false
	}
	return o.CustomNetworkId, true
}

// HasCustomNetworkId returns a boolean if a field has been set.
func (o *AdoptResourceInstanceRequest) HasCustomNetworkId() bool {
	if o != nil && !IsNil(o.CustomNetworkId) {
		return true
	}

	return false
}

// SetCustomNetworkId gets a reference to the given string and assigns it to the CustomNetworkId field.
func (o *AdoptResourceInstanceRequest) SetCustomNetworkId(v string) {
	o.CustomNetworkId = &v
}

// GetExternalPayerId returns the ExternalPayerId field value if set, zero value otherwise.
func (o *AdoptResourceInstanceRequest) GetExternalPayerId() string {
	if o == nil || IsNil(o.ExternalPayerId) {
		var ret string
		return ret
	}
	return *o.ExternalPayerId
}

// GetExternalPayerIdOk returns a tuple with the ExternalPayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetExternalPayerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPayerId) {
		return nil, false
	}
	return o.ExternalPayerId, true
}

// HasExternalPayerId returns a boolean if a field has been set.
func (o *AdoptResourceInstanceRequest) HasExternalPayerId() bool {
	if o != nil && !IsNil(o.ExternalPayerId) {
		return true
	}

	return false
}

// SetExternalPayerId gets a reference to the given string and assigns it to the ExternalPayerId field.
func (o *AdoptResourceInstanceRequest) SetExternalPayerId(v string) {
	o.ExternalPayerId = &v
}

// GetProductTierKey returns the ProductTierKey field value
func (o *AdoptResourceInstanceRequest) GetProductTierKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierKey
}

// GetProductTierKeyOk returns a tuple with the ProductTierKey field value
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetProductTierKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierKey, true
}

// SetProductTierKey sets field value
func (o *AdoptResourceInstanceRequest) SetProductTierKey(v string) {
	o.ProductTierKey = v
}

// GetProductTierVersion returns the ProductTierVersion field value if set, zero value otherwise.
func (o *AdoptResourceInstanceRequest) GetProductTierVersion() string {
	if o == nil || IsNil(o.ProductTierVersion) {
		var ret string
		return ret
	}
	return *o.ProductTierVersion
}

// GetProductTierVersionOk returns a tuple with the ProductTierVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetProductTierVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierVersion) {
		return nil, false
	}
	return o.ProductTierVersion, true
}

// HasProductTierVersion returns a boolean if a field has been set.
func (o *AdoptResourceInstanceRequest) HasProductTierVersion() bool {
	if o != nil && !IsNil(o.ProductTierVersion) {
		return true
	}

	return false
}

// SetProductTierVersion gets a reference to the given string and assigns it to the ProductTierVersion field.
func (o *AdoptResourceInstanceRequest) SetProductTierVersion(v string) {
	o.ProductTierVersion = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AdoptResourceInstanceRequest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AdoptResourceInstanceRequest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AdoptResourceInstanceRequest) SetRegion(v string) {
	o.Region = &v
}

// GetResourceConfiguration returns the ResourceConfiguration field value if set, zero value otherwise.
func (o *AdoptResourceInstanceRequest) GetResourceConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.ResourceConfiguration) {
		var ret map[string]interface{}
		return ret
	}
	return o.ResourceConfiguration
}

// GetResourceConfigurationOk returns a tuple with the ResourceConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetResourceConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResourceConfiguration) {
		return map[string]interface{}{}, false
	}
	return o.ResourceConfiguration, true
}

// HasResourceConfiguration returns a boolean if a field has been set.
func (o *AdoptResourceInstanceRequest) HasResourceConfiguration() bool {
	if o != nil && !IsNil(o.ResourceConfiguration) {
		return true
	}

	return false
}

// SetResourceConfiguration gets a reference to the given map[string]interface{} and assigns it to the ResourceConfiguration field.
func (o *AdoptResourceInstanceRequest) SetResourceConfiguration(v map[string]interface{}) {
	o.ResourceConfiguration = v
}

// GetServiceAPIVersion returns the ServiceAPIVersion field value
func (o *AdoptResourceInstanceRequest) GetServiceAPIVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAPIVersion
}

// GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field value
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetServiceAPIVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAPIVersion, true
}

// SetServiceAPIVersion sets field value
func (o *AdoptResourceInstanceRequest) SetServiceAPIVersion(v string) {
	o.ServiceAPIVersion = v
}

// GetServiceEnvironmentKey returns the ServiceEnvironmentKey field value
func (o *AdoptResourceInstanceRequest) GetServiceEnvironmentKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentKey
}

// GetServiceEnvironmentKeyOk returns a tuple with the ServiceEnvironmentKey field value
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetServiceEnvironmentKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentKey, true
}

// SetServiceEnvironmentKey sets field value
func (o *AdoptResourceInstanceRequest) SetServiceEnvironmentKey(v string) {
	o.ServiceEnvironmentKey = v
}

// GetServiceKey returns the ServiceKey field value
func (o *AdoptResourceInstanceRequest) GetServiceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetServiceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceKey, true
}

// SetServiceKey sets field value
func (o *AdoptResourceInstanceRequest) SetServiceKey(v string) {
	o.ServiceKey = v
}

// GetServiceModelKey returns the ServiceModelKey field value
func (o *AdoptResourceInstanceRequest) GetServiceModelKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceModelKey
}

// GetServiceModelKeyOk returns a tuple with the ServiceModelKey field value
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetServiceModelKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceModelKey, true
}

// SetServiceModelKey sets field value
func (o *AdoptResourceInstanceRequest) SetServiceModelKey(v string) {
	o.ServiceModelKey = v
}

// GetServiceProviderId returns the ServiceProviderId field value
func (o *AdoptResourceInstanceRequest) GetServiceProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceProviderId
}

// GetServiceProviderIdOk returns a tuple with the ServiceProviderId field value
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetServiceProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceProviderId, true
}

// SetServiceProviderId sets field value
func (o *AdoptResourceInstanceRequest) SetServiceProviderId(v string) {
	o.ServiceProviderId = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *AdoptResourceInstanceRequest) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *AdoptResourceInstanceRequest) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *AdoptResourceInstanceRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetToken returns the Token field value
func (o *AdoptResourceInstanceRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AdoptResourceInstanceRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AdoptResourceInstanceRequest) SetToken(v string) {
	o.Token = v
}

func (o AdoptResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdoptResourceInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if !IsNil(o.CustomNetworkId) {
		toSerialize["custom_network_id"] = o.CustomNetworkId
	}
	if !IsNil(o.ExternalPayerId) {
		toSerialize["externalPayerId"] = o.ExternalPayerId
	}
	toSerialize["productTierKey"] = o.ProductTierKey
	if !IsNil(o.ProductTierVersion) {
		toSerialize["productTierVersion"] = o.ProductTierVersion
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.ResourceConfiguration) {
		toSerialize["resourceConfiguration"] = o.ResourceConfiguration
	}
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

func (o *AdoptResourceInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productTierKey",
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

	varAdoptResourceInstanceRequest := _AdoptResourceInstanceRequest{}

	err = json.Unmarshal(data, &varAdoptResourceInstanceRequest)

	if err != nil {
		return err
	}

	*o = AdoptResourceInstanceRequest(varAdoptResourceInstanceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloud_provider")
		delete(additionalProperties, "custom_network_id")
		delete(additionalProperties, "externalPayerId")
		delete(additionalProperties, "productTierKey")
		delete(additionalProperties, "productTierVersion")
		delete(additionalProperties, "region")
		delete(additionalProperties, "resourceConfiguration")
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

type NullableAdoptResourceInstanceRequest struct {
	value *AdoptResourceInstanceRequest
	isSet bool
}

func (v NullableAdoptResourceInstanceRequest) Get() *AdoptResourceInstanceRequest {
	return v.value
}

func (v *NullableAdoptResourceInstanceRequest) Set(val *AdoptResourceInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdoptResourceInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdoptResourceInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdoptResourceInstanceRequest(val *AdoptResourceInstanceRequest) *NullableAdoptResourceInstanceRequest {
	return &NullableAdoptResourceInstanceRequest{value: val, isSet: true}
}

func (v NullableAdoptResourceInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdoptResourceInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


