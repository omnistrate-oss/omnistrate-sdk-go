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

// checks if the UpdateUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserRequest{}

// UpdateUserRequest struct for UpdateUserRequest
type UpdateUserRequest struct {
	Address *Address `json:"address,omitempty"`
	// The User ID
	Id string `json:"id"`
	// The name of the user
	Name *string `json:"name,omitempty"`
	// The cookie policy for the org that this user owns in an HTML format
	OrgCookiePolicy *string `json:"orgCookiePolicy,omitempty"`
	// The description of the org that this user owns
	OrgDescription *string `json:"orgDescription,omitempty"`
	// The favicon of the org that this user owns
	OrgFavIconURL *string `json:"orgFavIconURL,omitempty"`
	// The logo of the org that this user owns
	OrgLogoURL *string `json:"orgLogoURL,omitempty"`
	// The org name that this user owns
	OrgName *string `json:"orgName,omitempty"`
	// The privacy policy for the org that this user owns in an HTML format
	OrgPrivacyPolicy *string `json:"orgPrivacyPolicy,omitempty"`
	// The support email of the org that this user owns
	OrgSupportEmail *string `json:"orgSupportEmail,omitempty"`
	// The terms of use for the org that this user owns in an HTML format
	OrgTermsOfUse *string `json:"orgTermsOfUse,omitempty"`
	// The url of the org that this user owns
	OrgURL *string `json:"orgURL,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _UpdateUserRequest UpdateUserRequest

// NewUpdateUserRequest instantiates a new UpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequest(id string, token string) *UpdateUserRequest {
	this := UpdateUserRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestWithDefaults() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *UpdateUserRequest) SetAddress(v Address) {
	o.Address = &v
}

// GetId returns the Id field value
func (o *UpdateUserRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateUserRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateUserRequest) SetName(v string) {
	o.Name = &v
}

// GetOrgCookiePolicy returns the OrgCookiePolicy field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetOrgCookiePolicy() string {
	if o == nil || IsNil(o.OrgCookiePolicy) {
		var ret string
		return ret
	}
	return *o.OrgCookiePolicy
}

// GetOrgCookiePolicyOk returns a tuple with the OrgCookiePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetOrgCookiePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.OrgCookiePolicy) {
		return nil, false
	}
	return o.OrgCookiePolicy, true
}

// HasOrgCookiePolicy returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasOrgCookiePolicy() bool {
	if o != nil && !IsNil(o.OrgCookiePolicy) {
		return true
	}

	return false
}

// SetOrgCookiePolicy gets a reference to the given string and assigns it to the OrgCookiePolicy field.
func (o *UpdateUserRequest) SetOrgCookiePolicy(v string) {
	o.OrgCookiePolicy = &v
}

// GetOrgDescription returns the OrgDescription field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetOrgDescription() string {
	if o == nil || IsNil(o.OrgDescription) {
		var ret string
		return ret
	}
	return *o.OrgDescription
}

// GetOrgDescriptionOk returns a tuple with the OrgDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetOrgDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.OrgDescription) {
		return nil, false
	}
	return o.OrgDescription, true
}

// HasOrgDescription returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasOrgDescription() bool {
	if o != nil && !IsNil(o.OrgDescription) {
		return true
	}

	return false
}

// SetOrgDescription gets a reference to the given string and assigns it to the OrgDescription field.
func (o *UpdateUserRequest) SetOrgDescription(v string) {
	o.OrgDescription = &v
}

// GetOrgFavIconURL returns the OrgFavIconURL field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetOrgFavIconURL() string {
	if o == nil || IsNil(o.OrgFavIconURL) {
		var ret string
		return ret
	}
	return *o.OrgFavIconURL
}

// GetOrgFavIconURLOk returns a tuple with the OrgFavIconURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetOrgFavIconURLOk() (*string, bool) {
	if o == nil || IsNil(o.OrgFavIconURL) {
		return nil, false
	}
	return o.OrgFavIconURL, true
}

// HasOrgFavIconURL returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasOrgFavIconURL() bool {
	if o != nil && !IsNil(o.OrgFavIconURL) {
		return true
	}

	return false
}

// SetOrgFavIconURL gets a reference to the given string and assigns it to the OrgFavIconURL field.
func (o *UpdateUserRequest) SetOrgFavIconURL(v string) {
	o.OrgFavIconURL = &v
}

// GetOrgLogoURL returns the OrgLogoURL field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetOrgLogoURL() string {
	if o == nil || IsNil(o.OrgLogoURL) {
		var ret string
		return ret
	}
	return *o.OrgLogoURL
}

// GetOrgLogoURLOk returns a tuple with the OrgLogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetOrgLogoURLOk() (*string, bool) {
	if o == nil || IsNil(o.OrgLogoURL) {
		return nil, false
	}
	return o.OrgLogoURL, true
}

// HasOrgLogoURL returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasOrgLogoURL() bool {
	if o != nil && !IsNil(o.OrgLogoURL) {
		return true
	}

	return false
}

// SetOrgLogoURL gets a reference to the given string and assigns it to the OrgLogoURL field.
func (o *UpdateUserRequest) SetOrgLogoURL(v string) {
	o.OrgLogoURL = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetOrgName() string {
	if o == nil || IsNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetOrgNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrgName) {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasOrgName() bool {
	if o != nil && !IsNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UpdateUserRequest) SetOrgName(v string) {
	o.OrgName = &v
}

// GetOrgPrivacyPolicy returns the OrgPrivacyPolicy field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetOrgPrivacyPolicy() string {
	if o == nil || IsNil(o.OrgPrivacyPolicy) {
		var ret string
		return ret
	}
	return *o.OrgPrivacyPolicy
}

// GetOrgPrivacyPolicyOk returns a tuple with the OrgPrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetOrgPrivacyPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.OrgPrivacyPolicy) {
		return nil, false
	}
	return o.OrgPrivacyPolicy, true
}

// HasOrgPrivacyPolicy returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasOrgPrivacyPolicy() bool {
	if o != nil && !IsNil(o.OrgPrivacyPolicy) {
		return true
	}

	return false
}

// SetOrgPrivacyPolicy gets a reference to the given string and assigns it to the OrgPrivacyPolicy field.
func (o *UpdateUserRequest) SetOrgPrivacyPolicy(v string) {
	o.OrgPrivacyPolicy = &v
}

// GetOrgSupportEmail returns the OrgSupportEmail field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetOrgSupportEmail() string {
	if o == nil || IsNil(o.OrgSupportEmail) {
		var ret string
		return ret
	}
	return *o.OrgSupportEmail
}

// GetOrgSupportEmailOk returns a tuple with the OrgSupportEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetOrgSupportEmailOk() (*string, bool) {
	if o == nil || IsNil(o.OrgSupportEmail) {
		return nil, false
	}
	return o.OrgSupportEmail, true
}

// HasOrgSupportEmail returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasOrgSupportEmail() bool {
	if o != nil && !IsNil(o.OrgSupportEmail) {
		return true
	}

	return false
}

// SetOrgSupportEmail gets a reference to the given string and assigns it to the OrgSupportEmail field.
func (o *UpdateUserRequest) SetOrgSupportEmail(v string) {
	o.OrgSupportEmail = &v
}

// GetOrgTermsOfUse returns the OrgTermsOfUse field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetOrgTermsOfUse() string {
	if o == nil || IsNil(o.OrgTermsOfUse) {
		var ret string
		return ret
	}
	return *o.OrgTermsOfUse
}

// GetOrgTermsOfUseOk returns a tuple with the OrgTermsOfUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetOrgTermsOfUseOk() (*string, bool) {
	if o == nil || IsNil(o.OrgTermsOfUse) {
		return nil, false
	}
	return o.OrgTermsOfUse, true
}

// HasOrgTermsOfUse returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasOrgTermsOfUse() bool {
	if o != nil && !IsNil(o.OrgTermsOfUse) {
		return true
	}

	return false
}

// SetOrgTermsOfUse gets a reference to the given string and assigns it to the OrgTermsOfUse field.
func (o *UpdateUserRequest) SetOrgTermsOfUse(v string) {
	o.OrgTermsOfUse = &v
}

// GetOrgURL returns the OrgURL field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetOrgURL() string {
	if o == nil || IsNil(o.OrgURL) {
		var ret string
		return ret
	}
	return *o.OrgURL
}

// GetOrgURLOk returns a tuple with the OrgURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetOrgURLOk() (*string, bool) {
	if o == nil || IsNil(o.OrgURL) {
		return nil, false
	}
	return o.OrgURL, true
}

// HasOrgURL returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasOrgURL() bool {
	if o != nil && !IsNil(o.OrgURL) {
		return true
	}

	return false
}

// SetOrgURL gets a reference to the given string and assigns it to the OrgURL field.
func (o *UpdateUserRequest) SetOrgURL(v string) {
	o.OrgURL = &v
}

// GetToken returns the Token field value
func (o *UpdateUserRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateUserRequest) SetToken(v string) {
	o.Token = v
}

func (o UpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrgCookiePolicy) {
		toSerialize["orgCookiePolicy"] = o.OrgCookiePolicy
	}
	if !IsNil(o.OrgDescription) {
		toSerialize["orgDescription"] = o.OrgDescription
	}
	if !IsNil(o.OrgFavIconURL) {
		toSerialize["orgFavIconURL"] = o.OrgFavIconURL
	}
	if !IsNil(o.OrgLogoURL) {
		toSerialize["orgLogoURL"] = o.OrgLogoURL
	}
	if !IsNil(o.OrgName) {
		toSerialize["orgName"] = o.OrgName
	}
	if !IsNil(o.OrgPrivacyPolicy) {
		toSerialize["orgPrivacyPolicy"] = o.OrgPrivacyPolicy
	}
	if !IsNil(o.OrgSupportEmail) {
		toSerialize["orgSupportEmail"] = o.OrgSupportEmail
	}
	if !IsNil(o.OrgTermsOfUse) {
		toSerialize["orgTermsOfUse"] = o.OrgTermsOfUse
	}
	if !IsNil(o.OrgURL) {
		toSerialize["orgURL"] = o.OrgURL
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUpdateUserRequest := _UpdateUserRequest{}

	err = json.Unmarshal(data, &varUpdateUserRequest)

	if err != nil {
		return err
	}

	*o = UpdateUserRequest(varUpdateUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "orgCookiePolicy")
		delete(additionalProperties, "orgDescription")
		delete(additionalProperties, "orgFavIconURL")
		delete(additionalProperties, "orgLogoURL")
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "orgPrivacyPolicy")
		delete(additionalProperties, "orgSupportEmail")
		delete(additionalProperties, "orgTermsOfUse")
		delete(additionalProperties, "orgURL")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateUserRequest struct {
	value *UpdateUserRequest
	isSet bool
}

func (v NullableUpdateUserRequest) Get() *UpdateUserRequest {
	return v.value
}

func (v *NullableUpdateUserRequest) Set(val *UpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequest(val *UpdateUserRequest) *NullableUpdateUserRequest {
	return &NullableUpdateUserRequest{value: val, isSet: true}
}

func (v NullableUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


