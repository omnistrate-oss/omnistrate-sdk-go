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
	// The cloud provider of the deployment cell.
	CloudProvider string `json:"cloudProvider"`
	// The deployment cell description.
	Description string `json:"description"`
	// The deployment cell ID.
	Id string `json:"id"`
	// The region code of the deployment cell.
	RegionCode string `json:"regionCode"`
	// The service environment ID of the deployment cell.
	ServiceEnvironmentID string `json:"serviceEnvironmentID"`
	// The service environment name of the deployment cell.
	ServiceEnvironmentName string `json:"serviceEnvironmentName"`
	// The service environment type of the deployment cell.
	ServiceEnvironmentType *string `json:"serviceEnvironmentType,omitempty"`
	// The service ID of the deployment cell.
	ServiceID string `json:"serviceID"`
	// The service name of the deployment cell.
	ServiceName string `json:"serviceName"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentCellSearchRecord DeploymentCellSearchRecord

// NewDeploymentCellSearchRecord instantiates a new DeploymentCellSearchRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentCellSearchRecord(cloudProvider string, description string, id string, regionCode string, serviceEnvironmentID string, serviceEnvironmentName string, serviceID string, serviceName string) *DeploymentCellSearchRecord {
	this := DeploymentCellSearchRecord{}
	this.CloudProvider = cloudProvider
	this.Description = description
	this.Id = id
	this.RegionCode = regionCode
	this.ServiceEnvironmentID = serviceEnvironmentID
	this.ServiceEnvironmentName = serviceEnvironmentName
	this.ServiceID = serviceID
	this.ServiceName = serviceName
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

// GetServiceEnvironmentID returns the ServiceEnvironmentID field value
func (o *DeploymentCellSearchRecord) GetServiceEnvironmentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentID
}

// GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field value
// and a boolean to check if the value has been set.
func (o *DeploymentCellSearchRecord) GetServiceEnvironmentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentID, true
}

// SetServiceEnvironmentID sets field value
func (o *DeploymentCellSearchRecord) SetServiceEnvironmentID(v string) {
	o.ServiceEnvironmentID = v
}

// GetServiceEnvironmentName returns the ServiceEnvironmentName field value
func (o *DeploymentCellSearchRecord) GetServiceEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentName
}

// GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field value
// and a boolean to check if the value has been set.
func (o *DeploymentCellSearchRecord) GetServiceEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentName, true
}

// SetServiceEnvironmentName sets field value
func (o *DeploymentCellSearchRecord) SetServiceEnvironmentName(v string) {
	o.ServiceEnvironmentName = v
}

// GetServiceEnvironmentType returns the ServiceEnvironmentType field value if set, zero value otherwise.
func (o *DeploymentCellSearchRecord) GetServiceEnvironmentType() string {
	if o == nil || IsNil(o.ServiceEnvironmentType) {
		var ret string
		return ret
	}
	return *o.ServiceEnvironmentType
}

// GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentCellSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceEnvironmentType) {
		return nil, false
	}
	return o.ServiceEnvironmentType, true
}

// HasServiceEnvironmentType returns a boolean if a field has been set.
func (o *DeploymentCellSearchRecord) HasServiceEnvironmentType() bool {
	if o != nil && !IsNil(o.ServiceEnvironmentType) {
		return true
	}

	return false
}

// SetServiceEnvironmentType gets a reference to the given string and assigns it to the ServiceEnvironmentType field.
func (o *DeploymentCellSearchRecord) SetServiceEnvironmentType(v string) {
	o.ServiceEnvironmentType = &v
}

// GetServiceID returns the ServiceID field value
func (o *DeploymentCellSearchRecord) GetServiceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value
// and a boolean to check if the value has been set.
func (o *DeploymentCellSearchRecord) GetServiceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceID, true
}

// SetServiceID sets field value
func (o *DeploymentCellSearchRecord) SetServiceID(v string) {
	o.ServiceID = v
}

// GetServiceName returns the ServiceName field value
func (o *DeploymentCellSearchRecord) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *DeploymentCellSearchRecord) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *DeploymentCellSearchRecord) SetServiceName(v string) {
	o.ServiceName = v
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
	toSerialize["serviceEnvironmentID"] = o.ServiceEnvironmentID
	toSerialize["serviceEnvironmentName"] = o.ServiceEnvironmentName
	if !IsNil(o.ServiceEnvironmentType) {
		toSerialize["serviceEnvironmentType"] = o.ServiceEnvironmentType
	}
	toSerialize["serviceID"] = o.ServiceID
	toSerialize["serviceName"] = o.ServiceName

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
		"serviceEnvironmentID",
		"serviceEnvironmentName",
		"serviceID",
		"serviceName",
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
		delete(additionalProperties, "serviceEnvironmentID")
		delete(additionalProperties, "serviceEnvironmentName")
		delete(additionalProperties, "serviceEnvironmentType")
		delete(additionalProperties, "serviceID")
		delete(additionalProperties, "serviceName")
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


