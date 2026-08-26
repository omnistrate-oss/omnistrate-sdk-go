# MarketplaceChannelFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Which attribute the finding is about, so a form can put the message next to the input rather than at the top of the page | [optional] 
**Message** | **string** |  | 
**Severity** | **string** | BLOCKING prevents the channel being enabled. WARNING does not, and is worth saying anyway | 

## Methods

### NewMarketplaceChannelFinding

`func NewMarketplaceChannelFinding(message string, severity string, ) *MarketplaceChannelFinding`

NewMarketplaceChannelFinding instantiates a new MarketplaceChannelFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceChannelFindingWithDefaults

`func NewMarketplaceChannelFindingWithDefaults() *MarketplaceChannelFinding`

NewMarketplaceChannelFindingWithDefaults instantiates a new MarketplaceChannelFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *MarketplaceChannelFinding) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MarketplaceChannelFinding) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MarketplaceChannelFinding) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MarketplaceChannelFinding) HasField() bool`

HasField returns a boolean if a field has been set.

### GetMessage

`func (o *MarketplaceChannelFinding) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MarketplaceChannelFinding) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MarketplaceChannelFinding) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSeverity

`func (o *MarketplaceChannelFinding) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MarketplaceChannelFinding) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MarketplaceChannelFinding) SetSeverity(v string)`

SetSeverity sets Severity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


