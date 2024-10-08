# RestoreResourceInstanceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkType** | Pointer to **string** | The network type | [optional] 
**TargetRestoreTime** | **string** | The target restore time | 

## Methods

### NewRestoreResourceInstanceRequestBody

`func NewRestoreResourceInstanceRequestBody(targetRestoreTime string, ) *RestoreResourceInstanceRequestBody`

NewRestoreResourceInstanceRequestBody instantiates a new RestoreResourceInstanceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreResourceInstanceRequestBodyWithDefaults

`func NewRestoreResourceInstanceRequestBodyWithDefaults() *RestoreResourceInstanceRequestBody`

NewRestoreResourceInstanceRequestBodyWithDefaults instantiates a new RestoreResourceInstanceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkType

`func (o *RestoreResourceInstanceRequestBody) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *RestoreResourceInstanceRequestBody) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *RestoreResourceInstanceRequestBody) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *RestoreResourceInstanceRequestBody) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetTargetRestoreTime

`func (o *RestoreResourceInstanceRequestBody) GetTargetRestoreTime() string`

GetTargetRestoreTime returns the TargetRestoreTime field if non-nil, zero value otherwise.

### GetTargetRestoreTimeOk

`func (o *RestoreResourceInstanceRequestBody) GetTargetRestoreTimeOk() (*string, bool)`

GetTargetRestoreTimeOk returns a tuple with the TargetRestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRestoreTime

`func (o *RestoreResourceInstanceRequestBody) SetTargetRestoreTime(v string)`

SetTargetRestoreTime sets TargetRestoreTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


