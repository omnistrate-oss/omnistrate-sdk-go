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

// checks if the GitConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitConfiguration{}

// GitConfiguration struct for GitConfiguration
type GitConfiguration struct {
	// The access token
	AccessToken *string `json:"accessToken,omitempty"`
	// The commit SHA to checkout
	CommitSHA *string `json:"commitSHA,omitempty"`
	// The reference name of the repository
	ReferenceName string `json:"referenceName"`
	// The URL of the repository
	RepositoryUrl string `json:"repositoryUrl"`
	// The name of github user
	UserName *string `json:"userName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GitConfiguration GitConfiguration

// NewGitConfiguration instantiates a new GitConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitConfiguration(referenceName string, repositoryUrl string) *GitConfiguration {
	this := GitConfiguration{}
	this.ReferenceName = referenceName
	this.RepositoryUrl = repositoryUrl
	return &this
}

// NewGitConfigurationWithDefaults instantiates a new GitConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitConfigurationWithDefaults() *GitConfiguration {
	this := GitConfiguration{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *GitConfiguration) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitConfiguration) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *GitConfiguration) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *GitConfiguration) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetCommitSHA returns the CommitSHA field value if set, zero value otherwise.
func (o *GitConfiguration) GetCommitSHA() string {
	if o == nil || IsNil(o.CommitSHA) {
		var ret string
		return ret
	}
	return *o.CommitSHA
}

// GetCommitSHAOk returns a tuple with the CommitSHA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitConfiguration) GetCommitSHAOk() (*string, bool) {
	if o == nil || IsNil(o.CommitSHA) {
		return nil, false
	}
	return o.CommitSHA, true
}

// HasCommitSHA returns a boolean if a field has been set.
func (o *GitConfiguration) HasCommitSHA() bool {
	if o != nil && !IsNil(o.CommitSHA) {
		return true
	}

	return false
}

// SetCommitSHA gets a reference to the given string and assigns it to the CommitSHA field.
func (o *GitConfiguration) SetCommitSHA(v string) {
	o.CommitSHA = &v
}

// GetReferenceName returns the ReferenceName field value
func (o *GitConfiguration) GetReferenceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value
// and a boolean to check if the value has been set.
func (o *GitConfiguration) GetReferenceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceName, true
}

// SetReferenceName sets field value
func (o *GitConfiguration) SetReferenceName(v string) {
	o.ReferenceName = v
}

// GetRepositoryUrl returns the RepositoryUrl field value
func (o *GitConfiguration) GetRepositoryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value
// and a boolean to check if the value has been set.
func (o *GitConfiguration) GetRepositoryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryUrl, true
}

// SetRepositoryUrl sets field value
func (o *GitConfiguration) SetRepositoryUrl(v string) {
	o.RepositoryUrl = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *GitConfiguration) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitConfiguration) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *GitConfiguration) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *GitConfiguration) SetUserName(v string) {
	o.UserName = &v
}

func (o GitConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.CommitSHA) {
		toSerialize["commitSHA"] = o.CommitSHA
	}
	toSerialize["referenceName"] = o.ReferenceName
	toSerialize["repositoryUrl"] = o.RepositoryUrl
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GitConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"referenceName",
		"repositoryUrl",
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

	varGitConfiguration := _GitConfiguration{}

	err = json.Unmarshal(data, &varGitConfiguration)

	if err != nil {
		return err
	}

	*o = GitConfiguration(varGitConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessToken")
		delete(additionalProperties, "commitSHA")
		delete(additionalProperties, "referenceName")
		delete(additionalProperties, "repositoryUrl")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGitConfiguration struct {
	value *GitConfiguration
	isSet bool
}

func (v NullableGitConfiguration) Get() *GitConfiguration {
	return v.value
}

func (v *NullableGitConfiguration) Set(val *GitConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableGitConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableGitConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitConfiguration(val *GitConfiguration) *NullableGitConfiguration {
	return &NullableGitConfiguration{value: val, isSet: true}
}

func (v NullableGitConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


