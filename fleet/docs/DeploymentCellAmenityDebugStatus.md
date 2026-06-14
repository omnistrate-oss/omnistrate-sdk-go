# DeploymentCellAmenityDebugStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredStatus** | **string** | The desired state of the amenity on this deployment cell | 
**Generation** | **int64** | The amenity desired-state generation | 
**IsManaged** | Pointer to **bool** | Whether the amenity is managed by Omnistrate | [optional] 
**LastError** | Pointer to **string** | The last apply error for this amenity | [optional] 
**Name** | **string** | The amenity name | 
**Source** | Pointer to **string** | The source that last wrote the desired state | [optional] 
**SourceCloudProviderId** | Pointer to **string** | The cloud provider template used as source | [optional] 
**SourceEnvironmentType** | Pointer to **string** | The environment template used as source | [optional] 
**Type** | **string** | The amenity type | 
**WorkflowId** | Pointer to **string** | The workflow ID applying or last applying this amenity | [optional] 
**WorkflowRunId** | Pointer to **string** | The workflow run ID applying or last applying this amenity | [optional] 

## Methods

### NewDeploymentCellAmenityDebugStatus

`func NewDeploymentCellAmenityDebugStatus(desiredStatus string, generation int64, name string, type_ string, ) *DeploymentCellAmenityDebugStatus`

NewDeploymentCellAmenityDebugStatus instantiates a new DeploymentCellAmenityDebugStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellAmenityDebugStatusWithDefaults

`func NewDeploymentCellAmenityDebugStatusWithDefaults() *DeploymentCellAmenityDebugStatus`

NewDeploymentCellAmenityDebugStatusWithDefaults instantiates a new DeploymentCellAmenityDebugStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredStatus

`func (o *DeploymentCellAmenityDebugStatus) GetDesiredStatus() string`

GetDesiredStatus returns the DesiredStatus field if non-nil, zero value otherwise.

### GetDesiredStatusOk

`func (o *DeploymentCellAmenityDebugStatus) GetDesiredStatusOk() (*string, bool)`

GetDesiredStatusOk returns a tuple with the DesiredStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredStatus

`func (o *DeploymentCellAmenityDebugStatus) SetDesiredStatus(v string)`

SetDesiredStatus sets DesiredStatus field to given value.


### GetGeneration

`func (o *DeploymentCellAmenityDebugStatus) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *DeploymentCellAmenityDebugStatus) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *DeploymentCellAmenityDebugStatus) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.


### GetIsManaged

`func (o *DeploymentCellAmenityDebugStatus) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *DeploymentCellAmenityDebugStatus) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *DeploymentCellAmenityDebugStatus) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *DeploymentCellAmenityDebugStatus) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetLastError

`func (o *DeploymentCellAmenityDebugStatus) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *DeploymentCellAmenityDebugStatus) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *DeploymentCellAmenityDebugStatus) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *DeploymentCellAmenityDebugStatus) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetName

`func (o *DeploymentCellAmenityDebugStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentCellAmenityDebugStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentCellAmenityDebugStatus) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *DeploymentCellAmenityDebugStatus) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeploymentCellAmenityDebugStatus) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeploymentCellAmenityDebugStatus) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DeploymentCellAmenityDebugStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceCloudProviderId

`func (o *DeploymentCellAmenityDebugStatus) GetSourceCloudProviderId() string`

GetSourceCloudProviderId returns the SourceCloudProviderId field if non-nil, zero value otherwise.

### GetSourceCloudProviderIdOk

`func (o *DeploymentCellAmenityDebugStatus) GetSourceCloudProviderIdOk() (*string, bool)`

GetSourceCloudProviderIdOk returns a tuple with the SourceCloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCloudProviderId

`func (o *DeploymentCellAmenityDebugStatus) SetSourceCloudProviderId(v string)`

SetSourceCloudProviderId sets SourceCloudProviderId field to given value.

### HasSourceCloudProviderId

`func (o *DeploymentCellAmenityDebugStatus) HasSourceCloudProviderId() bool`

HasSourceCloudProviderId returns a boolean if a field has been set.

### GetSourceEnvironmentType

`func (o *DeploymentCellAmenityDebugStatus) GetSourceEnvironmentType() string`

GetSourceEnvironmentType returns the SourceEnvironmentType field if non-nil, zero value otherwise.

### GetSourceEnvironmentTypeOk

`func (o *DeploymentCellAmenityDebugStatus) GetSourceEnvironmentTypeOk() (*string, bool)`

GetSourceEnvironmentTypeOk returns a tuple with the SourceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironmentType

`func (o *DeploymentCellAmenityDebugStatus) SetSourceEnvironmentType(v string)`

SetSourceEnvironmentType sets SourceEnvironmentType field to given value.

### HasSourceEnvironmentType

`func (o *DeploymentCellAmenityDebugStatus) HasSourceEnvironmentType() bool`

HasSourceEnvironmentType returns a boolean if a field has been set.

### GetType

`func (o *DeploymentCellAmenityDebugStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentCellAmenityDebugStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentCellAmenityDebugStatus) SetType(v string)`

SetType sets Type field to given value.


### GetWorkflowId

`func (o *DeploymentCellAmenityDebugStatus) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *DeploymentCellAmenityDebugStatus) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *DeploymentCellAmenityDebugStatus) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *DeploymentCellAmenityDebugStatus) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowRunId

`func (o *DeploymentCellAmenityDebugStatus) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *DeploymentCellAmenityDebugStatus) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *DeploymentCellAmenityDebugStatus) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.

### HasWorkflowRunId

`func (o *DeploymentCellAmenityDebugStatus) HasWorkflowRunId() bool`

HasWorkflowRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


