# FleetDeleteUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | JWT token used to perform authorization | 
**UserId** | **string** | ID of a User | 

## Methods

### NewFleetDeleteUserRequest

`func NewFleetDeleteUserRequest(token string, userId string, ) *FleetDeleteUserRequest`

NewFleetDeleteUserRequest instantiates a new FleetDeleteUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDeleteUserRequestWithDefaults

`func NewFleetDeleteUserRequestWithDefaults() *FleetDeleteUserRequest`

NewFleetDeleteUserRequestWithDefaults instantiates a new FleetDeleteUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *FleetDeleteUserRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetDeleteUserRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetDeleteUserRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserId

`func (o *FleetDeleteUserRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FleetDeleteUserRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FleetDeleteUserRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


