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

// checks if the FleetCustomNetworkInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetCustomNetworkInstance{}

// FleetCustomNetworkInstance struct for FleetCustomNetworkInstance
type FleetCustomNetworkInstance struct {
	// The AWS account ID
	AwsAccountID *string `json:"awsAccountID,omitempty"`
	// The Azure subscription ID
	AzureSubscriptionID *string `json:"azureSubscriptionID,omitempty"`
	// The Azure tenant ID
	AzureTenantID *string `json:"azureTenantID,omitempty"`
	// The ID of the network within cloud provider account
	CloudProviderNativeNetworkId *string `json:"cloudProviderNativeNetworkId,omitempty"`
	// The GCP project ID
	GcpProjectID *string `json:"gcpProjectID,omitempty"`
	// The GCP project number
	GcpProjectNumber *string `json:"gcpProjectNumber,omitempty"`
	// ID of a Host Cluster
	HostClusterID *string `json:"hostClusterID,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FleetCustomNetworkInstance FleetCustomNetworkInstance

// NewFleetCustomNetworkInstance instantiates a new FleetCustomNetworkInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetCustomNetworkInstance() *FleetCustomNetworkInstance {
	this := FleetCustomNetworkInstance{}
	return &this
}

// NewFleetCustomNetworkInstanceWithDefaults instantiates a new FleetCustomNetworkInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetCustomNetworkInstanceWithDefaults() *FleetCustomNetworkInstance {
	this := FleetCustomNetworkInstance{}
	return &this
}

// GetAwsAccountID returns the AwsAccountID field value if set, zero value otherwise.
func (o *FleetCustomNetworkInstance) GetAwsAccountID() string {
	if o == nil || IsNil(o.AwsAccountID) {
		var ret string
		return ret
	}
	return *o.AwsAccountID
}

// GetAwsAccountIDOk returns a tuple with the AwsAccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetworkInstance) GetAwsAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountID) {
		return nil, false
	}
	return o.AwsAccountID, true
}

// HasAwsAccountID returns a boolean if a field has been set.
func (o *FleetCustomNetworkInstance) HasAwsAccountID() bool {
	if o != nil && !IsNil(o.AwsAccountID) {
		return true
	}

	return false
}

// SetAwsAccountID gets a reference to the given string and assigns it to the AwsAccountID field.
func (o *FleetCustomNetworkInstance) SetAwsAccountID(v string) {
	o.AwsAccountID = &v
}

// GetAzureSubscriptionID returns the AzureSubscriptionID field value if set, zero value otherwise.
func (o *FleetCustomNetworkInstance) GetAzureSubscriptionID() string {
	if o == nil || IsNil(o.AzureSubscriptionID) {
		var ret string
		return ret
	}
	return *o.AzureSubscriptionID
}

// GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetworkInstance) GetAzureSubscriptionIDOk() (*string, bool) {
	if o == nil || IsNil(o.AzureSubscriptionID) {
		return nil, false
	}
	return o.AzureSubscriptionID, true
}

// HasAzureSubscriptionID returns a boolean if a field has been set.
func (o *FleetCustomNetworkInstance) HasAzureSubscriptionID() bool {
	if o != nil && !IsNil(o.AzureSubscriptionID) {
		return true
	}

	return false
}

// SetAzureSubscriptionID gets a reference to the given string and assigns it to the AzureSubscriptionID field.
func (o *FleetCustomNetworkInstance) SetAzureSubscriptionID(v string) {
	o.AzureSubscriptionID = &v
}

// GetAzureTenantID returns the AzureTenantID field value if set, zero value otherwise.
func (o *FleetCustomNetworkInstance) GetAzureTenantID() string {
	if o == nil || IsNil(o.AzureTenantID) {
		var ret string
		return ret
	}
	return *o.AzureTenantID
}

// GetAzureTenantIDOk returns a tuple with the AzureTenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetworkInstance) GetAzureTenantIDOk() (*string, bool) {
	if o == nil || IsNil(o.AzureTenantID) {
		return nil, false
	}
	return o.AzureTenantID, true
}

// HasAzureTenantID returns a boolean if a field has been set.
func (o *FleetCustomNetworkInstance) HasAzureTenantID() bool {
	if o != nil && !IsNil(o.AzureTenantID) {
		return true
	}

	return false
}

// SetAzureTenantID gets a reference to the given string and assigns it to the AzureTenantID field.
func (o *FleetCustomNetworkInstance) SetAzureTenantID(v string) {
	o.AzureTenantID = &v
}

// GetCloudProviderNativeNetworkId returns the CloudProviderNativeNetworkId field value if set, zero value otherwise.
func (o *FleetCustomNetworkInstance) GetCloudProviderNativeNetworkId() string {
	if o == nil || IsNil(o.CloudProviderNativeNetworkId) {
		var ret string
		return ret
	}
	return *o.CloudProviderNativeNetworkId
}

// GetCloudProviderNativeNetworkIdOk returns a tuple with the CloudProviderNativeNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetworkInstance) GetCloudProviderNativeNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProviderNativeNetworkId) {
		return nil, false
	}
	return o.CloudProviderNativeNetworkId, true
}

// HasCloudProviderNativeNetworkId returns a boolean if a field has been set.
func (o *FleetCustomNetworkInstance) HasCloudProviderNativeNetworkId() bool {
	if o != nil && !IsNil(o.CloudProviderNativeNetworkId) {
		return true
	}

	return false
}

// SetCloudProviderNativeNetworkId gets a reference to the given string and assigns it to the CloudProviderNativeNetworkId field.
func (o *FleetCustomNetworkInstance) SetCloudProviderNativeNetworkId(v string) {
	o.CloudProviderNativeNetworkId = &v
}

// GetGcpProjectID returns the GcpProjectID field value if set, zero value otherwise.
func (o *FleetCustomNetworkInstance) GetGcpProjectID() string {
	if o == nil || IsNil(o.GcpProjectID) {
		var ret string
		return ret
	}
	return *o.GcpProjectID
}

// GetGcpProjectIDOk returns a tuple with the GcpProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetworkInstance) GetGcpProjectIDOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectID) {
		return nil, false
	}
	return o.GcpProjectID, true
}

// HasGcpProjectID returns a boolean if a field has been set.
func (o *FleetCustomNetworkInstance) HasGcpProjectID() bool {
	if o != nil && !IsNil(o.GcpProjectID) {
		return true
	}

	return false
}

// SetGcpProjectID gets a reference to the given string and assigns it to the GcpProjectID field.
func (o *FleetCustomNetworkInstance) SetGcpProjectID(v string) {
	o.GcpProjectID = &v
}

// GetGcpProjectNumber returns the GcpProjectNumber field value if set, zero value otherwise.
func (o *FleetCustomNetworkInstance) GetGcpProjectNumber() string {
	if o == nil || IsNil(o.GcpProjectNumber) {
		var ret string
		return ret
	}
	return *o.GcpProjectNumber
}

// GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetworkInstance) GetGcpProjectNumberOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectNumber) {
		return nil, false
	}
	return o.GcpProjectNumber, true
}

// HasGcpProjectNumber returns a boolean if a field has been set.
func (o *FleetCustomNetworkInstance) HasGcpProjectNumber() bool {
	if o != nil && !IsNil(o.GcpProjectNumber) {
		return true
	}

	return false
}

// SetGcpProjectNumber gets a reference to the given string and assigns it to the GcpProjectNumber field.
func (o *FleetCustomNetworkInstance) SetGcpProjectNumber(v string) {
	o.GcpProjectNumber = &v
}

// GetHostClusterID returns the HostClusterID field value if set, zero value otherwise.
func (o *FleetCustomNetworkInstance) GetHostClusterID() string {
	if o == nil || IsNil(o.HostClusterID) {
		var ret string
		return ret
	}
	return *o.HostClusterID
}

// GetHostClusterIDOk returns a tuple with the HostClusterID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetCustomNetworkInstance) GetHostClusterIDOk() (*string, bool) {
	if o == nil || IsNil(o.HostClusterID) {
		return nil, false
	}
	return o.HostClusterID, true
}

// HasHostClusterID returns a boolean if a field has been set.
func (o *FleetCustomNetworkInstance) HasHostClusterID() bool {
	if o != nil && !IsNil(o.HostClusterID) {
		return true
	}

	return false
}

// SetHostClusterID gets a reference to the given string and assigns it to the HostClusterID field.
func (o *FleetCustomNetworkInstance) SetHostClusterID(v string) {
	o.HostClusterID = &v
}

func (o FleetCustomNetworkInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetCustomNetworkInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsAccountID) {
		toSerialize["awsAccountID"] = o.AwsAccountID
	}
	if !IsNil(o.AzureSubscriptionID) {
		toSerialize["azureSubscriptionID"] = o.AzureSubscriptionID
	}
	if !IsNil(o.AzureTenantID) {
		toSerialize["azureTenantID"] = o.AzureTenantID
	}
	if !IsNil(o.CloudProviderNativeNetworkId) {
		toSerialize["cloudProviderNativeNetworkId"] = o.CloudProviderNativeNetworkId
	}
	if !IsNil(o.GcpProjectID) {
		toSerialize["gcpProjectID"] = o.GcpProjectID
	}
	if !IsNil(o.GcpProjectNumber) {
		toSerialize["gcpProjectNumber"] = o.GcpProjectNumber
	}
	if !IsNil(o.HostClusterID) {
		toSerialize["hostClusterID"] = o.HostClusterID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FleetCustomNetworkInstance) UnmarshalJSON(data []byte) (err error) {
	varFleetCustomNetworkInstance := _FleetCustomNetworkInstance{}

	err = json.Unmarshal(data, &varFleetCustomNetworkInstance)

	if err != nil {
		return err
	}

	*o = FleetCustomNetworkInstance(varFleetCustomNetworkInstance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "awsAccountID")
		delete(additionalProperties, "azureSubscriptionID")
		delete(additionalProperties, "azureTenantID")
		delete(additionalProperties, "cloudProviderNativeNetworkId")
		delete(additionalProperties, "gcpProjectID")
		delete(additionalProperties, "gcpProjectNumber")
		delete(additionalProperties, "hostClusterID")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFleetCustomNetworkInstance struct {
	value *FleetCustomNetworkInstance
	isSet bool
}

func (v NullableFleetCustomNetworkInstance) Get() *FleetCustomNetworkInstance {
	return v.value
}

func (v *NullableFleetCustomNetworkInstance) Set(val *FleetCustomNetworkInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetCustomNetworkInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetCustomNetworkInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetCustomNetworkInstance(val *FleetCustomNetworkInstance) *NullableFleetCustomNetworkInstance {
	return &NullableFleetCustomNetworkInstance{value: val, isSet: true}
}

func (v NullableFleetCustomNetworkInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetCustomNetworkInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


