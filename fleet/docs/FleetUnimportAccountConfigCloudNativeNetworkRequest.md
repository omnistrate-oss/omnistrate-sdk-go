# FleetUnimportAccountConfigCloudNativeNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworkId** | **string** | The cloud provider network ID (e.g. AWS VPC ID) to unimport. Rejected with HTTP 400 if the network is currently in use by a host cluster. | 
**Id** | **string** | ID of an Account Config | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetUnimportAccountConfigCloudNativeNetworkRequest

`func NewFleetUnimportAccountConfigCloudNativeNetworkRequest(cloudNativeNetworkId string, id string, token string, ) *FleetUnimportAccountConfigCloudNativeNetworkRequest`

NewFleetUnimportAccountConfigCloudNativeNetworkRequest instantiates a new FleetUnimportAccountConfigCloudNativeNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUnimportAccountConfigCloudNativeNetworkRequestWithDefaults

`func NewFleetUnimportAccountConfigCloudNativeNetworkRequestWithDefaults() *FleetUnimportAccountConfigCloudNativeNetworkRequest`

NewFleetUnimportAccountConfigCloudNativeNetworkRequestWithDefaults instantiates a new FleetUnimportAccountConfigCloudNativeNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworkId

`func (o *FleetUnimportAccountConfigCloudNativeNetworkRequest) GetCloudNativeNetworkId() string`

GetCloudNativeNetworkId returns the CloudNativeNetworkId field if non-nil, zero value otherwise.

### GetCloudNativeNetworkIdOk

`func (o *FleetUnimportAccountConfigCloudNativeNetworkRequest) GetCloudNativeNetworkIdOk() (*string, bool)`

GetCloudNativeNetworkIdOk returns a tuple with the CloudNativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworkId

`func (o *FleetUnimportAccountConfigCloudNativeNetworkRequest) SetCloudNativeNetworkId(v string)`

SetCloudNativeNetworkId sets CloudNativeNetworkId field to given value.


### GetId

`func (o *FleetUnimportAccountConfigCloudNativeNetworkRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetUnimportAccountConfigCloudNativeNetworkRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetUnimportAccountConfigCloudNativeNetworkRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *FleetUnimportAccountConfigCloudNativeNetworkRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetUnimportAccountConfigCloudNativeNetworkRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetUnimportAccountConfigCloudNativeNetworkRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


