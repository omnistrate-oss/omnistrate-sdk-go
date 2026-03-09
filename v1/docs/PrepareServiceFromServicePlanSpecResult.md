# PrepareServiceFromServicePlanSpecResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactUploadingTasks** | Pointer to [**[]ArtifactUploadingTask**](ArtifactUploadingTask.md) | List of artifact uploading tasks to be performed by the caller. Each task contains the artifact type and the presigned URL to upload the artifact to. | [optional] 
**IsNewProductTierCreated** | Pointer to **bool** | Indicates if a new product tier was created | [optional] 
**ProductTierID** | Pointer to **string** | ID of a Product Tier | [optional] 
**ServiceEnvironmentID** | Pointer to **string** | ID of a Service Environment | [optional] 
**ServiceID** | Pointer to **string** | ID of a Service | [optional] 

## Methods

### NewPrepareServiceFromServicePlanSpecResult

`func NewPrepareServiceFromServicePlanSpecResult() *PrepareServiceFromServicePlanSpecResult`

NewPrepareServiceFromServicePlanSpecResult instantiates a new PrepareServiceFromServicePlanSpecResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareServiceFromServicePlanSpecResultWithDefaults

`func NewPrepareServiceFromServicePlanSpecResultWithDefaults() *PrepareServiceFromServicePlanSpecResult`

NewPrepareServiceFromServicePlanSpecResultWithDefaults instantiates a new PrepareServiceFromServicePlanSpecResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactUploadingTasks

`func (o *PrepareServiceFromServicePlanSpecResult) GetArtifactUploadingTasks() []ArtifactUploadingTask`

GetArtifactUploadingTasks returns the ArtifactUploadingTasks field if non-nil, zero value otherwise.

### GetArtifactUploadingTasksOk

`func (o *PrepareServiceFromServicePlanSpecResult) GetArtifactUploadingTasksOk() (*[]ArtifactUploadingTask, bool)`

GetArtifactUploadingTasksOk returns a tuple with the ArtifactUploadingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactUploadingTasks

`func (o *PrepareServiceFromServicePlanSpecResult) SetArtifactUploadingTasks(v []ArtifactUploadingTask)`

SetArtifactUploadingTasks sets ArtifactUploadingTasks field to given value.

### HasArtifactUploadingTasks

`func (o *PrepareServiceFromServicePlanSpecResult) HasArtifactUploadingTasks() bool`

HasArtifactUploadingTasks returns a boolean if a field has been set.

### GetIsNewProductTierCreated

`func (o *PrepareServiceFromServicePlanSpecResult) GetIsNewProductTierCreated() bool`

GetIsNewProductTierCreated returns the IsNewProductTierCreated field if non-nil, zero value otherwise.

### GetIsNewProductTierCreatedOk

`func (o *PrepareServiceFromServicePlanSpecResult) GetIsNewProductTierCreatedOk() (*bool, bool)`

GetIsNewProductTierCreatedOk returns a tuple with the IsNewProductTierCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewProductTierCreated

`func (o *PrepareServiceFromServicePlanSpecResult) SetIsNewProductTierCreated(v bool)`

SetIsNewProductTierCreated sets IsNewProductTierCreated field to given value.

### HasIsNewProductTierCreated

`func (o *PrepareServiceFromServicePlanSpecResult) HasIsNewProductTierCreated() bool`

HasIsNewProductTierCreated returns a boolean if a field has been set.

### GetProductTierID

`func (o *PrepareServiceFromServicePlanSpecResult) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *PrepareServiceFromServicePlanSpecResult) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *PrepareServiceFromServicePlanSpecResult) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.

### HasProductTierID

`func (o *PrepareServiceFromServicePlanSpecResult) HasProductTierID() bool`

HasProductTierID returns a boolean if a field has been set.

### GetServiceEnvironmentID

`func (o *PrepareServiceFromServicePlanSpecResult) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *PrepareServiceFromServicePlanSpecResult) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *PrepareServiceFromServicePlanSpecResult) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.

### HasServiceEnvironmentID

`func (o *PrepareServiceFromServicePlanSpecResult) HasServiceEnvironmentID() bool`

HasServiceEnvironmentID returns a boolean if a field has been set.

### GetServiceID

`func (o *PrepareServiceFromServicePlanSpecResult) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *PrepareServiceFromServicePlanSpecResult) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *PrepareServiceFromServicePlanSpecResult) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *PrepareServiceFromServicePlanSpecResult) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


