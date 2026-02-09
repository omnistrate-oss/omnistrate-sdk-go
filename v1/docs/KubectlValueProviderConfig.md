# KubectlValueProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | The kubectl command to execute | 
**Context** | Pointer to **string** | Kubernetes context for the command | [optional] 
**Environment** | Pointer to **map[string]string** | Environment variables for the command | [optional] 
**Namespace** | Pointer to **string** | Kubernetes namespace for the command | [optional] 
**RetryCount** | Pointer to **int64** | Number of retry attempts | [optional] [default to 3]
**Timeout** | Pointer to **int64** | Command timeout in seconds | [optional] [default to 30]

## Methods

### NewKubectlValueProviderConfig

`func NewKubectlValueProviderConfig(command string, ) *KubectlValueProviderConfig`

NewKubectlValueProviderConfig instantiates a new KubectlValueProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubectlValueProviderConfigWithDefaults

`func NewKubectlValueProviderConfigWithDefaults() *KubectlValueProviderConfig`

NewKubectlValueProviderConfigWithDefaults instantiates a new KubectlValueProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *KubectlValueProviderConfig) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *KubectlValueProviderConfig) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *KubectlValueProviderConfig) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetContext

`func (o *KubectlValueProviderConfig) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *KubectlValueProviderConfig) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *KubectlValueProviderConfig) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *KubectlValueProviderConfig) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetEnvironment

`func (o *KubectlValueProviderConfig) GetEnvironment() map[string]string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KubectlValueProviderConfig) GetEnvironmentOk() (*map[string]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KubectlValueProviderConfig) SetEnvironment(v map[string]string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KubectlValueProviderConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNamespace

`func (o *KubectlValueProviderConfig) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubectlValueProviderConfig) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubectlValueProviderConfig) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubectlValueProviderConfig) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRetryCount

`func (o *KubectlValueProviderConfig) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *KubectlValueProviderConfig) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *KubectlValueProviderConfig) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *KubectlValueProviderConfig) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetTimeout

`func (o *KubectlValueProviderConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *KubectlValueProviderConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *KubectlValueProviderConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *KubectlValueProviderConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


