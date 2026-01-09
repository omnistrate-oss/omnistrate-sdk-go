# DescribeAuditEventResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSource** | Pointer to **string** | The event source | [optional] 
**Id** | **string** | ID of a Event | 
**Message** | **string** | Resource Instance of message | 
**OrgId** | Pointer to **string** | ID of an Org | [optional] 
**OrgName** | Pointer to **string** | The organization name of the user that caused the event | [optional] 
**ResourceInstanceId** | **string** | Instance Id of the resource instance | 
**ResourceName** | **string** | Name of the resource | 
**SubscriptionId** | **string** | The subscription ID | 
**Time** | **string** | The event time | 
**UserAgent** | Pointer to **string** | The User-Agent string of the client that caused the event | [optional] 
**UserId** | Pointer to **string** | ID of a User | [optional] 
**UserName** | Pointer to **string** | The user name of the user that caused the event | [optional] 
**WorkflowFailures** | Pointer to [**[]WorkflowFailure**](WorkflowFailure.md) | The list of workflow events that indicate failures | [optional] 

## Methods

### NewDescribeAuditEventResult

`func NewDescribeAuditEventResult(id string, message string, resourceInstanceId string, resourceName string, subscriptionId string, time string, ) *DescribeAuditEventResult`

NewDescribeAuditEventResult instantiates a new DescribeAuditEventResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAuditEventResultWithDefaults

`func NewDescribeAuditEventResultWithDefaults() *DescribeAuditEventResult`

NewDescribeAuditEventResultWithDefaults instantiates a new DescribeAuditEventResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSource

`func (o *DescribeAuditEventResult) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *DescribeAuditEventResult) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *DescribeAuditEventResult) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *DescribeAuditEventResult) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetId

`func (o *DescribeAuditEventResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeAuditEventResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeAuditEventResult) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *DescribeAuditEventResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DescribeAuditEventResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DescribeAuditEventResult) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrgId

`func (o *DescribeAuditEventResult) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DescribeAuditEventResult) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DescribeAuditEventResult) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DescribeAuditEventResult) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *DescribeAuditEventResult) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *DescribeAuditEventResult) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *DescribeAuditEventResult) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *DescribeAuditEventResult) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetResourceInstanceId

`func (o *DescribeAuditEventResult) GetResourceInstanceId() string`

GetResourceInstanceId returns the ResourceInstanceId field if non-nil, zero value otherwise.

### GetResourceInstanceIdOk

`func (o *DescribeAuditEventResult) GetResourceInstanceIdOk() (*string, bool)`

GetResourceInstanceIdOk returns a tuple with the ResourceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInstanceId

`func (o *DescribeAuditEventResult) SetResourceInstanceId(v string)`

SetResourceInstanceId sets ResourceInstanceId field to given value.


### GetResourceName

`func (o *DescribeAuditEventResult) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *DescribeAuditEventResult) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *DescribeAuditEventResult) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetSubscriptionId

`func (o *DescribeAuditEventResult) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DescribeAuditEventResult) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DescribeAuditEventResult) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTime

`func (o *DescribeAuditEventResult) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DescribeAuditEventResult) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DescribeAuditEventResult) SetTime(v string)`

SetTime sets Time field to given value.


### GetUserAgent

`func (o *DescribeAuditEventResult) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *DescribeAuditEventResult) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *DescribeAuditEventResult) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *DescribeAuditEventResult) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserId

`func (o *DescribeAuditEventResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DescribeAuditEventResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DescribeAuditEventResult) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DescribeAuditEventResult) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *DescribeAuditEventResult) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DescribeAuditEventResult) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DescribeAuditEventResult) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DescribeAuditEventResult) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetWorkflowFailures

`func (o *DescribeAuditEventResult) GetWorkflowFailures() []WorkflowFailure`

GetWorkflowFailures returns the WorkflowFailures field if non-nil, zero value otherwise.

### GetWorkflowFailuresOk

`func (o *DescribeAuditEventResult) GetWorkflowFailuresOk() (*[]WorkflowFailure, bool)`

GetWorkflowFailuresOk returns a tuple with the WorkflowFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFailures

`func (o *DescribeAuditEventResult) SetWorkflowFailures(v []WorkflowFailure)`

SetWorkflowFailures sets WorkflowFailures field to given value.

### HasWorkflowFailures

`func (o *DescribeAuditEventResult) HasWorkflowFailures() bool`

HasWorkflowFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


