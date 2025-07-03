# ValuesFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitConfiguration** | [**GitConfiguration**](GitConfiguration.md) |  | 
**Path** | **string** | The relative file path from repository root | 

## Methods

### NewValuesFile

`func NewValuesFile(gitConfiguration GitConfiguration, path string, ) *ValuesFile`

NewValuesFile instantiates a new ValuesFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuesFileWithDefaults

`func NewValuesFileWithDefaults() *ValuesFile`

NewValuesFileWithDefaults instantiates a new ValuesFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitConfiguration

`func (o *ValuesFile) GetGitConfiguration() GitConfiguration`

GetGitConfiguration returns the GitConfiguration field if non-nil, zero value otherwise.

### GetGitConfigurationOk

`func (o *ValuesFile) GetGitConfigurationOk() (*GitConfiguration, bool)`

GetGitConfigurationOk returns a tuple with the GitConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitConfiguration

`func (o *ValuesFile) SetGitConfiguration(v GitConfiguration)`

SetGitConfiguration sets GitConfiguration field to given value.


### GetPath

`func (o *ValuesFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ValuesFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ValuesFile) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


