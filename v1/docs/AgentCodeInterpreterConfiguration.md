# AgentCodeInterpreterConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPackages** | Pointer to **[]string** | List of allowed Python packages | [optional] 
**Enabled** | Pointer to **bool** | Enable code interpreter for the agent | [optional] [default to false]
**Provider** | Pointer to **string** | Code interpreter provider | [optional] [default to "agentcore"]
**Timeout** | Pointer to **int64** | Code execution timeout in seconds | [optional] [default to 60]

## Methods

### NewAgentCodeInterpreterConfiguration

`func NewAgentCodeInterpreterConfiguration() *AgentCodeInterpreterConfiguration`

NewAgentCodeInterpreterConfiguration instantiates a new AgentCodeInterpreterConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodeInterpreterConfigurationWithDefaults

`func NewAgentCodeInterpreterConfigurationWithDefaults() *AgentCodeInterpreterConfiguration`

NewAgentCodeInterpreterConfigurationWithDefaults instantiates a new AgentCodeInterpreterConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPackages

`func (o *AgentCodeInterpreterConfiguration) GetAllowedPackages() []string`

GetAllowedPackages returns the AllowedPackages field if non-nil, zero value otherwise.

### GetAllowedPackagesOk

`func (o *AgentCodeInterpreterConfiguration) GetAllowedPackagesOk() (*[]string, bool)`

GetAllowedPackagesOk returns a tuple with the AllowedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPackages

`func (o *AgentCodeInterpreterConfiguration) SetAllowedPackages(v []string)`

SetAllowedPackages sets AllowedPackages field to given value.

### HasAllowedPackages

`func (o *AgentCodeInterpreterConfiguration) HasAllowedPackages() bool`

HasAllowedPackages returns a boolean if a field has been set.

### GetEnabled

`func (o *AgentCodeInterpreterConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentCodeInterpreterConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentCodeInterpreterConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentCodeInterpreterConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProvider

`func (o *AgentCodeInterpreterConfiguration) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AgentCodeInterpreterConfiguration) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AgentCodeInterpreterConfiguration) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AgentCodeInterpreterConfiguration) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetTimeout

`func (o *AgentCodeInterpreterConfiguration) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AgentCodeInterpreterConfiguration) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AgentCodeInterpreterConfiguration) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *AgentCodeInterpreterConfiguration) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


