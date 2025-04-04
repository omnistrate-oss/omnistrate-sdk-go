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

// checks if the CreateProxyResourceInstanceRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProxyResourceInstanceRequest2{}

// CreateProxyResourceInstanceRequest2 struct for CreateProxyResourceInstanceRequest2
type CreateProxyResourceInstanceRequest2 struct {
	// The cloud provider name
	CloudProvider *string `json:"cloud_provider,omitempty"`
	// The region code
	Region *string `json:"region,omitempty"`
	// The request parameters
	RequestParams interface{} `json:"requestParams,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateProxyResourceInstanceRequest2 CreateProxyResourceInstanceRequest2

// NewCreateProxyResourceInstanceRequest2 instantiates a new CreateProxyResourceInstanceRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProxyResourceInstanceRequest2() *CreateProxyResourceInstanceRequest2 {
	this := CreateProxyResourceInstanceRequest2{}
	return &this
}

// NewCreateProxyResourceInstanceRequest2WithDefaults instantiates a new CreateProxyResourceInstanceRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProxyResourceInstanceRequest2WithDefaults() *CreateProxyResourceInstanceRequest2 {
	this := CreateProxyResourceInstanceRequest2{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *CreateProxyResourceInstanceRequest2) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProxyResourceInstanceRequest2) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *CreateProxyResourceInstanceRequest2) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *CreateProxyResourceInstanceRequest2) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CreateProxyResourceInstanceRequest2) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProxyResourceInstanceRequest2) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CreateProxyResourceInstanceRequest2) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CreateProxyResourceInstanceRequest2) SetRegion(v string) {
	o.Region = &v
}

// GetRequestParams returns the RequestParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProxyResourceInstanceRequest2) GetRequestParams() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RequestParams
}

// GetRequestParamsOk returns a tuple with the RequestParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProxyResourceInstanceRequest2) GetRequestParamsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RequestParams) {
		return nil, false
	}
	return &o.RequestParams, true
}

// HasRequestParams returns a boolean if a field has been set.
func (o *CreateProxyResourceInstanceRequest2) HasRequestParams() bool {
	if o != nil && !IsNil(o.RequestParams) {
		return true
	}

	return false
}

// SetRequestParams gets a reference to the given interface{} and assigns it to the RequestParams field.
func (o *CreateProxyResourceInstanceRequest2) SetRequestParams(v interface{}) {
	o.RequestParams = v
}

func (o CreateProxyResourceInstanceRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProxyResourceInstanceRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if o.RequestParams != nil {
		toSerialize["requestParams"] = o.RequestParams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateProxyResourceInstanceRequest2) UnmarshalJSON(data []byte) (err error) {
	varCreateProxyResourceInstanceRequest2 := _CreateProxyResourceInstanceRequest2{}

	err = json.Unmarshal(data, &varCreateProxyResourceInstanceRequest2)

	if err != nil {
		return err
	}

	*o = CreateProxyResourceInstanceRequest2(varCreateProxyResourceInstanceRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloud_provider")
		delete(additionalProperties, "region")
		delete(additionalProperties, "requestParams")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateProxyResourceInstanceRequest2 struct {
	value *CreateProxyResourceInstanceRequest2
	isSet bool
}

func (v NullableCreateProxyResourceInstanceRequest2) Get() *CreateProxyResourceInstanceRequest2 {
	return v.value
}

func (v *NullableCreateProxyResourceInstanceRequest2) Set(val *CreateProxyResourceInstanceRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProxyResourceInstanceRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProxyResourceInstanceRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProxyResourceInstanceRequest2(val *CreateProxyResourceInstanceRequest2) *NullableCreateProxyResourceInstanceRequest2 {
	return &NullableCreateProxyResourceInstanceRequest2{value: val, isSet: true}
}

func (v NullableCreateProxyResourceInstanceRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProxyResourceInstanceRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


