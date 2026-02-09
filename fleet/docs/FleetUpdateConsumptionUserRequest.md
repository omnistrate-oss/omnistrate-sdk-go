# FleetUpdateConsumptionUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]string** | Additional attributes of the user. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**UserId** | **string** | ID of a User | 

## Methods

### NewFleetUpdateConsumptionUserRequest

`func NewFleetUpdateConsumptionUserRequest(token string, userId string, ) *FleetUpdateConsumptionUserRequest`

NewFleetUpdateConsumptionUserRequest instantiates a new FleetUpdateConsumptionUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateConsumptionUserRequestWithDefaults

`func NewFleetUpdateConsumptionUserRequestWithDefaults() *FleetUpdateConsumptionUserRequest`

NewFleetUpdateConsumptionUserRequestWithDefaults instantiates a new FleetUpdateConsumptionUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *FleetUpdateConsumptionUserRequest) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FleetUpdateConsumptionUserRequest) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FleetUpdateConsumptionUserRequest) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *FleetUpdateConsumptionUserRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetToken

`func (o *FleetUpdateConsumptionUserRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetUpdateConsumptionUserRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetUpdateConsumptionUserRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserId

`func (o *FleetUpdateConsumptionUserRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FleetUpdateConsumptionUserRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FleetUpdateConsumptionUserRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


