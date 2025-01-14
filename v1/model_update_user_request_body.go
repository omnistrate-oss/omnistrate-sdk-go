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

// checks if the UpdateUserRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserRequestBody{}

// UpdateUserRequestBody struct for UpdateUserRequestBody
type UpdateUserRequestBody struct {
	Address *Address `json:"address,omitempty"`
	// The name of the user
	Name *string `json:"name,omitempty"`
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
	AdditionalProperties map[string]interface{}
}

type _UpdateUserRequestBody UpdateUserRequestBody

// NewUpdateUserRequestBody instantiates a new UpdateUserRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequestBody() *UpdateUserRequestBody {
	this := UpdateUserRequestBody{}
	return &this
}

// NewUpdateUserRequestBodyWithDefaults instantiates a new UpdateUserRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestBodyWithDefaults() *UpdateUserRequestBody {
	this := UpdateUserRequestBody{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateUserRequestBody) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestBody) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *UpdateUserRequestBody) SetAddress(v Address) {
	o.Address = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateUserRequestBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateUserRequestBody) SetName(v string) {
	o.Name = &v
}

// GetOrgDescription returns the OrgDescription field value if set, zero value otherwise.
func (o *UpdateUserRequestBody) GetOrgDescription() string {
	if o == nil || IsNil(o.OrgDescription) {
		var ret string
		return ret
	}
	return *o.OrgDescription
}

// GetOrgDescriptionOk returns a tuple with the OrgDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestBody) GetOrgDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.OrgDescription) {
		return nil, false
	}
	return o.OrgDescription, true
}

// SetOrgDescription gets a reference to the given string and assigns it to the OrgDescription field.
func (o *UpdateUserRequestBody) SetOrgDescription(v string) {
	o.OrgDescription = &v
}

// GetOrgFavIconURL returns the OrgFavIconURL field value if set, zero value otherwise.
func (o *UpdateUserRequestBody) GetOrgFavIconURL() string {
	if o == nil || IsNil(o.OrgFavIconURL) {
		var ret string
		return ret
	}
	return *o.OrgFavIconURL
}

// GetOrgFavIconURLOk returns a tuple with the OrgFavIconURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestBody) GetOrgFavIconURLOk() (*string, bool) {
	if o == nil || IsNil(o.OrgFavIconURL) {
		return nil, false
	}
	return o.OrgFavIconURL, true
}

// SetOrgFavIconURL gets a reference to the given string and assigns it to the OrgFavIconURL field.
func (o *UpdateUserRequestBody) SetOrgFavIconURL(v string) {
	o.OrgFavIconURL = &v
}

// GetOrgLogoURL returns the OrgLogoURL field value if set, zero value otherwise.
func (o *UpdateUserRequestBody) GetOrgLogoURL() string {
	if o == nil || IsNil(o.OrgLogoURL) {
		var ret string
		return ret
	}
	return *o.OrgLogoURL
}

// GetOrgLogoURLOk returns a tuple with the OrgLogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestBody) GetOrgLogoURLOk() (*string, bool) {
	if o == nil || IsNil(o.OrgLogoURL) {
		return nil, false
	}
	return o.OrgLogoURL, true
}

// SetOrgLogoURL gets a reference to the given string and assigns it to the OrgLogoURL field.
func (o *UpdateUserRequestBody) SetOrgLogoURL(v string) {
	o.OrgLogoURL = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UpdateUserRequestBody) GetOrgName() string {
	if o == nil || IsNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestBody) GetOrgNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrgName) {
		return nil, false
	}
	return o.OrgName, true
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UpdateUserRequestBody) SetOrgName(v string) {
	o.OrgName = &v
}

// GetOrgPrivacyPolicy returns the OrgPrivacyPolicy field value if set, zero value otherwise.
func (o *UpdateUserRequestBody) GetOrgPrivacyPolicy() string {
	if o == nil || IsNil(o.OrgPrivacyPolicy) {
		var ret string
		return ret
	}
	return *o.OrgPrivacyPolicy
}

