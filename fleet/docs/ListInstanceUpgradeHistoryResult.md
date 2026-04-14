# ListInstanceUpgradeHistoryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Upgrades** | [**[]InstanceUpgradeHistory**](InstanceUpgradeHistory.md) | The list of historical upgrades for this instance. | 

## Methods

### NewListInstanceUpgradeHistoryResult

`func NewListInstanceUpgradeHistoryResult(environmentId string, instanceId string, serviceId string, upgrades []InstanceUpgradeHistory, ) *ListInstanceUpgradeHistoryResult`

NewListInstanceUpgradeHistoryResult instantiates a new ListInstanceUpgradeHistoryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceUpgradeHistoryResultWithDefaults

`func NewListInstanceUpgradeHistoryResultWithDefaults() *ListInstanceUpgradeHistoryResult`

NewListInstanceUpgradeHistoryResultWithDefaults instantiates a new ListInstanceUpgradeHistoryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ListInstanceUpgradeHistoryResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ListInstanceUpgradeHistoryResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ListInstanceUpgradeHistoryResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *ListInstanceUpgradeHistoryResult) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListInstanceUpgradeHistoryResult) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListInstanceUpgradeHistoryResult) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetNextPageToken

`func (o *ListInstanceUpgradeHistoryResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListInstanceUpgradeHistoryResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListInstanceUpgradeHistoryResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListInstanceUpgradeHistoryResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetServiceId

`func (o *ListInstanceUpgradeHistoryResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListInstanceUpgradeHistoryResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListInstanceUpgradeHistoryResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetUpgrades

`func (o *ListInstanceUpgradeHistoryResult) GetUpgrades() []InstanceUpgradeHistory`

GetUpgrades returns the Upgrades field if non-nil, zero value otherwise.

### GetUpgradesOk

`func (o *ListInstanceUpgradeHistoryResult) GetUpgradesOk() (*[]InstanceUpgradeHistory, bool)`

GetUpgradesOk returns a tuple with the Upgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrades

`func (o *ListInstanceUpgradeHistoryResult) SetUpgrades(v []InstanceUpgradeHistory)`

SetUpgrades sets Upgrades field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


