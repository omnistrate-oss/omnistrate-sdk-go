# DeregisterActionHookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a resource | 
**Scope** | **string** | The scope of the hook | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**Type** | **string** | The type of hook to execute | 

## Methods

### NewDeregisterActionHookRequest

`func NewDeregisterActionHookRequest(id string, scope string, serviceId string, token string, type_ string, ) *DeregisterActionHookRequest`

NewDeregisterActionHookRequest instantiates a new DeregisterActionHookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeregisterActionHookRequestWithDefaults

`func NewDeregisterActionHookRequestWithDefaults() *DeregisterActionHookRequest`

NewDeregisterActionHookRequestWithDefaults instantiates a new DeregisterActionHookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeregisterActionHookRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeregisterActionHookRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeregisterActionHookRequest) SetId(v string)`

SetId sets Id field to given value.


### GetScope

`func (o *DeregisterActionHookRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DeregisterActionHookRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DeregisterActionHookRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetServiceId

`func (o *DeregisterActionHookRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeregisterActionHookRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeregisterActionHookRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *DeregisterActionHookRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeregisterActionHookRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeregisterActionHookRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *DeregisterActionHookRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeregisterActionHookRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeregisterActionHookRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


