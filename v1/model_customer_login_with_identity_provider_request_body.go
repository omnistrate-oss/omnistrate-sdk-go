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

// checks if the CustomerLoginWithIdentityProviderRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerLoginWithIdentityProviderRequestBody{}

// CustomerLoginWithIdentityProviderRequestBody struct for CustomerLoginWithIdentityProviderRequestBody
type CustomerLoginWithIdentityProviderRequestBody struct {
	// The authorization code from the Identity Provider
	AuthorizationCode string `json:"authorizationCode"`
	CompanyDescription *string `json:"companyDescription,omitempty"`
	CompanyUrl *string `json:"companyUrl,omitempty"`
	// The environment type of the portal that the customer is signing in to
	EnvironmentType *string `json:"environmentType,omitempty"`
	// The name of the Identity Provider
	IdentityProviderName string `json:"identityProviderName"`
	// Email address that the user was invited with
	InvitedEmail *string `json:"invitedEmail,omitempty"`
	LegalCompanyName *string `json:"legalCompanyName,omitempty"`
	// The redirect URI used to get the authorization code
	RedirectUri *string `json:"redirectUri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerLoginWithIdentityProviderRequestBody CustomerLoginWithIdentityProviderRequestBody

// NewCustomerLoginWithIdentityProviderRequestBody instantiates a new CustomerLoginWithIdentityProviderRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerLoginWithIdentityProviderRequestBody(authorizationCode string, identityProviderName string) *CustomerLoginWithIdentityProviderRequestBody {
	this := CustomerLoginWithIdentityProviderRequestBody{}
	this.AuthorizationCode = authorizationCode
	this.IdentityProviderName = identityProviderName
	return &this
}

// NewCustomerLoginWithIdentityProviderRequestBodyWithDefaults instantiates a new CustomerLoginWithIdentityProviderRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerLoginWithIdentityProviderRequestBodyWithDefaults() *CustomerLoginWithIdentityProviderRequestBody {
	this := CustomerLoginWithIdentityProviderRequestBody{}
	return &this
}

// GetAuthorizationCode returns the AuthorizationCode field value
func (o *CustomerLoginWithIdentityProviderRequestBody) GetAuthorizationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value
// and a boolean to check if the value has been set.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationCode, true
}

// SetAuthorizationCode sets field value
func (o *CustomerLoginWithIdentityProviderRequestBody) SetAuthorizationCode(v string) {
	o.AuthorizationCode = v
}

// GetCompanyDescription returns the CompanyDescription field value if set, zero value otherwise.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetCompanyDescription() string {
	if o == nil || IsNil(o.CompanyDescription) {
		var ret string
		return ret
	}
	return *o.CompanyDescription
}

// GetCompanyDescriptionOk returns a tuple with the CompanyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetCompanyDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyDescription) {
		return nil, false
	}
	return o.CompanyDescription, true
}

// SetCompanyDescription gets a reference to the given string and assigns it to the CompanyDescription field.
func (o *CustomerLoginWithIdentityProviderRequestBody) SetCompanyDescription(v string) {
	o.CompanyDescription = &v
}

// GetCompanyUrl returns the CompanyUrl field value if set, zero value otherwise.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetCompanyUrl() string {
	if o == nil || IsNil(o.CompanyUrl) {
		var ret string
		return ret
	}
	return *o.CompanyUrl
}

// GetCompanyUrlOk returns a tuple with the CompanyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetCompanyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyUrl) {
		return nil, false
	}
	return o.CompanyUrl, true
}

// SetCompanyUrl gets a reference to the given string and assigns it to the CompanyUrl field.
func (o *CustomerLoginWithIdentityProviderRequestBody) SetCompanyUrl(v string) {
	o.CompanyUrl = &v
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetEnvironmentType() string {
	if o == nil || IsNil(o.EnvironmentType) {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentType) {
		return nil, false
	}
	return o.EnvironmentType, true
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *CustomerLoginWithIdentityProviderRequestBody) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetIdentityProviderName returns the IdentityProviderName field value
func (o *CustomerLoginWithIdentityProviderRequestBody) GetIdentityProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityProviderName
}

// GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field value
// and a boolean to check if the value has been set.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetIdentityProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProviderName, true
}

// SetIdentityProviderName sets field value
func (o *CustomerLoginWithIdentityProviderRequestBody) SetIdentityProviderName(v string) {
	o.IdentityProviderName = v
}

// GetInvitedEmail returns the InvitedEmail field value if set, zero value otherwise.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetInvitedEmail() string {
	if o == nil || IsNil(o.InvitedEmail) {
		var ret string
		return ret
	}
	return *o.InvitedEmail
}

// GetInvitedEmailOk returns a tuple with the InvitedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetInvitedEmailOk() (*string, bool) {
	if o == nil || IsNil(o.InvitedEmail) {
		return nil, false
	}
	return o.InvitedEmail, true
}

// SetInvitedEmail gets a reference to the given string and assigns it to the InvitedEmail field.
func (o *CustomerLoginWithIdentityProviderRequestBody) SetInvitedEmail(v string) {
	o.InvitedEmail = &v
}

// GetLegalCompanyName returns the LegalCompanyName field value if set, zero value otherwise.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetLegalCompanyName() string {
	if o == nil || IsNil(o.LegalCompanyName) {
		var ret string
		return ret
	}
	return *o.LegalCompanyName
}

// GetLegalCompanyNameOk returns a tuple with the LegalCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetLegalCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalCompanyName) {
		return nil, false
	}
	return o.LegalCompanyName, true
}

// SetLegalCompanyName gets a reference to the given string and assigns it to the LegalCompanyName field.
func (o *CustomerLoginWithIdentityProviderRequestBody) SetLegalCompanyName(v string) {
	o.LegalCompanyName = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetRedirectUri() string {
	if o == nil || IsNil(o.RedirectUri) {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLoginWithIdentityProviderRequestBody) GetRedirectUriOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUri) {
		return nil, false
	}
	return o.RedirectUri, true
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *CustomerLoginWithIdentityProviderRequestBody) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

func (o CustomerLoginWithIdentityProviderRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerLoginWithIdentityProviderRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorizationCode"] = o.AuthorizationCode
	if !IsNil(o.CompanyDescription) {
		toSerialize["companyDescription"] = o.CompanyDescription
	}
	if !IsNil(o.CompanyUrl) {
		toSerialize["companyUrl"] = o.CompanyUrl
	}
	if !IsNil(o.EnvironmentType) {
		toSerialize["environmentType"] = o.EnvironmentType
	}
	toSerialize["identityProviderName"] = o.IdentityProviderName
	if !IsNil(o.InvitedEmail) {
		toSerialize["invitedEmail"] = o.InvitedEmail
	}
	if !IsNil(o.LegalCompanyName) {
		toSerialize["legalCompanyName"] = o.LegalCompanyName
	}
	if !IsNil(o.RedirectUri) {
		toSerialize["redirectUri"] = o.RedirectUri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerLoginWithIdentityProviderRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorizationCode",
		"identityProviderName",
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

	varCustomerLoginWithIdentityProviderRequestBody := _CustomerLoginWithIdentityProviderRequestBody{}

	err = json.Unmarshal(data, &varCustomerLoginWithIdentityProviderRequestBody)

	if err != nil {
		return err
	}

	*o = CustomerLoginWithIdentityProviderRequestBody(varCustomerLoginWithIdentityProviderRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorizationCode")
		delete(additionalProperties, "companyDescription")
		delete(additionalProperties, "companyUrl")
		delete(additionalProperties, "environmentType")
		delete(additionalProperties, "identityProviderName")
		delete(additionalProperties, "invitedEmail")
		delete(additionalProperties, "legalCompanyName")
		delete(additionalProperties, "redirectUri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerLoginWithIdentityProviderRequestBody struct {
	value *CustomerLoginWithIdentityProviderRequestBody
	isSet bool
}

func (v NullableCustomerLoginWithIdentityProviderRequestBody) Get() *CustomerLoginWithIdentityProviderRequestBody {
	return v.value
}

func (v *NullableCustomerLoginWithIdentityProviderRequestBody) Set(val *CustomerLoginWithIdentityProviderRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerLoginWithIdentityProviderRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerLoginWithIdentityProviderRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerLoginWithIdentityProviderRequestBody(val *CustomerLoginWithIdentityProviderRequestBody) *NullableCustomerLoginWithIdentityProviderRequestBody {
	return &NullableCustomerLoginWithIdentityProviderRequestBody{value: val, isSet: true}
}

func (v NullableCustomerLoginWithIdentityProviderRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerLoginWithIdentityProviderRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


