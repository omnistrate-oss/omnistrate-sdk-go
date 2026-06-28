# ResourceInstanceFilterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Name of the Infra Provider | [optional] 
**EnvironmentName** | Pointer to **string** | Filter by service environment name. | [optional] 
**InstanceId** | Pointer to **string** | ID of a Resource Instance | [optional] 
**ProductTierName** | Pointer to **string** | Filter by product tier name. | [optional] 
**ProductTierVersion** | Pointer to **string** | Filter by product tier version. | [optional] 
**RegionCode** | Pointer to **string** | Filter by region code. | [optional] 
**ResourceName** | Pointer to **string** | Filter by resource name. | [optional] 
**ServiceName** | Pointer to **string** | Filter by service name. | [optional] 
**Status** | Pointer to **string** | The status of an operation | [optional] 
**SubscriptionId** | Pointer to **string** | ID of a Subscription | [optional] 

## Methods

### NewResourceInstanceFilterGroup

`func NewResourceInstanceFilterGroup() *ResourceInstanceFilterGroup`

NewResourceInstanceFilterGroup instantiates a new ResourceInstanceFilterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInstanceFilterGroupWithDefaults

`func NewResourceInstanceFilterGroupWithDefaults() *ResourceInstanceFilterGroup`

NewResourceInstanceFilterGroupWithDefaults instantiates a new ResourceInstanceFilterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ResourceInstanceFilterGroup) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ResourceInstanceFilterGroup) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ResourceInstanceFilterGroup) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ResourceInstanceFilterGroup) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *ResourceInstanceFilterGroup) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *ResourceInstanceFilterGroup) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *ResourceInstanceFilterGroup) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *ResourceInstanceFilterGroup) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetInstanceId

`func (o *ResourceInstanceFilterGroup) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ResourceInstanceFilterGroup) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ResourceInstanceFilterGroup) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ResourceInstanceFilterGroup) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetProductTierName

`func (o *ResourceInstanceFilterGroup) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *ResourceInstanceFilterGroup) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *ResourceInstanceFilterGroup) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.

### HasProductTierName

`func (o *ResourceInstanceFilterGroup) HasProductTierName() bool`

HasProductTierName returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *ResourceInstanceFilterGroup) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *ResourceInstanceFilterGroup) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *ResourceInstanceFilterGroup) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *ResourceInstanceFilterGroup) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetRegionCode

`func (o *ResourceInstanceFilterGroup) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ResourceInstanceFilterGroup) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ResourceInstanceFilterGroup) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ResourceInstanceFilterGroup) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetResourceName

`func (o *ResourceInstanceFilterGroup) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceInstanceFilterGroup) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceInstanceFilterGroup) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ResourceInstanceFilterGroup) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetServiceName

`func (o *ResourceInstanceFilterGroup) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ResourceInstanceFilterGroup) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ResourceInstanceFilterGroup) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ResourceInstanceFilterGroup) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceInstanceFilterGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceInstanceFilterGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceInstanceFilterGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceInstanceFilterGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *ResourceInstanceFilterGroup) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ResourceInstanceFilterGroup) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ResourceInstanceFilterGroup) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ResourceInstanceFilterGroup) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


