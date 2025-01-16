# RestoreResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkType** | Pointer to **string** | The network type | [optional] 
**TargetRestoreTime** | **string** | The target restore time | 

## Methods

### NewRestoreResourceInstanceRequest2

`func NewRestoreResourceInstanceRequest2(targetRestoreTime string, ) *RestoreResourceInstanceRequest2`

NewRestoreResourceInstanceRequest2 instantiates a new RestoreResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreResourceInstanceRequest2WithDefaults

`func NewRestoreResourceInstanceRequest2WithDefaults() *RestoreResourceInstanceRequest2`

NewRestoreResourceInstanceRequest2WithDefaults instantiates a new RestoreResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *RestoreResourceInstanceRequest2) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *RestoreResourceInstanceRequest2) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *RestoreResourceInstanceRequest2) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *RestoreResourceInstanceRequest2) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetTargetRestoreTime

`func (o *RestoreResourceInstanceRequest2) GetTargetRestoreTime() string`

GetTargetRestoreTime returns the TargetRestoreTime field if non-nil, zero value otherwise.

### GetTargetRestoreTimeOk

`func (o *RestoreResourceInstanceRequest2) GetTargetRestoreTimeOk() (*string, bool)`

GetTargetRestoreTimeOk returns a tuple with the TargetRestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRestoreTime

`func (o *RestoreResourceInstanceRequest2) SetTargetRestoreTime(v string)`

SetTargetRestoreTime sets TargetRestoreTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


