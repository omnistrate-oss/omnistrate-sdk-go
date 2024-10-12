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

// checks if the DescribeAccountConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeAccountConfigResult{}

// DescribeAccountConfigResult struct for DescribeAccountConfigResult
type DescribeAccountConfigResult struct {
	// The AWS account ID
	AwsAccountID *string `json:"awsAccountID,omitempty"`
	// The security role ARN or service account ARN that grants access to operate the infra
	AwsBootstrapRoleARN *string `json:"awsBootstrapRoleARN,omitempty"`
	// The URL to the CloudFormation template (no LoadBalancer policy version)
	AwsCloudFormationNoLBTemplateURL *string `json:"awsCloudFormationNoLBTemplateURL,omitempty"`
	// The URL to the CloudFormation template
	AwsCloudFormationTemplateURL *string `json:"awsCloudFormationTemplateURL,omitempty"`
	// The BYOA instance IDs that this account config is tied to
	ByoaInstanceIDs []string `json:"byoaInstanceIDs,omitempty"`
	// Cloud Provider ID to operate on
	CloudProviderId string `json:"cloudProviderId"`
	// The description for the account
	Description string `json:"description"`
	// The GCP project ID
	GcpProjectID *string `json:"gcpProjectID,omitempty"`
	// The GCP project number
	GcpProjectNumber *string `json:"gcpProjectNumber,omitempty"`
	// The GCP service account email
	GcpServiceAccountEmail *string `json:"gcpServiceAccountEmail,omitempty"`
	// Account Config ID to operate on
	Id string `json:"id"`
	// The name of the account
	Name string `json:"name"`
	// The status of the account
	Status string `json:"status"`
	// The status message of the account
	StatusMessage string `json:"statusMessage"`
	AdditionalProperties map[string]interface{}
}

type _DescribeAccountConfigResult DescribeAccountConfigResult

// NewDescribeAccountConfigResult instantiates a new DescribeAccountConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeAccountConfigResult(cloudProviderId string, description string, id string, name string, status string, statusMessage string) *DescribeAccountConfigResult {
	this := DescribeAccountConfigResult{}
	this.CloudProviderId = cloudProviderId
	this.Description = description
	this.Id = id
	this.Name = name
	this.Status = status
	this.StatusMessage = statusMessage
	return &this
}

// NewDescribeAccountConfigResultWithDefaults instantiates a new DescribeAccountConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeAccountConfigResultWithDefaults() *DescribeAccountConfigResult {
	this := DescribeAccountConfigResult{}
	return &this
}

// GetAwsAccountID returns the AwsAccountID field value if set, zero value otherwise.
func (o *DescribeAccountConfigResult) GetAwsAccountID() string {
	if o == nil || IsNil(o.AwsAccountID) {
		var ret string
		return ret
	}
	return *o.AwsAccountID
}

// GetAwsAccountIDOk returns a tuple with the AwsAccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetAwsAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountID) {
		return nil, false
	}
	return o.AwsAccountID, true
}

// SetAwsAccountID gets a reference to the given string and assigns it to the AwsAccountID field.
func (o *DescribeAccountConfigResult) SetAwsAccountID(v string) {
	o.AwsAccountID = &v
}

// GetAwsBootstrapRoleARN returns the AwsBootstrapRoleARN field value if set, zero value otherwise.
func (o *DescribeAccountConfigResult) GetAwsBootstrapRoleARN() string {
	if o == nil || IsNil(o.AwsBootstrapRoleARN) {
		var ret string
		return ret
	}
	return *o.AwsBootstrapRoleARN
}

// GetAwsBootstrapRoleARNOk returns a tuple with the AwsBootstrapRoleARN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetAwsBootstrapRoleARNOk() (*string, bool) {
	if o == nil || IsNil(o.AwsBootstrapRoleARN) {
		return nil, false
	}
	return o.AwsBootstrapRoleARN, true
}

// SetAwsBootstrapRoleARN gets a reference to the given string and assigns it to the AwsBootstrapRoleARN field.
func (o *DescribeAccountConfigResult) SetAwsBootstrapRoleARN(v string) {
	o.AwsBootstrapRoleARN = &v
}

// GetAwsCloudFormationNoLBTemplateURL returns the AwsCloudFormationNoLBTemplateURL field value if set, zero value otherwise.
func (o *DescribeAccountConfigResult) GetAwsCloudFormationNoLBTemplateURL() string {
	if o == nil || IsNil(o.AwsCloudFormationNoLBTemplateURL) {
		var ret string
		return ret
	}
	return *o.AwsCloudFormationNoLBTemplateURL
}

