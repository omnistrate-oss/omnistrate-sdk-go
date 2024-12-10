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

// checks if the ListServiceWorkflowsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServiceWorkflowsRequest{}

// ListServiceWorkflowsRequest struct for ListServiceWorkflowsRequest
type ListServiceWorkflowsRequest struct {
	// ID of a Service Environment
	EnvironmentId string `json:"environmentId"`
	// ID of a Resource Instance
	InstanceId *string `json:"instanceId,omitempty"`
	// The next token to use for pagination
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The number of resources to return per page
	PageSize *int64 `json:"pageSize,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ListServiceWorkflowsRequest ListServiceWorkflowsRequest

// NewListServiceWorkflowsRequest instantiates a new ListServiceWorkflowsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServiceWorkflowsRequest(environmentId string, serviceId string, token string) *ListServiceWorkflowsRequest {
	this := ListServiceWorkflowsRequest{}
	this.EnvironmentId = environmentId
	this.ServiceId = serviceId
	this.Token = token
	return &this
}

// NewListServiceWorkflowsRequestWithDefaults instantiates a new ListServiceWorkflowsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServiceWorkflowsRequestWithDefaults() *ListServiceWorkflowsRequest {
	this := ListServiceWorkflowsRequest{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *ListServiceWorkflowsRequest) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *ListServiceWorkflowsRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *ListServiceWorkflowsRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *ListServiceWorkflowsRequest) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceWorkflowsRequest) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ListServiceWorkflowsRequest) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *ListServiceWorkflowsRequest) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListServiceWorkflowsRequest) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceWorkflowsRequest) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListServiceWorkflowsRequest) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListServiceWorkflowsRequest) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListServiceWorkflowsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceWorkflowsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListServiceWorkflowsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ListServiceWorkflowsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetServiceId returns the ServiceId field value
func (o *ListServiceWorkflowsRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ListServiceWorkflowsRequest) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ListServiceWorkflowsRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetToken returns the Token field value
func (o *ListServiceWorkflowsRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ListServiceWorkflowsRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ListServiceWorkflowsRequest) SetToken(v string) {
	o.Token = v
}

func (o ListServiceWorkflowsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListServiceWorkflowsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentId"] = o.EnvironmentId
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListServiceWorkflowsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentId",
		"serviceId",
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

	varListServiceWorkflowsRequest := _ListServiceWorkflowsRequest{}

	err = json.Unmarshal(data, &varListServiceWorkflowsRequest)

	if err != nil {
		return err
	}

	*o = ListServiceWorkflowsRequest(varListServiceWorkflowsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListServiceWorkflowsRequest struct {
	value *ListServiceWorkflowsRequest
	isSet bool
}

func (v NullableListServiceWorkflowsRequest) Get() *ListServiceWorkflowsRequest {
	return v.value
}

func (v *NullableListServiceWorkflowsRequest) Set(val *ListServiceWorkflowsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListServiceWorkflowsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListServiceWorkflowsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServiceWorkflowsRequest(val *ListServiceWorkflowsRequest) *NullableListServiceWorkflowsRequest {
	return &NullableListServiceWorkflowsRequest{value: val, isSet: true}
}

func (v NullableListServiceWorkflowsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServiceWorkflowsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


