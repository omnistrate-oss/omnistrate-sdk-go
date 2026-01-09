# OutputParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The parameter display name | [optional] 
**Key** | Pointer to **string** | The parameter key | [optional] 
**Type** | Pointer to **string** | The parameter type | [optional] 
**Value** | Pointer to **interface{}** | The parameter value | [optional] 

## Methods

### NewOutputParameter

`func NewOutputParameter() *OutputParameter`

NewOutputParameter instantiates a new OutputParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputParameterWithDefaults

`func NewOutputParameterWithDefaults() *OutputParameter`

NewOutputParameterWithDefaults instantiates a new OutputParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *OutputParameter) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OutputParameter) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OutputParameter) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OutputParameter) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetKey

`func (o *OutputParameter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OutputParameter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OutputParameter) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OutputParameter) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *OutputParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputParameter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OutputParameter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *OutputParameter) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OutputParameter) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OutputParameter) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *OutputParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *OutputParameter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *OutputParameter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


