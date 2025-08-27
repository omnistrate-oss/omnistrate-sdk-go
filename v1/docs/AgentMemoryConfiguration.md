# AgentMemoryConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable memory store for agent context persistence | [optional] [default to false]

## Methods

### NewAgentMemoryConfiguration

`func NewAgentMemoryConfiguration() *AgentMemoryConfiguration`

NewAgentMemoryConfiguration instantiates a new AgentMemoryConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentMemoryConfigurationWithDefaults

`func NewAgentMemoryConfigurationWithDefaults() *AgentMemoryConfiguration`

NewAgentMemoryConfigurationWithDefaults instantiates a new AgentMemoryConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AgentMemoryConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentMemoryConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentMemoryConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentMemoryConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


