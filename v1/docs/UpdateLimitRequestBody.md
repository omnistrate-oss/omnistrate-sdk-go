# UpdateLimitRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the limit | [optional] 
**Name** | Pointer to **string** | Name of the limit | [optional] 
**Value** | **int64** | Value of the limit being enforced | 

## Methods

### NewUpdateLimitRequestBody

`func NewUpdateLimitRequestBody(value int64, ) *UpdateLimitRequestBody`

NewUpdateLimitRequestBody instantiates a new UpdateLimitRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLimitRequestBodyWithDefaults

`func NewUpdateLimitRequestBodyWithDefaults() *UpdateLimitRequestBody`

NewUpdateLimitRequestBodyWithDefaults instantiates a new UpdateLimitRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateLimitRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLimitRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLimitRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLimitRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateLimitRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLimitRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLimitRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLimitRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *UpdateLimitRequestBody) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateLimitRequestBody) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateLimitRequestBody) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


