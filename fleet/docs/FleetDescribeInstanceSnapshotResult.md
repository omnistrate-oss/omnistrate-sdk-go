# FleetDescribeInstanceSnapshotResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Name of the Infra Provider | [optional] 
**CompleteTime** | Pointer to **string** | The snapshot time | [optional] 
**CreatedTime** | Pointer to **string** | The snapshot creation time | [optional] 
**Encrypted** | Pointer to **bool** | Whether the snapshot is encrypted | [optional] 
**EnvironmentId** | Pointer to **string** | ID of a Service Environment | [optional] 
**OutputParams** | Pointer to [**[]OutputParameter**](OutputParameter.md) | Custom output parameters | [optional] 
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ProductTierName** | Pointer to **string** | The product tier name | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version | [optional] 
**ProductTierVersionDescription** | Pointer to **string** | The product tier version description | [optional] 
**Progress** | Pointer to **int64** | The backup progress. 0-100 | [optional] 
**Region** | Pointer to **string** | The region name where the snapshot is stored | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**ServiceName** | Pointer to **string** | The service name | [optional] 
**SnapshotId** | Pointer to **string** | ID of a Resource Instance Snapshot | [optional] 
**SnapshotType** | Pointer to **string** | The snapshot type | [optional] 
**SourceInstanceId** | Pointer to **string** | ID of a Resource Instance | [optional] 
**Status** | Pointer to **string** | The snapshot status | [optional] 
**SubscriptionId** | Pointer to **string** | ID of a Subscription | [optional] 
**SubscriptionOwnerOrgId** | Pointer to **string** | ID of an Org | [optional] 
**SubscriptionOwnerOrgName** | Pointer to **string** | The subscription owner org name | [optional] 
**SubscriptionOwnerUserId** | Pointer to **string** | ID of a User | [optional] 
**SubscriptionOwnerUserName** | Pointer to **string** | The subscription owner user name | [optional] 

## Methods

### NewFleetDescribeInstanceSnapshotResult

`func NewFleetDescribeInstanceSnapshotResult() *FleetDescribeInstanceSnapshotResult`

NewFleetDescribeInstanceSnapshotResult instantiates a new FleetDescribeInstanceSnapshotResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDescribeInstanceSnapshotResultWithDefaults

`func NewFleetDescribeInstanceSnapshotResultWithDefaults() *FleetDescribeInstanceSnapshotResult`

NewFleetDescribeInstanceSnapshotResultWithDefaults instantiates a new FleetDescribeInstanceSnapshotResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *FleetDescribeInstanceSnapshotResult) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *FleetDescribeInstanceSnapshotResult) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *FleetDescribeInstanceSnapshotResult) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *FleetDescribeInstanceSnapshotResult) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCompleteTime

`func (o *FleetDescribeInstanceSnapshotResult) GetCompleteTime() string`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *FleetDescribeInstanceSnapshotResult) GetCompleteTimeOk() (*string, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *FleetDescribeInstanceSnapshotResult) SetCompleteTime(v string)`

SetCompleteTime sets CompleteTime field to given value.

### HasCompleteTime

`func (o *FleetDescribeInstanceSnapshotResult) HasCompleteTime() bool`

HasCompleteTime returns a boolean if a field has been set.

### GetCreatedTime

`func (o *FleetDescribeInstanceSnapshotResult) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FleetDescribeInstanceSnapshotResult) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FleetDescribeInstanceSnapshotResult) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *FleetDescribeInstanceSnapshotResult) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEncrypted

