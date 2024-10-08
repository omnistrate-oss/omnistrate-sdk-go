# FleetDescribeEventResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | The service environment ID this workflow belongs to. | 
**EventSource** | Pointer to **string** | The event source | [optional] 
**Id** | **string** | The ID of the event | 
**InstanceId** | **string** | The resource instance ID. | 
**Message** | **string** | Resource Instance of message | 
**OrgId** | Pointer to **string** | The organization ID of the user that caused the event | [optional] 
**OrgName** | Pointer to **string** | The organization name of the user that caused the event | [optional] 
**ResourceName** | **string** | Name of the resource | 
**ServiceId** | **string** | The service ID this workflow belongs to. | 
**Time** | **string** | The event time | 
**UserId** | Pointer to **string** | The ID of the user that caused the event | [optional] 
**UserName** | Pointer to **string** | The user name of the user that caused the event | [optional] 
**WorkflowFailures** | Pointer to [**[]WorkflowFailure**](WorkflowFailure.md) | The list of workflow events that indicate failures | [optional] 
**WorkflowId** | Pointer to **string** | The workflow ID | [optional] 

## Methods

### NewFleetDescribeEventResult

`func NewFleetDescribeEventResult(environmentId string, id string, instanceId string, message string, resourceName string, serviceId string, time string, ) *FleetDescribeEventResult`

NewFleetDescribeEventResult instantiates a new FleetDescribeEventResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDescribeEventResultWithDefaults

`func NewFleetDescribeEventResultWithDefaults() *FleetDescribeEventResult`

NewFleetDescribeEventResultWithDefaults instantiates a new FleetDescribeEventResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetDescribeEventResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetDescribeEventResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetDescribeEventResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEventSource

`func (o *FleetDescribeEventResult) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *FleetDescribeEventResult) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *FleetDescribeEventResult) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *FleetDescribeEventResult) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetId

`func (o *FleetDescribeEventResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetDescribeEventResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetDescribeEventResult) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *FleetDescribeEventResult) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetDescribeEventResult) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetDescribeEventResult) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetMessage

`func (o *FleetDescribeEventResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FleetDescribeEventResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FleetDescribeEventResult) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrgId

`func (o *FleetDescribeEventResult) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *FleetDescribeEventResult) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *FleetDescribeEventResult) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *FleetDescribeEventResult) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *FleetDescribeEventResult) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *FleetDescribeEventResult) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *FleetDescribeEventResult) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *FleetDescribeEventResult) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetResourceName

`func (o *FleetDescribeEventResult) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *FleetDescribeEventResult) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *FleetDescribeEventResult) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetServiceId

`func (o *FleetDescribeEventResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetDescribeEventResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetDescribeEventResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTime

`func (o *FleetDescribeEventResult) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *FleetDescribeEventResult) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *FleetDescribeEventResult) SetTime(v string)`

SetTime sets Time field to given value.


### GetUserId

`func (o *FleetDescribeEventResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FleetDescribeEventResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FleetDescribeEventResult) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FleetDescribeEventResult) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *FleetDescribeEventResult) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FleetDescribeEventResult) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FleetDescribeEventResult) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *FleetDescribeEventResult) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetWorkflowFailures

`func (o *FleetDescribeEventResult) GetWorkflowFailures() []WorkflowFailure`

GetWorkflowFailures returns the WorkflowFailures field if non-nil, zero value otherwise.

### GetWorkflowFailuresOk

`func (o *FleetDescribeEventResult) GetWorkflowFailuresOk() (*[]WorkflowFailure, bool)`

GetWorkflowFailuresOk returns a tuple with the WorkflowFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFailures

`func (o *FleetDescribeEventResult) SetWorkflowFailures(v []WorkflowFailure)`

SetWorkflowFailures sets WorkflowFailures field to given value.

### HasWorkflowFailures

`func (o *FleetDescribeEventResult) HasWorkflowFailures() bool`

HasWorkflowFailures returns a boolean if a field has been set.

### GetWorkflowId

`func (o *FleetDescribeEventResult) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *FleetDescribeEventResult) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *FleetDescribeEventResult) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *FleetDescribeEventResult) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