// GetOrgPrivacyPolicyOk returns a tuple with the OrgPrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestBody) GetOrgPrivacyPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.OrgPrivacyPolicy) {
		return nil, false
	}
	return o.OrgPrivacyPolicy, true
}

// SetOrgPrivacyPolicy gets a reference to the given string and assigns it to the OrgPrivacyPolicy field.
func (o *UpdateUserRequestBody) SetOrgPrivacyPolicy(v string) {
	o.OrgPrivacyPolicy = &v
}

// GetOrgSupportEmail returns the OrgSupportEmail field value if set, zero value otherwise.
func (o *UpdateUserRequestBody) GetOrgSupportEmail() string {
	if o == nil || IsNil(o.OrgSupportEmail) {
		var ret string
		return ret
	}
	return *o.OrgSupportEmail
}

// GetOrgSupportEmailOk returns a tuple with the OrgSupportEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestBody) GetOrgSupportEmailOk() (*string, bool) {
	if o == nil || IsNil(o.OrgSupportEmail) {
		return nil, false
	}
	return o.OrgSupportEmail, true
}

// SetOrgSupportEmail gets a reference to the given string and assigns it to the OrgSupportEmail field.
func (o *UpdateUserRequestBody) SetOrgSupportEmail(v string) {
	o.OrgSupportEmail = &v
}

// GetOrgTermsOfUse returns the OrgTermsOfUse field value if set, zero value otherwise.
func (o *UpdateUserRequestBody) GetOrgTermsOfUse() string {
	if o == nil || IsNil(o.OrgTermsOfUse) {
		var ret string
		return ret
	}
	return *o.OrgTermsOfUse
}

// GetOrgTermsOfUseOk returns a tuple with the OrgTermsOfUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestBody) GetOrgTermsOfUseOk() (*string, bool) {
	if o == nil || IsNil(o.OrgTermsOfUse) {
		return nil, false
	}
	return o.OrgTermsOfUse, true
}

// SetOrgTermsOfUse gets a reference to the given string and assigns it to the OrgTermsOfUse field.
func (o *UpdateUserRequestBody) SetOrgTermsOfUse(v string) {
	o.OrgTermsOfUse = &v
}

// GetOrgURL returns the OrgURL field value if set, zero value otherwise.
func (o *UpdateUserRequestBody) GetOrgURL() string {
	if o == nil || IsNil(o.OrgURL) {
		var ret string
		return ret
	}
	return *o.OrgURL
}

// GetOrgURLOk returns a tuple with the OrgURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestBody) GetOrgURLOk() (*string, bool) {
	if o == nil || IsNil(o.OrgURL) {
		return nil, false
	}
	return o.OrgURL, true
}

// SetOrgURL gets a reference to the given string and assigns it to the OrgURL field.
func (o *UpdateUserRequestBody) SetOrgURL(v string) {
	o.OrgURL = &v
}

func (o UpdateUserRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateUserRequestBody) UnmarshalJSON(data []byte) (err error) {
	varUpdateUserRequestBody := _UpdateUserRequestBody{}

	err = json.Unmarshal(data, &varUpdateUserRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateUserRequestBody(varUpdateUserRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "orgDescription")
		delete(additionalProperties, "orgFavIconURL")
		delete(additionalProperties, "orgLogoURL")
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "orgPrivacyPolicy")
		delete(additionalProperties, "orgSupportEmail")
		delete(additionalProperties, "orgTermsOfUse")
		delete(additionalProperties, "orgURL")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateUserRequestBody struct {
	value *UpdateUserRequestBody
	isSet bool
}

func (v NullableUpdateUserRequestBody) Get() *UpdateUserRequestBody {
	return v.value
}

func (v *NullableUpdateUserRequestBody) Set(val *UpdateUserRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequestBody(val *UpdateUserRequestBody) *NullableUpdateUserRequestBody {
	return &NullableUpdateUserRequestBody{value: val, isSet: true}
}

func (v NullableUpdateUserRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


