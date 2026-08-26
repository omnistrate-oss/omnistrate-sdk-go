# ConnectMarketplaceChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoConfirmIsv** | Pointer to **bool** |  | [optional] 
**BillingCredentials** | Pointer to **map[string]string** | Only for a channel whose billing credential differs from its contract-read credential. Omitted when sharesCredentials is true | [optional] 
**Channel** | **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | 
**Credentials** | Pointer to **map[string]string** | The channel&#39;s own credential fields. Write only: encrypted on arrival and never returned by any read | [optional] 
**DefaultProductTierId** | Pointer to **string** |  | [optional] 
**DefaultServiceEnvironmentId** | Pointer to **string** |  | [optional] 
**DefaultServiceId** | Pointer to **string** |  | [optional] 
**DimensionMap** | Pointer to **map[string]string** |  | [optional] 
**Enabled** | Pointer to **bool** | Enabling is refused unless the config is complete: a callback URL, a mapped plan, and a plan that lists the marketplace billing provider. Refusing here rather than at the first purchase is the difference between a configuration error and a paying buyer nobody can serve | [optional] 
**IsvCallbackUrl** | Pointer to **string** | Where the buyer&#39;s browser lands. https only, no query and no fragment: Omnistrate appends its own parameter, and the code it appends is a bearer value that plaintext would expose in transit | [optional] 
**IsvConfirmTimeoutSeconds** | Pointer to **int64** |  | [optional] 
**PlanMap** | Pointer to [**map[string]MarketplacePlanMapping**](MarketplacePlanMapping.md) |  | [optional] 
**PortBReceiverUrl** | Pointer to **string** | Where signed webhooks are delivered. Same shape rules as the callback, plus the address rules: a private, loopback, link-local or metadata address is refused here rather than discovered at the first purchase | [optional] 
**PortBSigningSecret** | Pointer to **string** | The secret webhook deliveries to portBReceiverUrl are signed with, so your receiver can prove they came from us. Write only: encrypted on arrival and returned by no read, so keep a copy. At least 32 bytes, because deliveries are signed with HMAC-SHA256 and a key shorter than the hash weakens it.  REQUIRED THE FIRST TIME, and omitted afterwards to keep the stored one. Supplying a value IS a rotation, so sending one on every connect replaces the secret your receiver verifies with, every time. Connecting a channel that already exists is how an ISV edits it, which made that the common case rather than the rare one.  The rule is that every channel ends this call with a secret: one was supplied, or one is already stored. A connect with neither is refused, and the refusal says so | [optional] 
**SyntheticEmailDomain** | Pointer to **string** |  | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewConnectMarketplaceChannelRequest

`func NewConnectMarketplaceChannelRequest(channel string, token string, ) *ConnectMarketplaceChannelRequest`

NewConnectMarketplaceChannelRequest instantiates a new ConnectMarketplaceChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectMarketplaceChannelRequestWithDefaults

`func NewConnectMarketplaceChannelRequestWithDefaults() *ConnectMarketplaceChannelRequest`

NewConnectMarketplaceChannelRequestWithDefaults instantiates a new ConnectMarketplaceChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoConfirmIsv

`func (o *ConnectMarketplaceChannelRequest) GetAutoConfirmIsv() bool`

GetAutoConfirmIsv returns the AutoConfirmIsv field if non-nil, zero value otherwise.

### GetAutoConfirmIsvOk

`func (o *ConnectMarketplaceChannelRequest) GetAutoConfirmIsvOk() (*bool, bool)`

GetAutoConfirmIsvOk returns a tuple with the AutoConfirmIsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfirmIsv

`func (o *ConnectMarketplaceChannelRequest) SetAutoConfirmIsv(v bool)`

SetAutoConfirmIsv sets AutoConfirmIsv field to given value.

### HasAutoConfirmIsv

`func (o *ConnectMarketplaceChannelRequest) HasAutoConfirmIsv() bool`

HasAutoConfirmIsv returns a boolean if a field has been set.

### GetBillingCredentials

`func (o *ConnectMarketplaceChannelRequest) GetBillingCredentials() map[string]string`

GetBillingCredentials returns the BillingCredentials field if non-nil, zero value otherwise.

### GetBillingCredentialsOk

`func (o *ConnectMarketplaceChannelRequest) GetBillingCredentialsOk() (*map[string]string, bool)`

GetBillingCredentialsOk returns a tuple with the BillingCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCredentials

`func (o *ConnectMarketplaceChannelRequest) SetBillingCredentials(v map[string]string)`

SetBillingCredentials sets BillingCredentials field to given value.

### HasBillingCredentials

`func (o *ConnectMarketplaceChannelRequest) HasBillingCredentials() bool`

