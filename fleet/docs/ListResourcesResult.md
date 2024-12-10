# ListResourcesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**Resources** | [**[]Resource**](Resource.md) | List of resources. | 
**ServiceId** | **string** | ID of a Service | 

## Methods

### NewListResourcesResult

`func NewListResourcesResult(environmentId string, resources []Resource, serviceId string, ) *ListResourcesResult`

NewListResourcesResult instantiates a new ListResourcesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourcesResultWithDefaults

`func NewListResourcesResultWithDefaults() *ListResourcesResult`

NewListResourcesResultWithDefaults instantiates a new ListResourcesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ListResourcesResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ListResourcesResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ListResourcesResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetNextPageToken

`func (o *ListResourcesResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListResourcesResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListResourcesResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListResourcesResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResources

`func (o *ListResourcesResult) GetResources() []Resource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ListResourcesResult) GetResourcesOk() (*[]Resource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ListResourcesResult) SetResources(v []Resource)`

SetResources sets Resources field to given value.


### GetServiceId

`func (o *ListResourcesResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListResourcesResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListResourcesResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


