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

// checks if the FleetDescribeAccountConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FleetDescribeAccountConfigResult{}

// FleetDescribeAccountConfigResult Account configuration including the role required to access
type FleetDescribeAccountConfigResult struct {
	// The AWS account ID
	AwsAccountID *string `json:"awsAccountID,omitempty"`
	// The security role ARN or service account ARN that grants access to operate the infra
	AwsBootstrapRoleARN *string `json:"awsBootstrapRoleARN,omitempty"`
	// The URL to the CloudFormation template (no LoadBalancer policy version)
	AwsCloudFormationNoLBTemplateURL *string `json:"awsCloudFormationNoLBTemplateURL,omitempty"`
	// The URL to the CloudFormation template
	AwsCloudFormationTemplateURL *string `json:"awsCloudFormationTemplateURL,omitempty"`
	// The Azure bootstrap shell command
	AzureBootstrapShellCommand *string `json:"azureBootstrapShellCommand,omitempty"`
	// The Azure disconnect shell command
	AzureDisconnectShellCommand *string `json:"azureDisconnectShellCommand,omitempty"`
	// The Azure subscription ID
	AzureSubscriptionID *string `json:"azureSubscriptionID,omitempty"`
	// The Azure tenant ID
	AzureTenantID *string `json:"azureTenantID,omitempty"`
	// The BYOA instance IDs that this account config is tied to
	ByoaInstanceIDs []string `json:"byoaInstanceIDs,omitempty"`
	// ID of an CloudProvider
	CloudProviderId string `json:"cloudProviderId"`
	// The description for the account
	Description string `json:"description"`
	// The GCP bootstrap shell command
	GcpBootstrapShellCommand *string `json:"gcpBootstrapShellCommand,omitempty"`
	// The GCP disconnect shell command
	GcpDisconnectShellCommand *string `json:"gcpDisconnectShellCommand,omitempty"`
	// The GCP project ID
	GcpProjectID *string `json:"gcpProjectID,omitempty"`
	// The GCP project number
	GcpProjectNumber *string `json:"gcpProjectNumber,omitempty"`
	// The GCP service account email
	GcpServiceAccountEmail *string `json:"gcpServiceAccountEmail,omitempty"`
	// ID of an Account Config
	Id string `json:"id"`
	// The name of the account
	Name string `json:"name"`
	// The status of the account configuration
	Status string `json:"status"`
	// The status message of the account
	StatusMessage string `json:"statusMessage"`
	AdditionalProperties map[string]interface{}
}

type _FleetDescribeAccountConfigResult FleetDescribeAccountConfigResult

// NewFleetDescribeAccountConfigResult instantiates a new FleetDescribeAccountConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetDescribeAccountConfigResult(cloudProviderId string, description string, id string, name string, status string, statusMessage string) *FleetDescribeAccountConfigResult {
	this := FleetDescribeAccountConfigResult{}
	this.CloudProviderId = cloudProviderId
	this.Description = description
	this.Id = id
	this.Name = name
	this.Status = status
	this.StatusMessage = statusMessage
	return &this
}

// NewFleetDescribeAccountConfigResultWithDefaults instantiates a new FleetDescribeAccountConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetDescribeAccountConfigResultWithDefaults() *FleetDescribeAccountConfigResult {
	this := FleetDescribeAccountConfigResult{}
	return &this
}

// GetAwsAccountID returns the AwsAccountID field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetAwsAccountID() string {
	if o == nil || IsNil(o.AwsAccountID) {
		var ret string
		return ret
	}
	return *o.AwsAccountID
}

// GetAwsAccountIDOk returns a tuple with the AwsAccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetAwsAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountID) {
		return nil, false
	}
	return o.AwsAccountID, true
}

// HasAwsAccountID returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasAwsAccountID() bool {
	if o != nil && !IsNil(o.AwsAccountID) {
		return true
	}

	return false
}

// SetAwsAccountID gets a reference to the given string and assigns it to the AwsAccountID field.
func (o *FleetDescribeAccountConfigResult) SetAwsAccountID(v string) {
	o.AwsAccountID = &v
}

// GetAwsBootstrapRoleARN returns the AwsBootstrapRoleARN field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetAwsBootstrapRoleARN() string {
	if o == nil || IsNil(o.AwsBootstrapRoleARN) {
		var ret string
		return ret
	}
	return *o.AwsBootstrapRoleARN
}

