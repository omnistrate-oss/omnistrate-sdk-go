# DescribeEventResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSource** | Pointer to **string** | The event source | [optional] 
**Id** | **string** | The ID of the event | 
**Message** | **string** | Resource Instance of message | 
**OrgId** | Pointer to **string** | The organization ID of the user that caused the event | [optional] 
**OrgName** | Pointer to **string** | The organization name of the user that caused the event | [optional] 
**ResourceInstanceId** | **string** | Instance Id of the resource instance | 
**ResourceName** | **string** | Name of the resource | 
**Time** | **string** | The event time | 
**UserId** | Pointer to **string** | The ID of the user that caused the event | [optional] 
**UserName** | Pointer to **string** | The user name of the user that caused the event | [optional] 
**WorkflowFailures** | Pointer to [**[]WorkflowFailure**](WorkflowFailure.md) | The list of workflow events that indicate failures | [optional] 

## Methods

### NewDescribeEventResult

`func NewDescribeEventResult(id string, message string, resourceInstanceId string, resourceName string, time string, ) *DescribeEventResult`

NewDescribeEventResult instantiates a new DescribeEventResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeEventResultWithDefaults

`func NewDescribeEventResultWithDefaults() *DescribeEventResult`

NewDescribeEventResultWithDefaults instantiates a new DescribeEventResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSource

`func (o *DescribeEventResult) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *DescribeEventResult) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *DescribeEventResult) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *DescribeEventResult) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetId

`func (o *DescribeEventResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeEventResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeEventResult) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *DescribeEventResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DescribeEventResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DescribeEventResult) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrgId

`func (o *DescribeEventResult) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DescribeEventResult) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DescribeEventResult) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DescribeEventResult) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *DescribeEventResult) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *DescribeEventResult) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *DescribeEventResult) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *DescribeEventResult) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetResourceInstanceId

`func (o *DescribeEventResult) GetResourceInstanceId() string`

GetResourceInstanceId returns the ResourceInstanceId field if non-nil, zero value otherwise.

### GetResourceInstanceIdOk

`func (o *DescribeEventResult) GetResourceInstanceIdOk() (*string, bool)`

GetResourceInstanceIdOk returns a tuple with the ResourceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInstanceId

`func (o *DescribeEventResult) SetResourceInstanceId(v string)`

SetResourceInstanceId sets ResourceInstanceId field to given value.


### GetResourceName

`func (o *DescribeEventResult) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *DescribeEventResult) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *DescribeEventResult) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetTime

`func (o *DescribeEventResult) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DescribeEventResult) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DescribeEventResult) SetTime(v string)`

SetTime sets Time field to given value.


### GetUserId

`func (o *DescribeEventResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DescribeEventResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DescribeEventResult) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DescribeEventResult) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *DescribeEventResult) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DescribeEventResult) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DescribeEventResult) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DescribeEventResult) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetWorkflowFailures

`func (o *DescribeEventResult) GetWorkflowFailures() []WorkflowFailure`

GetWorkflowFailures returns the WorkflowFailures field if non-nil, zero value otherwise.

### GetWorkflowFailuresOk

`func (o *DescribeEventResult) GetWorkflowFailuresOk() (*[]WorkflowFailure, bool)`

GetWorkflowFailuresOk returns a tuple with the WorkflowFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFailures

`func (o *DescribeEventResult) SetWorkflowFailures(v []WorkflowFailure)`

SetWorkflowFailures sets WorkflowFailures field to given value.

### HasWorkflowFailures

`func (o *DescribeEventResult) HasWorkflowFailures() bool`

HasWorkflowFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


