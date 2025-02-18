# DescribeResourceInstanceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | True if this resource instance has associated infrastructure deployed | [optional] 
**AutoscalingEnabled** | Pointer to **bool** | Whether the instance has autoscaling enabled | [optional] 
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**BackupStatus** | Pointer to [**BackupStatus**](BackupStatus.md) |  | [optional] 
**CloudProvider** | Pointer to **string** | The cloud provider name | [optional] 
**CreatedAt** | Pointer to **string** | The instance creation time | [optional] 
**CreatedByUserId** | Pointer to **string** | ID of a User | [optional] 
**CreatedByUserName** | Pointer to **string** | The name of the user that created the resource instance | [optional] 
**CurrentReplicas** | Pointer to **string** | The current number of replicas | [optional] 
**CustomNetworkDetail** | Pointer to [**CustomNetworkResourceDetail**](CustomNetworkResourceDetail.md) |  | [optional] 
**DetailedNetworkTopology** | Pointer to **map[string]interface{}** | The detailed network topology | [optional] 
**ExternalPayerId** | Pointer to **string** | The external payer id to record which customer should pay for this resource instance | [optional] 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**HighAvailability** | Pointer to **bool** | Whether the instance is High Availability | [optional] 
**Id** | Pointer to **string** | The instance ID | [optional] 
**InstanceLoadStatus** | Pointer to **string** | The load status of a pod | [optional] 
**KubernetesDashboardEndpoint** | Pointer to [**KubernetesDashboardEndpoint**](KubernetesDashboardEndpoint.md) |  | [optional] 
**LastModifiedAt** | Pointer to **string** | The instance update time | [optional] 
**MaxReplicas** | Pointer to **string** | The maximum number of replicas | [optional] 
**MinReplicas** | Pointer to **string** | The minimum number of replicas | [optional] 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**ProductTierFeatures** | Pointer to **map[string]interface{}** | The product tier features | [optional] 
**Region** | Pointer to **string** | The region code | [optional] 
**ResourceID** | Pointer to **string** | ID of a resource | [optional] 
**ResultParams** | Pointer to **interface{}** | Custom result parameters | [optional] 
**ServerlessEnabled** | Pointer to **bool** | Whether the instance has serverless enabled | [optional] 
**Status** | Pointer to **string** | The status of an operation | [optional] 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 

## Methods

### NewDescribeResourceInstanceResult

`func NewDescribeResourceInstanceResult() *DescribeResourceInstanceResult`

NewDescribeResourceInstanceResult instantiates a new DescribeResourceInstanceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceInstanceResultWithDefaults

`func NewDescribeResourceInstanceResultWithDefaults() *DescribeResourceInstanceResult`

NewDescribeResourceInstanceResultWithDefaults instantiates a new DescribeResourceInstanceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DescribeResourceInstanceResult) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DescribeResourceInstanceResult) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DescribeResourceInstanceResult) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DescribeResourceInstanceResult) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAutoscalingEnabled

`func (o *DescribeResourceInstanceResult) GetAutoscalingEnabled() bool`

GetAutoscalingEnabled returns the AutoscalingEnabled field if non-nil, zero value otherwise.

### GetAutoscalingEnabledOk

`func (o *DescribeResourceInstanceResult) GetAutoscalingEnabledOk() (*bool, bool)`

GetAutoscalingEnabledOk returns a tuple with the AutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingEnabled

`func (o *DescribeResourceInstanceResult) SetAutoscalingEnabled(v bool)`

SetAutoscalingEnabled sets AutoscalingEnabled field to given value.

### HasAutoscalingEnabled

`func (o *DescribeResourceInstanceResult) HasAutoscalingEnabled() bool`

HasAutoscalingEnabled returns a boolean if a field has been set.

### GetAwsAccountID

