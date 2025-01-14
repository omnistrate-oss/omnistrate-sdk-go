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

// checks if the ListEventsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEventsResult{}

// ListEventsResult struct for ListEventsResult
type ListEventsResult struct {
	// The list of events
	Events []DescribeEventResult `json:"events,omitempty"`
	// The list of event IDs
	Ids []string `json:"ids"`
	// The next token to use for pagination
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListEventsResult ListEventsResult

// NewListEventsResult instantiates a new ListEventsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEventsResult(ids []string) *ListEventsResult {
	this := ListEventsResult{}
	this.Ids = ids
	return &this
}

// NewListEventsResultWithDefaults instantiates a new ListEventsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEventsResultWithDefaults() *ListEventsResult {
	this := ListEventsResult{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ListEventsResult) GetEvents() []DescribeEventResult {
	if o == nil || IsNil(o.Events) {
		var ret []DescribeEventResult
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventsResult) GetEventsOk() ([]DescribeEventResult, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// SetEvents gets a reference to the given []DescribeEventResult and assigns it to the Events field.
func (o *ListEventsResult) SetEvents(v []DescribeEventResult) {
	o.Events = v
}

// GetIds returns the Ids field value
func (o *ListEventsResult) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *ListEventsResult) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *ListEventsResult) SetIds(v []string) {
	o.Ids = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListEventsResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventsResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListEventsResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ListEventsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEventsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	toSerialize["ids"] = o.Ids
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListEventsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ids",
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

	varListEventsResult := _ListEventsResult{}

	err = json.Unmarshal(data, &varListEventsResult)

	if err != nil {
		return err
	}

	*o = ListEventsResult(varListEventsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "events")
		delete(additionalProperties, "ids")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListEventsResult struct {
	value *ListEventsResult
	isSet bool
}

func (v NullableListEventsResult) Get() *ListEventsResult {
	return v.value
}

func (v *NullableListEventsResult) Set(val *ListEventsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListEventsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListEventsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEventsResult(val *ListEventsResult) *NullableListEventsResult {
	return &NullableListEventsResult{value: val, isSet: true}
}

func (v NullableListEventsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEventsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


