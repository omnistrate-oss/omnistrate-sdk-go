# FleetCreateResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider name | [optional] 
**CustomNetworkId** | Pointer to **string** | Custom network for resource | [optional] 
**CustomTags** | Pointer to [**[]CustomTag**](CustomTag.md) | The custom tags for the resource instance. | [optional] 
**ExternalPayerId** | Pointer to **string** | The external payer id to record which customer should pay for this resource instance. This will override the subscription level external payer id if set. | [optional] 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**OnpremPlatform** | Pointer to **string** | OnPrem platform | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version | [optional] 
**Region** | Pointer to **string** | The region code | [optional] 
**RequestParams** | Pointer to **interface{}** | The request parameters | [optional] 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 

## Methods

### NewFleetCreateResourceInstanceRequest2

`func NewFleetCreateResourceInstanceRequest2() *FleetCreateResourceInstanceRequest2`

NewFleetCreateResourceInstanceRequest2 instantiates a new FleetCreateResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCreateResourceInstanceRequest2WithDefaults

`func NewFleetCreateResourceInstanceRequest2WithDefaults() *FleetCreateResourceInstanceRequest2`

NewFleetCreateResourceInstanceRequest2WithDefaults instantiates a new FleetCreateResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *FleetCreateResourceInstanceRequest2) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *FleetCreateResourceInstanceRequest2) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *FleetCreateResourceInstanceRequest2) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *FleetCreateResourceInstanceRequest2) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCustomNetworkId

`func (o *FleetCreateResourceInstanceRequest2) GetCustomNetworkId() string`

GetCustomNetworkId returns the CustomNetworkId field if non-nil, zero value otherwise.

### GetCustomNetworkIdOk

`func (o *FleetCreateResourceInstanceRequest2) GetCustomNetworkIdOk() (*string, bool)`

GetCustomNetworkIdOk returns a tuple with the CustomNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworkId

`func (o *FleetCreateResourceInstanceRequest2) SetCustomNetworkId(v string)`

SetCustomNetworkId sets CustomNetworkId field to given value.

### HasCustomNetworkId

`func (o *FleetCreateResourceInstanceRequest2) HasCustomNetworkId() bool`

HasCustomNetworkId returns a boolean if a field has been set.

### GetCustomTags

`func (o *FleetCreateResourceInstanceRequest2) GetCustomTags() []CustomTag`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *FleetCreateResourceInstanceRequest2) GetCustomTagsOk() (*[]CustomTag, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *FleetCreateResourceInstanceRequest2) SetCustomTags(v []CustomTag)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *FleetCreateResourceInstanceRequest2) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### GetExternalPayerId

`func (o *FleetCreateResourceInstanceRequest2) GetExternalPayerId() string`

GetExternalPayerId returns the ExternalPayerId field if non-nil, zero value otherwise.

### GetExternalPayerIdOk

`func (o *FleetCreateResourceInstanceRequest2) GetExternalPayerIdOk() (*string, bool)`

GetExternalPayerIdOk returns a tuple with the ExternalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayerId

`func (o *FleetCreateResourceInstanceRequest2) SetExternalPayerId(v string)`

SetExternalPayerId sets ExternalPayerId field to given value.

### HasExternalPayerId

`func (o *FleetCreateResourceInstanceRequest2) HasExternalPayerId() bool`

HasExternalPayerId returns a boolean if a field has been set.

### GetNetworkType

`func (o *FleetCreateResourceInstanceRequest2) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *FleetCreateResourceInstanceRequest2) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *FleetCreateResourceInstanceRequest2) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *FleetCreateResourceInstanceRequest2) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetOnpremPlatform

`func (o *FleetCreateResourceInstanceRequest2) GetOnpremPlatform() string`

GetOnpremPlatform returns the OnpremPlatform field if non-nil, zero value otherwise.

### GetOnpremPlatformOk

`func (o *FleetCreateResourceInstanceRequest2) GetOnpremPlatformOk() (*string, bool)`

GetOnpremPlatformOk returns a tuple with the OnpremPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnpremPlatform

`func (o *FleetCreateResourceInstanceRequest2) SetOnpremPlatform(v string)`

SetOnpremPlatform sets OnpremPlatform field to given value.

### HasOnpremPlatform

`func (o *FleetCreateResourceInstanceRequest2) HasOnpremPlatform() bool`

HasOnpremPlatform returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *FleetCreateResourceInstanceRequest2) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *FleetCreateResourceInstanceRequest2) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *FleetCreateResourceInstanceRequest2) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *FleetCreateResourceInstanceRequest2) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetRegion

`func (o *FleetCreateResourceInstanceRequest2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FleetCreateResourceInstanceRequest2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FleetCreateResourceInstanceRequest2) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FleetCreateResourceInstanceRequest2) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRequestParams

`func (o *FleetCreateResourceInstanceRequest2) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *FleetCreateResourceInstanceRequest2) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *FleetCreateResourceInstanceRequest2) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *FleetCreateResourceInstanceRequest2) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *FleetCreateResourceInstanceRequest2) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *FleetCreateResourceInstanceRequest2) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil
### GetSubscriptionId

`func (o *FleetCreateResourceInstanceRequest2) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *FleetCreateResourceInstanceRequest2) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *FleetCreateResourceInstanceRequest2) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *FleetCreateResourceInstanceRequest2) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


