# MarketplaceBillingBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overridable** | **bool** | False everywhere today. Exists so a channel that genuinely permits a choice can say so, rather than every consumer assuming none ever will | 
**Provider** | **string** | The billing provider that collects. Every marketplace channel reports MARKETPLACE today | 
**Settlement** | **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | 
**SharesCredentials** | **bool** | True when the credential that reads contracts is also the credential that reports usage | 

## Methods

### NewMarketplaceBillingBinding

`func NewMarketplaceBillingBinding(overridable bool, provider string, settlement string, sharesCredentials bool, ) *MarketplaceBillingBinding`

NewMarketplaceBillingBinding instantiates a new MarketplaceBillingBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceBillingBindingWithDefaults

`func NewMarketplaceBillingBindingWithDefaults() *MarketplaceBillingBinding`

NewMarketplaceBillingBindingWithDefaults instantiates a new MarketplaceBillingBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverridable

`func (o *MarketplaceBillingBinding) GetOverridable() bool`

GetOverridable returns the Overridable field if non-nil, zero value otherwise.

### GetOverridableOk

`func (o *MarketplaceBillingBinding) GetOverridableOk() (*bool, bool)`

GetOverridableOk returns a tuple with the Overridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridable

`func (o *MarketplaceBillingBinding) SetOverridable(v bool)`

SetOverridable sets Overridable field to given value.


### GetProvider

`func (o *MarketplaceBillingBinding) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MarketplaceBillingBinding) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MarketplaceBillingBinding) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSettlement

`func (o *MarketplaceBillingBinding) GetSettlement() string`

GetSettlement returns the Settlement field if non-nil, zero value otherwise.

### GetSettlementOk

`func (o *MarketplaceBillingBinding) GetSettlementOk() (*string, bool)`

GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlement

`func (o *MarketplaceBillingBinding) SetSettlement(v string)`

SetSettlement sets Settlement field to given value.


### GetSharesCredentials

`func (o *MarketplaceBillingBinding) GetSharesCredentials() bool`

GetSharesCredentials returns the SharesCredentials field if non-nil, zero value otherwise.

### GetSharesCredentialsOk

`func (o *MarketplaceBillingBinding) GetSharesCredentialsOk() (*bool, bool)`

GetSharesCredentialsOk returns a tuple with the SharesCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesCredentials

`func (o *MarketplaceBillingBinding) SetSharesCredentials(v bool)`

SetSharesCredentials sets SharesCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


