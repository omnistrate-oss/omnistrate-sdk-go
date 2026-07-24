# FleetInvokeResourceInstanceWorkflowVerbRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestParams** | Pointer to **interface{}** | Input parameters for the custom workflow. System workflow operations do not require user supplied parameters. | [optional] 
**ResourceId** | **string** | The resource ID. | 
**WaitForCompletionSeconds** | Pointer to **int64** | Block up to N seconds (max 60) and return the completed execution inline. 0/absent &#x3D; async. | [optional] 

## Methods

### NewFleetInvokeResourceInstanceWorkflowVerbRequest2

`func NewFleetInvokeResourceInstanceWorkflowVerbRequest2(resourceId string, ) *FleetInvokeResourceInstanceWorkflowVerbRequest2`

NewFleetInvokeResourceInstanceWorkflowVerbRequest2 instantiates a new FleetInvokeResourceInstanceWorkflowVerbRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetInvokeResourceInstanceWorkflowVerbRequest2WithDefaults

`func NewFleetInvokeResourceInstanceWorkflowVerbRequest2WithDefaults() *FleetInvokeResourceInstanceWorkflowVerbRequest2`

NewFleetInvokeResourceInstanceWorkflowVerbRequest2WithDefaults instantiates a new FleetInvokeResourceInstanceWorkflowVerbRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestParams

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil
### GetResourceId

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetWaitForCompletionSeconds

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) GetWaitForCompletionSeconds() int64`

GetWaitForCompletionSeconds returns the WaitForCompletionSeconds field if non-nil, zero value otherwise.

### GetWaitForCompletionSecondsOk

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) GetWaitForCompletionSecondsOk() (*int64, bool)`

GetWaitForCompletionSecondsOk returns a tuple with the WaitForCompletionSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCompletionSeconds

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) SetWaitForCompletionSeconds(v int64)`

SetWaitForCompletionSeconds sets WaitForCompletionSeconds field to given value.

### HasWaitForCompletionSeconds

`func (o *FleetInvokeResourceInstanceWorkflowVerbRequest2) HasWaitForCompletionSeconds() bool`

HasWaitForCompletionSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


