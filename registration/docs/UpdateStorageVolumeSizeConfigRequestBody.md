# UpdateStorageVolumeSizeConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceStorageSizeGi** | **string** | The storage size (in Gi) provisioned for the configured instance storage type | 

## Methods

### NewUpdateStorageVolumeSizeConfigRequestBody

`func NewUpdateStorageVolumeSizeConfigRequestBody(instanceStorageSizeGi string, ) *UpdateStorageVolumeSizeConfigRequestBody`

NewUpdateStorageVolumeSizeConfigRequestBody instantiates a new UpdateStorageVolumeSizeConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageVolumeSizeConfigRequestBodyWithDefaults

`func NewUpdateStorageVolumeSizeConfigRequestBodyWithDefaults() *UpdateStorageVolumeSizeConfigRequestBody`

NewUpdateStorageVolumeSizeConfigRequestBodyWithDefaults instantiates a new UpdateStorageVolumeSizeConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceStorageSizeGi

`func (o *UpdateStorageVolumeSizeConfigRequestBody) GetInstanceStorageSizeGi() string`

GetInstanceStorageSizeGi returns the InstanceStorageSizeGi field if non-nil, zero value otherwise.

### GetInstanceStorageSizeGiOk

`func (o *UpdateStorageVolumeSizeConfigRequestBody) GetInstanceStorageSizeGiOk() (*string, bool)`

GetInstanceStorageSizeGiOk returns a tuple with the InstanceStorageSizeGi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStorageSizeGi

`func (o *UpdateStorageVolumeSizeConfigRequestBody) SetInstanceStorageSizeGi(v string)`

SetInstanceStorageSizeGi sets InstanceStorageSizeGi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


