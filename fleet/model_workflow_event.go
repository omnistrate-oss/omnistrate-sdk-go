/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WorkflowEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowEvent{}

// WorkflowEvent struct for WorkflowEvent
type WorkflowEvent struct {
	// Time of the event
	EventTime string `json:"eventTime"`
	// Type of the event
	EventType string `json:"eventType"`
	// Details of the event
	Message string `json:"message"`
}

type _WorkflowEvent WorkflowEvent

// NewWorkflowEvent instantiates a new WorkflowEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowEvent(eventTime string, eventType string, message string) *WorkflowEvent {
	this := WorkflowEvent{}
	this.EventTime = eventTime
	this.EventType = eventType
	this.Message = message
	return &this
}

// NewWorkflowEventWithDefaults instantiates a new WorkflowEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowEventWithDefaults() *WorkflowEvent {
	this := WorkflowEvent{}
	return &this
}

// GetEventTime returns the EventTime field value
func (o *WorkflowEvent) GetEventTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *WorkflowEvent) GetEventTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *WorkflowEvent) SetEventTime(v string) {
	o.EventTime = v
}

// GetEventType returns the EventType field value
func (o *WorkflowEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *WorkflowEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *WorkflowEvent) SetEventType(v string) {
	o.EventType = v
}

// GetMessage returns the Message field value
func (o *WorkflowEvent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *WorkflowEvent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *WorkflowEvent) SetMessage(v string) {
	o.Message = v
}

func (o WorkflowEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventTime"] = o.EventTime
	toSerialize["eventType"] = o.EventType
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *WorkflowEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventTime",
		"eventType",
		"message",
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

	varWorkflowEvent := _WorkflowEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowEvent)

	if err != nil {
		return err
	}

	*o = WorkflowEvent(varWorkflowEvent)

	return err
}

type NullableWorkflowEvent struct {
	value *WorkflowEvent
	isSet bool
}

func (v NullableWorkflowEvent) Get() *WorkflowEvent {
	return v.value
}

func (v *NullableWorkflowEvent) Set(val *WorkflowEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowEvent(val *WorkflowEvent) *NullableWorkflowEvent {
	return &NullableWorkflowEvent{value: val, isSet: true}
}

func (v NullableWorkflowEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


