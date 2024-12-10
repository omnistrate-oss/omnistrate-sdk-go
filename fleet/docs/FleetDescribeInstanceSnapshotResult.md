# FleetDescribeInstanceSnapshotResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompleteTime** | **string** | The snapshot time | 
**CreatedTime** | **string** | The snapshot creation time | 
**Encrypted** | **bool** | Whether the snapshot is encrypted | 
**EnvironmentId** | **string** | ID of a Service Environment | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProductTierVersion** | **string** | The product tier version | 
**Progress** | **int64** | The backup progress. 0-100 | 
**ServiceId** | **string** | ID of a Service | 
**SnapshotId** | **string** | ID of a Resource Instance Snapshot | 
**SourceInstanceId** | **string** | ID of a Resource Instance | 
**Status** | **string** | The snapshot status | 

## Methods

### NewFleetDescribeInstanceSnapshotResult

`func NewFleetDescribeInstanceSnapshotResult(completeTime string, createdTime string, encrypted bool, environmentId string, productTierId string, productTierVersion string, progress int64, serviceId string, snapshotId string, sourceInstanceId string, status string, ) *FleetDescribeInstanceSnapshotResult`

NewFleetDescribeInstanceSnapshotResult instantiates a new FleetDescribeInstanceSnapshotResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDescribeInstanceSnapshotResultWithDefaults

`func NewFleetDescribeInstanceSnapshotResultWithDefaults() *FleetDescribeInstanceSnapshotResult`

NewFleetDescribeInstanceSnapshotResultWithDefaults instantiates a new FleetDescribeInstanceSnapshotResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleteTime

`func (o *FleetDescribeInstanceSnapshotResult) GetCompleteTime() string`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *FleetDescribeInstanceSnapshotResult) GetCompleteTimeOk() (*string, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *FleetDescribeInstanceSnapshotResult) SetCompleteTime(v string)`

SetCompleteTime sets CompleteTime field to given value.


### GetCreatedTime

`func (o *FleetDescribeInstanceSnapshotResult) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FleetDescribeInstanceSnapshotResult) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FleetDescribeInstanceSnapshotResult) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.


### GetEncrypted

`func (o *FleetDescribeInstanceSnapshotResult) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *FleetDescribeInstanceSnapshotResult) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *FleetDescribeInstanceSnapshotResult) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.


### GetEnvironmentId

`func (o *FleetDescribeInstanceSnapshotResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetDescribeInstanceSnapshotResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetProductTierId

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *FleetDescribeInstanceSnapshotResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierVersion

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *FleetDescribeInstanceSnapshotResult) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.


### GetProgress

`func (o *FleetDescribeInstanceSnapshotResult) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FleetDescribeInstanceSnapshotResult) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FleetDescribeInstanceSnapshotResult) SetProgress(v int64)`

SetProgress sets Progress field to given value.


### GetServiceId

`func (o *FleetDescribeInstanceSnapshotResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetDescribeInstanceSnapshotResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSnapshotId

`func (o *FleetDescribeInstanceSnapshotResult) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *FleetDescribeInstanceSnapshotResult) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetSourceInstanceId

`func (o *FleetDescribeInstanceSnapshotResult) GetSourceInstanceId() string`

GetSourceInstanceId returns the SourceInstanceId field if non-nil, zero value otherwise.

### GetSourceInstanceIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetSourceInstanceIdOk() (*string, bool)`

GetSourceInstanceIdOk returns a tuple with the SourceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstanceId

`func (o *FleetDescribeInstanceSnapshotResult) SetSourceInstanceId(v string)`

SetSourceInstanceId sets SourceInstanceId field to given value.


### GetStatus

`func (o *FleetDescribeInstanceSnapshotResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetDescribeInstanceSnapshotResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetDescribeInstanceSnapshotResult) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


