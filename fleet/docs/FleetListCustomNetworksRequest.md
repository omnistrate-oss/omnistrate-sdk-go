# FleetListCustomNetworksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | Pointer to **string** | Name of the Infra Provider | [optional] 
**CloudProviderRegion** | Pointer to **string** | The region of the cloud provider that the instance is running in. | [optional] 
**CustomNetworksOnly** | Pointer to **bool** | Flag indicating whether to return only custom networks, or to include default and imported networks as well | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetListCustomNetworksRequest

`func NewFleetListCustomNetworksRequest(token string, ) *FleetListCustomNetworksRequest`

NewFleetListCustomNetworksRequest instantiates a new FleetListCustomNetworksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListCustomNetworksRequestWithDefaults

`func NewFleetListCustomNetworksRequestWithDefaults() *FleetListCustomNetworksRequest`

NewFleetListCustomNetworksRequestWithDefaults instantiates a new FleetListCustomNetworksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *FleetListCustomNetworksRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *FleetListCustomNetworksRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *FleetListCustomNetworksRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.

### HasCloudProviderName

`func (o *FleetListCustomNetworksRequest) HasCloudProviderName() bool`

HasCloudProviderName returns a boolean if a field has been set.

### GetCloudProviderRegion

`func (o *FleetListCustomNetworksRequest) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *FleetListCustomNetworksRequest) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *FleetListCustomNetworksRequest) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.

### HasCloudProviderRegion

`func (o *FleetListCustomNetworksRequest) HasCloudProviderRegion() bool`

HasCloudProviderRegion returns a boolean if a field has been set.

### GetCustomNetworksOnly

`func (o *FleetListCustomNetworksRequest) GetCustomNetworksOnly() bool`

GetCustomNetworksOnly returns the CustomNetworksOnly field if non-nil, zero value otherwise.

### GetCustomNetworksOnlyOk

`func (o *FleetListCustomNetworksRequest) GetCustomNetworksOnlyOk() (*bool, bool)`

GetCustomNetworksOnlyOk returns a tuple with the CustomNetworksOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworksOnly

`func (o *FleetListCustomNetworksRequest) SetCustomNetworksOnly(v bool)`

SetCustomNetworksOnly sets CustomNetworksOnly field to given value.

### HasCustomNetworksOnly

`func (o *FleetListCustomNetworksRequest) HasCustomNetworksOnly() bool`

HasCustomNetworksOnly returns a boolean if a field has been set.

### GetToken

`func (o *FleetListCustomNetworksRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetListCustomNetworksRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetListCustomNetworksRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


