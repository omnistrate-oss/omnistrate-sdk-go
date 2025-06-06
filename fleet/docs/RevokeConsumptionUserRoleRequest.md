# RevokeConsumptionUserRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**RoleType** | **string** | Type of the role | 
**SubscriptionId** | **string** | The subscription ID | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRevokeConsumptionUserRoleRequest

`func NewRevokeConsumptionUserRoleRequest(email string, roleType string, subscriptionId string, token string, ) *RevokeConsumptionUserRoleRequest`

NewRevokeConsumptionUserRoleRequest instantiates a new RevokeConsumptionUserRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeConsumptionUserRoleRequestWithDefaults

`func NewRevokeConsumptionUserRoleRequestWithDefaults() *RevokeConsumptionUserRoleRequest`

NewRevokeConsumptionUserRoleRequestWithDefaults instantiates a new RevokeConsumptionUserRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RevokeConsumptionUserRoleRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RevokeConsumptionUserRoleRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RevokeConsumptionUserRoleRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoleType

`func (o *RevokeConsumptionUserRoleRequest) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *RevokeConsumptionUserRoleRequest) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *RevokeConsumptionUserRoleRequest) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.


### GetSubscriptionId

`func (o *RevokeConsumptionUserRoleRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RevokeConsumptionUserRoleRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RevokeConsumptionUserRoleRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetToken

`func (o *RevokeConsumptionUserRoleRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RevokeConsumptionUserRoleRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RevokeConsumptionUserRoleRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


