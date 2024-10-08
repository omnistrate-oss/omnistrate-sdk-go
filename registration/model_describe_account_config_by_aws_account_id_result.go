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

// checks if the DescribeAccountConfigByAWSAccountIDResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeAccountConfigByAWSAccountIDResult{}

// DescribeAccountConfigByAWSAccountIDResult struct for DescribeAccountConfigByAWSAccountIDResult
type DescribeAccountConfigByAWSAccountIDResult struct {
	// The AWS account ID
	AwsAccountID string `json:"awsAccountID"`
	// The security role ARN or service account ARN that grants access to operate the infra
	AwsBootstrapRoleARN string `json:"awsBootstrapRoleARN"`
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
	// Account Config ID to operate on
	Id string `json:"id"`
	// The name of the account
	Name string `json:"name"`
	// The status of the account
	Status string `json:"status"`
	// The status message of the account
	StatusMessage string `json:"statusMessage"`
}

type _DescribeAccountConfigByAWSAccountIDResult DescribeAccountConfigByAWSAccountIDResult

// NewDescribeAccountConfigByAWSAccountIDResult instantiates a new DescribeAccountConfigByAWSAccountIDResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeAccountConfigByAWSAccountIDResult(awsAccountID string, awsBootstrapRoleARN string, cloudProviderId string, description string, id string, name string, status string, statusMessage string) *DescribeAccountConfigByAWSAccountIDResult {
	this := DescribeAccountConfigByAWSAccountIDResult{}
	this.AwsAccountID = awsAccountID
	this.AwsBootstrapRoleARN = awsBootstrapRoleARN
	this.CloudProviderId = cloudProviderId
	this.Description = description
	this.Id = id
	this.Name = name
	this.Status = status
	this.StatusMessage = statusMessage
	return &this
}

// NewDescribeAccountConfigByAWSAccountIDResultWithDefaults instantiates a new DescribeAccountConfigByAWSAccountIDResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeAccountConfigByAWSAccountIDResultWithDefaults() *DescribeAccountConfigByAWSAccountIDResult {
	this := DescribeAccountConfigByAWSAccountIDResult{}
	return &this
}

// GetAwsAccountID returns the AwsAccountID field value
func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsAccountID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccountID
}

// GetAwsAccountIDOk returns a tuple with the AwsAccountID field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsAccountIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountID, true
}

// SetAwsAccountID sets field value
func (o *DescribeAccountConfigByAWSAccountIDResult) SetAwsAccountID(v string) {
	o.AwsAccountID = v
}

// GetAwsBootstrapRoleARN returns the AwsBootstrapRoleARN field value
func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsBootstrapRoleARN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsBootstrapRoleARN
}

// GetAwsBootstrapRoleARNOk returns a tuple with the AwsBootstrapRoleARN field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsBootstrapRoleARNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsBootstrapRoleARN, true
}

// SetAwsBootstrapRoleARN sets field value
func (o *DescribeAccountConfigByAWSAccountIDResult) SetAwsBootstrapRoleARN(v string) {
	o.AwsBootstrapRoleARN = v
}

// GetAwsCloudFormationNoLBTemplateURL returns the AwsCloudFormationNoLBTemplateURL field value if set, zero value otherwise.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsCloudFormationNoLBTemplateURL() string {
	if o == nil || IsNil(o.AwsCloudFormationNoLBTemplateURL) {
		var ret string
		return ret
	}
	return *o.AwsCloudFormationNoLBTemplateURL
}

// GetAwsCloudFormationNoLBTemplateURLOk returns a tuple with the AwsCloudFormationNoLBTemplateURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsCloudFormationNoLBTemplateURLOk() (*string, bool) {
	if o == nil || IsNil(o.AwsCloudFormationNoLBTemplateURL) {
		return nil, false
	}
	return o.AwsCloudFormationNoLBTemplateURL, true
}

// HasAwsCloudFormationNoLBTemplateURL returns a boolean if a field has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) HasAwsCloudFormationNoLBTemplateURL() bool {
	if o != nil && !IsNil(o.AwsCloudFormationNoLBTemplateURL) {
		return true
	}

	return false
}

// SetAwsCloudFormationNoLBTemplateURL gets a reference to the given string and assigns it to the AwsCloudFormationNoLBTemplateURL field.
func (o *DescribeAccountConfigByAWSAccountIDResult) SetAwsCloudFormationNoLBTemplateURL(v string) {
	o.AwsCloudFormationNoLBTemplateURL = &v
}

// GetAwsCloudFormationTemplateURL returns the AwsCloudFormationTemplateURL field value if set, zero value otherwise.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsCloudFormationTemplateURL() string {
	if o == nil || IsNil(o.AwsCloudFormationTemplateURL) {
		var ret string
		return ret
	}
	return *o.AwsCloudFormationTemplateURL
}

