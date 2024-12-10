# ListFleetResourceInstancesResultInternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | ID of a Service Environment | [optional] 
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**ResourceInstances** | [**[]ResourceInstance**](ResourceInstance.md) | The list of resource instances. | 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 

## Methods

### NewListFleetResourceInstancesResultInternal

`func NewListFleetResourceInstancesResultInternal(resourceInstances []ResourceInstance, ) *ListFleetResourceInstancesResultInternal`

NewListFleetResourceInstancesResultInternal instantiates a new ListFleetResourceInstancesResultInternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFleetResourceInstancesResultInternalWithDefaults

`func NewListFleetResourceInstancesResultInternalWithDefaults() *ListFleetResourceInstancesResultInternal`

NewListFleetResourceInstancesResultInternalWithDefaults instantiates a new ListFleetResourceInstancesResultInternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ListFleetResourceInstancesResultInternal) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ListFleetResourceInstancesResultInternal) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ListFleetResourceInstancesResultInternal) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ListFleetResourceInstancesResultInternal) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListFleetResourceInstancesResultInternal) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListFleetResourceInstancesResultInternal) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListFleetResourceInstancesResultInternal) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListFleetResourceInstancesResultInternal) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResourceInstances

`func (o *ListFleetResourceInstancesResultInternal) GetResourceInstances() []ResourceInstance`

GetResourceInstances returns the ResourceInstances field if non-nil, zero value otherwise.

### GetResourceInstancesOk

`func (o *ListFleetResourceInstancesResultInternal) GetResourceInstancesOk() (*[]ResourceInstance, bool)`

GetResourceInstancesOk returns a tuple with the ResourceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInstances

`func (o *ListFleetResourceInstancesResultInternal) SetResourceInstances(v []ResourceInstance)`

SetResourceInstances sets ResourceInstances field to given value.


### GetServiceId

`func (o *ListFleetResourceInstancesResultInternal) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListFleetResourceInstancesResultInternal) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListFleetResourceInstancesResultInternal) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ListFleetResourceInstancesResultInternal) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


