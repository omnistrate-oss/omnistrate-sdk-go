# RemoveAccountConfigFromServiceModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigId** | **string** | ID of an Account Config | 
**Id** | **string** | ID of a Service Model | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRemoveAccountConfigFromServiceModelRequest

`func NewRemoveAccountConfigFromServiceModelRequest(accountConfigId string, id string, serviceId string, token string, ) *RemoveAccountConfigFromServiceModelRequest`

NewRemoveAccountConfigFromServiceModelRequest instantiates a new RemoveAccountConfigFromServiceModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveAccountConfigFromServiceModelRequestWithDefaults

`func NewRemoveAccountConfigFromServiceModelRequestWithDefaults() *RemoveAccountConfigFromServiceModelRequest`

NewRemoveAccountConfigFromServiceModelRequestWithDefaults instantiates a new RemoveAccountConfigFromServiceModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigId

`func (o *RemoveAccountConfigFromServiceModelRequest) GetAccountConfigId() string`

GetAccountConfigId returns the AccountConfigId field if non-nil, zero value otherwise.

### GetAccountConfigIdOk

`func (o *RemoveAccountConfigFromServiceModelRequest) GetAccountConfigIdOk() (*string, bool)`

GetAccountConfigIdOk returns a tuple with the AccountConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigId

`func (o *RemoveAccountConfigFromServiceModelRequest) SetAccountConfigId(v string)`

SetAccountConfigId sets AccountConfigId field to given value.


### GetId

`func (o *RemoveAccountConfigFromServiceModelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoveAccountConfigFromServiceModelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoveAccountConfigFromServiceModelRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *RemoveAccountConfigFromServiceModelRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RemoveAccountConfigFromServiceModelRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RemoveAccountConfigFromServiceModelRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *RemoveAccountConfigFromServiceModelRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RemoveAccountConfigFromServiceModelRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RemoveAccountConfigFromServiceModelRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


