# BackupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EarliestRestoreTime** | **string** | The earliest restore time | 

## Methods

### NewBackupStatus

`func NewBackupStatus(earliestRestoreTime string, ) *BackupStatus`

NewBackupStatus instantiates a new BackupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStatusWithDefaults

`func NewBackupStatusWithDefaults() *BackupStatus`

NewBackupStatusWithDefaults instantiates a new BackupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


