# OnboardingTerraformExecutionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to **[]string** | OCI permissions used to create Terraform execution policy statements. Supported for oci. | [optional] 
**PolicyDocument** | Pointer to **string** | AWS IAM policy document for the Terraform execution identity. Supported for aws. | [optional] 
**Roles** | Pointer to [**[]OnboardingTerraformExecutionRole**](OnboardingTerraformExecutionRole.md) | Roles assigned to the Terraform execution identity. Supported for gcp and azure. | [optional] 

## Methods

### NewOnboardingTerraformExecutionPolicy

`func NewOnboardingTerraformExecutionPolicy() *OnboardingTerraformExecutionPolicy`

NewOnboardingTerraformExecutionPolicy instantiates a new OnboardingTerraformExecutionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingTerraformExecutionPolicyWithDefaults

`func NewOnboardingTerraformExecutionPolicyWithDefaults() *OnboardingTerraformExecutionPolicy`

NewOnboardingTerraformExecutionPolicyWithDefaults instantiates a new OnboardingTerraformExecutionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *OnboardingTerraformExecutionPolicy) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OnboardingTerraformExecutionPolicy) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OnboardingTerraformExecutionPolicy) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OnboardingTerraformExecutionPolicy) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPolicyDocument

`func (o *OnboardingTerraformExecutionPolicy) GetPolicyDocument() string`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *OnboardingTerraformExecutionPolicy) GetPolicyDocumentOk() (*string, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *OnboardingTerraformExecutionPolicy) SetPolicyDocument(v string)`

SetPolicyDocument sets PolicyDocument field to given value.

### HasPolicyDocument

`func (o *OnboardingTerraformExecutionPolicy) HasPolicyDocument() bool`

HasPolicyDocument returns a boolean if a field has been set.

### GetRoles

`func (o *OnboardingTerraformExecutionPolicy) GetRoles() []OnboardingTerraformExecutionRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OnboardingTerraformExecutionPolicy) GetRolesOk() (*[]OnboardingTerraformExecutionRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OnboardingTerraformExecutionPolicy) SetRoles(v []OnboardingTerraformExecutionRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OnboardingTerraformExecutionPolicy) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


