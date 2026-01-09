# FleetListAllInstanceSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ServiceId** | **string** | ID of a Service | 
**SnapshotType** | Pointer to **string** | The snapshot type to filter by | [optional] 
**SubscriptionId** | Pointer to **string** | ID of a Subscription | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetListAllInstanceSnapshotRequest

`func NewFleetListAllInstanceSnapshotRequest(environmentId string, serviceId string, token string, ) *FleetListAllInstanceSnapshotRequest`

NewFleetListAllInstanceSnapshotRequest instantiates a new FleetListAllInstanceSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListAllInstanceSnapshotRequestWithDefaults

`func NewFleetListAllInstanceSnapshotRequestWithDefaults() *FleetListAllInstanceSnapshotRequest`

NewFleetListAllInstanceSnapshotRequestWithDefaults instantiates a new FleetListAllInstanceSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetListAllInstanceSnapshotRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetListAllInstanceSnapshotRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetListAllInstanceSnapshotRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetProductTierId

`func (o *FleetListAllInstanceSnapshotRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *FleetListAllInstanceSnapshotRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *FleetListAllInstanceSnapshotRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *FleetListAllInstanceSnapshotRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetListAllInstanceSnapshotRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetListAllInstanceSnapshotRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetListAllInstanceSnapshotRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSnapshotType

`func (o *FleetListAllInstanceSnapshotRequest) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *FleetListAllInstanceSnapshotRequest) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *FleetListAllInstanceSnapshotRequest) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *FleetListAllInstanceSnapshotRequest) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *FleetListAllInstanceSnapshotRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *FleetListAllInstanceSnapshotRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *FleetListAllInstanceSnapshotRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *FleetListAllInstanceSnapshotRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetToken

`func (o *FleetListAllInstanceSnapshotRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetListAllInstanceSnapshotRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetListAllInstanceSnapshotRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


