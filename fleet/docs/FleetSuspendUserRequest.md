# FleetSuspendUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | JWT token used to perform authorization | 
**UserId** | **string** | ID of a User | 

## Methods

### NewFleetSuspendUserRequest

`func NewFleetSuspendUserRequest(token string, userId string, ) *FleetSuspendUserRequest`

NewFleetSuspendUserRequest instantiates a new FleetSuspendUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetSuspendUserRequestWithDefaults

`func NewFleetSuspendUserRequestWithDefaults() *FleetSuspendUserRequest`

NewFleetSuspendUserRequestWithDefaults instantiates a new FleetSuspendUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *FleetSuspendUserRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetSuspendUserRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetSuspendUserRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserId

`func (o *FleetSuspendUserRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FleetSuspendUserRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FleetSuspendUserRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