// GetAwsBootstrapRoleARNOk returns a tuple with the AwsBootstrapRoleARN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetAwsBootstrapRoleARNOk() (*string, bool) {
	if o == nil || IsNil(o.AwsBootstrapRoleARN) {
		return nil, false
	}
	return o.AwsBootstrapRoleARN, true
}

// HasAwsBootstrapRoleARN returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasAwsBootstrapRoleARN() bool {
	if o != nil && !IsNil(o.AwsBootstrapRoleARN) {
		return true
	}

	return false
}

// SetAwsBootstrapRoleARN gets a reference to the given string and assigns it to the AwsBootstrapRoleARN field.
func (o *FleetDescribeAccountConfigResult) SetAwsBootstrapRoleARN(v string) {
	o.AwsBootstrapRoleARN = &v
}

// GetAwsCloudFormationNoLBTemplateURL returns the AwsCloudFormationNoLBTemplateURL field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetAwsCloudFormationNoLBTemplateURL() string {
	if o == nil || IsNil(o.AwsCloudFormationNoLBTemplateURL) {
		var ret string
		return ret
	}
	return *o.AwsCloudFormationNoLBTemplateURL
}

// GetAwsCloudFormationNoLBTemplateURLOk returns a tuple with the AwsCloudFormationNoLBTemplateURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetAwsCloudFormationNoLBTemplateURLOk() (*string, bool) {
	if o == nil || IsNil(o.AwsCloudFormationNoLBTemplateURL) {
		return nil, false
	}
	return o.AwsCloudFormationNoLBTemplateURL, true
}

// HasAwsCloudFormationNoLBTemplateURL returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasAwsCloudFormationNoLBTemplateURL() bool {
	if o != nil && !IsNil(o.AwsCloudFormationNoLBTemplateURL) {
		return true
	}

	return false
}

// SetAwsCloudFormationNoLBTemplateURL gets a reference to the given string and assigns it to the AwsCloudFormationNoLBTemplateURL field.
func (o *FleetDescribeAccountConfigResult) SetAwsCloudFormationNoLBTemplateURL(v string) {
	o.AwsCloudFormationNoLBTemplateURL = &v
}

// GetAwsCloudFormationTemplateURL returns the AwsCloudFormationTemplateURL field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetAwsCloudFormationTemplateURL() string {
	if o == nil || IsNil(o.AwsCloudFormationTemplateURL) {
		var ret string
		return ret
	}
	return *o.AwsCloudFormationTemplateURL
}

// GetAwsCloudFormationTemplateURLOk returns a tuple with the AwsCloudFormationTemplateURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetAwsCloudFormationTemplateURLOk() (*string, bool) {
	if o == nil || IsNil(o.AwsCloudFormationTemplateURL) {
		return nil, false
	}
	return o.AwsCloudFormationTemplateURL, true
}

// HasAwsCloudFormationTemplateURL returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasAwsCloudFormationTemplateURL() bool {
	if o != nil && !IsNil(o.AwsCloudFormationTemplateURL) {
		return true
	}

	return false
}

// SetAwsCloudFormationTemplateURL gets a reference to the given string and assigns it to the AwsCloudFormationTemplateURL field.
func (o *FleetDescribeAccountConfigResult) SetAwsCloudFormationTemplateURL(v string) {
	o.AwsCloudFormationTemplateURL = &v
}

// GetAzureBootstrapShellCommand returns the AzureBootstrapShellCommand field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetAzureBootstrapShellCommand() string {
	if o == nil || IsNil(o.AzureBootstrapShellCommand) {
		var ret string
		return ret
	}
	return *o.AzureBootstrapShellCommand
}

// GetAzureBootstrapShellCommandOk returns a tuple with the AzureBootstrapShellCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetAzureBootstrapShellCommandOk() (*string, bool) {
	if o == nil || IsNil(o.AzureBootstrapShellCommand) {
		return nil, false
	}
	return o.AzureBootstrapShellCommand, true
}

// HasAzureBootstrapShellCommand returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasAzureBootstrapShellCommand() bool {
	if o != nil && !IsNil(o.AzureBootstrapShellCommand) {
		return true
	}

	return false
}

// SetAzureBootstrapShellCommand gets a reference to the given string and assigns it to the AzureBootstrapShellCommand field.
func (o *FleetDescribeAccountConfigResult) SetAzureBootstrapShellCommand(v string) {
	o.AzureBootstrapShellCommand = &v
}

