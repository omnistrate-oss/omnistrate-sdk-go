# DebugHostClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmenityArtifacts** | Pointer to [**[]DeploymentCellAmenityArtifactSummary**](DeploymentCellAmenityArtifactSummary.md) | Compressed amenity artifact summaries available for debugging | [optional] 
**AmenityStatuses** | Pointer to [**[]DeploymentCellAmenityDebugStatus**](DeploymentCellAmenityDebugStatus.md) | Amenity desired-state statuses for this deployment cell | [optional] 
**CustomHelmExecutionLogsBase64** | Pointer to **map[string]string** | Custom Helm execution logs for the host cluster, keyed by namespace | [optional] 
**Template** | Pointer to [**DeploymentCellAmenityTemplateDebugInfo**](DeploymentCellAmenityTemplateDebugInfo.md) |  | [optional] 

## Methods

### NewDebugHostClusterResult

`func NewDebugHostClusterResult() *DebugHostClusterResult`

NewDebugHostClusterResult instantiates a new DebugHostClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebugHostClusterResultWithDefaults

`func NewDebugHostClusterResultWithDefaults() *DebugHostClusterResult`

NewDebugHostClusterResultWithDefaults instantiates a new DebugHostClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmenityArtifacts

`func (o *DebugHostClusterResult) GetAmenityArtifacts() []DeploymentCellAmenityArtifactSummary`

GetAmenityArtifacts returns the AmenityArtifacts field if non-nil, zero value otherwise.

### GetAmenityArtifactsOk

`func (o *DebugHostClusterResult) GetAmenityArtifactsOk() (*[]DeploymentCellAmenityArtifactSummary, bool)`

GetAmenityArtifactsOk returns a tuple with the AmenityArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenityArtifacts

`func (o *DebugHostClusterResult) SetAmenityArtifacts(v []DeploymentCellAmenityArtifactSummary)`

SetAmenityArtifacts sets AmenityArtifacts field to given value.

### HasAmenityArtifacts

`func (o *DebugHostClusterResult) HasAmenityArtifacts() bool`

HasAmenityArtifacts returns a boolean if a field has been set.

### GetAmenityStatuses

`func (o *DebugHostClusterResult) GetAmenityStatuses() []DeploymentCellAmenityDebugStatus`

GetAmenityStatuses returns the AmenityStatuses field if non-nil, zero value otherwise.

### GetAmenityStatusesOk

`func (o *DebugHostClusterResult) GetAmenityStatusesOk() (*[]DeploymentCellAmenityDebugStatus, bool)`

GetAmenityStatusesOk returns a tuple with the AmenityStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenityStatuses

`func (o *DebugHostClusterResult) SetAmenityStatuses(v []DeploymentCellAmenityDebugStatus)`

SetAmenityStatuses sets AmenityStatuses field to given value.

### HasAmenityStatuses

`func (o *DebugHostClusterResult) HasAmenityStatuses() bool`

HasAmenityStatuses returns a boolean if a field has been set.

### GetCustomHelmExecutionLogsBase64

`func (o *DebugHostClusterResult) GetCustomHelmExecutionLogsBase64() map[string]string`

GetCustomHelmExecutionLogsBase64 returns the CustomHelmExecutionLogsBase64 field if non-nil, zero value otherwise.

### GetCustomHelmExecutionLogsBase64Ok

`func (o *DebugHostClusterResult) GetCustomHelmExecutionLogsBase64Ok() (*map[string]string, bool)`

GetCustomHelmExecutionLogsBase64Ok returns a tuple with the CustomHelmExecutionLogsBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHelmExecutionLogsBase64

`func (o *DebugHostClusterResult) SetCustomHelmExecutionLogsBase64(v map[string]string)`

SetCustomHelmExecutionLogsBase64 sets CustomHelmExecutionLogsBase64 field to given value.

### HasCustomHelmExecutionLogsBase64

`func (o *DebugHostClusterResult) HasCustomHelmExecutionLogsBase64() bool`

HasCustomHelmExecutionLogsBase64 returns a boolean if a field has been set.

### GetTemplate

`func (o *DebugHostClusterResult) GetTemplate() DeploymentCellAmenityTemplateDebugInfo`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *DebugHostClusterResult) GetTemplateOk() (*DeploymentCellAmenityTemplateDebugInfo, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *DebugHostClusterResult) SetTemplate(v DeploymentCellAmenityTemplateDebugInfo)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *DebugHostClusterResult) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


