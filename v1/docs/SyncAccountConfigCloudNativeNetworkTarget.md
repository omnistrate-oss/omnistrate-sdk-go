# SyncAccountConfigCloudNativeNetworkTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworkId** | Pointer to **string** | Optional cloud provider network ID (e.g. AWS VPC ID). When omitted, every VPC in the region is enumerated. | [optional] 
**IncludeHostClusters** | Pointer to **bool** | Whether to include host clusters when refreshing this target. | [optional] 
**Region** | **string** | The cloud region where the network resides | 

## Methods

### NewSyncAccountConfigCloudNativeNetworkTarget

`func NewSyncAccountConfigCloudNativeNetworkTarget(region string, ) *SyncAccountConfigCloudNativeNetworkTarget`

NewSyncAccountConfigCloudNativeNetworkTarget instantiates a new SyncAccountConfigCloudNativeNetworkTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncAccountConfigCloudNativeNetworkTargetWithDefaults

`func NewSyncAccountConfigCloudNativeNetworkTargetWithDefaults() *SyncAccountConfigCloudNativeNetworkTarget`

NewSyncAccountConfigCloudNativeNetworkTargetWithDefaults instantiates a new SyncAccountConfigCloudNativeNetworkTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworkId

`func (o *SyncAccountConfigCloudNativeNetworkTarget) GetCloudNativeNetworkId() string`

GetCloudNativeNetworkId returns the CloudNativeNetworkId field if non-nil, zero value otherwise.

### GetCloudNativeNetworkIdOk

`func (o *SyncAccountConfigCloudNativeNetworkTarget) GetCloudNativeNetworkIdOk() (*string, bool)`

GetCloudNativeNetworkIdOk returns a tuple with the CloudNativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworkId

`func (o *SyncAccountConfigCloudNativeNetworkTarget) SetCloudNativeNetworkId(v string)`

SetCloudNativeNetworkId sets CloudNativeNetworkId field to given value.

### HasCloudNativeNetworkId

`func (o *SyncAccountConfigCloudNativeNetworkTarget) HasCloudNativeNetworkId() bool`

HasCloudNativeNetworkId returns a boolean if a field has been set.

### GetIncludeHostClusters

`func (o *SyncAccountConfigCloudNativeNetworkTarget) GetIncludeHostClusters() bool`

GetIncludeHostClusters returns the IncludeHostClusters field if non-nil, zero value otherwise.

### GetIncludeHostClustersOk

`func (o *SyncAccountConfigCloudNativeNetworkTarget) GetIncludeHostClustersOk() (*bool, bool)`

GetIncludeHostClustersOk returns a tuple with the IncludeHostClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHostClusters

`func (o *SyncAccountConfigCloudNativeNetworkTarget) SetIncludeHostClusters(v bool)`

SetIncludeHostClusters sets IncludeHostClusters field to given value.

### HasIncludeHostClusters

`func (o *SyncAccountConfigCloudNativeNetworkTarget) HasIncludeHostClusters() bool`

HasIncludeHostClusters returns a boolean if a field has been set.

### GetRegion

`func (o *SyncAccountConfigCloudNativeNetworkTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SyncAccountConfigCloudNativeNetworkTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SyncAccountConfigCloudNativeNetworkTarget) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


