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

// checks if the DescribeAccountConfigByGCPProjectIDResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeAccountConfigByGCPProjectIDResult{}

// DescribeAccountConfigByGCPProjectIDResult struct for DescribeAccountConfigByGCPProjectIDResult
type DescribeAccountConfigByGCPProjectIDResult struct {
	// The BYOA instance IDs that this account config is tied to
	ByoaInstanceIDs []string `json:"byoaInstanceIDs,omitempty"`
	// Cloud Provider ID to operate on
	CloudProviderId string `json:"cloudProviderId"`
	// The description for the account
	Description string `json:"description"`
	// The GCP project ID
	GcpProjectID string `json:"gcpProjectID"`
	// The GCP project number
	GcpProjectNumber string `json:"gcpProjectNumber"`
	// The GCP service account email
	GcpServiceAccountEmail string `json:"gcpServiceAccountEmail"`
	// Account Config ID to operate on
	Id string `json:"id"`
	// The name of the account
	Name string `json:"name"`
	// The status of the account
	Status string `json:"status"`
	// The status message of the account
	StatusMessage string `json:"statusMessage"`
}

type _DescribeAccountConfigByGCPProjectIDResult DescribeAccountConfigByGCPProjectIDResult

// NewDescribeAccountConfigByGCPProjectIDResult instantiates a new DescribeAccountConfigByGCPProjectIDResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeAccountConfigByGCPProjectIDResult(cloudProviderId string, description string, gcpProjectID string, gcpProjectNumber string, gcpServiceAccountEmail string, id string, name string, status string, statusMessage string) *DescribeAccountConfigByGCPProjectIDResult {
	this := DescribeAccountConfigByGCPProjectIDResult{}
	this.CloudProviderId = cloudProviderId
	this.Description = description
	this.GcpProjectID = gcpProjectID
	this.GcpProjectNumber = gcpProjectNumber
	this.GcpServiceAccountEmail = gcpServiceAccountEmail
	this.Id = id
	this.Name = name
	this.Status = status
	this.StatusMessage = statusMessage
	return &this
}

// NewDescribeAccountConfigByGCPProjectIDResultWithDefaults instantiates a new DescribeAccountConfigByGCPProjectIDResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeAccountConfigByGCPProjectIDResultWithDefaults() *DescribeAccountConfigByGCPProjectIDResult {
	this := DescribeAccountConfigByGCPProjectIDResult{}
	return &this
}

// GetByoaInstanceIDs returns the ByoaInstanceIDs field value if set, zero value otherwise.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetByoaInstanceIDs() []string {
	if o == nil || IsNil(o.ByoaInstanceIDs) {
		var ret []string
		return ret
	}
	return o.ByoaInstanceIDs
}

// GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetByoaInstanceIDsOk() ([]string, bool) {
	if o == nil || IsNil(o.ByoaInstanceIDs) {
		return nil, false
	}
	return o.ByoaInstanceIDs, true
}

// HasByoaInstanceIDs returns a boolean if a field has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) HasByoaInstanceIDs() bool {
	if o != nil && !IsNil(o.ByoaInstanceIDs) {
		return true
	}

	return false
}

// SetByoaInstanceIDs gets a reference to the given []string and assigns it to the ByoaInstanceIDs field.
func (o *DescribeAccountConfigByGCPProjectIDResult) SetByoaInstanceIDs(v []string) {
	o.ByoaInstanceIDs = v
}

// GetCloudProviderId returns the CloudProviderId field value
func (o *DescribeAccountConfigByGCPProjectIDResult) GetCloudProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderId
}

// GetCloudProviderIdOk returns a tuple with the CloudProviderId field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetCloudProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderId, true
}

// SetCloudProviderId sets field value
func (o *DescribeAccountConfigByGCPProjectIDResult) SetCloudProviderId(v string) {
	o.CloudProviderId = v
}

// GetDescription returns the Description field value
func (o *DescribeAccountConfigByGCPProjectIDResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeAccountConfigByGCPProjectIDResult) SetDescription(v string) {
	o.Description = v
}

