# UpdateLimitRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the limit | [optional] 
**Name** | Pointer to **string** | Name of the limit | [optional] 
**Value** | **int64** | Value of the limit being enforced | 

## Methods

### NewUpdateLimitRequest2

`func NewUpdateLimitRequest2(value int64, ) *UpdateLimitRequest2`

NewUpdateLimitRequest2 instantiates a new UpdateLimitRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLimitRequest2WithDefaults

`func NewUpdateLimitRequest2WithDefaults() *UpdateLimitRequest2`

NewUpdateLimitRequest2WithDefaults instantiates a new UpdateLimitRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateLimitRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLimitRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLimitRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLimitRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateLimitRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLimitRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLimitRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLimitRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *UpdateLimitRequest2) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateLimitRequest2) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateLimitRequest2) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


