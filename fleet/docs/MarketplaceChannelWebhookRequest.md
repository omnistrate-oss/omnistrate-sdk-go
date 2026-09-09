# MarketplaceChannelWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | 
**ServiceProviderOrgId** | **string** | The ISV whose route this is. Also in the path, and compared against the organization the body claims: a delivery that is genuine for one ISV must not be accepted when replayed at another&#39;s route | 
**Signature** | Pointer to **string** | The X-Suger-Signature-256 header: sha256&#x3D; followed by the lowercase hex HMAC of the raw body | [optional] 

## Methods

### NewMarketplaceChannelWebhookRequest

`func NewMarketplaceChannelWebhookRequest(channel string, serviceProviderOrgId string, ) *MarketplaceChannelWebhookRequest`

NewMarketplaceChannelWebhookRequest instantiates a new MarketplaceChannelWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceChannelWebhookRequestWithDefaults

`func NewMarketplaceChannelWebhookRequestWithDefaults() *MarketplaceChannelWebhookRequest`

NewMarketplaceChannelWebhookRequestWithDefaults instantiates a new MarketplaceChannelWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *MarketplaceChannelWebhookRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MarketplaceChannelWebhookRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MarketplaceChannelWebhookRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetServiceProviderOrgId

`func (o *MarketplaceChannelWebhookRequest) GetServiceProviderOrgId() string`

GetServiceProviderOrgId returns the ServiceProviderOrgId field if non-nil, zero value otherwise.

### GetServiceProviderOrgIdOk

`func (o *MarketplaceChannelWebhookRequest) GetServiceProviderOrgIdOk() (*string, bool)`

GetServiceProviderOrgIdOk returns a tuple with the ServiceProviderOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderOrgId

`func (o *MarketplaceChannelWebhookRequest) SetServiceProviderOrgId(v string)`

SetServiceProviderOrgId sets ServiceProviderOrgId field to given value.


### GetSignature

`func (o *MarketplaceChannelWebhookRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *MarketplaceChannelWebhookRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *MarketplaceChannelWebhookRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *MarketplaceChannelWebhookRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


