# OnboardingTerraformExecutionRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The role name. | 
**Type** | Pointer to **string** | The role type. Default is used for GCP IAM roles and Azure RBAC roles. Entra is used for Azure Entra roles. | [optional] 

## Methods

### NewOnboardingTerraformExecutionRole

`func NewOnboardingTerraformExecutionRole(name string, ) *OnboardingTerraformExecutionRole`

NewOnboardingTerraformExecutionRole instantiates a new OnboardingTerraformExecutionRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingTerraformExecutionRoleWithDefaults

`func NewOnboardingTerraformExecutionRoleWithDefaults() *OnboardingTerraformExecutionRole`

NewOnboardingTerraformExecutionRoleWithDefaults instantiates a new OnboardingTerraformExecutionRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OnboardingTerraformExecutionRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnboardingTerraformExecutionRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnboardingTerraformExecutionRole) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *OnboardingTerraformExecutionRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnboardingTerraformExecutionRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnboardingTerraformExecutionRole) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OnboardingTerraformExecutionRole) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


