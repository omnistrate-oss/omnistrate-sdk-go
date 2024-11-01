# DescribeResourceInstanceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | True if this resource instance has associated infrastructure deployed | 
**AutoscalingEnabled** | Pointer to **bool** | Whether the instance has autoscaling enabled | [optional] 
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**BackupStatus** | Pointer to [**BackupStatus**](BackupStatus.md) |  | [optional] 
**CloudProvider** | Pointer to **string** | The cloud provider name | [optional] 
**CreatedAt** | **string** | The instance creation time | 
**CreatedByUserId** | **string** | The user ID that created the resource instance | 
**CreatedByUserName** | **string** | The name of the user that created the resource instance | 
**CustomNetworkDetail** | Pointer to [**CustomNetworkResourceDetail**](CustomNetworkResourceDetail.md) |  | [optional] 
**DetailedNetworkTopology** | Pointer to [**map[string]ResourceNetworkTopologyResult**](ResourceNetworkTopologyResult.md) | The detailed network topology | [optional] 
**ExternalPayerId** | Pointer to **string** | The external payer id to record which customer should pay for this resource instance | [optional] 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**HighAvailability** | Pointer to **bool** | Whether the instance is High Availability | [optional] 
**Id** | **string** | The instance ID | 
**LastModifiedAt** | **string** | The instance update time | 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**ProductTierFeatures** | **map[string]interface{}** | The product tier features | 
**Region** | Pointer to **string** | The region code | [optional] 
**ResultParams** | **interface{}** | Custom result parameters | 
**Status** | **string** | The instance status | 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 

## Methods

### NewDescribeResourceInstanceResult

`func NewDescribeResourceInstanceResult(active bool, createdAt string, createdByUserId string, createdByUserName string, id string, lastModifiedAt string, productTierFeatures map[string]interface{}, resultParams interface{}, status string, ) *DescribeResourceInstanceResult`

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

`func (o *DescribeResourceInstanceResult) GetDetailedNetworkTopology() map[string]ResourceNetworkTopologyResult`

GetDetailedNetworkTopology returns the DetailedNetworkTopology field if non-nil, zero value otherwise.

### GetDetailedNetworkTopologyOk

`func (o *DescribeResourceInstanceResult) GetDetailedNetworkTopologyOk() (*map[string]ResourceNetworkTopologyResult, bool)`

GetDetailedNetworkTopologyOk returns a tuple with the DetailedNetworkTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedNetworkTopology

`func (o *DescribeResourceInstanceResult) SetDetailedNetworkTopology(v map[string]ResourceNetworkTopologyResult)`

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


### SetResultParamsNil

`func (o *DescribeResourceInstanceResult) SetResultParamsNil(b bool)`

 SetResultParamsNil sets the value for ResultParams to be an explicit nil

### UnsetResultParams
`func (o *DescribeResourceInstanceResult) UnsetResultParams()`

UnsetResultParams ensures that no value is present for ResultParams, not even an explicit nil
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


