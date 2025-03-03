/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ListServiceProviderEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServiceProviderEvents{}

// ListServiceProviderEvents struct for ListServiceProviderEvents
type ListServiceProviderEvents struct {
	// End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start.
	EndDate *time.Time `json:"endDate,omitempty"`
	// The type of service environment
	EnvironmentType *string `json:"environmentType,omitempty"`
	// The event types to filter by
	EventTypes []string `json:"eventTypes,omitempty"`
	// ID of a Resource Instance
	InstanceID *string `json:"instanceID,omitempty"`
	// The next token to use for pagination
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The number of resources to return per page
	PageSize *int64 `json:"pageSize,omitempty"`
	// ID of a Product Tier
	ProductTierID *string `json:"productTierID,omitempty"`
	// ID of a Service Environment
	ServiceEnvironmentID *string `json:"serviceEnvironmentID,omitempty"`
	// ID of a Service
	ServiceID *string `json:"serviceID,omitempty"`
	// Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start.
	StartDate *time.Time `json:"startDate,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListServiceProviderEvents ListServiceProviderEvents

// NewListServiceProviderEvents instantiates a new ListServiceProviderEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServiceProviderEvents(token string) *ListServiceProviderEvents {
	this := ListServiceProviderEvents{}
	this.Token = token
	return &this
}

// NewListServiceProviderEventsWithDefaults instantiates a new ListServiceProviderEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServiceProviderEventsWithDefaults() *ListServiceProviderEvents {
	this := ListServiceProviderEvents{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ListServiceProviderEvents) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ListServiceProviderEvents) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *ListServiceProviderEvents) GetEnvironmentType() string {
	if o == nil || IsNil(o.EnvironmentType) {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentType) {
		return nil, false
	}
	return o.EnvironmentType, true
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *ListServiceProviderEvents) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *ListServiceProviderEvents) GetEventTypes() []string {
	if o == nil || IsNil(o.EventTypes) {
		var ret []string
		return ret
	}
	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetEventTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.EventTypes) {
		return nil, false
	}
	return o.EventTypes, true
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *ListServiceProviderEvents) SetEventTypes(v []string) {
	o.EventTypes = v
}

// GetInstanceID returns the InstanceID field value if set, zero value otherwise.
func (o *ListServiceProviderEvents) GetInstanceID() string {
	if o == nil || IsNil(o.InstanceID) {
		var ret string
		return ret
	}
	return *o.InstanceID
}

// GetInstanceIDOk returns a tuple with the InstanceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetInstanceIDOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceID) {
		return nil, false
	}
	return o.InstanceID, true
}

// SetInstanceID gets a reference to the given string and assigns it to the InstanceID field.
func (o *ListServiceProviderEvents) SetInstanceID(v string) {
	o.InstanceID = &v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListServiceProviderEvents) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListServiceProviderEvents) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListServiceProviderEvents) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ListServiceProviderEvents) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetProductTierID returns the ProductTierID field value if set, zero value otherwise.
func (o *ListServiceProviderEvents) GetProductTierID() string {
	if o == nil || IsNil(o.ProductTierID) {
		var ret string
		return ret
	}
	return *o.ProductTierID
}

// GetProductTierIDOk returns a tuple with the ProductTierID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetProductTierIDOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierID) {
		return nil, false
	}
	return o.ProductTierID, true
}

// SetProductTierID gets a reference to the given string and assigns it to the ProductTierID field.
func (o *ListServiceProviderEvents) SetProductTierID(v string) {
	o.ProductTierID = &v
}

// GetServiceEnvironmentID returns the ServiceEnvironmentID field value if set, zero value otherwise.
func (o *ListServiceProviderEvents) GetServiceEnvironmentID() string {
	if o == nil || IsNil(o.ServiceEnvironmentID) {
		var ret string
		return ret
	}
	return *o.ServiceEnvironmentID
}

// GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetServiceEnvironmentIDOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceEnvironmentID) {
		return nil, false
	}
	return o.ServiceEnvironmentID, true
}

// SetServiceEnvironmentID gets a reference to the given string and assigns it to the ServiceEnvironmentID field.
func (o *ListServiceProviderEvents) SetServiceEnvironmentID(v string) {
	o.ServiceEnvironmentID = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *ListServiceProviderEvents) GetServiceID() string {
	if o == nil || IsNil(o.ServiceID) {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetServiceIDOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceID) {
		return nil, false
	}
	return o.ServiceID, true
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *ListServiceProviderEvents) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ListServiceProviderEvents) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ListServiceProviderEvents) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetToken returns the Token field value
func (o *ListServiceProviderEvents) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListServiceProviderEvents) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListServiceProviderEvents) SetToken(v string) {
	o.Token = v
}

func (o ListServiceProviderEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListServiceProviderEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.EnvironmentType) {
		toSerialize["environmentType"] = o.EnvironmentType
	}
	if !IsNil(o.EventTypes) {
		toSerialize["eventTypes"] = o.EventTypes
	}
	if !IsNil(o.InstanceID) {
		toSerialize["instanceID"] = o.InstanceID
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.ProductTierID) {
		toSerialize["productTierID"] = o.ProductTierID
	}
	if !IsNil(o.ServiceEnvironmentID) {
		toSerialize["serviceEnvironmentID"] = o.ServiceEnvironmentID
	}
	if !IsNil(o.ServiceID) {
		toSerialize["serviceID"] = o.ServiceID
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListServiceProviderEvents) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
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

	varListServiceProviderEvents := _ListServiceProviderEvents{}

	err = json.Unmarshal(data, &varListServiceProviderEvents)

	if err != nil {
		return err
	}

	*o = ListServiceProviderEvents(varListServiceProviderEvents)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "environmentType")
		delete(additionalProperties, "eventTypes")
		delete(additionalProperties, "instanceID")
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "productTierID")
		delete(additionalProperties, "serviceEnvironmentID")
		delete(additionalProperties, "serviceID")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListServiceProviderEvents struct {
	value *ListServiceProviderEvents
	isSet bool
}

func (v NullableListServiceProviderEvents) Get() *ListServiceProviderEvents {
	return v.value
}

func (v *NullableListServiceProviderEvents) Set(val *ListServiceProviderEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableListServiceProviderEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableListServiceProviderEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServiceProviderEvents(val *ListServiceProviderEvents) *NullableListServiceProviderEvents {
	return &NullableListServiceProviderEvents{value: val, isSet: true}
}

func (v NullableListServiceProviderEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServiceProviderEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


