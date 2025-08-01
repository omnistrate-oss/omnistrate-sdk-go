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

// checks if the AdoptHostClusterRequest2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdoptHostClusterRequest2{}

// AdoptHostClusterRequest2 struct for AdoptHostClusterRequest2
type AdoptHostClusterRequest2 struct {
	CloudProvider string `json:"cloudProvider"`
	// Email of the customer who owns the host cluster in case this is a BYOA host cluster
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// Description of the host cluster
	Description string `json:"description"`
	// ID of the host cluster to adopt
	Id string `json:"id"`
	// The actual region name of the host cluster
	Region string `json:"region"`
	AdditionalProperties map[string]interface{}
}

type _AdoptHostClusterRequest2 AdoptHostClusterRequest2

// NewAdoptHostClusterRequest2 instantiates a new AdoptHostClusterRequest2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdoptHostClusterRequest2(cloudProvider string, description string, id string, region string) *AdoptHostClusterRequest2 {
	this := AdoptHostClusterRequest2{}
	this.CloudProvider = cloudProvider
	this.Description = description
	this.Id = id
	this.Region = region
	return &this
}

// NewAdoptHostClusterRequest2WithDefaults instantiates a new AdoptHostClusterRequest2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdoptHostClusterRequest2WithDefaults() *AdoptHostClusterRequest2 {
	this := AdoptHostClusterRequest2{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *AdoptHostClusterRequest2) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *AdoptHostClusterRequest2) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *AdoptHostClusterRequest2) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise.
func (o *AdoptHostClusterRequest2) GetCustomerEmail() string {
	if o == nil || IsNil(o.CustomerEmail) {
		var ret string
		return ret
	}
	return *o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdoptHostClusterRequest2) GetCustomerEmailOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerEmail) {
		return nil, false
	}
	return o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *AdoptHostClusterRequest2) HasCustomerEmail() bool {
	if o != nil && !IsNil(o.CustomerEmail) {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given string and assigns it to the CustomerEmail field.
func (o *AdoptHostClusterRequest2) SetCustomerEmail(v string) {
	o.CustomerEmail = &v
}

// GetDescription returns the Description field value
func (o *AdoptHostClusterRequest2) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AdoptHostClusterRequest2) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AdoptHostClusterRequest2) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *AdoptHostClusterRequest2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdoptHostClusterRequest2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdoptHostClusterRequest2) SetId(v string) {
	o.Id = v
}

// GetRegion returns the Region field value
func (o *AdoptHostClusterRequest2) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AdoptHostClusterRequest2) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AdoptHostClusterRequest2) SetRegion(v string) {
	o.Region = v
}

func (o AdoptHostClusterRequest2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdoptHostClusterRequest2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProvider"] = o.CloudProvider
	if !IsNil(o.CustomerEmail) {
		toSerialize["customerEmail"] = o.CustomerEmail
	}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["region"] = o.Region

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdoptHostClusterRequest2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProvider",
		"description",
		"id",
		"region",
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

	varAdoptHostClusterRequest2 := _AdoptHostClusterRequest2{}

	err = json.Unmarshal(data, &varAdoptHostClusterRequest2)

	if err != nil {
		return err
	}

	*o = AdoptHostClusterRequest2(varAdoptHostClusterRequest2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProvider")
		delete(additionalProperties, "customerEmail")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdoptHostClusterRequest2 struct {
	value *AdoptHostClusterRequest2
	isSet bool
}

func (v NullableAdoptHostClusterRequest2) Get() *AdoptHostClusterRequest2 {
	return v.value
}

func (v *NullableAdoptHostClusterRequest2) Set(val *AdoptHostClusterRequest2) {
	v.value = val
	v.isSet = true
}

func (v NullableAdoptHostClusterRequest2) IsSet() bool {
	return v.isSet
}

func (v *NullableAdoptHostClusterRequest2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdoptHostClusterRequest2(val *AdoptHostClusterRequest2) *NullableAdoptHostClusterRequest2 {
	return &NullableAdoptHostClusterRequest2{value: val, isSet: true}
}

func (v NullableAdoptHostClusterRequest2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdoptHostClusterRequest2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


