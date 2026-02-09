# CreateOutputParameterRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the output variable being exported | 
**GenericCommandValueProvider** | Pointer to [**GenericCommandValueProviderConfig**](GenericCommandValueProviderConfig.md) |  | [optional] 
**Key** | **string** | Key of the output variable being exported | 
**KubectlValueProvider** | Pointer to [**KubectlValueProviderConfig**](KubectlValueProviderConfig.md) |  | [optional] 
**Name** | **string** | External name of the output variable being exported | 
**ResourceId** | **string** | The ID of the resource that this output parameter belongs to | 
**Value** | Pointer to **string** | Value of the output variable being exported | [optional] 
**ValueRef** | Pointer to **string** | Reference to another variable that will be used to set the value of the output variable being exported | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOutputParameterRequest2

`func NewCreateOutputParameterRequest2(description string, key string, name string, resourceId string, ) *CreateOutputParameterRequest2`

NewCreateOutputParameterRequest2 instantiates a new CreateOutputParameterRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOutputParameterRequest2WithDefaults

`func NewCreateOutputParameterRequest2WithDefaults() *CreateOutputParameterRequest2`

NewCreateOutputParameterRequest2WithDefaults instantiates a new CreateOutputParameterRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateOutputParameterRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOutputParameterRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOutputParameterRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGenericCommandValueProvider

`func (o *CreateOutputParameterRequest2) GetGenericCommandValueProvider() GenericCommandValueProviderConfig`

GetGenericCommandValueProvider returns the GenericCommandValueProvider field if non-nil, zero value otherwise.

### GetGenericCommandValueProviderOk

`func (o *CreateOutputParameterRequest2) GetGenericCommandValueProviderOk() (*GenericCommandValueProviderConfig, bool)`

GetGenericCommandValueProviderOk returns a tuple with the GenericCommandValueProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericCommandValueProvider

`func (o *CreateOutputParameterRequest2) SetGenericCommandValueProvider(v GenericCommandValueProviderConfig)`

SetGenericCommandValueProvider sets GenericCommandValueProvider field to given value.

### HasGenericCommandValueProvider

`func (o *CreateOutputParameterRequest2) HasGenericCommandValueProvider() bool`

HasGenericCommandValueProvider returns a boolean if a field has been set.

### GetKey

`func (o *CreateOutputParameterRequest2) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateOutputParameterRequest2) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateOutputParameterRequest2) SetKey(v string)`

SetKey sets Key field to given value.


### GetKubectlValueProvider

`func (o *CreateOutputParameterRequest2) GetKubectlValueProvider() KubectlValueProviderConfig`

GetKubectlValueProvider returns the KubectlValueProvider field if non-nil, zero value otherwise.

### GetKubectlValueProviderOk

`func (o *CreateOutputParameterRequest2) GetKubectlValueProviderOk() (*KubectlValueProviderConfig, bool)`

GetKubectlValueProviderOk returns a tuple with the KubectlValueProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubectlValueProvider

`func (o *CreateOutputParameterRequest2) SetKubectlValueProvider(v KubectlValueProviderConfig)`

SetKubectlValueProvider sets KubectlValueProvider field to given value.

### HasKubectlValueProvider

`func (o *CreateOutputParameterRequest2) HasKubectlValueProvider() bool`

HasKubectlValueProvider returns a boolean if a field has been set.

### GetName

`func (o *CreateOutputParameterRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOutputParameterRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOutputParameterRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetResourceId

`func (o *CreateOutputParameterRequest2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CreateOutputParameterRequest2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CreateOutputParameterRequest2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetValue

`func (o *CreateOutputParameterRequest2) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateOutputParameterRequest2) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateOutputParameterRequest2) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateOutputParameterRequest2) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueRef

`func (o *CreateOutputParameterRequest2) GetValueRef() string`

GetValueRef returns the ValueRef field if non-nil, zero value otherwise.

### GetValueRefOk

`func (o *CreateOutputParameterRequest2) GetValueRefOk() (*string, bool)`

GetValueRefOk returns a tuple with the ValueRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRef

`func (o *CreateOutputParameterRequest2) SetValueRef(v string)`

SetValueRef sets ValueRef field to given value.

### HasValueRef

`func (o *CreateOutputParameterRequest2) HasValueRef() bool`

HasValueRef returns a boolean if a field has been set.

### GetValueType

`func (o *CreateOutputParameterRequest2) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CreateOutputParameterRequest2) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CreateOutputParameterRequest2) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *CreateOutputParameterRequest2) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


