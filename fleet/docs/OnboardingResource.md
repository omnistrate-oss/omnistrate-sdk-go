# OnboardingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactIDs** | Pointer to **[]string** | The artifact IDs associated with this resource. | [optional] 
**ExecutionOrder** | Pointer to **int64** | The execution order of the resource. | [optional] 
**HelmChartConfiguration** | Pointer to [**OnboardingHelmChartConfiguration**](OnboardingHelmChartConfiguration.md) |  | [optional] 
**InputVariables** | Pointer to [**[]OnboardingResourceInputVariable**](OnboardingResourceInputVariable.md) | Input variables for this resource. | [optional] 
**KustomizeConfiguration** | Pointer to [**OnboardingKustomizeConfiguration**](OnboardingKustomizeConfiguration.md) |  | [optional] 
**Name** | **string** | The resource name. | 
**OperatorCRDConfiguration** | Pointer to [**OnboardingOperatorCRDConfiguration**](OnboardingOperatorCRDConfiguration.md) |  | [optional] 
**OutputVariables** | Pointer to [**[]OnboardingResourceOutputVariable**](OnboardingResourceOutputVariable.md) | Output variables for this resource. | [optional] 
**TerraformConfigurations** | Pointer to [**OnboardingTerraformConfigurations**](OnboardingTerraformConfigurations.md) |  | [optional] 
**TerraformExecutionPolicies** | Pointer to [**map[string]OnboardingTerraformExecutionPolicy**](OnboardingTerraformExecutionPolicy.md) | Terraform execution policy metadata per cloud provider. | [optional] 
**Type** | **string** | The resource type. | 

## Methods

### NewOnboardingResource

`func NewOnboardingResource(name string, type_ string, ) *OnboardingResource`

NewOnboardingResource instantiates a new OnboardingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingResourceWithDefaults

`func NewOnboardingResourceWithDefaults() *OnboardingResource`

NewOnboardingResourceWithDefaults instantiates a new OnboardingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactIDs

`func (o *OnboardingResource) GetArtifactIDs() []string`

GetArtifactIDs returns the ArtifactIDs field if non-nil, zero value otherwise.

### GetArtifactIDsOk

`func (o *OnboardingResource) GetArtifactIDsOk() (*[]string, bool)`

GetArtifactIDsOk returns a tuple with the ArtifactIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactIDs

`func (o *OnboardingResource) SetArtifactIDs(v []string)`

SetArtifactIDs sets ArtifactIDs field to given value.

### HasArtifactIDs

`func (o *OnboardingResource) HasArtifactIDs() bool`

HasArtifactIDs returns a boolean if a field has been set.

### GetExecutionOrder

`func (o *OnboardingResource) GetExecutionOrder() int64`

GetExecutionOrder returns the ExecutionOrder field if non-nil, zero value otherwise.

### GetExecutionOrderOk

`func (o *OnboardingResource) GetExecutionOrderOk() (*int64, bool)`

GetExecutionOrderOk returns a tuple with the ExecutionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOrder

`func (o *OnboardingResource) SetExecutionOrder(v int64)`

SetExecutionOrder sets ExecutionOrder field to given value.

### HasExecutionOrder

`func (o *OnboardingResource) HasExecutionOrder() bool`

HasExecutionOrder returns a boolean if a field has been set.

### GetHelmChartConfiguration

`func (o *OnboardingResource) GetHelmChartConfiguration() OnboardingHelmChartConfiguration`

GetHelmChartConfiguration returns the HelmChartConfiguration field if non-nil, zero value otherwise.

### GetHelmChartConfigurationOk

`func (o *OnboardingResource) GetHelmChartConfigurationOk() (*OnboardingHelmChartConfiguration, bool)`

GetHelmChartConfigurationOk returns a tuple with the HelmChartConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmChartConfiguration

`func (o *OnboardingResource) SetHelmChartConfiguration(v OnboardingHelmChartConfiguration)`

SetHelmChartConfiguration sets HelmChartConfiguration field to given value.

### HasHelmChartConfiguration

`func (o *OnboardingResource) HasHelmChartConfiguration() bool`

HasHelmChartConfiguration returns a boolean if a field has been set.

### GetInputVariables

`func (o *OnboardingResource) GetInputVariables() []OnboardingResourceInputVariable`

GetInputVariables returns the InputVariables field if non-nil, zero value otherwise.

### GetInputVariablesOk

`func (o *OnboardingResource) GetInputVariablesOk() (*[]OnboardingResourceInputVariable, bool)`

GetInputVariablesOk returns a tuple with the InputVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVariables

`func (o *OnboardingResource) SetInputVariables(v []OnboardingResourceInputVariable)`

SetInputVariables sets InputVariables field to given value.

### HasInputVariables

