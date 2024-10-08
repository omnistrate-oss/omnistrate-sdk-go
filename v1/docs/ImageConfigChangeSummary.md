# ImageConfigChangeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeState** | Pointer to **string** | The pending change state for the infra configuration | [optional] 
**ImageConfigId** | Pointer to **string** | The ID of the image configuration that this resource refers to | [optional] 
**ImageName** | Pointer to **string** | Name of the container image | [optional] 

## Methods

### NewImageConfigChangeSummary

`func NewImageConfigChangeSummary() *ImageConfigChangeSummary`

NewImageConfigChangeSummary instantiates a new ImageConfigChangeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageConfigChangeSummaryWithDefaults

`func NewImageConfigChangeSummaryWithDefaults() *ImageConfigChangeSummary`

NewImageConfigChangeSummaryWithDefaults instantiates a new ImageConfigChangeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeState

`func (o *ImageConfigChangeSummary) GetChangeState() string`

GetChangeState returns the ChangeState field if non-nil, zero value otherwise.

### GetChangeStateOk

`func (o *ImageConfigChangeSummary) GetChangeStateOk() (*string, bool)`

GetChangeStateOk returns a tuple with the ChangeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeState

`func (o *ImageConfigChangeSummary) SetChangeState(v string)`

SetChangeState sets ChangeState field to given value.

### HasChangeState

`func (o *ImageConfigChangeSummary) HasChangeState() bool`

HasChangeState returns a boolean if a field has been set.

### GetImageConfigId

`func (o *ImageConfigChangeSummary) GetImageConfigId() string`

GetImageConfigId returns the ImageConfigId field if non-nil, zero value otherwise.

### GetImageConfigIdOk

`func (o *ImageConfigChangeSummary) GetImageConfigIdOk() (*string, bool)`

GetImageConfigIdOk returns a tuple with the ImageConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfigId

`func (o *ImageConfigChangeSummary) SetImageConfigId(v string)`

SetImageConfigId sets ImageConfigId field to given value.

### HasImageConfigId

`func (o *ImageConfigChangeSummary) HasImageConfigId() bool`

HasImageConfigId returns a boolean if a field has been set.

### GetImageName

`func (o *ImageConfigChangeSummary) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ImageConfigChangeSummary) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ImageConfigChangeSummary) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ImageConfigChangeSummary) HasImageName() bool`

HasImageName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


