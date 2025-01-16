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

// checks if the ServiceProviderEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProviderEvent{}

// ServiceProviderEvent struct for ServiceProviderEvent
type ServiceProviderEvent struct {
	// The type of the alert
	AlertType string `json:"alertType"`
	// The category of the service provider specific event
	EventCategory string `json:"eventCategory"`
	// ID of a Event
	EventID string `json:"eventID"`
	// The event payload for a service provider
	EventPayload map[string]interface{} `json:"eventPayload"`
	// The type of the service provider specific event
	EventType string `json:"eventType"`
	// The expiry time of the event
	ExpiryTime string `json:"expiryTime"`
	// ID of a Resource Instance
	InstanceID *string `json:"instanceID,omitempty"`
	// The version of the plan
	PlanVersion *string `json:"planVersion,omitempty"`
	// The priority of the event
	Priority string `json:"priority"`
	// The name of the resource
	ResourceName *string `json:"resourceName,omitempty"`
	// The scope of the event
	Scope string `json:"scope"`
	// ID of a Service Environment
	ServiceEnvironmentID *string `json:"serviceEnvironmentID,omitempty"`
	// ID of a Service
	ServiceID *string `json:"serviceID,omitempty"`
	// The name of the service
	ServiceName *string `json:"serviceName,omitempty"`
	// The name of the service plan
	ServicePlanName *string `json:"servicePlanName,omitempty"`
	// The event time
	Time string `json:"time"`
	AdditionalProperties map[string]interface{}
}

type _ServiceProviderEvent ServiceProviderEvent

// NewServiceProviderEvent instantiates a new ServiceProviderEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProviderEvent(alertType string, eventCategory string, eventID string, eventPayload map[string]interface{}, eventType string, expiryTime string, priority string, scope string, time string) *ServiceProviderEvent {
	this := ServiceProviderEvent{}
	this.AlertType = alertType
	this.EventCategory = eventCategory
	this.EventID = eventID
	this.EventPayload = eventPayload
	this.EventType = eventType
	this.ExpiryTime = expiryTime
	this.Priority = priority
	this.Scope = scope
	this.Time = time
	return &this
}

// NewServiceProviderEventWithDefaults instantiates a new ServiceProviderEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProviderEventWithDefaults() *ServiceProviderEvent {
	this := ServiceProviderEvent{}
	return &this
}

// GetAlertType returns the AlertType field value
func (o *ServiceProviderEvent) GetAlertType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetAlertTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertType, true
}

// SetAlertType sets field value
func (o *ServiceProviderEvent) SetAlertType(v string) {
	o.AlertType = v
}

// GetEventCategory returns the EventCategory field value
func (o *ServiceProviderEvent) GetEventCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventCategory
}

// GetEventCategoryOk returns a tuple with the EventCategory field value
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetEventCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCategory, true
}

// SetEventCategory sets field value
func (o *ServiceProviderEvent) SetEventCategory(v string) {
	o.EventCategory = v
}

// GetEventID returns the EventID field value
func (o *ServiceProviderEvent) GetEventID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetEventIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventID, true
}

// SetEventID sets field value
func (o *ServiceProviderEvent) SetEventID(v string) {
	o.EventID = v
}

// GetEventPayload returns the EventPayload field value
func (o *ServiceProviderEvent) GetEventPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.EventPayload
}

// GetEventPayloadOk returns a tuple with the EventPayload field value
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetEventPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.EventPayload, true
}

// SetEventPayload sets field value
func (o *ServiceProviderEvent) SetEventPayload(v map[string]interface{}) {
	o.EventPayload = v
}

// GetEventType returns the EventType field value
func (o *ServiceProviderEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *ServiceProviderEvent) SetEventType(v string) {
	o.EventType = v
}

// GetExpiryTime returns the ExpiryTime field value
func (o *ServiceProviderEvent) GetExpiryTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetExpiryTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryTime, true
}

// SetExpiryTime sets field value
func (o *ServiceProviderEvent) SetExpiryTime(v string) {
	o.ExpiryTime = v
}

// GetInstanceID returns the InstanceID field value if set, zero value otherwise.
func (o *ServiceProviderEvent) GetInstanceID() string {
	if o == nil || IsNil(o.InstanceID) {
		var ret string
		return ret
	}
	return *o.InstanceID
}

// GetInstanceIDOk returns a tuple with the InstanceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetInstanceIDOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceID) {
		return nil, false
	}
	return o.InstanceID, true
}

// SetInstanceID gets a reference to the given string and assigns it to the InstanceID field.
func (o *ServiceProviderEvent) SetInstanceID(v string) {
	o.InstanceID = &v
}

// GetPlanVersion returns the PlanVersion field value if set, zero value otherwise.
func (o *ServiceProviderEvent) GetPlanVersion() string {
	if o == nil || IsNil(o.PlanVersion) {
		var ret string
		return ret
	}
	return *o.PlanVersion
}

// GetPlanVersionOk returns a tuple with the PlanVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetPlanVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PlanVersion) {
		return nil, false
	}
	return o.PlanVersion, true
}