// GetAwsCloudFormationNoLBTemplateURLOk returns a tuple with the AwsCloudFormationNoLBTemplateURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetAwsCloudFormationNoLBTemplateURLOk() (*string, bool) {
	if o == nil || IsNil(o.AwsCloudFormationNoLBTemplateURL) {
		return nil, false
	}
	return o.AwsCloudFormationNoLBTemplateURL, true
}

// SetAwsCloudFormationNoLBTemplateURL gets a reference to the given string and assigns it to the AwsCloudFormationNoLBTemplateURL field.
func (o *DescribeAccountConfigResult) SetAwsCloudFormationNoLBTemplateURL(v string) {
	o.AwsCloudFormationNoLBTemplateURL = &v
}

// GetAwsCloudFormationTemplateURL returns the AwsCloudFormationTemplateURL field value if set, zero value otherwise.
func (o *DescribeAccountConfigResult) GetAwsCloudFormationTemplateURL() string {
	if o == nil || IsNil(o.AwsCloudFormationTemplateURL) {
		var ret string
		return ret
	}
	return *o.AwsCloudFormationTemplateURL
}

// GetAwsCloudFormationTemplateURLOk returns a tuple with the AwsCloudFormationTemplateURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetAwsCloudFormationTemplateURLOk() (*string, bool) {
	if o == nil || IsNil(o.AwsCloudFormationTemplateURL) {
		return nil, false
	}
	return o.AwsCloudFormationTemplateURL, true
}

// SetAwsCloudFormationTemplateURL gets a reference to the given string and assigns it to the AwsCloudFormationTemplateURL field.
func (o *DescribeAccountConfigResult) SetAwsCloudFormationTemplateURL(v string) {
	o.AwsCloudFormationTemplateURL = &v
}

// GetByoaInstanceIDs returns the ByoaInstanceIDs field value if set, zero value otherwise.
func (o *DescribeAccountConfigResult) GetByoaInstanceIDs() []string {
	if o == nil || IsNil(o.ByoaInstanceIDs) {
		var ret []string
		return ret
	}
	return o.ByoaInstanceIDs
}

// GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetByoaInstanceIDsOk() ([]string, bool) {
	if o == nil || IsNil(o.ByoaInstanceIDs) {
		return nil, false
	}
	return o.ByoaInstanceIDs, true
}

// SetByoaInstanceIDs gets a reference to the given []string and assigns it to the ByoaInstanceIDs field.
func (o *DescribeAccountConfigResult) SetByoaInstanceIDs(v []string) {
	o.ByoaInstanceIDs = v
}

// GetCloudProviderId returns the CloudProviderId field value
func (o *DescribeAccountConfigResult) GetCloudProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderId
}

// GetCloudProviderIdOk returns a tuple with the CloudProviderId field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetCloudProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderId, true
}

// SetCloudProviderId sets field value
func (o *DescribeAccountConfigResult) SetCloudProviderId(v string) {
	o.CloudProviderId = v
}

// GetDescription returns the Description field value
func (o *DescribeAccountConfigResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeAccountConfigResult) SetDescription(v string) {
	o.Description = v
}

// GetGcpProjectID returns the GcpProjectID field value if set, zero value otherwise.
func (o *DescribeAccountConfigResult) GetGcpProjectID() string {
	if o == nil || IsNil(o.GcpProjectID) {
		var ret string
		return ret
	}
	return *o.GcpProjectID
}

// GetGcpProjectIDOk returns a tuple with the GcpProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetGcpProjectIDOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectID) {
		return nil, false
	}
	return o.GcpProjectID, true
}

// SetGcpProjectID gets a reference to the given string and assigns it to the GcpProjectID field.
func (o *DescribeAccountConfigResult) SetGcpProjectID(v string) {
	o.GcpProjectID = &v
}

// GetGcpProjectNumber returns the GcpProjectNumber field value if set, zero value otherwise.
func (o *DescribeAccountConfigResult) GetGcpProjectNumber() string {
	if o == nil || IsNil(o.GcpProjectNumber) {
		var ret string
		return ret
	}
	return *o.GcpProjectNumber
}

// GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetGcpProjectNumberOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectNumber) {
		return nil, false
	}
	return o.GcpProjectNumber, true
}

