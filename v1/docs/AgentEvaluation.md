# AgentEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | **[]string** | Metrics to evaluate | 
**Name** | **string** | Name of the evaluation | 
**Schedule** | **string** | Evaluation schedule | 

## Methods

### NewAgentEvaluation

`func NewAgentEvaluation(metrics []string, name string, schedule string, ) *AgentEvaluation`

NewAgentEvaluation instantiates a new AgentEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentEvaluationWithDefaults

`func NewAgentEvaluationWithDefaults() *AgentEvaluation`

NewAgentEvaluationWithDefaults instantiates a new AgentEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *AgentEvaluation) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AgentEvaluation) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AgentEvaluation) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.


### GetName

`func (o *AgentEvaluation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentEvaluation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentEvaluation) SetName(v string)`

SetName sets Name field to given value.


### GetSchedule

`func (o *AgentEvaluation) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AgentEvaluation) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AgentEvaluation) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


