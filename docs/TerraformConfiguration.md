# TerraformConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitConfiguration** | Pointer to [**GitConfiguration**](GitConfiguration.md) |  | [optional] 
**TerraformPath** | **string** | The path to the terraform files directory | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


