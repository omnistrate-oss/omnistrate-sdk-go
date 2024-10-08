/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package omnistrategosdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DescribeUserResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeUserResult{}

// DescribeUserResult struct for DescribeUserResult
type DescribeUserResult struct {
	Address *Address `json:"address,omitempty"`
	// The user creation time
	CreatedAt string `json:"createdAt"`
	// The email of the user
	Email string `json:"email"`
	// The User ID
	Id string `json:"id"`
	// The user update time
	LastModifiedAt string `json:"lastModifiedAt"`
	// The name of the user
	Name string `json:"name"`
	// The description of the org that this user owns
	OrgDescription string `json:"orgDescription"`
	// The favicon of the org that this user owns
	OrgFavIconURL string `json:"orgFavIconURL"`
	// The ID of the org that this user owns
	OrgId string `json:"orgId"`
	// The logo of the org that this user owns
	OrgLogoURL string `json:"orgLogoURL"`
	// The org name that this user owns
	OrgName string `json:"orgName"`
	// The privacy policy for the org that this user owns
	OrgPrivacyPolicy string `json:"orgPrivacyPolicy"`
	// The support email of the org that this user owns
	OrgSupportEmail string `json:"orgSupportEmail"`
	// The terms of use for the org that this user owns
	OrgTermsOfUse string `json:"orgTermsOfUse"`
	// The url of the org that this user owns
	OrgURL string `json:"orgURL"`
	// The name of the plan of the user
	PlanName *string `json:"planName,omitempty"`
	// The role type of the user
	RoleType string `json:"roleType"`
}

type _DescribeUserResult DescribeUserResult

// NewDescribeUserResult instantiates a new DescribeUserResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeUserResult(createdAt string, email string, id string, lastModifiedAt string, name string, orgDescription string, orgFavIconURL string, orgId string, orgLogoURL string, orgName string, orgPrivacyPolicy string, orgSupportEmail string, orgTermsOfUse string, orgURL string, roleType string) *DescribeUserResult {
	this := DescribeUserResult{}
	this.CreatedAt = createdAt
	this.Email = email
	this.Id = id
	this.LastModifiedAt = lastModifiedAt
	this.Name = name
	this.OrgDescription = orgDescription
	this.OrgFavIconURL = orgFavIconURL
	this.OrgId = orgId
	this.OrgLogoURL = orgLogoURL
	this.OrgName = orgName
	this.OrgPrivacyPolicy = orgPrivacyPolicy
	this.OrgSupportEmail = orgSupportEmail
	this.OrgTermsOfUse = orgTermsOfUse
	this.OrgURL = orgURL
	this.RoleType = roleType
	return &this
}

// NewDescribeUserResultWithDefaults instantiates a new DescribeUserResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeUserResultWithDefaults() *DescribeUserResult {
	this := DescribeUserResult{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DescribeUserResult) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DescribeUserResult) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *DescribeUserResult) SetAddress(v Address) {
	o.Address = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DescribeUserResult) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DescribeUserResult) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetEmail returns the Email field value
func (o *DescribeUserResult) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *DescribeUserResult) SetEmail(v string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *DescribeUserResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeUserResult) SetId(v string) {
	o.Id = v
}

// GetLastModifiedAt returns the LastModifiedAt field value
func (o *DescribeUserResult) GetLastModifiedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastModifiedAt
}

// GetLastModifiedAtOk returns a tuple with the LastModifiedAt field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetLastModifiedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedAt, true
}

// SetLastModifiedAt sets field value
func (o *DescribeUserResult) SetLastModifiedAt(v string) {
	o.LastModifiedAt = v
}

// GetName returns the Name field value
func (o *DescribeUserResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeUserResult) SetName(v string) {
	o.Name = v
}

// GetOrgDescription returns the OrgDescription field value
func (o *DescribeUserResult) GetOrgDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgDescription
}

// GetOrgDescriptionOk returns a tuple with the OrgDescription field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetOrgDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgDescription, true
}

// SetOrgDescription sets field value
func (o *DescribeUserResult) SetOrgDescription(v string) {
	o.OrgDescription = v
}

// GetOrgFavIconURL returns the OrgFavIconURL field value
func (o *DescribeUserResult) GetOrgFavIconURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgFavIconURL
}

// GetOrgFavIconURLOk returns a tuple with the OrgFavIconURL field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetOrgFavIconURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgFavIconURL, true
}

// SetOrgFavIconURL sets field value
func (o *DescribeUserResult) SetOrgFavIconURL(v string) {
	o.OrgFavIconURL = v
}

// GetOrgId returns the OrgId field value
func (o *DescribeUserResult) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *DescribeUserResult) SetOrgId(v string) {
	o.OrgId = v
}

// GetOrgLogoURL returns the OrgLogoURL field value
func (o *DescribeUserResult) GetOrgLogoURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgLogoURL
}

