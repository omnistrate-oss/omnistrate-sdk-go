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

// checks if the ResourceInstanceSearchRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceInstanceSearchRecord{}

// ResourceInstanceSearchRecord struct for ResourceInstanceSearchRecord
type ResourceInstanceSearchRecord struct {
	// Name of the Infra Provider
	CloudProvider string `json:"cloudProvider"`
	// The instance description.
	Description string `json:"description"`
	// The resource instance ID.
	Id string `json:"id"`
	// Is the proxy managed by Omnistrate.
	Managed *bool `json:"managed,omitempty"`
	// The managed resource type of the proxy instance.
	ManagedResourceType *string `json:"managedResourceType,omitempty"`
	// The ports registration status of the ports based proxy instance.
	PortsRegistrationStatus *map[string][]int64 `json:"portsRegistrationStatus,omitempty"`
	// ID of a Product Tier
	ProductTierId string `json:"productTierId"`
	// The product tier name of the instance.
	ProductTierName *string `json:"productTierName,omitempty"`
	// The product tier version of the instance.
	ProductTierVersion *string `json:"productTierVersion,omitempty"`
	// The proxy type.
	ProxyType *string `json:"proxyType,omitempty"`
	// The region code where the instance is hosted.
	RegionCode string `json:"regionCode"`
	// ID of a resource
	ResourceId *string `json:"resourceId,omitempty"`
	// The name of the resource for the instance.
	ResourceName string `json:"resourceName"`
	// ID of a Service Environment
	ServiceEnvironmentId string `json:"serviceEnvironmentId"`
	// The service environment name of the instance.
	ServiceEnvironmentName string `json:"serviceEnvironmentName"`
	// The type of service environment
	ServiceEnvironmentType *string `json:"serviceEnvironmentType,omitempty"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// The service name of the instance.
	ServiceName string `json:"serviceName"`
	// The status of an operation
	Status string `json:"status"`
	// The instance status description.
	StatusDescription string `json:"statusDescription"`
	// ID of a Subscription
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceInstanceSearchRecord ResourceInstanceSearchRecord

// NewResourceInstanceSearchRecord instantiates a new ResourceInstanceSearchRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceInstanceSearchRecord(cloudProvider string, description string, id string, productTierId string, regionCode string, resourceName string, serviceEnvironmentId string, serviceEnvironmentName string, serviceId string, serviceName string, status string, statusDescription string) *ResourceInstanceSearchRecord {
	this := ResourceInstanceSearchRecord{}
	this.CloudProvider = cloudProvider
	this.Description = description
	this.Id = id
	this.ProductTierId = productTierId
	this.RegionCode = regionCode
	this.ResourceName = resourceName
	this.ServiceEnvironmentId = serviceEnvironmentId
	this.ServiceEnvironmentName = serviceEnvironmentName
	this.ServiceId = serviceId
	this.ServiceName = serviceName
	this.Status = status
	this.StatusDescription = statusDescription
	return &this
}

// NewResourceInstanceSearchRecordWithDefaults instantiates a new ResourceInstanceSearchRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceInstanceSearchRecordWithDefaults() *ResourceInstanceSearchRecord {
	this := ResourceInstanceSearchRecord{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *ResourceInstanceSearchRecord) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ResourceInstanceSearchRecord) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetDescription returns the Description field value
func (o *ResourceInstanceSearchRecord) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ResourceInstanceSearchRecord) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *ResourceInstanceSearchRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceInstanceSearchRecord) SetId(v string) {
	o.Id = v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *ResourceInstanceSearchRecord) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *ResourceInstanceSearchRecord) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *ResourceInstanceSearchRecord) SetManaged(v bool) {
	o.Managed = &v
}

// GetManagedResourceType returns the ManagedResourceType field value if set, zero value otherwise.
func (o *ResourceInstanceSearchRecord) GetManagedResourceType() string {
	if o == nil || IsNil(o.ManagedResourceType) {
		var ret string
		return ret
	}
	return *o.ManagedResourceType
}

// GetManagedResourceTypeOk returns a tuple with the ManagedResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetManagedResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedResourceType) {
		return nil, false
	}
	return o.ManagedResourceType, true
}

// HasManagedResourceType returns a boolean if a field has been set.
func (o *ResourceInstanceSearchRecord) HasManagedResourceType() bool {
	if o != nil && !IsNil(o.ManagedResourceType) {
		return true
	}

	return false
}

// SetManagedResourceType gets a reference to the given string and assigns it to the ManagedResourceType field.
func (o *ResourceInstanceSearchRecord) SetManagedResourceType(v string) {
	o.ManagedResourceType = &v
}

// GetPortsRegistrationStatus returns the PortsRegistrationStatus field value if set, zero value otherwise.
func (o *ResourceInstanceSearchRecord) GetPortsRegistrationStatus() map[string][]int64 {
	if o == nil || IsNil(o.PortsRegistrationStatus) {
		var ret map[string][]int64
		return ret
	}
	return *o.PortsRegistrationStatus
}

// GetPortsRegistrationStatusOk returns a tuple with the PortsRegistrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetPortsRegistrationStatusOk() (*map[string][]int64, bool) {
	if o == nil || IsNil(o.PortsRegistrationStatus) {
		return nil, false
	}
	return o.PortsRegistrationStatus, true
}

// HasPortsRegistrationStatus returns a boolean if a field has been set.
func (o *ResourceInstanceSearchRecord) HasPortsRegistrationStatus() bool {
	if o != nil && !IsNil(o.PortsRegistrationStatus) {
		return true
	}

	return false
}

// SetPortsRegistrationStatus gets a reference to the given map[string][]int64 and assigns it to the PortsRegistrationStatus field.
func (o *ResourceInstanceSearchRecord) SetPortsRegistrationStatus(v map[string][]int64) {
	o.PortsRegistrationStatus = &v
}

// GetProductTierId returns the ProductTierId field value
func (o *ResourceInstanceSearchRecord) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *ResourceInstanceSearchRecord) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetProductTierName returns the ProductTierName field value if set, zero value otherwise.
func (o *ResourceInstanceSearchRecord) GetProductTierName() string {
	if o == nil || IsNil(o.ProductTierName) {
		var ret string
		return ret
	}
	return *o.ProductTierName
}

// GetProductTierNameOk returns a tuple with the ProductTierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetProductTierNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierName) {
		return nil, false
	}
	return o.ProductTierName, true
}

// HasProductTierName returns a boolean if a field has been set.
func (o *ResourceInstanceSearchRecord) HasProductTierName() bool {
	if o != nil && !IsNil(o.ProductTierName) {
		return true
	}

	return false
}

// SetProductTierName gets a reference to the given string and assigns it to the ProductTierName field.
func (o *ResourceInstanceSearchRecord) SetProductTierName(v string) {
	o.ProductTierName = &v
}

// GetProductTierVersion returns the ProductTierVersion field value if set, zero value otherwise.
func (o *ResourceInstanceSearchRecord) GetProductTierVersion() string {
	if o == nil || IsNil(o.ProductTierVersion) {
		var ret string
		return ret
	}
	return *o.ProductTierVersion
}

// GetProductTierVersionOk returns a tuple with the ProductTierVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetProductTierVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierVersion) {
		return nil, false
	}
	return o.ProductTierVersion, true
}

// HasProductTierVersion returns a boolean if a field has been set.
func (o *ResourceInstanceSearchRecord) HasProductTierVersion() bool {
	if o != nil && !IsNil(o.ProductTierVersion) {
		return true
	}

	return false
}

// SetProductTierVersion gets a reference to the given string and assigns it to the ProductTierVersion field.
func (o *ResourceInstanceSearchRecord) SetProductTierVersion(v string) {
	o.ProductTierVersion = &v
}

// GetProxyType returns the ProxyType field value if set, zero value otherwise.
func (o *ResourceInstanceSearchRecord) GetProxyType() string {
	if o == nil || IsNil(o.ProxyType) {
		var ret string
		return ret
	}
	return *o.ProxyType
}

// GetProxyTypeOk returns a tuple with the ProxyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetProxyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyType) {
		return nil, false
	}
	return o.ProxyType, true
}

// HasProxyType returns a boolean if a field has been set.
func (o *ResourceInstanceSearchRecord) HasProxyType() bool {
	if o != nil && !IsNil(o.ProxyType) {
		return true
	}

	return false
}

// SetProxyType gets a reference to the given string and assigns it to the ProxyType field.
func (o *ResourceInstanceSearchRecord) SetProxyType(v string) {
	o.ProxyType = &v
}

// GetRegionCode returns the RegionCode field value
func (o *ResourceInstanceSearchRecord) GetRegionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetRegionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionCode, true
}

// SetRegionCode sets field value
func (o *ResourceInstanceSearchRecord) SetRegionCode(v string) {
	o.RegionCode = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *ResourceInstanceSearchRecord) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *ResourceInstanceSearchRecord) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *ResourceInstanceSearchRecord) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceName returns the ResourceName field value
func (o *ResourceInstanceSearchRecord) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *ResourceInstanceSearchRecord) SetResourceName(v string) {
	o.ResourceName = v
}

// GetServiceEnvironmentId returns the ServiceEnvironmentId field value
func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentId
}

// GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentId, true
}

// SetServiceEnvironmentId sets field value
func (o *ResourceInstanceSearchRecord) SetServiceEnvironmentId(v string) {
	o.ServiceEnvironmentId = v
}

// GetServiceEnvironmentName returns the ServiceEnvironmentName field value
func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentName
}

// GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentName, true
}

// SetServiceEnvironmentName sets field value
func (o *ResourceInstanceSearchRecord) SetServiceEnvironmentName(v string) {
	o.ServiceEnvironmentName = v
}

// GetServiceEnvironmentType returns the ServiceEnvironmentType field value if set, zero value otherwise.
func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentType() string {
	if o == nil || IsNil(o.ServiceEnvironmentType) {
		var ret string
		return ret
	}
	return *o.ServiceEnvironmentType
}

// GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceEnvironmentType) {
		return nil, false
	}
	return o.ServiceEnvironmentType, true
}

// HasServiceEnvironmentType returns a boolean if a field has been set.
func (o *ResourceInstanceSearchRecord) HasServiceEnvironmentType() bool {
	if o != nil && !IsNil(o.ServiceEnvironmentType) {
		return true
	}

	return false
}

// SetServiceEnvironmentType gets a reference to the given string and assigns it to the ServiceEnvironmentType field.
func (o *ResourceInstanceSearchRecord) SetServiceEnvironmentType(v string) {
	o.ServiceEnvironmentType = &v
}

// GetServiceId returns the ServiceId field value
func (o *ResourceInstanceSearchRecord) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ResourceInstanceSearchRecord) SetServiceId(v string) {
	o.ServiceId = v
}

// GetServiceName returns the ServiceName field value
func (o *ResourceInstanceSearchRecord) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *ResourceInstanceSearchRecord) SetServiceName(v string) {
	o.ServiceName = v
}

// GetStatus returns the Status field value
func (o *ResourceInstanceSearchRecord) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResourceInstanceSearchRecord) SetStatus(v string) {
	o.Status = v
}

// GetStatusDescription returns the StatusDescription field value
func (o *ResourceInstanceSearchRecord) GetStatusDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetStatusDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDescription, true
}

// SetStatusDescription sets field value
func (o *ResourceInstanceSearchRecord) SetStatusDescription(v string) {
	o.StatusDescription = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *ResourceInstanceSearchRecord) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInstanceSearchRecord) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *ResourceInstanceSearchRecord) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *ResourceInstanceSearchRecord) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o ResourceInstanceSearchRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceInstanceSearchRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProvider"] = o.CloudProvider
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	if !IsNil(o.ManagedResourceType) {
		toSerialize["managedResourceType"] = o.ManagedResourceType
	}
	if !IsNil(o.PortsRegistrationStatus) {
		toSerialize["portsRegistrationStatus"] = o.PortsRegistrationStatus
	}
	toSerialize["productTierId"] = o.ProductTierId
	if !IsNil(o.ProductTierName) {
		toSerialize["productTierName"] = o.ProductTierName
	}
	if !IsNil(o.ProductTierVersion) {
		toSerialize["productTierVersion"] = o.ProductTierVersion
	}
	if !IsNil(o.ProxyType) {
		toSerialize["proxyType"] = o.ProxyType
	}
	toSerialize["regionCode"] = o.RegionCode
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	toSerialize["resourceName"] = o.ResourceName
	toSerialize["serviceEnvironmentId"] = o.ServiceEnvironmentId
	toSerialize["serviceEnvironmentName"] = o.ServiceEnvironmentName
	if !IsNil(o.ServiceEnvironmentType) {
		toSerialize["serviceEnvironmentType"] = o.ServiceEnvironmentType
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["status"] = o.Status
	toSerialize["statusDescription"] = o.StatusDescription
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceInstanceSearchRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProvider",
		"description",
		"id",
		"productTierId",
		"regionCode",
		"resourceName",
		"serviceEnvironmentId",
		"serviceEnvironmentName",
		"serviceId",
		"serviceName",
		"status",
		"statusDescription",
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

	varResourceInstanceSearchRecord := _ResourceInstanceSearchRecord{}

	err = json.Unmarshal(data, &varResourceInstanceSearchRecord)

	if err != nil {
		return err
	}

	*o = ResourceInstanceSearchRecord(varResourceInstanceSearchRecord)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProvider")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "managedResourceType")
		delete(additionalProperties, "portsRegistrationStatus")
		delete(additionalProperties, "productTierId")
		delete(additionalProperties, "productTierName")
		delete(additionalProperties, "productTierVersion")
		delete(additionalProperties, "proxyType")
		delete(additionalProperties, "regionCode")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceName")
		delete(additionalProperties, "serviceEnvironmentId")
		delete(additionalProperties, "serviceEnvironmentName")
		delete(additionalProperties, "serviceEnvironmentType")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "serviceName")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusDescription")
		delete(additionalProperties, "subscriptionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceInstanceSearchRecord struct {
	value *ResourceInstanceSearchRecord
	isSet bool
}

func (v NullableResourceInstanceSearchRecord) Get() *ResourceInstanceSearchRecord {
	return v.value
}

func (v *NullableResourceInstanceSearchRecord) Set(val *ResourceInstanceSearchRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceInstanceSearchRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceInstanceSearchRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceInstanceSearchRecord(val *ResourceInstanceSearchRecord) *NullableResourceInstanceSearchRecord {
	return &NullableResourceInstanceSearchRecord{value: val, isSet: true}
}

func (v NullableResourceInstanceSearchRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceInstanceSearchRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


