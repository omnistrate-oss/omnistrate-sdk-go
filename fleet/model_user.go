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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	// The time the user was created.
	CreatedAt string `json:"createdAt"`
	// The user email.
	Email string `json:"email"`
	// Is the user enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The service environment ID this workflow belongs to.
	EnvironmentId string `json:"environmentId"`
	// The number of active instances the user has.
	InstanceCount int64 `json:"instanceCount"`
	// The last modified time of the user.
	LastModifiedAt string `json:"lastModifiedAt"`
	// The user ID of the last modifier.
	LastModifiedByUserID *string `json:"lastModifiedByUserID,omitempty"`
	// The user name of the last modifier.
	LastModifiedByUserName *string `json:"lastModifiedByUserName,omitempty"`
	// The organization ID.
	OrgId string `json:"orgId"`
	// The organization name.
	OrgName string `json:"orgName"`
	// The organization URL.
	OrgUrl *string `json:"orgUrl,omitempty"`
	// The service ID this workflow belongs to.
	ServiceId string `json:"serviceId"`
	// The status of the user.
	Status string `json:"status"`
	// The number of subscriptions the user has.
	SubscriptionCount *int64 `json:"subscriptionCount,omitempty"`
	// Token to validate the user, if the user is not enabled.
	Token *string `json:"token,omitempty"`
	// The user ID.
	UserId string `json:"userId"`
	// The user name.
	UserName string `json:"userName"`
	// The user subscription role.
	UserSubscriptionRole *string `json:"userSubscriptionRole,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(createdAt string, email string, environmentId string, instanceCount int64, lastModifiedAt string, orgId string, orgName string, serviceId string, status string, userId string, userName string) *User {
	this := User{}
	this.CreatedAt = createdAt
	this.Email = email
	this.EnvironmentId = environmentId
	this.InstanceCount = instanceCount
	this.LastModifiedAt = lastModifiedAt
	this.OrgId = orgId
	this.OrgName = orgName
	this.ServiceId = serviceId
	this.Status = status
	this.UserId = userId
	this.UserName = userName
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *User) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *User) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *User) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *User) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *User) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *User) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *User) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *User) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetInstanceCount returns the InstanceCount field value
func (o *User) GetInstanceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *User) GetInstanceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value
func (o *User) SetInstanceCount(v int64) {
	o.InstanceCount = v
}

// GetLastModifiedAt returns the LastModifiedAt field value
func (o *User) GetLastModifiedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastModifiedAt
}

// GetLastModifiedAtOk returns a tuple with the LastModifiedAt field value
// and a boolean to check if the value has been set.
func (o *User) GetLastModifiedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedAt, true
}

// SetLastModifiedAt sets field value
func (o *User) SetLastModifiedAt(v string) {
	o.LastModifiedAt = v
}

// GetLastModifiedByUserID returns the LastModifiedByUserID field value if set, zero value otherwise.
func (o *User) GetLastModifiedByUserID() string {
	if o == nil || IsNil(o.LastModifiedByUserID) {
		var ret string
		return ret
	}
	return *o.LastModifiedByUserID
}

// GetLastModifiedByUserIDOk returns a tuple with the LastModifiedByUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastModifiedByUserIDOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedByUserID) {
		return nil, false
	}
	return o.LastModifiedByUserID, true
}

// HasLastModifiedByUserID returns a boolean if a field has been set.
func (o *User) HasLastModifiedByUserID() bool {
	if o != nil && !IsNil(o.LastModifiedByUserID) {
		return true
	}

	return false
}

// SetLastModifiedByUserID gets a reference to the given string and assigns it to the LastModifiedByUserID field.
func (o *User) SetLastModifiedByUserID(v string) {
	o.LastModifiedByUserID = &v
}

// GetLastModifiedByUserName returns the LastModifiedByUserName field value if set, zero value otherwise.
func (o *User) GetLastModifiedByUserName() string {
	if o == nil || IsNil(o.LastModifiedByUserName) {
		var ret string
		return ret
	}
	return *o.LastModifiedByUserName
}

// GetLastModifiedByUserNameOk returns a tuple with the LastModifiedByUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastModifiedByUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedByUserName) {
		return nil, false
	}
	return o.LastModifiedByUserName, true
}

// HasLastModifiedByUserName returns a boolean if a field has been set.
func (o *User) HasLastModifiedByUserName() bool {
	if o != nil && !IsNil(o.LastModifiedByUserName) {
		return true
	}

	return false
}

// SetLastModifiedByUserName gets a reference to the given string and assigns it to the LastModifiedByUserName field.
func (o *User) SetLastModifiedByUserName(v string) {
	o.LastModifiedByUserName = &v
}

// GetOrgId returns the OrgId field value
func (o *User) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *User) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *User) SetOrgId(v string) {
	o.OrgId = v
}

// GetOrgName returns the OrgName field value
func (o *User) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *User) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *User) SetOrgName(v string) {
	o.OrgName = v
}

// GetOrgUrl returns the OrgUrl field value if set, zero value otherwise.
func (o *User) GetOrgUrl() string {
	if o == nil || IsNil(o.OrgUrl) {
		var ret string
		return ret
	}
	return *o.OrgUrl
}

// GetOrgUrlOk returns a tuple with the OrgUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOrgUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OrgUrl) {
		return nil, false
	}
	return o.OrgUrl, true
}

// HasOrgUrl returns a boolean if a field has been set.
func (o *User) HasOrgUrl() bool {
	if o != nil && !IsNil(o.OrgUrl) {
		return true
	}

	return false
}

// SetOrgUrl gets a reference to the given string and assigns it to the OrgUrl field.
func (o *User) SetOrgUrl(v string) {
	o.OrgUrl = &v
}

// GetServiceId returns the ServiceId field value
func (o *User) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *User) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *User) SetServiceId(v string) {
	o.ServiceId = v
}

// GetStatus returns the Status field value
func (o *User) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *User) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *User) SetStatus(v string) {
	o.Status = v
}

// GetSubscriptionCount returns the SubscriptionCount field value if set, zero value otherwise.
func (o *User) GetSubscriptionCount() int64 {
	if o == nil || IsNil(o.SubscriptionCount) {
		var ret int64
		return ret
	}
	return *o.SubscriptionCount
}

// GetSubscriptionCountOk returns a tuple with the SubscriptionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSubscriptionCountOk() (*int64, bool) {
	if o == nil || IsNil(o.SubscriptionCount) {
		return nil, false
	}
	return o.SubscriptionCount, true
}

// HasSubscriptionCount returns a boolean if a field has been set.
func (o *User) HasSubscriptionCount() bool {
	if o != nil && !IsNil(o.SubscriptionCount) {
		return true
	}

	return false
}

// SetSubscriptionCount gets a reference to the given int64 and assigns it to the SubscriptionCount field.
func (o *User) SetSubscriptionCount(v int64) {
	o.SubscriptionCount = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *User) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *User) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *User) SetToken(v string) {
	o.Token = &v
}

// GetUserId returns the UserId field value
func (o *User) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *User) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *User) SetUserId(v string) {
	o.UserId = v
}

// GetUserName returns the UserName field value
func (o *User) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *User) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *User) SetUserName(v string) {
	o.UserName = v
}

// GetUserSubscriptionRole returns the UserSubscriptionRole field value if set, zero value otherwise.
func (o *User) GetUserSubscriptionRole() string {
	if o == nil || IsNil(o.UserSubscriptionRole) {
		var ret string
		return ret
	}
	return *o.UserSubscriptionRole
}

// GetUserSubscriptionRoleOk returns a tuple with the UserSubscriptionRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUserSubscriptionRoleOk() (*string, bool) {
	if o == nil || IsNil(o.UserSubscriptionRole) {
		return nil, false
	}
	return o.UserSubscriptionRole, true
}

// HasUserSubscriptionRole returns a boolean if a field has been set.
func (o *User) HasUserSubscriptionRole() bool {
	if o != nil && !IsNil(o.UserSubscriptionRole) {
		return true
	}

	return false
}

// SetUserSubscriptionRole gets a reference to the given string and assigns it to the UserSubscriptionRole field.
func (o *User) SetUserSubscriptionRole(v string) {
	o.UserSubscriptionRole = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["email"] = o.Email
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["instanceCount"] = o.InstanceCount
	toSerialize["lastModifiedAt"] = o.LastModifiedAt
	if !IsNil(o.LastModifiedByUserID) {
		toSerialize["lastModifiedByUserID"] = o.LastModifiedByUserID
	}
	if !IsNil(o.LastModifiedByUserName) {
		toSerialize["lastModifiedByUserName"] = o.LastModifiedByUserName
	}
	toSerialize["orgId"] = o.OrgId
	toSerialize["orgName"] = o.OrgName
	if !IsNil(o.OrgUrl) {
		toSerialize["orgUrl"] = o.OrgUrl
	}
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["status"] = o.Status
	if !IsNil(o.SubscriptionCount) {
		toSerialize["subscriptionCount"] = o.SubscriptionCount
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["userId"] = o.UserId
	toSerialize["userName"] = o.UserName
	if !IsNil(o.UserSubscriptionRole) {
		toSerialize["userSubscriptionRole"] = o.UserSubscriptionRole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *User) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"email",
		"environmentId",
		"instanceCount",
		"lastModifiedAt",
		"orgId",
		"orgName",
		"serviceId",
		"status",
		"userId",
		"userName",
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

	varUser := _User{}

	err = json.Unmarshal(data, &varUser)

	if err != nil {
		return err
	}

	*o = User(varUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "email")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "environmentId")
		delete(additionalProperties, "instanceCount")
		delete(additionalProperties, "lastModifiedAt")
		delete(additionalProperties, "lastModifiedByUserID")
		delete(additionalProperties, "lastModifiedByUserName")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "orgUrl")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subscriptionCount")
		delete(additionalProperties, "token")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "userName")
		delete(additionalProperties, "userSubscriptionRole")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


