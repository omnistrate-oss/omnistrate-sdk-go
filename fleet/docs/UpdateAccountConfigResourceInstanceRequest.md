# UpdateAccountConfigResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disconnect** | Pointer to **bool** | Disconnect account config instance or not | [optional] 
**Id** | **string** | The instance ID | 
**ServiceId** | **string** | The service ID | 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateAccountConfigResourceInstanceRequest

`func NewUpdateAccountConfigResourceInstanceRequest(id string, serviceId string, token string, ) *UpdateAccountConfigResourceInstanceRequest`

NewUpdateAccountConfigResourceInstanceRequest instantiates a new UpdateAccountConfigResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountConfigResourceInstanceRequestWithDefaults

`func NewUpdateAccountConfigResourceInstanceRequestWithDefaults() *UpdateAccountConfigResourceInstanceRequest`

NewUpdateAccountConfigResourceInstanceRequestWithDefaults instantiates a new UpdateAccountConfigResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisconnect

`func (o *UpdateAccountConfigResourceInstanceRequest) GetDisconnect() bool`

GetDisconnect returns the Disconnect field if non-nil, zero value otherwise.

### GetDisconnectOk

`func (o *UpdateAccountConfigResourceInstanceRequest) GetDisconnectOk() (*bool, bool)`

GetDisconnectOk returns a tuple with the Disconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnect

`func (o *UpdateAccountConfigResourceInstanceRequest) SetDisconnect(v bool)`

SetDisconnect sets Disconnect field to given value.

### HasDisconnect

`func (o *UpdateAccountConfigResourceInstanceRequest) HasDisconnect() bool`

HasDisconnect returns a boolean if a field has been set.

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

### HasSubscriptionId

`func (o *UpdateAccountConfigResourceInstanceRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

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


