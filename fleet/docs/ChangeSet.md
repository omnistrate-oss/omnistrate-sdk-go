# ChangeSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategorizedResourceChanges** | [**map[string]ChangeSummary**](ChangeSummary.md) | Summary of all changes within a resource for each category (ie. image, infra, product tier feature, etc.) | 
**DeploymentArtifactChanges** | Pointer to **string** | State of the configuration change | [optional] 
**ImageConfigChanges** | Pointer to [**ImageConfigChangeSummary**](ImageConfigChangeSummary.md) |  | [optional] 
**InfraConfigChanges** | Pointer to [**InfraConfigChangeSummary**](InfraConfigChangeSummary.md) |  | [optional] 
**OverallResourceStatus** | **string** | Summary status of the changes | 
**ProductTierFeatureChanges** | Pointer to **string** | State of the configuration change | [optional] 
**ResourceChanges** | Pointer to **string** | State of the configuration change | [optional] 
**ResourceName** | Pointer to **string** | The name of the resource | [optional] 

## Methods

### NewChangeSet

`func NewChangeSet(categorizedResourceChanges map[string]ChangeSummary, overallResourceStatus string, ) *ChangeSet`

NewChangeSet instantiates a new ChangeSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeSetWithDefaults

`func NewChangeSetWithDefaults() *ChangeSet`

NewChangeSetWithDefaults instantiates a new ChangeSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategorizedResourceChanges

`func (o *ChangeSet) GetCategorizedResourceChanges() map[string]ChangeSummary`

GetCategorizedResourceChanges returns the CategorizedResourceChanges field if non-nil, zero value otherwise.

### GetCategorizedResourceChangesOk

`func (o *ChangeSet) GetCategorizedResourceChangesOk() (*map[string]ChangeSummary, bool)`

GetCategorizedResourceChangesOk returns a tuple with the CategorizedResourceChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorizedResourceChanges

`func (o *ChangeSet) SetCategorizedResourceChanges(v map[string]ChangeSummary)`

SetCategorizedResourceChanges sets CategorizedResourceChanges field to given value.


### GetDeploymentArtifactChanges

`func (o *ChangeSet) GetDeploymentArtifactChanges() string`

GetDeploymentArtifactChanges returns the DeploymentArtifactChanges field if non-nil, zero value otherwise.

### GetDeploymentArtifactChangesOk

`func (o *ChangeSet) GetDeploymentArtifactChangesOk() (*string, bool)`

GetDeploymentArtifactChangesOk returns a tuple with the DeploymentArtifactChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentArtifactChanges

`func (o *ChangeSet) SetDeploymentArtifactChanges(v string)`

SetDeploymentArtifactChanges sets DeploymentArtifactChanges field to given value.

### HasDeploymentArtifactChanges

`func (o *ChangeSet) HasDeploymentArtifactChanges() bool`

HasDeploymentArtifactChanges returns a boolean if a field has been set.

### GetImageConfigChanges

`func (o *ChangeSet) GetImageConfigChanges() ImageConfigChangeSummary`

GetImageConfigChanges returns the ImageConfigChanges field if non-nil, zero value otherwise.

### GetImageConfigChangesOk

`func (o *ChangeSet) GetImageConfigChangesOk() (*ImageConfigChangeSummary, bool)`

GetImageConfigChangesOk returns a tuple with the ImageConfigChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfigChanges

`func (o *ChangeSet) SetImageConfigChanges(v ImageConfigChangeSummary)`

SetImageConfigChanges sets ImageConfigChanges field to given value.

### HasImageConfigChanges

`func (o *ChangeSet) HasImageConfigChanges() bool`

HasImageConfigChanges returns a boolean if a field has been set.

### GetInfraConfigChanges

`func (o *ChangeSet) GetInfraConfigChanges() InfraConfigChangeSummary`

GetInfraConfigChanges returns the InfraConfigChanges field if non-nil, zero value otherwise.

### GetInfraConfigChangesOk

`func (o *ChangeSet) GetInfraConfigChangesOk() (*InfraConfigChangeSummary, bool)`

GetInfraConfigChangesOk returns a tuple with the InfraConfigChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigChanges

`func (o *ChangeSet) SetInfraConfigChanges(v InfraConfigChangeSummary)`

SetInfraConfigChanges sets InfraConfigChanges field to given value.

### HasInfraConfigChanges

`func (o *ChangeSet) HasInfraConfigChanges() bool`

HasInfraConfigChanges returns a boolean if a field has been set.

### GetOverallResourceStatus

`func (o *ChangeSet) GetOverallResourceStatus() string`

GetOverallResourceStatus returns the OverallResourceStatus field if non-nil, zero value otherwise.

### GetOverallResourceStatusOk

`func (o *ChangeSet) GetOverallResourceStatusOk() (*string, bool)`

GetOverallResourceStatusOk returns a tuple with the OverallResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallResourceStatus

`func (o *ChangeSet) SetOverallResourceStatus(v string)`

SetOverallResourceStatus sets OverallResourceStatus field to given value.


### GetProductTierFeatureChanges

`func (o *ChangeSet) GetProductTierFeatureChanges() string`

GetProductTierFeatureChanges returns the ProductTierFeatureChanges field if non-nil, zero value otherwise.

### GetProductTierFeatureChangesOk

`func (o *ChangeSet) GetProductTierFeatureChangesOk() (*string, bool)`

GetProductTierFeatureChangesOk returns a tuple with the ProductTierFeatureChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierFeatureChanges

`func (o *ChangeSet) SetProductTierFeatureChanges(v string)`

SetProductTierFeatureChanges sets ProductTierFeatureChanges field to given value.

### HasProductTierFeatureChanges

`func (o *ChangeSet) HasProductTierFeatureChanges() bool`

HasProductTierFeatureChanges returns a boolean if a field has been set.

### GetResourceChanges

`func (o *ChangeSet) GetResourceChanges() string`

GetResourceChanges returns the ResourceChanges field if non-nil, zero value otherwise.

### GetResourceChangesOk

`func (o *ChangeSet) GetResourceChangesOk() (*string, bool)`

GetResourceChangesOk returns a tuple with the ResourceChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceChanges

`func (o *ChangeSet) SetResourceChanges(v string)`

SetResourceChanges sets ResourceChanges field to given value.

### HasResourceChanges

`func (o *ChangeSet) HasResourceChanges() bool`

HasResourceChanges returns a boolean if a field has been set.

### GetResourceName

`func (o *ChangeSet) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ChangeSet) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ChangeSet) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ChangeSet) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


