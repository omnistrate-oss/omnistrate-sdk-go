# AgentVectorStoreConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable vector store for the agent | [optional] [default to false]

## Methods

### NewAgentVectorStoreConfiguration

`func NewAgentVectorStoreConfiguration() *AgentVectorStoreConfiguration`

NewAgentVectorStoreConfiguration instantiates a new AgentVectorStoreConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentVectorStoreConfigurationWithDefaults

`func NewAgentVectorStoreConfigurationWithDefaults() *AgentVectorStoreConfiguration`

NewAgentVectorStoreConfigurationWithDefaults instantiates a new AgentVectorStoreConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AgentVectorStoreConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentVectorStoreConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentVectorStoreConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentVectorStoreConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


