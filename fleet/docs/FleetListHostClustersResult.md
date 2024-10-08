# FleetListHostClustersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusters** | [**[]FleetDescribeHostClusterResult**](FleetDescribeHostClusterResult.md) | List of host clusters | 
**Ids** | **[]string** | List of host cluster IDs | 
**NextPageToken** | Pointer to **string** | The next token to use when listing host clusters | [optional] 

## Methods

### NewFleetListHostClustersResult

`func NewFleetListHostClustersResult(hostClusters []FleetDescribeHostClusterResult, ids []string, ) *FleetListHostClustersResult`

NewFleetListHostClustersResult instantiates a new FleetListHostClustersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListHostClustersResultWithDefaults

`func NewFleetListHostClustersResultWithDefaults() *FleetListHostClustersResult`

NewFleetListHostClustersResultWithDefaults instantiates a new FleetListHostClustersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusters

`func (o *FleetListHostClustersResult) GetHostClusters() []FleetDescribeHostClusterResult`

GetHostClusters returns the HostClusters field if non-nil, zero value otherwise.

### GetHostClustersOk

`func (o *FleetListHostClustersResult) GetHostClustersOk() (*[]FleetDescribeHostClusterResult, bool)`

GetHostClustersOk returns a tuple with the HostClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusters

`func (o *FleetListHostClustersResult) SetHostClusters(v []FleetDescribeHostClusterResult)`

SetHostClusters sets HostClusters field to given value.


### GetIds

`func (o *FleetListHostClustersResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *FleetListHostClustersResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *FleetListHostClustersResult) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetNextPageToken

`func (o *FleetListHostClustersResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FleetListHostClustersResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FleetListHostClustersResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FleetListHostClustersResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


