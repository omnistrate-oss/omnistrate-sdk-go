# SyncAccountConfigCloudNativeNetworksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of an Account Config | 
**Regions** | Pointer to **[]string** | Cloud regions to discover CloudNativeNetworks in. If not provided, all regions from the service plan are used. | [optional] 
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


### GetRegions

`func (o *SyncAccountConfigCloudNativeNetworksRequest) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *SyncAccountConfigCloudNativeNetworksRequest) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *SyncAccountConfigCloudNativeNetworksRequest) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *SyncAccountConfigCloudNativeNetworksRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

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


