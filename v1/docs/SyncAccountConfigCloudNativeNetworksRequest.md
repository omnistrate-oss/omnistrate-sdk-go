# SyncAccountConfigCloudNativeNetworksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworks** | Pointer to [**[]SyncAccountConfigCloudNativeNetworkTarget**](SyncAccountConfigCloudNativeNetworkTarget.md) | Optional list of (region, cloudNativeNetworkId) targets to sync. Each target is {region (required), cloudNativeNetworkId (optional)}: with the network ID set, only that VPC is re-validated; with the network ID omitted, every VPC in the region is enumerated. If the entire list is empty, the sync sweeps every region the account is enabled in (derived from the service plan). | [optional] 
**Id** | **string** | ID of an Account Config | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewSyncAccountConfigCloudNativeNetworksRequest

`func NewSyncAccountConfigCloudNativeNetworksRequest(id string, token string, ) *SyncAccountConfigCloudNativeNetworksRequest`

NewSyncAccountConfigCloudNativeNetworksRequest instantiates a new SyncAccountConfigCloudNativeNetworksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncAccountConfigCloudNativeNetworksRequestWithDefaults

`func NewSyncAccountConfigCloudNativeNetworksRequestWithDefaults() *SyncAccountConfigCloudNativeNetworksRequest`

NewSyncAccountConfigCloudNativeNetworksRequestWithDefaults instantiates a new SyncAccountConfigCloudNativeNetworksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworks

`func (o *SyncAccountConfigCloudNativeNetworksRequest) GetCloudNativeNetworks() []SyncAccountConfigCloudNativeNetworkTarget`

GetCloudNativeNetworks returns the CloudNativeNetworks field if non-nil, zero value otherwise.

### GetCloudNativeNetworksOk

`func (o *SyncAccountConfigCloudNativeNetworksRequest) GetCloudNativeNetworksOk() (*[]SyncAccountConfigCloudNativeNetworkTarget, bool)`

GetCloudNativeNetworksOk returns a tuple with the CloudNativeNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworks

`func (o *SyncAccountConfigCloudNativeNetworksRequest) SetCloudNativeNetworks(v []SyncAccountConfigCloudNativeNetworkTarget)`

SetCloudNativeNetworks sets CloudNativeNetworks field to given value.

### HasCloudNativeNetworks

`func (o *SyncAccountConfigCloudNativeNetworksRequest) HasCloudNativeNetworks() bool`

HasCloudNativeNetworks returns a boolean if a field has been set.

### GetId

`func (o *SyncAccountConfigCloudNativeNetworksRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncAccountConfigCloudNativeNetworksRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncAccountConfigCloudNativeNetworksRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *SyncAccountConfigCloudNativeNetworksRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SyncAccountConfigCloudNativeNetworksRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SyncAccountConfigCloudNativeNetworksRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