// GetOrgLogoURLOk returns a tuple with the OrgLogoURL field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetOrgLogoURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgLogoURL, true
}

// SetOrgLogoURL sets field value
func (o *DescribeUserResult) SetOrgLogoURL(v string) {
	o.OrgLogoURL = v
}

// GetOrgName returns the OrgName field value
func (o *DescribeUserResult) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *DescribeUserResult) SetOrgName(v string) {
	o.OrgName = v
}

// GetOrgPrivacyPolicy returns the OrgPrivacyPolicy field value
func (o *DescribeUserResult) GetOrgPrivacyPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgPrivacyPolicy
}

// GetOrgPrivacyPolicyOk returns a tuple with the OrgPrivacyPolicy field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetOrgPrivacyPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgPrivacyPolicy, true
}

// SetOrgPrivacyPolicy sets field value
func (o *DescribeUserResult) SetOrgPrivacyPolicy(v string) {
	o.OrgPrivacyPolicy = v
}

// GetOrgSupportEmail returns the OrgSupportEmail field value
func (o *DescribeUserResult) GetOrgSupportEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgSupportEmail
}

// GetOrgSupportEmailOk returns a tuple with the OrgSupportEmail field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetOrgSupportEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgSupportEmail, true
}

// SetOrgSupportEmail sets field value
func (o *DescribeUserResult) SetOrgSupportEmail(v string) {
	o.OrgSupportEmail = v
}

// GetOrgTermsOfUse returns the OrgTermsOfUse field value
func (o *DescribeUserResult) GetOrgTermsOfUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgTermsOfUse
}

// GetOrgTermsOfUseOk returns a tuple with the OrgTermsOfUse field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetOrgTermsOfUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgTermsOfUse, true
}

// SetOrgTermsOfUse sets field value
func (o *DescribeUserResult) SetOrgTermsOfUse(v string) {
	o.OrgTermsOfUse = v
}

// GetOrgURL returns the OrgURL field value
func (o *DescribeUserResult) GetOrgURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgURL
}

// GetOrgURLOk returns a tuple with the OrgURL field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetOrgURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgURL, true
}

// SetOrgURL sets field value
func (o *DescribeUserResult) SetOrgURL(v string) {
	o.OrgURL = v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *DescribeUserResult) GetPlanName() string {
	if o == nil || IsNil(o.PlanName) {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlanName) {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *DescribeUserResult) HasPlanName() bool {
	if o != nil && !IsNil(o.PlanName) {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *DescribeUserResult) SetPlanName(v string) {
	o.PlanName = &v
}

// GetRoleType returns the RoleType field value
func (o *DescribeUserResult) GetRoleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value
// and a boolean to check if the value has been set.
func (o *DescribeUserResult) GetRoleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleType, true
}

// SetRoleType sets field value
func (o *DescribeUserResult) SetRoleType(v string) {
	o.RoleType = v
}

func (o DescribeUserResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeUserResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["email"] = o.Email
	toSerialize["id"] = o.Id
	toSerialize["lastModifiedAt"] = o.LastModifiedAt
	toSerialize["name"] = o.Name
	toSerialize["orgDescription"] = o.OrgDescription
	toSerialize["orgFavIconURL"] = o.OrgFavIconURL
	toSerialize["orgId"] = o.OrgId
	toSerialize["orgLogoURL"] = o.OrgLogoURL
	toSerialize["orgName"] = o.OrgName
	toSerialize["orgPrivacyPolicy"] = o.OrgPrivacyPolicy
	toSerialize["orgSupportEmail"] = o.OrgSupportEmail
	toSerialize["orgTermsOfUse"] = o.OrgTermsOfUse
	toSerialize["orgURL"] = o.OrgURL
	if !IsNil(o.PlanName) {
		toSerialize["planName"] = o.PlanName
	}
	toSerialize["roleType"] = o.RoleType
	return toSerialize, nil
}

func (o *DescribeUserResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"email",
		"id",
		"lastModifiedAt",
		"name",
		"orgDescription",
		"orgFavIconURL",
		"orgId",
		"orgLogoURL",
		"orgName",
		"orgPrivacyPolicy",
		"orgSupportEmail",
		"orgTermsOfUse",
		"orgURL",
		"roleType",
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

	varDescribeUserResult := _DescribeUserResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeUserResult)

	if err != nil {
		return err
	}

	*o = DescribeUserResult(varDescribeUserResult)

	return err
}

type NullableDescribeUserResult struct {
	value *DescribeUserResult
	isSet bool
}

func (v NullableDescribeUserResult) Get() *DescribeUserResult {
	return v.value
}

func (v *NullableDescribeUserResult) Set(val *DescribeUserResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeUserResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeUserResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeUserResult(val *DescribeUserResult) *NullableDescribeUserResult {
	return &NullableDescribeUserResult{value: val, isSet: true}
}

func (v NullableDescribeUserResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeUserResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


