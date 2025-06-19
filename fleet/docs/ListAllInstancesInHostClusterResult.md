# ListAllInstancesInHostClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]ResourceInstance**](ResourceInstance.md) | All resource instances in the given host cluster | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 

## Methods

### NewListAllInstancesInHostClusterResult

`func NewListAllInstancesInHostClusterResult(instances []ResourceInstance, ) *ListAllInstancesInHostClusterResult`

NewListAllInstancesInHostClusterResult instantiates a new ListAllInstancesInHostClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllInstancesInHostClusterResultWithDefaults

`func NewListAllInstancesInHostClusterResultWithDefaults() *ListAllInstancesInHostClusterResult`

NewListAllInstancesInHostClusterResultWithDefaults instantiates a new ListAllInstancesInHostClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *ListAllInstancesInHostClusterResult) GetInstances() []ResourceInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListAllInstancesInHostClusterResult) GetInstancesOk() (*[]ResourceInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListAllInstancesInHostClusterResult) SetInstances(v []ResourceInstance)`

SetInstances sets Instances field to given value.


### GetNextPageToken

`func (o *ListAllInstancesInHostClusterResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListAllInstancesInHostClusterResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListAllInstancesInHostClusterResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListAllInstancesInHostClusterResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


