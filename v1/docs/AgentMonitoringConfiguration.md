# AgentMonitoringConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable monitoring for the agent | [optional] [default to false]

## Methods

### NewAgentMonitoringConfiguration

`func NewAgentMonitoringConfiguration() *AgentMonitoringConfiguration`

NewAgentMonitoringConfiguration instantiates a new AgentMonitoringConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentMonitoringConfigurationWithDefaults

`func NewAgentMonitoringConfigurationWithDefaults() *AgentMonitoringConfiguration`

NewAgentMonitoringConfigurationWithDefaults instantiates a new AgentMonitoringConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AgentMonitoringConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentMonitoringConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentMonitoringConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentMonitoringConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


