# DescribeInstanceSnapshotResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Name of the Infra Provider | [optional] 
**CompleteTime** | Pointer to **string** | The snapshot time | [optional] 
**CreatedTime** | Pointer to **string** | The snapshot creation time | [optional] 
**Encrypted** | Pointer to **bool** | Whether the snapshot is encrypted | [optional] 
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

### NewDescribeInstanceSnapshotResult

`func NewDescribeInstanceSnapshotResult() *DescribeInstanceSnapshotResult`

NewDescribeInstanceSnapshotResult instantiates a new DescribeInstanceSnapshotResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeInstanceSnapshotResultWithDefaults

`func NewDescribeInstanceSnapshotResultWithDefaults() *DescribeInstanceSnapshotResult`

NewDescribeInstanceSnapshotResultWithDefaults instantiates a new DescribeInstanceSnapshotResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DescribeInstanceSnapshotResult) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DescribeInstanceSnapshotResult) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DescribeInstanceSnapshotResult) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DescribeInstanceSnapshotResult) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCompleteTime

`func (o *DescribeInstanceSnapshotResult) GetCompleteTime() string`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *DescribeInstanceSnapshotResult) GetCompleteTimeOk() (*string, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *DescribeInstanceSnapshotResult) SetCompleteTime(v string)`

SetCompleteTime sets CompleteTime field to given value.

### HasCompleteTime

`func (o *DescribeInstanceSnapshotResult) HasCompleteTime() bool`

HasCompleteTime returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DescribeInstanceSnapshotResult) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DescribeInstanceSnapshotResult) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DescribeInstanceSnapshotResult) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DescribeInstanceSnapshotResult) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEncrypted

`func (o *DescribeInstanceSnapshotResult) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *DescribeInstanceSnapshotResult) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *DescribeInstanceSnapshotResult) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *DescribeInstanceSnapshotResult) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetOutputParams

`func (o *DescribeInstanceSnapshotResult) GetOutputParams() []OutputParameter`

GetOutputParams returns the OutputParams field if non-nil, zero value otherwise.

### GetOutputParamsOk

`func (o *DescribeInstanceSnapshotResult) GetOutputParamsOk() (*[]OutputParameter, bool)`

GetOutputParamsOk returns a tuple with the OutputParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParams

`func (o *DescribeInstanceSnapshotResult) SetOutputParams(v []OutputParameter)`

SetOutputParams sets OutputParams field to given value.

### HasOutputParams

`func (o *DescribeInstanceSnapshotResult) HasOutputParams() bool`

HasOutputParams returns a boolean if a field has been set.

### GetProductTierId

`func (o *DescribeInstanceSnapshotResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DescribeInstanceSnapshotResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DescribeInstanceSnapshotResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *DescribeInstanceSnapshotResult) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetProductTierName

`func (o *DescribeInstanceSnapshotResult) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *DescribeInstanceSnapshotResult) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *DescribeInstanceSnapshotResult) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.

### HasProductTierName

`func (o *DescribeInstanceSnapshotResult) HasProductTierName() bool`

HasProductTierName returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *DescribeInstanceSnapshotResult) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *DescribeInstanceSnapshotResult) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *DescribeInstanceSnapshotResult) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *DescribeInstanceSnapshotResult) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetProductTierVersionDescription

`func (o *DescribeInstanceSnapshotResult) GetProductTierVersionDescription() string`

GetProductTierVersionDescription returns the ProductTierVersionDescription field if non-nil, zero value otherwise.

### GetProductTierVersionDescriptionOk

`func (o *DescribeInstanceSnapshotResult) GetProductTierVersionDescriptionOk() (*string, bool)`

GetProductTierVersionDescriptionOk returns a tuple with the ProductTierVersionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersionDescription

`func (o *DescribeInstanceSnapshotResult) SetProductTierVersionDescription(v string)`

SetProductTierVersionDescription sets ProductTierVersionDescription field to given value.

### HasProductTierVersionDescription

`func (o *DescribeInstanceSnapshotResult) HasProductTierVersionDescription() bool`

HasProductTierVersionDescription returns a boolean if a field has been set.

### GetProgress

`func (o *DescribeInstanceSnapshotResult) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DescribeInstanceSnapshotResult) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DescribeInstanceSnapshotResult) SetProgress(v int64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DescribeInstanceSnapshotResult) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRegion

`func (o *DescribeInstanceSnapshotResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DescribeInstanceSnapshotResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DescribeInstanceSnapshotResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DescribeInstanceSnapshotResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServiceId

`func (o *DescribeInstanceSnapshotResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeInstanceSnapshotResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeInstanceSnapshotResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DescribeInstanceSnapshotResult) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *DescribeInstanceSnapshotResult) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DescribeInstanceSnapshotResult) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DescribeInstanceSnapshotResult) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *DescribeInstanceSnapshotResult) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSnapshotId

`func (o *DescribeInstanceSnapshotResult) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *DescribeInstanceSnapshotResult) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *DescribeInstanceSnapshotResult) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *DescribeInstanceSnapshotResult) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetSnapshotType

