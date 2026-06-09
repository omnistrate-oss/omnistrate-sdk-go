# DescribeProductTierWorkspaceArtifactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**Version** | Pointer to **string** | The product tier version to use. If omitted, the latest product tier workspace artifacts are used. | [optional] 
**WorkspaceArtifactId** | **string** | ID of a workspace artifact archive | 

## Methods

### NewDescribeProductTierWorkspaceArtifactsRequest

`func NewDescribeProductTierWorkspaceArtifactsRequest(id string, serviceId string, token string, workspaceArtifactId string, ) *DescribeProductTierWorkspaceArtifactsRequest`

NewDescribeProductTierWorkspaceArtifactsRequest instantiates a new DescribeProductTierWorkspaceArtifactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProductTierWorkspaceArtifactsRequestWithDefaults

`func NewDescribeProductTierWorkspaceArtifactsRequestWithDefaults() *DescribeProductTierWorkspaceArtifactsRequest`

NewDescribeProductTierWorkspaceArtifactsRequestWithDefaults instantiates a new DescribeProductTierWorkspaceArtifactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DescribeProductTierWorkspaceArtifactsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeProductTierWorkspaceArtifactsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeProductTierWorkspaceArtifactsRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *DescribeProductTierWorkspaceArtifactsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeProductTierWorkspaceArtifactsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeProductTierWorkspaceArtifactsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *DescribeProductTierWorkspaceArtifactsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeProductTierWorkspaceArtifactsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeProductTierWorkspaceArtifactsRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVersion

`func (o *DescribeProductTierWorkspaceArtifactsRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DescribeProductTierWorkspaceArtifactsRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DescribeProductTierWorkspaceArtifactsRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DescribeProductTierWorkspaceArtifactsRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkspaceArtifactId

`func (o *DescribeProductTierWorkspaceArtifactsRequest) GetWorkspaceArtifactId() string`

GetWorkspaceArtifactId returns the WorkspaceArtifactId field if non-nil, zero value otherwise.

### GetWorkspaceArtifactIdOk

`func (o *DescribeProductTierWorkspaceArtifactsRequest) GetWorkspaceArtifactIdOk() (*string, bool)`

GetWorkspaceArtifactIdOk returns a tuple with the WorkspaceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceArtifactId

`func (o *DescribeProductTierWorkspaceArtifactsRequest) SetWorkspaceArtifactId(v string)`

SetWorkspaceArtifactId sets WorkspaceArtifactId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


