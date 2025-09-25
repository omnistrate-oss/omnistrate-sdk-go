# GenericCommandValueProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | Arguments for the command | [optional] 
**Command** | **string** | The command to execute | 
**Environment** | Pointer to **map[string]string** | Environment variables for the command | [optional] 
**RetryCount** | Pointer to **int64** | Number of retry attempts | [optional] [default to 3]
**Shell** | Pointer to **string** | Shell to use for command execution | [optional] [default to "/bin/sh"]
**Timeout** | Pointer to **int64** | Command timeout in seconds | [optional] [default to 30]
**WorkingDir** | Pointer to **string** | Working directory for command execution | [optional] 

## Methods

### NewGenericCommandValueProviderConfig

`func NewGenericCommandValueProviderConfig(command string, ) *GenericCommandValueProviderConfig`

NewGenericCommandValueProviderConfig instantiates a new GenericCommandValueProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericCommandValueProviderConfigWithDefaults

`func NewGenericCommandValueProviderConfigWithDefaults() *GenericCommandValueProviderConfig`

NewGenericCommandValueProviderConfigWithDefaults instantiates a new GenericCommandValueProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *GenericCommandValueProviderConfig) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *GenericCommandValueProviderConfig) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *GenericCommandValueProviderConfig) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *GenericCommandValueProviderConfig) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCommand

`func (o *GenericCommandValueProviderConfig) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *GenericCommandValueProviderConfig) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *GenericCommandValueProviderConfig) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetEnvironment

`func (o *GenericCommandValueProviderConfig) GetEnvironment() map[string]string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GenericCommandValueProviderConfig) GetEnvironmentOk() (*map[string]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GenericCommandValueProviderConfig) SetEnvironment(v map[string]string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GenericCommandValueProviderConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRetryCount

`func (o *GenericCommandValueProviderConfig) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *GenericCommandValueProviderConfig) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *GenericCommandValueProviderConfig) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *GenericCommandValueProviderConfig) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetShell

`func (o *GenericCommandValueProviderConfig) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *GenericCommandValueProviderConfig) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *GenericCommandValueProviderConfig) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *GenericCommandValueProviderConfig) HasShell() bool`

HasShell returns a boolean if a field has been set.

### GetTimeout

`func (o *GenericCommandValueProviderConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GenericCommandValueProviderConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GenericCommandValueProviderConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GenericCommandValueProviderConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetWorkingDir

`func (o *GenericCommandValueProviderConfig) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *GenericCommandValueProviderConfig) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *GenericCommandValueProviderConfig) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.

### HasWorkingDir

`func (o *GenericCommandValueProviderConfig) HasWorkingDir() bool`

HasWorkingDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


