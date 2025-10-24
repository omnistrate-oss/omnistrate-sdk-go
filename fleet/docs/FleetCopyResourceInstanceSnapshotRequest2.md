# FleetCopyResourceInstanceSnapshotRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceSnapshotId** | Pointer to **string** | The source snapshot ID | [optional] 
**TargetRegion** | **string** | The target region to copy the snapshot to | 

## Methods

### NewFleetCopyResourceInstanceSnapshotRequest2

`func NewFleetCopyResourceInstanceSnapshotRequest2(targetRegion string, ) *FleetCopyResourceInstanceSnapshotRequest2`

NewFleetCopyResourceInstanceSnapshotRequest2 instantiates a new FleetCopyResourceInstanceSnapshotRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCopyResourceInstanceSnapshotRequest2WithDefaults

`func NewFleetCopyResourceInstanceSnapshotRequest2WithDefaults() *FleetCopyResourceInstanceSnapshotRequest2`

NewFleetCopyResourceInstanceSnapshotRequest2WithDefaults instantiates a new FleetCopyResourceInstanceSnapshotRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceSnapshotId

`func (o *FleetCopyResourceInstanceSnapshotRequest2) GetSourceSnapshotId() string`

GetSourceSnapshotId returns the SourceSnapshotId field if non-nil, zero value otherwise.

### GetSourceSnapshotIdOk

`func (o *FleetCopyResourceInstanceSnapshotRequest2) GetSourceSnapshotIdOk() (*string, bool)`

GetSourceSnapshotIdOk returns a tuple with the SourceSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotId

`func (o *FleetCopyResourceInstanceSnapshotRequest2) SetSourceSnapshotId(v string)`

SetSourceSnapshotId sets SourceSnapshotId field to given value.

### HasSourceSnapshotId

`func (o *FleetCopyResourceInstanceSnapshotRequest2) HasSourceSnapshotId() bool`

HasSourceSnapshotId returns a boolean if a field has been set.

### GetTargetRegion

`func (o *FleetCopyResourceInstanceSnapshotRequest2) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### GetTargetRegionOk

`func (o *FleetCopyResourceInstanceSnapshotRequest2) GetTargetRegionOk() (*string, bool)`

GetTargetRegionOk returns a tuple with the TargetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegion

`func (o *FleetCopyResourceInstanceSnapshotRequest2) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