// GetAzureDisconnectShellCommand returns the AzureDisconnectShellCommand field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetAzureDisconnectShellCommand() string {
	if o == nil || IsNil(o.AzureDisconnectShellCommand) {
		var ret string
		return ret
	}
	return *o.AzureDisconnectShellCommand
}

// GetAzureDisconnectShellCommandOk returns a tuple with the AzureDisconnectShellCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetAzureDisconnectShellCommandOk() (*string, bool) {
	if o == nil || IsNil(o.AzureDisconnectShellCommand) {
		return nil, false
	}
	return o.AzureDisconnectShellCommand, true
}

// HasAzureDisconnectShellCommand returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasAzureDisconnectShellCommand() bool {
	if o != nil && !IsNil(o.AzureDisconnectShellCommand) {
		return true
	}

	return false
}

// SetAzureDisconnectShellCommand gets a reference to the given string and assigns it to the AzureDisconnectShellCommand field.
func (o *FleetDescribeAccountConfigResult) SetAzureDisconnectShellCommand(v string) {
	o.AzureDisconnectShellCommand = &v
}

// GetAzureSubscriptionID returns the AzureSubscriptionID field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetAzureSubscriptionID() string {
	if o == nil || IsNil(o.AzureSubscriptionID) {
		var ret string
		return ret
	}
	return *o.AzureSubscriptionID
}

// GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetAzureSubscriptionIDOk() (*string, bool) {
	if o == nil || IsNil(o.AzureSubscriptionID) {
		return nil, false
	}
	return o.AzureSubscriptionID, true
}

// HasAzureSubscriptionID returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasAzureSubscriptionID() bool {
	if o != nil && !IsNil(o.AzureSubscriptionID) {
		return true
	}

	return false
}

// SetAzureSubscriptionID gets a reference to the given string and assigns it to the AzureSubscriptionID field.
func (o *FleetDescribeAccountConfigResult) SetAzureSubscriptionID(v string) {
	o.AzureSubscriptionID = &v
}

// GetAzureTenantID returns the AzureTenantID field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetAzureTenantID() string {
	if o == nil || IsNil(o.AzureTenantID) {
		var ret string
		return ret
	}
	return *o.AzureTenantID
}

// GetAzureTenantIDOk returns a tuple with the AzureTenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetAzureTenantIDOk() (*string, bool) {
	if o == nil || IsNil(o.AzureTenantID) {
		return nil, false
	}
	return o.AzureTenantID, true
}

// HasAzureTenantID returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasAzureTenantID() bool {
	if o != nil && !IsNil(o.AzureTenantID) {
		return true
	}

	return false
}

// SetAzureTenantID gets a reference to the given string and assigns it to the AzureTenantID field.
func (o *FleetDescribeAccountConfigResult) SetAzureTenantID(v string) {
	o.AzureTenantID = &v
}

// GetByoaInstanceIDs returns the ByoaInstanceIDs field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetByoaInstanceIDs() []string {
	if o == nil || IsNil(o.ByoaInstanceIDs) {
		var ret []string
		return ret
	}
	return o.ByoaInstanceIDs
}

// GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetByoaInstanceIDsOk() ([]string, bool) {
	if o == nil || IsNil(o.ByoaInstanceIDs) {
		return nil, false
	}
	return o.ByoaInstanceIDs, true
}

// HasByoaInstanceIDs returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasByoaInstanceIDs() bool {
	if o != nil && !IsNil(o.ByoaInstanceIDs) {
		return true
	}

	return false
}

// SetByoaInstanceIDs gets a reference to the given []string and assigns it to the ByoaInstanceIDs field.
func (o *FleetDescribeAccountConfigResult) SetByoaInstanceIDs(v []string) {
	o.ByoaInstanceIDs = v
}

// GetCloudProviderId returns the CloudProviderId field value
func (o *FleetDescribeAccountConfigResult) GetCloudProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderId
}

// GetCloudProviderIdOk returns a tuple with the CloudProviderId field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetCloudProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderId, true
}

// SetCloudProviderId sets field value
func (o *FleetDescribeAccountConfigResult) SetCloudProviderId(v string) {
	o.CloudProviderId = v
}

// GetDescription returns the Description field value
func (o *FleetDescribeAccountConfigResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FleetDescribeAccountConfigResult) SetDescription(v string) {
	o.Description = v
}

