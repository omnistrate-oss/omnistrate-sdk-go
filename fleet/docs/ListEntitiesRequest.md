# ListEntitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | Type of entities to list (e.g., NAMESPACE, SERVICE, POD, etc.) | 
**HostClusterID** | **string** | ID of a Host Cluster | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListEntitiesRequest

`func NewListEntitiesRequest(entityType string, hostClusterID string, token string, ) *ListEntitiesRequest`

NewListEntitiesRequest instantiates a new ListEntitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntitiesRequestWithDefaults

`func NewListEntitiesRequestWithDefaults() *ListEntitiesRequest`

NewListEntitiesRequestWithDefaults instantiates a new ListEntitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *ListEntitiesRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ListEntitiesRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ListEntitiesRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetHostClusterID

`func (o *ListEntitiesRequest) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *ListEntitiesRequest) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *ListEntitiesRequest) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetToken

`func (o *ListEntitiesRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListEntitiesRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListEntitiesRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


