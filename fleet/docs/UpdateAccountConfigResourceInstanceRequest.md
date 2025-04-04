# UpdateAccountConfigResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The instance ID | 
**ServiceId** | **string** | The service ID | 
**SetConnection** | Pointer to **bool** | set account config instance connection | [optional] 
**SubscriptionId** | **string** | The subscription ID | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateAccountConfigResourceInstanceRequest

`func NewUpdateAccountConfigResourceInstanceRequest(id string, serviceId string, subscriptionId string, token string, ) *UpdateAccountConfigResourceInstanceRequest`

NewUpdateAccountConfigResourceInstanceRequest instantiates a new UpdateAccountConfigResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountConfigResourceInstanceRequestWithDefaults

`func NewUpdateAccountConfigResourceInstanceRequestWithDefaults() *UpdateAccountConfigResourceInstanceRequest`

NewUpdateAccountConfigResourceInstanceRequestWithDefaults instantiates a new UpdateAccountConfigResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAccountConfigResourceInstanceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAccountConfigResourceInstanceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAccountConfigResourceInstanceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *UpdateAccountConfigResourceInstanceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateAccountConfigResourceInstanceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateAccountConfigResourceInstanceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSetConnection

`func (o *UpdateAccountConfigResourceInstanceRequest) GetSetConnection() bool`

GetSetConnection returns the SetConnection field if non-nil, zero value otherwise.

### GetSetConnectionOk

`func (o *UpdateAccountConfigResourceInstanceRequest) GetSetConnectionOk() (*bool, bool)`

GetSetConnectionOk returns a tuple with the SetConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetConnection

`func (o *UpdateAccountConfigResourceInstanceRequest) SetSetConnection(v bool)`

SetSetConnection sets SetConnection field to given value.

### HasSetConnection

`func (o *UpdateAccountConfigResourceInstanceRequest) HasSetConnection() bool`

HasSetConnection returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UpdateAccountConfigResourceInstanceRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UpdateAccountConfigResourceInstanceRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UpdateAccountConfigResourceInstanceRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetToken

`func (o *UpdateAccountConfigResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAccountConfigResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAccountConfigResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


