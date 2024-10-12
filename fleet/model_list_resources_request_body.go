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

// checks if the ListResourcesRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResourcesRequestBody{}

// ListResourcesRequestBody struct for ListResourcesRequestBody
type ListResourcesRequestBody struct {
	// Token to use for the next request
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The number of resources to return per page.
	PageSize *int64 `json:"pageSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListResourcesRequestBody ListResourcesRequestBody

// NewListResourcesRequestBody instantiates a new ListResourcesRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResourcesRequestBody() *ListResourcesRequestBody {
	this := ListResourcesRequestBody{}
	return &this
}

// NewListResourcesRequestBodyWithDefaults instantiates a new ListResourcesRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResourcesRequestBodyWithDefaults() *ListResourcesRequestBody {
	this := ListResourcesRequestBody{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListResourcesRequestBody) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResourcesRequestBody) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListResourcesRequestBody) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListResourcesRequestBody) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListResourcesRequestBody) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResourcesRequestBody) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListResourcesRequestBody) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ListResourcesRequestBody) SetPageSize(v int64) {
	o.PageSize = &v
}

func (o ListResourcesRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResourcesRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListResourcesRequestBody) UnmarshalJSON(data []byte) (err error) {
	varListResourcesRequestBody := _ListResourcesRequestBody{}

	err = json.Unmarshal(data, &varListResourcesRequestBody)

	if err != nil {
		return err
	}

	*o = ListResourcesRequestBody(varListResourcesRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "pageSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListResourcesRequestBody struct {
	value *ListResourcesRequestBody
	isSet bool
}

func (v NullableListResourcesRequestBody) Get() *ListResourcesRequestBody {
	return v.value
}

func (v *NullableListResourcesRequestBody) Set(val *ListResourcesRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableListResourcesRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableListResourcesRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResourcesRequestBody(val *ListResourcesRequestBody) *NullableListResourcesRequestBody {
	return &NullableListResourcesRequestBody{value: val, isSet: true}
}

func (v NullableListResourcesRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResourcesRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


