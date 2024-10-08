# CreateOutputParameterRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the output variable being exported | 
**Key** | **string** | Key of the output variable being exported | 
**Name** | **string** | External name of the output variable being exported | 
**ResourceId** | **string** | The ID of the resource that this output parameter belongs to | 
**Value** | Pointer to **string** | Value of the output variable being exported | [optional] 
**ValueRef** | Pointer to **string** | Reference to another variable that will be used to set the value of the output variable being exported | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOutputParameterRequestBody

`func NewCreateOutputParameterRequestBody(description string, key string, name string, resourceId string, ) *CreateOutputParameterRequestBody`

NewCreateOutputParameterRequestBody instantiates a new CreateOutputParameterRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOutputParameterRequestBodyWithDefaults

`func NewCreateOutputParameterRequestBodyWithDefaults() *CreateOutputParameterRequestBody`

NewCreateOutputParameterRequestBodyWithDefaults instantiates a new CreateOutputParameterRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateOutputParameterRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOutputParameterRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOutputParameterRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetKey

`func (o *CreateOutputParameterRequestBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateOutputParameterRequestBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateOutputParameterRequestBody) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateOutputParameterRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOutputParameterRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOutputParameterRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetResourceId

`func (o *CreateOutputParameterRequestBody) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CreateOutputParameterRequestBody) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CreateOutputParameterRequestBody) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetValue

`func (o *CreateOutputParameterRequestBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateOutputParameterRequestBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateOutputParameterRequestBody) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateOutputParameterRequestBody) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueRef

`func (o *CreateOutputParameterRequestBody) GetValueRef() string`

GetValueRef returns the ValueRef field if non-nil, zero value otherwise.

### GetValueRefOk

`func (o *CreateOutputParameterRequestBody) GetValueRefOk() (*string, bool)`

GetValueRefOk returns a tuple with the ValueRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRef

`func (o *CreateOutputParameterRequestBody) SetValueRef(v string)`

SetValueRef sets ValueRef field to given value.

### HasValueRef

`func (o *CreateOutputParameterRequestBody) HasValueRef() bool`

HasValueRef returns a boolean if a field has been set.

### GetValueType

`func (o *CreateOutputParameterRequestBody) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CreateOutputParameterRequestBody) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CreateOutputParameterRequestBody) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *CreateOutputParameterRequestBody) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


