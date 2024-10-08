/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserSearchRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSearchRecord{}

// UserSearchRecord struct for UserSearchRecord
type UserSearchRecord struct {
	// The user email.
	Email string `json:"email"`
	// Is the user an external user.
	External bool `json:"external"`
	// The user ID.
	Id string `json:"id"`
	// The user name.
	Name string `json:"name"`
	// The organization name.
	OrgName string `json:"orgName"`
}

type _UserSearchRecord UserSearchRecord

// NewUserSearchRecord instantiates a new UserSearchRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSearchRecord(email string, external bool, id string, name string, orgName string) *UserSearchRecord {
	this := UserSearchRecord{}
	this.Email = email
	this.External = external
	this.Id = id
	this.Name = name
	this.OrgName = orgName
	return &this
}

// NewUserSearchRecordWithDefaults instantiates a new UserSearchRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSearchRecordWithDefaults() *UserSearchRecord {
	this := UserSearchRecord{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserSearchRecord) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserSearchRecord) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserSearchRecord) SetEmail(v string) {
	o.Email = v
}

// GetExternal returns the External field value
func (o *UserSearchRecord) GetExternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.External
}

// GetExternalOk returns a tuple with the External field value
// and a boolean to check if the value has been set.
func (o *UserSearchRecord) GetExternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.External, true
}

// SetExternal sets field value
func (o *UserSearchRecord) SetExternal(v bool) {
	o.External = v
}

// GetId returns the Id field value
func (o *UserSearchRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSearchRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSearchRecord) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *UserSearchRecord) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserSearchRecord) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserSearchRecord) SetName(v string) {
	o.Name = v
}

// GetOrgName returns the OrgName field value
func (o *UserSearchRecord) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *UserSearchRecord) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *UserSearchRecord) SetOrgName(v string) {
	o.OrgName = v
}

func (o UserSearchRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSearchRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["external"] = o.External
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["orgName"] = o.OrgName
	return toSerialize, nil
}

func (o *UserSearchRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"external",
		"id",
		"name",
		"orgName",
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

	varUserSearchRecord := _UserSearchRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSearchRecord)

	if err != nil {
		return err
	}

	*o = UserSearchRecord(varUserSearchRecord)

	return err
}

type NullableUserSearchRecord struct {
	value *UserSearchRecord
	isSet bool
}

func (v NullableUserSearchRecord) Get() *UserSearchRecord {
	return v.value
}

func (v *NullableUserSearchRecord) Set(val *UserSearchRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSearchRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSearchRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSearchRecord(val *UserSearchRecord) *NullableUserSearchRecord {
	return &NullableUserSearchRecord{value: val, isSet: true}
}

func (v NullableUserSearchRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSearchRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