`func (o *DescribeResourceInstanceResult) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *DescribeResourceInstanceResult) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *DescribeResourceInstanceResult) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *DescribeResourceInstanceResult) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetBackupStatus

`func (o *DescribeResourceInstanceResult) GetBackupStatus() BackupStatus`

GetBackupStatus returns the BackupStatus field if non-nil, zero value otherwise.

### GetBackupStatusOk

`func (o *DescribeResourceInstanceResult) GetBackupStatusOk() (*BackupStatus, bool)`

GetBackupStatusOk returns a tuple with the BackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStatus

`func (o *DescribeResourceInstanceResult) SetBackupStatus(v BackupStatus)`

SetBackupStatus sets BackupStatus field to given value.

### HasBackupStatus

`func (o *DescribeResourceInstanceResult) HasBackupStatus() bool`

HasBackupStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *DescribeResourceInstanceResult) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DescribeResourceInstanceResult) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DescribeResourceInstanceResult) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DescribeResourceInstanceResult) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DescribeResourceInstanceResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DescribeResourceInstanceResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DescribeResourceInstanceResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DescribeResourceInstanceResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *DescribeResourceInstanceResult) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *DescribeResourceInstanceResult) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *DescribeResourceInstanceResult) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *DescribeResourceInstanceResult) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedByUserName

`func (o *DescribeResourceInstanceResult) GetCreatedByUserName() string`

GetCreatedByUserName returns the CreatedByUserName field if non-nil, zero value otherwise.

### GetCreatedByUserNameOk

`func (o *DescribeResourceInstanceResult) GetCreatedByUserNameOk() (*string, bool)`

GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserName

`func (o *DescribeResourceInstanceResult) SetCreatedByUserName(v string)`

SetCreatedByUserName sets CreatedByUserName field to given value.

### HasCreatedByUserName

`func (o *DescribeResourceInstanceResult) HasCreatedByUserName() bool`

HasCreatedByUserName returns a boolean if a field has been set.

### GetCurrentReplicas

`func (o *DescribeResourceInstanceResult) GetCurrentReplicas() string`

GetCurrentReplicas returns the CurrentReplicas field if non-nil, zero value otherwise.

### GetCurrentReplicasOk

`func (o *DescribeResourceInstanceResult) GetCurrentReplicasOk() (*string, bool)`

GetCurrentReplicasOk returns a tuple with the CurrentReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReplicas

`func (o *DescribeResourceInstanceResult) SetCurrentReplicas(v string)`

SetCurrentReplicas sets CurrentReplicas field to given value.

### HasCurrentReplicas

`func (o *DescribeResourceInstanceResult) HasCurrentReplicas() bool`

HasCurrentReplicas returns a boolean if a field has been set.

### GetCustomNetworkDetail

`func (o *DescribeResourceInstanceResult) GetCustomNetworkDetail() CustomNetworkResourceDetail`

GetCustomNetworkDetail returns the CustomNetworkDetail field if non-nil, zero value otherwise.

### GetCustomNetworkDetailOk

`func (o *DescribeResourceInstanceResult) GetCustomNetworkDetailOk() (*CustomNetworkResourceDetail, bool)`

GetCustomNetworkDetailOk returns a tuple with the CustomNetworkDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworkDetail

`func (o *DescribeResourceInstanceResult) SetCustomNetworkDetail(v CustomNetworkResourceDetail)`

SetCustomNetworkDetail sets CustomNetworkDetail field to given value.

### HasCustomNetworkDetail

`func (o *DescribeResourceInstanceResult) HasCustomNetworkDetail() bool`

HasCustomNetworkDetail returns a boolean if a field has been set.

### GetDetailedNetworkTopology

`func (o *DescribeResourceInstanceResult) GetDetailedNetworkTopology() map[string]interface{}`

GetDetailedNetworkTopology returns the DetailedNetworkTopology field if non-nil, zero value otherwise.

### GetDetailedNetworkTopologyOk

`func (o *DescribeResourceInstanceResult) GetDetailedNetworkTopologyOk() (*map[string]interface{}, bool)`

GetDetailedNetworkTopologyOk returns a tuple with the DetailedNetworkTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedNetworkTopology

`func (o *DescribeResourceInstanceResult) SetDetailedNetworkTopology(v map[string]interface{})`

SetDetailedNetworkTopology sets DetailedNetworkTopology field to given value.

### HasDetailedNetworkTopology

`func (o *DescribeResourceInstanceResult) HasDetailedNetworkTopology() bool`

HasDetailedNetworkTopology returns a boolean if a field has been set.

### GetExternalPayerId

`func (o *DescribeResourceInstanceResult) GetExternalPayerId() string`

GetExternalPayerId returns the ExternalPayerId field if non-nil, zero value otherwise.

### GetExternalPayerIdOk

`func (o *DescribeResourceInstanceResult) GetExternalPayerIdOk() (*string, bool)`

GetExternalPayerIdOk returns a tuple with the ExternalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayerId

`func (o *DescribeResourceInstanceResult) SetExternalPayerId(v string)`

SetExternalPayerId sets ExternalPayerId field to given value.

### HasExternalPayerId

`func (o *DescribeResourceInstanceResult) HasExternalPayerId() bool`

HasExternalPayerId returns a boolean if a field has been set.

### GetGcpProjectID

`func (o *DescribeResourceInstanceResult) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *DescribeResourceInstanceResult) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *DescribeResourceInstanceResult) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.

