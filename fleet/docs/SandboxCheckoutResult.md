# SandboxCheckoutResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheControl** | **string** | Always no-store, so a back-button replay re-runs the arrival rather than replaying a cached one. The arrival itself is idempotent on the contract&#39;s reference | 
**Location** | **string** | The landing route for this organization on the simulated channel, carrying the same query parameters Suger appends to a vendor&#39;s Product Fulfillment URL. Following it is what a marketplace redirect does to a real buyer | 

## Methods

### NewSandboxCheckoutResult

`func NewSandboxCheckoutResult(cacheControl string, location string, ) *SandboxCheckoutResult`

NewSandboxCheckoutResult instantiates a new SandboxCheckoutResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCheckoutResultWithDefaults

`func NewSandboxCheckoutResultWithDefaults() *SandboxCheckoutResult`

NewSandboxCheckoutResultWithDefaults instantiates a new SandboxCheckoutResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheControl

`func (o *SandboxCheckoutResult) GetCacheControl() string`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *SandboxCheckoutResult) GetCacheControlOk() (*string, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *SandboxCheckoutResult) SetCacheControl(v string)`

SetCacheControl sets CacheControl field to given value.


### GetLocation

`func (o *SandboxCheckoutResult) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SandboxCheckoutResult) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SandboxCheckoutResult) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


