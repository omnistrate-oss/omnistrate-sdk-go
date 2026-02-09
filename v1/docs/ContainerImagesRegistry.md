# ContainerImagesRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ContainerImageRegistryCredentials**](ContainerImageRegistryCredentials.md) |  | [optional] 
**RegistryURL** | **string** | The URL of the container image registry (if not provided, defaults to the DockerHub) | 
**RepositoryName** | **string** | The name of the container image repository | 

## Methods

### NewContainerImagesRegistry

`func NewContainerImagesRegistry(registryURL string, repositoryName string, ) *ContainerImagesRegistry`

NewContainerImagesRegistry instantiates a new ContainerImagesRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImagesRegistryWithDefaults

`func NewContainerImagesRegistryWithDefaults() *ContainerImagesRegistry`

NewContainerImagesRegistryWithDefaults instantiates a new ContainerImagesRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ContainerImagesRegistry) GetCredentials() ContainerImageRegistryCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ContainerImagesRegistry) GetCredentialsOk() (*ContainerImageRegistryCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ContainerImagesRegistry) SetCredentials(v ContainerImageRegistryCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ContainerImagesRegistry) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetRegistryURL

`func (o *ContainerImagesRegistry) GetRegistryURL() string`

GetRegistryURL returns the RegistryURL field if non-nil, zero value otherwise.

### GetRegistryURLOk

`func (o *ContainerImagesRegistry) GetRegistryURLOk() (*string, bool)`

GetRegistryURLOk returns a tuple with the RegistryURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryURL

`func (o *ContainerImagesRegistry) SetRegistryURL(v string)`

SetRegistryURL sets RegistryURL field to given value.


### GetRepositoryName

`func (o *ContainerImagesRegistry) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *ContainerImagesRegistry) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *ContainerImagesRegistry) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


