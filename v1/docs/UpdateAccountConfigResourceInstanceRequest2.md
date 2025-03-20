# UpdateAccountConfigResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | The service ID | 
**SetConnection** | Pointer to **bool** | set account config instance connection | [optional] 
**SubscriptionId** | **string** | The subscription ID | 

## Methods

### NewUpdateAccountConfigResourceInstanceRequest2

`func NewUpdateAccountConfigResourceInstanceRequest2(serviceId string, subscriptionId string, ) *UpdateAccountConfigResourceInstanceRequest2`

NewUpdateAccountConfigResourceInstanceRequest2 instantiates a new UpdateAccountConfigResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountConfigResourceInstanceRequest2WithDefaults

`func NewUpdateAccountConfigResourceInstanceRequest2WithDefaults() *UpdateAccountConfigResourceInstanceRequest2`

NewUpdateAccountConfigResourceInstanceRequest2WithDefaults instantiates a new UpdateAccountConfigResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *UpdateAccountConfigResourceInstanceRequest2) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateAccountConfigResourceInstanceRequest2) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateAccountConfigResourceInstanceRequest2) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSetConnection

`func (o *UpdateAccountConfigResourceInstanceRequest2) GetSetConnection() bool`

GetSetConnection returns the SetConnection field if non-nil, zero value otherwise.

### GetSetConnectionOk

`func (o *UpdateAccountConfigResourceInstanceRequest2) GetSetConnectionOk() (*bool, bool)`

GetSetConnectionOk returns a tuple with the SetConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetConnection

`func (o *UpdateAccountConfigResourceInstanceRequest2) SetSetConnection(v bool)`

SetSetConnection sets SetConnection field to given value.

### HasSetConnection

`func (o *UpdateAccountConfigResourceInstanceRequest2) HasSetConnection() bool`

HasSetConnection returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UpdateAccountConfigResourceInstanceRequest2) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UpdateAccountConfigResourceInstanceRequest2) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UpdateAccountConfigResourceInstanceRequest2) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


