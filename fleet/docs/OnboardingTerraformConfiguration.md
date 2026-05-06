# OnboardingTerraformConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactsLocalPath** | Pointer to **string** | The local path where the terraform artifacts are stored. | [optional] 
**GitConfiguration** | Pointer to [**GitConfiguration**](GitConfiguration.md) |  | [optional] 
**PrivateKeyPEM** | Pointer to **string** | The private key PEM for terraform execution. | [optional] 
**PrivateModuleGitAccessTokens** | Pointer to **map[string]string** | Git access tokens for private modules. | [optional] 
**PublicKeyID** | Pointer to **string** | The public key ID paired with the terraform service account private key. | [optional] 
**RequiredOutputs** | Pointer to [**[]OnboardingTerraformOutput**](OnboardingTerraformOutput.md) | The required output keys to export. | [optional] 
**ServiceAccountID** | Pointer to **string** | The service account ID for terraform execution. | [optional] 
**TerraformExecutionIdentity** | Pointer to **string** | The identity to use for terraform execution. | [optional] 
**TerraformPath** | **string** | The path to the terraform files directory. | 
**VariablesValuesFileOverride** | Pointer to **string** | The variables values file override in base64 format. | [optional] 

## Methods

### NewOnboardingTerraformConfiguration

`func NewOnboardingTerraformConfiguration(terraformPath string, ) *OnboardingTerraformConfiguration`

NewOnboardingTerraformConfiguration instantiates a new OnboardingTerraformConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingTerraformConfigurationWithDefaults

`func NewOnboardingTerraformConfigurationWithDefaults() *OnboardingTerraformConfiguration`

NewOnboardingTerraformConfigurationWithDefaults instantiates a new OnboardingTerraformConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactsLocalPath

`func (o *OnboardingTerraformConfiguration) GetArtifactsLocalPath() string`

GetArtifactsLocalPath returns the ArtifactsLocalPath field if non-nil, zero value otherwise.

### GetArtifactsLocalPathOk

`func (o *OnboardingTerraformConfiguration) GetArtifactsLocalPathOk() (*string, bool)`

GetArtifactsLocalPathOk returns a tuple with the ArtifactsLocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactsLocalPath

`func (o *OnboardingTerraformConfiguration) SetArtifactsLocalPath(v string)`

SetArtifactsLocalPath sets ArtifactsLocalPath field to given value.

### HasArtifactsLocalPath

`func (o *OnboardingTerraformConfiguration) HasArtifactsLocalPath() bool`

HasArtifactsLocalPath returns a boolean if a field has been set.

### GetGitConfiguration

`func (o *OnboardingTerraformConfiguration) GetGitConfiguration() GitConfiguration`

GetGitConfiguration returns the GitConfiguration field if non-nil, zero value otherwise.

### GetGitConfigurationOk

`func (o *OnboardingTerraformConfiguration) GetGitConfigurationOk() (*GitConfiguration, bool)`

GetGitConfigurationOk returns a tuple with the GitConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitConfiguration

`func (o *OnboardingTerraformConfiguration) SetGitConfiguration(v GitConfiguration)`

SetGitConfiguration sets GitConfiguration field to given value.

### HasGitConfiguration

`func (o *OnboardingTerraformConfiguration) HasGitConfiguration() bool`

HasGitConfiguration returns a boolean if a field has been set.

### GetPrivateKeyPEM

`func (o *OnboardingTerraformConfiguration) GetPrivateKeyPEM() string`

GetPrivateKeyPEM returns the PrivateKeyPEM field if non-nil, zero value otherwise.

### GetPrivateKeyPEMOk

`func (o *OnboardingTerraformConfiguration) GetPrivateKeyPEMOk() (*string, bool)`

GetPrivateKeyPEMOk returns a tuple with the PrivateKeyPEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPEM

`func (o *OnboardingTerraformConfiguration) SetPrivateKeyPEM(v string)`

SetPrivateKeyPEM sets PrivateKeyPEM field to given value.

### HasPrivateKeyPEM

`func (o *OnboardingTerraformConfiguration) HasPrivateKeyPEM() bool`

HasPrivateKeyPEM returns a boolean if a field has been set.

### GetPrivateModuleGitAccessTokens

`func (o *OnboardingTerraformConfiguration) GetPrivateModuleGitAccessTokens() map[string]string`

GetPrivateModuleGitAccessTokens returns the PrivateModuleGitAccessTokens field if non-nil, zero value otherwise.

### GetPrivateModuleGitAccessTokensOk

`func (o *OnboardingTerraformConfiguration) GetPrivateModuleGitAccessTokensOk() (*map[string]string, bool)`

GetPrivateModuleGitAccessTokensOk returns a tuple with the PrivateModuleGitAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateModuleGitAccessTokens

`func (o *OnboardingTerraformConfiguration) SetPrivateModuleGitAccessTokens(v map[string]string)`

