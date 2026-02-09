# TerraformConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactsLocalPath** | Pointer to **string** | The local path to store artifacts generated during Terraform execution | [optional] 
**GitConfiguration** | Pointer to [**GitConfiguration**](GitConfiguration.md) |  | [optional] 
**PrivateModuleGitAccessTokens** | Pointer to **map[string]string** | The git access tokens for private modules | [optional] 
**RequiredOutputKeys** | Pointer to **[]string** | The required output keys to export | [optional] 
**RequiredOutputs** | Pointer to [**[]TerraformOutput**](TerraformOutput.md) | The required output keys to export | [optional] 
**TerraformExecutionIdentity** | Pointer to **string** | The identity to use for terraform execution | [optional] 
**TerraformPath** | **string** | The path to the terraform files directory | 
**VariablesValuesFileOverride** | Pointer to **string** | The variables values file override in base64 format for the Terraform configuration | [optional] 

## Methods

### NewTerraformConfiguration

`func NewTerraformConfiguration(terraformPath string, ) *TerraformConfiguration`

NewTerraformConfiguration instantiates a new TerraformConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformConfigurationWithDefaults

`func NewTerraformConfigurationWithDefaults() *TerraformConfiguration`

NewTerraformConfigurationWithDefaults instantiates a new TerraformConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactsLocalPath

`func (o *TerraformConfiguration) GetArtifactsLocalPath() string`

GetArtifactsLocalPath returns the ArtifactsLocalPath field if non-nil, zero value otherwise.

### GetArtifactsLocalPathOk

`func (o *TerraformConfiguration) GetArtifactsLocalPathOk() (*string, bool)`

GetArtifactsLocalPathOk returns a tuple with the ArtifactsLocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactsLocalPath

`func (o *TerraformConfiguration) SetArtifactsLocalPath(v string)`

SetArtifactsLocalPath sets ArtifactsLocalPath field to given value.

### HasArtifactsLocalPath

`func (o *TerraformConfiguration) HasArtifactsLocalPath() bool`

HasArtifactsLocalPath returns a boolean if a field has been set.

### GetGitConfiguration

`func (o *TerraformConfiguration) GetGitConfiguration() GitConfiguration`

GetGitConfiguration returns the GitConfiguration field if non-nil, zero value otherwise.

### GetGitConfigurationOk

`func (o *TerraformConfiguration) GetGitConfigurationOk() (*GitConfiguration, bool)`

GetGitConfigurationOk returns a tuple with the GitConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitConfiguration

`func (o *TerraformConfiguration) SetGitConfiguration(v GitConfiguration)`

SetGitConfiguration sets GitConfiguration field to given value.

### HasGitConfiguration

`func (o *TerraformConfiguration) HasGitConfiguration() bool`

HasGitConfiguration returns a boolean if a field has been set.

### GetPrivateModuleGitAccessTokens

`func (o *TerraformConfiguration) GetPrivateModuleGitAccessTokens() map[string]string`

GetPrivateModuleGitAccessTokens returns the PrivateModuleGitAccessTokens field if non-nil, zero value otherwise.

### GetPrivateModuleGitAccessTokensOk

`func (o *TerraformConfiguration) GetPrivateModuleGitAccessTokensOk() (*map[string]string, bool)`

GetPrivateModuleGitAccessTokensOk returns a tuple with the PrivateModuleGitAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateModuleGitAccessTokens

`func (o *TerraformConfiguration) SetPrivateModuleGitAccessTokens(v map[string]string)`

SetPrivateModuleGitAccessTokens sets PrivateModuleGitAccessTokens field to given value.

### HasPrivateModuleGitAccessTokens

`func (o *TerraformConfiguration) HasPrivateModuleGitAccessTokens() bool`

HasPrivateModuleGitAccessTokens returns a boolean if a field has been set.

### GetRequiredOutputKeys

`func (o *TerraformConfiguration) GetRequiredOutputKeys() []string`

GetRequiredOutputKeys returns the RequiredOutputKeys field if non-nil, zero value otherwise.

### GetRequiredOutputKeysOk

`func (o *TerraformConfiguration) GetRequiredOutputKeysOk() (*[]string, bool)`

GetRequiredOutputKeysOk returns a tuple with the RequiredOutputKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredOutputKeys

`func (o *TerraformConfiguration) SetRequiredOutputKeys(v []string)`

SetRequiredOutputKeys sets RequiredOutputKeys field to given value.

### HasRequiredOutputKeys

`func (o *TerraformConfiguration) HasRequiredOutputKeys() bool`

HasRequiredOutputKeys returns a boolean if a field has been set.

### GetRequiredOutputs

`func (o *TerraformConfiguration) GetRequiredOutputs() []TerraformOutput`

GetRequiredOutputs returns the RequiredOutputs field if non-nil, zero value otherwise.

### GetRequiredOutputsOk

`func (o *TerraformConfiguration) GetRequiredOutputsOk() (*[]TerraformOutput, bool)`

GetRequiredOutputsOk returns a tuple with the RequiredOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredOutputs

`func (o *TerraformConfiguration) SetRequiredOutputs(v []TerraformOutput)`

SetRequiredOutputs sets RequiredOutputs field to given value.

### HasRequiredOutputs

`func (o *TerraformConfiguration) HasRequiredOutputs() bool`

HasRequiredOutputs returns a boolean if a field has been set.

### GetTerraformExecutionIdentity

`func (o *TerraformConfiguration) GetTerraformExecutionIdentity() string`

GetTerraformExecutionIdentity returns the TerraformExecutionIdentity field if non-nil, zero value otherwise.

### GetTerraformExecutionIdentityOk

`func (o *TerraformConfiguration) GetTerraformExecutionIdentityOk() (*string, bool)`

GetTerraformExecutionIdentityOk returns a tuple with the TerraformExecutionIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformExecutionIdentity

`func (o *TerraformConfiguration) SetTerraformExecutionIdentity(v string)`

SetTerraformExecutionIdentity sets TerraformExecutionIdentity field to given value.

### HasTerraformExecutionIdentity

`func (o *TerraformConfiguration) HasTerraformExecutionIdentity() bool`

HasTerraformExecutionIdentity returns a boolean if a field has been set.

### GetTerraformPath

`func (o *TerraformConfiguration) GetTerraformPath() string`

GetTerraformPath returns the TerraformPath field if non-nil, zero value otherwise.

### GetTerraformPathOk

`func (o *TerraformConfiguration) GetTerraformPathOk() (*string, bool)`

GetTerraformPathOk returns a tuple with the TerraformPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformPath

`func (o *TerraformConfiguration) SetTerraformPath(v string)`

SetTerraformPath sets TerraformPath field to given value.


### GetVariablesValuesFileOverride

`func (o *TerraformConfiguration) GetVariablesValuesFileOverride() string`

GetVariablesValuesFileOverride returns the VariablesValuesFileOverride field if non-nil, zero value otherwise.

### GetVariablesValuesFileOverrideOk

`func (o *TerraformConfiguration) GetVariablesValuesFileOverrideOk() (*string, bool)`

GetVariablesValuesFileOverrideOk returns a tuple with the VariablesValuesFileOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesValuesFileOverride

`func (o *TerraformConfiguration) SetVariablesValuesFileOverride(v string)`

SetVariablesValuesFileOverride sets VariablesValuesFileOverride field to given value.

### HasVariablesValuesFileOverride

`func (o *TerraformConfiguration) HasVariablesValuesFileOverride() bool`

HasVariablesValuesFileOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


