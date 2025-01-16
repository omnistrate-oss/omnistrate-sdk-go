# UpdateOutputParameterRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the output variable being exported | [optional] 
**Name** | Pointer to **string** | External name of the output variable being exported | [optional] 
**Value** | Pointer to **string** | Value of the output variable being exported | [optional] 
**ValueRef** | Pointer to **string** | Reference to an input variable that will be used to set the value of the output variable being exported | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateOutputParameterRequest2

`func NewUpdateOutputParameterRequest2() *UpdateOutputParameterRequest2`

NewUpdateOutputParameterRequest2 instantiates a new UpdateOutputParameterRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOutputParameterRequest2WithDefaults

`func NewUpdateOutputParameterRequest2WithDefaults() *UpdateOutputParameterRequest2`

NewUpdateOutputParameterRequest2WithDefaults instantiates a new UpdateOutputParameterRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateOutputParameterRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOutputParameterRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOutputParameterRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOutputParameterRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateOutputParameterRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOutputParameterRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOutputParameterRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOutputParameterRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *UpdateOutputParameterRequest2) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateOutputParameterRequest2) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateOutputParameterRequest2) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateOutputParameterRequest2) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueRef

`func (o *UpdateOutputParameterRequest2) GetValueRef() string`

GetValueRef returns the ValueRef field if non-nil, zero value otherwise.

### GetValueRefOk

`func (o *UpdateOutputParameterRequest2) GetValueRefOk() (*string, bool)`

GetValueRefOk returns a tuple with the ValueRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRef

`func (o *UpdateOutputParameterRequest2) SetValueRef(v string)`

SetValueRef sets ValueRef field to given value.

### HasValueRef

`func (o *UpdateOutputParameterRequest2) HasValueRef() bool`

HasValueRef returns a boolean if a field has been set.

### GetValueType

`func (o *UpdateOutputParameterRequest2) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *UpdateOutputParameterRequest2) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *UpdateOutputParameterRequest2) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *UpdateOutputParameterRequest2) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


