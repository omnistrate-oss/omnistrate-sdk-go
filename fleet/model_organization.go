/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
)

// checks if the Organization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organization{}

// Organization struct for Organization
type Organization struct {
	// Email address to reach for support-related queries.
	Email *string `json:"email,omitempty"`
	// ID of an Org
	OrgId *string `json:"orgId,omitempty"`
	// The organization name.
	OrgName *string `json:"orgName,omitempty"`
	// The organization's URL.
	OrgUrl *string `json:"orgUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Organization Organization

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization() *Organization {
	this := Organization{}
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Organization) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Organization) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Organization) SetEmail(v string) {
	o.Email = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Organization) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Organization) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Organization) SetOrgId(v string) {
	o.OrgId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Organization) GetOrgName() string {
	if o == nil || IsNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetOrgNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrgName) {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Organization) HasOrgName() bool {
	if o != nil && !IsNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Organization) SetOrgName(v string) {
	o.OrgName = &v
}

// GetOrgUrl returns the OrgUrl field value if set, zero value otherwise.
func (o *Organization) GetOrgUrl() string {
	if o == nil || IsNil(o.OrgUrl) {
		var ret string
		return ret
	}
	return *o.OrgUrl
}

// GetOrgUrlOk returns a tuple with the OrgUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetOrgUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OrgUrl) {
		return nil, false
	}
	return o.OrgUrl, true
}

// HasOrgUrl returns a boolean if a field has been set.
func (o *Organization) HasOrgUrl() bool {
	if o != nil && !IsNil(o.OrgUrl) {
		return true
	}

	return false
}

// SetOrgUrl gets a reference to the given string and assigns it to the OrgUrl field.
func (o *Organization) SetOrgUrl(v string) {
	o.OrgUrl = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !IsNil(o.OrgName) {
		toSerialize["orgName"] = o.OrgName
	}
	if !IsNil(o.OrgUrl) {
		toSerialize["orgUrl"] = o.OrgUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Organization) UnmarshalJSON(data []byte) (err error) {
	varOrganization := _Organization{}

	err = json.Unmarshal(data, &varOrganization)

	if err != nil {
		return err
	}

	*o = Organization(varOrganization)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "orgUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


