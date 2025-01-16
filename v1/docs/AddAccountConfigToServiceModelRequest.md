# AddAccountConfigToServiceModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigId** | **string** | ID of an Account Config | 
**Id** | **string** | ID of a Service Model | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewAddAccountConfigToServiceModelRequest

`func NewAddAccountConfigToServiceModelRequest(accountConfigId string, id string, serviceId string, token string, ) *AddAccountConfigToServiceModelRequest`

NewAddAccountConfigToServiceModelRequest instantiates a new AddAccountConfigToServiceModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccountConfigToServiceModelRequestWithDefaults

`func NewAddAccountConfigToServiceModelRequestWithDefaults() *AddAccountConfigToServiceModelRequest`

NewAddAccountConfigToServiceModelRequestWithDefaults instantiates a new AddAccountConfigToServiceModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigId

`func (o *AddAccountConfigToServiceModelRequest) GetAccountConfigId() string`

GetAccountConfigId returns the AccountConfigId field if non-nil, zero value otherwise.

### GetAccountConfigIdOk

`func (o *AddAccountConfigToServiceModelRequest) GetAccountConfigIdOk() (*string, bool)`

GetAccountConfigIdOk returns a tuple with the AccountConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigId

`func (o *AddAccountConfigToServiceModelRequest) SetAccountConfigId(v string)`

SetAccountConfigId sets AccountConfigId field to given value.


### GetId

`func (o *AddAccountConfigToServiceModelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddAccountConfigToServiceModelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddAccountConfigToServiceModelRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *AddAccountConfigToServiceModelRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AddAccountConfigToServiceModelRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AddAccountConfigToServiceModelRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *AddAccountConfigToServiceModelRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddAccountConfigToServiceModelRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddAccountConfigToServiceModelRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


