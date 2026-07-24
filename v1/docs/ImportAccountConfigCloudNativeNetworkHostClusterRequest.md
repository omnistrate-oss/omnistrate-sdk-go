# ImportAccountConfigCloudNativeNetworkHostClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworkId** | **string** | The cloud provider network ID (e.g. AWS VPC ID) that contains the host cluster to import | 
**HostClusterName** | **string** | The cloud provider host cluster name to import from this cloud native network | 
**Id** | **string** | ID of an Account Config | 
**Region** | **string** | The deployment region where the host cluster resides | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewImportAccountConfigCloudNativeNetworkHostClusterRequest

`func NewImportAccountConfigCloudNativeNetworkHostClusterRequest(cloudNativeNetworkId string, hostClusterName string, id string, region string, token string, ) *ImportAccountConfigCloudNativeNetworkHostClusterRequest`

NewImportAccountConfigCloudNativeNetworkHostClusterRequest instantiates a new ImportAccountConfigCloudNativeNetworkHostClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportAccountConfigCloudNativeNetworkHostClusterRequestWithDefaults

`func NewImportAccountConfigCloudNativeNetworkHostClusterRequestWithDefaults() *ImportAccountConfigCloudNativeNetworkHostClusterRequest`

NewImportAccountConfigCloudNativeNetworkHostClusterRequestWithDefaults instantiates a new ImportAccountConfigCloudNativeNetworkHostClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworkId

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) GetCloudNativeNetworkId() string`

GetCloudNativeNetworkId returns the CloudNativeNetworkId field if non-nil, zero value otherwise.

### GetCloudNativeNetworkIdOk

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) GetCloudNativeNetworkIdOk() (*string, bool)`

GetCloudNativeNetworkIdOk returns a tuple with the CloudNativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworkId

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) SetCloudNativeNetworkId(v string)`

SetCloudNativeNetworkId sets CloudNativeNetworkId field to given value.


### GetHostClusterName

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) GetHostClusterName() string`

GetHostClusterName returns the HostClusterName field if non-nil, zero value otherwise.

### GetHostClusterNameOk

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) GetHostClusterNameOk() (*string, bool)`

GetHostClusterNameOk returns a tuple with the HostClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterName

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) SetHostClusterName(v string)`

SetHostClusterName sets HostClusterName field to given value.


### GetId

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetRegion

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetToken

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ImportAccountConfigCloudNativeNetworkHostClusterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


