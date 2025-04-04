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

// checks if the CreateAccountConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountConfigRequest{}

// CreateAccountConfigRequest Account configuration including the role required to access
type CreateAccountConfigRequest struct {
	// The AWS access key
	AwsAccessKey *string `json:"awsAccessKey,omitempty"`
	// The AWS account ID
	AwsAccountID *string `json:"awsAccountID,omitempty"`
	// The security role ARN or service account ARN that grants access to operate the infra
	AwsBootstrapRoleARN *string `json:"awsBootstrapRoleARN,omitempty"`
	// The AWS secret key
	AwsSecretKey *string `json:"awsSecretKey,omitempty"`
	// The Azure subscription ID
	AzureSubscriptionID *string `json:"azureSubscriptionID,omitempty"`
	// The Azure tenant ID
	AzureTenantID *string `json:"azureTenantID,omitempty"`
	// The BYOA instance ID that this account config is tied to
	ByoaInstanceID *string `json:"byoaInstanceID,omitempty"`
	// ID of an CloudProvider
	CloudProviderId string `json:"cloudProviderId"`
	// The description for the account
	Description string `json:"description"`
	// The GCP project ID
	GcpProjectID *string `json:"gcpProjectID,omitempty"`
	// The GCP project number
	GcpProjectNumber *string `json:"gcpProjectNumber,omitempty"`
	// The GCP service account email
	GcpServiceAccountEmail *string `json:"gcpServiceAccountEmail,omitempty"`
	// The GCP service account key
	GcpServiceAccountKey *string `json:"gcpServiceAccountKey,omitempty"`
	// The name of the account
	Name string `json:"name"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _CreateAccountConfigRequest CreateAccountConfigRequest

// NewCreateAccountConfigRequest instantiates a new CreateAccountConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountConfigRequest(cloudProviderId string, description string, name string, token string) *CreateAccountConfigRequest {
	this := CreateAccountConfigRequest{}
	this.CloudProviderId = cloudProviderId
	this.Description = description
	this.Name = name
	this.Token = token
	return &this
}

// NewCreateAccountConfigRequestWithDefaults instantiates a new CreateAccountConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountConfigRequestWithDefaults() *CreateAccountConfigRequest {
	this := CreateAccountConfigRequest{}
	return &this
}

// GetAwsAccessKey returns the AwsAccessKey field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetAwsAccessKey() string {
	if o == nil || IsNil(o.AwsAccessKey) {
		var ret string
		return ret
	}
	return *o.AwsAccessKey
}

// GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetAwsAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccessKey) {
		return nil, false
	}
	return o.AwsAccessKey, true
}

// SetAwsAccessKey gets a reference to the given string and assigns it to the AwsAccessKey field.
func (o *CreateAccountConfigRequest) SetAwsAccessKey(v string) {
	o.AwsAccessKey = &v
}

// GetAwsAccountID returns the AwsAccountID field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetAwsAccountID() string {
	if o == nil || IsNil(o.AwsAccountID) {
		var ret string
		return ret
	}
	return *o.AwsAccountID
}

// GetAwsAccountIDOk returns a tuple with the AwsAccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetAwsAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountID) {
		return nil, false
	}
	return o.AwsAccountID, true
}

// SetAwsAccountID gets a reference to the given string and assigns it to the AwsAccountID field.
func (o *CreateAccountConfigRequest) SetAwsAccountID(v string) {
	o.AwsAccountID = &v
}

// GetAwsBootstrapRoleARN returns the AwsBootstrapRoleARN field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetAwsBootstrapRoleARN() string {
	if o == nil || IsNil(o.AwsBootstrapRoleARN) {
		var ret string
		return ret
	}
	return *o.AwsBootstrapRoleARN
}

// GetAwsBootstrapRoleARNOk returns a tuple with the AwsBootstrapRoleARN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetAwsBootstrapRoleARNOk() (*string, bool) {
	if o == nil || IsNil(o.AwsBootstrapRoleARN) {
		return nil, false
	}
	return o.AwsBootstrapRoleARN, true
}

// SetAwsBootstrapRoleARN gets a reference to the given string and assigns it to the AwsBootstrapRoleARN field.
func (o *CreateAccountConfigRequest) SetAwsBootstrapRoleARN(v string) {
	o.AwsBootstrapRoleARN = &v
}

// GetAwsSecretKey returns the AwsSecretKey field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetAwsSecretKey() string {
	if o == nil || IsNil(o.AwsSecretKey) {
		var ret string
		return ret
	}
	return *o.AwsSecretKey
}

// GetAwsSecretKeyOk returns a tuple with the AwsSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetAwsSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AwsSecretKey) {
		return nil, false
	}
	return o.AwsSecretKey, true
}

// SetAwsSecretKey gets a reference to the given string and assigns it to the AwsSecretKey field.
func (o *CreateAccountConfigRequest) SetAwsSecretKey(v string) {
	o.AwsSecretKey = &v
}

// GetAzureSubscriptionID returns the AzureSubscriptionID field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetAzureSubscriptionID() string {
	if o == nil || IsNil(o.AzureSubscriptionID) {
		var ret string
		return ret
	}
	return *o.AzureSubscriptionID
}

// GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetAzureSubscriptionIDOk() (*string, bool) {
	if o == nil || IsNil(o.AzureSubscriptionID) {
		return nil, false
	}
	return o.AzureSubscriptionID, true
}

// SetAzureSubscriptionID gets a reference to the given string and assigns it to the AzureSubscriptionID field.
func (o *CreateAccountConfigRequest) SetAzureSubscriptionID(v string) {
	o.AzureSubscriptionID = &v
}

// GetAzureTenantID returns the AzureTenantID field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetAzureTenantID() string {
	if o == nil || IsNil(o.AzureTenantID) {
		var ret string
		return ret
	}
	return *o.AzureTenantID
}

// GetAzureTenantIDOk returns a tuple with the AzureTenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetAzureTenantIDOk() (*string, bool) {
	if o == nil || IsNil(o.AzureTenantID) {
		return nil, false
	}
	return o.AzureTenantID, true
}

// SetAzureTenantID gets a reference to the given string and assigns it to the AzureTenantID field.
func (o *CreateAccountConfigRequest) SetAzureTenantID(v string) {
	o.AzureTenantID = &v
}

// GetByoaInstanceID returns the ByoaInstanceID field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetByoaInstanceID() string {
	if o == nil || IsNil(o.ByoaInstanceID) {
		var ret string
		return ret
	}
	return *o.ByoaInstanceID
}

// GetByoaInstanceIDOk returns a tuple with the ByoaInstanceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetByoaInstanceIDOk() (*string, bool) {
	if o == nil || IsNil(o.ByoaInstanceID) {
		return nil, false
	}
	return o.ByoaInstanceID, true
}

// SetByoaInstanceID gets a reference to the given string and assigns it to the ByoaInstanceID field.
func (o *CreateAccountConfigRequest) SetByoaInstanceID(v string) {
	o.ByoaInstanceID = &v
}

// GetCloudProviderId returns the CloudProviderId field value
func (o *CreateAccountConfigRequest) GetCloudProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProviderId
}

// GetCloudProviderIdOk returns a tuple with the CloudProviderId field value
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetCloudProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProviderId, true
}

// SetCloudProviderId sets field value
func (o *CreateAccountConfigRequest) SetCloudProviderId(v string) {
	o.CloudProviderId = v
}

// GetDescription returns the Description field value
func (o *CreateAccountConfigRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateAccountConfigRequest) SetDescription(v string) {
	o.Description = v
}

// GetGcpProjectID returns the GcpProjectID field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetGcpProjectID() string {
	if o == nil || IsNil(o.GcpProjectID) {
		var ret string
		return ret
	}
	return *o.GcpProjectID
}

// GetGcpProjectIDOk returns a tuple with the GcpProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetGcpProjectIDOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectID) {
		return nil, false
	}
	return o.GcpProjectID, true
}

// SetGcpProjectID gets a reference to the given string and assigns it to the GcpProjectID field.
func (o *CreateAccountConfigRequest) SetGcpProjectID(v string) {
	o.GcpProjectID = &v
}

// GetGcpProjectNumber returns the GcpProjectNumber field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetGcpProjectNumber() string {
	if o == nil || IsNil(o.GcpProjectNumber) {
		var ret string
		return ret
	}
	return *o.GcpProjectNumber
}

// GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetGcpProjectNumberOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectNumber) {
		return nil, false
	}
	return o.GcpProjectNumber, true
}

// SetGcpProjectNumber gets a reference to the given string and assigns it to the GcpProjectNumber field.
func (o *CreateAccountConfigRequest) SetGcpProjectNumber(v string) {
	o.GcpProjectNumber = &v
}

// GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetGcpServiceAccountEmail() string {
	if o == nil || IsNil(o.GcpServiceAccountEmail) {
		var ret string
		return ret
	}
	return *o.GcpServiceAccountEmail
}

// GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetGcpServiceAccountEmailOk() (*string, bool) {
	if o == nil || IsNil(o.GcpServiceAccountEmail) {
		return nil, false
	}
	return o.GcpServiceAccountEmail, true
}

// SetGcpServiceAccountEmail gets a reference to the given string and assigns it to the GcpServiceAccountEmail field.
func (o *CreateAccountConfigRequest) SetGcpServiceAccountEmail(v string) {
	o.GcpServiceAccountEmail = &v
}

// GetGcpServiceAccountKey returns the GcpServiceAccountKey field value if set, zero value otherwise.
func (o *CreateAccountConfigRequest) GetGcpServiceAccountKey() string {
	if o == nil || IsNil(o.GcpServiceAccountKey) {
		var ret string
		return ret
	}
	return *o.GcpServiceAccountKey
}

// GetGcpServiceAccountKeyOk returns a tuple with the GcpServiceAccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetGcpServiceAccountKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GcpServiceAccountKey) {
		return nil, false
	}
	return o.GcpServiceAccountKey, true
}

// SetGcpServiceAccountKey gets a reference to the given string and assigns it to the GcpServiceAccountKey field.
func (o *CreateAccountConfigRequest) SetGcpServiceAccountKey(v string) {
	o.GcpServiceAccountKey = &v
}

// GetName returns the Name field value
func (o *CreateAccountConfigRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAccountConfigRequest) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value
func (o *CreateAccountConfigRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CreateAccountConfigRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CreateAccountConfigRequest) SetToken(v string) {
	o.Token = v
}

func (o CreateAccountConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsAccessKey) {
		toSerialize["awsAccessKey"] = o.AwsAccessKey
	}
	if !IsNil(o.AwsAccountID) {
		toSerialize["awsAccountID"] = o.AwsAccountID
	}
	if !IsNil(o.AwsBootstrapRoleARN) {
		toSerialize["awsBootstrapRoleARN"] = o.AwsBootstrapRoleARN
	}
	if !IsNil(o.AwsSecretKey) {
		toSerialize["awsSecretKey"] = o.AwsSecretKey
	}
	if !IsNil(o.AzureSubscriptionID) {
		toSerialize["azureSubscriptionID"] = o.AzureSubscriptionID
	}
	if !IsNil(o.AzureTenantID) {
		toSerialize["azureTenantID"] = o.AzureTenantID
	}
	if !IsNil(o.ByoaInstanceID) {
		toSerialize["byoaInstanceID"] = o.ByoaInstanceID
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
	if !IsNil(o.GcpServiceAccountKey) {
		toSerialize["gcpServiceAccountKey"] = o.GcpServiceAccountKey
	}
	toSerialize["name"] = o.Name
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAccountConfigRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudProviderId",
		"description",
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

	varCreateAccountConfigRequest := _CreateAccountConfigRequest{}

	err = json.Unmarshal(data, &varCreateAccountConfigRequest)

	if err != nil {
		return err
	}

	*o = CreateAccountConfigRequest(varCreateAccountConfigRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "awsAccessKey")
		delete(additionalProperties, "awsAccountID")
		delete(additionalProperties, "awsBootstrapRoleARN")
		delete(additionalProperties, "awsSecretKey")
		delete(additionalProperties, "azureSubscriptionID")
		delete(additionalProperties, "azureTenantID")
		delete(additionalProperties, "byoaInstanceID")
		delete(additionalProperties, "cloudProviderId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "gcpProjectID")
		delete(additionalProperties, "gcpProjectNumber")
		delete(additionalProperties, "gcpServiceAccountEmail")
		delete(additionalProperties, "gcpServiceAccountKey")
		delete(additionalProperties, "name")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAccountConfigRequest struct {
	value *CreateAccountConfigRequest
	isSet bool
}

func (v NullableCreateAccountConfigRequest) Get() *CreateAccountConfigRequest {
	return v.value
}

func (v *NullableCreateAccountConfigRequest) Set(val *CreateAccountConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountConfigRequest(val *CreateAccountConfigRequest) *NullableCreateAccountConfigRequest {
	return &NullableCreateAccountConfigRequest{value: val, isSet: true}
}

func (v NullableCreateAccountConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


