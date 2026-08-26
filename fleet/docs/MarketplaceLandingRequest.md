# MarketplaceLandingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | 
**OfferType** | Pointer to **string** | Suger&#39;s offer type for the purchase. Recorded for audit only | [optional] 
**Partner** | Pointer to **string** | Which cloud the purchase originated on, as Suger reports it. Recorded for audit; the authoritative value comes from the readback | [optional] 
**ServiceProviderOrgId** | **string** | Which ISV organization&#39;s listing was purchased. A SELECTOR for whose stored channel credentials perform the server-side readback, not a credential: it grants nothing, and substituting another organization&#39;s id resolves the token against an account where it does not exist | 
**SugerEntitlementId** | Pointer to **string** | Suger&#39;s entitlement identifier, appended by Suger&#39;s signup redirect. A pointer to be read back, never believed | [optional] 

## Methods

### NewMarketplaceLandingRequest

`func NewMarketplaceLandingRequest(channel string, serviceProviderOrgId string, ) *MarketplaceLandingRequest`

NewMarketplaceLandingRequest instantiates a new MarketplaceLandingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceLandingRequestWithDefaults

`func NewMarketplaceLandingRequestWithDefaults() *MarketplaceLandingRequest`

NewMarketplaceLandingRequestWithDefaults instantiates a new MarketplaceLandingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *MarketplaceLandingRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MarketplaceLandingRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MarketplaceLandingRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetOfferType

`func (o *MarketplaceLandingRequest) GetOfferType() string`

GetOfferType returns the OfferType field if non-nil, zero value otherwise.

### GetOfferTypeOk

`func (o *MarketplaceLandingRequest) GetOfferTypeOk() (*string, bool)`

GetOfferTypeOk returns a tuple with the OfferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferType

`func (o *MarketplaceLandingRequest) SetOfferType(v string)`

SetOfferType sets OfferType field to given value.

### HasOfferType

`func (o *MarketplaceLandingRequest) HasOfferType() bool`

HasOfferType returns a boolean if a field has been set.

### GetPartner

`func (o *MarketplaceLandingRequest) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *MarketplaceLandingRequest) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *MarketplaceLandingRequest) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *MarketplaceLandingRequest) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetServiceProviderOrgId

`func (o *MarketplaceLandingRequest) GetServiceProviderOrgId() string`

GetServiceProviderOrgId returns the ServiceProviderOrgId field if non-nil, zero value otherwise.

### GetServiceProviderOrgIdOk

`func (o *MarketplaceLandingRequest) GetServiceProviderOrgIdOk() (*string, bool)`

GetServiceProviderOrgIdOk returns a tuple with the ServiceProviderOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderOrgId

`func (o *MarketplaceLandingRequest) SetServiceProviderOrgId(v string)`

SetServiceProviderOrgId sets ServiceProviderOrgId field to given value.


### GetSugerEntitlementId

`func (o *MarketplaceLandingRequest) GetSugerEntitlementId() string`

GetSugerEntitlementId returns the SugerEntitlementId field if non-nil, zero value otherwise.

### GetSugerEntitlementIdOk

`func (o *MarketplaceLandingRequest) GetSugerEntitlementIdOk() (*string, bool)`

GetSugerEntitlementIdOk returns a tuple with the SugerEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSugerEntitlementId

`func (o *MarketplaceLandingRequest) SetSugerEntitlementId(v string)`

SetSugerEntitlementId sets SugerEntitlementId field to given value.

### HasSugerEntitlementId

`func (o *MarketplaceLandingRequest) HasSugerEntitlementId() bool`

HasSugerEntitlementId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