// GetGcpBootstrapShellCommand returns the GcpBootstrapShellCommand field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetGcpBootstrapShellCommand() string {
	if o == nil || IsNil(o.GcpBootstrapShellCommand) {
		var ret string
		return ret
	}
	return *o.GcpBootstrapShellCommand
}

// GetGcpBootstrapShellCommandOk returns a tuple with the GcpBootstrapShellCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetGcpBootstrapShellCommandOk() (*string, bool) {
	if o == nil || IsNil(o.GcpBootstrapShellCommand) {
		return nil, false
	}
	return o.GcpBootstrapShellCommand, true
}

// HasGcpBootstrapShellCommand returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasGcpBootstrapShellCommand() bool {
	if o != nil && !IsNil(o.GcpBootstrapShellCommand) {
		return true
	}

	return false
}

// SetGcpBootstrapShellCommand gets a reference to the given string and assigns it to the GcpBootstrapShellCommand field.
func (o *FleetDescribeAccountConfigResult) SetGcpBootstrapShellCommand(v string) {
	o.GcpBootstrapShellCommand = &v
}

// GetGcpDisconnectShellCommand returns the GcpDisconnectShellCommand field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetGcpDisconnectShellCommand() string {
	if o == nil || IsNil(o.GcpDisconnectShellCommand) {
		var ret string
		return ret
	}
	return *o.GcpDisconnectShellCommand
}

// GetGcpDisconnectShellCommandOk returns a tuple with the GcpDisconnectShellCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetGcpDisconnectShellCommandOk() (*string, bool) {
	if o == nil || IsNil(o.GcpDisconnectShellCommand) {
		return nil, false
	}
	return o.GcpDisconnectShellCommand, true
}

// HasGcpDisconnectShellCommand returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasGcpDisconnectShellCommand() bool {
	if o != nil && !IsNil(o.GcpDisconnectShellCommand) {
		return true
	}

	return false
}

// SetGcpDisconnectShellCommand gets a reference to the given string and assigns it to the GcpDisconnectShellCommand field.
func (o *FleetDescribeAccountConfigResult) SetGcpDisconnectShellCommand(v string) {
	o.GcpDisconnectShellCommand = &v
}

// GetGcpProjectID returns the GcpProjectID field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetGcpProjectID() string {
	if o == nil || IsNil(o.GcpProjectID) {
		var ret string
		return ret
	}
	return *o.GcpProjectID
}

// GetGcpProjectIDOk returns a tuple with the GcpProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetGcpProjectIDOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectID) {
		return nil, false
	}
	return o.GcpProjectID, true
}

// HasGcpProjectID returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasGcpProjectID() bool {
	if o != nil && !IsNil(o.GcpProjectID) {
		return true
	}

	return false
}

// SetGcpProjectID gets a reference to the given string and assigns it to the GcpProjectID field.
func (o *FleetDescribeAccountConfigResult) SetGcpProjectID(v string) {
	o.GcpProjectID = &v
}

// GetGcpProjectNumber returns the GcpProjectNumber field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetGcpProjectNumber() string {
	if o == nil || IsNil(o.GcpProjectNumber) {
		var ret string
		return ret
	}
	return *o.GcpProjectNumber
}

// GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetGcpProjectNumberOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectNumber) {
		return nil, false
	}
	return o.GcpProjectNumber, true
}

// HasGcpProjectNumber returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasGcpProjectNumber() bool {
	if o != nil && !IsNil(o.GcpProjectNumber) {
		return true
	}

	return false
}

// SetGcpProjectNumber gets a reference to the given string and assigns it to the GcpProjectNumber field.
func (o *FleetDescribeAccountConfigResult) SetGcpProjectNumber(v string) {
	o.GcpProjectNumber = &v
}

// GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field value if set, zero value otherwise.
func (o *FleetDescribeAccountConfigResult) GetGcpServiceAccountEmail() string {
	if o == nil || IsNil(o.GcpServiceAccountEmail) {
		var ret string
		return ret
	}
	return *o.GcpServiceAccountEmail
}

// GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetGcpServiceAccountEmailOk() (*string, bool) {
	if o == nil || IsNil(o.GcpServiceAccountEmail) {
		return nil, false
	}
	return o.GcpServiceAccountEmail, true
}

