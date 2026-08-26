# DenyMarketplaceFulfillmentRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Why the buyer was refused. Recorded on the audit event and shown to an operator, because a refused marketplace purchase is a conversation somebody will have to have | [optional] 

## Methods

### NewDenyMarketplaceFulfillmentRequest2

`func NewDenyMarketplaceFulfillmentRequest2() *DenyMarketplaceFulfillmentRequest2`

NewDenyMarketplaceFulfillmentRequest2 instantiates a new DenyMarketplaceFulfillmentRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDenyMarketplaceFulfillmentRequest2WithDefaults

`func NewDenyMarketplaceFulfillmentRequest2WithDefaults() *DenyMarketplaceFulfillmentRequest2`

NewDenyMarketplaceFulfillmentRequest2WithDefaults instantiates a new DenyMarketplaceFulfillmentRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *DenyMarketplaceFulfillmentRequest2) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DenyMarketplaceFulfillmentRequest2) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DenyMarketplaceFulfillmentRequest2) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DenyMarketplaceFulfillmentRequest2) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