// SetPlanVersion gets a reference to the given string and assigns it to the PlanVersion field.
func (o *ServiceProviderEvent) SetPlanVersion(v string) {
	o.PlanVersion = &v
}

// GetPriority returns the Priority field value
func (o *ServiceProviderEvent) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *ServiceProviderEvent) SetPriority(v string) {
	o.Priority = v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *ServiceProviderEvent) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *ServiceProviderEvent) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetScope returns the Scope field value
func (o *ServiceProviderEvent) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *ServiceProviderEvent) SetScope(v string) {
	o.Scope = v
}

// GetServiceEnvironmentID returns the ServiceEnvironmentID field value if set, zero value otherwise.
func (o *ServiceProviderEvent) GetServiceEnvironmentID() string {
	if o == nil || IsNil(o.ServiceEnvironmentID) {
		var ret string
		return ret
	}
	return *o.ServiceEnvironmentID
}

// GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetServiceEnvironmentIDOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceEnvironmentID) {
		return nil, false
	}
	return o.ServiceEnvironmentID, true
}

// SetServiceEnvironmentID gets a reference to the given string and assigns it to the ServiceEnvironmentID field.
func (o *ServiceProviderEvent) SetServiceEnvironmentID(v string) {
	o.ServiceEnvironmentID = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *ServiceProviderEvent) GetServiceID() string {
	if o == nil || IsNil(o.ServiceID) {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetServiceIDOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceID) {
		return nil, false
	}
	return o.ServiceID, true
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *ServiceProviderEvent) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ServiceProviderEvent) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ServiceProviderEvent) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServicePlanName returns the ServicePlanName field value if set, zero value otherwise.
func (o *ServiceProviderEvent) GetServicePlanName() string {
	if o == nil || IsNil(o.ServicePlanName) {
		var ret string
		return ret
	}
	return *o.ServicePlanName
}

// GetServicePlanNameOk returns a tuple with the ServicePlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetServicePlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePlanName) {
		return nil, false
	}
	return o.ServicePlanName, true
}

// SetServicePlanName gets a reference to the given string and assigns it to the ServicePlanName field.
func (o *ServiceProviderEvent) SetServicePlanName(v string) {
	o.ServicePlanName = &v
}

// GetTime returns the Time field value
func (o *ServiceProviderEvent) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *ServiceProviderEvent) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *ServiceProviderEvent) SetTime(v string) {
	o.Time = v
}

func (o ServiceProviderEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProviderEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alertType"] = o.AlertType
	toSerialize["eventCategory"] = o.EventCategory
	toSerialize["eventID"] = o.EventID
	toSerialize["eventPayload"] = o.EventPayload
	toSerialize["eventType"] = o.EventType
	toSerialize["expiryTime"] = o.ExpiryTime
	if !IsNil(o.InstanceID) {
		toSerialize["instanceID"] = o.InstanceID
	}
	if !IsNil(o.PlanVersion) {
		toSerialize["planVersion"] = o.PlanVersion
	}
	toSerialize["priority"] = o.Priority
	if !IsNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}
	toSerialize["scope"] = o.Scope
	if !IsNil(o.ServiceEnvironmentID) {
		toSerialize["serviceEnvironmentID"] = o.ServiceEnvironmentID
	}
	if !IsNil(o.ServiceID) {
		toSerialize["serviceID"] = o.ServiceID
	}
	if !IsNil(o.ServiceName) {
		toSerialize["serviceName"] = o.ServiceName
	}
	if !IsNil(o.ServicePlanName) {
		toSerialize["servicePlanName"] = o.ServicePlanName
	}
	toSerialize["time"] = o.Time

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceProviderEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alertType",
		"eventCategory",
		"eventID",
		"eventPayload",
		"eventType",
		"expiryTime",
		"priority",
		"scope",
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

	varServiceProviderEvent := _ServiceProviderEvent{}

	err = json.Unmarshal(data, &varServiceProviderEvent)

	if err != nil {
		return err
	}

	*o = ServiceProviderEvent(varServiceProviderEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alertType")
		delete(additionalProperties, "eventCategory")
		delete(additionalProperties, "eventID")
		delete(additionalProperties, "eventPayload")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "expiryTime")
		delete(additionalProperties, "instanceID")
		delete(additionalProperties, "planVersion")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "resourceName")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "serviceEnvironmentID")
		delete(additionalProperties, "serviceID")
		delete(additionalProperties, "serviceName")
		delete(additionalProperties, "servicePlanName")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceProviderEvent struct {
	value *ServiceProviderEvent
	isSet bool
}

func (v NullableServiceProviderEvent) Get() *ServiceProviderEvent {
	return v.value
}

func (v *NullableServiceProviderEvent) Set(val *ServiceProviderEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProviderEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProviderEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProviderEvent(val *ServiceProviderEvent) *NullableServiceProviderEvent {
	return &NullableServiceProviderEvent{value: val, isSet: true}
}

func (v NullableServiceProviderEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProviderEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


