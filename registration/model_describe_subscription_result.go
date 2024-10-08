/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DescribeSubscriptionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeSubscriptionResult{}

// DescribeSubscriptionResult struct for DescribeSubscriptionResult
type DescribeSubscriptionResult struct {
	// The org ID of the account config identity
	AccountConfigIdentityId string `json:"accountConfigIdentityId"`
	// List of cloud provider names
	CloudProviderNames []string `json:"cloudProviderNames"`
	// The time that this subscription was created
	CreatedAt string `json:"createdAt"`
	// Whether this is the default subscription for the user
	DefaultSubscription bool `json:"defaultSubscription"`
	// The subscription ID
	Id string `json:"id"`
	// The product tier ID that this subscription is tied to
	ProductTierId string `json:"productTierId"`
	// The name of the product tier
	ProductTierName string `json:"productTierName"`
	// The role type of the caller user in this subscription
	RoleType string `json:"roleType"`
	// The user ID that this subscription belong to
	RootUserId string `json:"rootUserId"`
	// The service ID that this subscription is tied to
	ServiceId string `json:"serviceId"`
	// The logo for the service
	ServiceLogoURL string `json:"serviceLogoURL"`
	// The name of the service
	ServiceName string `json:"serviceName"`
	// The org ID of the org that owns the service
	ServiceOrgId string `json:"serviceOrgId"`
	// The name of the org that owns the service
	ServiceOrgName string `json:"serviceOrgName"`
	// The status of the subscription
	Status string `json:"status"`
	// The name of the subscription owner user
	SubscriptionOwnerName string `json:"subscriptionOwnerName"`
}

type _DescribeSubscriptionResult DescribeSubscriptionResult

// NewDescribeSubscriptionResult instantiates a new DescribeSubscriptionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeSubscriptionResult(accountConfigIdentityId string, cloudProviderNames []string, createdAt string, defaultSubscription bool, id string, productTierId string, productTierName string, roleType string, rootUserId string, serviceId string, serviceLogoURL string, serviceName string, serviceOrgId string, serviceOrgName string, status string, subscriptionOwnerName string) *DescribeSubscriptionResult {
	this := DescribeSubscriptionResult{}
	this.AccountConfigIdentityId = accountConfigIdentityId
	this.CloudProviderNames = cloudProviderNames
	this.CreatedAt = createdAt
	this.DefaultSubscription = defaultSubscription
	this.Id = id
	this.ProductTierId = productTierId
	this.ProductTierName = productTierName
	this.RoleType = roleType
	this.RootUserId = rootUserId
	this.ServiceId = serviceId
	this.ServiceLogoURL = serviceLogoURL
	this.ServiceName = serviceName
	this.ServiceOrgId = serviceOrgId
	this.ServiceOrgName = serviceOrgName
	this.Status = status
	this.SubscriptionOwnerName = subscriptionOwnerName
	return &this
}

// NewDescribeSubscriptionResultWithDefaults instantiates a new DescribeSubscriptionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeSubscriptionResultWithDefaults() *DescribeSubscriptionResult {
	this := DescribeSubscriptionResult{}
	return &this
}

// GetAccountConfigIdentityId returns the AccountConfigIdentityId field value
func (o *DescribeSubscriptionResult) GetAccountConfigIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountConfigIdentityId
}

// GetAccountConfigIdentityIdOk returns a tuple with the AccountConfigIdentityId field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetAccountConfigIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountConfigIdentityId, true
}

// SetAccountConfigIdentityId sets field value
func (o *DescribeSubscriptionResult) SetAccountConfigIdentityId(v string) {
	o.AccountConfigIdentityId = v
}

// GetCloudProviderNames returns the CloudProviderNames field value
func (o *DescribeSubscriptionResult) GetCloudProviderNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CloudProviderNames
}

// GetCloudProviderNamesOk returns a tuple with the CloudProviderNames field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetCloudProviderNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudProviderNames, true
}

// SetCloudProviderNames sets field value
func (o *DescribeSubscriptionResult) SetCloudProviderNames(v []string) {
	o.CloudProviderNames = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DescribeSubscriptionResult) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DescribeSubscriptionResult) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDefaultSubscription returns the DefaultSubscription field value
func (o *DescribeSubscriptionResult) GetDefaultSubscription() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultSubscription
}

// GetDefaultSubscriptionOk returns a tuple with the DefaultSubscription field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetDefaultSubscriptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSubscription, true
}

// SetDefaultSubscription sets field value
func (o *DescribeSubscriptionResult) SetDefaultSubscription(v bool) {
	o.DefaultSubscription = v
}

// GetId returns the Id field value
func (o *DescribeSubscriptionResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeSubscriptionResult) SetId(v string) {
	o.Id = v
}

// GetProductTierId returns the ProductTierId field value
func (o *DescribeSubscriptionResult) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *DescribeSubscriptionResult) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetProductTierName returns the ProductTierName field value
func (o *DescribeSubscriptionResult) GetProductTierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierName
}

// GetProductTierNameOk returns a tuple with the ProductTierName field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetProductTierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierName, true
}

