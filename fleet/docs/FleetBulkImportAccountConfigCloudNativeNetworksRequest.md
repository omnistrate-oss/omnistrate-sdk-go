# FleetBulkImportAccountConfigCloudNativeNetworksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworks** | [**[]FleetAccountConfigCloudNativeNetworkOperation**](FleetAccountConfigCloudNativeNetworkOperation.md) | List of cloud native network operations to perform | 
**Id** | **string** | ID of an Account Config | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetBulkImportAccountConfigCloudNativeNetworksRequest

`func NewFleetBulkImportAccountConfigCloudNativeNetworksRequest(cloudNativeNetworks []FleetAccountConfigCloudNativeNetworkOperation, id string, token string, ) *FleetBulkImportAccountConfigCloudNativeNetworksRequest`

NewFleetBulkImportAccountConfigCloudNativeNetworksRequest instantiates a new FleetBulkImportAccountConfigCloudNativeNetworksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetBulkImportAccountConfigCloudNativeNetworksRequestWithDefaults

`func NewFleetBulkImportAccountConfigCloudNativeNetworksRequestWithDefaults() *FleetBulkImportAccountConfigCloudNativeNetworksRequest`

NewFleetBulkImportAccountConfigCloudNativeNetworksRequestWithDefaults instantiates a new FleetBulkImportAccountConfigCloudNativeNetworksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworks

`func (o *FleetBulkImportAccountConfigCloudNativeNetworksRequest) GetCloudNativeNetworks() []FleetAccountConfigCloudNativeNetworkOperation`

GetCloudNativeNetworks returns the CloudNativeNetworks field if non-nil, zero value otherwise.

### GetCloudNativeNetworksOk

`func (o *FleetBulkImportAccountConfigCloudNativeNetworksRequest) GetCloudNativeNetworksOk() (*[]FleetAccountConfigCloudNativeNetworkOperation, bool)`

GetCloudNativeNetworksOk returns a tuple with the CloudNativeNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworks

`func (o *FleetBulkImportAccountConfigCloudNativeNetworksRequest) SetCloudNativeNetworks(v []FleetAccountConfigCloudNativeNetworkOperation)`

SetCloudNativeNetworks sets CloudNativeNetworks field to given value.


### GetId

`func (o *FleetBulkImportAccountConfigCloudNativeNetworksRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetBulkImportAccountConfigCloudNativeNetworksRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetBulkImportAccountConfigCloudNativeNetworksRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *FleetBulkImportAccountConfigCloudNativeNetworksRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetBulkImportAccountConfigCloudNativeNetworksRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetBulkImportAccountConfigCloudNativeNetworksRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


