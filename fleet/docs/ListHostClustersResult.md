# ListHostClustersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusters** | [**[]HostCluster**](HostCluster.md) |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListHostClustersResult

`func NewListHostClustersResult(hostClusters []HostCluster, token string, ) *ListHostClustersResult`

NewListHostClustersResult instantiates a new ListHostClustersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHostClustersResultWithDefaults

`func NewListHostClustersResultWithDefaults() *ListHostClustersResult`

NewListHostClustersResultWithDefaults instantiates a new ListHostClustersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusters

`func (o *ListHostClustersResult) GetHostClusters() []HostCluster`

GetHostClusters returns the HostClusters field if non-nil, zero value otherwise.

### GetHostClustersOk

`func (o *ListHostClustersResult) GetHostClustersOk() (*[]HostCluster, bool)`

GetHostClustersOk returns a tuple with the HostClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusters

`func (o *ListHostClustersResult) SetHostClusters(v []HostCluster)`

SetHostClusters sets HostClusters field to given value.


### GetToken

`func (o *ListHostClustersResult) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListHostClustersResult) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListHostClustersResult) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