SetPrivateModuleGitAccessTokens sets PrivateModuleGitAccessTokens field to given value.

### HasPrivateModuleGitAccessTokens

`func (o *OnboardingTerraformConfiguration) HasPrivateModuleGitAccessTokens() bool`

HasPrivateModuleGitAccessTokens returns a boolean if a field has been set.

### GetPublicKeyID

`func (o *OnboardingTerraformConfiguration) GetPublicKeyID() string`

GetPublicKeyID returns the PublicKeyID field if non-nil, zero value otherwise.

### GetPublicKeyIDOk

`func (o *OnboardingTerraformConfiguration) GetPublicKeyIDOk() (*string, bool)`

GetPublicKeyIDOk returns a tuple with the PublicKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyID

`func (o *OnboardingTerraformConfiguration) SetPublicKeyID(v string)`

SetPublicKeyID sets PublicKeyID field to given value.

### HasPublicKeyID

`func (o *OnboardingTerraformConfiguration) HasPublicKeyID() bool`

HasPublicKeyID returns a boolean if a field has been set.

### GetRequiredOutputs

`func (o *OnboardingTerraformConfiguration) GetRequiredOutputs() []OnboardingTerraformOutput`

GetRequiredOutputs returns the RequiredOutputs field if non-nil, zero value otherwise.

### GetRequiredOutputsOk

`func (o *OnboardingTerraformConfiguration) GetRequiredOutputsOk() (*[]OnboardingTerraformOutput, bool)`

GetRequiredOutputsOk returns a tuple with the RequiredOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredOutputs

`func (o *OnboardingTerraformConfiguration) SetRequiredOutputs(v []OnboardingTerraformOutput)`

SetRequiredOutputs sets RequiredOutputs field to given value.

### HasRequiredOutputs

`func (o *OnboardingTerraformConfiguration) HasRequiredOutputs() bool`

HasRequiredOutputs returns a boolean if a field has been set.

### GetServiceAccountID

`func (o *OnboardingTerraformConfiguration) GetServiceAccountID() string`

GetServiceAccountID returns the ServiceAccountID field if non-nil, zero value otherwise.

### GetServiceAccountIDOk

`func (o *OnboardingTerraformConfiguration) GetServiceAccountIDOk() (*string, bool)`

GetServiceAccountIDOk returns a tuple with the ServiceAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountID

`func (o *OnboardingTerraformConfiguration) SetServiceAccountID(v string)`

SetServiceAccountID sets ServiceAccountID field to given value.

### HasServiceAccountID

`func (o *OnboardingTerraformConfiguration) HasServiceAccountID() bool`

HasServiceAccountID returns a boolean if a field has been set.

### GetTerraformExecutionIdentity

`func (o *OnboardingTerraformConfiguration) GetTerraformExecutionIdentity() string`

GetTerraformExecutionIdentity returns the TerraformExecutionIdentity field if non-nil, zero value otherwise.

### GetTerraformExecutionIdentityOk

`func (o *OnboardingTerraformConfiguration) GetTerraformExecutionIdentityOk() (*string, bool)`

GetTerraformExecutionIdentityOk returns a tuple with the TerraformExecutionIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformExecutionIdentity

`func (o *OnboardingTerraformConfiguration) SetTerraformExecutionIdentity(v string)`

SetTerraformExecutionIdentity sets TerraformExecutionIdentity field to given value.

### HasTerraformExecutionIdentity

`func (o *OnboardingTerraformConfiguration) HasTerraformExecutionIdentity() bool`

HasTerraformExecutionIdentity returns a boolean if a field has been set.

### GetTerraformPath

`func (o *OnboardingTerraformConfiguration) GetTerraformPath() string`

GetTerraformPath returns the TerraformPath field if non-nil, zero value otherwise.

### GetTerraformPathOk

`func (o *OnboardingTerraformConfiguration) GetTerraformPathOk() (*string, bool)`

GetTerraformPathOk returns a tuple with the TerraformPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformPath

`func (o *OnboardingTerraformConfiguration) SetTerraformPath(v string)`

SetTerraformPath sets TerraformPath field to given value.


### GetVariablesValuesFileOverride

`func (o *OnboardingTerraformConfiguration) GetVariablesValuesFileOverride() string`

GetVariablesValuesFileOverride returns the VariablesValuesFileOverride field if non-nil, zero value otherwise.

### GetVariablesValuesFileOverrideOk

`func (o *OnboardingTerraformConfiguration) GetVariablesValuesFileOverrideOk() (*string, bool)`

GetVariablesValuesFileOverrideOk returns a tuple with the VariablesValuesFileOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesValuesFileOverride

`func (o *OnboardingTerraformConfiguration) SetVariablesValuesFileOverride(v string)`

SetVariablesValuesFileOverride sets VariablesValuesFileOverride field to given value.

### HasVariablesValuesFileOverride

`func (o *OnboardingTerraformConfiguration) HasVariablesValuesFileOverride() bool`

HasVariablesValuesFileOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


