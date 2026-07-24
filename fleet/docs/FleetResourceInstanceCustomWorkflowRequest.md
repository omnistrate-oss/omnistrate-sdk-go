# FleetResourceInstanceCustomWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**RequestParams** | Pointer to **interface{}** | Input parameters for the custom workflow. System workflow operations do not require user supplied parameters. | [optional] 
**ResourceId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**WaitForCompletionSeconds** | Pointer to **int64** | Block up to N seconds (max 60) and return the completed execution inline. 0/absent &#x3D; async. | [optional] 
**WorkflowId** | **string** | ID of a Custom Workflow | 

## Methods

### NewFleetResourceInstanceCustomWorkflowRequest

`func NewFleetResourceInstanceCustomWorkflowRequest(environmentId string, instanceId string, resourceId string, serviceId string, token string, workflowId string, ) *FleetResourceInstanceCustomWorkflowRequest`

NewFleetResourceInstanceCustomWorkflowRequest instantiates a new FleetResourceInstanceCustomWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetResourceInstanceCustomWorkflowRequestWithDefaults

`func NewFleetResourceInstanceCustomWorkflowRequestWithDefaults() *FleetResourceInstanceCustomWorkflowRequest`

NewFleetResourceInstanceCustomWorkflowRequestWithDefaults instantiates a new FleetResourceInstanceCustomWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetResourceInstanceCustomWorkflowRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetResourceInstanceCustomWorkflowRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetRequestParams

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *FleetResourceInstanceCustomWorkflowRequest) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *FleetResourceInstanceCustomWorkflowRequest) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *FleetResourceInstanceCustomWorkflowRequest) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *FleetResourceInstanceCustomWorkflowRequest) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil
### GetResourceId

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FleetResourceInstanceCustomWorkflowRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetServiceId

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetResourceInstanceCustomWorkflowRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetResourceInstanceCustomWorkflowRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetWaitForCompletionSeconds

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetWaitForCompletionSeconds() int64`

GetWaitForCompletionSeconds returns the WaitForCompletionSeconds field if non-nil, zero value otherwise.

### GetWaitForCompletionSecondsOk

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetWaitForCompletionSecondsOk() (*int64, bool)`

GetWaitForCompletionSecondsOk returns a tuple with the WaitForCompletionSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCompletionSeconds

`func (o *FleetResourceInstanceCustomWorkflowRequest) SetWaitForCompletionSeconds(v int64)`

SetWaitForCompletionSeconds sets WaitForCompletionSeconds field to given value.

### HasWaitForCompletionSeconds

`func (o *FleetResourceInstanceCustomWorkflowRequest) HasWaitForCompletionSeconds() bool`

HasWaitForCompletionSeconds returns a boolean if a field has been set.

### GetWorkflowId

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *FleetResourceInstanceCustomWorkflowRequest) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *FleetResourceInstanceCustomWorkflowRequest) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


