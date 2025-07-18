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

// checks if the EnableBillingProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableBillingProviderRequest{}

// EnableBillingProviderRequest struct for EnableBillingProviderRequest
type EnableBillingProviderRequest struct {
	// The URL to the balance due page
	BalanceDueLink *string `json:"balanceDueLink,omitempty"`
	// The billing provider type
	BillingProviderType string `json:"billingProviderType"`
	// The URL of the logo for the billing provider. Only present when the billing provider is BRING_YOUR_OWN
	LogoURL *string `json:"logoURL,omitempty"`
	// A custom name for the billing provider. Only present when the billing provider is BRING_YOUR_OWN
	Name *string `json:"name,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _EnableBillingProviderRequest EnableBillingProviderRequest

// NewEnableBillingProviderRequest instantiates a new EnableBillingProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableBillingProviderRequest(billingProviderType string, token string) *EnableBillingProviderRequest {
	this := EnableBillingProviderRequest{}
	this.BillingProviderType = billingProviderType
	this.Token = token
	return &this
}

// NewEnableBillingProviderRequestWithDefaults instantiates a new EnableBillingProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableBillingProviderRequestWithDefaults() *EnableBillingProviderRequest {
	this := EnableBillingProviderRequest{}
	return &this
}

// GetBalanceDueLink returns the BalanceDueLink field value if set, zero value otherwise.
func (o *EnableBillingProviderRequest) GetBalanceDueLink() string {
	if o == nil || IsNil(o.BalanceDueLink) {
		var ret string
		return ret
	}
	return *o.BalanceDueLink
}

// GetBalanceDueLinkOk returns a tuple with the BalanceDueLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableBillingProviderRequest) GetBalanceDueLinkOk() (*string, bool) {
	if o == nil || IsNil(o.BalanceDueLink) {
		return nil, false
	}
	return o.BalanceDueLink, true
}

// HasBalanceDueLink returns a boolean if a field has been set.
func (o *EnableBillingProviderRequest) HasBalanceDueLink() bool {
	if o != nil && !IsNil(o.BalanceDueLink) {
		return true
	}

	return false
}

// SetBalanceDueLink gets a reference to the given string and assigns it to the BalanceDueLink field.
func (o *EnableBillingProviderRequest) SetBalanceDueLink(v string) {
	o.BalanceDueLink = &v
}

// GetBillingProviderType returns the BillingProviderType field value
func (o *EnableBillingProviderRequest) GetBillingProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingProviderType
}

// GetBillingProviderTypeOk returns a tuple with the BillingProviderType field value
// and a boolean to check if the value has been set.
func (o *EnableBillingProviderRequest) GetBillingProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingProviderType, true
}

// SetBillingProviderType sets field value
func (o *EnableBillingProviderRequest) SetBillingProviderType(v string) {
	o.BillingProviderType = v
}

// GetLogoURL returns the LogoURL field value if set, zero value otherwise.
func (o *EnableBillingProviderRequest) GetLogoURL() string {
	if o == nil || IsNil(o.LogoURL) {
		var ret string
		return ret
	}
	return *o.LogoURL
}

// GetLogoURLOk returns a tuple with the LogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableBillingProviderRequest) GetLogoURLOk() (*string, bool) {
	if o == nil || IsNil(o.LogoURL) {
		return nil, false
	}
	return o.LogoURL, true
}

// HasLogoURL returns a boolean if a field has been set.
func (o *EnableBillingProviderRequest) HasLogoURL() bool {
	if o != nil && !IsNil(o.LogoURL) {
		return true
	}

	return false
}

// SetLogoURL gets a reference to the given string and assigns it to the LogoURL field.
func (o *EnableBillingProviderRequest) SetLogoURL(v string) {
	o.LogoURL = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnableBillingProviderRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableBillingProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnableBillingProviderRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnableBillingProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetToken returns the Token field value
func (o *EnableBillingProviderRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *EnableBillingProviderRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *EnableBillingProviderRequest) SetToken(v string) {
	o.Token = v
}

func (o EnableBillingProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableBillingProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BalanceDueLink) {
		toSerialize["balanceDueLink"] = o.BalanceDueLink
	}
	toSerialize["billingProviderType"] = o.BillingProviderType
	if !IsNil(o.LogoURL) {
		toSerialize["logoURL"] = o.LogoURL
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnableBillingProviderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billingProviderType",
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

	varEnableBillingProviderRequest := _EnableBillingProviderRequest{}

	err = json.Unmarshal(data, &varEnableBillingProviderRequest)

	if err != nil {
		return err
	}

	*o = EnableBillingProviderRequest(varEnableBillingProviderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "balanceDueLink")
		delete(additionalProperties, "billingProviderType")
		delete(additionalProperties, "logoURL")
		delete(additionalProperties, "name")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnableBillingProviderRequest struct {
	value *EnableBillingProviderRequest
	isSet bool
}

func (v NullableEnableBillingProviderRequest) Get() *EnableBillingProviderRequest {
	return v.value
}

func (v *NullableEnableBillingProviderRequest) Set(val *EnableBillingProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableBillingProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableBillingProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableBillingProviderRequest(val *EnableBillingProviderRequest) *NullableEnableBillingProviderRequest {
	return &NullableEnableBillingProviderRequest{value: val, isSet: true}
}

func (v NullableEnableBillingProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableBillingProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