`func (o *FleetDescribeInstanceSnapshotResult) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *FleetDescribeInstanceSnapshotResult) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *FleetDescribeInstanceSnapshotResult) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *FleetDescribeInstanceSnapshotResult) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *FleetDescribeInstanceSnapshotResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetDescribeInstanceSnapshotResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *FleetDescribeInstanceSnapshotResult) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetOutputParams

`func (o *FleetDescribeInstanceSnapshotResult) GetOutputParams() []OutputParameter`

GetOutputParams returns the OutputParams field if non-nil, zero value otherwise.

### GetOutputParamsOk

`func (o *FleetDescribeInstanceSnapshotResult) GetOutputParamsOk() (*[]OutputParameter, bool)`

GetOutputParamsOk returns a tuple with the OutputParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParams

`func (o *FleetDescribeInstanceSnapshotResult) SetOutputParams(v []OutputParameter)`

SetOutputParams sets OutputParams field to given value.

### HasOutputParams

`func (o *FleetDescribeInstanceSnapshotResult) HasOutputParams() bool`

HasOutputParams returns a boolean if a field has been set.

### GetProductTierId

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *FleetDescribeInstanceSnapshotResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *FleetDescribeInstanceSnapshotResult) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetProductTierName

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *FleetDescribeInstanceSnapshotResult) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.

### HasProductTierName

`func (o *FleetDescribeInstanceSnapshotResult) HasProductTierName() bool`

HasProductTierName returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *FleetDescribeInstanceSnapshotResult) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *FleetDescribeInstanceSnapshotResult) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetProductTierVersionDescription

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierVersionDescription() string`

GetProductTierVersionDescription returns the ProductTierVersionDescription field if non-nil, zero value otherwise.

### GetProductTierVersionDescriptionOk

`func (o *FleetDescribeInstanceSnapshotResult) GetProductTierVersionDescriptionOk() (*string, bool)`

GetProductTierVersionDescriptionOk returns a tuple with the ProductTierVersionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersionDescription

`func (o *FleetDescribeInstanceSnapshotResult) SetProductTierVersionDescription(v string)`

SetProductTierVersionDescription sets ProductTierVersionDescription field to given value.

### HasProductTierVersionDescription

`func (o *FleetDescribeInstanceSnapshotResult) HasProductTierVersionDescription() bool`

HasProductTierVersionDescription returns a boolean if a field has been set.

### GetProgress

`func (o *FleetDescribeInstanceSnapshotResult) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FleetDescribeInstanceSnapshotResult) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FleetDescribeInstanceSnapshotResult) SetProgress(v int64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *FleetDescribeInstanceSnapshotResult) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRegion

`func (o *FleetDescribeInstanceSnapshotResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FleetDescribeInstanceSnapshotResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FleetDescribeInstanceSnapshotResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FleetDescribeInstanceSnapshotResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetDescribeInstanceSnapshotResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetDescribeInstanceSnapshotResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *FleetDescribeInstanceSnapshotResult) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *FleetDescribeInstanceSnapshotResult) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *FleetDescribeInstanceSnapshotResult) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *FleetDescribeInstanceSnapshotResult) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *FleetDescribeInstanceSnapshotResult) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSnapshotId

`func (o *FleetDescribeInstanceSnapshotResult) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *FleetDescribeInstanceSnapshotResult) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *FleetDescribeInstanceSnapshotResult) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetSnapshotType

`func (o *FleetDescribeInstanceSnapshotResult) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *FleetDescribeInstanceSnapshotResult) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *FleetDescribeInstanceSnapshotResult) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *FleetDescribeInstanceSnapshotResult) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSourceInstanceId

`func (o *FleetDescribeInstanceSnapshotResult) GetSourceInstanceId() string`

GetSourceInstanceId returns the SourceInstanceId field if non-nil, zero value otherwise.

### GetSourceInstanceIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetSourceInstanceIdOk() (*string, bool)`

GetSourceInstanceIdOk returns a tuple with the SourceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstanceId

`func (o *FleetDescribeInstanceSnapshotResult) SetSourceInstanceId(v string)`

SetSourceInstanceId sets SourceInstanceId field to given value.

### HasSourceInstanceId

`func (o *FleetDescribeInstanceSnapshotResult) HasSourceInstanceId() bool`

HasSourceInstanceId returns a boolean if a field has been set.

### GetStatus

