# DenyMarketplaceFulfillmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Omnistrate contract identifier | 
**Reason** | Pointer to **string** | Why the buyer was refused. Recorded on the audit event and shown to an operator, because a refused marketplace purchase is a conversation somebody will have to have | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDenyMarketplaceFulfillmentRequest

`func NewDenyMarketplaceFulfillmentRequest(id string, token string, ) *DenyMarketplaceFulfillmentRequest`

NewDenyMarketplaceFulfillmentRequest instantiates a new DenyMarketplaceFulfillmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDenyMarketplaceFulfillmentRequestWithDefaults

`func NewDenyMarketplaceFulfillmentRequestWithDefaults() *DenyMarketplaceFulfillmentRequest`

NewDenyMarketplaceFulfillmentRequestWithDefaults instantiates a new DenyMarketplaceFulfillmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DenyMarketplaceFulfillmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DenyMarketplaceFulfillmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DenyMarketplaceFulfillmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *DenyMarketplaceFulfillmentRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DenyMarketplaceFulfillmentRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DenyMarketplaceFulfillmentRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DenyMarketplaceFulfillmentRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetToken

`func (o *DenyMarketplaceFulfillmentRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DenyMarketplaceFulfillmentRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DenyMarketplaceFulfillmentRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


