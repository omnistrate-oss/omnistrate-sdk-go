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

// checks if the UpdateSaaSPortalRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSaaSPortalRequestBody{}

// UpdateSaaSPortalRequestBody struct for UpdateSaaSPortalRequestBody
type UpdateSaaSPortalRequestBody struct {
	// The custom domain for the SaaS portal
	CustomDomain *string `json:"customDomain,omitempty"`
	EmailConfig *SaaSPortalEmailConfig `json:"emailConfig,omitempty"`
	// The Google Analytics tag ID for the SaaS portal
	GoogleAnalyticsTagID *string `json:"googleAnalyticsTagID,omitempty"`
	ImageConfig *SaaSPortalImageConfig `json:"imageConfig,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateSaaSPortalRequestBody UpdateSaaSPortalRequestBody

// NewUpdateSaaSPortalRequestBody instantiates a new UpdateSaaSPortalRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSaaSPortalRequestBody() *UpdateSaaSPortalRequestBody {
	this := UpdateSaaSPortalRequestBody{}
	return &this
}

// NewUpdateSaaSPortalRequestBodyWithDefaults instantiates a new UpdateSaaSPortalRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSaaSPortalRequestBodyWithDefaults() *UpdateSaaSPortalRequestBody {
	this := UpdateSaaSPortalRequestBody{}
	return &this
}

// GetCustomDomain returns the CustomDomain field value if set, zero value otherwise.
func (o *UpdateSaaSPortalRequestBody) GetCustomDomain() string {
	if o == nil || IsNil(o.CustomDomain) {
		var ret string
		return ret
	}
	return *o.CustomDomain
}

// GetCustomDomainOk returns a tuple with the CustomDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSaaSPortalRequestBody) GetCustomDomainOk() (*string, bool) {
	if o == nil || IsNil(o.CustomDomain) {
		return nil, false
	}
	return o.CustomDomain, true
}

// SetCustomDomain gets a reference to the given string and assigns it to the CustomDomain field.
func (o *UpdateSaaSPortalRequestBody) SetCustomDomain(v string) {
	o.CustomDomain = &v
}

// GetEmailConfig returns the EmailConfig field value if set, zero value otherwise.
func (o *UpdateSaaSPortalRequestBody) GetEmailConfig() SaaSPortalEmailConfig {
	if o == nil || IsNil(o.EmailConfig) {
		var ret SaaSPortalEmailConfig
		return ret
	}
	return *o.EmailConfig
}

// GetEmailConfigOk returns a tuple with the EmailConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSaaSPortalRequestBody) GetEmailConfigOk() (*SaaSPortalEmailConfig, bool) {
	if o == nil || IsNil(o.EmailConfig) {
		return nil, false
	}
	return o.EmailConfig, true
}

// SetEmailConfig gets a reference to the given SaaSPortalEmailConfig and assigns it to the EmailConfig field.
func (o *UpdateSaaSPortalRequestBody) SetEmailConfig(v SaaSPortalEmailConfig) {
	o.EmailConfig = &v
}

// GetGoogleAnalyticsTagID returns the GoogleAnalyticsTagID field value if set, zero value otherwise.
func (o *UpdateSaaSPortalRequestBody) GetGoogleAnalyticsTagID() string {
	if o == nil || IsNil(o.GoogleAnalyticsTagID) {
		var ret string
		return ret
	}
	return *o.GoogleAnalyticsTagID
}

// GetGoogleAnalyticsTagIDOk returns a tuple with the GoogleAnalyticsTagID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSaaSPortalRequestBody) GetGoogleAnalyticsTagIDOk() (*string, bool) {
	if o == nil || IsNil(o.GoogleAnalyticsTagID) {
		return nil, false
	}
	return o.GoogleAnalyticsTagID, true
}

// SetGoogleAnalyticsTagID gets a reference to the given string and assigns it to the GoogleAnalyticsTagID field.
func (o *UpdateSaaSPortalRequestBody) SetGoogleAnalyticsTagID(v string) {
	o.GoogleAnalyticsTagID = &v
}

// GetImageConfig returns the ImageConfig field value if set, zero value otherwise.
func (o *UpdateSaaSPortalRequestBody) GetImageConfig() SaaSPortalImageConfig {
	if o == nil || IsNil(o.ImageConfig) {
		var ret SaaSPortalImageConfig
		return ret
	}
	return *o.ImageConfig
}

// GetImageConfigOk returns a tuple with the ImageConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSaaSPortalRequestBody) GetImageConfigOk() (*SaaSPortalImageConfig, bool) {
	if o == nil || IsNil(o.ImageConfig) {
		return nil, false
	}
	return o.ImageConfig, true
}

// SetImageConfig gets a reference to the given SaaSPortalImageConfig and assigns it to the ImageConfig field.
func (o *UpdateSaaSPortalRequestBody) SetImageConfig(v SaaSPortalImageConfig) {
	o.ImageConfig = &v
}

func (o UpdateSaaSPortalRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSaaSPortalRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomDomain) {
		toSerialize["customDomain"] = o.CustomDomain
	}
	if !IsNil(o.EmailConfig) {
		toSerialize["emailConfig"] = o.EmailConfig
	}
	if !IsNil(o.GoogleAnalyticsTagID) {
		toSerialize["googleAnalyticsTagID"] = o.GoogleAnalyticsTagID
	}
	if !IsNil(o.ImageConfig) {
		toSerialize["imageConfig"] = o.ImageConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateSaaSPortalRequestBody) UnmarshalJSON(data []byte) (err error) {
	varUpdateSaaSPortalRequestBody := _UpdateSaaSPortalRequestBody{}

	err = json.Unmarshal(data, &varUpdateSaaSPortalRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateSaaSPortalRequestBody(varUpdateSaaSPortalRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customDomain")
		delete(additionalProperties, "emailConfig")
		delete(additionalProperties, "googleAnalyticsTagID")
		delete(additionalProperties, "imageConfig")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateSaaSPortalRequestBody struct {
	value *UpdateSaaSPortalRequestBody
	isSet bool
}

func (v NullableUpdateSaaSPortalRequestBody) Get() *UpdateSaaSPortalRequestBody {
	return v.value
}

func (v *NullableUpdateSaaSPortalRequestBody) Set(val *UpdateSaaSPortalRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSaaSPortalRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSaaSPortalRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSaaSPortalRequestBody(val *UpdateSaaSPortalRequestBody) *NullableUpdateSaaSPortalRequestBody {
	return &NullableUpdateSaaSPortalRequestBody{value: val, isSet: true}
}

func (v NullableUpdateSaaSPortalRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSaaSPortalRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


