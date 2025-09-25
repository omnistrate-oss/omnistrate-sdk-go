# UpdateOutputParameterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the output variable being exported | [optional] 
**GenericCommandValueProvider** | Pointer to [**GenericCommandValueProviderConfig**](GenericCommandValueProviderConfig.md) |  | [optional] 
**Id** | **string** | ID of an Output Parameter | 
**KubectlValueProvider** | Pointer to [**KubectlValueProviderConfig**](KubectlValueProviderConfig.md) |  | [optional] 
**Name** | Pointer to **string** | External name of the output variable being exported | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**Value** | Pointer to **string** | Value of the output variable being exported | [optional] 
**ValueRef** | Pointer to **string** | Reference to an input variable that will be used to set the value of the output variable being exported | [optional] 
**ValueType** | Pointer to **string** | Type of the variable encoding the value | [optional] 

## Methods

### NewUpdateOutputParameterRequest

`func NewUpdateOutputParameterRequest(id string, serviceId string, token string, ) *UpdateOutputParameterRequest`

NewUpdateOutputParameterRequest instantiates a new UpdateOutputParameterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOutputParameterRequestWithDefaults

`func NewUpdateOutputParameterRequestWithDefaults() *UpdateOutputParameterRequest`

NewUpdateOutputParameterRequestWithDefaults instantiates a new UpdateOutputParameterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateOutputParameterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOutputParameterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOutputParameterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOutputParameterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenericCommandValueProvider

`func (o *UpdateOutputParameterRequest) GetGenericCommandValueProvider() GenericCommandValueProviderConfig`

GetGenericCommandValueProvider returns the GenericCommandValueProvider field if non-nil, zero value otherwise.

### GetGenericCommandValueProviderOk

`func (o *UpdateOutputParameterRequest) GetGenericCommandValueProviderOk() (*GenericCommandValueProviderConfig, bool)`

GetGenericCommandValueProviderOk returns a tuple with the GenericCommandValueProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericCommandValueProvider

`func (o *UpdateOutputParameterRequest) SetGenericCommandValueProvider(v GenericCommandValueProviderConfig)`

SetGenericCommandValueProvider sets GenericCommandValueProvider field to given value.

### HasGenericCommandValueProvider

`func (o *UpdateOutputParameterRequest) HasGenericCommandValueProvider() bool`

HasGenericCommandValueProvider returns a boolean if a field has been set.

### GetId

`func (o *UpdateOutputParameterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOutputParameterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOutputParameterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetKubectlValueProvider

`func (o *UpdateOutputParameterRequest) GetKubectlValueProvider() KubectlValueProviderConfig`

GetKubectlValueProvider returns the KubectlValueProvider field if non-nil, zero value otherwise.

### GetKubectlValueProviderOk

`func (o *UpdateOutputParameterRequest) GetKubectlValueProviderOk() (*KubectlValueProviderConfig, bool)`

GetKubectlValueProviderOk returns a tuple with the KubectlValueProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubectlValueProvider

`func (o *UpdateOutputParameterRequest) SetKubectlValueProvider(v KubectlValueProviderConfig)`

SetKubectlValueProvider sets KubectlValueProvider field to given value.

### HasKubectlValueProvider

`func (o *UpdateOutputParameterRequest) HasKubectlValueProvider() bool`

HasKubectlValueProvider returns a boolean if a field has been set.

### GetName

`func (o *UpdateOutputParameterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOutputParameterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOutputParameterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOutputParameterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceId

`func (o *UpdateOutputParameterRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateOutputParameterRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateOutputParameterRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *UpdateOutputParameterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateOutputParameterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateOutputParameterRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetValue

`func (o *UpdateOutputParameterRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateOutputParameterRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateOutputParameterRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateOutputParameterRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueRef

`func (o *UpdateOutputParameterRequest) GetValueRef() string`

GetValueRef returns the ValueRef field if non-nil, zero value otherwise.

### GetValueRefOk

`func (o *UpdateOutputParameterRequest) GetValueRefOk() (*string, bool)`

GetValueRefOk returns a tuple with the ValueRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRef

`func (o *UpdateOutputParameterRequest) SetValueRef(v string)`

SetValueRef sets ValueRef field to given value.

### HasValueRef

`func (o *UpdateOutputParameterRequest) HasValueRef() bool`

HasValueRef returns a boolean if a field has been set.

### GetValueType

`func (o *UpdateOutputParameterRequest) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *UpdateOutputParameterRequest) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *UpdateOutputParameterRequest) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *UpdateOutputParameterRequest) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


