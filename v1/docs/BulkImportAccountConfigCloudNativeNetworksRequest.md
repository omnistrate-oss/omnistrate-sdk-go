# BulkImportAccountConfigCloudNativeNetworksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworks** | [**[]AccountConfigCloudNativeNetworkOperation**](AccountConfigCloudNativeNetworkOperation.md) | List of cloud native network operations to perform | 
**Id** | **string** | ID of an Account Config | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewBulkImportAccountConfigCloudNativeNetworksRequest

`func NewBulkImportAccountConfigCloudNativeNetworksRequest(cloudNativeNetworks []AccountConfigCloudNativeNetworkOperation, id string, token string, ) *BulkImportAccountConfigCloudNativeNetworksRequest`

NewBulkImportAccountConfigCloudNativeNetworksRequest instantiates a new BulkImportAccountConfigCloudNativeNetworksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkImportAccountConfigCloudNativeNetworksRequestWithDefaults

`func NewBulkImportAccountConfigCloudNativeNetworksRequestWithDefaults() *BulkImportAccountConfigCloudNativeNetworksRequest`

NewBulkImportAccountConfigCloudNativeNetworksRequestWithDefaults instantiates a new BulkImportAccountConfigCloudNativeNetworksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworks

`func (o *BulkImportAccountConfigCloudNativeNetworksRequest) GetCloudNativeNetworks() []AccountConfigCloudNativeNetworkOperation`

GetCloudNativeNetworks returns the CloudNativeNetworks field if non-nil, zero value otherwise.

### GetCloudNativeNetworksOk

`func (o *BulkImportAccountConfigCloudNativeNetworksRequest) GetCloudNativeNetworksOk() (*[]AccountConfigCloudNativeNetworkOperation, bool)`

GetCloudNativeNetworksOk returns a tuple with the CloudNativeNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworks

`func (o *BulkImportAccountConfigCloudNativeNetworksRequest) SetCloudNativeNetworks(v []AccountConfigCloudNativeNetworkOperation)`

SetCloudNativeNetworks sets CloudNativeNetworks field to given value.


### GetId

`func (o *BulkImportAccountConfigCloudNativeNetworksRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkImportAccountConfigCloudNativeNetworksRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkImportAccountConfigCloudNativeNetworksRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *BulkImportAccountConfigCloudNativeNetworksRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BulkImportAccountConfigCloudNativeNetworksRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BulkImportAccountConfigCloudNativeNetworksRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


