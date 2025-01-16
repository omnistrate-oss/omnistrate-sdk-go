/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// checks if the BuildServiceFromComposeSpecRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildServiceFromComposeSpecRequest{}

// BuildServiceFromComposeSpecRequest struct for BuildServiceFromComposeSpecRequest
type BuildServiceFromComposeSpecRequest struct {
	// Configs for the service. Key is the compose spec name of the config and value is base64 encoded config content
	Configs *map[string]string `json:"configs,omitempty"`
	// A brief description of the service
	Description *string `json:"description,omitempty"`
	// The environment to build the service in
	Environment *string `json:"environment,omitempty"`
	// The type of service environment
	EnvironmentType *string `json:"environmentType,omitempty"`
	// Base64 encoded Compose Spec YAML in docker compose format
	FileContent string `json:"fileContent"`
	// Name of the Service
	Name string `json:"name"`
	// Release the service after building
	Release *bool `json:"release,omitempty"`
	// Release the service as preferred
	ReleaseAsPreferred *bool `json:"releaseAsPreferred,omitempty"`
	// Release version name
	ReleaseVersionName *string `json:"releaseVersionName,omitempty"`
	// Secrets for the service. Key is the compose spec name of the secret and value is base64 encoded secret content
	Secrets *map[string]string `json:"secrets,omitempty"`
	// The logo for the service
	ServiceLogoURL *string `json:"serviceLogoURL,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _BuildServiceFromComposeSpecRequest BuildServiceFromComposeSpecRequest

// NewBuildServiceFromComposeSpecRequest instantiates a new BuildServiceFromComposeSpecRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildServiceFromComposeSpecRequest(fileContent string, name string, token string) *BuildServiceFromComposeSpecRequest {
	this := BuildServiceFromComposeSpecRequest{}
	this.FileContent = fileContent
	this.Name = name
	this.Token = token
	return &this
}

// NewBuildServiceFromComposeSpecRequestWithDefaults instantiates a new BuildServiceFromComposeSpecRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildServiceFromComposeSpecRequestWithDefaults() *BuildServiceFromComposeSpecRequest {
	this := BuildServiceFromComposeSpecRequest{}
	return &this
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *BuildServiceFromComposeSpecRequest) GetConfigs() map[string]string {
	if o == nil || IsNil(o.Configs) {
		var ret map[string]string
		return ret
	}
	return *o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetConfigsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Configs) {
		return nil, false
	}
	return o.Configs, true
}

// SetConfigs gets a reference to the given map[string]string and assigns it to the Configs field.
func (o *BuildServiceFromComposeSpecRequest) SetConfigs(v map[string]string) {
	o.Configs = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BuildServiceFromComposeSpecRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BuildServiceFromComposeSpecRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *BuildServiceFromComposeSpecRequest) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *BuildServiceFromComposeSpecRequest) SetEnvironment(v string) {
	o.Environment = &v
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *BuildServiceFromComposeSpecRequest) GetEnvironmentType() string {
	if o == nil || IsNil(o.EnvironmentType) {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentType) {
		return nil, false
	}
	return o.EnvironmentType, true
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *BuildServiceFromComposeSpecRequest) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetFileContent returns the FileContent field value
func (o *BuildServiceFromComposeSpecRequest) GetFileContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileContent
}

// GetFileContentOk returns a tuple with the FileContent field value
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetFileContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileContent, true
}

// SetFileContent sets field value
func (o *BuildServiceFromComposeSpecRequest) SetFileContent(v string) {
	o.FileContent = v
}

// GetName returns the Name field value
func (o *BuildServiceFromComposeSpecRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BuildServiceFromComposeSpecRequest) SetName(v string) {
	o.Name = v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *BuildServiceFromComposeSpecRequest) GetRelease() bool {
	if o == nil || IsNil(o.Release) {
		var ret bool
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetReleaseOk() (*bool, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// SetRelease gets a reference to the given bool and assigns it to the Release field.
func (o *BuildServiceFromComposeSpecRequest) SetRelease(v bool) {
	o.Release = &v
}

// GetReleaseAsPreferred returns the ReleaseAsPreferred field value if set, zero value otherwise.
func (o *BuildServiceFromComposeSpecRequest) GetReleaseAsPreferred() bool {
	if o == nil || IsNil(o.ReleaseAsPreferred) {
		var ret bool
		return ret
	}
	return *o.ReleaseAsPreferred
}

// GetReleaseAsPreferredOk returns a tuple with the ReleaseAsPreferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetReleaseAsPreferredOk() (*bool, bool) {
	if o == nil || IsNil(o.ReleaseAsPreferred) {
		return nil, false
	}
	return o.ReleaseAsPreferred, true
}

// SetReleaseAsPreferred gets a reference to the given bool and assigns it to the ReleaseAsPreferred field.
func (o *BuildServiceFromComposeSpecRequest) SetReleaseAsPreferred(v bool) {
	o.ReleaseAsPreferred = &v
}

// GetReleaseVersionName returns the ReleaseVersionName field value if set, zero value otherwise.
func (o *BuildServiceFromComposeSpecRequest) GetReleaseVersionName() string {
	if o == nil || IsNil(o.ReleaseVersionName) {
		var ret string
		return ret
	}
	return *o.ReleaseVersionName
}

// GetReleaseVersionNameOk returns a tuple with the ReleaseVersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetReleaseVersionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseVersionName) {
		return nil, false
	}
	return o.ReleaseVersionName, true
}

// SetReleaseVersionName gets a reference to the given string and assigns it to the ReleaseVersionName field.
func (o *BuildServiceFromComposeSpecRequest) SetReleaseVersionName(v string) {
	o.ReleaseVersionName = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *BuildServiceFromComposeSpecRequest) GetSecrets() map[string]string {
	if o == nil || IsNil(o.Secrets) {
		var ret map[string]string
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetSecretsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}
	return o.Secrets, true
}

// SetSecrets gets a reference to the given map[string]string and assigns it to the Secrets field.
func (o *BuildServiceFromComposeSpecRequest) SetSecrets(v map[string]string) {
	o.Secrets = &v
}

// GetServiceLogoURL returns the ServiceLogoURL field value if set, zero value otherwise.
func (o *BuildServiceFromComposeSpecRequest) GetServiceLogoURL() string {
	if o == nil || IsNil(o.ServiceLogoURL) {
		var ret string
		return ret
	}
	return *o.ServiceLogoURL
}

// GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetServiceLogoURLOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLogoURL) {
		return nil, false
	}
	return o.ServiceLogoURL, true
}

// SetServiceLogoURL gets a reference to the given string and assigns it to the ServiceLogoURL field.
func (o *BuildServiceFromComposeSpecRequest) SetServiceLogoURL(v string) {
	o.ServiceLogoURL = &v
}

// GetToken returns the Token field value
func (o *BuildServiceFromComposeSpecRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *BuildServiceFromComposeSpecRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *BuildServiceFromComposeSpecRequest) SetToken(v string) {
	o.Token = v
}

func (o BuildServiceFromComposeSpecRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildServiceFromComposeSpecRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configs) {
		toSerialize["configs"] = o.Configs
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.EnvironmentType) {
		toSerialize["environmentType"] = o.EnvironmentType
	}
	toSerialize["fileContent"] = o.FileContent
	toSerialize["name"] = o.Name
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	if !IsNil(o.ReleaseAsPreferred) {
		toSerialize["releaseAsPreferred"] = o.ReleaseAsPreferred
	}
	if !IsNil(o.ReleaseVersionName) {
		toSerialize["releaseVersionName"] = o.ReleaseVersionName
	}
	if !IsNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	if !IsNil(o.ServiceLogoURL) {
		toSerialize["serviceLogoURL"] = o.ServiceLogoURL
	}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BuildServiceFromComposeSpecRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileContent",
		"name",
		"token",
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

	varBuildServiceFromComposeSpecRequest := _BuildServiceFromComposeSpecRequest{}

	err = json.Unmarshal(data, &varBuildServiceFromComposeSpecRequest)

	if err != nil {
		return err
	}

	*o = BuildServiceFromComposeSpecRequest(varBuildServiceFromComposeSpecRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configs")
		delete(additionalProperties, "description")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "environmentType")
		delete(additionalProperties, "fileContent")
		delete(additionalProperties, "name")
		delete(additionalProperties, "release")
		delete(additionalProperties, "releaseAsPreferred")
		delete(additionalProperties, "releaseVersionName")
		delete(additionalProperties, "secrets")
		delete(additionalProperties, "serviceLogoURL")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBuildServiceFromComposeSpecRequest struct {
	value *BuildServiceFromComposeSpecRequest
	isSet bool
}

func (v NullableBuildServiceFromComposeSpecRequest) Get() *BuildServiceFromComposeSpecRequest {
	return v.value
}

func (v *NullableBuildServiceFromComposeSpecRequest) Set(val *BuildServiceFromComposeSpecRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildServiceFromComposeSpecRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildServiceFromComposeSpecRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildServiceFromComposeSpecRequest(val *BuildServiceFromComposeSpecRequest) *NullableBuildServiceFromComposeSpecRequest {
	return &NullableBuildServiceFromComposeSpecRequest{value: val, isSet: true}
}

func (v NullableBuildServiceFromComposeSpecRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildServiceFromComposeSpecRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


