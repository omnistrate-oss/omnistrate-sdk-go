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

// checks if the UserSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSubscription{}

// UserSubscription struct for UserSubscription
type UserSubscription struct {
	// List of cloud provider names
	CloudProviderNames []string `json:"cloudProviderNames,omitempty"`
	// Whether this is the default subscription for the user
	DefaultSubscription *bool `json:"defaultSubscription,omitempty"`
	// [DEPRECATED] The email of the user
	Email *string `json:"email,omitempty"`
	// The number of active instances in the subscription
	InstanceCount *int64 `json:"instanceCount,omitempty"`
	// [DEPRECATED] The name of the user
	Name *string `json:"name,omitempty"`
	// ID of a Product Tier
	ProductTierId *string `json:"productTierId,omitempty"`
	// The name of the product tier
	ProductTierName *string `json:"productTierName,omitempty"`
	// Type of the role
	RoleType *string `json:"roleType,omitempty"`
	// ID of a Service Environment
	ServiceEnvironmentId *string `json:"serviceEnvironmentId,omitempty"`
	// ID of a Service
	ServiceId *string `json:"serviceId,omitempty"`
	// The logo for the service
	ServiceLogoURL *string `json:"serviceLogoURL,omitempty"`
	// The name of the service
	ServiceName *string `json:"serviceName,omitempty"`
	// The date the user joined the subscription
	SubscriptionDate *string `json:"subscriptionDate,omitempty"`
	// ID of a Subscription
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// The name of the subscription owner user
	SubscriptionOwnerName *string `json:"subscriptionOwnerName,omitempty"`
	// [DEPRECATED] The User ID
	UserId *string `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSubscription UserSubscription

// NewUserSubscription instantiates a new UserSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSubscription() *UserSubscription {
	this := UserSubscription{}
	return &this
}

// NewUserSubscriptionWithDefaults instantiates a new UserSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSubscriptionWithDefaults() *UserSubscription {
	this := UserSubscription{}
	return &this
}

// GetCloudProviderNames returns the CloudProviderNames field value if set, zero value otherwise.
func (o *UserSubscription) GetCloudProviderNames() []string {
	if o == nil || IsNil(o.CloudProviderNames) {
		var ret []string
		return ret
	}
	return o.CloudProviderNames
}

// GetCloudProviderNamesOk returns a tuple with the CloudProviderNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetCloudProviderNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CloudProviderNames) {
		return nil, false
	}
	return o.CloudProviderNames, true
}

// HasCloudProviderNames returns a boolean if a field has been set.
func (o *UserSubscription) HasCloudProviderNames() bool {
	if o != nil && !IsNil(o.CloudProviderNames) {
		return true
	}

	return false
}

// SetCloudProviderNames gets a reference to the given []string and assigns it to the CloudProviderNames field.
func (o *UserSubscription) SetCloudProviderNames(v []string) {
	o.CloudProviderNames = v
}

// GetDefaultSubscription returns the DefaultSubscription field value if set, zero value otherwise.
func (o *UserSubscription) GetDefaultSubscription() bool {
	if o == nil || IsNil(o.DefaultSubscription) {
		var ret bool
		return ret
	}
	return *o.DefaultSubscription
}

// GetDefaultSubscriptionOk returns a tuple with the DefaultSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetDefaultSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultSubscription) {
		return nil, false
	}
	return o.DefaultSubscription, true
}

// HasDefaultSubscription returns a boolean if a field has been set.
func (o *UserSubscription) HasDefaultSubscription() bool {
	if o != nil && !IsNil(o.DefaultSubscription) {
		return true
	}

	return false
}

// SetDefaultSubscription gets a reference to the given bool and assigns it to the DefaultSubscription field.
func (o *UserSubscription) SetDefaultSubscription(v bool) {
	o.DefaultSubscription = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserSubscription) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserSubscription) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserSubscription) SetEmail(v string) {
	o.Email = &v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *UserSubscription) GetInstanceCount() int64 {
	if o == nil || IsNil(o.InstanceCount) {
		var ret int64
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetInstanceCountOk() (*int64, bool) {
	if o == nil || IsNil(o.InstanceCount) {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *UserSubscription) HasInstanceCount() bool {
	if o != nil && !IsNil(o.InstanceCount) {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int64 and assigns it to the InstanceCount field.
func (o *UserSubscription) SetInstanceCount(v int64) {
	o.InstanceCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSubscription) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSubscription) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSubscription) SetName(v string) {
	o.Name = &v
}

// GetProductTierId returns the ProductTierId field value if set, zero value otherwise.
func (o *UserSubscription) GetProductTierId() string {
	if o == nil || IsNil(o.ProductTierId) {
		var ret string
		return ret
	}
	return *o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetProductTierIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierId) {
		return nil, false
	}
	return o.ProductTierId, true
}

// HasProductTierId returns a boolean if a field has been set.
func (o *UserSubscription) HasProductTierId() bool {
	if o != nil && !IsNil(o.ProductTierId) {
		return true
	}

	return false
}

// SetProductTierId gets a reference to the given string and assigns it to the ProductTierId field.
func (o *UserSubscription) SetProductTierId(v string) {
	o.ProductTierId = &v
}

// GetProductTierName returns the ProductTierName field value if set, zero value otherwise.
func (o *UserSubscription) GetProductTierName() string {
	if o == nil || IsNil(o.ProductTierName) {
		var ret string
		return ret
	}
	return *o.ProductTierName
}

// GetProductTierNameOk returns a tuple with the ProductTierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetProductTierNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierName) {
		return nil, false
	}
	return o.ProductTierName, true
}

// HasProductTierName returns a boolean if a field has been set.
func (o *UserSubscription) HasProductTierName() bool {
	if o != nil && !IsNil(o.ProductTierName) {
		return true
	}

	return false
}

// SetProductTierName gets a reference to the given string and assigns it to the ProductTierName field.
func (o *UserSubscription) SetProductTierName(v string) {
	o.ProductTierName = &v
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *UserSubscription) GetRoleType() string {
	if o == nil || IsNil(o.RoleType) {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetRoleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RoleType) {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *UserSubscription) HasRoleType() bool {
	if o != nil && !IsNil(o.RoleType) {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *UserSubscription) SetRoleType(v string) {
	o.RoleType = &v
}

// GetServiceEnvironmentId returns the ServiceEnvironmentId field value if set, zero value otherwise.
func (o *UserSubscription) GetServiceEnvironmentId() string {
	if o == nil || IsNil(o.ServiceEnvironmentId) {
		var ret string
		return ret
	}
	return *o.ServiceEnvironmentId
}

// GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetServiceEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceEnvironmentId) {
		return nil, false
	}
	return o.ServiceEnvironmentId, true
}

// HasServiceEnvironmentId returns a boolean if a field has been set.
func (o *UserSubscription) HasServiceEnvironmentId() bool {
	if o != nil && !IsNil(o.ServiceEnvironmentId) {
		return true
	}

	return false
}

// SetServiceEnvironmentId gets a reference to the given string and assigns it to the ServiceEnvironmentId field.
func (o *UserSubscription) SetServiceEnvironmentId(v string) {
	o.ServiceEnvironmentId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *UserSubscription) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *UserSubscription) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *UserSubscription) SetServiceId(v string) {
	o.ServiceId = &v
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

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *UserSubscription) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *UserSubscription) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *UserSubscription) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetSubscriptionDate returns the SubscriptionDate field value if set, zero value otherwise.
func (o *UserSubscription) GetSubscriptionDate() string {
	if o == nil || IsNil(o.SubscriptionDate) {
		var ret string
		return ret
	}
	return *o.SubscriptionDate
}

// GetSubscriptionDateOk returns a tuple with the SubscriptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetSubscriptionDateOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionDate) {
		return nil, false
	}
	return o.SubscriptionDate, true
}

// HasSubscriptionDate returns a boolean if a field has been set.
func (o *UserSubscription) HasSubscriptionDate() bool {
	if o != nil && !IsNil(o.SubscriptionDate) {
		return true
	}

	return false
}

// SetSubscriptionDate gets a reference to the given string and assigns it to the SubscriptionDate field.
func (o *UserSubscription) SetSubscriptionDate(v string) {
	o.SubscriptionDate = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UserSubscription) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UserSubscription) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UserSubscription) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetSubscriptionOwnerName returns the SubscriptionOwnerName field value if set, zero value otherwise.
func (o *UserSubscription) GetSubscriptionOwnerName() string {
	if o == nil || IsNil(o.SubscriptionOwnerName) {
		var ret string
		return ret
	}
	return *o.SubscriptionOwnerName
}

// GetSubscriptionOwnerNameOk returns a tuple with the SubscriptionOwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetSubscriptionOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionOwnerName) {
		return nil, false
	}
	return o.SubscriptionOwnerName, true
}

// HasSubscriptionOwnerName returns a boolean if a field has been set.
func (o *UserSubscription) HasSubscriptionOwnerName() bool {
	if o != nil && !IsNil(o.SubscriptionOwnerName) {
		return true
	}

	return false
}

// SetSubscriptionOwnerName gets a reference to the given string and assigns it to the SubscriptionOwnerName field.
func (o *UserSubscription) SetSubscriptionOwnerName(v string) {
	o.SubscriptionOwnerName = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserSubscription) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSubscription) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserSubscription) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserSubscription) SetUserId(v string) {
	o.UserId = &v
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
	if !IsNil(o.CloudProviderNames) {
		toSerialize["cloudProviderNames"] = o.CloudProviderNames
	}
	if !IsNil(o.DefaultSubscription) {
		toSerialize["defaultSubscription"] = o.DefaultSubscription
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.InstanceCount) {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProductTierId) {
		toSerialize["productTierId"] = o.ProductTierId
	}
	if !IsNil(o.ProductTierName) {
		toSerialize["productTierName"] = o.ProductTierName
	}
	if !IsNil(o.RoleType) {
		toSerialize["roleType"] = o.RoleType
	}
	if !IsNil(o.ServiceEnvironmentId) {
		toSerialize["serviceEnvironmentId"] = o.ServiceEnvironmentId
	}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.ServiceLogoURL) {
		toSerialize["serviceLogoURL"] = o.ServiceLogoURL
	}
	if !IsNil(o.ServiceName) {
		toSerialize["serviceName"] = o.ServiceName
	}
	if !IsNil(o.SubscriptionDate) {
		toSerialize["subscriptionDate"] = o.SubscriptionDate
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.SubscriptionOwnerName) {
		toSerialize["subscriptionOwnerName"] = o.SubscriptionOwnerName
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSubscription) UnmarshalJSON(data []byte) (err error) {
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


