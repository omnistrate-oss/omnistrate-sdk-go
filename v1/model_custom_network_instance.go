/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the CustomNetworkInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomNetworkInstance{}

// CustomNetworkInstance struct for CustomNetworkInstance
type CustomNetworkInstance struct {
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

type _CustomNetworkInstance CustomNetworkInstance

// NewCustomNetworkInstance instantiates a new CustomNetworkInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomNetworkInstance() *CustomNetworkInstance {
	this := CustomNetworkInstance{}
	return &this
}

// NewCustomNetworkInstanceWithDefaults instantiates a new CustomNetworkInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomNetworkInstanceWithDefaults() *CustomNetworkInstance {
	this := CustomNetworkInstance{}
	return &this
}

// GetAwsAccountID returns the AwsAccountID field value if set, zero value otherwise.
func (o *CustomNetworkInstance) GetAwsAccountID() string {
	if o == nil || IsNil(o.AwsAccountID) {
		var ret string
		return ret
	}
	return *o.AwsAccountID
}

// GetAwsAccountIDOk returns a tuple with the AwsAccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNetworkInstance) GetAwsAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountID) {
		return nil, false
	}
	return o.AwsAccountID, true
}

// SetAwsAccountID gets a reference to the given string and assigns it to the AwsAccountID field.
func (o *CustomNetworkInstance) SetAwsAccountID(v string) {
	o.AwsAccountID = &v
}

// GetAzureSubscriptionID returns the AzureSubscriptionID field value if set, zero value otherwise.
func (o *CustomNetworkInstance) GetAzureSubscriptionID() string {
	if o == nil || IsNil(o.AzureSubscriptionID) {
		var ret string
		return ret
	}
	return *o.AzureSubscriptionID
}

// GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNetworkInstance) GetAzureSubscriptionIDOk() (*string, bool) {
	if o == nil || IsNil(o.AzureSubscriptionID) {
		return nil, false
	}
	return o.AzureSubscriptionID, true
}

// SetAzureSubscriptionID gets a reference to the given string and assigns it to the AzureSubscriptionID field.
func (o *CustomNetworkInstance) SetAzureSubscriptionID(v string) {
	o.AzureSubscriptionID = &v
}

// GetAzureTenantID returns the AzureTenantID field value if set, zero value otherwise.
func (o *CustomNetworkInstance) GetAzureTenantID() string {
	if o == nil || IsNil(o.AzureTenantID) {
		var ret string
		return ret
	}
	return *o.AzureTenantID
}

// GetAzureTenantIDOk returns a tuple with the AzureTenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNetworkInstance) GetAzureTenantIDOk() (*string, bool) {
	if o == nil || IsNil(o.AzureTenantID) {
		return nil, false
	}
	return o.AzureTenantID, true
}

// SetAzureTenantID gets a reference to the given string and assigns it to the AzureTenantID field.
func (o *CustomNetworkInstance) SetAzureTenantID(v string) {
	o.AzureTenantID = &v
}

// GetCloudProviderNativeNetworkId returns the CloudProviderNativeNetworkId field value if set, zero value otherwise.
func (o *CustomNetworkInstance) GetCloudProviderNativeNetworkId() string {
	if o == nil || IsNil(o.CloudProviderNativeNetworkId) {
		var ret string
		return ret
	}
	return *o.CloudProviderNativeNetworkId
}

// GetCloudProviderNativeNetworkIdOk returns a tuple with the CloudProviderNativeNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNetworkInstance) GetCloudProviderNativeNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProviderNativeNetworkId) {
		return nil, false
	}
	return o.CloudProviderNativeNetworkId, true
}

// SetCloudProviderNativeNetworkId gets a reference to the given string and assigns it to the CloudProviderNativeNetworkId field.
func (o *CustomNetworkInstance) SetCloudProviderNativeNetworkId(v string) {
	o.CloudProviderNativeNetworkId = &v
}

// GetGcpProjectID returns the GcpProjectID field value if set, zero value otherwise.
func (o *CustomNetworkInstance) GetGcpProjectID() string {
	if o == nil || IsNil(o.GcpProjectID) {
		var ret string
		return ret
	}
	return *o.GcpProjectID
}

// GetGcpProjectIDOk returns a tuple with the GcpProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNetworkInstance) GetGcpProjectIDOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectID) {
		return nil, false
	}
	return o.GcpProjectID, true
}

// SetGcpProjectID gets a reference to the given string and assigns it to the GcpProjectID field.
func (o *CustomNetworkInstance) SetGcpProjectID(v string) {
	o.GcpProjectID = &v
}

// GetGcpProjectNumber returns the GcpProjectNumber field value if set, zero value otherwise.
func (o *CustomNetworkInstance) GetGcpProjectNumber() string {
	if o == nil || IsNil(o.GcpProjectNumber) {
		var ret string
		return ret
	}
	return *o.GcpProjectNumber
}

// GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNetworkInstance) GetGcpProjectNumberOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectNumber) {
		return nil, false
	}
	return o.GcpProjectNumber, true
}

// SetGcpProjectNumber gets a reference to the given string and assigns it to the GcpProjectNumber field.
func (o *CustomNetworkInstance) SetGcpProjectNumber(v string) {
	o.GcpProjectNumber = &v
}

// GetHostClusterID returns the HostClusterID field value if set, zero value otherwise.
func (o *CustomNetworkInstance) GetHostClusterID() string {
	if o == nil || IsNil(o.HostClusterID) {
		var ret string
		return ret
	}
	return *o.HostClusterID
}

// GetHostClusterIDOk returns a tuple with the HostClusterID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNetworkInstance) GetHostClusterIDOk() (*string, bool) {
	if o == nil || IsNil(o.HostClusterID) {
		return nil, false
	}
	return o.HostClusterID, true
}

// SetHostClusterID gets a reference to the given string and assigns it to the HostClusterID field.
func (o *CustomNetworkInstance) SetHostClusterID(v string) {
	o.HostClusterID = &v
}

func (o CustomNetworkInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomNetworkInstance) ToMap() (map[string]interface{}, error) {
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

func (o *CustomNetworkInstance) UnmarshalJSON(data []byte) (err error) {
	varCustomNetworkInstance := _CustomNetworkInstance{}

	err = json.Unmarshal(data, &varCustomNetworkInstance)

	if err != nil {
		return err
	}

	*o = CustomNetworkInstance(varCustomNetworkInstance)

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

type NullableCustomNetworkInstance struct {
	value *CustomNetworkInstance
	isSet bool
}

func (v NullableCustomNetworkInstance) Get() *CustomNetworkInstance {
	return v.value
}

func (v *NullableCustomNetworkInstance) Set(val *CustomNetworkInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomNetworkInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomNetworkInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomNetworkInstance(val *CustomNetworkInstance) *NullableCustomNetworkInstance {
	return &NullableCustomNetworkInstance{value: val, isSet: true}
}

func (v NullableCustomNetworkInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomNetworkInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


