# UpdateMarketplaceChannelRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoConfirmIsv** | Pointer to **bool** |  | [optional] 
**BillingCredentials** | Pointer to **map[string]string** |  | [optional] 
**Credentials** | Pointer to **map[string]string** | Replaces the stored credential when present. Omit it to leave the credential untouched, which is what makes it safe to change one unrelated field | [optional] 
**DefaultProductTierId** | Pointer to **string** |  | [optional] 
**DefaultServiceEnvironmentId** | Pointer to **string** |  | [optional] 
**DefaultServiceId** | Pointer to **string** |  | [optional] 
**DimensionMap** | Pointer to **map[string]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**IsvCallbackUrl** | Pointer to **string** |  | [optional] 
**IsvConfirmTimeoutSeconds** | Pointer to **int64** |  | [optional] 
**PlanMap** | Pointer to [**map[string]MarketplacePlanMapping**](MarketplacePlanMapping.md) |  | [optional] 
**PortBReceiverUrl** | Pointer to **string** |  | [optional] 
**PortBSigningSecret** | Pointer to **string** | Replaces the signing secret. Omit it to leave the current one alone, which is what makes it safe to change an unrelated field: supplying one here IS the rotation, and the outgoing secret stays valid for an overlap window so a receiver that has not installed the new one is not cut off the instant you rotate. At least 32 bytes. THERE IS NO WAY TO REMOVE ONE. Omitting it keeps the stored secret and an empty string is refused, because a channel without a signing secret delivers events no conforming receiver accepts, and the ability to reach that state is not worth having | [optional] 
**SyntheticEmailDomain** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateMarketplaceChannelRequest2

`func NewUpdateMarketplaceChannelRequest2() *UpdateMarketplaceChannelRequest2`

NewUpdateMarketplaceChannelRequest2 instantiates a new UpdateMarketplaceChannelRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMarketplaceChannelRequest2WithDefaults

`func NewUpdateMarketplaceChannelRequest2WithDefaults() *UpdateMarketplaceChannelRequest2`

NewUpdateMarketplaceChannelRequest2WithDefaults instantiates a new UpdateMarketplaceChannelRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoConfirmIsv

`func (o *UpdateMarketplaceChannelRequest2) GetAutoConfirmIsv() bool`

GetAutoConfirmIsv returns the AutoConfirmIsv field if non-nil, zero value otherwise.

### GetAutoConfirmIsvOk

`func (o *UpdateMarketplaceChannelRequest2) GetAutoConfirmIsvOk() (*bool, bool)`

GetAutoConfirmIsvOk returns a tuple with the AutoConfirmIsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfirmIsv

`func (o *UpdateMarketplaceChannelRequest2) SetAutoConfirmIsv(v bool)`

SetAutoConfirmIsv sets AutoConfirmIsv field to given value.

### HasAutoConfirmIsv

`func (o *UpdateMarketplaceChannelRequest2) HasAutoConfirmIsv() bool`

HasAutoConfirmIsv returns a boolean if a field has been set.

### GetBillingCredentials

`func (o *UpdateMarketplaceChannelRequest2) GetBillingCredentials() map[string]string`

GetBillingCredentials returns the BillingCredentials field if non-nil, zero value otherwise.

### GetBillingCredentialsOk

`func (o *UpdateMarketplaceChannelRequest2) GetBillingCredentialsOk() (*map[string]string, bool)`

GetBillingCredentialsOk returns a tuple with the BillingCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCredentials

`func (o *UpdateMarketplaceChannelRequest2) SetBillingCredentials(v map[string]string)`

SetBillingCredentials sets BillingCredentials field to given value.

### HasBillingCredentials

`func (o *UpdateMarketplaceChannelRequest2) HasBillingCredentials() bool`

HasBillingCredentials returns a boolean if a field has been set.

### GetCredentials

