# OnboardingHelmRuntimeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableHooks** | Pointer to **bool** | Disable Helm hooks. | [optional] 
**DisableReconciliation** | Pointer to **bool** | Flag to disable drift reconciliation for the Helm package. | [optional] 
**Recreate** | Pointer to **bool** | Recreate the Helm package if it already exists. | [optional] 
**ResetThenReuseValues** | Pointer to **bool** | Reset then reuse values before applying. | [optional] 
**ResetValues** | Pointer to **bool** | Reset values before applying. | [optional] 
**ReuseValues** | Pointer to **bool** | Reuse values before applying. | [optional] 
**SkipCRDs** | Pointer to **bool** | Skip CRDs for the Helm package. | [optional] 
**TimeoutNanos** | Pointer to **int64** | Timeout (nanos) for the Helm package to be deployed. | [optional] 
**UpgradeCRDs** | Pointer to **bool** | Upgrade CRDs for the Helm package. | [optional] 
**Wait** | Pointer to **bool** | Wait for the Helm package to be deployed. | [optional] 
**WaitForJobs** | Pointer to **bool** | Wait for all jobs to be completed. | [optional] 

## Methods

### NewOnboardingHelmRuntimeConfiguration

`func NewOnboardingHelmRuntimeConfiguration() *OnboardingHelmRuntimeConfiguration`

NewOnboardingHelmRuntimeConfiguration instantiates a new OnboardingHelmRuntimeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingHelmRuntimeConfigurationWithDefaults

`func NewOnboardingHelmRuntimeConfigurationWithDefaults() *OnboardingHelmRuntimeConfiguration`

NewOnboardingHelmRuntimeConfigurationWithDefaults instantiates a new OnboardingHelmRuntimeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableHooks

`func (o *OnboardingHelmRuntimeConfiguration) GetDisableHooks() bool`

GetDisableHooks returns the DisableHooks field if non-nil, zero value otherwise.

### GetDisableHooksOk

`func (o *OnboardingHelmRuntimeConfiguration) GetDisableHooksOk() (*bool, bool)`

GetDisableHooksOk returns a tuple with the DisableHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHooks

`func (o *OnboardingHelmRuntimeConfiguration) SetDisableHooks(v bool)`

SetDisableHooks sets DisableHooks field to given value.

### HasDisableHooks

`func (o *OnboardingHelmRuntimeConfiguration) HasDisableHooks() bool`

HasDisableHooks returns a boolean if a field has been set.

### GetDisableReconciliation

`func (o *OnboardingHelmRuntimeConfiguration) GetDisableReconciliation() bool`

GetDisableReconciliation returns the DisableReconciliation field if non-nil, zero value otherwise.

### GetDisableReconciliationOk

`func (o *OnboardingHelmRuntimeConfiguration) GetDisableReconciliationOk() (*bool, bool)`

GetDisableReconciliationOk returns a tuple with the DisableReconciliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReconciliation

`func (o *OnboardingHelmRuntimeConfiguration) SetDisableReconciliation(v bool)`

SetDisableReconciliation sets DisableReconciliation field to given value.

### HasDisableReconciliation

`func (o *OnboardingHelmRuntimeConfiguration) HasDisableReconciliation() bool`

HasDisableReconciliation returns a boolean if a field has been set.

### GetRecreate

`func (o *OnboardingHelmRuntimeConfiguration) GetRecreate() bool`

GetRecreate returns the Recreate field if non-nil, zero value otherwise.

### GetRecreateOk

`func (o *OnboardingHelmRuntimeConfiguration) GetRecreateOk() (*bool, bool)`

GetRecreateOk returns a tuple with the Recreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecreate

`func (o *OnboardingHelmRuntimeConfiguration) SetRecreate(v bool)`

SetRecreate sets Recreate field to given value.

### HasRecreate

`func (o *OnboardingHelmRuntimeConfiguration) HasRecreate() bool`

HasRecreate returns a boolean if a field has been set.

### GetResetThenReuseValues

`func (o *OnboardingHelmRuntimeConfiguration) GetResetThenReuseValues() bool`

GetResetThenReuseValues returns the ResetThenReuseValues field if non-nil, zero value otherwise.

### GetResetThenReuseValuesOk

`func (o *OnboardingHelmRuntimeConfiguration) GetResetThenReuseValuesOk() (*bool, bool)`

GetResetThenReuseValuesOk returns a tuple with the ResetThenReuseValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetThenReuseValues

`func (o *OnboardingHelmRuntimeConfiguration) SetResetThenReuseValues(v bool)`

SetResetThenReuseValues sets ResetThenReuseValues field to given value.

### HasResetThenReuseValues

`func (o *OnboardingHelmRuntimeConfiguration) HasResetThenReuseValues() bool`

HasResetThenReuseValues returns a boolean if a field has been set.

### GetResetValues

`func (o *OnboardingHelmRuntimeConfiguration) GetResetValues() bool`

