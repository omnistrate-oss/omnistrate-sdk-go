# ListAllSubscriptionUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListAllSubscriptionUsersRequest

`func NewListAllSubscriptionUsersRequest(token string, ) *ListAllSubscriptionUsersRequest`

NewListAllSubscriptionUsersRequest instantiates a new ListAllSubscriptionUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllSubscriptionUsersRequestWithDefaults

`func NewListAllSubscriptionUsersRequestWithDefaults() *ListAllSubscriptionUsersRequest`

NewListAllSubscriptionUsersRequestWithDefaults instantiates a new ListAllSubscriptionUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *ListAllSubscriptionUsersRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *ListAllSubscriptionUsersRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *ListAllSubscriptionUsersRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *ListAllSubscriptionUsersRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetToken

`func (o *ListAllSubscriptionUsersRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAllSubscriptionUsersRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAllSubscriptionUsersRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


