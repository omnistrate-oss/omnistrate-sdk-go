# EFSFileSystemConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerformanceMode** | **string** | The performance mode of the EFS file system | 
**ProvisionedThroughputInMibps** | Pointer to **float64** | The throughput, measured in MiBps, that you want to provision for the EFS file system, if throughput mode if provisioned | [optional] 
**ThroughputMode** | **string** | The throughput mode of the EFS file system | 

## Methods

### NewEFSFileSystemConfiguration

`func NewEFSFileSystemConfiguration(performanceMode string, throughputMode string, ) *EFSFileSystemConfiguration`

NewEFSFileSystemConfiguration instantiates a new EFSFileSystemConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEFSFileSystemConfigurationWithDefaults

`func NewEFSFileSystemConfigurationWithDefaults() *EFSFileSystemConfiguration`

NewEFSFileSystemConfigurationWithDefaults instantiates a new EFSFileSystemConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerformanceMode

`func (o *EFSFileSystemConfiguration) GetPerformanceMode() string`

GetPerformanceMode returns the PerformanceMode field if non-nil, zero value otherwise.

### GetPerformanceModeOk

`func (o *EFSFileSystemConfiguration) GetPerformanceModeOk() (*string, bool)`

GetPerformanceModeOk returns a tuple with the PerformanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceMode

`func (o *EFSFileSystemConfiguration) SetPerformanceMode(v string)`

SetPerformanceMode sets PerformanceMode field to given value.


### GetProvisionedThroughputInMibps

`func (o *EFSFileSystemConfiguration) GetProvisionedThroughputInMibps() float64`

GetProvisionedThroughputInMibps returns the ProvisionedThroughputInMibps field if non-nil, zero value otherwise.

### GetProvisionedThroughputInMibpsOk

`func (o *EFSFileSystemConfiguration) GetProvisionedThroughputInMibpsOk() (*float64, bool)`

GetProvisionedThroughputInMibpsOk returns a tuple with the ProvisionedThroughputInMibps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedThroughputInMibps

`func (o *EFSFileSystemConfiguration) SetProvisionedThroughputInMibps(v float64)`

SetProvisionedThroughputInMibps sets ProvisionedThroughputInMibps field to given value.

### HasProvisionedThroughputInMibps

`func (o *EFSFileSystemConfiguration) HasProvisionedThroughputInMibps() bool`

HasProvisionedThroughputInMibps returns a boolean if a field has been set.

### GetThroughputMode

`func (o *EFSFileSystemConfiguration) GetThroughputMode() string`

GetThroughputMode returns the ThroughputMode field if non-nil, zero value otherwise.

### GetThroughputModeOk

`func (o *EFSFileSystemConfiguration) GetThroughputModeOk() (*string, bool)`

GetThroughputModeOk returns a tuple with the ThroughputMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputMode

`func (o *EFSFileSystemConfiguration) SetThroughputMode(v string)`

SetThroughputMode sets ThroughputMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


