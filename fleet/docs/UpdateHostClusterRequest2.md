# UpdateHostClusterRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingAmenities** | Pointer to [**[]Amenity**](Amenity.md) | The pending amenities for the host cluster | [optional] 
**SyncWithOrgTemplate** | Pointer to **bool** | Whether to sync the host cluster with the org template | [optional] 

## Methods

### NewUpdateHostClusterRequest2

`func NewUpdateHostClusterRequest2() *UpdateHostClusterRequest2`

NewUpdateHostClusterRequest2 instantiates a new UpdateHostClusterRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostClusterRequest2WithDefaults

`func NewUpdateHostClusterRequest2WithDefaults() *UpdateHostClusterRequest2`

NewUpdateHostClusterRequest2WithDefaults instantiates a new UpdateHostClusterRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingAmenities

`func (o *UpdateHostClusterRequest2) GetPendingAmenities() []Amenity`

GetPendingAmenities returns the PendingAmenities field if non-nil, zero value otherwise.

### GetPendingAmenitiesOk

`func (o *UpdateHostClusterRequest2) GetPendingAmenitiesOk() (*[]Amenity, bool)`

GetPendingAmenitiesOk returns a tuple with the PendingAmenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmenities

`func (o *UpdateHostClusterRequest2) SetPendingAmenities(v []Amenity)`

SetPendingAmenities sets PendingAmenities field to given value.

### HasPendingAmenities

`func (o *UpdateHostClusterRequest2) HasPendingAmenities() bool`

HasPendingAmenities returns a boolean if a field has been set.

### GetSyncWithOrgTemplate

`func (o *UpdateHostClusterRequest2) GetSyncWithOrgTemplate() bool`

GetSyncWithOrgTemplate returns the SyncWithOrgTemplate field if non-nil, zero value otherwise.

### GetSyncWithOrgTemplateOk

`func (o *UpdateHostClusterRequest2) GetSyncWithOrgTemplateOk() (*bool, bool)`

GetSyncWithOrgTemplateOk returns a tuple with the SyncWithOrgTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncWithOrgTemplate

`func (o *UpdateHostClusterRequest2) SetSyncWithOrgTemplate(v bool)`

SetSyncWithOrgTemplate sets SyncWithOrgTemplate field to given value.

### HasSyncWithOrgTemplate

`func (o *UpdateHostClusterRequest2) HasSyncWithOrgTemplate() bool`

HasSyncWithOrgTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