GetResetValues returns the ResetValues field if non-nil, zero value otherwise.

### GetResetValuesOk

`func (o *OnboardingHelmRuntimeConfiguration) GetResetValuesOk() (*bool, bool)`

GetResetValuesOk returns a tuple with the ResetValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetValues

`func (o *OnboardingHelmRuntimeConfiguration) SetResetValues(v bool)`

SetResetValues sets ResetValues field to given value.

### HasResetValues

`func (o *OnboardingHelmRuntimeConfiguration) HasResetValues() bool`

HasResetValues returns a boolean if a field has been set.

### GetReuseValues

`func (o *OnboardingHelmRuntimeConfiguration) GetReuseValues() bool`

GetReuseValues returns the ReuseValues field if non-nil, zero value otherwise.

### GetReuseValuesOk

`func (o *OnboardingHelmRuntimeConfiguration) GetReuseValuesOk() (*bool, bool)`

GetReuseValuesOk returns a tuple with the ReuseValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseValues

`func (o *OnboardingHelmRuntimeConfiguration) SetReuseValues(v bool)`

SetReuseValues sets ReuseValues field to given value.

### HasReuseValues

`func (o *OnboardingHelmRuntimeConfiguration) HasReuseValues() bool`

HasReuseValues returns a boolean if a field has been set.

### GetSkipCRDs

`func (o *OnboardingHelmRuntimeConfiguration) GetSkipCRDs() bool`

GetSkipCRDs returns the SkipCRDs field if non-nil, zero value otherwise.

### GetSkipCRDsOk

`func (o *OnboardingHelmRuntimeConfiguration) GetSkipCRDsOk() (*bool, bool)`

GetSkipCRDsOk returns a tuple with the SkipCRDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCRDs

`func (o *OnboardingHelmRuntimeConfiguration) SetSkipCRDs(v bool)`

SetSkipCRDs sets SkipCRDs field to given value.

### HasSkipCRDs

`func (o *OnboardingHelmRuntimeConfiguration) HasSkipCRDs() bool`

HasSkipCRDs returns a boolean if a field has been set.

### GetTimeoutNanos

`func (o *OnboardingHelmRuntimeConfiguration) GetTimeoutNanos() int64`

GetTimeoutNanos returns the TimeoutNanos field if non-nil, zero value otherwise.

### GetTimeoutNanosOk

`func (o *OnboardingHelmRuntimeConfiguration) GetTimeoutNanosOk() (*int64, bool)`

GetTimeoutNanosOk returns a tuple with the TimeoutNanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutNanos

`func (o *OnboardingHelmRuntimeConfiguration) SetTimeoutNanos(v int64)`

SetTimeoutNanos sets TimeoutNanos field to given value.

### HasTimeoutNanos

`func (o *OnboardingHelmRuntimeConfiguration) HasTimeoutNanos() bool`

HasTimeoutNanos returns a boolean if a field has been set.

### GetUpgradeCRDs

`func (o *OnboardingHelmRuntimeConfiguration) GetUpgradeCRDs() bool`

GetUpgradeCRDs returns the UpgradeCRDs field if non-nil, zero value otherwise.

### GetUpgradeCRDsOk

`func (o *OnboardingHelmRuntimeConfiguration) GetUpgradeCRDsOk() (*bool, bool)`

GetUpgradeCRDsOk returns a tuple with the UpgradeCRDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeCRDs

`func (o *OnboardingHelmRuntimeConfiguration) SetUpgradeCRDs(v bool)`

SetUpgradeCRDs sets UpgradeCRDs field to given value.

### HasUpgradeCRDs

`func (o *OnboardingHelmRuntimeConfiguration) HasUpgradeCRDs() bool`

HasUpgradeCRDs returns a boolean if a field has been set.

### GetWait

`func (o *OnboardingHelmRuntimeConfiguration) GetWait() bool`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *OnboardingHelmRuntimeConfiguration) GetWaitOk() (*bool, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *OnboardingHelmRuntimeConfiguration) SetWait(v bool)`

SetWait sets Wait field to given value.

### HasWait

`func (o *OnboardingHelmRuntimeConfiguration) HasWait() bool`

HasWait returns a boolean if a field has been set.

### GetWaitForJobs

`func (o *OnboardingHelmRuntimeConfiguration) GetWaitForJobs() bool`

GetWaitForJobs returns the WaitForJobs field if non-nil, zero value otherwise.

### GetWaitForJobsOk

`func (o *OnboardingHelmRuntimeConfiguration) GetWaitForJobsOk() (*bool, bool)`

GetWaitForJobsOk returns a tuple with the WaitForJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForJobs

`func (o *OnboardingHelmRuntimeConfiguration) SetWaitForJobs(v bool)`

SetWaitForJobs sets WaitForJobs field to given value.

### HasWaitForJobs

`func (o *OnboardingHelmRuntimeConfiguration) HasWaitForJobs() bool`

HasWaitForJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


