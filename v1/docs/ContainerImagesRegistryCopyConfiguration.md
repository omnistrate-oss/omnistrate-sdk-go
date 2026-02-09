# ContainerImagesRegistryCopyConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]ContainerImage**](ContainerImage.md) | The list of container images to copy from source to target | [optional] 
**PullMode** | Pointer to **string** | Mode for the image pull | [optional] 
**PullSource** | Pointer to [**ContainerImagesRegistry**](ContainerImagesRegistry.md) |  | [optional] 
**PushTarget** | Pointer to [**ContainerImagesRegistry**](ContainerImagesRegistry.md) |  | [optional] 

## Methods

### NewContainerImagesRegistryCopyConfiguration

`func NewContainerImagesRegistryCopyConfiguration() *ContainerImagesRegistryCopyConfiguration`

NewContainerImagesRegistryCopyConfiguration instantiates a new ContainerImagesRegistryCopyConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImagesRegistryCopyConfigurationWithDefaults

`func NewContainerImagesRegistryCopyConfigurationWithDefaults() *ContainerImagesRegistryCopyConfiguration`

NewContainerImagesRegistryCopyConfigurationWithDefaults instantiates a new ContainerImagesRegistryCopyConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *ContainerImagesRegistryCopyConfiguration) GetImages() []ContainerImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ContainerImagesRegistryCopyConfiguration) GetImagesOk() (*[]ContainerImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ContainerImagesRegistryCopyConfiguration) SetImages(v []ContainerImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *ContainerImagesRegistryCopyConfiguration) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetPullMode

`func (o *ContainerImagesRegistryCopyConfiguration) GetPullMode() string`

GetPullMode returns the PullMode field if non-nil, zero value otherwise.

### GetPullModeOk

`func (o *ContainerImagesRegistryCopyConfiguration) GetPullModeOk() (*string, bool)`

GetPullModeOk returns a tuple with the PullMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullMode

`func (o *ContainerImagesRegistryCopyConfiguration) SetPullMode(v string)`

SetPullMode sets PullMode field to given value.

### HasPullMode

`func (o *ContainerImagesRegistryCopyConfiguration) HasPullMode() bool`

HasPullMode returns a boolean if a field has been set.

### GetPullSource

`func (o *ContainerImagesRegistryCopyConfiguration) GetPullSource() ContainerImagesRegistry`

GetPullSource returns the PullSource field if non-nil, zero value otherwise.

### GetPullSourceOk

`func (o *ContainerImagesRegistryCopyConfiguration) GetPullSourceOk() (*ContainerImagesRegistry, bool)`

GetPullSourceOk returns a tuple with the PullSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullSource

`func (o *ContainerImagesRegistryCopyConfiguration) SetPullSource(v ContainerImagesRegistry)`

SetPullSource sets PullSource field to given value.

### HasPullSource

`func (o *ContainerImagesRegistryCopyConfiguration) HasPullSource() bool`

HasPullSource returns a boolean if a field has been set.

### GetPushTarget

`func (o *ContainerImagesRegistryCopyConfiguration) GetPushTarget() ContainerImagesRegistry`

GetPushTarget returns the PushTarget field if non-nil, zero value otherwise.

### GetPushTargetOk

`func (o *ContainerImagesRegistryCopyConfiguration) GetPushTargetOk() (*ContainerImagesRegistry, bool)`

GetPushTargetOk returns a tuple with the PushTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTarget

`func (o *ContainerImagesRegistryCopyConfiguration) SetPushTarget(v ContainerImagesRegistry)`

SetPushTarget sets PushTarget field to given value.

### HasPushTarget

`func (o *ContainerImagesRegistryCopyConfiguration) HasPushTarget() bool`

HasPushTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


