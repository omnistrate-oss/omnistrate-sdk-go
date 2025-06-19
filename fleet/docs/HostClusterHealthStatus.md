# HostClusterHealthStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedEntities** | Pointer to [**[]EntityHealth**](EntityHealth.md) | List of entities that have failed health checks | [optional] 
**OverallStatus** | **string** | The heath status of a resource | 
**StatusMessage** | Pointer to **string** | Detailed message about the health status of the host cluster | [optional] 
**TotalNumberOfEntities** | **int64** | Total number of entities under management in the host cluster | 
**TotalNumberOfEntitiesByType** | **map[string]int64** | Map of entity types to their counts | 
**TotalNumberOfFailedEntities** | **int64** | Number of entities that have failed health checks | 
**TotalNumberOfFailedEntitiesByType** | **map[string]int64** | Map of entity types to their counts that have failed health checks | 
**TotalNumberOfHealthyEntities** | **int64** | Number of entities that are healthy | 
**TotalNumberOfHealthyEntitiesByType** | **map[string]int64** | Map of entity types to their counts that are healthy | 

## Methods

### NewHostClusterHealthStatus

`func NewHostClusterHealthStatus(overallStatus string, totalNumberOfEntities int64, totalNumberOfEntitiesByType map[string]int64, totalNumberOfFailedEntities int64, totalNumberOfFailedEntitiesByType map[string]int64, totalNumberOfHealthyEntities int64, totalNumberOfHealthyEntitiesByType map[string]int64, ) *HostClusterHealthStatus`

NewHostClusterHealthStatus instantiates a new HostClusterHealthStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostClusterHealthStatusWithDefaults

`func NewHostClusterHealthStatusWithDefaults() *HostClusterHealthStatus`

NewHostClusterHealthStatusWithDefaults instantiates a new HostClusterHealthStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedEntities

`func (o *HostClusterHealthStatus) GetFailedEntities() []EntityHealth`

GetFailedEntities returns the FailedEntities field if non-nil, zero value otherwise.

### GetFailedEntitiesOk

`func (o *HostClusterHealthStatus) GetFailedEntitiesOk() (*[]EntityHealth, bool)`

GetFailedEntitiesOk returns a tuple with the FailedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedEntities

`func (o *HostClusterHealthStatus) SetFailedEntities(v []EntityHealth)`

SetFailedEntities sets FailedEntities field to given value.

### HasFailedEntities

`func (o *HostClusterHealthStatus) HasFailedEntities() bool`

HasFailedEntities returns a boolean if a field has been set.

### GetOverallStatus

`func (o *HostClusterHealthStatus) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *HostClusterHealthStatus) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *HostClusterHealthStatus) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.


### GetStatusMessage

`func (o *HostClusterHealthStatus) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *HostClusterHealthStatus) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *HostClusterHealthStatus) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *HostClusterHealthStatus) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetTotalNumberOfEntities

`func (o *HostClusterHealthStatus) GetTotalNumberOfEntities() int64`

GetTotalNumberOfEntities returns the TotalNumberOfEntities field if non-nil, zero value otherwise.

### GetTotalNumberOfEntitiesOk

`func (o *HostClusterHealthStatus) GetTotalNumberOfEntitiesOk() (*int64, bool)`

GetTotalNumberOfEntitiesOk returns a tuple with the TotalNumberOfEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfEntities

`func (o *HostClusterHealthStatus) SetTotalNumberOfEntities(v int64)`

SetTotalNumberOfEntities sets TotalNumberOfEntities field to given value.


### GetTotalNumberOfEntitiesByType

`func (o *HostClusterHealthStatus) GetTotalNumberOfEntitiesByType() map[string]int64`

GetTotalNumberOfEntitiesByType returns the TotalNumberOfEntitiesByType field if non-nil, zero value otherwise.

### GetTotalNumberOfEntitiesByTypeOk

