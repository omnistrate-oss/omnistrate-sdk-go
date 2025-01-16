# InviteConsumptionUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**RoleType** | **string** | Type of the role | 
**SubscriptionId** | **string** | The subscription ID | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewInviteConsumptionUserRequest

`func NewInviteConsumptionUserRequest(email string, roleType string, subscriptionId string, token string, ) *InviteConsumptionUserRequest`

NewInviteConsumptionUserRequest instantiates a new InviteConsumptionUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteConsumptionUserRequestWithDefaults

`func NewInviteConsumptionUserRequestWithDefaults() *InviteConsumptionUserRequest`

NewInviteConsumptionUserRequestWithDefaults instantiates a new InviteConsumptionUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteConsumptionUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteConsumptionUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteConsumptionUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoleType

`func (o *InviteConsumptionUserRequest) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *InviteConsumptionUserRequest) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *InviteConsumptionUserRequest) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.


### GetSubscriptionId

`func (o *InviteConsumptionUserRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *InviteConsumptionUserRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *InviteConsumptionUserRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetToken

`func (o *InviteConsumptionUserRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InviteConsumptionUserRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InviteConsumptionUserRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


