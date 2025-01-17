# DeploymentCellHealthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusterID** | Pointer to **string** | ID of a Host Cluster | [optional] 
**ServiceEnvironmentID** | Pointer to **string** | ID of a Service Environment | [optional] 
**ServiceID** | Pointer to **string** | ID of a Service | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDeploymentCellHealthRequest

`func NewDeploymentCellHealthRequest(token string, ) *DeploymentCellHealthRequest`

NewDeploymentCellHealthRequest instantiates a new DeploymentCellHealthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellHealthRequestWithDefaults

`func NewDeploymentCellHealthRequestWithDefaults() *DeploymentCellHealthRequest`

NewDeploymentCellHealthRequestWithDefaults instantiates a new DeploymentCellHealthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusterID

`func (o *DeploymentCellHealthRequest) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *DeploymentCellHealthRequest) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *DeploymentCellHealthRequest) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.

### HasHostClusterID

`func (o *DeploymentCellHealthRequest) HasHostClusterID() bool`

HasHostClusterID returns a boolean if a field has been set.

### GetServiceEnvironmentID

`func (o *DeploymentCellHealthRequest) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *DeploymentCellHealthRequest) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *DeploymentCellHealthRequest) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.

### HasServiceEnvironmentID

`func (o *DeploymentCellHealthRequest) HasServiceEnvironmentID() bool`

HasServiceEnvironmentID returns a boolean if a field has been set.

### GetServiceID

`func (o *DeploymentCellHealthRequest) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *DeploymentCellHealthRequest) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *DeploymentCellHealthRequest) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *DeploymentCellHealthRequest) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetToken

`func (o *DeploymentCellHealthRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeploymentCellHealthRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeploymentCellHealthRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