// SetProductTierName sets field value
func (o *DescribeSubscriptionResult) SetProductTierName(v string) {
	o.ProductTierName = v
}

// GetRoleType returns the RoleType field value
func (o *DescribeSubscriptionResult) GetRoleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetRoleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleType, true
}

// SetRoleType sets field value
func (o *DescribeSubscriptionResult) SetRoleType(v string) {
	o.RoleType = v
}

// GetRootUserId returns the RootUserId field value
func (o *DescribeSubscriptionResult) GetRootUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootUserId
}

// GetRootUserIdOk returns a tuple with the RootUserId field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetRootUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootUserId, true
}

// SetRootUserId sets field value
func (o *DescribeSubscriptionResult) SetRootUserId(v string) {
	o.RootUserId = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeSubscriptionResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeSubscriptionResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetServiceLogoURL returns the ServiceLogoURL field value
func (o *DescribeSubscriptionResult) GetServiceLogoURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceLogoURL
}

// GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetServiceLogoURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceLogoURL, true
}

// SetServiceLogoURL sets field value
func (o *DescribeSubscriptionResult) SetServiceLogoURL(v string) {
	o.ServiceLogoURL = v
}

// GetServiceName returns the ServiceName field value
func (o *DescribeSubscriptionResult) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *DescribeSubscriptionResult) SetServiceName(v string) {
	o.ServiceName = v
}

// GetServiceOrgId returns the ServiceOrgId field value
func (o *DescribeSubscriptionResult) GetServiceOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceOrgId
}

// GetServiceOrgIdOk returns a tuple with the ServiceOrgId field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetServiceOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceOrgId, true
}

// SetServiceOrgId sets field value
func (o *DescribeSubscriptionResult) SetServiceOrgId(v string) {
	o.ServiceOrgId = v
}

// GetServiceOrgName returns the ServiceOrgName field value
func (o *DescribeSubscriptionResult) GetServiceOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceOrgName
}

// GetServiceOrgNameOk returns a tuple with the ServiceOrgName field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetServiceOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceOrgName, true
}

// SetServiceOrgName sets field value
func (o *DescribeSubscriptionResult) SetServiceOrgName(v string) {
	o.ServiceOrgName = v
}

// GetStatus returns the Status field value
func (o *DescribeSubscriptionResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DescribeSubscriptionResult) SetStatus(v string) {
	o.Status = v
}

// GetSubscriptionOwnerName returns the SubscriptionOwnerName field value
func (o *DescribeSubscriptionResult) GetSubscriptionOwnerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionOwnerName
}

// GetSubscriptionOwnerNameOk returns a tuple with the SubscriptionOwnerName field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionResult) GetSubscriptionOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionOwnerName, true
}

// SetSubscriptionOwnerName sets field value
func (o *DescribeSubscriptionResult) SetSubscriptionOwnerName(v string) {
	o.SubscriptionOwnerName = v
}

func (o DescribeSubscriptionResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeSubscriptionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountConfigIdentityId"] = o.AccountConfigIdentityId
	toSerialize["cloudProviderNames"] = o.CloudProviderNames
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["defaultSubscription"] = o.DefaultSubscription
	toSerialize["id"] = o.Id
	toSerialize["productTierId"] = o.ProductTierId
	toSerialize["productTierName"] = o.ProductTierName
	toSerialize["roleType"] = o.RoleType
	toSerialize["rootUserId"] = o.RootUserId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["serviceLogoURL"] = o.ServiceLogoURL
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["serviceOrgId"] = o.ServiceOrgId
	toSerialize["serviceOrgName"] = o.ServiceOrgName
	toSerialize["status"] = o.Status
	toSerialize["subscriptionOwnerName"] = o.SubscriptionOwnerName
	return toSerialize, nil
}

func (o *DescribeSubscriptionResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountConfigIdentityId",
		"cloudProviderNames",
		"createdAt",
		"defaultSubscription",
		"id",
		"productTierId",
		"productTierName",
		"roleType",
		"rootUserId",
		"serviceId",
		"serviceLogoURL",
		"serviceName",
		"serviceOrgId",
		"serviceOrgName",
		"status",
		"subscriptionOwnerName",
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

	varDescribeSubscriptionResult := _DescribeSubscriptionResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeSubscriptionResult)

	if err != nil {
		return err
	}

	*o = DescribeSubscriptionResult(varDescribeSubscriptionResult)

	return err
}

type NullableDescribeSubscriptionResult struct {
	value *DescribeSubscriptionResult
	isSet bool
}

func (v NullableDescribeSubscriptionResult) Get() *DescribeSubscriptionResult {
	return v.value
}

func (v *NullableDescribeSubscriptionResult) Set(val *DescribeSubscriptionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeSubscriptionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeSubscriptionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeSubscriptionResult(val *DescribeSubscriptionResult) *NullableDescribeSubscriptionResult {
	return &NullableDescribeSubscriptionResult{value: val, isSet: true}
}

func (v NullableDescribeSubscriptionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeSubscriptionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


