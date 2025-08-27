# ContainerImageLocator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ContainerImageRegistryCredentials**](ContainerImageRegistryCredentials.md) |  | [optional] 
**ImageName** | **string** | The name of the container image | 
**ImageTag** | **string** | The tag of the container image | 
**RegistryURL** | Pointer to **string** | The URL of the container image registry (if not provided, defaults to the DockerHub) | [optional] 
**RepositoryName** | **string** | The name of the container image repository | 

## Methods

### NewContainerImageLocator

`func NewContainerImageLocator(imageName string, imageTag string, repositoryName string, ) *ContainerImageLocator`

NewContainerImageLocator instantiates a new ContainerImageLocator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageLocatorWithDefaults

`func NewContainerImageLocatorWithDefaults() *ContainerImageLocator`

NewContainerImageLocatorWithDefaults instantiates a new ContainerImageLocator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ContainerImageLocator) GetCredentials() ContainerImageRegistryCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ContainerImageLocator) GetCredentialsOk() (*ContainerImageRegistryCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ContainerImageLocator) SetCredentials(v ContainerImageRegistryCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ContainerImageLocator) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetImageName

`func (o *ContainerImageLocator) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ContainerImageLocator) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ContainerImageLocator) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetImageTag

`func (o *ContainerImageLocator) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ContainerImageLocator) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ContainerImageLocator) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.


### GetRegistryURL

`func (o *ContainerImageLocator) GetRegistryURL() string`

GetRegistryURL returns the RegistryURL field if non-nil, zero value otherwise.

### GetRegistryURLOk

`func (o *ContainerImageLocator) GetRegistryURLOk() (*string, bool)`

GetRegistryURLOk returns a tuple with the RegistryURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryURL

`func (o *ContainerImageLocator) SetRegistryURL(v string)`

SetRegistryURL sets RegistryURL field to given value.

### HasRegistryURL

`func (o *ContainerImageLocator) HasRegistryURL() bool`

HasRegistryURL returns a boolean if a field has been set.

### GetRepositoryName

`func (o *ContainerImageLocator) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *ContainerImageLocator) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *ContainerImageLocator) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


