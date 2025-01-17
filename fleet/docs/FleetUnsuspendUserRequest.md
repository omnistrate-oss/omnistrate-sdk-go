# FleetUnsuspendUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | JWT token used to perform authorization | 
**UserId** | **string** | ID of a User | 

## Methods

### NewFleetUnsuspendUserRequest

`func NewFleetUnsuspendUserRequest(token string, userId string, ) *FleetUnsuspendUserRequest`

NewFleetUnsuspendUserRequest instantiates a new FleetUnsuspendUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUnsuspendUserRequestWithDefaults

`func NewFleetUnsuspendUserRequestWithDefaults() *FleetUnsuspendUserRequest`

NewFleetUnsuspendUserRequestWithDefaults instantiates a new FleetUnsuspendUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *FleetUnsuspendUserRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetUnsuspendUserRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetUnsuspendUserRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserId

`func (o *FleetUnsuspendUserRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FleetUnsuspendUserRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FleetUnsuspendUserRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


