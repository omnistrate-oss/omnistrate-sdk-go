# TerraformDeploymentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationFiles** | Pointer to **map[string]string** | Deployed files for Terraform deployment | [optional] 
**DeploymentErrors** | Pointer to **string** | Errors encountered during the Terraform deployment | [optional] 

## Methods

### NewTerraformDeploymentConfiguration

`func NewTerraformDeploymentConfiguration() *TerraformDeploymentConfiguration`

NewTerraformDeploymentConfiguration instantiates a new TerraformDeploymentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformDeploymentConfigurationWithDefaults

`func NewTerraformDeploymentConfigurationWithDefaults() *TerraformDeploymentConfiguration`

NewTerraformDeploymentConfigurationWithDefaults instantiates a new TerraformDeploymentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationFiles

`func (o *TerraformDeploymentConfiguration) GetConfigurationFiles() map[string]string`

GetConfigurationFiles returns the ConfigurationFiles field if non-nil, zero value otherwise.

### GetConfigurationFilesOk

`func (o *TerraformDeploymentConfiguration) GetConfigurationFilesOk() (*map[string]string, bool)`

GetConfigurationFilesOk returns a tuple with the ConfigurationFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFiles

`func (o *TerraformDeploymentConfiguration) SetConfigurationFiles(v map[string]string)`

SetConfigurationFiles sets ConfigurationFiles field to given value.

### HasConfigurationFiles

`func (o *TerraformDeploymentConfiguration) HasConfigurationFiles() bool`

HasConfigurationFiles returns a boolean if a field has been set.

### GetDeploymentErrors

`func (o *TerraformDeploymentConfiguration) GetDeploymentErrors() string`

GetDeploymentErrors returns the DeploymentErrors field if non-nil, zero value otherwise.

### GetDeploymentErrorsOk

`func (o *TerraformDeploymentConfiguration) GetDeploymentErrorsOk() (*string, bool)`

GetDeploymentErrorsOk returns a tuple with the DeploymentErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentErrors

`func (o *TerraformDeploymentConfiguration) SetDeploymentErrors(v string)`

SetDeploymentErrors sets DeploymentErrors field to given value.

### HasDeploymentErrors

`func (o *TerraformDeploymentConfiguration) HasDeploymentErrors() bool`

HasDeploymentErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


