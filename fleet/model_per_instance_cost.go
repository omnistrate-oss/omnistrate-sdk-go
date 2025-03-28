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

// checks if the PerInstanceCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerInstanceCost{}

// PerInstanceCost struct for PerInstanceCost
type PerInstanceCost struct {
	// Name of the Infra Provider
	CloudProviderName string `json:"cloudProviderName"`
	// ID of a Resource Instance
	InstanceID string `json:"instanceID"`
	// Whether the instance is deleted
	IsDeleted bool `json:"isDeleted"`
	// ID of a Product Tier
	ProductTierID string `json:"productTierID"`
	// The name of the product tier
	ProductTierName string `json:"productTierName"`
	// The tenancy type of the product tier
	ProductTierTenancyType string `json:"productTierTenancyType"`
	// The name of the region
	RegionName string `json:"regionName"`
	// ID of a Service Environment
	ServiceEnvironmentID string `json:"serviceEnvironmentID"`
	// ID of a Service
	ServiceID string `json:"serviceID"`
	// ID of a Subscription
	SubscriptionID string `json:"subscriptionID"`
	// The total cost of the instance
	TotalCost float64 `json:"totalCost"`
	// Percentage of the instance resource utilization
	Utilization float64 `json:"utilization"`
	AdditionalProperties map[string]interface{}
}

type _PerInstanceCost PerInstanceCost

// NewPerInstanceCost instantiates a new PerInstanceCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerInstanceCost(cloudProviderName string, instanceID string, isDeleted bool, productTierID string, productTierName string, productTierTenancyType string, regionName string, serviceEnvironmentID string, serviceID string, subscriptionID string, totalCost float64, utilization float64) *PerInstanceCost {
	this := PerInstanceCost{}
	this.CloudProviderName = cloudProviderName
	this.InstanceID = instanceID
	this.IsDeleted = isDeleted
	this.ProductTierID = productTierID
	this.ProductTierName = productTierName
	this.ProductTierTenancyType = productTierTenancyType
	this.RegionName = regionName
	this.ServiceEnvironmentID = serviceEnvironmentID
	this.ServiceID = serviceID
	this.SubscriptionID = subscriptionID
	this.TotalCost = totalCost
	this.Utilization = utilization
	return &this
}

// NewPerInstanceCostWithDefaults instantiates a new PerInstanceCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerInstanceCostWithDefaults() *PerInstanceCost {
	this := PerInstanceCost{}
	return &this
}

// GetCloudProviderName returns the CloudProviderName field value
func (o *PerInstanceCost) GetCloudProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderName
}

// GetCloudProviderNameOk returns a tuple with the CloudProviderName field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetCloudProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderName, true
}

// SetCloudProviderName sets field value
func (o *PerInstanceCost) SetCloudProviderName(v string) {
	o.CloudProviderName = v
}

// GetInstanceID returns the InstanceID field value
func (o *PerInstanceCost) GetInstanceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceID
}

// GetInstanceIDOk returns a tuple with the InstanceID field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetInstanceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceID, true
}

// SetInstanceID sets field value
func (o *PerInstanceCost) SetInstanceID(v string) {
	o.InstanceID = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *PerInstanceCost) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *PerInstanceCost) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetProductTierID returns the ProductTierID field value
func (o *PerInstanceCost) GetProductTierID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierID
}

// GetProductTierIDOk returns a tuple with the ProductTierID field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetProductTierIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierID, true
}

// SetProductTierID sets field value
func (o *PerInstanceCost) SetProductTierID(v string) {
	o.ProductTierID = v
}

// GetProductTierName returns the ProductTierName field value
func (o *PerInstanceCost) GetProductTierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierName
}

// GetProductTierNameOk returns a tuple with the ProductTierName field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetProductTierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierName, true
}

// SetProductTierName sets field value
func (o *PerInstanceCost) SetProductTierName(v string) {
	o.ProductTierName = v
}

// GetProductTierTenancyType returns the ProductTierTenancyType field value
func (o *PerInstanceCost) GetProductTierTenancyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierTenancyType
}

// GetProductTierTenancyTypeOk returns a tuple with the ProductTierTenancyType field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetProductTierTenancyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierTenancyType, true
}

