# OperatorCRDConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HelmChartDependencies** | Pointer to [**[]OperatorHelmChartDependency**](OperatorHelmChartDependency.md) | The helm chart dependencies for the CRD (including charts necessary to manage the operator) - Optional | [optional] 
**OutputParameters** | Pointer to **map[string]string** | The output parameters to export to the user from the CRD | [optional] 
**ReadinessConditions** | Pointer to **map[string]interface{}** | The readiness conditions to check for the CRD | [optional] 
**SupplementalFiles** | Pointer to **[]string** | The supplemental files to apply with the CRD | [optional] 
**Template** | **string** | The template of the CRD to apply on every deployment | 

## Methods

### NewOperatorCRDConfiguration

`func NewOperatorCRDConfiguration(template string, ) *OperatorCRDConfiguration`

NewOperatorCRDConfiguration instantiates a new OperatorCRDConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorCRDConfigurationWithDefaults

`func NewOperatorCRDConfigurationWithDefaults() *OperatorCRDConfiguration`

NewOperatorCRDConfigurationWithDefaults instantiates a new OperatorCRDConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelmChartDependencies

`func (o *OperatorCRDConfiguration) GetHelmChartDependencies() []OperatorHelmChartDependency`

GetHelmChartDependencies returns the HelmChartDependencies field if non-nil, zero value otherwise.

### GetHelmChartDependenciesOk

`func (o *OperatorCRDConfiguration) GetHelmChartDependenciesOk() (*[]OperatorHelmChartDependency, bool)`

GetHelmChartDependenciesOk returns a tuple with the HelmChartDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartDependencies

`func (o *OperatorCRDConfiguration) SetHelmChartDependencies(v []OperatorHelmChartDependency)`

SetHelmChartDependencies sets HelmChartDependencies field to given value.

### HasHelmChartDependencies

`func (o *OperatorCRDConfiguration) HasHelmChartDependencies() bool`

HasHelmChartDependencies returns a boolean if a field has been set.

### GetOutputParameters

`func (o *OperatorCRDConfiguration) GetOutputParameters() map[string]string`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *OperatorCRDConfiguration) GetOutputParametersOk() (*map[string]string, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *OperatorCRDConfiguration) SetOutputParameters(v map[string]string)`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *OperatorCRDConfiguration) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### GetReadinessConditions

`func (o *OperatorCRDConfiguration) GetReadinessConditions() map[string]interface{}`

GetReadinessConditions returns the ReadinessConditions field if non-nil, zero value otherwise.

### GetReadinessConditionsOk

`func (o *OperatorCRDConfiguration) GetReadinessConditionsOk() (*map[string]interface{}, bool)`

GetReadinessConditionsOk returns a tuple with the ReadinessConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessConditions

`func (o *OperatorCRDConfiguration) SetReadinessConditions(v map[string]interface{})`

SetReadinessConditions sets ReadinessConditions field to given value.

### HasReadinessConditions

`func (o *OperatorCRDConfiguration) HasReadinessConditions() bool`

HasReadinessConditions returns a boolean if a field has been set.

### GetSupplementalFiles

`func (o *OperatorCRDConfiguration) GetSupplementalFiles() []string`

GetSupplementalFiles returns the SupplementalFiles field if non-nil, zero value otherwise.

### GetSupplementalFilesOk

`func (o *OperatorCRDConfiguration) GetSupplementalFilesOk() (*[]string, bool)`

GetSupplementalFilesOk returns a tuple with the SupplementalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalFiles

`func (o *OperatorCRDConfiguration) SetSupplementalFiles(v []string)`

SetSupplementalFiles sets SupplementalFiles field to given value.

### HasSupplementalFiles

`func (o *OperatorCRDConfiguration) HasSupplementalFiles() bool`

HasSupplementalFiles returns a boolean if a field has been set.

### GetTemplate

`func (o *OperatorCRDConfiguration) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *OperatorCRDConfiguration) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *OperatorCRDConfiguration) SetTemplate(v string)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