### HasGcpProjectID

`func (o *DescribeResourceInstanceResult) HasGcpProjectID() bool`

HasGcpProjectID returns a boolean if a field has been set.

### GetHighAvailability

`func (o *DescribeResourceInstanceResult) GetHighAvailability() bool`

GetHighAvailability returns the HighAvailability field if non-nil, zero value otherwise.

### GetHighAvailabilityOk

`func (o *DescribeResourceInstanceResult) GetHighAvailabilityOk() (*bool, bool)`

GetHighAvailabilityOk returns a tuple with the HighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailability

`func (o *DescribeResourceInstanceResult) SetHighAvailability(v bool)`

SetHighAvailability sets HighAvailability field to given value.

### HasHighAvailability

`func (o *DescribeResourceInstanceResult) HasHighAvailability() bool`

HasHighAvailability returns a boolean if a field has been set.

### GetId

`func (o *DescribeResourceInstanceResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeResourceInstanceResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeResourceInstanceResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DescribeResourceInstanceResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceLoadStatus

`func (o *DescribeResourceInstanceResult) GetInstanceLoadStatus() string`

GetInstanceLoadStatus returns the InstanceLoadStatus field if non-nil, zero value otherwise.

### GetInstanceLoadStatusOk

`func (o *DescribeResourceInstanceResult) GetInstanceLoadStatusOk() (*string, bool)`

GetInstanceLoadStatusOk returns a tuple with the InstanceLoadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLoadStatus

`func (o *DescribeResourceInstanceResult) SetInstanceLoadStatus(v string)`

SetInstanceLoadStatus sets InstanceLoadStatus field to given value.

### HasInstanceLoadStatus

`func (o *DescribeResourceInstanceResult) HasInstanceLoadStatus() bool`

HasInstanceLoadStatus returns a boolean if a field has been set.

### GetKubernetesDashboardEndpoint

`func (o *DescribeResourceInstanceResult) GetKubernetesDashboardEndpoint() KubernetesDashboardEndpoint`

GetKubernetesDashboardEndpoint returns the KubernetesDashboardEndpoint field if non-nil, zero value otherwise.

### GetKubernetesDashboardEndpointOk

`func (o *DescribeResourceInstanceResult) GetKubernetesDashboardEndpointOk() (*KubernetesDashboardEndpoint, bool)`

GetKubernetesDashboardEndpointOk returns a tuple with the KubernetesDashboardEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesDashboardEndpoint

`func (o *DescribeResourceInstanceResult) SetKubernetesDashboardEndpoint(v KubernetesDashboardEndpoint)`

SetKubernetesDashboardEndpoint sets KubernetesDashboardEndpoint field to given value.

### HasKubernetesDashboardEndpoint

`func (o *DescribeResourceInstanceResult) HasKubernetesDashboardEndpoint() bool`

HasKubernetesDashboardEndpoint returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *DescribeResourceInstanceResult) GetLastModifiedAt() string`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *DescribeResourceInstanceResult) GetLastModifiedAtOk() (*string, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *DescribeResourceInstanceResult) SetLastModifiedAt(v string)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *DescribeResourceInstanceResult) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *DescribeResourceInstanceResult) GetMaxReplicas() string`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *DescribeResourceInstanceResult) GetMaxReplicasOk() (*string, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *DescribeResourceInstanceResult) SetMaxReplicas(v string)`

SetMaxReplicas sets MaxReplicas field to given value.

### HasMaxReplicas

