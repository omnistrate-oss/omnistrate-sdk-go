# TerraformOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exported** | **bool** | Whether the output is exported | 
**Key** | **string** | The key of the output | 

## Methods

### NewTerraformOutput

`func NewTerraformOutput(exported bool, key string, ) *TerraformOutput`

NewTerraformOutput instantiates a new TerraformOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformOutputWithDefaults

`func NewTerraformOutputWithDefaults() *TerraformOutput`

NewTerraformOutputWithDefaults instantiates a new TerraformOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExported

`func (o *TerraformOutput) GetExported() bool`

GetExported returns the Exported field if non-nil, zero value otherwise.

### GetExportedOk

`func (o *TerraformOutput) GetExportedOk() (*bool, bool)`

GetExportedOk returns a tuple with the Exported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExported

`func (o *TerraformOutput) SetExported(v bool)`

SetExported sets Exported field to given value.


### GetKey

`func (o *TerraformOutput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TerraformOutput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TerraformOutput) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


