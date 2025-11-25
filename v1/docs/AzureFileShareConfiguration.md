# AzureFileShareConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Redundancy** | Pointer to **string** | The redundancy level of the Azure file share | [optional] 
**ShareQuota** | **int32** | The share quota for the Azure file share | 
**Tier** | **string** | The tier of the Azure file share | 

## Methods

### NewAzureFileShareConfiguration

`func NewAzureFileShareConfiguration(shareQuota int32, tier string, ) *AzureFileShareConfiguration`

NewAzureFileShareConfiguration instantiates a new AzureFileShareConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureFileShareConfigurationWithDefaults

`func NewAzureFileShareConfigurationWithDefaults() *AzureFileShareConfiguration`

NewAzureFileShareConfigurationWithDefaults instantiates a new AzureFileShareConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedundancy

`func (o *AzureFileShareConfiguration) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *AzureFileShareConfiguration) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *AzureFileShareConfiguration) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *AzureFileShareConfiguration) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetShareQuota

`func (o *AzureFileShareConfiguration) GetShareQuota() int32`

GetShareQuota returns the ShareQuota field if non-nil, zero value otherwise.

### GetShareQuotaOk

`func (o *AzureFileShareConfiguration) GetShareQuotaOk() (*int32, bool)`

GetShareQuotaOk returns a tuple with the ShareQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareQuota

`func (o *AzureFileShareConfiguration) SetShareQuota(v int32)`

SetShareQuota sets ShareQuota field to given value.


### GetTier

`func (o *AzureFileShareConfiguration) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *AzureFileShareConfiguration) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *AzureFileShareConfiguration) SetTier(v string)`

SetTier sets Tier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


