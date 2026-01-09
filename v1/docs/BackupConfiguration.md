# BackupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPeriodInHours** | **int64** | The period in hours to take backups | 
**BackupRetentionInDays** | **int64** | The number of days to retain backups | 
**SnapshotBeforeDeletion** | Pointer to **bool** | Controls whether a final manual snapshot is automatically created before resource deletion. Defaults to false (disabled) if not specified. | [optional] 

## Methods

### NewBackupConfiguration

`func NewBackupConfiguration(backupPeriodInHours int64, backupRetentionInDays int64, ) *BackupConfiguration`

NewBackupConfiguration instantiates a new BackupConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupConfigurationWithDefaults

`func NewBackupConfigurationWithDefaults() *BackupConfiguration`

NewBackupConfigurationWithDefaults instantiates a new BackupConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPeriodInHours

`func (o *BackupConfiguration) GetBackupPeriodInHours() int64`

GetBackupPeriodInHours returns the BackupPeriodInHours field if non-nil, zero value otherwise.

### GetBackupPeriodInHoursOk

`func (o *BackupConfiguration) GetBackupPeriodInHoursOk() (*int64, bool)`

GetBackupPeriodInHoursOk returns a tuple with the BackupPeriodInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPeriodInHours

`func (o *BackupConfiguration) SetBackupPeriodInHours(v int64)`

SetBackupPeriodInHours sets BackupPeriodInHours field to given value.


### GetBackupRetentionInDays

`func (o *BackupConfiguration) GetBackupRetentionInDays() int64`

GetBackupRetentionInDays returns the BackupRetentionInDays field if non-nil, zero value otherwise.

### GetBackupRetentionInDaysOk

`func (o *BackupConfiguration) GetBackupRetentionInDaysOk() (*int64, bool)`

GetBackupRetentionInDaysOk returns a tuple with the BackupRetentionInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionInDays

`func (o *BackupConfiguration) SetBackupRetentionInDays(v int64)`

SetBackupRetentionInDays sets BackupRetentionInDays field to given value.


### GetSnapshotBeforeDeletion

`func (o *BackupConfiguration) GetSnapshotBeforeDeletion() bool`

GetSnapshotBeforeDeletion returns the SnapshotBeforeDeletion field if non-nil, zero value otherwise.

### GetSnapshotBeforeDeletionOk

`func (o *BackupConfiguration) GetSnapshotBeforeDeletionOk() (*bool, bool)`

GetSnapshotBeforeDeletionOk returns a tuple with the SnapshotBeforeDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotBeforeDeletion

`func (o *BackupConfiguration) SetSnapshotBeforeDeletion(v bool)`

SetSnapshotBeforeDeletion sets SnapshotBeforeDeletion field to given value.

### HasSnapshotBeforeDeletion

`func (o *BackupConfiguration) HasSnapshotBeforeDeletion() bool`

HasSnapshotBeforeDeletion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


