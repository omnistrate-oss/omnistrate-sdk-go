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

// checks if the ListServiceOfferingsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServiceOfferingsResult{}

// ListServiceOfferingsResult struct for ListServiceOfferingsResult
type ListServiceOfferingsResult struct {
	// Token to use for the next request
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The service IDs
	ServiceIds []string `json:"serviceIds"`
	// List of services
	Services []DescribeServiceOfferingResult `json:"services,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListServiceOfferingsResult ListServiceOfferingsResult

// NewListServiceOfferingsResult instantiates a new ListServiceOfferingsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServiceOfferingsResult(serviceIds []string) *ListServiceOfferingsResult {
	this := ListServiceOfferingsResult{}
	this.ServiceIds = serviceIds
	return &this
}

// NewListServiceOfferingsResultWithDefaults instantiates a new ListServiceOfferingsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServiceOfferingsResultWithDefaults() *ListServiceOfferingsResult {
	this := ListServiceOfferingsResult{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListServiceOfferingsResult) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceOfferingsResult) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListServiceOfferingsResult) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListServiceOfferingsResult) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetServiceIds returns the ServiceIds field value
func (o *ListServiceOfferingsResult) GetServiceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ServiceIds
}

// GetServiceIdsOk returns a tuple with the ServiceIds field value
// and a boolean to check if the value has been set.
func (o *ListServiceOfferingsResult) GetServiceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceIds, true
}

// SetServiceIds sets field value
func (o *ListServiceOfferingsResult) SetServiceIds(v []string) {
	o.ServiceIds = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ListServiceOfferingsResult) GetServices() []DescribeServiceOfferingResult {
	if o == nil || IsNil(o.Services) {
		var ret []DescribeServiceOfferingResult
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceOfferingsResult) GetServicesOk() ([]DescribeServiceOfferingResult, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ListServiceOfferingsResult) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []DescribeServiceOfferingResult and assigns it to the Services field.
func (o *ListServiceOfferingsResult) SetServices(v []DescribeServiceOfferingResult) {
	o.Services = v
}

func (o ListServiceOfferingsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListServiceOfferingsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	toSerialize["serviceIds"] = o.ServiceIds
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListServiceOfferingsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serviceIds",
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

	varListServiceOfferingsResult := _ListServiceOfferingsResult{}

	err = json.Unmarshal(data, &varListServiceOfferingsResult)

	if err != nil {
		return err
	}

	*o = ListServiceOfferingsResult(varListServiceOfferingsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nextPageToken")
		delete(additionalProperties, "serviceIds")
		delete(additionalProperties, "services")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListServiceOfferingsResult struct {
	value *ListServiceOfferingsResult
	isSet bool
}

func (v NullableListServiceOfferingsResult) Get() *ListServiceOfferingsResult {
	return v.value
}

func (v *NullableListServiceOfferingsResult) Set(val *ListServiceOfferingsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListServiceOfferingsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListServiceOfferingsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServiceOfferingsResult(val *ListServiceOfferingsResult) *NullableListServiceOfferingsResult {
	return &NullableListServiceOfferingsResult{value: val, isSet: true}
}

func (v NullableListServiceOfferingsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServiceOfferingsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


