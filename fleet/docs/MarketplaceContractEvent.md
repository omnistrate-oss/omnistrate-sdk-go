# MarketplaceContractEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | **time.Time** |  | 
**Detail** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Source** | **string** | Who caused it. ISV events are the inbound calls the partner made | 
**Type** | **string** |  | 

## Methods

### NewMarketplaceContractEvent

`func NewMarketplaceContractEvent(at time.Time, id string, source string, type_ string, ) *MarketplaceContractEvent`

NewMarketplaceContractEvent instantiates a new MarketplaceContractEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceContractEventWithDefaults

`func NewMarketplaceContractEventWithDefaults() *MarketplaceContractEvent`

NewMarketplaceContractEventWithDefaults instantiates a new MarketplaceContractEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketplaceContractEvent) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketplaceContractEvent) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketplaceContractEvent) SetAt(v time.Time)`

SetAt sets At field to given value.


### GetDetail

`func (o *MarketplaceContractEvent) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *MarketplaceContractEvent) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *MarketplaceContractEvent) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *MarketplaceContractEvent) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetId

`func (o *MarketplaceContractEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceContractEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceContractEvent) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *MarketplaceContractEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketplaceContractEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketplaceContractEvent) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *MarketplaceContractEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketplaceContractEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketplaceContractEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


