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

// checks if the UserSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSubscription{}

// UserSubscription struct for UserSubscription
type UserSubscription struct {
	// List of cloud provider names
	CloudProviderNames []string `json:"cloudProviderNames"`
	// Whether this is the default subscription for the user
	DefaultSubscription bool `json:"defaultSubscription"`
	// The email of the user
	Email string `json:"email"`
	// The number of active instances in the subscription
	InstanceCount int64 `json:"instanceCount"`
	// The name of the user
	Name string `json:"name"`
	// ID of a Product Tier
	ProductTierId string `json:"productTierId"`
	// The name of the product tier
	ProductTierName string `json:"productTierName"`
	// Type of the role
	RoleType string `json:"roleType"`
	// ID of a Service Environment
	ServiceEnvironmentId string `json:"serviceEnvironmentId"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// The logo for the service
	ServiceLogoURL *string `json:"serviceLogoURL,omitempty"`
	// The name of the service
	ServiceName string `json:"serviceName"`
	// The date the user joined the subscription
	SubscriptionDate string `json:"subscriptionDate"`
	// ID of a Subscription
	SubscriptionId string `json:"subscriptionId"`
	// The name of the subscription owner user
	SubscriptionOwnerName string `json:"subscriptionOwnerName"`
	// The User ID
	UserId string `json:"userId"`
	AdditionalProperties map[string]interface{}
}

type _UserSubscription UserSubscription

// NewUserSubscription instantiates a new UserSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSubscription(cloudProviderNames []string, defaultSubscription bool, email string, instanceCount int64, name string, productTierId string, productTierName string, roleType string, serviceEnvironmentId string, serviceId string, serviceName string, subscriptionDate string, subscriptionId string, subscriptionOwnerName string, userId string) *UserSubscription {
	this := UserSubscription{}
	this.CloudProviderNames = cloudProviderNames
	this.DefaultSubscription = defaultSubscription
	this.Email = email
	this.InstanceCount = instanceCount
	this.Name = name
	this.ProductTierId = productTierId
	this.ProductTierName = productTierName
	this.RoleType = roleType
	this.ServiceEnvironmentId = serviceEnvironmentId
	this.ServiceId = serviceId
	this.ServiceName = serviceName
	this.SubscriptionDate = subscriptionDate
	this.SubscriptionId = subscriptionId
	this.SubscriptionOwnerName = subscriptionOwnerName
	this.UserId = userId
	return &this
}

// NewUserSubscriptionWithDefaults instantiates a new UserSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSubscriptionWithDefaults() *UserSubscription {
	this := UserSubscription{}
	return &this
}

// GetCloudProviderNames returns the CloudProviderNames field value
func (o *UserSubscription) GetCloudProviderNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CloudProviderNames
}

// GetCloudProviderNamesOk returns a tuple with the CloudProviderNames field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetCloudProviderNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudProviderNames, true
}

// SetCloudProviderNames sets field value
func (o *UserSubscription) SetCloudProviderNames(v []string) {
	o.CloudProviderNames = v
}

// GetDefaultSubscription returns the DefaultSubscription field value
func (o *UserSubscription) GetDefaultSubscription() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultSubscription
}

// GetDefaultSubscriptionOk returns a tuple with the DefaultSubscription field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetDefaultSubscriptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSubscription, true
}

// SetDefaultSubscription sets field value
func (o *UserSubscription) SetDefaultSubscription(v bool) {
	o.DefaultSubscription = v
}

// GetEmail returns the Email field value
func (o *UserSubscription) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserSubscription) SetEmail(v string) {
	o.Email = v
}

// GetInstanceCount returns the InstanceCount field value
func (o *UserSubscription) GetInstanceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetInstanceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value
func (o *UserSubscription) SetInstanceCount(v int64) {
	o.InstanceCount = v
}

// GetName returns the Name field value
func (o *UserSubscription) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserSubscription) SetName(v string) {
	o.Name = v
}

// GetProductTierId returns the ProductTierId field value
func (o *UserSubscription) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *UserSubscription) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetProductTierName returns the ProductTierName field value
func (o *UserSubscription) GetProductTierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierName
}

// GetProductTierNameOk returns a tuple with the ProductTierName field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetProductTierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierName, true
}

// SetProductTierName sets field value
func (o *UserSubscription) SetProductTierName(v string) {
	o.ProductTierName = v
}

// GetRoleType returns the RoleType field value
func (o *UserSubscription) GetRoleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetRoleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleType, true
}

// SetRoleType sets field value
func (o *UserSubscription) SetRoleType(v string) {
	o.RoleType = v
}

// GetServiceEnvironmentId returns the ServiceEnvironmentId field value
func (o *UserSubscription) GetServiceEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceEnvironmentId
}

// GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetServiceEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceEnvironmentId, true
}

// SetServiceEnvironmentId sets field value
func (o *UserSubscription) SetServiceEnvironmentId(v string) {
	o.ServiceEnvironmentId = v
}

// GetServiceId returns the ServiceId field value
func (o *UserSubscription) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *UserSubscription) SetServiceId(v string) {
	o.ServiceId = v
}

// GetServiceLogoURL returns the ServiceLogoURL field value if set, zero value otherwise.
func (o *UserSubscription) GetServiceLogoURL() string {
	if o == nil || IsNil(o.ServiceLogoURL) {
		var ret string
		return ret
	}
	return *o.ServiceLogoURL
}

// GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetServiceLogoURLOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLogoURL) {
		return nil, false
	}
	return o.ServiceLogoURL, true
}

// HasServiceLogoURL returns a boolean if a field has been set.
func (o *UserSubscription) HasServiceLogoURL() bool {
	if o != nil && !IsNil(o.ServiceLogoURL) {
		return true
	}

	return false
}

// SetServiceLogoURL gets a reference to the given string and assigns it to the ServiceLogoURL field.
func (o *UserSubscription) SetServiceLogoURL(v string) {
	o.ServiceLogoURL = &v
}

// GetServiceName returns the ServiceName field value
func (o *UserSubscription) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *UserSubscription) SetServiceName(v string) {
	o.ServiceName = v
}

// GetSubscriptionDate returns the SubscriptionDate field value
func (o *UserSubscription) GetSubscriptionDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionDate
}

// GetSubscriptionDateOk returns a tuple with the SubscriptionDate field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetSubscriptionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionDate, true
}

// SetSubscriptionDate sets field value
func (o *UserSubscription) SetSubscriptionDate(v string) {
	o.SubscriptionDate = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UserSubscription) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UserSubscription) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetSubscriptionOwnerName returns the SubscriptionOwnerName field value
func (o *UserSubscription) GetSubscriptionOwnerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionOwnerName
}

// GetSubscriptionOwnerNameOk returns a tuple with the SubscriptionOwnerName field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetSubscriptionOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionOwnerName, true
}

// SetSubscriptionOwnerName sets field value
func (o *UserSubscription) SetSubscriptionOwnerName(v string) {
	o.SubscriptionOwnerName = v
}

// GetUserId returns the UserId field value
func (o *UserSubscription) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserSubscription) SetUserId(v string) {
	o.UserId = v
}

func (o UserSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProviderNames"] = o.CloudProviderNames
	toSerialize["defaultSubscription"] = o.DefaultSubscription
	toSerialize["email"] = o.Email
	toSerialize["instanceCount"] = o.InstanceCount
	toSerialize["name"] = o.Name
	toSerialize["productTierId"] = o.ProductTierId
	toSerialize["productTierName"] = o.ProductTierName
	toSerialize["roleType"] = o.RoleType
	toSerialize["serviceEnvironmentId"] = o.ServiceEnvironmentId
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.ServiceLogoURL) {
		toSerialize["serviceLogoURL"] = o.ServiceLogoURL
	}
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["subscriptionDate"] = o.SubscriptionDate
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["subscriptionOwnerName"] = o.SubscriptionOwnerName
	toSerialize["userId"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSubscription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderNames",
		"defaultSubscription",
		"email",
		"instanceCount",
		"name",
		"productTierId",
		"productTierName",
		"roleType",
		"serviceEnvironmentId",
		"serviceId",
		"serviceName",
		"subscriptionDate",
		"subscriptionId",
		"subscriptionOwnerName",
		"userId",
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

	varUserSubscription := _UserSubscription{}

	err = json.Unmarshal(data, &varUserSubscription)

	if err != nil {
		return err
	}

	*o = UserSubscription(varUserSubscription)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloudProviderNames")
		delete(additionalProperties, "defaultSubscription")
		delete(additionalProperties, "email")
		delete(additionalProperties, "instanceCount")
		delete(additionalProperties, "name")
		delete(additionalProperties, "productTierId")
		delete(additionalProperties, "productTierName")
		delete(additionalProperties, "roleType")
		delete(additionalProperties, "serviceEnvironmentId")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "serviceLogoURL")
		delete(additionalProperties, "serviceName")
		delete(additionalProperties, "subscriptionDate")
		delete(additionalProperties, "subscriptionId")
		delete(additionalProperties, "subscriptionOwnerName")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserSubscription struct {
	value *UserSubscription
	isSet bool
}

func (v NullableUserSubscription) Get() *UserSubscription {
	return v.value
}

func (v *NullableUserSubscription) Set(val *UserSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSubscription(val *UserSubscription) *NullableUserSubscription {
	return &NullableUserSubscription{value: val, isSet: true}
}

func (v NullableUserSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


