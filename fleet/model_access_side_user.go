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

// checks if the AccessSideUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessSideUser{}

// AccessSideUser struct for AccessSideUser
type AccessSideUser struct {
	// The time the user was created.
	CreatedAt string `json:"createdAt"`
	// The user email.
	Email string `json:"email"`
	// Is the user enabled.
	Enabled bool `json:"enabled"`
	// The number of active instances the user has.
	InstanceCount int64 `json:"instanceCount"`
	// The last modified time of the user.
	LastModifiedAt string `json:"lastModifiedAt"`
	// ID of a User
	LastModifiedByUserID *string `json:"lastModifiedByUserID,omitempty"`
	// The user name of the last modifier.
	LastModifiedByUserName *string `json:"lastModifiedByUserName,omitempty"`
	// ID of an Org
	OrgId string `json:"orgId"`
	// The organization name.
	OrgName string `json:"orgName"`
	// The organization URL.
	OrgUrl *string `json:"orgUrl,omitempty"`
	// The status of the user.
	Status string `json:"status"`
	// The number of subscriptions the user has.
	SubscriptionCount int64 `json:"subscriptionCount"`
	// Token to validate the user, if the user is not enabled.
	Token *string `json:"token,omitempty"`
	// ID of a User
	UserId string `json:"userId"`
	// The user name.
	UserName string `json:"userName"`
	AdditionalProperties map[string]interface{}
}

type _AccessSideUser AccessSideUser

// NewAccessSideUser instantiates a new AccessSideUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessSideUser(createdAt string, email string, enabled bool, instanceCount int64, lastModifiedAt string, orgId string, orgName string, status string, subscriptionCount int64, userId string, userName string) *AccessSideUser {
	this := AccessSideUser{}
	this.CreatedAt = createdAt
	this.Email = email
	this.Enabled = enabled
	this.InstanceCount = instanceCount
	this.LastModifiedAt = lastModifiedAt
	this.OrgId = orgId
	this.OrgName = orgName
	this.Status = status
	this.SubscriptionCount = subscriptionCount
	this.UserId = userId
	this.UserName = userName
	return &this
}

// NewAccessSideUserWithDefaults instantiates a new AccessSideUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessSideUserWithDefaults() *AccessSideUser {
	this := AccessSideUser{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *AccessSideUser) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AccessSideUser) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetEmail returns the Email field value
func (o *AccessSideUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AccessSideUser) SetEmail(v string) {
	o.Email = v
}

// GetEnabled returns the Enabled field value
func (o *AccessSideUser) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AccessSideUser) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInstanceCount returns the InstanceCount field value
func (o *AccessSideUser) GetInstanceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetInstanceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value
func (o *AccessSideUser) SetInstanceCount(v int64) {
	o.InstanceCount = v
}

// GetLastModifiedAt returns the LastModifiedAt field value
func (o *AccessSideUser) GetLastModifiedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastModifiedAt
}

// GetLastModifiedAtOk returns a tuple with the LastModifiedAt field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetLastModifiedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedAt, true
}

// SetLastModifiedAt sets field value
func (o *AccessSideUser) SetLastModifiedAt(v string) {
	o.LastModifiedAt = v
}

// GetLastModifiedByUserID returns the LastModifiedByUserID field value if set, zero value otherwise.
func (o *AccessSideUser) GetLastModifiedByUserID() string {
	if o == nil || IsNil(o.LastModifiedByUserID) {
		var ret string
		return ret
	}
	return *o.LastModifiedByUserID
}

// GetLastModifiedByUserIDOk returns a tuple with the LastModifiedByUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetLastModifiedByUserIDOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedByUserID) {
		return nil, false
	}
	return o.LastModifiedByUserID, true
}

// HasLastModifiedByUserID returns a boolean if a field has been set.
func (o *AccessSideUser) HasLastModifiedByUserID() bool {
	if o != nil && !IsNil(o.LastModifiedByUserID) {
		return true
	}

	return false
}

// SetLastModifiedByUserID gets a reference to the given string and assigns it to the LastModifiedByUserID field.
func (o *AccessSideUser) SetLastModifiedByUserID(v string) {
	o.LastModifiedByUserID = &v
}

// GetLastModifiedByUserName returns the LastModifiedByUserName field value if set, zero value otherwise.
func (o *AccessSideUser) GetLastModifiedByUserName() string {
	if o == nil || IsNil(o.LastModifiedByUserName) {
		var ret string
		return ret
	}
	return *o.LastModifiedByUserName
}

