# ListHostClustersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigId** | Pointer to **string** | ID of an Account Config | [optional] 
**IncludeProvisionerClusters** | Pointer to **bool** | Include provisioner clusters in the response | [optional] 
**RegionId** | Pointer to **string** | ID of a Region | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListHostClustersRequest

`func NewListHostClustersRequest(token string, ) *ListHostClustersRequest`

NewListHostClustersRequest instantiates a new ListHostClustersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHostClustersRequestWithDefaults

`func NewListHostClustersRequestWithDefaults() *ListHostClustersRequest`

NewListHostClustersRequestWithDefaults instantiates a new ListHostClustersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigId

`func (o *ListHostClustersRequest) GetAccountConfigId() string`

GetAccountConfigId returns the AccountConfigId field if non-nil, zero value otherwise.

### GetAccountConfigIdOk

`func (o *ListHostClustersRequest) GetAccountConfigIdOk() (*string, bool)`

GetAccountConfigIdOk returns a tuple with the AccountConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigId

`func (o *ListHostClustersRequest) SetAccountConfigId(v string)`

SetAccountConfigId sets AccountConfigId field to given value.

### HasAccountConfigId

`func (o *ListHostClustersRequest) HasAccountConfigId() bool`

HasAccountConfigId returns a boolean if a field has been set.

### GetIncludeProvisionerClusters

`func (o *ListHostClustersRequest) GetIncludeProvisionerClusters() bool`

GetIncludeProvisionerClusters returns the IncludeProvisionerClusters field if non-nil, zero value otherwise.

### GetIncludeProvisionerClustersOk

`func (o *ListHostClustersRequest) GetIncludeProvisionerClustersOk() (*bool, bool)`

GetIncludeProvisionerClustersOk returns a tuple with the IncludeProvisionerClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProvisionerClusters

`func (o *ListHostClustersRequest) SetIncludeProvisionerClusters(v bool)`

SetIncludeProvisionerClusters sets IncludeProvisionerClusters field to given value.

### HasIncludeProvisionerClusters

`func (o *ListHostClustersRequest) HasIncludeProvisionerClusters() bool`

HasIncludeProvisionerClusters returns a boolean if a field has been set.

### GetRegionId

`func (o *ListHostClustersRequest) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ListHostClustersRequest) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ListHostClustersRequest) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ListHostClustersRequest) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetToken

`func (o *ListHostClustersRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListHostClustersRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListHostClustersRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


