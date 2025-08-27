# ContainerImageConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**ContainerImageLocator**](ContainerImageLocator.md) |  | [optional] 
**Target** | Pointer to [**ContainerImageLocator**](ContainerImageLocator.md) |  | [optional] 

## Methods

### NewContainerImageConfiguration

`func NewContainerImageConfiguration() *ContainerImageConfiguration`

NewContainerImageConfiguration instantiates a new ContainerImageConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageConfigurationWithDefaults

`func NewContainerImageConfigurationWithDefaults() *ContainerImageConfiguration`

NewContainerImageConfigurationWithDefaults instantiates a new ContainerImageConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ContainerImageConfiguration) GetSource() ContainerImageLocator`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContainerImageConfiguration) GetSourceOk() (*ContainerImageLocator, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContainerImageConfiguration) SetSource(v ContainerImageLocator)`

SetSource sets Source field to given value.

### HasSource

`func (o *ContainerImageConfiguration) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *ContainerImageConfiguration) GetTarget() ContainerImageLocator`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ContainerImageConfiguration) GetTargetOk() (*ContainerImageLocator, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ContainerImageConfiguration) SetTarget(v ContainerImageLocator)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ContainerImageConfiguration) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