// SetProductTierTenancyType sets field value
func (o *PerInstanceCost) SetProductTierTenancyType(v string) {
	o.ProductTierTenancyType = v
}

// GetRegionName returns the RegionName field value
func (o *PerInstanceCost) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *PerInstanceCost) SetRegionName(v string) {
	o.RegionName = v
}

// GetServiceEnvironmentID returns the ServiceEnvironmentID field value
func (o *PerInstanceCost) GetServiceEnvironmentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentID
}

// GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetServiceEnvironmentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentID, true
}

// SetServiceEnvironmentID sets field value
func (o *PerInstanceCost) SetServiceEnvironmentID(v string) {
	o.ServiceEnvironmentID = v
}

// GetServiceID returns the ServiceID field value
func (o *PerInstanceCost) GetServiceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetServiceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceID, true
}

// SetServiceID sets field value
func (o *PerInstanceCost) SetServiceID(v string) {
	o.ServiceID = v
}

// GetSubscriptionID returns the SubscriptionID field value
func (o *PerInstanceCost) GetSubscriptionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionID
}

// GetSubscriptionIDOk returns a tuple with the SubscriptionID field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetSubscriptionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionID, true
}

// SetSubscriptionID sets field value
func (o *PerInstanceCost) SetSubscriptionID(v string) {
	o.SubscriptionID = v
}

// GetTotalCost returns the TotalCost field value
func (o *PerInstanceCost) GetTotalCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetTotalCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCost, true
}

// SetTotalCost sets field value
func (o *PerInstanceCost) SetTotalCost(v float64) {
	o.TotalCost = v
}

// GetUtilization returns the Utilization field value
func (o *PerInstanceCost) GetUtilization() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value
// and a boolean to check if the value has been set.
func (o *PerInstanceCost) GetUtilizationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Utilization, true
}

// SetUtilization sets field value
func (o *PerInstanceCost) SetUtilization(v float64) {
	o.Utilization = v
}

func (o PerInstanceCost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerInstanceCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderName"] = o.CloudProviderName
	toSerialize["instanceID"] = o.InstanceID
	toSerialize["isDeleted"] = o.IsDeleted
	toSerialize["productTierID"] = o.ProductTierID
	toSerialize["productTierName"] = o.ProductTierName
	toSerialize["productTierTenancyType"] = o.ProductTierTenancyType
	toSerialize["regionName"] = o.RegionName
	toSerialize["serviceEnvironmentID"] = o.ServiceEnvironmentID
	toSerialize["serviceID"] = o.ServiceID
	toSerialize["subscriptionID"] = o.SubscriptionID
	toSerialize["totalCost"] = o.TotalCost
	toSerialize["utilization"] = o.Utilization

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PerInstanceCost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderName",
		"instanceID",
		"isDeleted",
		"productTierID",
		"productTierName",
		"productTierTenancyType",
		"regionName",
		"serviceEnvironmentID",
		"serviceID",
		"subscriptionID",
		"totalCost",
		"utilization",
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

	varPerInstanceCost := _PerInstanceCost{}

	err = json.Unmarshal(data, &varPerInstanceCost)

	if err != nil {
		return err
	}

	*o = PerInstanceCost(varPerInstanceCost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProviderName")
		delete(additionalProperties, "instanceID")
		delete(additionalProperties, "isDeleted")
		delete(additionalProperties, "productTierID")
		delete(additionalProperties, "productTierName")
		delete(additionalProperties, "productTierTenancyType")
		delete(additionalProperties, "regionName")
		delete(additionalProperties, "serviceEnvironmentID")
		delete(additionalProperties, "serviceID")
		delete(additionalProperties, "subscriptionID")
		delete(additionalProperties, "totalCost")
		delete(additionalProperties, "utilization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePerInstanceCost struct {
	value *PerInstanceCost
	isSet bool
}

func (v NullablePerInstanceCost) Get() *PerInstanceCost {
	return v.value
}

func (v *NullablePerInstanceCost) Set(val *PerInstanceCost) {
	v.value = val
	v.isSet = true
}

func (v NullablePerInstanceCost) IsSet() bool {
	return v.isSet
}

func (v *NullablePerInstanceCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerInstanceCost(val *PerInstanceCost) *NullablePerInstanceCost {
	return &NullablePerInstanceCost{value: val, isSet: true}
}

func (v NullablePerInstanceCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerInstanceCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