// GetLastModifiedByUserNameOk returns a tuple with the LastModifiedByUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetLastModifiedByUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedByUserName) {
		return nil, false
	}
	return o.LastModifiedByUserName, true
}

// HasLastModifiedByUserName returns a boolean if a field has been set.
func (o *AccessSideUser) HasLastModifiedByUserName() bool {
	if o != nil && !IsNil(o.LastModifiedByUserName) {
		return true
	}

	return false
}

// SetLastModifiedByUserName gets a reference to the given string and assigns it to the LastModifiedByUserName field.
func (o *AccessSideUser) SetLastModifiedByUserName(v string) {
	o.LastModifiedByUserName = &v
}

// GetOrgId returns the OrgId field value
func (o *AccessSideUser) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *AccessSideUser) SetOrgId(v string) {
	o.OrgId = v
}

// GetOrgName returns the OrgName field value
func (o *AccessSideUser) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *AccessSideUser) SetOrgName(v string) {
	o.OrgName = v
}

// GetOrgUrl returns the OrgUrl field value if set, zero value otherwise.
func (o *AccessSideUser) GetOrgUrl() string {
	if o == nil || IsNil(o.OrgUrl) {
		var ret string
		return ret
	}
	return *o.OrgUrl
}

// GetOrgUrlOk returns a tuple with the OrgUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetOrgUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OrgUrl) {
		return nil, false
	}
	return o.OrgUrl, true
}

// HasOrgUrl returns a boolean if a field has been set.
func (o *AccessSideUser) HasOrgUrl() bool {
	if o != nil && !IsNil(o.OrgUrl) {
		return true
	}

	return false
}

// SetOrgUrl gets a reference to the given string and assigns it to the OrgUrl field.
func (o *AccessSideUser) SetOrgUrl(v string) {
	o.OrgUrl = &v
}

// GetStatus returns the Status field value
func (o *AccessSideUser) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccessSideUser) SetStatus(v string) {
	o.Status = v
}

// GetSubscriptionCount returns the SubscriptionCount field value
func (o *AccessSideUser) GetSubscriptionCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SubscriptionCount
}

// GetSubscriptionCountOk returns a tuple with the SubscriptionCount field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetSubscriptionCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionCount, true
}

// SetSubscriptionCount sets field value
func (o *AccessSideUser) SetSubscriptionCount(v int64) {
	o.SubscriptionCount = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AccessSideUser) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AccessSideUser) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AccessSideUser) SetToken(v string) {
	o.Token = &v
}

// GetUserId returns the UserId field value
func (o *AccessSideUser) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AccessSideUser) SetUserId(v string) {
	o.UserId = v
}

// GetUserName returns the UserName field value
func (o *AccessSideUser) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *AccessSideUser) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *AccessSideUser) SetUserName(v string) {
	o.UserName = v
}

func (o AccessSideUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessSideUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["email"] = o.Email
	toSerialize["enabled"] = o.Enabled
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
	toSerialize["status"] = o.Status
	toSerialize["subscriptionCount"] = o.SubscriptionCount
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["userId"] = o.UserId
	toSerialize["userName"] = o.UserName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessSideUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"email",
		"enabled",
		"instanceCount",
		"lastModifiedAt",
		"orgId",
		"orgName",
		"status",
		"subscriptionCount",
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

	varAccessSideUser := _AccessSideUser{}

	err = json.Unmarshal(data, &varAccessSideUser)

	if err != nil {
		return err
	}

	*o = AccessSideUser(varAccessSideUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "email")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "instanceCount")
		delete(additionalProperties, "lastModifiedAt")
		delete(additionalProperties, "lastModifiedByUserID")
		delete(additionalProperties, "lastModifiedByUserName")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "orgUrl")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subscriptionCount")
		delete(additionalProperties, "token")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessSideUser struct {
	value *AccessSideUser
	isSet bool
}

func (v NullableAccessSideUser) Get() *AccessSideUser {
	return v.value
}

func (v *NullableAccessSideUser) Set(val *AccessSideUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessSideUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessSideUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessSideUser(val *AccessSideUser) *NullableAccessSideUser {
	return &NullableAccessSideUser{value: val, isSet: true}
}

func (v NullableAccessSideUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessSideUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


