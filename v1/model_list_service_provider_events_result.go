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

// checks if the ListServiceProviderEventsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServiceProviderEventsResult{}

// ListServiceProviderEventsResult struct for ListServiceProviderEventsResult
type ListServiceProviderEventsResult struct {
	// List of events
	Events []ServiceProviderEvent `json:"events"`
	EventsSummary ServiceProviderEventSummary `json:"eventsSummary"`
	// The next token to use for pagination
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListServiceProviderEventsResult ListServiceProviderEventsResult

// NewListServiceProviderEventsResult instantiates a new ListServiceProviderEventsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServiceProviderEventsResult(events []ServiceProviderEvent, eventsSummary ServiceProviderEventSummary) *ListServiceProviderEventsResult {
	this := ListServiceProviderEventsResult{}
	this.Events = events
	this.EventsSummary = eventsSummary
	return &this
}

// NewListServiceProviderEventsResultWithDefaults instantiates a new ListServiceProviderEventsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServiceProviderEventsResultWithDefaults() *ListServiceProviderEventsResult {
	this := ListServiceProviderEventsResult{}
	return &this
}

// GetEvents returns the Events field value
func (o *ListServiceProviderEventsResult) GetEvents() []ServiceProviderEvent {
	if o == nil {
		var ret []ServiceProviderEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEventsResult) GetEventsOk() ([]ServiceProviderEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *ListServiceProviderEventsResult) SetEvents(v []ServiceProviderEvent) {
	o.Events = v
}

// GetEventsSummary returns the EventsSummary field value
func (o *ListServiceProviderEventsResult) GetEventsSummary() ServiceProviderEventSummary {
	if o == nil {
		var ret ServiceProviderEventSummary
		return ret
	}

	return o.EventsSummary
}

// GetEventsSummaryOk returns a tuple with the EventsSummary field value
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEventsResult) GetEventsSummaryOk() (*ServiceProviderEventSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventsSummary, true
}

// SetEventsSummary sets field value
func (o *ListServiceProviderEventsResult) SetEventsSummary(v ServiceProviderEventSummary) {
	o.EventsSummary = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListServiceProviderEventsResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEventsResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListServiceProviderEventsResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ListServiceProviderEventsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListServiceProviderEventsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	toSerialize["eventsSummary"] = o.EventsSummary
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListServiceProviderEventsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
		"eventsSummary",
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

	varListServiceProviderEventsResult := _ListServiceProviderEventsResult{}

	err = json.Unmarshal(data, &varListServiceProviderEventsResult)

	if err != nil {
		return err
	}

	*o = ListServiceProviderEventsResult(varListServiceProviderEventsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "events")
		delete(additionalProperties, "eventsSummary")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListServiceProviderEventsResult struct {
	value *ListServiceProviderEventsResult
	isSet bool
}

func (v NullableListServiceProviderEventsResult) Get() *ListServiceProviderEventsResult {
	return v.value
}

func (v *NullableListServiceProviderEventsResult) Set(val *ListServiceProviderEventsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListServiceProviderEventsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListServiceProviderEventsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServiceProviderEventsResult(val *ListServiceProviderEventsResult) *NullableListServiceProviderEventsResult {
	return &NullableListServiceProviderEventsResult{value: val, isSet: true}
}

func (v NullableListServiceProviderEventsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServiceProviderEventsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


