# FleetSyncAccountConfigCloudNativeNetworksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworks** | Pointer to [**[]FleetSyncAccountConfigCloudNativeNetworkTarget**](FleetSyncAccountConfigCloudNativeNetworkTarget.md) | Optional list of (region, cloudNativeNetworkId) targets to sync. Each target is {region (required), cloudNativeNetworkId (optional)}: with the network ID set, only that network is re-validated; with the network ID omitted, every supported network in the region is enumerated. If the entire list is empty, the sync sweeps every region the account is enabled in (derived from the service plan). | [optional] 
**Id** | **string** | ID of an Account Config | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetSyncAccountConfigCloudNativeNetworksRequest

`func NewFleetSyncAccountConfigCloudNativeNetworksRequest(id string, token string, ) *FleetSyncAccountConfigCloudNativeNetworksRequest`

NewFleetSyncAccountConfigCloudNativeNetworksRequest instantiates a new FleetSyncAccountConfigCloudNativeNetworksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetSyncAccountConfigCloudNativeNetworksRequestWithDefaults

`func NewFleetSyncAccountConfigCloudNativeNetworksRequestWithDefaults() *FleetSyncAccountConfigCloudNativeNetworksRequest`

NewFleetSyncAccountConfigCloudNativeNetworksRequestWithDefaults instantiates a new FleetSyncAccountConfigCloudNativeNetworksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworks

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest) GetCloudNativeNetworks() []FleetSyncAccountConfigCloudNativeNetworkTarget`

GetCloudNativeNetworks returns the CloudNativeNetworks field if non-nil, zero value otherwise.

### GetCloudNativeNetworksOk

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest) GetCloudNativeNetworksOk() (*[]FleetSyncAccountConfigCloudNativeNetworkTarget, bool)`

GetCloudNativeNetworksOk returns a tuple with the CloudNativeNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworks

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest) SetCloudNativeNetworks(v []FleetSyncAccountConfigCloudNativeNetworkTarget)`

SetCloudNativeNetworks sets CloudNativeNetworks field to given value.

### HasCloudNativeNetworks

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest) HasCloudNativeNetworks() bool`

HasCloudNativeNetworks returns a boolean if a field has been set.

### GetId

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


