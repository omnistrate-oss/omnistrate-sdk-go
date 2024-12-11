# FleetListAllUsersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**ServiceId** | Pointer to **string** | The service ID of the users. | [optional] 
**Users** | [**[]AccessSideUser**](AccessSideUser.md) | List of access-side users. | 

## Methods

### NewFleetListAllUsersResult

`func NewFleetListAllUsersResult(users []AccessSideUser, ) *FleetListAllUsersResult`

NewFleetListAllUsersResult instantiates a new FleetListAllUsersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListAllUsersResultWithDefaults

`func NewFleetListAllUsersResultWithDefaults() *FleetListAllUsersResult`

NewFleetListAllUsersResultWithDefaults instantiates a new FleetListAllUsersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *FleetListAllUsersResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FleetListAllUsersResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FleetListAllUsersResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FleetListAllUsersResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetListAllUsersResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetListAllUsersResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetListAllUsersResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *FleetListAllUsersResult) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetUsers

`func (o *FleetListAllUsersResult) GetUsers() []AccessSideUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FleetListAllUsersResult) GetUsersOk() (*[]AccessSideUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FleetListAllUsersResult) SetUsers(v []AccessSideUser)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


