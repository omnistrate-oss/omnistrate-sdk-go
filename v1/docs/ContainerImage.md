# ContainerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | **string** | The name of the container image | 
**ImageTag** | **string** | The tag of the container image | 

## Methods

### NewContainerImage

`func NewContainerImage(imageName string, imageTag string, ) *ContainerImage`

NewContainerImage instantiates a new ContainerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageWithDefaults

`func NewContainerImageWithDefaults() *ContainerImage`

NewContainerImageWithDefaults instantiates a new ContainerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *ContainerImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ContainerImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ContainerImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetImageTag

`func (o *ContainerImage) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ContainerImage) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ContainerImage) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


