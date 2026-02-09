# CopyInstanceSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceSnapshotId** | **string** | ID of a Resource Instance Snapshot | 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 
**TargetRegion** | **string** | The target region to copy the snapshot to | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCopyInstanceSnapshotRequest

`func NewCopyInstanceSnapshotRequest(sourceSnapshotId string, targetRegion string, token string, ) *CopyInstanceSnapshotRequest`

NewCopyInstanceSnapshotRequest instantiates a new CopyInstanceSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyInstanceSnapshotRequestWithDefaults

`func NewCopyInstanceSnapshotRequestWithDefaults() *CopyInstanceSnapshotRequest`

NewCopyInstanceSnapshotRequestWithDefaults instantiates a new CopyInstanceSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceSnapshotId

`func (o *CopyInstanceSnapshotRequest) GetSourceSnapshotId() string`

GetSourceSnapshotId returns the SourceSnapshotId field if non-nil, zero value otherwise.

### GetSourceSnapshotIdOk

`func (o *CopyInstanceSnapshotRequest) GetSourceSnapshotIdOk() (*string, bool)`

GetSourceSnapshotIdOk returns a tuple with the SourceSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotId

`func (o *CopyInstanceSnapshotRequest) SetSourceSnapshotId(v string)`

SetSourceSnapshotId sets SourceSnapshotId field to given value.


### GetSubscriptionId

`func (o *CopyInstanceSnapshotRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CopyInstanceSnapshotRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CopyInstanceSnapshotRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CopyInstanceSnapshotRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTargetRegion

`func (o *CopyInstanceSnapshotRequest) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### GetTargetRegionOk

`func (o *CopyInstanceSnapshotRequest) GetTargetRegionOk() (*string, bool)`

GetTargetRegionOk returns a tuple with the TargetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegion

`func (o *CopyInstanceSnapshotRequest) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.


### GetToken

`func (o *CopyInstanceSnapshotRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CopyInstanceSnapshotRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CopyInstanceSnapshotRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


