# FleetSyncAccountConfigCloudNativeNetworksRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworks** | Pointer to [**[]FleetSyncAccountConfigCloudNativeNetworkTarget**](FleetSyncAccountConfigCloudNativeNetworkTarget.md) | Optional list of (region, cloudNativeNetworkId) targets to sync. Each target is {region (required), cloudNativeNetworkId (optional)}: with the network ID set, only that network is re-validated; with the network ID omitted, every supported network in the region is enumerated. If the entire list is empty, the sync sweeps every region the account is enabled in (derived from the service plan). | [optional] 

## Methods

### NewFleetSyncAccountConfigCloudNativeNetworksRequest2

`func NewFleetSyncAccountConfigCloudNativeNetworksRequest2() *FleetSyncAccountConfigCloudNativeNetworksRequest2`

NewFleetSyncAccountConfigCloudNativeNetworksRequest2 instantiates a new FleetSyncAccountConfigCloudNativeNetworksRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetSyncAccountConfigCloudNativeNetworksRequest2WithDefaults

`func NewFleetSyncAccountConfigCloudNativeNetworksRequest2WithDefaults() *FleetSyncAccountConfigCloudNativeNetworksRequest2`

NewFleetSyncAccountConfigCloudNativeNetworksRequest2WithDefaults instantiates a new FleetSyncAccountConfigCloudNativeNetworksRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworks

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest2) GetCloudNativeNetworks() []FleetSyncAccountConfigCloudNativeNetworkTarget`

GetCloudNativeNetworks returns the CloudNativeNetworks field if non-nil, zero value otherwise.

### GetCloudNativeNetworksOk

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest2) GetCloudNativeNetworksOk() (*[]FleetSyncAccountConfigCloudNativeNetworkTarget, bool)`

GetCloudNativeNetworksOk returns a tuple with the CloudNativeNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworks

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest2) SetCloudNativeNetworks(v []FleetSyncAccountConfigCloudNativeNetworkTarget)`

SetCloudNativeNetworks sets CloudNativeNetworks field to given value.

### HasCloudNativeNetworks

`func (o *FleetSyncAccountConfigCloudNativeNetworksRequest2) HasCloudNativeNetworks() bool`

HasCloudNativeNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


