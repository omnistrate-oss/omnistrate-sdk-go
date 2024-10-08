# SetEnvironmentVariablesRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariable**](EnvironmentVariable.md) | The environment variables that this resource requires | [optional] 

## Methods

### NewSetEnvironmentVariablesRequestBody

`func NewSetEnvironmentVariablesRequestBody() *SetEnvironmentVariablesRequestBody`

NewSetEnvironmentVariablesRequestBody instantiates a new SetEnvironmentVariablesRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetEnvironmentVariablesRequestBodyWithDefaults

`func NewSetEnvironmentVariablesRequestBodyWithDefaults() *SetEnvironmentVariablesRequestBody`

NewSetEnvironmentVariablesRequestBodyWithDefaults instantiates a new SetEnvironmentVariablesRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentVariables

`func (o *SetEnvironmentVariablesRequestBody) GetEnvironmentVariables() []EnvironmentVariable`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *SetEnvironmentVariablesRequestBody) GetEnvironmentVariablesOk() (*[]EnvironmentVariable, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *SetEnvironmentVariablesRequestBody) SetEnvironmentVariables(v []EnvironmentVariable)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *SetEnvironmentVariablesRequestBody) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


