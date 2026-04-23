# ListAllSnapshotsInHostClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusterId** | **string** | ID of a Host Cluster | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListAllSnapshotsInHostClusterRequest

`func NewListAllSnapshotsInHostClusterRequest(hostClusterId string, token string, ) *ListAllSnapshotsInHostClusterRequest`

NewListAllSnapshotsInHostClusterRequest instantiates a new ListAllSnapshotsInHostClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllSnapshotsInHostClusterRequestWithDefaults

`func NewListAllSnapshotsInHostClusterRequestWithDefaults() *ListAllSnapshotsInHostClusterRequest`

NewListAllSnapshotsInHostClusterRequestWithDefaults instantiates a new ListAllSnapshotsInHostClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusterId

`func (o *ListAllSnapshotsInHostClusterRequest) GetHostClusterId() string`

GetHostClusterId returns the HostClusterId field if non-nil, zero value otherwise.

### GetHostClusterIdOk

`func (o *ListAllSnapshotsInHostClusterRequest) GetHostClusterIdOk() (*string, bool)`

GetHostClusterIdOk returns a tuple with the HostClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterId

`func (o *ListAllSnapshotsInHostClusterRequest) SetHostClusterId(v string)`

SetHostClusterId sets HostClusterId field to given value.


### GetNextPageToken

`func (o *ListAllSnapshotsInHostClusterRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListAllSnapshotsInHostClusterRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListAllSnapshotsInHostClusterRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListAllSnapshotsInHostClusterRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListAllSnapshotsInHostClusterRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListAllSnapshotsInHostClusterRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListAllSnapshotsInHostClusterRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListAllSnapshotsInHostClusterRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetToken

`func (o *ListAllSnapshotsInHostClusterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAllSnapshotsInHostClusterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAllSnapshotsInHostClusterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


