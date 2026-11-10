# ConfirmMarketplaceFulfillmentRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalReference** | Pointer to **string** | Your own identifier for the tenant you provisioned, recorded on the audit event. Optional, and worth sending: it is what lets a support conversation start from your id rather than ours | [optional] 
**OrgId** | Pointer to **string** | The existing customer organization in which to create a dedicated subscription. Optional; omitting it creates the subscription in the contract&#39;s buyer organization. Existing subscriptions and their instances are not automatically associated | [optional] 

## Methods

### NewConfirmMarketplaceFulfillmentRequest2

`func NewConfirmMarketplaceFulfillmentRequest2() *ConfirmMarketplaceFulfillmentRequest2`

NewConfirmMarketplaceFulfillmentRequest2 instantiates a new ConfirmMarketplaceFulfillmentRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmMarketplaceFulfillmentRequest2WithDefaults

`func NewConfirmMarketplaceFulfillmentRequest2WithDefaults() *ConfirmMarketplaceFulfillmentRequest2`

NewConfirmMarketplaceFulfillmentRequest2WithDefaults instantiates a new ConfirmMarketplaceFulfillmentRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalReference

`func (o *ConfirmMarketplaceFulfillmentRequest2) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *ConfirmMarketplaceFulfillmentRequest2) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *ConfirmMarketplaceFulfillmentRequest2) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *ConfirmMarketplaceFulfillmentRequest2) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetOrgId

`func (o *ConfirmMarketplaceFulfillmentRequest2) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ConfirmMarketplaceFulfillmentRequest2) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ConfirmMarketplaceFulfillmentRequest2) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ConfirmMarketplaceFulfillmentRequest2) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


