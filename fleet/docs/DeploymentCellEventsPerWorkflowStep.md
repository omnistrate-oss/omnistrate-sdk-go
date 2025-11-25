# DeploymentCellEventsPerWorkflowStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]DeploymentCellWorkflowEvent**](DeploymentCellWorkflowEvent.md) | List of events for the workflow step | 
**StepName** | **string** | The name of the workflow step | 

## Methods

### NewDeploymentCellEventsPerWorkflowStep

`func NewDeploymentCellEventsPerWorkflowStep(events []DeploymentCellWorkflowEvent, stepName string, ) *DeploymentCellEventsPerWorkflowStep`

NewDeploymentCellEventsPerWorkflowStep instantiates a new DeploymentCellEventsPerWorkflowStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellEventsPerWorkflowStepWithDefaults

`func NewDeploymentCellEventsPerWorkflowStepWithDefaults() *DeploymentCellEventsPerWorkflowStep`

NewDeploymentCellEventsPerWorkflowStepWithDefaults instantiates a new DeploymentCellEventsPerWorkflowStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *DeploymentCellEventsPerWorkflowStep) GetEvents() []DeploymentCellWorkflowEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *DeploymentCellEventsPerWorkflowStep) GetEventsOk() (*[]DeploymentCellWorkflowEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *DeploymentCellEventsPerWorkflowStep) SetEvents(v []DeploymentCellWorkflowEvent)`

SetEvents sets Events field to given value.


### GetStepName

`func (o *DeploymentCellEventsPerWorkflowStep) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *DeploymentCellEventsPerWorkflowStep) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *DeploymentCellEventsPerWorkflowStep) SetStepName(v string)`

SetStepName sets StepName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


