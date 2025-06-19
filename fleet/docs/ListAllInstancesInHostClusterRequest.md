# ListAllInstancesInHostClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusterId** | **string** | ID of a Host Cluster | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListAllInstancesInHostClusterRequest

`func NewListAllInstancesInHostClusterRequest(hostClusterId string, token string, ) *ListAllInstancesInHostClusterRequest`

NewListAllInstancesInHostClusterRequest instantiates a new ListAllInstancesInHostClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllInstancesInHostClusterRequestWithDefaults

`func NewListAllInstancesInHostClusterRequestWithDefaults() *ListAllInstancesInHostClusterRequest`

NewListAllInstancesInHostClusterRequestWithDefaults instantiates a new ListAllInstancesInHostClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusterId

`func (o *ListAllInstancesInHostClusterRequest) GetHostClusterId() string`

GetHostClusterId returns the HostClusterId field if non-nil, zero value otherwise.

### GetHostClusterIdOk

`func (o *ListAllInstancesInHostClusterRequest) GetHostClusterIdOk() (*string, bool)`

GetHostClusterIdOk returns a tuple with the HostClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterId

`func (o *ListAllInstancesInHostClusterRequest) SetHostClusterId(v string)`

SetHostClusterId sets HostClusterId field to given value.


### GetNextPageToken

`func (o *ListAllInstancesInHostClusterRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListAllInstancesInHostClusterRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListAllInstancesInHostClusterRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListAllInstancesInHostClusterRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListAllInstancesInHostClusterRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListAllInstancesInHostClusterRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListAllInstancesInHostClusterRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListAllInstancesInHostClusterRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetToken

`func (o *ListAllInstancesInHostClusterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAllInstancesInHostClusterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAllInstancesInHostClusterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


