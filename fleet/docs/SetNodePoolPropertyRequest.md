# SetNodePoolPropertyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusterID** | **string** | ID of a Host Cluster | 
**MaxSize** | **int64** | New maximum size for the node pool | 
**NodePoolName** | **string** | Unique identifier of the node pool to update | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewSetNodePoolPropertyRequest

`func NewSetNodePoolPropertyRequest(hostClusterID string, maxSize int64, nodePoolName string, token string, ) *SetNodePoolPropertyRequest`

NewSetNodePoolPropertyRequest instantiates a new SetNodePoolPropertyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNodePoolPropertyRequestWithDefaults

`func NewSetNodePoolPropertyRequestWithDefaults() *SetNodePoolPropertyRequest`

NewSetNodePoolPropertyRequestWithDefaults instantiates a new SetNodePoolPropertyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusterID

`func (o *SetNodePoolPropertyRequest) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *SetNodePoolPropertyRequest) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *SetNodePoolPropertyRequest) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetMaxSize

`func (o *SetNodePoolPropertyRequest) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *SetNodePoolPropertyRequest) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *SetNodePoolPropertyRequest) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.


### GetNodePoolName

`func (o *SetNodePoolPropertyRequest) GetNodePoolName() string`

GetNodePoolName returns the NodePoolName field if non-nil, zero value otherwise.

### GetNodePoolNameOk

`func (o *SetNodePoolPropertyRequest) GetNodePoolNameOk() (*string, bool)`

GetNodePoolNameOk returns a tuple with the NodePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePoolName

`func (o *SetNodePoolPropertyRequest) SetNodePoolName(v string)`

SetNodePoolName sets NodePoolName field to given value.


### GetToken

`func (o *SetNodePoolPropertyRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SetNodePoolPropertyRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SetNodePoolPropertyRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


