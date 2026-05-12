# OnboardingOperatorCRDConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalEndpoints** | Pointer to [**map[string]OnboardingEndpoint**](OnboardingEndpoint.md) | Endpoints from the Operator CRD deployment to expose. | [optional] 
**DisableReconciliation** | Pointer to **bool** | Flag to disable drift reconciliation. | [optional] 
**HelmChartDependencies** | Pointer to [**[]OnboardingChartDependency**](OnboardingChartDependency.md) | Helm chart dependencies for the operator. | [optional] 
**OutputParameters** | Pointer to **map[string]string** | Output parameters to export from the CRD. | [optional] 
**ReadinessCondition** | Pointer to **map[string]interface{}** | Readiness conditions to check for the CRD. | [optional] 
**SupplementalFiles** | Pointer to **[]string** | Supplemental files to apply with the CRD. | [optional] 
**Template** | **string** | The template of the CRD to apply on every deployment. | 

## Methods

### NewOnboardingOperatorCRDConfiguration

`func NewOnboardingOperatorCRDConfiguration(template string, ) *OnboardingOperatorCRDConfiguration`

NewOnboardingOperatorCRDConfiguration instantiates a new OnboardingOperatorCRDConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingOperatorCRDConfigurationWithDefaults

`func NewOnboardingOperatorCRDConfigurationWithDefaults() *OnboardingOperatorCRDConfiguration`

NewOnboardingOperatorCRDConfigurationWithDefaults instantiates a new OnboardingOperatorCRDConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalEndpoints

`func (o *OnboardingOperatorCRDConfiguration) GetAdditionalEndpoints() map[string]OnboardingEndpoint`

GetAdditionalEndpoints returns the AdditionalEndpoints field if non-nil, zero value otherwise.

### GetAdditionalEndpointsOk

`func (o *OnboardingOperatorCRDConfiguration) GetAdditionalEndpointsOk() (*map[string]OnboardingEndpoint, bool)`

GetAdditionalEndpointsOk returns a tuple with the AdditionalEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEndpoints

`func (o *OnboardingOperatorCRDConfiguration) SetAdditionalEndpoints(v map[string]OnboardingEndpoint)`

SetAdditionalEndpoints sets AdditionalEndpoints field to given value.

### HasAdditionalEndpoints

`func (o *OnboardingOperatorCRDConfiguration) HasAdditionalEndpoints() bool`

HasAdditionalEndpoints returns a boolean if a field has been set.

### GetDisableReconciliation

`func (o *OnboardingOperatorCRDConfiguration) GetDisableReconciliation() bool`

GetDisableReconciliation returns the DisableReconciliation field if non-nil, zero value otherwise.

### GetDisableReconciliationOk

`func (o *OnboardingOperatorCRDConfiguration) GetDisableReconciliationOk() (*bool, bool)`

GetDisableReconciliationOk returns a tuple with the DisableReconciliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReconciliation

`func (o *OnboardingOperatorCRDConfiguration) SetDisableReconciliation(v bool)`

SetDisableReconciliation sets DisableReconciliation field to given value.

### HasDisableReconciliation

`func (o *OnboardingOperatorCRDConfiguration) HasDisableReconciliation() bool`

HasDisableReconciliation returns a boolean if a field has been set.

### GetHelmChartDependencies

`func (o *OnboardingOperatorCRDConfiguration) GetHelmChartDependencies() []OnboardingChartDependency`

GetHelmChartDependencies returns the HelmChartDependencies field if non-nil, zero value otherwise.

### GetHelmChartDependenciesOk

`func (o *OnboardingOperatorCRDConfiguration) GetHelmChartDependenciesOk() (*[]OnboardingChartDependency, bool)`

GetHelmChartDependenciesOk returns a tuple with the HelmChartDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartDependencies

`func (o *OnboardingOperatorCRDConfiguration) SetHelmChartDependencies(v []OnboardingChartDependency)`

SetHelmChartDependencies sets HelmChartDependencies field to given value.

### HasHelmChartDependencies

`func (o *OnboardingOperatorCRDConfiguration) HasHelmChartDependencies() bool`

HasHelmChartDependencies returns a boolean if a field has been set.

### GetOutputParameters

`func (o *OnboardingOperatorCRDConfiguration) GetOutputParameters() map[string]string`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *OnboardingOperatorCRDConfiguration) GetOutputParametersOk() (*map[string]string, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *OnboardingOperatorCRDConfiguration) SetOutputParameters(v map[string]string)`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *OnboardingOperatorCRDConfiguration) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### GetReadinessCondition

`func (o *OnboardingOperatorCRDConfiguration) GetReadinessCondition() map[string]interface{}`

GetReadinessCondition returns the ReadinessCondition field if non-nil, zero value otherwise.

### GetReadinessConditionOk

`func (o *OnboardingOperatorCRDConfiguration) GetReadinessConditionOk() (*map[string]interface{}, bool)`

GetReadinessConditionOk returns a tuple with the ReadinessCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessCondition

`func (o *OnboardingOperatorCRDConfiguration) SetReadinessCondition(v map[string]interface{})`

SetReadinessCondition sets ReadinessCondition field to given value.

### HasReadinessCondition

`func (o *OnboardingOperatorCRDConfiguration) HasReadinessCondition() bool`

HasReadinessCondition returns a boolean if a field has been set.

### GetSupplementalFiles

`func (o *OnboardingOperatorCRDConfiguration) GetSupplementalFiles() []string`

GetSupplementalFiles returns the SupplementalFiles field if non-nil, zero value otherwise.

### GetSupplementalFilesOk

`func (o *OnboardingOperatorCRDConfiguration) GetSupplementalFilesOk() (*[]string, bool)`

GetSupplementalFilesOk returns a tuple with the SupplementalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalFiles

`func (o *OnboardingOperatorCRDConfiguration) SetSupplementalFiles(v []string)`

SetSupplementalFiles sets SupplementalFiles field to given value.

### HasSupplementalFiles

`func (o *OnboardingOperatorCRDConfiguration) HasSupplementalFiles() bool`

HasSupplementalFiles returns a boolean if a field has been set.

### GetTemplate

`func (o *OnboardingOperatorCRDConfiguration) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *OnboardingOperatorCRDConfiguration) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *OnboardingOperatorCRDConfiguration) SetTemplate(v string)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