// HasGcpServiceAccountEmail returns a boolean if a field has been set.
func (o *FleetDescribeAccountConfigResult) HasGcpServiceAccountEmail() bool {
	if o != nil && !IsNil(o.GcpServiceAccountEmail) {
		return true
	}

	return false
}

// SetGcpServiceAccountEmail gets a reference to the given string and assigns it to the GcpServiceAccountEmail field.
func (o *FleetDescribeAccountConfigResult) SetGcpServiceAccountEmail(v string) {
	o.GcpServiceAccountEmail = &v
}

// GetId returns the Id field value
func (o *FleetDescribeAccountConfigResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FleetDescribeAccountConfigResult) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FleetDescribeAccountConfigResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FleetDescribeAccountConfigResult) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *FleetDescribeAccountConfigResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FleetDescribeAccountConfigResult) SetStatus(v string) {
	o.Status = v
}

// GetStatusMessage returns the StatusMessage field value
func (o *FleetDescribeAccountConfigResult) GetStatusMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value
// and a boolean to check if the value has been set.
func (o *FleetDescribeAccountConfigResult) GetStatusMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusMessage, true
}

// SetStatusMessage sets field value
func (o *FleetDescribeAccountConfigResult) SetStatusMessage(v string) {
	o.StatusMessage = v
}

func (o FleetDescribeAccountConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FleetDescribeAccountConfigResult) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AzureBootstrapShellCommand) {
		toSerialize["azureBootstrapShellCommand"] = o.AzureBootstrapShellCommand
	}
	if !IsNil(o.AzureDisconnectShellCommand) {
		toSerialize["azureDisconnectShellCommand"] = o.AzureDisconnectShellCommand
	}
	if !IsNil(o.AzureSubscriptionID) {
		toSerialize["azureSubscriptionID"] = o.AzureSubscriptionID
	}
	if !IsNil(o.AzureTenantID) {
		toSerialize["azureTenantID"] = o.AzureTenantID
	}
	if !IsNil(o.ByoaInstanceIDs) {
		toSerialize["byoaInstanceIDs"] = o.ByoaInstanceIDs
	}
	toSerialize["cloudProviderId"] = o.CloudProviderId
	toSerialize["description"] = o.Description
	if !IsNil(o.GcpBootstrapShellCommand) {
		toSerialize["gcpBootstrapShellCommand"] = o.GcpBootstrapShellCommand
	}
	if !IsNil(o.GcpDisconnectShellCommand) {
		toSerialize["gcpDisconnectShellCommand"] = o.GcpDisconnectShellCommand
	}
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

func (o *FleetDescribeAccountConfigResult) UnmarshalJSON(data []byte) (err error) {
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

	varFleetDescribeAccountConfigResult := _FleetDescribeAccountConfigResult{}

	err = json.Unmarshal(data, &varFleetDescribeAccountConfigResult)

	if err != nil {
		return err
	}

	*o = FleetDescribeAccountConfigResult(varFleetDescribeAccountConfigResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "awsAccountID")
		delete(additionalProperties, "awsBootstrapRoleARN")
		delete(additionalProperties, "awsCloudFormationNoLBTemplateURL")
		delete(additionalProperties, "awsCloudFormationTemplateURL")
		delete(additionalProperties, "azureBootstrapShellCommand")
		delete(additionalProperties, "azureDisconnectShellCommand")
		delete(additionalProperties, "azureSubscriptionID")
		delete(additionalProperties, "azureTenantID")
		delete(additionalProperties, "byoaInstanceIDs")
		delete(additionalProperties, "cloudProviderId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "gcpBootstrapShellCommand")
		delete(additionalProperties, "gcpDisconnectShellCommand")
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

type NullableFleetDescribeAccountConfigResult struct {
	value *FleetDescribeAccountConfigResult
	isSet bool
}

func (v NullableFleetDescribeAccountConfigResult) Get() *FleetDescribeAccountConfigResult {
	return v.value
}

func (v *NullableFleetDescribeAccountConfigResult) Set(val *FleetDescribeAccountConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetDescribeAccountConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetDescribeAccountConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetDescribeAccountConfigResult(val *FleetDescribeAccountConfigResult) *NullableFleetDescribeAccountConfigResult {
	return &NullableFleetDescribeAccountConfigResult{value: val, isSet: true}
}

func (v NullableFleetDescribeAccountConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetDescribeAccountConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


