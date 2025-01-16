# CreateServiceFromComposeSpecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the service | 
**FileContent** | **string** | Base64 encoded Compose Spec YAML in docker compose format | 
**FileFormat** | **string** | MIME type of file format | 
**FileName** | **string** | Name of compose spec YAML file that is uploaded | 
**Name** | **string** | Name of the Service | 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateServiceFromComposeSpecRequest

`func NewCreateServiceFromComposeSpecRequest(description string, fileContent string, fileFormat string, fileName string, name string, token string, ) *CreateServiceFromComposeSpecRequest`

NewCreateServiceFromComposeSpecRequest instantiates a new CreateServiceFromComposeSpecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceFromComposeSpecRequestWithDefaults

`func NewCreateServiceFromComposeSpecRequestWithDefaults() *CreateServiceFromComposeSpecRequest`

NewCreateServiceFromComposeSpecRequestWithDefaults instantiates a new CreateServiceFromComposeSpecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateServiceFromComposeSpecRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceFromComposeSpecRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceFromComposeSpecRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFileContent

`func (o *CreateServiceFromComposeSpecRequest) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *CreateServiceFromComposeSpecRequest) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *CreateServiceFromComposeSpecRequest) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.


### GetFileFormat

`func (o *CreateServiceFromComposeSpecRequest) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *CreateServiceFromComposeSpecRequest) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *CreateServiceFromComposeSpecRequest) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.


### GetFileName

`func (o *CreateServiceFromComposeSpecRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CreateServiceFromComposeSpecRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CreateServiceFromComposeSpecRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetName

`func (o *CreateServiceFromComposeSpecRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceFromComposeSpecRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceFromComposeSpecRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServiceLogoURL

`func (o *CreateServiceFromComposeSpecRequest) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *CreateServiceFromComposeSpecRequest) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *CreateServiceFromComposeSpecRequest) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *CreateServiceFromComposeSpecRequest) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.

### GetToken

`func (o *CreateServiceFromComposeSpecRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateServiceFromComposeSpecRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateServiceFromComposeSpecRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