`func (o *DescribeResourceInstanceResult) HasMaxReplicas() bool`

HasMaxReplicas returns a boolean if a field has been set.

### GetMinReplicas

`func (o *DescribeResourceInstanceResult) GetMinReplicas() string`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *DescribeResourceInstanceResult) GetMinReplicasOk() (*string, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *DescribeResourceInstanceResult) SetMinReplicas(v string)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *DescribeResourceInstanceResult) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetNetworkType

`func (o *DescribeResourceInstanceResult) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *DescribeResourceInstanceResult) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *DescribeResourceInstanceResult) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *DescribeResourceInstanceResult) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetProductTierFeatures

`func (o *DescribeResourceInstanceResult) GetProductTierFeatures() map[string]interface{}`

GetProductTierFeatures returns the ProductTierFeatures field if non-nil, zero value otherwise.

### GetProductTierFeaturesOk

`func (o *DescribeResourceInstanceResult) GetProductTierFeaturesOk() (*map[string]interface{}, bool)`

GetProductTierFeaturesOk returns a tuple with the ProductTierFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierFeatures

`func (o *DescribeResourceInstanceResult) SetProductTierFeatures(v map[string]interface{})`

SetProductTierFeatures sets ProductTierFeatures field to given value.

### HasProductTierFeatures

`func (o *DescribeResourceInstanceResult) HasProductTierFeatures() bool`

HasProductTierFeatures returns a boolean if a field has been set.

### GetRegion

`func (o *DescribeResourceInstanceResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DescribeResourceInstanceResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DescribeResourceInstanceResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DescribeResourceInstanceResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResourceID

`func (o *DescribeResourceInstanceResult) GetResourceID() string`

GetResourceID returns the ResourceID field if non-nil, zero value otherwise.

### GetResourceIDOk

`func (o *DescribeResourceInstanceResult) GetResourceIDOk() (*string, bool)`

GetResourceIDOk returns a tuple with the ResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceID

`func (o *DescribeResourceInstanceResult) SetResourceID(v string)`

SetResourceID sets ResourceID field to given value.

### HasResourceID

`func (o *DescribeResourceInstanceResult) HasResourceID() bool`

HasResourceID returns a boolean if a field has been set.

### GetResultParams

`func (o *DescribeResourceInstanceResult) GetResultParams() interface{}`

GetResultParams returns the ResultParams field if non-nil, zero value otherwise.

### GetResultParamsOk

`func (o *DescribeResourceInstanceResult) GetResultParamsOk() (*interface{}, bool)`

GetResultParamsOk returns a tuple with the ResultParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultParams

`func (o *DescribeResourceInstanceResult) SetResultParams(v interface{})`

SetResultParams sets ResultParams field to given value.

### HasResultParams

`func (o *DescribeResourceInstanceResult) HasResultParams() bool`

HasResultParams returns a boolean if a field has been set.

### SetResultParamsNil

`func (o *DescribeResourceInstanceResult) SetResultParamsNil(b bool)`

 SetResultParamsNil sets the value for ResultParams to be an explicit nil

### UnsetResultParams
`func (o *DescribeResourceInstanceResult) UnsetResultParams()`

UnsetResultParams ensures that no value is present for ResultParams, not even an explicit nil
### GetServerlessEnabled

`func (o *DescribeResourceInstanceResult) GetServerlessEnabled() bool`

GetServerlessEnabled returns the ServerlessEnabled field if non-nil, zero value otherwise.

### GetServerlessEnabledOk

`func (o *DescribeResourceInstanceResult) GetServerlessEnabledOk() (*bool, bool)`

GetServerlessEnabledOk returns a tuple with the ServerlessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessEnabled

`func (o *DescribeResourceInstanceResult) SetServerlessEnabled(v bool)`

SetServerlessEnabled sets ServerlessEnabled field to given value.

### HasServerlessEnabled

`func (o *DescribeResourceInstanceResult) HasServerlessEnabled() bool`

HasServerlessEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeResourceInstanceResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeResourceInstanceResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeResourceInstanceResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeResourceInstanceResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DescribeResourceInstanceResult) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DescribeResourceInstanceResult) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DescribeResourceInstanceResult) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DescribeResourceInstanceResult) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


