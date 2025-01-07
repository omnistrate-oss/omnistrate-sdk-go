# HelmRuntimeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableHooks** | **bool** | Disable Helm hooks | 
**Recreate** | **bool** | Recreate the Helm package if it already exists | 
**ResetThenReuseValues** | **bool** | Reset then reuse values for the Helm package before applying | 
**ResetValues** | **bool** | Reset values for the Helm package before applying | 
**ReuseValues** | **bool** | Reuse values for the Helm package before applying | 
**SkipCRDs** | **bool** | Skip CRDs for the Helm package | 
**TimeoutNanos** | **int64** | Timeout (nanos) for the Helm package to be deployed | 
**UpgradeCRDs** | **bool** | Upgrade CRDs for the Helm package | 
**Wait** | **bool** | Wait for the Helm package to be deployed | 
**WaitForJobs** | **bool** | Wait for all jobs to be completed | 

## Methods

### NewHelmRuntimeConfiguration

`func NewHelmRuntimeConfiguration(disableHooks bool, recreate bool, resetThenReuseValues bool, resetValues bool, reuseValues bool, skipCRDs bool, timeoutNanos int64, upgradeCRDs bool, wait bool, waitForJobs bool, ) *HelmRuntimeConfiguration`

NewHelmRuntimeConfiguration instantiates a new HelmRuntimeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRuntimeConfigurationWithDefaults

`func NewHelmRuntimeConfigurationWithDefaults() *HelmRuntimeConfiguration`

NewHelmRuntimeConfigurationWithDefaults instantiates a new HelmRuntimeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableHooks

`func (o *HelmRuntimeConfiguration) GetDisableHooks() bool`

GetDisableHooks returns the DisableHooks field if non-nil, zero value otherwise.

### GetDisableHooksOk

`func (o *HelmRuntimeConfiguration) GetDisableHooksOk() (*bool, bool)`

GetDisableHooksOk returns a tuple with the DisableHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHooks

`func (o *HelmRuntimeConfiguration) SetDisableHooks(v bool)`

SetDisableHooks sets DisableHooks field to given value.


### GetRecreate

`func (o *HelmRuntimeConfiguration) GetRecreate() bool`

GetRecreate returns the Recreate field if non-nil, zero value otherwise.

### GetRecreateOk

`func (o *HelmRuntimeConfiguration) GetRecreateOk() (*bool, bool)`

GetRecreateOk returns a tuple with the Recreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecreate

`func (o *HelmRuntimeConfiguration) SetRecreate(v bool)`

SetRecreate sets Recreate field to given value.


### GetResetThenReuseValues

`func (o *HelmRuntimeConfiguration) GetResetThenReuseValues() bool`

GetResetThenReuseValues returns the ResetThenReuseValues field if non-nil, zero value otherwise.

### GetResetThenReuseValuesOk

`func (o *HelmRuntimeConfiguration) GetResetThenReuseValuesOk() (*bool, bool)`

GetResetThenReuseValuesOk returns a tuple with the ResetThenReuseValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetThenReuseValues

`func (o *HelmRuntimeConfiguration) SetResetThenReuseValues(v bool)`

SetResetThenReuseValues sets ResetThenReuseValues field to given value.


### GetResetValues

`func (o *HelmRuntimeConfiguration) GetResetValues() bool`

GetResetValues returns the ResetValues field if non-nil, zero value otherwise.

### GetResetValuesOk

`func (o *HelmRuntimeConfiguration) GetResetValuesOk() (*bool, bool)`

GetResetValuesOk returns a tuple with the ResetValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetValues

`func (o *HelmRuntimeConfiguration) SetResetValues(v bool)`

SetResetValues sets ResetValues field to given value.


### GetReuseValues

`func (o *HelmRuntimeConfiguration) GetReuseValues() bool`

GetReuseValues returns the ReuseValues field if non-nil, zero value otherwise.

### GetReuseValuesOk

`func (o *HelmRuntimeConfiguration) GetReuseValuesOk() (*bool, bool)`

GetReuseValuesOk returns a tuple with the ReuseValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseValues

`func (o *HelmRuntimeConfiguration) SetReuseValues(v bool)`

SetReuseValues sets ReuseValues field to given value.


### GetSkipCRDs

`func (o *HelmRuntimeConfiguration) GetSkipCRDs() bool`

GetSkipCRDs returns the SkipCRDs field if non-nil, zero value otherwise.

### GetSkipCRDsOk

`func (o *HelmRuntimeConfiguration) GetSkipCRDsOk() (*bool, bool)`

GetSkipCRDsOk returns a tuple with the SkipCRDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCRDs

`func (o *HelmRuntimeConfiguration) SetSkipCRDs(v bool)`

SetSkipCRDs sets SkipCRDs field to given value.


### GetTimeoutNanos

`func (o *HelmRuntimeConfiguration) GetTimeoutNanos() int64`

GetTimeoutNanos returns the TimeoutNanos field if non-nil, zero value otherwise.

### GetTimeoutNanosOk

`func (o *HelmRuntimeConfiguration) GetTimeoutNanosOk() (*int64, bool)`

GetTimeoutNanosOk returns a tuple with the TimeoutNanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutNanos

`func (o *HelmRuntimeConfiguration) SetTimeoutNanos(v int64)`

SetTimeoutNanos sets TimeoutNanos field to given value.


### GetUpgradeCRDs

`func (o *HelmRuntimeConfiguration) GetUpgradeCRDs() bool`

GetUpgradeCRDs returns the UpgradeCRDs field if non-nil, zero value otherwise.

### GetUpgradeCRDsOk

`func (o *HelmRuntimeConfiguration) GetUpgradeCRDsOk() (*bool, bool)`

GetUpgradeCRDsOk returns a tuple with the UpgradeCRDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeCRDs

`func (o *HelmRuntimeConfiguration) SetUpgradeCRDs(v bool)`

SetUpgradeCRDs sets UpgradeCRDs field to given value.


### GetWait

`func (o *HelmRuntimeConfiguration) GetWait() bool`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *HelmRuntimeConfiguration) GetWaitOk() (*bool, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *HelmRuntimeConfiguration) SetWait(v bool)`

SetWait sets Wait field to given value.


### GetWaitForJobs

`func (o *HelmRuntimeConfiguration) GetWaitForJobs() bool`

GetWaitForJobs returns the WaitForJobs field if non-nil, zero value otherwise.

### GetWaitForJobsOk

`func (o *HelmRuntimeConfiguration) GetWaitForJobsOk() (*bool, bool)`

GetWaitForJobsOk returns a tuple with the WaitForJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForJobs

`func (o *HelmRuntimeConfiguration) SetWaitForJobs(v bool)`

SetWaitForJobs sets WaitForJobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


