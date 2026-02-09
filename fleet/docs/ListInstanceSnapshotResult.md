# ListInstanceSnapshotResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | Pointer to [**[]DescribeInstanceSnapshotResult**](DescribeInstanceSnapshotResult.md) | The list of instance snapshots. | [optional] 

## Methods

### NewListInstanceSnapshotResult

`func NewListInstanceSnapshotResult() *ListInstanceSnapshotResult`

NewListInstanceSnapshotResult instantiates a new ListInstanceSnapshotResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceSnapshotResultWithDefaults

`func NewListInstanceSnapshotResultWithDefaults() *ListInstanceSnapshotResult`

NewListInstanceSnapshotResultWithDefaults instantiates a new ListInstanceSnapshotResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *ListInstanceSnapshotResult) GetSnapshots() []DescribeInstanceSnapshotResult`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ListInstanceSnapshotResult) GetSnapshotsOk() (*[]DescribeInstanceSnapshotResult, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ListInstanceSnapshotResult) SetSnapshots(v []DescribeInstanceSnapshotResult)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *ListInstanceSnapshotResult) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