`func (o *DescribeInstanceSnapshotResult) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *DescribeInstanceSnapshotResult) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *DescribeInstanceSnapshotResult) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *DescribeInstanceSnapshotResult) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSourceInstanceId

`func (o *DescribeInstanceSnapshotResult) GetSourceInstanceId() string`

GetSourceInstanceId returns the SourceInstanceId field if non-nil, zero value otherwise.

### GetSourceInstanceIdOk

`func (o *DescribeInstanceSnapshotResult) GetSourceInstanceIdOk() (*string, bool)`

GetSourceInstanceIdOk returns a tuple with the SourceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstanceId

`func (o *DescribeInstanceSnapshotResult) SetSourceInstanceId(v string)`

SetSourceInstanceId sets SourceInstanceId field to given value.

### HasSourceInstanceId

`func (o *DescribeInstanceSnapshotResult) HasSourceInstanceId() bool`

HasSourceInstanceId returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeInstanceSnapshotResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeInstanceSnapshotResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeInstanceSnapshotResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeInstanceSnapshotResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DescribeInstanceSnapshotResult) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DescribeInstanceSnapshotResult) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DescribeInstanceSnapshotResult) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DescribeInstanceSnapshotResult) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionOwnerOrgId

`func (o *DescribeInstanceSnapshotResult) GetSubscriptionOwnerOrgId() string`

GetSubscriptionOwnerOrgId returns the SubscriptionOwnerOrgId field if non-nil, zero value otherwise.

### GetSubscriptionOwnerOrgIdOk

`func (o *DescribeInstanceSnapshotResult) GetSubscriptionOwnerOrgIdOk() (*string, bool)`

GetSubscriptionOwnerOrgIdOk returns a tuple with the SubscriptionOwnerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerOrgId

`func (o *DescribeInstanceSnapshotResult) SetSubscriptionOwnerOrgId(v string)`

SetSubscriptionOwnerOrgId sets SubscriptionOwnerOrgId field to given value.

### HasSubscriptionOwnerOrgId

`func (o *DescribeInstanceSnapshotResult) HasSubscriptionOwnerOrgId() bool`

HasSubscriptionOwnerOrgId returns a boolean if a field has been set.

### GetSubscriptionOwnerOrgName

`func (o *DescribeInstanceSnapshotResult) GetSubscriptionOwnerOrgName() string`

GetSubscriptionOwnerOrgName returns the SubscriptionOwnerOrgName field if non-nil, zero value otherwise.

### GetSubscriptionOwnerOrgNameOk

`func (o *DescribeInstanceSnapshotResult) GetSubscriptionOwnerOrgNameOk() (*string, bool)`

GetSubscriptionOwnerOrgNameOk returns a tuple with the SubscriptionOwnerOrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerOrgName

`func (o *DescribeInstanceSnapshotResult) SetSubscriptionOwnerOrgName(v string)`

SetSubscriptionOwnerOrgName sets SubscriptionOwnerOrgName field to given value.

### HasSubscriptionOwnerOrgName

`func (o *DescribeInstanceSnapshotResult) HasSubscriptionOwnerOrgName() bool`

HasSubscriptionOwnerOrgName returns a boolean if a field has been set.

### GetSubscriptionOwnerUserId

`func (o *DescribeInstanceSnapshotResult) GetSubscriptionOwnerUserId() string`

GetSubscriptionOwnerUserId returns the SubscriptionOwnerUserId field if non-nil, zero value otherwise.

### GetSubscriptionOwnerUserIdOk

`func (o *DescribeInstanceSnapshotResult) GetSubscriptionOwnerUserIdOk() (*string, bool)`

GetSubscriptionOwnerUserIdOk returns a tuple with the SubscriptionOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerUserId

`func (o *DescribeInstanceSnapshotResult) SetSubscriptionOwnerUserId(v string)`

SetSubscriptionOwnerUserId sets SubscriptionOwnerUserId field to given value.

### HasSubscriptionOwnerUserId

`func (o *DescribeInstanceSnapshotResult) HasSubscriptionOwnerUserId() bool`

HasSubscriptionOwnerUserId returns a boolean if a field has been set.

### GetSubscriptionOwnerUserName

`func (o *DescribeInstanceSnapshotResult) GetSubscriptionOwnerUserName() string`

GetSubscriptionOwnerUserName returns the SubscriptionOwnerUserName field if non-nil, zero value otherwise.

### GetSubscriptionOwnerUserNameOk

`func (o *DescribeInstanceSnapshotResult) GetSubscriptionOwnerUserNameOk() (*string, bool)`

GetSubscriptionOwnerUserNameOk returns a tuple with the SubscriptionOwnerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerUserName

`func (o *DescribeInstanceSnapshotResult) SetSubscriptionOwnerUserName(v string)`

SetSubscriptionOwnerUserName sets SubscriptionOwnerUserName field to given value.

### HasSubscriptionOwnerUserName

`func (o *DescribeInstanceSnapshotResult) HasSubscriptionOwnerUserName() bool`

HasSubscriptionOwnerUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


