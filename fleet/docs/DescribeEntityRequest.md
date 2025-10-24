# DescribeEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityID** | **string** | Unique identifier of the entity to describe | 
**EntityType** | **string** | Type of the entity (e.g., NAMESPACE, SERVICE, POD, etc.) | 
**HostClusterID** | **string** | ID of a Host Cluster | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeEntityRequest

`func NewDescribeEntityRequest(entityID string, entityType string, hostClusterID string, token string, ) *DescribeEntityRequest`

NewDescribeEntityRequest instantiates a new DescribeEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeEntityRequestWithDefaults

`func NewDescribeEntityRequestWithDefaults() *DescribeEntityRequest`

NewDescribeEntityRequestWithDefaults instantiates a new DescribeEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityID

`func (o *DescribeEntityRequest) GetEntityID() string`

GetEntityID returns the EntityID field if non-nil, zero value otherwise.

### GetEntityIDOk

`func (o *DescribeEntityRequest) GetEntityIDOk() (*string, bool)`

GetEntityIDOk returns a tuple with the EntityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityID

`func (o *DescribeEntityRequest) SetEntityID(v string)`

SetEntityID sets EntityID field to given value.


### GetEntityType

`func (o *DescribeEntityRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DescribeEntityRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DescribeEntityRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetHostClusterID

`func (o *DescribeEntityRequest) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *DescribeEntityRequest) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *DescribeEntityRequest) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetToken

`func (o *DescribeEntityRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeEntityRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeEntityRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


