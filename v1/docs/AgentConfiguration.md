# AgentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeInterpreter** | [**AgentCodeInterpreterConfiguration**](AgentCodeInterpreterConfiguration.md) |  | 
**Dockerfile** | **string** | The Dockerfile path for the agent container | 
**Evaluations** | [**[]AgentEvaluation**](AgentEvaluation.md) | Evaluation configurations for the agent | 
**Memory** | [**AgentMemoryConfiguration**](AgentMemoryConfiguration.md) |  | 
**Monitoring** | [**AgentMonitoringConfiguration**](AgentMonitoringConfiguration.md) |  | 
**VectorStore** | [**AgentVectorStoreConfiguration**](AgentVectorStoreConfiguration.md) |  | 

## Methods

### NewAgentConfiguration

`func NewAgentConfiguration(codeInterpreter AgentCodeInterpreterConfiguration, dockerfile string, evaluations []AgentEvaluation, memory AgentMemoryConfiguration, monitoring AgentMonitoringConfiguration, vectorStore AgentVectorStoreConfiguration, ) *AgentConfiguration`

NewAgentConfiguration instantiates a new AgentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConfigurationWithDefaults

`func NewAgentConfigurationWithDefaults() *AgentConfiguration`

NewAgentConfigurationWithDefaults instantiates a new AgentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeInterpreter

`func (o *AgentConfiguration) GetCodeInterpreter() AgentCodeInterpreterConfiguration`

GetCodeInterpreter returns the CodeInterpreter field if non-nil, zero value otherwise.

### GetCodeInterpreterOk

`func (o *AgentConfiguration) GetCodeInterpreterOk() (*AgentCodeInterpreterConfiguration, bool)`

GetCodeInterpreterOk returns a tuple with the CodeInterpreter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeInterpreter

`func (o *AgentConfiguration) SetCodeInterpreter(v AgentCodeInterpreterConfiguration)`

SetCodeInterpreter sets CodeInterpreter field to given value.


### GetDockerfile

`func (o *AgentConfiguration) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *AgentConfiguration) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *AgentConfiguration) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.


### GetEvaluations

`func (o *AgentConfiguration) GetEvaluations() []AgentEvaluation`

GetEvaluations returns the Evaluations field if non-nil, zero value otherwise.

### GetEvaluationsOk

`func (o *AgentConfiguration) GetEvaluationsOk() (*[]AgentEvaluation, bool)`

GetEvaluationsOk returns a tuple with the Evaluations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluations

`func (o *AgentConfiguration) SetEvaluations(v []AgentEvaluation)`

SetEvaluations sets Evaluations field to given value.


### GetMemory

`func (o *AgentConfiguration) GetMemory() AgentMemoryConfiguration`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *AgentConfiguration) GetMemoryOk() (*AgentMemoryConfiguration, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *AgentConfiguration) SetMemory(v AgentMemoryConfiguration)`

SetMemory sets Memory field to given value.


### GetMonitoring

`func (o *AgentConfiguration) GetMonitoring() AgentMonitoringConfiguration`

GetMonitoring returns the Monitoring field if non-nil, zero value otherwise.

### GetMonitoringOk

`func (o *AgentConfiguration) GetMonitoringOk() (*AgentMonitoringConfiguration, bool)`

GetMonitoringOk returns a tuple with the Monitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoring

`func (o *AgentConfiguration) SetMonitoring(v AgentMonitoringConfiguration)`

SetMonitoring sets Monitoring field to given value.


### GetVectorStore

`func (o *AgentConfiguration) GetVectorStore() AgentVectorStoreConfiguration`

GetVectorStore returns the VectorStore field if non-nil, zero value otherwise.

### GetVectorStoreOk

`func (o *AgentConfiguration) GetVectorStoreOk() (*AgentVectorStoreConfiguration, bool)`

GetVectorStoreOk returns a tuple with the VectorStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorStore

`func (o *AgentConfiguration) SetVectorStore(v AgentVectorStoreConfiguration)`

SetVectorStore sets VectorStore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


