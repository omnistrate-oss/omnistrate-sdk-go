# FleetListUsersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | The service environment ID this workflow belongs to. | [optional] 
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**ServiceId** | Pointer to **string** | The service ID this workflow belongs to. | [optional] 
**Users** | [**[]User**](User.md) | List of active users using the service. | 

## Methods

### NewFleetListUsersResult

`func NewFleetListUsersResult(users []User, ) *FleetListUsersResult`

NewFleetListUsersResult instantiates a new FleetListUsersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListUsersResultWithDefaults

`func NewFleetListUsersResultWithDefaults() *FleetListUsersResult`

NewFleetListUsersResultWithDefaults instantiates a new FleetListUsersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetListUsersResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetListUsersResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetListUsersResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *FleetListUsersResult) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetNextPageToken

`func (o *FleetListUsersResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FleetListUsersResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FleetListUsersResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FleetListUsersResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetListUsersResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetListUsersResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetListUsersResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *FleetListUsersResult) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetUsers

`func (o *FleetListUsersResult) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FleetListUsersResult) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FleetListUsersResult) SetUsers(v []User)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


