# FleetImportAccountConfigCloudNativeNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworkId** | **string** | The provider-native network ID to import for deployments | 
**Id** | **string** | ID of an Account Config | 
**Region** | **string** | The deployment region whose validated subnetworks should be imported | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetImportAccountConfigCloudNativeNetworkRequest

`func NewFleetImportAccountConfigCloudNativeNetworkRequest(cloudNativeNetworkId string, id string, region string, token string, ) *FleetImportAccountConfigCloudNativeNetworkRequest`

NewFleetImportAccountConfigCloudNativeNetworkRequest instantiates a new FleetImportAccountConfigCloudNativeNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetImportAccountConfigCloudNativeNetworkRequestWithDefaults

`func NewFleetImportAccountConfigCloudNativeNetworkRequestWithDefaults() *FleetImportAccountConfigCloudNativeNetworkRequest`

NewFleetImportAccountConfigCloudNativeNetworkRequestWithDefaults instantiates a new FleetImportAccountConfigCloudNativeNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworkId

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) GetCloudNativeNetworkId() string`

GetCloudNativeNetworkId returns the CloudNativeNetworkId field if non-nil, zero value otherwise.

### GetCloudNativeNetworkIdOk

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) GetCloudNativeNetworkIdOk() (*string, bool)`

GetCloudNativeNetworkIdOk returns a tuple with the CloudNativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworkId

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) SetCloudNativeNetworkId(v string)`

SetCloudNativeNetworkId sets CloudNativeNetworkId field to given value.


### GetId

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) SetId(v string)`

SetId sets Id field to given value.


### GetRegion

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetToken

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetImportAccountConfigCloudNativeNetworkRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