// GetAwsCloudFormationTemplateURLOk returns a tuple with the AwsCloudFormationTemplateURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsCloudFormationTemplateURLOk() (*string, bool) {
	if o == nil || IsNil(o.AwsCloudFormationTemplateURL) {
		return nil, false
	}
	return o.AwsCloudFormationTemplateURL, true
}

// HasAwsCloudFormationTemplateURL returns a boolean if a field has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) HasAwsCloudFormationTemplateURL() bool {
	if o != nil && !IsNil(o.AwsCloudFormationTemplateURL) {
		return true
	}

	return false
}

// SetAwsCloudFormationTemplateURL gets a reference to the given string and assigns it to the AwsCloudFormationTemplateURL field.
func (o *DescribeAccountConfigByAWSAccountIDResult) SetAwsCloudFormationTemplateURL(v string) {
	o.AwsCloudFormationTemplateURL = &v
}

// GetByoaInstanceIDs returns the ByoaInstanceIDs field value if set, zero value otherwise.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetByoaInstanceIDs() []string {
	if o == nil || IsNil(o.ByoaInstanceIDs) {
		var ret []string
		return ret
	}
	return o.ByoaInstanceIDs
}

// GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetByoaInstanceIDsOk() ([]string, bool) {
	if o == nil || IsNil(o.ByoaInstanceIDs) {
		return nil, false
	}
	return o.ByoaInstanceIDs, true
}

// HasByoaInstanceIDs returns a boolean if a field has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) HasByoaInstanceIDs() bool {
	if o != nil && !IsNil(o.ByoaInstanceIDs) {
		return true
	}

	return false
}

// SetByoaInstanceIDs gets a reference to the given []string and assigns it to the ByoaInstanceIDs field.
func (o *DescribeAccountConfigByAWSAccountIDResult) SetByoaInstanceIDs(v []string) {
	o.ByoaInstanceIDs = v
}

// GetCloudProviderId returns the CloudProviderId field value
func (o *DescribeAccountConfigByAWSAccountIDResult) GetCloudProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderId
}

// GetCloudProviderIdOk returns a tuple with the CloudProviderId field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetCloudProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderId, true
}

// SetCloudProviderId sets field value
func (o *DescribeAccountConfigByAWSAccountIDResult) SetCloudProviderId(v string) {
	o.CloudProviderId = v
}

// GetDescription returns the Description field value
func (o *DescribeAccountConfigByAWSAccountIDResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeAccountConfigByAWSAccountIDResult) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *DescribeAccountConfigByAWSAccountIDResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeAccountConfigByAWSAccountIDResult) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DescribeAccountConfigByAWSAccountIDResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeAccountConfigByAWSAccountIDResult) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *DescribeAccountConfigByAWSAccountIDResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DescribeAccountConfigByAWSAccountIDResult) SetStatus(v string) {
	o.Status = v
}

// GetStatusMessage returns the StatusMessage field value
func (o *DescribeAccountConfigByAWSAccountIDResult) GetStatusMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByAWSAccountIDResult) GetStatusMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusMessage, true
}

// SetStatusMessage sets field value
func (o *DescribeAccountConfigByAWSAccountIDResult) SetStatusMessage(v string) {
	o.StatusMessage = v
}

func (o DescribeAccountConfigByAWSAccountIDResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeAccountConfigByAWSAccountIDResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["awsAccountID"] = o.AwsAccountID
	toSerialize["awsBootstrapRoleARN"] = o.AwsBootstrapRoleARN
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
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["statusMessage"] = o.StatusMessage
	return toSerialize, nil
}

func (o *DescribeAccountConfigByAWSAccountIDResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"awsAccountID",
		"awsBootstrapRoleARN",
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

	varDescribeAccountConfigByAWSAccountIDResult := _DescribeAccountConfigByAWSAccountIDResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeAccountConfigByAWSAccountIDResult)

	if err != nil {
		return err
	}

	*o = DescribeAccountConfigByAWSAccountIDResult(varDescribeAccountConfigByAWSAccountIDResult)

	return err
}

type NullableDescribeAccountConfigByAWSAccountIDResult struct {
	value *DescribeAccountConfigByAWSAccountIDResult
	isSet bool
}

func (v NullableDescribeAccountConfigByAWSAccountIDResult) Get() *DescribeAccountConfigByAWSAccountIDResult {
	return v.value
}

func (v *NullableDescribeAccountConfigByAWSAccountIDResult) Set(val *DescribeAccountConfigByAWSAccountIDResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeAccountConfigByAWSAccountIDResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeAccountConfigByAWSAccountIDResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeAccountConfigByAWSAccountIDResult(val *DescribeAccountConfigByAWSAccountIDResult) *NullableDescribeAccountConfigByAWSAccountIDResult {
	return &NullableDescribeAccountConfigByAWSAccountIDResult{value: val, isSet: true}
}

func (v NullableDescribeAccountConfigByAWSAccountIDResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeAccountConfigByAWSAccountIDResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


