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

// checks if the UpdateIdentityProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIdentityProviderRequest{}

// UpdateIdentityProviderRequest struct for UpdateIdentityProviderRequest
type UpdateIdentityProviderRequest struct {
	// The authorization endpoint of the Identity Provider
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty"`
	// The Client ID of the Identity Provider
	ClientId *string `json:"clientId,omitempty"`
	// The Client Secret of the Identity Provider
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Whether the Identity Provider is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// The email identifiers to use for the Identity Provider
	EmailIdentifiers *string `json:"emailIdentifiers,omitempty"`
	// The type of environment for the Identity Provider
	EnvironmentType *string `json:"environmentType,omitempty"`
	// ID of an Identity Provider
	Id string `json:"id"`
	// The URL of the icon to use for the login button
	LoginButtonIconUrl *string `json:"loginButtonIconUrl,omitempty"`
	// The text to use for the login button
	LoginButtonText *string `json:"loginButtonText,omitempty"`
	// The name of the Identity Provider
	Name *string `json:"name,omitempty"`
	// The scopes to request from the Identity Provider
	Scopes *string `json:"scopes,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	// The token endpoint of the Identity Provider
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`
	// The user info endpoint of the Identity Provider
	UserInfoEndpoint *string `json:"userInfoEndpoint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIdentityProviderRequest UpdateIdentityProviderRequest

// NewUpdateIdentityProviderRequest instantiates a new UpdateIdentityProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIdentityProviderRequest(id string, token string) *UpdateIdentityProviderRequest {
	this := UpdateIdentityProviderRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewUpdateIdentityProviderRequestWithDefaults instantiates a new UpdateIdentityProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIdentityProviderRequestWithDefaults() *UpdateIdentityProviderRequest {
	this := UpdateIdentityProviderRequest{}
	return &this
}

// GetAuthorizationEndpoint returns the AuthorizationEndpoint field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetAuthorizationEndpoint() string {
	if o == nil || IsNil(o.AuthorizationEndpoint) {
		var ret string
		return ret
	}
	return *o.AuthorizationEndpoint
}

// GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetAuthorizationEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationEndpoint) {
		return nil, false
	}
	return o.AuthorizationEndpoint, true
}

// SetAuthorizationEndpoint gets a reference to the given string and assigns it to the AuthorizationEndpoint field.
func (o *UpdateIdentityProviderRequest) SetAuthorizationEndpoint(v string) {
	o.AuthorizationEndpoint = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *UpdateIdentityProviderRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *UpdateIdentityProviderRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *UpdateIdentityProviderRequest) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetEmailIdentifiers returns the EmailIdentifiers field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetEmailIdentifiers() string {
	if o == nil || IsNil(o.EmailIdentifiers) {
		var ret string
		return ret
	}
	return *o.EmailIdentifiers
}

// GetEmailIdentifiersOk returns a tuple with the EmailIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetEmailIdentifiersOk() (*string, bool) {
	if o == nil || IsNil(o.EmailIdentifiers) {
		return nil, false
	}
	return o.EmailIdentifiers, true
}

// SetEmailIdentifiers gets a reference to the given string and assigns it to the EmailIdentifiers field.
func (o *UpdateIdentityProviderRequest) SetEmailIdentifiers(v string) {
	o.EmailIdentifiers = &v
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetEnvironmentType() string {
	if o == nil || IsNil(o.EnvironmentType) {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentType) {
		return nil, false
	}
	return o.EnvironmentType, true
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *UpdateIdentityProviderRequest) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetId returns the Id field value
func (o *UpdateIdentityProviderRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateIdentityProviderRequest) SetId(v string) {
	o.Id = v
}

// GetLoginButtonIconUrl returns the LoginButtonIconUrl field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetLoginButtonIconUrl() string {
	if o == nil || IsNil(o.LoginButtonIconUrl) {
		var ret string
		return ret
	}
	return *o.LoginButtonIconUrl
}

// GetLoginButtonIconUrlOk returns a tuple with the LoginButtonIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetLoginButtonIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LoginButtonIconUrl) {
		return nil, false
	}
	return o.LoginButtonIconUrl, true
}

// SetLoginButtonIconUrl gets a reference to the given string and assigns it to the LoginButtonIconUrl field.
func (o *UpdateIdentityProviderRequest) SetLoginButtonIconUrl(v string) {
	o.LoginButtonIconUrl = &v
}

// GetLoginButtonText returns the LoginButtonText field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetLoginButtonText() string {
	if o == nil || IsNil(o.LoginButtonText) {
		var ret string
		return ret
	}
	return *o.LoginButtonText
}

// GetLoginButtonTextOk returns a tuple with the LoginButtonText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetLoginButtonTextOk() (*string, bool) {
	if o == nil || IsNil(o.LoginButtonText) {
		return nil, false
	}
	return o.LoginButtonText, true
}

// SetLoginButtonText gets a reference to the given string and assigns it to the LoginButtonText field.
func (o *UpdateIdentityProviderRequest) SetLoginButtonText(v string) {
	o.LoginButtonText = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateIdentityProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetScopes() string {
	if o == nil || IsNil(o.Scopes) {
		var ret string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetScopesOk() (*string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes gets a reference to the given string and assigns it to the Scopes field.
func (o *UpdateIdentityProviderRequest) SetScopes(v string) {
	o.Scopes = &v
}

// GetToken returns the Token field value
func (o *UpdateIdentityProviderRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateIdentityProviderRequest) SetToken(v string) {
	o.Token = v
}

// GetTokenEndpoint returns the TokenEndpoint field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetTokenEndpoint() string {
	if o == nil || IsNil(o.TokenEndpoint) {
		var ret string
		return ret
	}
	return *o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetTokenEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEndpoint) {
		return nil, false
	}
	return o.TokenEndpoint, true
}

// SetTokenEndpoint gets a reference to the given string and assigns it to the TokenEndpoint field.
func (o *UpdateIdentityProviderRequest) SetTokenEndpoint(v string) {
	o.TokenEndpoint = &v
}

// GetUserInfoEndpoint returns the UserInfoEndpoint field value if set, zero value otherwise.
func (o *UpdateIdentityProviderRequest) GetUserInfoEndpoint() string {
	if o == nil || IsNil(o.UserInfoEndpoint) {
		var ret string
		return ret
	}
	return *o.UserInfoEndpoint
}

// GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentityProviderRequest) GetUserInfoEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.UserInfoEndpoint) {
		return nil, false
	}
	return o.UserInfoEndpoint, true
}

// SetUserInfoEndpoint gets a reference to the given string and assigns it to the UserInfoEndpoint field.
func (o *UpdateIdentityProviderRequest) SetUserInfoEndpoint(v string) {
	o.UserInfoEndpoint = &v
}

func (o UpdateIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIdentityProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationEndpoint) {
		toSerialize["authorizationEndpoint"] = o.AuthorizationEndpoint
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.EmailIdentifiers) {
		toSerialize["emailIdentifiers"] = o.EmailIdentifiers
	}
	if !IsNil(o.EnvironmentType) {
		toSerialize["environmentType"] = o.EnvironmentType
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.LoginButtonIconUrl) {
		toSerialize["loginButtonIconUrl"] = o.LoginButtonIconUrl
	}
	if !IsNil(o.LoginButtonText) {
		toSerialize["loginButtonText"] = o.LoginButtonText
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.TokenEndpoint) {
		toSerialize["tokenEndpoint"] = o.TokenEndpoint
	}
	if !IsNil(o.UserInfoEndpoint) {
		toSerialize["userInfoEndpoint"] = o.UserInfoEndpoint
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIdentityProviderRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateIdentityProviderRequest := _UpdateIdentityProviderRequest{}

	err = json.Unmarshal(data, &varUpdateIdentityProviderRequest)

	if err != nil {
		return err
	}

	*o = UpdateIdentityProviderRequest(varUpdateIdentityProviderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorizationEndpoint")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "emailIdentifiers")
		delete(additionalProperties, "environmentType")
		delete(additionalProperties, "id")
		delete(additionalProperties, "loginButtonIconUrl")
		delete(additionalProperties, "loginButtonText")
		delete(additionalProperties, "name")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "token")
		delete(additionalProperties, "tokenEndpoint")
		delete(additionalProperties, "userInfoEndpoint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIdentityProviderRequest struct {
	value *UpdateIdentityProviderRequest
	isSet bool
}

func (v NullableUpdateIdentityProviderRequest) Get() *UpdateIdentityProviderRequest {
	return v.value
}

func (v *NullableUpdateIdentityProviderRequest) Set(val *UpdateIdentityProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIdentityProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIdentityProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIdentityProviderRequest(val *UpdateIdentityProviderRequest) *NullableUpdateIdentityProviderRequest {
	return &NullableUpdateIdentityProviderRequest{value: val, isSet: true}
}

func (v NullableUpdateIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIdentityProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