`func (o *FleetDescribeInstanceSnapshotResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetDescribeInstanceSnapshotResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetDescribeInstanceSnapshotResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FleetDescribeInstanceSnapshotResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *FleetDescribeInstanceSnapshotResult) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *FleetDescribeInstanceSnapshotResult) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *FleetDescribeInstanceSnapshotResult) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionOwnerOrgId

`func (o *FleetDescribeInstanceSnapshotResult) GetSubscriptionOwnerOrgId() string`

GetSubscriptionOwnerOrgId returns the SubscriptionOwnerOrgId field if non-nil, zero value otherwise.

### GetSubscriptionOwnerOrgIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetSubscriptionOwnerOrgIdOk() (*string, bool)`

GetSubscriptionOwnerOrgIdOk returns a tuple with the SubscriptionOwnerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerOrgId

`func (o *FleetDescribeInstanceSnapshotResult) SetSubscriptionOwnerOrgId(v string)`

SetSubscriptionOwnerOrgId sets SubscriptionOwnerOrgId field to given value.

### HasSubscriptionOwnerOrgId

`func (o *FleetDescribeInstanceSnapshotResult) HasSubscriptionOwnerOrgId() bool`

HasSubscriptionOwnerOrgId returns a boolean if a field has been set.

### GetSubscriptionOwnerOrgName

`func (o *FleetDescribeInstanceSnapshotResult) GetSubscriptionOwnerOrgName() string`

GetSubscriptionOwnerOrgName returns the SubscriptionOwnerOrgName field if non-nil, zero value otherwise.

### GetSubscriptionOwnerOrgNameOk

`func (o *FleetDescribeInstanceSnapshotResult) GetSubscriptionOwnerOrgNameOk() (*string, bool)`

GetSubscriptionOwnerOrgNameOk returns a tuple with the SubscriptionOwnerOrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerOrgName

`func (o *FleetDescribeInstanceSnapshotResult) SetSubscriptionOwnerOrgName(v string)`

SetSubscriptionOwnerOrgName sets SubscriptionOwnerOrgName field to given value.

### HasSubscriptionOwnerOrgName

`func (o *FleetDescribeInstanceSnapshotResult) HasSubscriptionOwnerOrgName() bool`

HasSubscriptionOwnerOrgName returns a boolean if a field has been set.

### GetSubscriptionOwnerUserId

`func (o *FleetDescribeInstanceSnapshotResult) GetSubscriptionOwnerUserId() string`

GetSubscriptionOwnerUserId returns the SubscriptionOwnerUserId field if non-nil, zero value otherwise.

### GetSubscriptionOwnerUserIdOk

`func (o *FleetDescribeInstanceSnapshotResult) GetSubscriptionOwnerUserIdOk() (*string, bool)`

GetSubscriptionOwnerUserIdOk returns a tuple with the SubscriptionOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerUserId

`func (o *FleetDescribeInstanceSnapshotResult) SetSubscriptionOwnerUserId(v string)`

SetSubscriptionOwnerUserId sets SubscriptionOwnerUserId field to given value.

### HasSubscriptionOwnerUserId

`func (o *FleetDescribeInstanceSnapshotResult) HasSubscriptionOwnerUserId() bool`

HasSubscriptionOwnerUserId returns a boolean if a field has been set.

### GetSubscriptionOwnerUserName

`func (o *FleetDescribeInstanceSnapshotResult) GetSubscriptionOwnerUserName() string`

GetSubscriptionOwnerUserName returns the SubscriptionOwnerUserName field if non-nil, zero value otherwise.

### GetSubscriptionOwnerUserNameOk

`func (o *FleetDescribeInstanceSnapshotResult) GetSubscriptionOwnerUserNameOk() (*string, bool)`

GetSubscriptionOwnerUserNameOk returns a tuple with the SubscriptionOwnerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerUserName

`func (o *FleetDescribeInstanceSnapshotResult) SetSubscriptionOwnerUserName(v string)`

SetSubscriptionOwnerUserName sets SubscriptionOwnerUserName field to given value.

### HasSubscriptionOwnerUserName

`func (o *FleetDescribeInstanceSnapshotResult) HasSubscriptionOwnerUserName() bool`

HasSubscriptionOwnerUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


