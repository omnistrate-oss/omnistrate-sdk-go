# RestoreFromInstanceSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomNetworkId** | Pointer to **string** | Custom network for resource | [optional] 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**SnapshotId** | **string** | ID of a Resource Instance Snapshot | 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRestoreFromInstanceSnapshotRequest

`func NewRestoreFromInstanceSnapshotRequest(snapshotId string, token string, ) *RestoreFromInstanceSnapshotRequest`

NewRestoreFromInstanceSnapshotRequest instantiates a new RestoreFromInstanceSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreFromInstanceSnapshotRequestWithDefaults

`func NewRestoreFromInstanceSnapshotRequestWithDefaults() *RestoreFromInstanceSnapshotRequest`

NewRestoreFromInstanceSnapshotRequestWithDefaults instantiates a new RestoreFromInstanceSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomNetworkId

`func (o *RestoreFromInstanceSnapshotRequest) GetCustomNetworkId() string`

GetCustomNetworkId returns the CustomNetworkId field if non-nil, zero value otherwise.

### GetCustomNetworkIdOk

`func (o *RestoreFromInstanceSnapshotRequest) GetCustomNetworkIdOk() (*string, bool)`

GetCustomNetworkIdOk returns a tuple with the CustomNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworkId

`func (o *RestoreFromInstanceSnapshotRequest) SetCustomNetworkId(v string)`

SetCustomNetworkId sets CustomNetworkId field to given value.

### HasCustomNetworkId

`func (o *RestoreFromInstanceSnapshotRequest) HasCustomNetworkId() bool`

HasCustomNetworkId returns a boolean if a field has been set.

### GetNetworkType

`func (o *RestoreFromInstanceSnapshotRequest) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *RestoreFromInstanceSnapshotRequest) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *RestoreFromInstanceSnapshotRequest) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *RestoreFromInstanceSnapshotRequest) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetSnapshotId

`func (o *RestoreFromInstanceSnapshotRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *RestoreFromInstanceSnapshotRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *RestoreFromInstanceSnapshotRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetSubscriptionId

`func (o *RestoreFromInstanceSnapshotRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RestoreFromInstanceSnapshotRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RestoreFromInstanceSnapshotRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RestoreFromInstanceSnapshotRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetToken

`func (o *RestoreFromInstanceSnapshotRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RestoreFromInstanceSnapshotRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RestoreFromInstanceSnapshotRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


