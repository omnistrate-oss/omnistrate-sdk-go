# ListResourceInstanceSnapshotsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | Pointer to [**[]InstanceSnapshot**](InstanceSnapshot.md) | The instance snapshots | [optional] 

## Methods

### NewListResourceInstanceSnapshotsResult

`func NewListResourceInstanceSnapshotsResult() *ListResourceInstanceSnapshotsResult`

NewListResourceInstanceSnapshotsResult instantiates a new ListResourceInstanceSnapshotsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourceInstanceSnapshotsResultWithDefaults

`func NewListResourceInstanceSnapshotsResultWithDefaults() *ListResourceInstanceSnapshotsResult`

NewListResourceInstanceSnapshotsResultWithDefaults instantiates a new ListResourceInstanceSnapshotsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *ListResourceInstanceSnapshotsResult) GetSnapshots() []InstanceSnapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ListResourceInstanceSnapshotsResult) GetSnapshotsOk() (*[]InstanceSnapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ListResourceInstanceSnapshotsResult) SetSnapshots(v []InstanceSnapshot)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *ListResourceInstanceSnapshotsResult) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


