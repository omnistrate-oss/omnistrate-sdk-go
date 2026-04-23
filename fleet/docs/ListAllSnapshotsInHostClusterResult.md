# ListAllSnapshotsInHostClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**Snapshots** | [**[]FleetDescribeInstanceSnapshotResult**](FleetDescribeInstanceSnapshotResult.md) | All instance snapshots in the given host cluster | 

## Methods

### NewListAllSnapshotsInHostClusterResult

`func NewListAllSnapshotsInHostClusterResult(snapshots []FleetDescribeInstanceSnapshotResult, ) *ListAllSnapshotsInHostClusterResult`

NewListAllSnapshotsInHostClusterResult instantiates a new ListAllSnapshotsInHostClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllSnapshotsInHostClusterResultWithDefaults

`func NewListAllSnapshotsInHostClusterResultWithDefaults() *ListAllSnapshotsInHostClusterResult`

NewListAllSnapshotsInHostClusterResultWithDefaults instantiates a new ListAllSnapshotsInHostClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListAllSnapshotsInHostClusterResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListAllSnapshotsInHostClusterResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListAllSnapshotsInHostClusterResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListAllSnapshotsInHostClusterResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetSnapshots

`func (o *ListAllSnapshotsInHostClusterResult) GetSnapshots() []FleetDescribeInstanceSnapshotResult`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ListAllSnapshotsInHostClusterResult) GetSnapshotsOk() (*[]FleetDescribeInstanceSnapshotResult, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ListAllSnapshotsInHostClusterResult) SetSnapshots(v []FleetDescribeInstanceSnapshotResult)`

SetSnapshots sets Snapshots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


