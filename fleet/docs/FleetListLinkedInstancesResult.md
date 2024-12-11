# FleetListLinkedInstancesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | The service environment ID this workflow belongs to. | 
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**ResourceInstances** | [**[]ResourceInstance**](ResourceInstance.md) | The list of resource instances. | 
**ServiceId** | **string** | The service ID this workflow belongs to. | 

## Methods

### NewFleetListLinkedInstancesResult

`func NewFleetListLinkedInstancesResult(environmentId string, resourceInstances []ResourceInstance, serviceId string, ) *FleetListLinkedInstancesResult`

NewFleetListLinkedInstancesResult instantiates a new FleetListLinkedInstancesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListLinkedInstancesResultWithDefaults

`func NewFleetListLinkedInstancesResultWithDefaults() *FleetListLinkedInstancesResult`

NewFleetListLinkedInstancesResultWithDefaults instantiates a new FleetListLinkedInstancesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetListLinkedInstancesResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetListLinkedInstancesResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetListLinkedInstancesResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetNextPageToken

`func (o *FleetListLinkedInstancesResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FleetListLinkedInstancesResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FleetListLinkedInstancesResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FleetListLinkedInstancesResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResourceInstances

`func (o *FleetListLinkedInstancesResult) GetResourceInstances() []ResourceInstance`

GetResourceInstances returns the ResourceInstances field if non-nil, zero value otherwise.

### GetResourceInstancesOk

`func (o *FleetListLinkedInstancesResult) GetResourceInstancesOk() (*[]ResourceInstance, bool)`

GetResourceInstancesOk returns a tuple with the ResourceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInstances

`func (o *FleetListLinkedInstancesResult) SetResourceInstances(v []ResourceInstance)`

SetResourceInstances sets ResourceInstances field to given value.


### GetServiceId

`func (o *FleetListLinkedInstancesResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetListLinkedInstancesResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetListLinkedInstancesResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


