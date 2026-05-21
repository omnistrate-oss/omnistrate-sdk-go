# FleetResourceInstanceCustomWorkflowRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestParams** | Pointer to **interface{}** | Input parameters for the custom workflow. System workflow operations do not require user supplied parameters. | [optional] 
**ResourceId** | **string** | The resource ID. | 

## Methods

### NewFleetResourceInstanceCustomWorkflowRequest2

`func NewFleetResourceInstanceCustomWorkflowRequest2(resourceId string, ) *FleetResourceInstanceCustomWorkflowRequest2`

NewFleetResourceInstanceCustomWorkflowRequest2 instantiates a new FleetResourceInstanceCustomWorkflowRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetResourceInstanceCustomWorkflowRequest2WithDefaults

`func NewFleetResourceInstanceCustomWorkflowRequest2WithDefaults() *FleetResourceInstanceCustomWorkflowRequest2`

NewFleetResourceInstanceCustomWorkflowRequest2WithDefaults instantiates a new FleetResourceInstanceCustomWorkflowRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestParams

`func (o *FleetResourceInstanceCustomWorkflowRequest2) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *FleetResourceInstanceCustomWorkflowRequest2) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *FleetResourceInstanceCustomWorkflowRequest2) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *FleetResourceInstanceCustomWorkflowRequest2) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *FleetResourceInstanceCustomWorkflowRequest2) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *FleetResourceInstanceCustomWorkflowRequest2) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil
### GetResourceId

`func (o *FleetResourceInstanceCustomWorkflowRequest2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FleetResourceInstanceCustomWorkflowRequest2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FleetResourceInstanceCustomWorkflowRequest2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


