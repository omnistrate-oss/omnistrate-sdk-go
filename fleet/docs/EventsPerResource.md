# EventsPerResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | Resource ID | 
**ResourceKey** | **string** | Resource Key | 
**ResourceName** | **string** | Resource Name | 
**WorkflowSteps** | Pointer to [**[]EventsPerWorkflowStep**](EventsPerWorkflowStep.md) | Per step workflow events for the resource | [optional] 

## Methods

### NewEventsPerResource

`func NewEventsPerResource(resourceId string, resourceKey string, resourceName string, ) *EventsPerResource`

NewEventsPerResource instantiates a new EventsPerResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsPerResourceWithDefaults

`func NewEventsPerResourceWithDefaults() *EventsPerResource`

NewEventsPerResourceWithDefaults instantiates a new EventsPerResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *EventsPerResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *EventsPerResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *EventsPerResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceKey

`func (o *EventsPerResource) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *EventsPerResource) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *EventsPerResource) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetResourceName

`func (o *EventsPerResource) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *EventsPerResource) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *EventsPerResource) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetWorkflowSteps

`func (o *EventsPerResource) GetWorkflowSteps() []EventsPerWorkflowStep`

GetWorkflowSteps returns the WorkflowSteps field if non-nil, zero value otherwise.

### GetWorkflowStepsOk

`func (o *EventsPerResource) GetWorkflowStepsOk() (*[]EventsPerWorkflowStep, bool)`

GetWorkflowStepsOk returns a tuple with the WorkflowSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSteps

`func (o *EventsPerResource) SetWorkflowSteps(v []EventsPerWorkflowStep)`

SetWorkflowSteps sets WorkflowSteps field to given value.

### HasWorkflowSteps

`func (o *EventsPerResource) HasWorkflowSteps() bool`

HasWorkflowSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


