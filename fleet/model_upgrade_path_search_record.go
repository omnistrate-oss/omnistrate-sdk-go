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

// checks if the UpgradePathSearchRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpgradePathSearchRecord{}

// UpgradePathSearchRecord struct for UpgradePathSearchRecord
type UpgradePathSearchRecord struct {
	// The ID of the upgrade path.
	Id string `json:"id"`
	// The product tier ID of the upgrade path.
	ProductTierID string `json:"productTierID"`
	// The product tier name of the upgrade path.
	ProductTierName string `json:"productTierName"`
	// The service environment ID of the upgrade path.
	ServiceEnvironmentId string `json:"serviceEnvironmentId"`
	// The service environment name of the upgrade path.
	ServiceEnvironmentName string `json:"serviceEnvironmentName"`
	// The service environment type of the upgrade path.
	ServiceEnvironmentType *string `json:"serviceEnvironmentType,omitempty"`
	// The service ID of the upgrade path.
	ServiceId string `json:"serviceId"`
	// The service name of the upgrade path.
	ServiceName string `json:"serviceName"`
	// The upgrade path status.
	Status string `json:"status"`
}

type _UpgradePathSearchRecord UpgradePathSearchRecord

// NewUpgradePathSearchRecord instantiates a new UpgradePathSearchRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradePathSearchRecord(id string, productTierID string, productTierName string, serviceEnvironmentId string, serviceEnvironmentName string, serviceId string, serviceName string, status string) *UpgradePathSearchRecord {
	this := UpgradePathSearchRecord{}
	this.Id = id
	this.ProductTierID = productTierID
	this.ProductTierName = productTierName
	this.ServiceEnvironmentId = serviceEnvironmentId
	this.ServiceEnvironmentName = serviceEnvironmentName
	this.ServiceId = serviceId
	this.ServiceName = serviceName
	this.Status = status
	return &this
}

// NewUpgradePathSearchRecordWithDefaults instantiates a new UpgradePathSearchRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradePathSearchRecordWithDefaults() *UpgradePathSearchRecord {
	this := UpgradePathSearchRecord{}
	return &this
}

// GetId returns the Id field value
func (o *UpgradePathSearchRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpgradePathSearchRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpgradePathSearchRecord) SetId(v string) {
	o.Id = v
}

// GetProductTierID returns the ProductTierID field value
func (o *UpgradePathSearchRecord) GetProductTierID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierID
}

// GetProductTierIDOk returns a tuple with the ProductTierID field value
// and a boolean to check if the value has been set.
func (o *UpgradePathSearchRecord) GetProductTierIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierID, true
}

// SetProductTierID sets field value
func (o *UpgradePathSearchRecord) SetProductTierID(v string) {
	o.ProductTierID = v
}

// GetProductTierName returns the ProductTierName field value
func (o *UpgradePathSearchRecord) GetProductTierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierName
}

// GetProductTierNameOk returns a tuple with the ProductTierName field value
// and a boolean to check if the value has been set.
func (o *UpgradePathSearchRecord) GetProductTierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierName, true
}

// SetProductTierName sets field value
func (o *UpgradePathSearchRecord) SetProductTierName(v string) {
	o.ProductTierName = v
}

// GetServiceEnvironmentId returns the ServiceEnvironmentId field value
func (o *UpgradePathSearchRecord) GetServiceEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentId
}

// GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field value
// and a boolean to check if the value has been set.
func (o *UpgradePathSearchRecord) GetServiceEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentId, true
}

// SetServiceEnvironmentId sets field value
func (o *UpgradePathSearchRecord) SetServiceEnvironmentId(v string) {
	o.ServiceEnvironmentId = v
}

// GetServiceEnvironmentName returns the ServiceEnvironmentName field value
func (o *UpgradePathSearchRecord) GetServiceEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentName
}

// GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field value
// and a boolean to check if the value has been set.
func (o *UpgradePathSearchRecord) GetServiceEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentName, true
}

// SetServiceEnvironmentName sets field value
func (o *UpgradePathSearchRecord) SetServiceEnvironmentName(v string) {
	o.ServiceEnvironmentName = v
}

// GetServiceEnvironmentType returns the ServiceEnvironmentType field value if set, zero value otherwise.
func (o *UpgradePathSearchRecord) GetServiceEnvironmentType() string {
	if o == nil || IsNil(o.ServiceEnvironmentType) {
		var ret string
		return ret
	}
	return *o.ServiceEnvironmentType
}

// GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradePathSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceEnvironmentType) {
		return nil, false
	}
	return o.ServiceEnvironmentType, true
}

// HasServiceEnvironmentType returns a boolean if a field has been set.
func (o *UpgradePathSearchRecord) HasServiceEnvironmentType() bool {
	if o != nil && !IsNil(o.ServiceEnvironmentType) {
		return true
	}

	return false
}

// SetServiceEnvironmentType gets a reference to the given string and assigns it to the ServiceEnvironmentType field.
func (o *UpgradePathSearchRecord) SetServiceEnvironmentType(v string) {
	o.ServiceEnvironmentType = &v
}

// GetServiceId returns the ServiceId field value
func (o *UpgradePathSearchRecord) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UpgradePathSearchRecord) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UpgradePathSearchRecord) SetServiceId(v string) {
	o.ServiceId = v
}

// GetServiceName returns the ServiceName field value
func (o *UpgradePathSearchRecord) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *UpgradePathSearchRecord) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *UpgradePathSearchRecord) SetServiceName(v string) {
	o.ServiceName = v
}

// GetStatus returns the Status field value
func (o *UpgradePathSearchRecord) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpgradePathSearchRecord) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpgradePathSearchRecord) SetStatus(v string) {
	o.Status = v
}

func (o UpgradePathSearchRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpgradePathSearchRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["productTierID"] = o.ProductTierID
	toSerialize["productTierName"] = o.ProductTierName
	toSerialize["serviceEnvironmentId"] = o.ServiceEnvironmentId
	toSerialize["serviceEnvironmentName"] = o.ServiceEnvironmentName
	if !IsNil(o.ServiceEnvironmentType) {
		toSerialize["serviceEnvironmentType"] = o.ServiceEnvironmentType
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *UpgradePathSearchRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"productTierID",
		"productTierName",
		"serviceEnvironmentId",
		"serviceEnvironmentName",
		"serviceId",
		"serviceName",
		"status",
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

	varUpgradePathSearchRecord := _UpgradePathSearchRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpgradePathSearchRecord)

	if err != nil {
		return err
	}

	*o = UpgradePathSearchRecord(varUpgradePathSearchRecord)

	return err
}

type NullableUpgradePathSearchRecord struct {
	value *UpgradePathSearchRecord
	isSet bool
}

func (v NullableUpgradePathSearchRecord) Get() *UpgradePathSearchRecord {
	return v.value
}

func (v *NullableUpgradePathSearchRecord) Set(val *UpgradePathSearchRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradePathSearchRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradePathSearchRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradePathSearchRecord(val *UpgradePathSearchRecord) *NullableUpgradePathSearchRecord {
	return &NullableUpgradePathSearchRecord{value: val, isSet: true}
}

func (v NullableUpgradePathSearchRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradePathSearchRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


