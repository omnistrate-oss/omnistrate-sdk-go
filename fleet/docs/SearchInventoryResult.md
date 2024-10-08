# SearchInventoryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCellResults** | Pointer to [**[]DeploymentCellSearchRecord**](DeploymentCellSearchRecord.md) | The deployment cell search results | [optional] 
**NotificationResults** | Pointer to [**[]NotificationSearchRecord**](NotificationSearchRecord.md) | The notification search results | [optional] 
**ProxyInstanceResults** | Pointer to [**[]ProxyInstanceSearchRecord**](ProxyInstanceSearchRecord.md) | The proxy instance search results | [optional] 
**ResourceInstanceResults** | Pointer to [**[]ResourceInstanceSearchRecord**](ResourceInstanceSearchRecord.md) | The resource instance search results | [optional] 
**ResourceResults** | Pointer to [**[]ResourceSearchRecord**](ResourceSearchRecord.md) | The resource search results | [optional] 
**ServerlessProxyResults** | Pointer to [**[]ServerlessProxySearchRecord**](ServerlessProxySearchRecord.md) | The serverless proxy search results | [optional] 
**ServicePlanResults** | Pointer to [**[]ServicePlanSearchRecord**](ServicePlanSearchRecord.md) | The service plan search results | [optional] 
**ServiceResults** | Pointer to [**[]ServiceSearchRecord**](ServiceSearchRecord.md) | The service search results | [optional] 
**SubscriptionResults** | Pointer to [**[]SubscriptionSearchRecord**](SubscriptionSearchRecord.md) | The subscription search results | [optional] 
**UpgradePathResults** | Pointer to [**[]UpgradePathSearchRecord**](UpgradePathSearchRecord.md) | The upgrade path search results | [optional] 
**UserResults** | Pointer to [**[]UserSearchRecord**](UserSearchRecord.md) | The user search results | [optional] 
**WorkflowResults** | Pointer to [**[]WorkflowSearchRecord**](WorkflowSearchRecord.md) | The workflow search results | [optional] 

## Methods

### NewSearchInventoryResult

`func NewSearchInventoryResult() *SearchInventoryResult`

NewSearchInventoryResult instantiates a new SearchInventoryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchInventoryResultWithDefaults

`func NewSearchInventoryResultWithDefaults() *SearchInventoryResult`

NewSearchInventoryResultWithDefaults instantiates a new SearchInventoryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCellResults

`func (o *SearchInventoryResult) GetDeploymentCellResults() []DeploymentCellSearchRecord`

GetDeploymentCellResults returns the DeploymentCellResults field if non-nil, zero value otherwise.

### GetDeploymentCellResultsOk

`func (o *SearchInventoryResult) GetDeploymentCellResultsOk() (*[]DeploymentCellSearchRecord, bool)`

GetDeploymentCellResultsOk returns a tuple with the DeploymentCellResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellResults

`func (o *SearchInventoryResult) SetDeploymentCellResults(v []DeploymentCellSearchRecord)`

SetDeploymentCellResults sets DeploymentCellResults field to given value.

### HasDeploymentCellResults

`func (o *SearchInventoryResult) HasDeploymentCellResults() bool`

HasDeploymentCellResults returns a boolean if a field has been set.

### GetNotificationResults

`func (o *SearchInventoryResult) GetNotificationResults() []NotificationSearchRecord`

GetNotificationResults returns the NotificationResults field if non-nil, zero value otherwise.

### GetNotificationResultsOk

`func (o *SearchInventoryResult) GetNotificationResultsOk() (*[]NotificationSearchRecord, bool)`

GetNotificationResultsOk returns a tuple with the NotificationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationResults

`func (o *SearchInventoryResult) SetNotificationResults(v []NotificationSearchRecord)`

SetNotificationResults sets NotificationResults field to given value.

### HasNotificationResults

`func (o *SearchInventoryResult) HasNotificationResults() bool`

HasNotificationResults returns a boolean if a field has been set.

