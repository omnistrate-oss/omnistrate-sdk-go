# DeleteEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityID** | **string** | Unique identifier of the entity to delete | 
**EntityType** | **string** | Type of the entity (e.g., NAMESPACE, SERVICE, POD, etc.) | 
**HostClusterID** | **string** | ID of a Host Cluster | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDeleteEntityRequest

`func NewDeleteEntityRequest(entityID string, entityType string, hostClusterID string, token string, ) *DeleteEntityRequest`

NewDeleteEntityRequest instantiates a new DeleteEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteEntityRequestWithDefaults

`func NewDeleteEntityRequestWithDefaults() *DeleteEntityRequest`

NewDeleteEntityRequestWithDefaults instantiates a new DeleteEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityID

`func (o *DeleteEntityRequest) GetEntityID() string`

GetEntityID returns the EntityID field if non-nil, zero value otherwise.

### GetEntityIDOk

`func (o *DeleteEntityRequest) GetEntityIDOk() (*string, bool)`

GetEntityIDOk returns a tuple with the EntityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityID

`func (o *DeleteEntityRequest) SetEntityID(v string)`

SetEntityID sets EntityID field to given value.


### GetEntityType

`func (o *DeleteEntityRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DeleteEntityRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DeleteEntityRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetHostClusterID

`func (o *DeleteEntityRequest) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *DeleteEntityRequest) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *DeleteEntityRequest) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetToken

`func (o *DeleteEntityRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteEntityRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteEntityRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


