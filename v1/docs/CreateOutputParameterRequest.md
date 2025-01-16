# CreateOutputParameterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the output variable being exported | 
**Key** | **string** | Key of the output variable being exported | 
**Name** | **string** | External name of the output variable being exported | 
**ResourceId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**Value** | Pointer to **string** | Value of the output variable being exported | [optional] 
**ValueRef** | Pointer to **string** | Reference to another variable that will be used to set the value of the output variable being exported | [optional] 
**ValueType** | Pointer to **string** | Type of the variable encoding the value | [optional] 

## Methods

### NewCreateOutputParameterRequest

`func NewCreateOutputParameterRequest(description string, key string, name string, resourceId string, serviceId string, token string, ) *CreateOutputParameterRequest`

NewCreateOutputParameterRequest instantiates a new CreateOutputParameterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOutputParameterRequestWithDefaults

`func NewCreateOutputParameterRequestWithDefaults() *CreateOutputParameterRequest`

NewCreateOutputParameterRequestWithDefaults instantiates a new CreateOutputParameterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateOutputParameterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOutputParameterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOutputParameterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetKey

`func (o *CreateOutputParameterRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateOutputParameterRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateOutputParameterRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateOutputParameterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOutputParameterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOutputParameterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetResourceId

`func (o *CreateOutputParameterRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CreateOutputParameterRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CreateOutputParameterRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetServiceId

`func (o *CreateOutputParameterRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateOutputParameterRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateOutputParameterRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *CreateOutputParameterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateOutputParameterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateOutputParameterRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetValue

`func (o *CreateOutputParameterRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateOutputParameterRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateOutputParameterRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateOutputParameterRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueRef

`func (o *CreateOutputParameterRequest) GetValueRef() string`

GetValueRef returns the ValueRef field if non-nil, zero value otherwise.

### GetValueRefOk

`func (o *CreateOutputParameterRequest) GetValueRefOk() (*string, bool)`

GetValueRefOk returns a tuple with the ValueRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRef

`func (o *CreateOutputParameterRequest) SetValueRef(v string)`

SetValueRef sets ValueRef field to given value.

### HasValueRef

`func (o *CreateOutputParameterRequest) HasValueRef() bool`

HasValueRef returns a boolean if a field has been set.

### GetValueType

`func (o *CreateOutputParameterRequest) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CreateOutputParameterRequest) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CreateOutputParameterRequest) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *CreateOutputParameterRequest) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


