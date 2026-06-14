# FleetImportAccountConfigCloudNativeNetworkHostClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworkId** | **string** | The cloud provider network ID (e.g. AWS VPC ID) that contains the host cluster to import | 
**HostClusterName** | **string** | The cloud provider host cluster name to import from this cloud native network | 
**Id** | **string** | ID of an Account Config | 
**Region** | **string** | The deployment region where the host cluster resides | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetImportAccountConfigCloudNativeNetworkHostClusterRequest

`func NewFleetImportAccountConfigCloudNativeNetworkHostClusterRequest(cloudNativeNetworkId string, hostClusterName string, id string, region string, token string, ) *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest`

NewFleetImportAccountConfigCloudNativeNetworkHostClusterRequest instantiates a new FleetImportAccountConfigCloudNativeNetworkHostClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetImportAccountConfigCloudNativeNetworkHostClusterRequestWithDefaults

`func NewFleetImportAccountConfigCloudNativeNetworkHostClusterRequestWithDefaults() *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest`

NewFleetImportAccountConfigCloudNativeNetworkHostClusterRequestWithDefaults instantiates a new FleetImportAccountConfigCloudNativeNetworkHostClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworkId

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) GetCloudNativeNetworkId() string`

GetCloudNativeNetworkId returns the CloudNativeNetworkId field if non-nil, zero value otherwise.

### GetCloudNativeNetworkIdOk

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) GetCloudNativeNetworkIdOk() (*string, bool)`

GetCloudNativeNetworkIdOk returns a tuple with the CloudNativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworkId

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) SetCloudNativeNetworkId(v string)`

SetCloudNativeNetworkId sets CloudNativeNetworkId field to given value.


### GetHostClusterName

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) GetHostClusterName() string`

GetHostClusterName returns the HostClusterName field if non-nil, zero value otherwise.

### GetHostClusterNameOk

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) GetHostClusterNameOk() (*string, bool)`

GetHostClusterNameOk returns a tuple with the HostClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterName

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) SetHostClusterName(v string)`

SetHostClusterName sets HostClusterName field to given value.


### GetId

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetRegion

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetToken

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetImportAccountConfigCloudNativeNetworkHostClusterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


