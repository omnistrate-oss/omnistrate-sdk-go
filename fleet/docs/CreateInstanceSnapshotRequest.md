# CreateInstanceSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | The instance ID | 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 
**TargetRegion** | Pointer to **string** | The target region to create the snapshot in. If not specified, use the same region as the instance | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateInstanceSnapshotRequest

`func NewCreateInstanceSnapshotRequest(instanceId string, token string, ) *CreateInstanceSnapshotRequest`

NewCreateInstanceSnapshotRequest instantiates a new CreateInstanceSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceSnapshotRequestWithDefaults

`func NewCreateInstanceSnapshotRequestWithDefaults() *CreateInstanceSnapshotRequest`

NewCreateInstanceSnapshotRequestWithDefaults instantiates a new CreateInstanceSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *CreateInstanceSnapshotRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CreateInstanceSnapshotRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CreateInstanceSnapshotRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetSubscriptionId

`func (o *CreateInstanceSnapshotRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateInstanceSnapshotRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateInstanceSnapshotRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateInstanceSnapshotRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTargetRegion

`func (o *CreateInstanceSnapshotRequest) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### GetTargetRegionOk

`func (o *CreateInstanceSnapshotRequest) GetTargetRegionOk() (*string, bool)`

GetTargetRegionOk returns a tuple with the TargetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegion

`func (o *CreateInstanceSnapshotRequest) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.

### HasTargetRegion

`func (o *CreateInstanceSnapshotRequest) HasTargetRegion() bool`

HasTargetRegion returns a boolean if a field has been set.

### GetToken

`func (o *CreateInstanceSnapshotRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateInstanceSnapshotRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateInstanceSnapshotRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


