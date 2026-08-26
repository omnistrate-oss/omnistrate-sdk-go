# MarketplaceLandingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheControl** | **string** | Always no-store. The Location carries a single-use credential, so a cached or back-button replay of this response must not hand the same code to anything again | 
**Location** | **string** | The ISV&#39;s configured callback URL with exactly one appended code parameter | 

## Methods

### NewMarketplaceLandingResult

`func NewMarketplaceLandingResult(cacheControl string, location string, ) *MarketplaceLandingResult`

NewMarketplaceLandingResult instantiates a new MarketplaceLandingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceLandingResultWithDefaults

`func NewMarketplaceLandingResultWithDefaults() *MarketplaceLandingResult`

NewMarketplaceLandingResultWithDefaults instantiates a new MarketplaceLandingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheControl

`func (o *MarketplaceLandingResult) GetCacheControl() string`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *MarketplaceLandingResult) GetCacheControlOk() (*string, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *MarketplaceLandingResult) SetCacheControl(v string)`

SetCacheControl sets CacheControl field to given value.


### GetLocation

`func (o *MarketplaceLandingResult) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MarketplaceLandingResult) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MarketplaceLandingResult) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