// GetGcpProjectID returns the GcpProjectID field value
func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpProjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GcpProjectID
}

// GetGcpProjectIDOk returns a tuple with the GcpProjectID field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpProjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpProjectID, true
}

// SetGcpProjectID sets field value
func (o *DescribeAccountConfigByGCPProjectIDResult) SetGcpProjectID(v string) {
	o.GcpProjectID = v
}

// GetGcpProjectNumber returns the GcpProjectNumber field value
func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpProjectNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GcpProjectNumber
}

// GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpProjectNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpProjectNumber, true
}

// SetGcpProjectNumber sets field value
func (o *DescribeAccountConfigByGCPProjectIDResult) SetGcpProjectNumber(v string) {
	o.GcpProjectNumber = v
}

// GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field value
func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpServiceAccountEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GcpServiceAccountEmail
}

// GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpServiceAccountEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpServiceAccountEmail, true
}

// SetGcpServiceAccountEmail sets field value
func (o *DescribeAccountConfigByGCPProjectIDResult) SetGcpServiceAccountEmail(v string) {
	o.GcpServiceAccountEmail = v
}

// GetId returns the Id field value
func (o *DescribeAccountConfigByGCPProjectIDResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeAccountConfigByGCPProjectIDResult) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DescribeAccountConfigByGCPProjectIDResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeAccountConfigByGCPProjectIDResult) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *DescribeAccountConfigByGCPProjectIDResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DescribeAccountConfigByGCPProjectIDResult) SetStatus(v string) {
	o.Status = v
}

// GetStatusMessage returns the StatusMessage field value
func (o *DescribeAccountConfigByGCPProjectIDResult) GetStatusMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value
// and a boolean to check if the value has been set.
func (o *DescribeAccountConfigByGCPProjectIDResult) GetStatusMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusMessage, true
}

// SetStatusMessage sets field value
func (o *DescribeAccountConfigByGCPProjectIDResult) SetStatusMessage(v string) {
	o.StatusMessage = v
}

func (o DescribeAccountConfigByGCPProjectIDResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeAccountConfigByGCPProjectIDResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ByoaInstanceIDs) {
		toSerialize["byoaInstanceIDs"] = o.ByoaInstanceIDs
	}
	toSerialize["cloudProviderId"] = o.CloudProviderId
	toSerialize["description"] = o.Description
	toSerialize["gcpProjectID"] = o.GcpProjectID
	toSerialize["gcpProjectNumber"] = o.GcpProjectNumber
	toSerialize["gcpServiceAccountEmail"] = o.GcpServiceAccountEmail
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["statusMessage"] = o.StatusMessage
	return toSerialize, nil
}

func (o *DescribeAccountConfigByGCPProjectIDResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderId",
		"description",
		"gcpProjectID",
		"gcpProjectNumber",
		"gcpServiceAccountEmail",
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

	varDescribeAccountConfigByGCPProjectIDResult := _DescribeAccountConfigByGCPProjectIDResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeAccountConfigByGCPProjectIDResult)

	if err != nil {
		return err
	}

	*o = DescribeAccountConfigByGCPProjectIDResult(varDescribeAccountConfigByGCPProjectIDResult)

	return err
}

type NullableDescribeAccountConfigByGCPProjectIDResult struct {
	value *DescribeAccountConfigByGCPProjectIDResult
	isSet bool
}

func (v NullableDescribeAccountConfigByGCPProjectIDResult) Get() *DescribeAccountConfigByGCPProjectIDResult {
	return v.value
}

func (v *NullableDescribeAccountConfigByGCPProjectIDResult) Set(val *DescribeAccountConfigByGCPProjectIDResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeAccountConfigByGCPProjectIDResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeAccountConfigByGCPProjectIDResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeAccountConfigByGCPProjectIDResult(val *DescribeAccountConfigByGCPProjectIDResult) *NullableDescribeAccountConfigByGCPProjectIDResult {
	return &NullableDescribeAccountConfigByGCPProjectIDResult{value: val, isSet: true}
}

func (v NullableDescribeAccountConfigByGCPProjectIDResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeAccountConfigByGCPProjectIDResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


