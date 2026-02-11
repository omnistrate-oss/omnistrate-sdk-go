# ListInstanceSnapshotResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Name of the Infra Provider | [optional] 
**CompleteTime** | Pointer to **string** | The snapshot time | [optional] 
**CreatedTime** | Pointer to **string** | The snapshot creation time | [optional] 
**Encrypted** | Pointer to **bool** | Whether the snapshot is encrypted | [optional] 
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

### NewListInstanceSnapshotResult

`func NewListInstanceSnapshotResult() *ListInstanceSnapshotResult`

NewListInstanceSnapshotResult instantiates a new ListInstanceSnapshotResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceSnapshotResultWithDefaults

`func NewListInstanceSnapshotResultWithDefaults() *ListInstanceSnapshotResult`

NewListInstanceSnapshotResultWithDefaults instantiates a new ListInstanceSnapshotResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ListInstanceSnapshotResult) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ListInstanceSnapshotResult) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ListInstanceSnapshotResult) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ListInstanceSnapshotResult) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCompleteTime

`func (o *ListInstanceSnapshotResult) GetCompleteTime() string`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *ListInstanceSnapshotResult) GetCompleteTimeOk() (*string, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *ListInstanceSnapshotResult) SetCompleteTime(v string)`

SetCompleteTime sets CompleteTime field to given value.

### HasCompleteTime

`func (o *ListInstanceSnapshotResult) HasCompleteTime() bool`

HasCompleteTime returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ListInstanceSnapshotResult) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ListInstanceSnapshotResult) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ListInstanceSnapshotResult) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ListInstanceSnapshotResult) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEncrypted

`func (o *ListInstanceSnapshotResult) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *ListInstanceSnapshotResult) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *ListInstanceSnapshotResult) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *ListInstanceSnapshotResult) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetProductTierId

`func (o *ListInstanceSnapshotResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ListInstanceSnapshotResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ListInstanceSnapshotResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *ListInstanceSnapshotResult) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetProductTierName

`func (o *ListInstanceSnapshotResult) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *ListInstanceSnapshotResult) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *ListInstanceSnapshotResult) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.

### HasProductTierName

`func (o *ListInstanceSnapshotResult) HasProductTierName() bool`

HasProductTierName returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *ListInstanceSnapshotResult) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *ListInstanceSnapshotResult) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *ListInstanceSnapshotResult) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *ListInstanceSnapshotResult) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetProductTierVersionDescription

`func (o *ListInstanceSnapshotResult) GetProductTierVersionDescription() string`

GetProductTierVersionDescription returns the ProductTierVersionDescription field if non-nil, zero value otherwise.

### GetProductTierVersionDescriptionOk

`func (o *ListInstanceSnapshotResult) GetProductTierVersionDescriptionOk() (*string, bool)`

GetProductTierVersionDescriptionOk returns a tuple with the ProductTierVersionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersionDescription

`func (o *ListInstanceSnapshotResult) SetProductTierVersionDescription(v string)`

SetProductTierVersionDescription sets ProductTierVersionDescription field to given value.

### HasProductTierVersionDescription

`func (o *ListInstanceSnapshotResult) HasProductTierVersionDescription() bool`

HasProductTierVersionDescription returns a boolean if a field has been set.

### GetProgress

`func (o *ListInstanceSnapshotResult) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ListInstanceSnapshotResult) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ListInstanceSnapshotResult) SetProgress(v int64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ListInstanceSnapshotResult) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRegion

`func (o *ListInstanceSnapshotResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ListInstanceSnapshotResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ListInstanceSnapshotResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ListInstanceSnapshotResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServiceId

`func (o *ListInstanceSnapshotResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListInstanceSnapshotResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListInstanceSnapshotResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ListInstanceSnapshotResult) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *ListInstanceSnapshotResult) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ListInstanceSnapshotResult) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ListInstanceSnapshotResult) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ListInstanceSnapshotResult) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSnapshotId

`func (o *ListInstanceSnapshotResult) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ListInstanceSnapshotResult) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ListInstanceSnapshotResult) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ListInstanceSnapshotResult) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetSnapshotType

`func (o *ListInstanceSnapshotResult) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *ListInstanceSnapshotResult) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *ListInstanceSnapshotResult) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *ListInstanceSnapshotResult) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSourceInstanceId

`func (o *ListInstanceSnapshotResult) GetSourceInstanceId() string`

GetSourceInstanceId returns the SourceInstanceId field if non-nil, zero value otherwise.

### GetSourceInstanceIdOk

`func (o *ListInstanceSnapshotResult) GetSourceInstanceIdOk() (*string, bool)`

GetSourceInstanceIdOk returns a tuple with the SourceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstanceId

`func (o *ListInstanceSnapshotResult) SetSourceInstanceId(v string)`

SetSourceInstanceId sets SourceInstanceId field to given value.

### HasSourceInstanceId

`func (o *ListInstanceSnapshotResult) HasSourceInstanceId() bool`

HasSourceInstanceId returns a boolean if a field has been set.

### GetStatus

`func (o *ListInstanceSnapshotResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListInstanceSnapshotResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListInstanceSnapshotResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListInstanceSnapshotResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *ListInstanceSnapshotResult) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ListInstanceSnapshotResult) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ListInstanceSnapshotResult) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ListInstanceSnapshotResult) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionOwnerOrgId

`func (o *ListInstanceSnapshotResult) GetSubscriptionOwnerOrgId() string`

GetSubscriptionOwnerOrgId returns the SubscriptionOwnerOrgId field if non-nil, zero value otherwise.

### GetSubscriptionOwnerOrgIdOk

`func (o *ListInstanceSnapshotResult) GetSubscriptionOwnerOrgIdOk() (*string, bool)`

GetSubscriptionOwnerOrgIdOk returns a tuple with the SubscriptionOwnerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerOrgId

`func (o *ListInstanceSnapshotResult) SetSubscriptionOwnerOrgId(v string)`

SetSubscriptionOwnerOrgId sets SubscriptionOwnerOrgId field to given value.

### HasSubscriptionOwnerOrgId

`func (o *ListInstanceSnapshotResult) HasSubscriptionOwnerOrgId() bool`

HasSubscriptionOwnerOrgId returns a boolean if a field has been set.

### GetSubscriptionOwnerOrgName

`func (o *ListInstanceSnapshotResult) GetSubscriptionOwnerOrgName() string`

GetSubscriptionOwnerOrgName returns the SubscriptionOwnerOrgName field if non-nil, zero value otherwise.

### GetSubscriptionOwnerOrgNameOk

`func (o *ListInstanceSnapshotResult) GetSubscriptionOwnerOrgNameOk() (*string, bool)`

GetSubscriptionOwnerOrgNameOk returns a tuple with the SubscriptionOwnerOrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerOrgName

`func (o *ListInstanceSnapshotResult) SetSubscriptionOwnerOrgName(v string)`

SetSubscriptionOwnerOrgName sets SubscriptionOwnerOrgName field to given value.

### HasSubscriptionOwnerOrgName

`func (o *ListInstanceSnapshotResult) HasSubscriptionOwnerOrgName() bool`

HasSubscriptionOwnerOrgName returns a boolean if a field has been set.

### GetSubscriptionOwnerUserId

`func (o *ListInstanceSnapshotResult) GetSubscriptionOwnerUserId() string`

GetSubscriptionOwnerUserId returns the SubscriptionOwnerUserId field if non-nil, zero value otherwise.

### GetSubscriptionOwnerUserIdOk

`func (o *ListInstanceSnapshotResult) GetSubscriptionOwnerUserIdOk() (*string, bool)`

GetSubscriptionOwnerUserIdOk returns a tuple with the SubscriptionOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerUserId

`func (o *ListInstanceSnapshotResult) SetSubscriptionOwnerUserId(v string)`

SetSubscriptionOwnerUserId sets SubscriptionOwnerUserId field to given value.

### HasSubscriptionOwnerUserId

`func (o *ListInstanceSnapshotResult) HasSubscriptionOwnerUserId() bool`

HasSubscriptionOwnerUserId returns a boolean if a field has been set.

### GetSubscriptionOwnerUserName

`func (o *ListInstanceSnapshotResult) GetSubscriptionOwnerUserName() string`

GetSubscriptionOwnerUserName returns the SubscriptionOwnerUserName field if non-nil, zero value otherwise.

### GetSubscriptionOwnerUserNameOk

`func (o *ListInstanceSnapshotResult) GetSubscriptionOwnerUserNameOk() (*string, bool)`

GetSubscriptionOwnerUserNameOk returns a tuple with the SubscriptionOwnerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerUserName

`func (o *ListInstanceSnapshotResult) SetSubscriptionOwnerUserName(v string)`

SetSubscriptionOwnerUserName sets SubscriptionOwnerUserName field to given value.

### HasSubscriptionOwnerUserName

`func (o *ListInstanceSnapshotResult) HasSubscriptionOwnerUserName() bool`

HasSubscriptionOwnerUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


