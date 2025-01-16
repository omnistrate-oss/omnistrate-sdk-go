# CreateServiceFromComposeSpecRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the service | 
**FileContent** | **string** | Base64 encoded Compose Spec YAML in docker compose format | 
**FileFormat** | **string** | MIME type of file format | 
**FileName** | **string** | Name of compose spec YAML file that is uploaded | 
**Name** | **string** | Name of the Service | 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 

## Methods

### NewCreateServiceFromComposeSpecRequest2

`func NewCreateServiceFromComposeSpecRequest2(description string, fileContent string, fileFormat string, fileName string, name string, ) *CreateServiceFromComposeSpecRequest2`

NewCreateServiceFromComposeSpecRequest2 instantiates a new CreateServiceFromComposeSpecRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceFromComposeSpecRequest2WithDefaults

`func NewCreateServiceFromComposeSpecRequest2WithDefaults() *CreateServiceFromComposeSpecRequest2`

NewCreateServiceFromComposeSpecRequest2WithDefaults instantiates a new CreateServiceFromComposeSpecRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateServiceFromComposeSpecRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceFromComposeSpecRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceFromComposeSpecRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFileContent

`func (o *CreateServiceFromComposeSpecRequest2) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *CreateServiceFromComposeSpecRequest2) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *CreateServiceFromComposeSpecRequest2) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetFileFormat

`func (o *CreateServiceFromComposeSpecRequest2) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *CreateServiceFromComposeSpecRequest2) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *CreateServiceFromComposeSpecRequest2) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.


### GetFileName

`func (o *CreateServiceFromComposeSpecRequest2) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CreateServiceFromComposeSpecRequest2) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CreateServiceFromComposeSpecRequest2) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetName

`func (o *CreateServiceFromComposeSpecRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceFromComposeSpecRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceFromComposeSpecRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetServiceLogoURL

`func (o *CreateServiceFromComposeSpecRequest2) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *CreateServiceFromComposeSpecRequest2) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *CreateServiceFromComposeSpecRequest2) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *CreateServiceFromComposeSpecRequest2) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


