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

// checks if the DeploymentCellSearchRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentCellSearchRecord{}

// DeploymentCellSearchRecord struct for DeploymentCellSearchRecord
type DeploymentCellSearchRecord struct {
	// Name of the Infra Provider
	CloudProvider string `json:"cloudProvider"`
	// The deployment cell description.
	Description string `json:"description"`
	// The deployment cell ID.
	Id string `json:"id"`
	// The region code of the deployment cell.
	RegionCode string `json:"regionCode"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentCellSearchRecord DeploymentCellSearchRecord

// NewDeploymentCellSearchRecord instantiates a new DeploymentCellSearchRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentCellSearchRecord(cloudProvider string, description string, id string, regionCode string) *DeploymentCellSearchRecord {
	this := DeploymentCellSearchRecord{}
	this.CloudProvider = cloudProvider
	this.Description = description
	this.Id = id
	this.RegionCode = regionCode
	return &this
}

// NewDeploymentCellSearchRecordWithDefaults instantiates a new DeploymentCellSearchRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentCellSearchRecordWithDefaults() *DeploymentCellSearchRecord {
	this := DeploymentCellSearchRecord{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *DeploymentCellSearchRecord) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *DeploymentCellSearchRecord) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *DeploymentCellSearchRecord) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetDescription returns the Description field value
func (o *DeploymentCellSearchRecord) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DeploymentCellSearchRecord) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DeploymentCellSearchRecord) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *DeploymentCellSearchRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentCellSearchRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploymentCellSearchRecord) SetId(v string) {
	o.Id = v
}

// GetRegionCode returns the RegionCode field value
func (o *DeploymentCellSearchRecord) GetRegionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value
// and a boolean to check if the value has been set.
func (o *DeploymentCellSearchRecord) GetRegionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionCode, true
}

// SetRegionCode sets field value
func (o *DeploymentCellSearchRecord) SetRegionCode(v string) {
	o.RegionCode = v
}

func (o DeploymentCellSearchRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentCellSearchRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProvider"] = o.CloudProvider
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["regionCode"] = o.RegionCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentCellSearchRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProvider",
		"description",
		"id",
		"regionCode",
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

	varDeploymentCellSearchRecord := _DeploymentCellSearchRecord{}

	err = json.Unmarshal(data, &varDeploymentCellSearchRecord)

	if err != nil {
		return err
	}

	*o = DeploymentCellSearchRecord(varDeploymentCellSearchRecord)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProvider")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "regionCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentCellSearchRecord struct {
	value *DeploymentCellSearchRecord
	isSet bool
}

func (v NullableDeploymentCellSearchRecord) Get() *DeploymentCellSearchRecord {
	return v.value
}

func (v *NullableDeploymentCellSearchRecord) Set(val *DeploymentCellSearchRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentCellSearchRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentCellSearchRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentCellSearchRecord(val *DeploymentCellSearchRecord) *NullableDeploymentCellSearchRecord {
	return &NullableDeploymentCellSearchRecord{value: val, isSet: true}
}

func (v NullableDeploymentCellSearchRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentCellSearchRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


