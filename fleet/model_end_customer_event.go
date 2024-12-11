/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"fmt"
)

// checks if the EndCustomerEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndCustomerEvent{}

// EndCustomerEvent struct for EndCustomerEvent
type EndCustomerEvent struct {
	// The ID of the event
	EventID string `json:"eventID"`
	// The event payload for a service provider.
	EventPayload map[string]interface{} `json:"eventPayload"`
	// The type of the event 
	EventType string `json:"eventType"`
	// Associated organization ID.
	OrgID string `json:"orgID"`
	// Associated organization name.
	OrgName string `json:"orgName"`
	// Associated organization URL.
	OrgURL string `json:"orgURL"`
	// The priority of the event
	Priority string `json:"priority"`
	// The event time
	Time string `json:"time"`
	// User email associated with the event.
	UserEmail *string `json:"userEmail,omitempty"`
	// User ID associated with the event.
	UserID *string `json:"userID,omitempty"`
	// Name of the user associated with the event.
	UserName *string `json:"userName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EndCustomerEvent EndCustomerEvent

// NewEndCustomerEvent instantiates a new EndCustomerEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndCustomerEvent(eventID string, eventPayload map[string]interface{}, eventType string, orgID string, orgName string, orgURL string, priority string, time string) *EndCustomerEvent {
	this := EndCustomerEvent{}
	this.EventID = eventID
	this.EventPayload = eventPayload
	this.EventType = eventType
	this.OrgID = orgID
	this.OrgName = orgName
	this.OrgURL = orgURL
	this.Priority = priority
	this.Time = time
	return &this
}

// NewEndCustomerEventWithDefaults instantiates a new EndCustomerEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndCustomerEventWithDefaults() *EndCustomerEvent {
	this := EndCustomerEvent{}
	return &this
}

// GetEventID returns the EventID field value
func (o *EndCustomerEvent) GetEventID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetEventIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventID, true
}

// SetEventID sets field value
func (o *EndCustomerEvent) SetEventID(v string) {
	o.EventID = v
}

// GetEventPayload returns the EventPayload field value
func (o *EndCustomerEvent) GetEventPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.EventPayload
}

// GetEventPayloadOk returns a tuple with the EventPayload field value
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetEventPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.EventPayload, true
}

// SetEventPayload sets field value
func (o *EndCustomerEvent) SetEventPayload(v map[string]interface{}) {
	o.EventPayload = v
}

// GetEventType returns the EventType field value
func (o *EndCustomerEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *EndCustomerEvent) SetEventType(v string) {
	o.EventType = v
}

// GetOrgID returns the OrgID field value
func (o *EndCustomerEvent) GetOrgID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetOrgIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgID, true
}

// SetOrgID sets field value
func (o *EndCustomerEvent) SetOrgID(v string) {
	o.OrgID = v
}

// GetOrgName returns the OrgName field value
func (o *EndCustomerEvent) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *EndCustomerEvent) SetOrgName(v string) {
	o.OrgName = v
}

// GetOrgURL returns the OrgURL field value
func (o *EndCustomerEvent) GetOrgURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgURL
}

// GetOrgURLOk returns a tuple with the OrgURL field value
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetOrgURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgURL, true
}

// SetOrgURL sets field value
func (o *EndCustomerEvent) SetOrgURL(v string) {
	o.OrgURL = v
}

// GetPriority returns the Priority field value
func (o *EndCustomerEvent) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *EndCustomerEvent) SetPriority(v string) {
	o.Priority = v
}

// GetTime returns the Time field value
func (o *EndCustomerEvent) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *EndCustomerEvent) SetTime(v string) {
	o.Time = v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *EndCustomerEvent) GetUserEmail() string {
	if o == nil || IsNil(o.UserEmail) {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetUserEmailOk() (*string, bool) {
	if o == nil || IsNil(o.UserEmail) {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *EndCustomerEvent) HasUserEmail() bool {
	if o != nil && !IsNil(o.UserEmail) {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *EndCustomerEvent) SetUserEmail(v string) {
	o.UserEmail = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *EndCustomerEvent) GetUserID() string {
	if o == nil || IsNil(o.UserID) {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetUserIDOk() (*string, bool) {
	if o == nil || IsNil(o.UserID) {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *EndCustomerEvent) HasUserID() bool {
	if o != nil && !IsNil(o.UserID) {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *EndCustomerEvent) SetUserID(v string) {
	o.UserID = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *EndCustomerEvent) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndCustomerEvent) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *EndCustomerEvent) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *EndCustomerEvent) SetUserName(v string) {
	o.UserName = &v
}

func (o EndCustomerEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndCustomerEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventID"] = o.EventID
	toSerialize["eventPayload"] = o.EventPayload
	toSerialize["eventType"] = o.EventType
	toSerialize["orgID"] = o.OrgID
	toSerialize["orgName"] = o.OrgName
	toSerialize["orgURL"] = o.OrgURL
	toSerialize["priority"] = o.Priority
	toSerialize["time"] = o.Time
	if !IsNil(o.UserEmail) {
		toSerialize["userEmail"] = o.UserEmail
	}
	if !IsNil(o.UserID) {
		toSerialize["userID"] = o.UserID
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EndCustomerEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventID",
		"eventPayload",
		"eventType",
		"orgID",
		"orgName",
		"orgURL",
		"priority",
		"time",
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

	varEndCustomerEvent := _EndCustomerEvent{}

	err = json.Unmarshal(data, &varEndCustomerEvent)

	if err != nil {
		return err
	}

	*o = EndCustomerEvent(varEndCustomerEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "eventID")
		delete(additionalProperties, "eventPayload")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "orgID")
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "orgURL")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "time")
		delete(additionalProperties, "userEmail")
		delete(additionalProperties, "userID")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEndCustomerEvent struct {
	value *EndCustomerEvent
	isSet bool
}

func (v NullableEndCustomerEvent) Get() *EndCustomerEvent {
	return v.value
}

func (v *NullableEndCustomerEvent) Set(val *EndCustomerEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableEndCustomerEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEndCustomerEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndCustomerEvent(val *EndCustomerEvent) *NullableEndCustomerEvent {
	return &NullableEndCustomerEvent{value: val, isSet: true}
}

func (v NullableEndCustomerEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndCustomerEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


