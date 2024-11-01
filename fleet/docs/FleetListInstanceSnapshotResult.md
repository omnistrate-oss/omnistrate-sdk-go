# FleetListInstanceSnapshotResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | Pointer to [**[]FleetDescribeInstanceSnapshotResult**](FleetDescribeInstanceSnapshotResult.md) | The list of instance snapshots. | [optional] 

## Methods

### NewFleetListInstanceSnapshotResult

`func NewFleetListInstanceSnapshotResult() *FleetListInstanceSnapshotResult`

NewFleetListInstanceSnapshotResult instantiates a new FleetListInstanceSnapshotResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListInstanceSnapshotResultWithDefaults

`func NewFleetListInstanceSnapshotResultWithDefaults() *FleetListInstanceSnapshotResult`

NewFleetListInstanceSnapshotResultWithDefaults instantiates a new FleetListInstanceSnapshotResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *FleetListInstanceSnapshotResult) GetSnapshots() []FleetDescribeInstanceSnapshotResult`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *FleetListInstanceSnapshotResult) GetSnapshotsOk() (*[]FleetDescribeInstanceSnapshotResult, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *FleetListInstanceSnapshotResult) SetSnapshots(v []FleetDescribeInstanceSnapshotResult)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *FleetListInstanceSnapshotResult) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