### GetProxyInstanceResults

`func (o *SearchInventoryResult) GetProxyInstanceResults() []ProxyInstanceSearchRecord`

GetProxyInstanceResults returns the ProxyInstanceResults field if non-nil, zero value otherwise.

### GetProxyInstanceResultsOk

`func (o *SearchInventoryResult) GetProxyInstanceResultsOk() (*[]ProxyInstanceSearchRecord, bool)`

GetProxyInstanceResultsOk returns a tuple with the ProxyInstanceResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyInstanceResults

`func (o *SearchInventoryResult) SetProxyInstanceResults(v []ProxyInstanceSearchRecord)`

SetProxyInstanceResults sets ProxyInstanceResults field to given value.

### HasProxyInstanceResults

`func (o *SearchInventoryResult) HasProxyInstanceResults() bool`

HasProxyInstanceResults returns a boolean if a field has been set.

### GetResourceInstanceResults

`func (o *SearchInventoryResult) GetResourceInstanceResults() []ResourceInstanceSearchRecord`

GetResourceInstanceResults returns the ResourceInstanceResults field if non-nil, zero value otherwise.

### GetResourceInstanceResultsOk

`func (o *SearchInventoryResult) GetResourceInstanceResultsOk() (*[]ResourceInstanceSearchRecord, bool)`

GetResourceInstanceResultsOk returns a tuple with the ResourceInstanceResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInstanceResults

`func (o *SearchInventoryResult) SetResourceInstanceResults(v []ResourceInstanceSearchRecord)`

SetResourceInstanceResults sets ResourceInstanceResults field to given value.

### HasResourceInstanceResults

`func (o *SearchInventoryResult) HasResourceInstanceResults() bool`

HasResourceInstanceResults returns a boolean if a field has been set.

### GetResourceResults

`func (o *SearchInventoryResult) GetResourceResults() []ResourceSearchRecord`

GetResourceResults returns the ResourceResults field if non-nil, zero value otherwise.

### GetResourceResultsOk

`func (o *SearchInventoryResult) GetResourceResultsOk() (*[]ResourceSearchRecord, bool)`

GetResourceResultsOk returns a tuple with the ResourceResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceResults

`func (o *SearchInventoryResult) SetResourceResults(v []ResourceSearchRecord)`

SetResourceResults sets ResourceResults field to given value.

### HasResourceResults

`func (o *SearchInventoryResult) HasResourceResults() bool`

HasResourceResults returns a boolean if a field has been set.

### GetServerlessProxyResults

`func (o *SearchInventoryResult) GetServerlessProxyResults() []ServerlessProxySearchRecord`

GetServerlessProxyResults returns the ServerlessProxyResults field if non-nil, zero value otherwise.

### GetServerlessProxyResultsOk

`func (o *SearchInventoryResult) GetServerlessProxyResultsOk() (*[]ServerlessProxySearchRecord, bool)`

GetServerlessProxyResultsOk returns a tuple with the ServerlessProxyResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessProxyResults

`func (o *SearchInventoryResult) SetServerlessProxyResults(v []ServerlessProxySearchRecord)`

SetServerlessProxyResults sets ServerlessProxyResults field to given value.

### HasServerlessProxyResults

`func (o *SearchInventoryResult) HasServerlessProxyResults() bool`

HasServerlessProxyResults returns a boolean if a field has been set.

### GetServicePlanResults

`func (o *SearchInventoryResult) GetServicePlanResults() []ServicePlanSearchRecord`

GetServicePlanResults returns the ServicePlanResults field if non-nil, zero value otherwise.

### GetServicePlanResultsOk

`func (o *SearchInventoryResult) GetServicePlanResultsOk() (*[]ServicePlanSearchRecord, bool)`

GetServicePlanResultsOk returns a tuple with the ServicePlanResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanResults

`func (o *SearchInventoryResult) SetServicePlanResults(v []ServicePlanSearchRecord)`

