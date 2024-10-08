# CreateServiceFromComposeSpecRequestBody

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

### NewCreateServiceFromComposeSpecRequestBody

`func NewCreateServiceFromComposeSpecRequestBody(description string, fileContent string, fileFormat string, fileName string, name string, ) *CreateServiceFromComposeSpecRequestBody`

NewCreateServiceFromComposeSpecRequestBody instantiates a new CreateServiceFromComposeSpecRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceFromComposeSpecRequestBodyWithDefaults

`func NewCreateServiceFromComposeSpecRequestBodyWithDefaults() *CreateServiceFromComposeSpecRequestBody`

NewCreateServiceFromComposeSpecRequestBodyWithDefaults instantiates a new CreateServiceFromComposeSpecRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateServiceFromComposeSpecRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceFromComposeSpecRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceFromComposeSpecRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFileContent

`func (o *CreateServiceFromComposeSpecRequestBody) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *CreateServiceFromComposeSpecRequestBody) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *CreateServiceFromComposeSpecRequestBody) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetFileFormat

`func (o *CreateServiceFromComposeSpecRequestBody) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *CreateServiceFromComposeSpecRequestBody) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *CreateServiceFromComposeSpecRequestBody) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.


### GetFileName

`func (o *CreateServiceFromComposeSpecRequestBody) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CreateServiceFromComposeSpecRequestBody) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CreateServiceFromComposeSpecRequestBody) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetName

`func (o *CreateServiceFromComposeSpecRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceFromComposeSpecRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceFromComposeSpecRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetServiceLogoURL

`func (o *CreateServiceFromComposeSpecRequestBody) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *CreateServiceFromComposeSpecRequestBody) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *CreateServiceFromComposeSpecRequestBody) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *CreateServiceFromComposeSpecRequestBody) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


