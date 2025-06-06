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

// checks if the ListAllOrganizationsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAllOrganizationsResult{}

// ListAllOrganizationsResult struct for ListAllOrganizationsResult
type ListAllOrganizationsResult struct {
	// Token to use for the next request
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// List of organizations.
	Organizations []Organization `json:"organizations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAllOrganizationsResult ListAllOrganizationsResult

// NewListAllOrganizationsResult instantiates a new ListAllOrganizationsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllOrganizationsResult() *ListAllOrganizationsResult {
	this := ListAllOrganizationsResult{}
	return &this
}

// NewListAllOrganizationsResultWithDefaults instantiates a new ListAllOrganizationsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllOrganizationsResultWithDefaults() *ListAllOrganizationsResult {
	this := ListAllOrganizationsResult{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListAllOrganizationsResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllOrganizationsResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListAllOrganizationsResult) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListAllOrganizationsResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *ListAllOrganizationsResult) GetOrganizations() []Organization {
	if o == nil || IsNil(o.Organizations) {
		var ret []Organization
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllOrganizationsResult) GetOrganizationsOk() ([]Organization, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *ListAllOrganizationsResult) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []Organization and assigns it to the Organizations field.
func (o *ListAllOrganizationsResult) SetOrganizations(v []Organization) {
	o.Organizations = v
}

func (o ListAllOrganizationsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAllOrganizationsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAllOrganizationsResult) UnmarshalJSON(data []byte) (err error) {
	varListAllOrganizationsResult := _ListAllOrganizationsResult{}

	err = json.Unmarshal(data, &varListAllOrganizationsResult)

	if err != nil {
		return err
	}

	*o = ListAllOrganizationsResult(varListAllOrganizationsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "organizations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAllOrganizationsResult struct {
	value *ListAllOrganizationsResult
	isSet bool
}

func (v NullableListAllOrganizationsResult) Get() *ListAllOrganizationsResult {
	return v.value
}

func (v *NullableListAllOrganizationsResult) Set(val *ListAllOrganizationsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllOrganizationsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllOrganizationsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllOrganizationsResult(val *ListAllOrganizationsResult) *NullableListAllOrganizationsResult {
	return &NullableListAllOrganizationsResult{value: val, isSet: true}
}

func (v NullableListAllOrganizationsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllOrganizationsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


