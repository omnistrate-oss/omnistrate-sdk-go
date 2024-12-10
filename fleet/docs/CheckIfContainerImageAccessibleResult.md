# CheckIfContainerImageAccessibleResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMsg** | Pointer to **string** | Error message if the image is not accessible | [optional] 
**ImageAccessible** | **bool** | Indicates if the image is accessible | 

## Methods

### NewCheckIfContainerImageAccessibleResult

`func NewCheckIfContainerImageAccessibleResult(imageAccessible bool, ) *CheckIfContainerImageAccessibleResult`

NewCheckIfContainerImageAccessibleResult instantiates a new CheckIfContainerImageAccessibleResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckIfContainerImageAccessibleResultWithDefaults

`func NewCheckIfContainerImageAccessibleResultWithDefaults() *CheckIfContainerImageAccessibleResult`

NewCheckIfContainerImageAccessibleResultWithDefaults instantiates a new CheckIfContainerImageAccessibleResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMsg

`func (o *CheckIfContainerImageAccessibleResult) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *CheckIfContainerImageAccessibleResult) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *CheckIfContainerImageAccessibleResult) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *CheckIfContainerImageAccessibleResult) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetImageAccessible

`func (o *CheckIfContainerImageAccessibleResult) GetImageAccessible() bool`

GetImageAccessible returns the ImageAccessible field if non-nil, zero value otherwise.

### GetImageAccessibleOk

`func (o *CheckIfContainerImageAccessibleResult) GetImageAccessibleOk() (*bool, bool)`

GetImageAccessibleOk returns a tuple with the ImageAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAccessible

`func (o *CheckIfContainerImageAccessibleResult) SetImageAccessible(v bool)`

SetImageAccessible sets ImageAccessible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


