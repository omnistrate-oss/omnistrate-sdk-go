# BackupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPeriodInHours** | **int64** | The backup period in hours | 
**BackupRetentionInDays** | **int64** | The backup retention in days | 
**EarliestRestoreTime** | Pointer to **string** | The earliest restore time | [optional] 
**LastBackupTime** | Pointer to **string** | The last backup time | [optional] 

## Methods

### NewBackupStatus

`func NewBackupStatus(backupPeriodInHours int64, backupRetentionInDays int64, ) *BackupStatus`

NewBackupStatus instantiates a new BackupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStatusWithDefaults

`func NewBackupStatusWithDefaults() *BackupStatus`

NewBackupStatusWithDefaults instantiates a new BackupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPeriodInHours

`func (o *BackupStatus) GetBackupPeriodInHours() int64`

GetBackupPeriodInHours returns the BackupPeriodInHours field if non-nil, zero value otherwise.

### GetBackupPeriodInHoursOk

`func (o *BackupStatus) GetBackupPeriodInHoursOk() (*int64, bool)`

GetBackupPeriodInHoursOk returns a tuple with the BackupPeriodInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPeriodInHours

`func (o *BackupStatus) SetBackupPeriodInHours(v int64)`

SetBackupPeriodInHours sets BackupPeriodInHours field to given value.


### GetBackupRetentionInDays

`func (o *BackupStatus) GetBackupRetentionInDays() int64`

GetBackupRetentionInDays returns the BackupRetentionInDays field if non-nil, zero value otherwise.

### GetBackupRetentionInDaysOk

`func (o *BackupStatus) GetBackupRetentionInDaysOk() (*int64, bool)`

GetBackupRetentionInDaysOk returns a tuple with the BackupRetentionInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionInDays

`func (o *BackupStatus) SetBackupRetentionInDays(v int64)`

SetBackupRetentionInDays sets BackupRetentionInDays field to given value.


### GetEarliestRestoreTime

`func (o *BackupStatus) GetEarliestRestoreTime() string`

GetEarliestRestoreTime returns the EarliestRestoreTime field if non-nil, zero value otherwise.

### GetEarliestRestoreTimeOk

`func (o *BackupStatus) GetEarliestRestoreTimeOk() (*string, bool)`

GetEarliestRestoreTimeOk returns a tuple with the EarliestRestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestRestoreTime

`func (o *BackupStatus) SetEarliestRestoreTime(v string)`

SetEarliestRestoreTime sets EarliestRestoreTime field to given value.

### HasEarliestRestoreTime

`func (o *BackupStatus) HasEarliestRestoreTime() bool`

HasEarliestRestoreTime returns a boolean if a field has been set.

### GetLastBackupTime

`func (o *BackupStatus) GetLastBackupTime() string`

GetLastBackupTime returns the LastBackupTime field if non-nil, zero value otherwise.

### GetLastBackupTimeOk

`func (o *BackupStatus) GetLastBackupTimeOk() (*string, bool)`

GetLastBackupTimeOk returns a tuple with the LastBackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupTime

`func (o *BackupStatus) SetLastBackupTime(v string)`

SetLastBackupTime sets LastBackupTime field to given value.

### HasLastBackupTime

`func (o *BackupStatus) HasLastBackupTime() bool`

HasLastBackupTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


