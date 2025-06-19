# TaintConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | The effect of the taint | 
**Key** | **string** | The key of the taint | 
**Value** | **string** | The value of the taint | 

## Methods

### NewTaintConfiguration

`func NewTaintConfiguration(effect string, key string, value string, ) *TaintConfiguration`

NewTaintConfiguration instantiates a new TaintConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaintConfigurationWithDefaults

`func NewTaintConfigurationWithDefaults() *TaintConfiguration`

NewTaintConfigurationWithDefaults instantiates a new TaintConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *TaintConfiguration) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *TaintConfiguration) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *TaintConfiguration) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetKey

`func (o *TaintConfiguration) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TaintConfiguration) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TaintConfiguration) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *TaintConfiguration) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TaintConfiguration) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TaintConfiguration) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


