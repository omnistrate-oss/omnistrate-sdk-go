# DescribeOutputParameterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the output variable being exported | 
**Id** | **string** | ID of an Output Parameter | 
**Key** | **string** | Key of the output variable being exported | 
**Name** | **string** | External name of the output variable being exported | 
**ResourceId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Value** | Pointer to **string** | Value of the output variable being exported | [optional] 
**ValueRef** | Pointer to **string** | Reference to an input variable that will be used to set the value of the output variable being exported | [optional] 
**ValueType** | Pointer to **string** | Type of the variable encoding the value | [optional] 

## Methods

### NewDescribeOutputParameterResult

`func NewDescribeOutputParameterResult(description string, id string, key string, name string, resourceId string, serviceId string, ) *DescribeOutputParameterResult`

NewDescribeOutputParameterResult instantiates a new DescribeOutputParameterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeOutputParameterResultWithDefaults

`func NewDescribeOutputParameterResultWithDefaults() *DescribeOutputParameterResult`

NewDescribeOutputParameterResultWithDefaults instantiates a new DescribeOutputParameterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DescribeOutputParameterResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeOutputParameterResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeOutputParameterResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeOutputParameterResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeOutputParameterResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeOutputParameterResult) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *DescribeOutputParameterResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DescribeOutputParameterResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DescribeOutputParameterResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *DescribeOutputParameterResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeOutputParameterResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeOutputParameterResult) SetName(v string)`

SetName sets Name field to given value.


### GetResourceId

`func (o *DescribeOutputParameterResult) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DescribeOutputParameterResult) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DescribeOutputParameterResult) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetServiceId

`func (o *DescribeOutputParameterResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeOutputParameterResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeOutputParameterResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetValue

`func (o *DescribeOutputParameterResult) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DescribeOutputParameterResult) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DescribeOutputParameterResult) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DescribeOutputParameterResult) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueRef

`func (o *DescribeOutputParameterResult) GetValueRef() string`

GetValueRef returns the ValueRef field if non-nil, zero value otherwise.

### GetValueRefOk

`func (o *DescribeOutputParameterResult) GetValueRefOk() (*string, bool)`

GetValueRefOk returns a tuple with the ValueRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRef

`func (o *DescribeOutputParameterResult) SetValueRef(v string)`

SetValueRef sets ValueRef field to given value.

### HasValueRef

`func (o *DescribeOutputParameterResult) HasValueRef() bool`

HasValueRef returns a boolean if a field has been set.

### GetValueType

`func (o *DescribeOutputParameterResult) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *DescribeOutputParameterResult) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *DescribeOutputParameterResult) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *DescribeOutputParameterResult) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


