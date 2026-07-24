# FleetInvokeResourceInstanceWorkflowVerbRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**RequestParams** | Pointer to **interface{}** | Input parameters for the custom workflow. System workflow operations do not require user supplied parameters. | [optional] 
**ResourceId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**Verb** | **string** | The provider-defined verb to invoke. The server resolves the verb to a custom workflow ID against the instance&#39;s plan version. | 
**WaitForCompletionSeconds** | Pointer to **int64** | Block up to N seconds (max 60) and return the completed execution inline. 0/absent &#x3D; async. | [optional] 

## Methods

### NewFleetInvokeResourceInstanceWorkflowVerbRequest

`func NewFleetInvokeResourceInstanceWorkflowVerbRequest(environmentId string, instanceId string, resourceId string, serviceId string, token string, verb string, ) *FleetInvokeResourceInstanceWorkflowVerbRequest`

NewFleetInvokeResourceInstanceWorkflowVerbRequest instantiates a new FleetInvokeResourceInstanceWorkflowVerbRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetInvokeResourceInstanceWorkflowVerbRequestWithDefaults

`func NewFleetInvokeResourceInstanceWorkflowVerbRequestWithDefaults() *FleetInvokeResourceInstanceWorkflowVerbRequest`

NewFleetInvokeResourceInstanceWorkflowVerbRequestWithDefaults instantiates a new FleetInvokeResourceInstanceWorkflowVerbRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetRequestParams

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil
### GetResourceId

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetServiceId

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVerb

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetWaitForCompletionSeconds

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetWaitForCompletionSeconds() int64`

GetWaitForCompletionSeconds returns the WaitForCompletionSeconds field if non-nil, zero value otherwise.

### GetWaitForCompletionSecondsOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) GetWaitForCompletionSecondsOk() (*int64, bool)`

GetWaitForCompletionSecondsOk returns a tuple with the WaitForCompletionSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCompletionSeconds

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) SetWaitForCompletionSeconds(v int64)`

SetWaitForCompletionSeconds sets WaitForCompletionSeconds field to given value.

### HasWaitForCompletionSeconds

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest) HasWaitForCompletionSeconds() bool`

HasWaitForCompletionSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


