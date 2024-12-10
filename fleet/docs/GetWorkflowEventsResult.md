# GetWorkflowEventsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**Id** | **string** | ID of the ServiceWorkflow | 
**Resources** | [**[]EventsPerResource**](EventsPerResource.md) | List of resources with workflow events. | 
**ServiceId** | **string** | ID of a Service | 

## Methods

### NewGetWorkflowEventsResult

`func NewGetWorkflowEventsResult(environmentId string, id string, resources []EventsPerResource, serviceId string, ) *GetWorkflowEventsResult`

NewGetWorkflowEventsResult instantiates a new GetWorkflowEventsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflowEventsResultWithDefaults

`func NewGetWorkflowEventsResultWithDefaults() *GetWorkflowEventsResult`

NewGetWorkflowEventsResultWithDefaults instantiates a new GetWorkflowEventsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *GetWorkflowEventsResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *GetWorkflowEventsResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *GetWorkflowEventsResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *GetWorkflowEventsResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWorkflowEventsResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWorkflowEventsResult) SetId(v string)`

SetId sets Id field to given value.


### GetResources

`func (o *GetWorkflowEventsResult) GetResources() []EventsPerResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GetWorkflowEventsResult) GetResourcesOk() (*[]EventsPerResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GetWorkflowEventsResult) SetResources(v []EventsPerResource)`

SetResources sets Resources field to given value.


### GetServiceId

`func (o *GetWorkflowEventsResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetWorkflowEventsResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetWorkflowEventsResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


