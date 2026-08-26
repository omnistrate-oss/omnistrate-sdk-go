# ConfirmMarketplaceFulfillmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalReference** | Pointer to **string** | Your own identifier for the tenant you provisioned, recorded on the audit event. Optional, and worth sending: it is what lets a support conversation start from your id rather than ours | [optional] 
**Id** | **string** | The Omnistrate contract identifier, as carried by marketplaceContractId on the contract.discovered event and returned by the handoff redeem | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewConfirmMarketplaceFulfillmentRequest

`func NewConfirmMarketplaceFulfillmentRequest(id string, token string, ) *ConfirmMarketplaceFulfillmentRequest`

NewConfirmMarketplaceFulfillmentRequest instantiates a new ConfirmMarketplaceFulfillmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmMarketplaceFulfillmentRequestWithDefaults

`func NewConfirmMarketplaceFulfillmentRequestWithDefaults() *ConfirmMarketplaceFulfillmentRequest`

NewConfirmMarketplaceFulfillmentRequestWithDefaults instantiates a new ConfirmMarketplaceFulfillmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalReference

`func (o *ConfirmMarketplaceFulfillmentRequest) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *ConfirmMarketplaceFulfillmentRequest) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *ConfirmMarketplaceFulfillmentRequest) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *ConfirmMarketplaceFulfillmentRequest) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetId

`func (o *ConfirmMarketplaceFulfillmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfirmMarketplaceFulfillmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfirmMarketplaceFulfillmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *ConfirmMarketplaceFulfillmentRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConfirmMarketplaceFulfillmentRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConfirmMarketplaceFulfillmentRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


