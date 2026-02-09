# FleetAuditEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**EventSource** | Pointer to **string** | The event source | [optional] 
**Id** | **string** | ID of a Event | 
**InstanceId** | **string** | ID of a Resource Instance | 
**Message** | **string** | Resource Instance of message | 
**OrgId** | Pointer to **string** | ID of an Org | [optional] 
**OrgName** | Pointer to **string** | The organization name of the user that caused the event | [optional] 
**PlanVersion** | **string** | The version of the plan | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ResourceName** | **string** | Name of the resource | 
**ServiceId** | **string** | ID of a Service | 
**ServiceName** | **string** | The service name | 
**ServicePlanName** | **string** | The name of the service plan | 
**Time** | **string** | ID of a Event | 
**UserAgent** | Pointer to **string** | The User-Agent string of the client that caused the event | [optional] 
**UserId** | Pointer to **string** | ID of a User | [optional] 
**UserName** | Pointer to **string** | The user name of the user that caused the event | [optional] 
**WorkflowFailures** | Pointer to [**[]WorkflowFailure**](WorkflowFailure.md) | The list of workflow events that indicate failures | [optional] 
**WorkflowId** | Pointer to **string** | ID of a Workflow | [optional] 

## Methods

### NewFleetAuditEvent

`func NewFleetAuditEvent(environmentId string, id string, instanceId string, message string, planVersion string, productTierId string, resourceName string, serviceId string, serviceName string, servicePlanName string, time string, ) *FleetAuditEvent`

NewFleetAuditEvent instantiates a new FleetAuditEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetAuditEventWithDefaults

`func NewFleetAuditEventWithDefaults() *FleetAuditEvent`

NewFleetAuditEventWithDefaults instantiates a new FleetAuditEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetAuditEvent) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetAuditEvent) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetAuditEvent) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEventSource

`func (o *FleetAuditEvent) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *FleetAuditEvent) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *FleetAuditEvent) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *FleetAuditEvent) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetId

`func (o *FleetAuditEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetAuditEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetAuditEvent) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *FleetAuditEvent) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetAuditEvent) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetAuditEvent) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetMessage

`func (o *FleetAuditEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FleetAuditEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FleetAuditEvent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrgId

`func (o *FleetAuditEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *FleetAuditEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *FleetAuditEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *FleetAuditEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *FleetAuditEvent) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *FleetAuditEvent) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *FleetAuditEvent) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *FleetAuditEvent) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPlanVersion

`func (o *FleetAuditEvent) GetPlanVersion() string`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *FleetAuditEvent) GetPlanVersionOk() (*string, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *FleetAuditEvent) SetPlanVersion(v string)`

SetPlanVersion sets PlanVersion field to given value.


### GetProductTierId

`func (o *FleetAuditEvent) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *FleetAuditEvent) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *FleetAuditEvent) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetResourceName

`func (o *FleetAuditEvent) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *FleetAuditEvent) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *FleetAuditEvent) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetServiceId

`func (o *FleetAuditEvent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetAuditEvent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetAuditEvent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceName

`func (o *FleetAuditEvent) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *FleetAuditEvent) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *FleetAuditEvent) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServicePlanName

`func (o *FleetAuditEvent) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *FleetAuditEvent) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *FleetAuditEvent) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.


### GetTime

`func (o *FleetAuditEvent) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *FleetAuditEvent) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *FleetAuditEvent) SetTime(v string)`

SetTime sets Time field to given value.


### GetUserAgent

`func (o *FleetAuditEvent) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *FleetAuditEvent) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *FleetAuditEvent) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *FleetAuditEvent) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserId

`func (o *FleetAuditEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FleetAuditEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FleetAuditEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FleetAuditEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *FleetAuditEvent) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FleetAuditEvent) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FleetAuditEvent) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *FleetAuditEvent) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetWorkflowFailures

`func (o *FleetAuditEvent) GetWorkflowFailures() []WorkflowFailure`

GetWorkflowFailures returns the WorkflowFailures field if non-nil, zero value otherwise.

### GetWorkflowFailuresOk

`func (o *FleetAuditEvent) GetWorkflowFailuresOk() (*[]WorkflowFailure, bool)`

GetWorkflowFailuresOk returns a tuple with the WorkflowFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFailures

`func (o *FleetAuditEvent) SetWorkflowFailures(v []WorkflowFailure)`

SetWorkflowFailures sets WorkflowFailures field to given value.

### HasWorkflowFailures

`func (o *FleetAuditEvent) HasWorkflowFailures() bool`

HasWorkflowFailures returns a boolean if a field has been set.

### GetWorkflowId

`func (o *FleetAuditEvent) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *FleetAuditEvent) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *FleetAuditEvent) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *FleetAuditEvent) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


