# UpdateOutputParameterRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the output variable being exported | [optional] 
**Name** | Pointer to **string** | External name of the output variable being exported | [optional] 
**Value** | Pointer to **string** | Value of the output variable being exported | [optional] 
**ValueRef** | Pointer to **string** | Reference to an input variable that will be used to set the value of the output variable being exported | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateOutputParameterRequestBody

`func NewUpdateOutputParameterRequestBody() *UpdateOutputParameterRequestBody`

NewUpdateOutputParameterRequestBody instantiates a new UpdateOutputParameterRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOutputParameterRequestBodyWithDefaults

`func NewUpdateOutputParameterRequestBodyWithDefaults() *UpdateOutputParameterRequestBody`

NewUpdateOutputParameterRequestBodyWithDefaults instantiates a new UpdateOutputParameterRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateOutputParameterRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOutputParameterRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOutputParameterRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOutputParameterRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateOutputParameterRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOutputParameterRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOutputParameterRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOutputParameterRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *UpdateOutputParameterRequestBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateOutputParameterRequestBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateOutputParameterRequestBody) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateOutputParameterRequestBody) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueRef

`func (o *UpdateOutputParameterRequestBody) GetValueRef() string`

GetValueRef returns the ValueRef field if non-nil, zero value otherwise.

### GetValueRefOk

`func (o *UpdateOutputParameterRequestBody) GetValueRefOk() (*string, bool)`

GetValueRefOk returns a tuple with the ValueRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRef

`func (o *UpdateOutputParameterRequestBody) SetValueRef(v string)`

SetValueRef sets ValueRef field to given value.

### HasValueRef

`func (o *UpdateOutputParameterRequestBody) HasValueRef() bool`

HasValueRef returns a boolean if a field has been set.

### GetValueType

`func (o *UpdateOutputParameterRequestBody) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *UpdateOutputParameterRequestBody) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *UpdateOutputParameterRequestBody) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *UpdateOutputParameterRequestBody) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