`func (o *UpdateMarketplaceChannelRequest2) GetCredentials() map[string]string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UpdateMarketplaceChannelRequest2) GetCredentialsOk() (*map[string]string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UpdateMarketplaceChannelRequest2) SetCredentials(v map[string]string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UpdateMarketplaceChannelRequest2) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDefaultProductTierId

`func (o *UpdateMarketplaceChannelRequest2) GetDefaultProductTierId() string`

GetDefaultProductTierId returns the DefaultProductTierId field if non-nil, zero value otherwise.

### GetDefaultProductTierIdOk

`func (o *UpdateMarketplaceChannelRequest2) GetDefaultProductTierIdOk() (*string, bool)`

GetDefaultProductTierIdOk returns a tuple with the DefaultProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProductTierId

`func (o *UpdateMarketplaceChannelRequest2) SetDefaultProductTierId(v string)`

SetDefaultProductTierId sets DefaultProductTierId field to given value.

### HasDefaultProductTierId

`func (o *UpdateMarketplaceChannelRequest2) HasDefaultProductTierId() bool`

HasDefaultProductTierId returns a boolean if a field has been set.

### GetDefaultServiceEnvironmentId

`func (o *UpdateMarketplaceChannelRequest2) GetDefaultServiceEnvironmentId() string`

GetDefaultServiceEnvironmentId returns the DefaultServiceEnvironmentId field if non-nil, zero value otherwise.

### GetDefaultServiceEnvironmentIdOk

`func (o *UpdateMarketplaceChannelRequest2) GetDefaultServiceEnvironmentIdOk() (*string, bool)`

GetDefaultServiceEnvironmentIdOk returns a tuple with the DefaultServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceEnvironmentId

`func (o *UpdateMarketplaceChannelRequest2) SetDefaultServiceEnvironmentId(v string)`

SetDefaultServiceEnvironmentId sets DefaultServiceEnvironmentId field to given value.

### HasDefaultServiceEnvironmentId

`func (o *UpdateMarketplaceChannelRequest2) HasDefaultServiceEnvironmentId() bool`

HasDefaultServiceEnvironmentId returns a boolean if a field has been set.

### GetDefaultServiceId

`func (o *UpdateMarketplaceChannelRequest2) GetDefaultServiceId() string`

GetDefaultServiceId returns the DefaultServiceId field if non-nil, zero value otherwise.

### GetDefaultServiceIdOk

`func (o *UpdateMarketplaceChannelRequest2) GetDefaultServiceIdOk() (*string, bool)`

GetDefaultServiceIdOk returns a tuple with the DefaultServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceId

`func (o *UpdateMarketplaceChannelRequest2) SetDefaultServiceId(v string)`

SetDefaultServiceId sets DefaultServiceId field to given value.

### HasDefaultServiceId

`func (o *UpdateMarketplaceChannelRequest2) HasDefaultServiceId() bool`

HasDefaultServiceId returns a boolean if a field has been set.

### GetDimensionMap

`func (o *UpdateMarketplaceChannelRequest2) GetDimensionMap() map[string]string`

GetDimensionMap returns the DimensionMap field if non-nil, zero value otherwise.

### GetDimensionMapOk

`func (o *UpdateMarketplaceChannelRequest2) GetDimensionMapOk() (*map[string]string, bool)`

GetDimensionMapOk returns a tuple with the DimensionMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionMap

`func (o *UpdateMarketplaceChannelRequest2) SetDimensionMap(v map[string]string)`

SetDimensionMap sets DimensionMap field to given value.

### HasDimensionMap

`func (o *UpdateMarketplaceChannelRequest2) HasDimensionMap() bool`

HasDimensionMap returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateMarketplaceChannelRequest2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateMarketplaceChannelRequest2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateMarketplaceChannelRequest2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateMarketplaceChannelRequest2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsvCallbackUrl

`func (o *UpdateMarketplaceChannelRequest2) GetIsvCallbackUrl() string`

GetIsvCallbackUrl returns the IsvCallbackUrl field if non-nil, zero value otherwise.

### GetIsvCallbackUrlOk

`func (o *UpdateMarketplaceChannelRequest2) GetIsvCallbackUrlOk() (*string, bool)`

GetIsvCallbackUrlOk returns a tuple with the IsvCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvCallbackUrl

`func (o *UpdateMarketplaceChannelRequest2) SetIsvCallbackUrl(v string)`

SetIsvCallbackUrl sets IsvCallbackUrl field to given value.

### HasIsvCallbackUrl

`func (o *UpdateMarketplaceChannelRequest2) HasIsvCallbackUrl() bool`

HasIsvCallbackUrl returns a boolean if a field has been set.

### GetIsvConfirmTimeoutSeconds

`func (o *UpdateMarketplaceChannelRequest2) GetIsvConfirmTimeoutSeconds() int64`

GetIsvConfirmTimeoutSeconds returns the IsvConfirmTimeoutSeconds field if non-nil, zero value otherwise.

### GetIsvConfirmTimeoutSecondsOk

`func (o *UpdateMarketplaceChannelRequest2) GetIsvConfirmTimeoutSecondsOk() (*int64, bool)`

GetIsvConfirmTimeoutSecondsOk returns a tuple with the IsvConfirmTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvConfirmTimeoutSeconds

`func (o *UpdateMarketplaceChannelRequest2) SetIsvConfirmTimeoutSeconds(v int64)`

SetIsvConfirmTimeoutSeconds sets IsvConfirmTimeoutSeconds field to given value.

### HasIsvConfirmTimeoutSeconds

`func (o *UpdateMarketplaceChannelRequest2) HasIsvConfirmTimeoutSeconds() bool`

HasIsvConfirmTimeoutSeconds returns a boolean if a field has been set.

### GetPlanMap

`func (o *UpdateMarketplaceChannelRequest2) GetPlanMap() map[string]MarketplacePlanMapping`

GetPlanMap returns the PlanMap field if non-nil, zero value otherwise.

### GetPlanMapOk

`func (o *UpdateMarketplaceChannelRequest2) GetPlanMapOk() (*map[string]MarketplacePlanMapping, bool)`

GetPlanMapOk returns a tuple with the PlanMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMap

`func (o *UpdateMarketplaceChannelRequest2) SetPlanMap(v map[string]MarketplacePlanMapping)`

SetPlanMap sets PlanMap field to given value.

### HasPlanMap

`func (o *UpdateMarketplaceChannelRequest2) HasPlanMap() bool`

HasPlanMap returns a boolean if a field has been set.

### GetPortBReceiverUrl

`func (o *UpdateMarketplaceChannelRequest2) GetPortBReceiverUrl() string`

GetPortBReceiverUrl returns the PortBReceiverUrl field if non-nil, zero value otherwise.

### GetPortBReceiverUrlOk

`func (o *UpdateMarketplaceChannelRequest2) GetPortBReceiverUrlOk() (*string, bool)`

GetPortBReceiverUrlOk returns a tuple with the PortBReceiverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortBReceiverUrl

`func (o *UpdateMarketplaceChannelRequest2) SetPortBReceiverUrl(v string)`

SetPortBReceiverUrl sets PortBReceiverUrl field to given value.

### HasPortBReceiverUrl

`func (o *UpdateMarketplaceChannelRequest2) HasPortBReceiverUrl() bool`

HasPortBReceiverUrl returns a boolean if a field has been set.

### GetPortBSigningSecret

`func (o *UpdateMarketplaceChannelRequest2) GetPortBSigningSecret() string`

GetPortBSigningSecret returns the PortBSigningSecret field if non-nil, zero value otherwise.

### GetPortBSigningSecretOk

`func (o *UpdateMarketplaceChannelRequest2) GetPortBSigningSecretOk() (*string, bool)`

GetPortBSigningSecretOk returns a tuple with the PortBSigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortBSigningSecret

`func (o *UpdateMarketplaceChannelRequest2) SetPortBSigningSecret(v string)`

SetPortBSigningSecret sets PortBSigningSecret field to given value.

### HasPortBSigningSecret

`func (o *UpdateMarketplaceChannelRequest2) HasPortBSigningSecret() bool`

HasPortBSigningSecret returns a boolean if a field has been set.

### GetSyntheticEmailDomain

`func (o *UpdateMarketplaceChannelRequest2) GetSyntheticEmailDomain() string`

GetSyntheticEmailDomain returns the SyntheticEmailDomain field if non-nil, zero value otherwise.

### GetSyntheticEmailDomainOk

`func (o *UpdateMarketplaceChannelRequest2) GetSyntheticEmailDomainOk() (*string, bool)`

GetSyntheticEmailDomainOk returns a tuple with the SyntheticEmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticEmailDomain

`func (o *UpdateMarketplaceChannelRequest2) SetSyntheticEmailDomain(v string)`

SetSyntheticEmailDomain sets SyntheticEmailDomain field to given value.

### HasSyntheticEmailDomain

`func (o *UpdateMarketplaceChannelRequest2) HasSyntheticEmailDomain() bool`

HasSyntheticEmailDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


