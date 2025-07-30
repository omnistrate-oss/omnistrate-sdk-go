# UpdateHostClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Host Cluster | 
**PendingAmenities** | Pointer to [**[]Amenity**](Amenity.md) | The pending amenities for the host cluster | [optional] 
**SyncWithOrgTemplate** | Pointer to **bool** | Whether to sync the host cluster with the org template | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateHostClusterRequest

`func NewUpdateHostClusterRequest(id string, token string, ) *UpdateHostClusterRequest`

NewUpdateHostClusterRequest instantiates a new UpdateHostClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostClusterRequestWithDefaults

`func NewUpdateHostClusterRequestWithDefaults() *UpdateHostClusterRequest`

NewUpdateHostClusterRequestWithDefaults instantiates a new UpdateHostClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateHostClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateHostClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateHostClusterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPendingAmenities

`func (o *UpdateHostClusterRequest) GetPendingAmenities() []Amenity`

GetPendingAmenities returns the PendingAmenities field if non-nil, zero value otherwise.

### GetPendingAmenitiesOk

`func (o *UpdateHostClusterRequest) GetPendingAmenitiesOk() (*[]Amenity, bool)`

GetPendingAmenitiesOk returns a tuple with the PendingAmenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmenities

`func (o *UpdateHostClusterRequest) SetPendingAmenities(v []Amenity)`

SetPendingAmenities sets PendingAmenities field to given value.

### HasPendingAmenities

`func (o *UpdateHostClusterRequest) HasPendingAmenities() bool`

HasPendingAmenities returns a boolean if a field has been set.

### GetSyncWithOrgTemplate

`func (o *UpdateHostClusterRequest) GetSyncWithOrgTemplate() bool`

GetSyncWithOrgTemplate returns the SyncWithOrgTemplate field if non-nil, zero value otherwise.

### GetSyncWithOrgTemplateOk

`func (o *UpdateHostClusterRequest) GetSyncWithOrgTemplateOk() (*bool, bool)`

GetSyncWithOrgTemplateOk returns a tuple with the SyncWithOrgTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncWithOrgTemplate

`func (o *UpdateHostClusterRequest) SetSyncWithOrgTemplate(v bool)`

SetSyncWithOrgTemplate sets SyncWithOrgTemplate field to given value.

### HasSyncWithOrgTemplate

`func (o *UpdateHostClusterRequest) HasSyncWithOrgTemplate() bool`

HasSyncWithOrgTemplate returns a boolean if a field has been set.

### GetToken

`func (o *UpdateHostClusterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateHostClusterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateHostClusterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


