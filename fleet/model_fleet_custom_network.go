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

// checks if the FleetCustomNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetCustomNetwork{}

// FleetCustomNetwork struct for FleetCustomNetwork
type FleetCustomNetwork struct {
	// CIDR block for the network
	Cidr string `json:"cidr"`
	// Name of the Infra Provider
	CloudProviderName string `json:"cloudProviderName"`
	// The region of the cloud provider that the instance is running in.
	CloudProviderRegion string `json:"cloudProviderRegion"`
	// ID of a custom network
	Id string `json:"id"`
	// User friendly network name to help distinguish networks with same CIDRs
	Name *string `json:"name,omitempty"`
	// Type of the network definition
	NetworkDefinitionType *string `json:"networkDefinitionType,omitempty"`
	NetworkFeaturesConfiguration *FleetNetworkFeaturesConfiguration `json:"networkFeaturesConfiguration,omitempty"`
	// List of network instances created within this custom network
	NetworkInstances []FleetCustomNetworkInstance `json:"networkInstances,omitempty"`
	// ID of the owning organization
	OwningOrgID string `json:"owningOrgID"`
	// Name of the owning organization
	OwningOrgName string `json:"owningOrgName"`
	// ID of the owning user
	OwningUserID *string `json:"owningUserID,omitempty"`
	// Name of the owning user
	OwningUserName *string `json:"owningUserName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FleetCustomNetwork FleetCustomNetwork

// NewFleetCustomNetwork instantiates a new FleetCustomNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetCustomNetwork(cidr string, cloudProviderName string, cloudProviderRegion string, id string, owningOrgID string, owningOrgName string) *FleetCustomNetwork {
	this := FleetCustomNetwork{}
	this.Cidr = cidr
	this.CloudProviderName = cloudProviderName
	this.CloudProviderRegion = cloudProviderRegion
	this.Id = id
	this.OwningOrgID = owningOrgID
	this.OwningOrgName = owningOrgName
	return &this
}

// NewFleetCustomNetworkWithDefaults instantiates a new FleetCustomNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetCustomNetworkWithDefaults() *FleetCustomNetwork {
	this := FleetCustomNetwork{}
	var cidr string = "10.0.0.0/16"
	this.Cidr = cidr
	return &this
}

// GetCidr returns the Cidr field value
func (o *FleetCustomNetwork) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *FleetCustomNetwork) SetCidr(v string) {
	o.Cidr = v
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *FleetCustomNetwork) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *FleetCustomNetwork) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

// GetCloudProviderRegion returns the CloudProviderRegion field value
func (o *FleetCustomNetwork) GetCloudProviderRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderRegion
}

// GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field value
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetCloudProviderRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderRegion, true
}

// SetCloudProviderRegion sets field value
func (o *FleetCustomNetwork) SetCloudProviderRegion(v string) {
	o.CloudProviderRegion = v
}

// GetId returns the Id field value
func (o *FleetCustomNetwork) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FleetCustomNetwork) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FleetCustomNetwork) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FleetCustomNetwork) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FleetCustomNetwork) SetName(v string) {
	o.Name = &v
}

// GetNetworkDefinitionType returns the NetworkDefinitionType field value if set, zero value otherwise.
func (o *FleetCustomNetwork) GetNetworkDefinitionType() string {
	if o == nil || IsNil(o.NetworkDefinitionType) {
		var ret string
		return ret
	}
	return *o.NetworkDefinitionType
}

// GetNetworkDefinitionTypeOk returns a tuple with the NetworkDefinitionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetNetworkDefinitionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkDefinitionType) {
		return nil, false
	}
	return o.NetworkDefinitionType, true
}

// HasNetworkDefinitionType returns a boolean if a field has been set.
func (o *FleetCustomNetwork) HasNetworkDefinitionType() bool {
	if o != nil && !IsNil(o.NetworkDefinitionType) {
		return true
	}

	return false
}

// SetNetworkDefinitionType gets a reference to the given string and assigns it to the NetworkDefinitionType field.
func (o *FleetCustomNetwork) SetNetworkDefinitionType(v string) {
	o.NetworkDefinitionType = &v
}

// GetNetworkFeaturesConfiguration returns the NetworkFeaturesConfiguration field value if set, zero value otherwise.
func (o *FleetCustomNetwork) GetNetworkFeaturesConfiguration() FleetNetworkFeaturesConfiguration {
	if o == nil || IsNil(o.NetworkFeaturesConfiguration) {
		var ret FleetNetworkFeaturesConfiguration
		return ret
	}
	return *o.NetworkFeaturesConfiguration
}

// GetNetworkFeaturesConfigurationOk returns a tuple with the NetworkFeaturesConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetNetworkFeaturesConfigurationOk() (*FleetNetworkFeaturesConfiguration, bool) {
	if o == nil || IsNil(o.NetworkFeaturesConfiguration) {
		return nil, false
	}
	return o.NetworkFeaturesConfiguration, true
}

// HasNetworkFeaturesConfiguration returns a boolean if a field has been set.
func (o *FleetCustomNetwork) HasNetworkFeaturesConfiguration() bool {
	if o != nil && !IsNil(o.NetworkFeaturesConfiguration) {
		return true
	}

	return false
}

// SetNetworkFeaturesConfiguration gets a reference to the given FleetNetworkFeaturesConfiguration and assigns it to the NetworkFeaturesConfiguration field.
func (o *FleetCustomNetwork) SetNetworkFeaturesConfiguration(v FleetNetworkFeaturesConfiguration) {
	o.NetworkFeaturesConfiguration = &v
}

// GetNetworkInstances returns the NetworkInstances field value if set, zero value otherwise.
func (o *FleetCustomNetwork) GetNetworkInstances() []FleetCustomNetworkInstance {
	if o == nil || IsNil(o.NetworkInstances) {
		var ret []FleetCustomNetworkInstance
		return ret
	}
	return o.NetworkInstances
}

// GetNetworkInstancesOk returns a tuple with the NetworkInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetNetworkInstancesOk() ([]FleetCustomNetworkInstance, bool) {
	if o == nil || IsNil(o.NetworkInstances) {
		return nil, false
	}
	return o.NetworkInstances, true
}

// HasNetworkInstances returns a boolean if a field has been set.
func (o *FleetCustomNetwork) HasNetworkInstances() bool {
	if o != nil && !IsNil(o.NetworkInstances) {
		return true
	}

	return false
}

// SetNetworkInstances gets a reference to the given []FleetCustomNetworkInstance and assigns it to the NetworkInstances field.
func (o *FleetCustomNetwork) SetNetworkInstances(v []FleetCustomNetworkInstance) {
	o.NetworkInstances = v
}

// GetOwningOrgID returns the OwningOrgID field value
func (o *FleetCustomNetwork) GetOwningOrgID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwningOrgID
}

// GetOwningOrgIDOk returns a tuple with the OwningOrgID field value
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetOwningOrgIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwningOrgID, true
}

// SetOwningOrgID sets field value
func (o *FleetCustomNetwork) SetOwningOrgID(v string) {
	o.OwningOrgID = v
}

// GetOwningOrgName returns the OwningOrgName field value
func (o *FleetCustomNetwork) GetOwningOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwningOrgName
}

// GetOwningOrgNameOk returns a tuple with the OwningOrgName field value
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetOwningOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwningOrgName, true
}

// SetOwningOrgName sets field value
func (o *FleetCustomNetwork) SetOwningOrgName(v string) {
	o.OwningOrgName = v
}

// GetOwningUserID returns the OwningUserID field value if set, zero value otherwise.
func (o *FleetCustomNetwork) GetOwningUserID() string {
	if o == nil || IsNil(o.OwningUserID) {
		var ret string
		return ret
	}
	return *o.OwningUserID
}

// GetOwningUserIDOk returns a tuple with the OwningUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetOwningUserIDOk() (*string, bool) {
	if o == nil || IsNil(o.OwningUserID) {
		return nil, false
	}
	return o.OwningUserID, true
}

// HasOwningUserID returns a boolean if a field has been set.
func (o *FleetCustomNetwork) HasOwningUserID() bool {
	if o != nil && !IsNil(o.OwningUserID) {
		return true
	}

	return false
}

// SetOwningUserID gets a reference to the given string and assigns it to the OwningUserID field.
func (o *FleetCustomNetwork) SetOwningUserID(v string) {
	o.OwningUserID = &v
}

// GetOwningUserName returns the OwningUserName field value if set, zero value otherwise.
func (o *FleetCustomNetwork) GetOwningUserName() string {
	if o == nil || IsNil(o.OwningUserName) {
		var ret string
		return ret
	}
	return *o.OwningUserName
}

// GetOwningUserNameOk returns a tuple with the OwningUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetwork) GetOwningUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwningUserName) {
		return nil, false
	}
	return o.OwningUserName, true
}

// HasOwningUserName returns a boolean if a field has been set.
func (o *FleetCustomNetwork) HasOwningUserName() bool {
	if o != nil && !IsNil(o.OwningUserName) {
		return true
	}

	return false
}

// SetOwningUserName gets a reference to the given string and assigns it to the OwningUserName field.
func (o *FleetCustomNetwork) SetOwningUserName(v string) {
	o.OwningUserName = &v
}

func (o FleetCustomNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetCustomNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cidr"] = o.Cidr
	toSerialize["cloudProviderName"] = o.CloudProviderName
	toSerialize["cloudProviderRegion"] = o.CloudProviderRegion
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetworkDefinitionType) {
		toSerialize["networkDefinitionType"] = o.NetworkDefinitionType
	}
	if !IsNil(o.NetworkFeaturesConfiguration) {
		toSerialize["networkFeaturesConfiguration"] = o.NetworkFeaturesConfiguration
	}
	if !IsNil(o.NetworkInstances) {
		toSerialize["networkInstances"] = o.NetworkInstances
	}
	toSerialize["owningOrgID"] = o.OwningOrgID
	toSerialize["owningOrgName"] = o.OwningOrgName
	if !IsNil(o.OwningUserID) {
		toSerialize["owningUserID"] = o.OwningUserID
	}
	if !IsNil(o.OwningUserName) {
		toSerialize["owningUserName"] = o.OwningUserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetCustomNetwork) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cidr",
		"cloudProviderName",
		"cloudProviderRegion",
		"id",
		"owningOrgID",
		"owningOrgName",
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

	varFleetCustomNetwork := _FleetCustomNetwork{}

	err = json.Unmarshal(data, &varFleetCustomNetwork)

	if err != nil {
		return err
	}

	*o = FleetCustomNetwork(varFleetCustomNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cidr")
		delete(additionalProperties, "cloudProviderName")
		delete(additionalProperties, "cloudProviderRegion")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "networkDefinitionType")
		delete(additionalProperties, "networkFeaturesConfiguration")
		delete(additionalProperties, "networkInstances")
		delete(additionalProperties, "owningOrgID")
		delete(additionalProperties, "owningOrgName")
		delete(additionalProperties, "owningUserID")
		delete(additionalProperties, "owningUserName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetCustomNetwork struct {
	value *FleetCustomNetwork
	isSet bool
}

func (v NullableFleetCustomNetwork) Get() *FleetCustomNetwork {
	return v.value
}

func (v *NullableFleetCustomNetwork) Set(val *FleetCustomNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetCustomNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetCustomNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetCustomNetwork(val *FleetCustomNetwork) *NullableFleetCustomNetwork {
	return &NullableFleetCustomNetwork{value: val, isSet: true}
}

func (v NullableFleetCustomNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetCustomNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


