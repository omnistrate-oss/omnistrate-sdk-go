# DescribeStorageConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the storage config | 
**Id** | **string** | ID of a Storage Config | 
**InfraConfigIDs** | Pointer to **[]string** | The list of infra config IDs associated with the compute config. | [optional] 
**Name** | **string** | Name of the storage config | 
**ServiceId** | **string** | ID of a Service | 
**Volumes** | **map[string]interface{}** | The storage volume config IDs and the corresponding mount path | 

## Methods

### NewDescribeStorageConfigResult

`func NewDescribeStorageConfigResult(description string, id string, name string, serviceId string, volumes map[string]interface{}, ) *DescribeStorageConfigResult`

NewDescribeStorageConfigResult instantiates a new DescribeStorageConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeStorageConfigResultWithDefaults

`func NewDescribeStorageConfigResultWithDefaults() *DescribeStorageConfigResult`

NewDescribeStorageConfigResultWithDefaults instantiates a new DescribeStorageConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DescribeStorageConfigResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeStorageConfigResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeStorageConfigResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeStorageConfigResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeStorageConfigResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeStorageConfigResult) SetId(v string)`

SetId sets Id field to given value.


### GetInfraConfigIDs

`func (o *DescribeStorageConfigResult) GetInfraConfigIDs() []string`

GetInfraConfigIDs returns the InfraConfigIDs field if non-nil, zero value otherwise.

### GetInfraConfigIDsOk

`func (o *DescribeStorageConfigResult) GetInfraConfigIDsOk() (*[]string, bool)`

GetInfraConfigIDsOk returns a tuple with the InfraConfigIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigIDs

`func (o *DescribeStorageConfigResult) SetInfraConfigIDs(v []string)`

SetInfraConfigIDs sets InfraConfigIDs field to given value.

### HasInfraConfigIDs

`func (o *DescribeStorageConfigResult) HasInfraConfigIDs() bool`

HasInfraConfigIDs returns a boolean if a field has been set.

### GetName

`func (o *DescribeStorageConfigResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeStorageConfigResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeStorageConfigResult) SetName(v string)`

SetName sets Name field to given value.


### GetServiceId

`func (o *DescribeStorageConfigResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeStorageConfigResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeStorageConfigResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetVolumes

`func (o *DescribeStorageConfigResult) GetVolumes() map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *DescribeStorageConfigResult) GetVolumesOk() (*map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *DescribeStorageConfigResult) SetVolumes(v map[string]interface{})`

SetVolumes sets Volumes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


