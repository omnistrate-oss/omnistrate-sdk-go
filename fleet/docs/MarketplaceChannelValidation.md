# MarketplaceChannelValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Findings** | [**[]MarketplaceChannelFinding**](MarketplaceChannelFinding.md) | Everything wrong, not just the first thing wrong | 
**Ok** | **bool** | True when the channel could be enabled right now | 

## Methods

### NewMarketplaceChannelValidation

`func NewMarketplaceChannelValidation(findings []MarketplaceChannelFinding, ok bool, ) *MarketplaceChannelValidation`

NewMarketplaceChannelValidation instantiates a new MarketplaceChannelValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceChannelValidationWithDefaults

`func NewMarketplaceChannelValidationWithDefaults() *MarketplaceChannelValidation`

NewMarketplaceChannelValidationWithDefaults instantiates a new MarketplaceChannelValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFindings

`func (o *MarketplaceChannelValidation) GetFindings() []MarketplaceChannelFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *MarketplaceChannelValidation) GetFindingsOk() (*[]MarketplaceChannelFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *MarketplaceChannelValidation) SetFindings(v []MarketplaceChannelFinding)`

SetFindings sets Findings field to given value.


### GetOk

`func (o *MarketplaceChannelValidation) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *MarketplaceChannelValidation) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *MarketplaceChannelValidation) SetOk(v bool)`

SetOk sets Ok field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