`func (o *OnboardingResource) HasInputVariables() bool`

HasInputVariables returns a boolean if a field has been set.

### GetKustomizeConfiguration

`func (o *OnboardingResource) GetKustomizeConfiguration() OnboardingKustomizeConfiguration`

GetKustomizeConfiguration returns the KustomizeConfiguration field if non-nil, zero value otherwise.

### GetKustomizeConfigurationOk

`func (o *OnboardingResource) GetKustomizeConfigurationOk() (*OnboardingKustomizeConfiguration, bool)`

GetKustomizeConfigurationOk returns a tuple with the KustomizeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizeConfiguration

`func (o *OnboardingResource) SetKustomizeConfiguration(v OnboardingKustomizeConfiguration)`

SetKustomizeConfiguration sets KustomizeConfiguration field to given value.

### HasKustomizeConfiguration

`func (o *OnboardingResource) HasKustomizeConfiguration() bool`

HasKustomizeConfiguration returns a boolean if a field has been set.

### GetName

`func (o *OnboardingResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnboardingResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnboardingResource) SetName(v string)`

SetName sets Name field to given value.


### GetOperatorCRDConfiguration

`func (o *OnboardingResource) GetOperatorCRDConfiguration() OnboardingOperatorCRDConfiguration`

GetOperatorCRDConfiguration returns the OperatorCRDConfiguration field if non-nil, zero value otherwise.

### GetOperatorCRDConfigurationOk

`func (o *OnboardingResource) GetOperatorCRDConfigurationOk() (*OnboardingOperatorCRDConfiguration, bool)`

GetOperatorCRDConfigurationOk returns a tuple with the OperatorCRDConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorCRDConfiguration

`func (o *OnboardingResource) SetOperatorCRDConfiguration(v OnboardingOperatorCRDConfiguration)`

SetOperatorCRDConfiguration sets OperatorCRDConfiguration field to given value.

### HasOperatorCRDConfiguration

`func (o *OnboardingResource) HasOperatorCRDConfiguration() bool`

HasOperatorCRDConfiguration returns a boolean if a field has been set.

### GetOutputVariables

`func (o *OnboardingResource) GetOutputVariables() []OnboardingResourceOutputVariable`

GetOutputVariables returns the OutputVariables field if non-nil, zero value otherwise.

### GetOutputVariablesOk

`func (o *OnboardingResource) GetOutputVariablesOk() (*[]OnboardingResourceOutputVariable, bool)`

GetOutputVariablesOk returns a tuple with the OutputVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputVariables

`func (o *OnboardingResource) SetOutputVariables(v []OnboardingResourceOutputVariable)`

SetOutputVariables sets OutputVariables field to given value.

### HasOutputVariables

`func (o *OnboardingResource) HasOutputVariables() bool`

HasOutputVariables returns a boolean if a field has been set.

### GetTerraformConfigurations

`func (o *OnboardingResource) GetTerraformConfigurations() OnboardingTerraformConfigurations`

GetTerraformConfigurations returns the TerraformConfigurations field if non-nil, zero value otherwise.

### GetTerraformConfigurationsOk

`func (o *OnboardingResource) GetTerraformConfigurationsOk() (*OnboardingTerraformConfigurations, bool)`

GetTerraformConfigurationsOk returns a tuple with the TerraformConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformConfigurations

`func (o *OnboardingResource) SetTerraformConfigurations(v OnboardingTerraformConfigurations)`

SetTerraformConfigurations sets TerraformConfigurations field to given value.

### HasTerraformConfigurations

`func (o *OnboardingResource) HasTerraformConfigurations() bool`

HasTerraformConfigurations returns a boolean if a field has been set.

### GetTerraformExecutionPolicies

`func (o *OnboardingResource) GetTerraformExecutionPolicies() map[string]OnboardingTerraformExecutionPolicy`

GetTerraformExecutionPolicies returns the TerraformExecutionPolicies field if non-nil, zero value otherwise.

### GetTerraformExecutionPoliciesOk

`func (o *OnboardingResource) GetTerraformExecutionPoliciesOk() (*map[string]OnboardingTerraformExecutionPolicy, bool)`

GetTerraformExecutionPoliciesOk returns a tuple with the TerraformExecutionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformExecutionPolicies

`func (o *OnboardingResource) SetTerraformExecutionPolicies(v map[string]OnboardingTerraformExecutionPolicy)`

SetTerraformExecutionPolicies sets TerraformExecutionPolicies field to given value.

### HasTerraformExecutionPolicies

`func (o *OnboardingResource) HasTerraformExecutionPolicies() bool`

HasTerraformExecutionPolicies returns a boolean if a field has been set.

### GetType

`func (o *OnboardingResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnboardingResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnboardingResource) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


