# MarketplaceChannelCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanCancel** | **bool** | Whether Omnistrate can cancel the contract, as opposed to only observing a cancellation | 
**CanHoldContract** | **bool** | Whether the contract can be held before the ISV confirms, so the buyer is not charged while they are being onboarded | 
**CanReportProgress** | **bool** | Whether a progress message can be shown to the buyer on the marketplace side | 
**UsageGate** | **string** | Whether usage may be reported before the channel has confirmed the contract is billable. INHERIT is treated as HARD wherever the difference decides whether to send | 

## Methods

### NewMarketplaceChannelCapabilities

`func NewMarketplaceChannelCapabilities(canCancel bool, canHoldContract bool, canReportProgress bool, usageGate string, ) *MarketplaceChannelCapabilities`

NewMarketplaceChannelCapabilities instantiates a new MarketplaceChannelCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceChannelCapabilitiesWithDefaults

`func NewMarketplaceChannelCapabilitiesWithDefaults() *MarketplaceChannelCapabilities`

NewMarketplaceChannelCapabilitiesWithDefaults instantiates a new MarketplaceChannelCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanCancel

`func (o *MarketplaceChannelCapabilities) GetCanCancel() bool`

GetCanCancel returns the CanCancel field if non-nil, zero value otherwise.

### GetCanCancelOk

`func (o *MarketplaceChannelCapabilities) GetCanCancelOk() (*bool, bool)`

GetCanCancelOk returns a tuple with the CanCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCancel

`func (o *MarketplaceChannelCapabilities) SetCanCancel(v bool)`

SetCanCancel sets CanCancel field to given value.


### GetCanHoldContract

`func (o *MarketplaceChannelCapabilities) GetCanHoldContract() bool`

GetCanHoldContract returns the CanHoldContract field if non-nil, zero value otherwise.

### GetCanHoldContractOk

`func (o *MarketplaceChannelCapabilities) GetCanHoldContractOk() (*bool, bool)`

GetCanHoldContractOk returns a tuple with the CanHoldContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanHoldContract

`func (o *MarketplaceChannelCapabilities) SetCanHoldContract(v bool)`

SetCanHoldContract sets CanHoldContract field to given value.


### GetCanReportProgress

`func (o *MarketplaceChannelCapabilities) GetCanReportProgress() bool`

GetCanReportProgress returns the CanReportProgress field if non-nil, zero value otherwise.

### GetCanReportProgressOk

`func (o *MarketplaceChannelCapabilities) GetCanReportProgressOk() (*bool, bool)`

GetCanReportProgressOk returns a tuple with the CanReportProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReportProgress

`func (o *MarketplaceChannelCapabilities) SetCanReportProgress(v bool)`

SetCanReportProgress sets CanReportProgress field to given value.


### GetUsageGate

`func (o *MarketplaceChannelCapabilities) GetUsageGate() string`

GetUsageGate returns the UsageGate field if non-nil, zero value otherwise.

### GetUsageGateOk

`func (o *MarketplaceChannelCapabilities) GetUsageGateOk() (*string, bool)`

GetUsageGateOk returns a tuple with the UsageGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageGate

`func (o *MarketplaceChannelCapabilities) SetUsageGate(v string)`

SetUsageGate sets UsageGate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


