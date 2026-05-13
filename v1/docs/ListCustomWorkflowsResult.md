# ListCustomWorkflowsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomWorkflows** | [**[]DescribeCustomWorkflowResult**](DescribeCustomWorkflowResult.md) | The list of custom workflows associated with the service plan | 

## Methods

### NewListCustomWorkflowsResult

`func NewListCustomWorkflowsResult(customWorkflows []DescribeCustomWorkflowResult, ) *ListCustomWorkflowsResult`

NewListCustomWorkflowsResult instantiates a new ListCustomWorkflowsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomWorkflowsResultWithDefaults

`func NewListCustomWorkflowsResultWithDefaults() *ListCustomWorkflowsResult`

NewListCustomWorkflowsResultWithDefaults instantiates a new ListCustomWorkflowsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomWorkflows

`func (o *ListCustomWorkflowsResult) GetCustomWorkflows() []DescribeCustomWorkflowResult`

GetCustomWorkflows returns the CustomWorkflows field if non-nil, zero value otherwise.

### GetCustomWorkflowsOk

`func (o *ListCustomWorkflowsResult) GetCustomWorkflowsOk() (*[]DescribeCustomWorkflowResult, bool)`

GetCustomWorkflowsOk returns a tuple with the CustomWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomWorkflows

`func (o *ListCustomWorkflowsResult) SetCustomWorkflows(v []DescribeCustomWorkflowResult)`

SetCustomWorkflows sets CustomWorkflows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