// SetGcpProjectNumber gets a reference to the given string and assigns it to the GcpProjectNumber field.
func (o *DescribeAccountConfigResult) SetGcpProjectNumber(v string) {
	o.GcpProjectNumber = &v
}

// GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field value if set, zero value otherwise.
func (o *DescribeAccountConfigResult) GetGcpServiceAccountEmail() string {
	if o == nil || IsNil(o.GcpServiceAccountEmail) {
		var ret string
		return ret
	}
	return *o.GcpServiceAccountEmail
}

// GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetGcpServiceAccountEmailOk() (*string, bool) {
	if o == nil || IsNil(o.GcpServiceAccountEmail) {
		return nil, false
	}
	return o.GcpServiceAccountEmail, true
}

// SetGcpServiceAccountEmail gets a reference to the given string and assigns it to the GcpServiceAccountEmail field.
func (o *DescribeAccountConfigResult) SetGcpServiceAccountEmail(v string) {
	o.GcpServiceAccountEmail = &v
}

// GetId returns the Id field value
func (o *DescribeAccountConfigResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeAccountConfigResult) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DescribeAccountConfigResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeAccountConfigResult) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *DescribeAccountConfigResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DescribeAccountConfigResult) SetStatus(v string) {
	o.Status = v
}

// GetStatusMessage returns the StatusMessage field value
func (o *DescribeAccountConfigResult) GetStatusMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigResult) GetStatusMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusMessage, true
}

// SetStatusMessage sets field value
func (o *DescribeAccountConfigResult) SetStatusMessage(v string) {
	o.StatusMessage = v
}

func (o DescribeAccountConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeAccountConfigResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsAccountID) {
		toSerialize["awsAccountID"] = o.AwsAccountID
	}
	if !IsNil(o.AwsBootstrapRoleARN) {
		toSerialize["awsBootstrapRoleARN"] = o.AwsBootstrapRoleARN
	}
	if !IsNil(o.AwsCloudFormationNoLBTemplateURL) {
		toSerialize["awsCloudFormationNoLBTemplateURL"] = o.AwsCloudFormationNoLBTemplateURL
	}
	if !IsNil(o.AwsCloudFormationTemplateURL) {
		toSerialize["awsCloudFormationTemplateURL"] = o.AwsCloudFormationTemplateURL
	}
	if !IsNil(o.ByoaInstanceIDs) {
		toSerialize["byoaInstanceIDs"] = o.ByoaInstanceIDs
	}
	toSerialize["cloudProviderId"] = o.CloudProviderId
	toSerialize["description"] = o.Description
	if !IsNil(o.GcpProjectID) {
		toSerialize["gcpProjectID"] = o.GcpProjectID
	}
	if !IsNil(o.GcpProjectNumber) {
		toSerialize["gcpProjectNumber"] = o.GcpProjectNumber
	}
	if !IsNil(o.GcpServiceAccountEmail) {
		toSerialize["gcpServiceAccountEmail"] = o.GcpServiceAccountEmail
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["statusMessage"] = o.StatusMessage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeAccountConfigResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderId",
		"description",
		"id",
		"name",
		"status",
		"statusMessage",
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

	varDescribeAccountConfigResult := _DescribeAccountConfigResult{}

	err = json.Unmarshal(data, &varDescribeAccountConfigResult)

	if err != nil {
		return err
	}

	*o = DescribeAccountConfigResult(varDescribeAccountConfigResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "awsAccountID")
		delete(additionalProperties, "awsBootstrapRoleARN")
		delete(additionalProperties, "awsCloudFormationNoLBTemplateURL")
		delete(additionalProperties, "awsCloudFormationTemplateURL")
		delete(additionalProperties, "byoaInstanceIDs")
		delete(additionalProperties, "cloudProviderId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "gcpProjectID")
		delete(additionalProperties, "gcpProjectNumber")
		delete(additionalProperties, "gcpServiceAccountEmail")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusMessage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeAccountConfigResult struct {
	value *DescribeAccountConfigResult
	isSet bool
}

func (v NullableDescribeAccountConfigResult) Get() *DescribeAccountConfigResult {
	return v.value
}

func (v *NullableDescribeAccountConfigResult) Set(val *DescribeAccountConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeAccountConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeAccountConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeAccountConfigResult(val *DescribeAccountConfigResult) *NullableDescribeAccountConfigResult {
	return &NullableDescribeAccountConfigResult{value: val, isSet: true}
}

func (v NullableDescribeAccountConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeAccountConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