HasBillingCredentials returns a boolean if a field has been set.

### GetChannel

`func (o *ConnectMarketplaceChannelRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ConnectMarketplaceChannelRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ConnectMarketplaceChannelRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCredentials

`func (o *ConnectMarketplaceChannelRequest) GetCredentials() map[string]string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ConnectMarketplaceChannelRequest) GetCredentialsOk() (*map[string]string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ConnectMarketplaceChannelRequest) SetCredentials(v map[string]string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ConnectMarketplaceChannelRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDefaultProductTierId

`func (o *ConnectMarketplaceChannelRequest) GetDefaultProductTierId() string`

GetDefaultProductTierId returns the DefaultProductTierId field if non-nil, zero value otherwise.

### GetDefaultProductTierIdOk

`func (o *ConnectMarketplaceChannelRequest) GetDefaultProductTierIdOk() (*string, bool)`

GetDefaultProductTierIdOk returns a tuple with the DefaultProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProductTierId

`func (o *ConnectMarketplaceChannelRequest) SetDefaultProductTierId(v string)`

SetDefaultProductTierId sets DefaultProductTierId field to given value.

### HasDefaultProductTierId

`func (o *ConnectMarketplaceChannelRequest) HasDefaultProductTierId() bool`

HasDefaultProductTierId returns a boolean if a field has been set.

### GetDefaultServiceEnvironmentId

`func (o *ConnectMarketplaceChannelRequest) GetDefaultServiceEnvironmentId() string`

GetDefaultServiceEnvironmentId returns the DefaultServiceEnvironmentId field if non-nil, zero value otherwise.

### GetDefaultServiceEnvironmentIdOk

`func (o *ConnectMarketplaceChannelRequest) GetDefaultServiceEnvironmentIdOk() (*string, bool)`

GetDefaultServiceEnvironmentIdOk returns a tuple with the DefaultServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceEnvironmentId

`func (o *ConnectMarketplaceChannelRequest) SetDefaultServiceEnvironmentId(v string)`

SetDefaultServiceEnvironmentId sets DefaultServiceEnvironmentId field to given value.

### HasDefaultServiceEnvironmentId

`func (o *ConnectMarketplaceChannelRequest) HasDefaultServiceEnvironmentId() bool`

HasDefaultServiceEnvironmentId returns a boolean if a field has been set.

### GetDefaultServiceId

`func (o *ConnectMarketplaceChannelRequest) GetDefaultServiceId() string`

GetDefaultServiceId returns the DefaultServiceId field if non-nil, zero value otherwise.

### GetDefaultServiceIdOk

`func (o *ConnectMarketplaceChannelRequest) GetDefaultServiceIdOk() (*string, bool)`

GetDefaultServiceIdOk returns a tuple with the DefaultServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceId

`func (o *ConnectMarketplaceChannelRequest) SetDefaultServiceId(v string)`

SetDefaultServiceId sets DefaultServiceId field to given value.

### HasDefaultServiceId

`func (o *ConnectMarketplaceChannelRequest) HasDefaultServiceId() bool`

HasDefaultServiceId returns a boolean if a field has been set.

### GetDimensionMap

`func (o *ConnectMarketplaceChannelRequest) GetDimensionMap() map[string]string`

GetDimensionMap returns the DimensionMap field if non-nil, zero value otherwise.

### GetDimensionMapOk

`func (o *ConnectMarketplaceChannelRequest) GetDimensionMapOk() (*map[string]string, bool)`

GetDimensionMapOk returns a tuple with the DimensionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionMap

`func (o *ConnectMarketplaceChannelRequest) SetDimensionMap(v map[string]string)`

SetDimensionMap sets DimensionMap field to given value.

### HasDimensionMap

`func (o *ConnectMarketplaceChannelRequest) HasDimensionMap() bool`

HasDimensionMap returns a boolean if a field has been set.

### GetEnabled

`func (o *ConnectMarketplaceChannelRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConnectMarketplaceChannelRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConnectMarketplaceChannelRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConnectMarketplaceChannelRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsvCallbackUrl

`func (o *ConnectMarketplaceChannelRequest) GetIsvCallbackUrl() string`

GetIsvCallbackUrl returns the IsvCallbackUrl field if non-nil, zero value otherwise.

### GetIsvCallbackUrlOk

`func (o *ConnectMarketplaceChannelRequest) GetIsvCallbackUrlOk() (*string, bool)`

GetIsvCallbackUrlOk returns a tuple with the IsvCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvCallbackUrl

`func (o *ConnectMarketplaceChannelRequest) SetIsvCallbackUrl(v string)`

SetIsvCallbackUrl sets IsvCallbackUrl field to given value.

### HasIsvCallbackUrl

`func (o *ConnectMarketplaceChannelRequest) HasIsvCallbackUrl() bool`

HasIsvCallbackUrl returns a boolean if a field has been set.

### GetIsvConfirmTimeoutSeconds

`func (o *ConnectMarketplaceChannelRequest) GetIsvConfirmTimeoutSeconds() int64`

GetIsvConfirmTimeoutSeconds returns the IsvConfirmTimeoutSeconds field if non-nil, zero value otherwise.

### GetIsvConfirmTimeoutSecondsOk

`func (o *ConnectMarketplaceChannelRequest) GetIsvConfirmTimeoutSecondsOk() (*int64, bool)`

GetIsvConfirmTimeoutSecondsOk returns a tuple with the IsvConfirmTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvConfirmTimeoutSeconds

`func (o *ConnectMarketplaceChannelRequest) SetIsvConfirmTimeoutSeconds(v int64)`

SetIsvConfirmTimeoutSeconds sets IsvConfirmTimeoutSeconds field to given value.

### HasIsvConfirmTimeoutSeconds

`func (o *ConnectMarketplaceChannelRequest) HasIsvConfirmTimeoutSeconds() bool`

HasIsvConfirmTimeoutSeconds returns a boolean if a field has been set.

### GetPlanMap

`func (o *ConnectMarketplaceChannelRequest) GetPlanMap() map[string]MarketplacePlanMapping`

GetPlanMap returns the PlanMap field if non-nil, zero value otherwise.

### GetPlanMapOk

`func (o *ConnectMarketplaceChannelRequest) GetPlanMapOk() (*map[string]MarketplacePlanMapping, bool)`

GetPlanMapOk returns a tuple with the PlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMap

`func (o *ConnectMarketplaceChannelRequest) SetPlanMap(v map[string]MarketplacePlanMapping)`

SetPlanMap sets PlanMap field to given value.

### HasPlanMap

`func (o *ConnectMarketplaceChannelRequest) HasPlanMap() bool`

HasPlanMap returns a boolean if a field has been set.

### GetPortBReceiverUrl

`func (o *ConnectMarketplaceChannelRequest) GetPortBReceiverUrl() string`

GetPortBReceiverUrl returns the PortBReceiverUrl field if non-nil, zero value otherwise.

### GetPortBReceiverUrlOk

`func (o *ConnectMarketplaceChannelRequest) GetPortBReceiverUrlOk() (*string, bool)`

GetPortBReceiverUrlOk returns a tuple with the PortBReceiverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortBReceiverUrl

`func (o *ConnectMarketplaceChannelRequest) SetPortBReceiverUrl(v string)`

SetPortBReceiverUrl sets PortBReceiverUrl field to given value.

### HasPortBReceiverUrl

`func (o *ConnectMarketplaceChannelRequest) HasPortBReceiverUrl() bool`

HasPortBReceiverUrl returns a boolean if a field has been set.

### GetPortBSigningSecret

`func (o *ConnectMarketplaceChannelRequest) GetPortBSigningSecret() string`

GetPortBSigningSecret returns the PortBSigningSecret field if non-nil, zero value otherwise.

### GetPortBSigningSecretOk

`func (o *ConnectMarketplaceChannelRequest) GetPortBSigningSecretOk() (*string, bool)`

GetPortBSigningSecretOk returns a tuple with the PortBSigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortBSigningSecret

`func (o *ConnectMarketplaceChannelRequest) SetPortBSigningSecret(v string)`

SetPortBSigningSecret sets PortBSigningSecret field to given value.

### HasPortBSigningSecret

`func (o *ConnectMarketplaceChannelRequest) HasPortBSigningSecret() bool`

HasPortBSigningSecret returns a boolean if a field has been set.

### GetSyntheticEmailDomain

`func (o *ConnectMarketplaceChannelRequest) GetSyntheticEmailDomain() string`

GetSyntheticEmailDomain returns the SyntheticEmailDomain field if non-nil, zero value otherwise.

### GetSyntheticEmailDomainOk

`func (o *ConnectMarketplaceChannelRequest) GetSyntheticEmailDomainOk() (*string, bool)`

GetSyntheticEmailDomainOk returns a tuple with the SyntheticEmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEmailDomain

`func (o *ConnectMarketplaceChannelRequest) SetSyntheticEmailDomain(v string)`

SetSyntheticEmailDomain sets SyntheticEmailDomain field to given value.

### HasSyntheticEmailDomain

`func (o *ConnectMarketplaceChannelRequest) HasSyntheticEmailDomain() bool`

HasSyntheticEmailDomain returns a boolean if a field has been set.

### GetToken

`func (o *ConnectMarketplaceChannelRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConnectMarketplaceChannelRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConnectMarketplaceChannelRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


