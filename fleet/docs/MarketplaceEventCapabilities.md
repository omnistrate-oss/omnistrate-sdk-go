# MarketplaceEventCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanCancel** | **bool** | False on every real channel today. Nothing should branch on it being true | 
**CanHoldContract** | **bool** | Whether the channel can keep the contract un-charged until Omnistrate commits it. False means the contract is already live and billing races onboarding | 
**CanReportProgress** | **bool** | Whether onboarding progress can be shown to the buyer on the marketplace | 
**UsageGate** | **string** | How the channel treats usage submitted before it opens the gate. HARD means early usage is dropped upstream with no error, which loses revenue silently. SOFT means early usage is accepted and attributed. INHERIT means the effective gate is whatever partner marketplace sits under an aggregated listing, and is treated as HARD wherever the distinction decides whether to send | 

## Methods

### NewMarketplaceEventCapabilities

`func NewMarketplaceEventCapabilities(canCancel bool, canHoldContract bool, canReportProgress bool, usageGate string, ) *MarketplaceEventCapabilities`

NewMarketplaceEventCapabilities instantiates a new MarketplaceEventCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceEventCapabilitiesWithDefaults

`func NewMarketplaceEventCapabilitiesWithDefaults() *MarketplaceEventCapabilities`

NewMarketplaceEventCapabilitiesWithDefaults instantiates a new MarketplaceEventCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanCancel

`func (o *MarketplaceEventCapabilities) GetCanCancel() bool`

GetCanCancel returns the CanCancel field if non-nil, zero value otherwise.

### GetCanCancelOk

`func (o *MarketplaceEventCapabilities) GetCanCancelOk() (*bool, bool)`

GetCanCancelOk returns a tuple with the CanCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCancel

`func (o *MarketplaceEventCapabilities) SetCanCancel(v bool)`

SetCanCancel sets CanCancel field to given value.


### GetCanHoldContract

`func (o *MarketplaceEventCapabilities) GetCanHoldContract() bool`

GetCanHoldContract returns the CanHoldContract field if non-nil, zero value otherwise.

### GetCanHoldContractOk

`func (o *MarketplaceEventCapabilities) GetCanHoldContractOk() (*bool, bool)`

GetCanHoldContractOk returns a tuple with the CanHoldContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanHoldContract

`func (o *MarketplaceEventCapabilities) SetCanHoldContract(v bool)`

SetCanHoldContract sets CanHoldContract field to given value.


### GetCanReportProgress

`func (o *MarketplaceEventCapabilities) GetCanReportProgress() bool`

GetCanReportProgress returns the CanReportProgress field if non-nil, zero value otherwise.

### GetCanReportProgressOk

`func (o *MarketplaceEventCapabilities) GetCanReportProgressOk() (*bool, bool)`

GetCanReportProgressOk returns a tuple with the CanReportProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReportProgress

`func (o *MarketplaceEventCapabilities) SetCanReportProgress(v bool)`

SetCanReportProgress sets CanReportProgress field to given value.


### GetUsageGate

`func (o *MarketplaceEventCapabilities) GetUsageGate() string`

GetUsageGate returns the UsageGate field if non-nil, zero value otherwise.

### GetUsageGateOk

`func (o *MarketplaceEventCapabilities) GetUsageGateOk() (*string, bool)`

GetUsageGateOk returns a tuple with the UsageGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageGate

`func (o *MarketplaceEventCapabilities) SetUsageGate(v string)`

SetUsageGate sets UsageGate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


