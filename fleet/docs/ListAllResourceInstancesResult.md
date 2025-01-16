# ListAllResourceInstancesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**ResourceInstances** | [**[]DescribeResourceInstanceResult**](DescribeResourceInstanceResult.md) | The list of resource instances | 

## Methods

### NewListAllResourceInstancesResult

`func NewListAllResourceInstancesResult(resourceInstances []DescribeResourceInstanceResult, ) *ListAllResourceInstancesResult`

NewListAllResourceInstancesResult instantiates a new ListAllResourceInstancesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllResourceInstancesResultWithDefaults

`func NewListAllResourceInstancesResultWithDefaults() *ListAllResourceInstancesResult`

NewListAllResourceInstancesResultWithDefaults instantiates a new ListAllResourceInstancesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListAllResourceInstancesResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListAllResourceInstancesResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListAllResourceInstancesResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListAllResourceInstancesResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetResourceInstances

`func (o *ListAllResourceInstancesResult) GetResourceInstances() []DescribeResourceInstanceResult`

GetResourceInstances returns the ResourceInstances field if non-nil, zero value otherwise.

### GetResourceInstancesOk

`func (o *ListAllResourceInstancesResult) GetResourceInstancesOk() (*[]DescribeResourceInstanceResult, bool)`

GetResourceInstancesOk returns a tuple with the ResourceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInstances

`func (o *ListAllResourceInstancesResult) SetResourceInstances(v []DescribeResourceInstanceResult)`

SetResourceInstances sets ResourceInstances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