`func (o *HostClusterHealthStatus) GetTotalNumberOfEntitiesByTypeOk() (*map[string]int64, bool)`

GetTotalNumberOfEntitiesByTypeOk returns a tuple with the TotalNumberOfEntitiesByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfEntitiesByType

`func (o *HostClusterHealthStatus) SetTotalNumberOfEntitiesByType(v map[string]int64)`

SetTotalNumberOfEntitiesByType sets TotalNumberOfEntitiesByType field to given value.


### GetTotalNumberOfFailedEntities

`func (o *HostClusterHealthStatus) GetTotalNumberOfFailedEntities() int64`

GetTotalNumberOfFailedEntities returns the TotalNumberOfFailedEntities field if non-nil, zero value otherwise.

### GetTotalNumberOfFailedEntitiesOk

`func (o *HostClusterHealthStatus) GetTotalNumberOfFailedEntitiesOk() (*int64, bool)`

GetTotalNumberOfFailedEntitiesOk returns a tuple with the TotalNumberOfFailedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfFailedEntities

`func (o *HostClusterHealthStatus) SetTotalNumberOfFailedEntities(v int64)`

SetTotalNumberOfFailedEntities sets TotalNumberOfFailedEntities field to given value.


### GetTotalNumberOfFailedEntitiesByType

`func (o *HostClusterHealthStatus) GetTotalNumberOfFailedEntitiesByType() map[string]int64`

GetTotalNumberOfFailedEntitiesByType returns the TotalNumberOfFailedEntitiesByType field if non-nil, zero value otherwise.

### GetTotalNumberOfFailedEntitiesByTypeOk

`func (o *HostClusterHealthStatus) GetTotalNumberOfFailedEntitiesByTypeOk() (*map[string]int64, bool)`

GetTotalNumberOfFailedEntitiesByTypeOk returns a tuple with the TotalNumberOfFailedEntitiesByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfFailedEntitiesByType

`func (o *HostClusterHealthStatus) SetTotalNumberOfFailedEntitiesByType(v map[string]int64)`

SetTotalNumberOfFailedEntitiesByType sets TotalNumberOfFailedEntitiesByType field to given value.


### GetTotalNumberOfHealthyEntities

`func (o *HostClusterHealthStatus) GetTotalNumberOfHealthyEntities() int64`

GetTotalNumberOfHealthyEntities returns the TotalNumberOfHealthyEntities field if non-nil, zero value otherwise.

### GetTotalNumberOfHealthyEntitiesOk

`func (o *HostClusterHealthStatus) GetTotalNumberOfHealthyEntitiesOk() (*int64, bool)`

GetTotalNumberOfHealthyEntitiesOk returns a tuple with the TotalNumberOfHealthyEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfHealthyEntities

`func (o *HostClusterHealthStatus) SetTotalNumberOfHealthyEntities(v int64)`

SetTotalNumberOfHealthyEntities sets TotalNumberOfHealthyEntities field to given value.


### GetTotalNumberOfHealthyEntitiesByType

`func (o *HostClusterHealthStatus) GetTotalNumberOfHealthyEntitiesByType() map[string]int64`

GetTotalNumberOfHealthyEntitiesByType returns the TotalNumberOfHealthyEntitiesByType field if non-nil, zero value otherwise.

### GetTotalNumberOfHealthyEntitiesByTypeOk

`func (o *HostClusterHealthStatus) GetTotalNumberOfHealthyEntitiesByTypeOk() (*map[string]int64, bool)`

GetTotalNumberOfHealthyEntitiesByTypeOk returns a tuple with the TotalNumberOfHealthyEntitiesByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfHealthyEntitiesByType

`func (o *HostClusterHealthStatus) SetTotalNumberOfHealthyEntitiesByType(v map[string]int64)`

SetTotalNumberOfHealthyEntitiesByType sets TotalNumberOfHealthyEntitiesByType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


