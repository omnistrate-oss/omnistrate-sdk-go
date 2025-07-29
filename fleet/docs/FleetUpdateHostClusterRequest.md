# FleetUpdateHostClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**Id** | **string** | ID of a Host Cluster | 
**PendingAmenities** | Pointer to [**[]Amenity**](Amenity.md) | The pending amenities for the host cluster | [optional] 
**ServiceId** | **string** | ID of a Service | 
**SyncWithOrgTemplate** | Pointer to **bool** | Whether to sync the host cluster with the org template | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetUpdateHostClusterRequest

`func NewFleetUpdateHostClusterRequest(environmentId string, id string, serviceId string, token string, ) *FleetUpdateHostClusterRequest`

NewFleetUpdateHostClusterRequest instantiates a new FleetUpdateHostClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateHostClusterRequestWithDefaults

`func NewFleetUpdateHostClusterRequestWithDefaults() *FleetUpdateHostClusterRequest`

NewFleetUpdateHostClusterRequestWithDefaults instantiates a new FleetUpdateHostClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetUpdateHostClusterRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetUpdateHostClusterRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetUpdateHostClusterRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *FleetUpdateHostClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetUpdateHostClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetUpdateHostClusterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPendingAmenities

`func (o *FleetUpdateHostClusterRequest) GetPendingAmenities() []Amenity`

GetPendingAmenities returns the PendingAmenities field if non-nil, zero value otherwise.

### GetPendingAmenitiesOk

`func (o *FleetUpdateHostClusterRequest) GetPendingAmenitiesOk() (*[]Amenity, bool)`

GetPendingAmenitiesOk returns a tuple with the PendingAmenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmenities

`func (o *FleetUpdateHostClusterRequest) SetPendingAmenities(v []Amenity)`

SetPendingAmenities sets PendingAmenities field to given value.

### HasPendingAmenities

`func (o *FleetUpdateHostClusterRequest) HasPendingAmenities() bool`

HasPendingAmenities returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetUpdateHostClusterRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetUpdateHostClusterRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetUpdateHostClusterRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSyncWithOrgTemplate

`func (o *FleetUpdateHostClusterRequest) GetSyncWithOrgTemplate() bool`

GetSyncWithOrgTemplate returns the SyncWithOrgTemplate field if non-nil, zero value otherwise.

### GetSyncWithOrgTemplateOk

`func (o *FleetUpdateHostClusterRequest) GetSyncWithOrgTemplateOk() (*bool, bool)`

GetSyncWithOrgTemplateOk returns a tuple with the SyncWithOrgTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncWithOrgTemplate

`func (o *FleetUpdateHostClusterRequest) SetSyncWithOrgTemplate(v bool)`

SetSyncWithOrgTemplate sets SyncWithOrgTemplate field to given value.

### HasSyncWithOrgTemplate

`func (o *FleetUpdateHostClusterRequest) HasSyncWithOrgTemplate() bool`

HasSyncWithOrgTemplate returns a boolean if a field has been set.

### GetToken

`func (o *FleetUpdateHostClusterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetUpdateHostClusterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetUpdateHostClusterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


