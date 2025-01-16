# ListCustomNetworksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | Pointer to **string** | Name of the Infra Provider | [optional] 
**CloudProviderRegion** | Pointer to **string** | The region of the cloud provider that the instance is running in. | [optional] 
**CustomNetworksOnly** | Pointer to **bool** | Flag indicating whether to return only custom networks, or to include default and imported networks as well | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListCustomNetworksRequest

`func NewListCustomNetworksRequest(token string, ) *ListCustomNetworksRequest`

NewListCustomNetworksRequest instantiates a new ListCustomNetworksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomNetworksRequestWithDefaults

`func NewListCustomNetworksRequestWithDefaults() *ListCustomNetworksRequest`

NewListCustomNetworksRequestWithDefaults instantiates a new ListCustomNetworksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *ListCustomNetworksRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *ListCustomNetworksRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *ListCustomNetworksRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.

### HasCloudProviderName

`func (o *ListCustomNetworksRequest) HasCloudProviderName() bool`

HasCloudProviderName returns a boolean if a field has been set.

### GetCloudProviderRegion

`func (o *ListCustomNetworksRequest) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *ListCustomNetworksRequest) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *ListCustomNetworksRequest) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.

### HasCloudProviderRegion

`func (o *ListCustomNetworksRequest) HasCloudProviderRegion() bool`

HasCloudProviderRegion returns a boolean if a field has been set.

### GetCustomNetworksOnly

`func (o *ListCustomNetworksRequest) GetCustomNetworksOnly() bool`

GetCustomNetworksOnly returns the CustomNetworksOnly field if non-nil, zero value otherwise.

### GetCustomNetworksOnlyOk

`func (o *ListCustomNetworksRequest) GetCustomNetworksOnlyOk() (*bool, bool)`

GetCustomNetworksOnlyOk returns a tuple with the CustomNetworksOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworksOnly

`func (o *ListCustomNetworksRequest) SetCustomNetworksOnly(v bool)`

SetCustomNetworksOnly sets CustomNetworksOnly field to given value.

### HasCustomNetworksOnly

`func (o *ListCustomNetworksRequest) HasCustomNetworksOnly() bool`

HasCustomNetworksOnly returns a boolean if a field has been set.

### GetToken

`func (o *ListCustomNetworksRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListCustomNetworksRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListCustomNetworksRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