SetServicePlanResults sets ServicePlanResults field to given value.

### HasServicePlanResults

`func (o *SearchInventoryResult) HasServicePlanResults() bool`

HasServicePlanResults returns a boolean if a field has been set.

### GetServiceResults

`func (o *SearchInventoryResult) GetServiceResults() []ServiceSearchRecord`

GetServiceResults returns the ServiceResults field if non-nil, zero value otherwise.

### GetServiceResultsOk

`func (o *SearchInventoryResult) GetServiceResultsOk() (*[]ServiceSearchRecord, bool)`

GetServiceResultsOk returns a tuple with the ServiceResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceResults

`func (o *SearchInventoryResult) SetServiceResults(v []ServiceSearchRecord)`

SetServiceResults sets ServiceResults field to given value.

### HasServiceResults

`func (o *SearchInventoryResult) HasServiceResults() bool`

HasServiceResults returns a boolean if a field has been set.

### GetSubscriptionResults

`func (o *SearchInventoryResult) GetSubscriptionResults() []SubscriptionSearchRecord`

GetSubscriptionResults returns the SubscriptionResults field if non-nil, zero value otherwise.

### GetSubscriptionResultsOk

`func (o *SearchInventoryResult) GetSubscriptionResultsOk() (*[]SubscriptionSearchRecord, bool)`

GetSubscriptionResultsOk returns a tuple with the SubscriptionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionResults

`func (o *SearchInventoryResult) SetSubscriptionResults(v []SubscriptionSearchRecord)`

SetSubscriptionResults sets SubscriptionResults field to given value.

### HasSubscriptionResults

`func (o *SearchInventoryResult) HasSubscriptionResults() bool`

HasSubscriptionResults returns a boolean if a field has been set.

### GetUpgradePathResults

`func (o *SearchInventoryResult) GetUpgradePathResults() []UpgradePathSearchRecord`

GetUpgradePathResults returns the UpgradePathResults field if non-nil, zero value otherwise.

### GetUpgradePathResultsOk

`func (o *SearchInventoryResult) GetUpgradePathResultsOk() (*[]UpgradePathSearchRecord, bool)`

GetUpgradePathResultsOk returns a tuple with the UpgradePathResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePathResults

`func (o *SearchInventoryResult) SetUpgradePathResults(v []UpgradePathSearchRecord)`

SetUpgradePathResults sets UpgradePathResults field to given value.

### HasUpgradePathResults

`func (o *SearchInventoryResult) HasUpgradePathResults() bool`

HasUpgradePathResults returns a boolean if a field has been set.

### GetUserResults

`func (o *SearchInventoryResult) GetUserResults() []UserSearchRecord`

GetUserResults returns the UserResults field if non-nil, zero value otherwise.

### GetUserResultsOk

`func (o *SearchInventoryResult) GetUserResultsOk() (*[]UserSearchRecord, bool)`

GetUserResultsOk returns a tuple with the UserResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResults

`func (o *SearchInventoryResult) SetUserResults(v []UserSearchRecord)`

SetUserResults sets UserResults field to given value.

### HasUserResults

`func (o *SearchInventoryResult) HasUserResults() bool`

HasUserResults returns a boolean if a field has been set.

### GetWorkflowResults

`func (o *SearchInventoryResult) GetWorkflowResults() []WorkflowSearchRecord`

GetWorkflowResults returns the WorkflowResults field if non-nil, zero value otherwise.

### GetWorkflowResultsOk

`func (o *SearchInventoryResult) GetWorkflowResultsOk() (*[]WorkflowSearchRecord, bool)`

GetWorkflowResultsOk returns a tuple with the WorkflowResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowResults

`func (o *SearchInventoryResult) SetWorkflowResults(v []WorkflowSearchRecord)`

SetWorkflowResults sets WorkflowResults field to given value.

### HasWorkflowResults

`func (o *SearchInventoryResult) HasWorkflowResults() bool`

HasWorkflowResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


