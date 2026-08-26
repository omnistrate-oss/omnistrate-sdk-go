# MarketplaceEventOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | The buyer organization created for this contract | 
**RootUserId** | **string** | The synthetic root user that owns the buyer organization | 
**SyntheticEmail** | **string** | The derived address the root user was created with. It is deterministic, derived from the channel and the buyer reference, and it is NOT guaranteed to be a deliverable mailbox. Several marketplaces do not return a buyer email at all, so this is an identifier rather than a contact address. Do not send mail to it | 

## Methods

### NewMarketplaceEventOrg

`func NewMarketplaceEventOrg(orgId string, rootUserId string, syntheticEmail string, ) *MarketplaceEventOrg`

NewMarketplaceEventOrg instantiates a new MarketplaceEventOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceEventOrgWithDefaults

`func NewMarketplaceEventOrgWithDefaults() *MarketplaceEventOrg`

NewMarketplaceEventOrgWithDefaults instantiates a new MarketplaceEventOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *MarketplaceEventOrg) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MarketplaceEventOrg) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MarketplaceEventOrg) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRootUserId

`func (o *MarketplaceEventOrg) GetRootUserId() string`

GetRootUserId returns the RootUserId field if non-nil, zero value otherwise.

### GetRootUserIdOk

`func (o *MarketplaceEventOrg) GetRootUserIdOk() (*string, bool)`

GetRootUserIdOk returns a tuple with the RootUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserId

`func (o *MarketplaceEventOrg) SetRootUserId(v string)`

SetRootUserId sets RootUserId field to given value.


### GetSyntheticEmail

`func (o *MarketplaceEventOrg) GetSyntheticEmail() string`

GetSyntheticEmail returns the SyntheticEmail field if non-nil, zero value otherwise.

### GetSyntheticEmailOk

`func (o *MarketplaceEventOrg) GetSyntheticEmailOk() (*string, bool)`

GetSyntheticEmailOk returns a tuple with the SyntheticEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEmail

`func (o *MarketplaceEventOrg) SetSyntheticEmail(v string)`

SetSyntheticEmail sets SyntheticEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


