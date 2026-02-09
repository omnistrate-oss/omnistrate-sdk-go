# InstanceSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompleteTime** | **string** | The snapshot time | 
**CreatedTime** | **string** | The snapshot creation time | 
**Encrypted** | **bool** | Whether the snapshot is encrypted | 
**Progress** | **int64** | The backup progress. 0-100 | 
**Region** | **string** | The region name where the snapshot is stored | 
**SnapshotId** | **string** | ID of a Resource Instance Snapshot | 
**SnapshotType** | **string** | The snapshot type | 
**Status** | **string** | The snapshot status | 

## Methods

### NewInstanceSnapshot

`func NewInstanceSnapshot(completeTime string, createdTime string, encrypted bool, progress int64, region string, snapshotId string, snapshotType string, status string, ) *InstanceSnapshot`

NewInstanceSnapshot instantiates a new InstanceSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceSnapshotWithDefaults

`func NewInstanceSnapshotWithDefaults() *InstanceSnapshot`

NewInstanceSnapshotWithDefaults instantiates a new InstanceSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleteTime

`func (o *InstanceSnapshot) GetCompleteTime() string`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *InstanceSnapshot) GetCompleteTimeOk() (*string, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *InstanceSnapshot) SetCompleteTime(v string)`

SetCompleteTime sets CompleteTime field to given value.


### GetCreatedTime

`func (o *InstanceSnapshot) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *InstanceSnapshot) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *InstanceSnapshot) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.


### GetEncrypted

`func (o *InstanceSnapshot) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *InstanceSnapshot) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *InstanceSnapshot) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.


### GetProgress

`func (o *InstanceSnapshot) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *InstanceSnapshot) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *InstanceSnapshot) SetProgress(v int64)`

SetProgress sets Progress field to given value.


### GetRegion

`func (o *InstanceSnapshot) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InstanceSnapshot) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InstanceSnapshot) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSnapshotId

`func (o *InstanceSnapshot) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *InstanceSnapshot) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *InstanceSnapshot) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetSnapshotType

`func (o *InstanceSnapshot) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *InstanceSnapshot) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *InstanceSnapshot) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.


### GetStatus

`func (o *InstanceSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


