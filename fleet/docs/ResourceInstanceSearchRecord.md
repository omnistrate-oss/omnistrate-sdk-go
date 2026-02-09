# ResourceInstanceSearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Name of the Infra Provider | 
**CustomTags** | Pointer to [**[]CustomTag**](CustomTag.md) | The custom tags associated with the resource instance. | [optional] 
**Description** | **string** | The instance description. | 
**Id** | **string** | The resource instance ID. | 
**Managed** | Pointer to **bool** | Is the proxy managed by Omnistrate. | [optional] 
**ManagedResourceType** | Pointer to **string** | The managed resource type of the proxy instance. | [optional] 
**PortsRegistrationStatus** | Pointer to **map[string][]int64** | The ports registration status of the ports based proxy instance. | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProductTierName** | Pointer to **string** | The product tier name of the instance. | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version of the instance. | [optional] 
**ProxyType** | Pointer to **string** | The proxy type. | [optional] 
**RegionCode** | **string** | The region code where the instance is hosted. | 
**ResourceId** | Pointer to **string** | ID of a resource | [optional] 
**ResourceName** | **string** | The name of the resource for the instance. | 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceEnvironmentName** | **string** | The service environment name of the instance. | 
**ServiceEnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**ServiceId** | **string** | ID of a Service | 
**ServiceName** | **string** | The service name of the instance. | 
**Status** | **string** | The status of an operation | 
**StatusDescription** | **string** | The instance status description. | 
**SubscriptionId** | Pointer to **string** | ID of a Subscription | [optional] 

## Methods

### NewResourceInstanceSearchRecord

`func NewResourceInstanceSearchRecord(cloudProvider string, description string, id string, productTierId string, regionCode string, resourceName string, serviceEnvironmentId string, serviceEnvironmentName string, serviceId string, serviceName string, status string, statusDescription string, ) *ResourceInstanceSearchRecord`

NewResourceInstanceSearchRecord instantiates a new ResourceInstanceSearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInstanceSearchRecordWithDefaults

`func NewResourceInstanceSearchRecordWithDefaults() *ResourceInstanceSearchRecord`

NewResourceInstanceSearchRecordWithDefaults instantiates a new ResourceInstanceSearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ResourceInstanceSearchRecord) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ResourceInstanceSearchRecord) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ResourceInstanceSearchRecord) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetCustomTags

`func (o *ResourceInstanceSearchRecord) GetCustomTags() []CustomTag`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *ResourceInstanceSearchRecord) GetCustomTagsOk() (*[]CustomTag, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *ResourceInstanceSearchRecord) SetCustomTags(v []CustomTag)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *ResourceInstanceSearchRecord) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### GetDescription

`func (o *ResourceInstanceSearchRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceInstanceSearchRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceInstanceSearchRecord) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *ResourceInstanceSearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceInstanceSearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceInstanceSearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetManaged

`func (o *ResourceInstanceSearchRecord) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ResourceInstanceSearchRecord) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ResourceInstanceSearchRecord) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ResourceInstanceSearchRecord) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetManagedResourceType

`func (o *ResourceInstanceSearchRecord) GetManagedResourceType() string`

GetManagedResourceType returns the ManagedResourceType field if non-nil, zero value otherwise.

### GetManagedResourceTypeOk

`func (o *ResourceInstanceSearchRecord) GetManagedResourceTypeOk() (*string, bool)`

GetManagedResourceTypeOk returns a tuple with the ManagedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedResourceType

`func (o *ResourceInstanceSearchRecord) SetManagedResourceType(v string)`

SetManagedResourceType sets ManagedResourceType field to given value.

### HasManagedResourceType

`func (o *ResourceInstanceSearchRecord) HasManagedResourceType() bool`

HasManagedResourceType returns a boolean if a field has been set.

### GetPortsRegistrationStatus

`func (o *ResourceInstanceSearchRecord) GetPortsRegistrationStatus() map[string][]int64`

GetPortsRegistrationStatus returns the PortsRegistrationStatus field if non-nil, zero value otherwise.

### GetPortsRegistrationStatusOk

`func (o *ResourceInstanceSearchRecord) GetPortsRegistrationStatusOk() (*map[string][]int64, bool)`

GetPortsRegistrationStatusOk returns a tuple with the PortsRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsRegistrationStatus

`func (o *ResourceInstanceSearchRecord) SetPortsRegistrationStatus(v map[string][]int64)`

SetPortsRegistrationStatus sets PortsRegistrationStatus field to given value.

### HasPortsRegistrationStatus

`func (o *ResourceInstanceSearchRecord) HasPortsRegistrationStatus() bool`

HasPortsRegistrationStatus returns a boolean if a field has been set.

### GetProductTierId

`func (o *ResourceInstanceSearchRecord) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ResourceInstanceSearchRecord) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ResourceInstanceSearchRecord) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierName

`func (o *ResourceInstanceSearchRecord) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *ResourceInstanceSearchRecord) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *ResourceInstanceSearchRecord) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.

### HasProductTierName

`func (o *ResourceInstanceSearchRecord) HasProductTierName() bool`

HasProductTierName returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *ResourceInstanceSearchRecord) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *ResourceInstanceSearchRecord) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *ResourceInstanceSearchRecord) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *ResourceInstanceSearchRecord) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetProxyType

`func (o *ResourceInstanceSearchRecord) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *ResourceInstanceSearchRecord) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *ResourceInstanceSearchRecord) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *ResourceInstanceSearchRecord) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetRegionCode

`func (o *ResourceInstanceSearchRecord) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ResourceInstanceSearchRecord) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ResourceInstanceSearchRecord) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


### GetResourceId

`func (o *ResourceInstanceSearchRecord) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceInstanceSearchRecord) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceInstanceSearchRecord) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ResourceInstanceSearchRecord) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *ResourceInstanceSearchRecord) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceInstanceSearchRecord) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceInstanceSearchRecord) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetServiceEnvironmentId

`func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *ResourceInstanceSearchRecord) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceEnvironmentName

`func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *ResourceInstanceSearchRecord) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceEnvironmentType

`func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentType() string`

GetServiceEnvironmentType returns the ServiceEnvironmentType field if non-nil, zero value otherwise.

### GetServiceEnvironmentTypeOk

`func (o *ResourceInstanceSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool)`

GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentType

`func (o *ResourceInstanceSearchRecord) SetServiceEnvironmentType(v string)`

SetServiceEnvironmentType sets ServiceEnvironmentType field to given value.

### HasServiceEnvironmentType

`func (o *ResourceInstanceSearchRecord) HasServiceEnvironmentType() bool`

HasServiceEnvironmentType returns a boolean if a field has been set.

### GetServiceId

`func (o *ResourceInstanceSearchRecord) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ResourceInstanceSearchRecord) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ResourceInstanceSearchRecord) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceName

`func (o *ResourceInstanceSearchRecord) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ResourceInstanceSearchRecord) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ResourceInstanceSearchRecord) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStatus

`func (o *ResourceInstanceSearchRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceInstanceSearchRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceInstanceSearchRecord) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *ResourceInstanceSearchRecord) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *ResourceInstanceSearchRecord) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *ResourceInstanceSearchRecord) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetSubscriptionId

`func (o *ResourceInstanceSearchRecord) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ResourceInstanceSearchRecord) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ResourceInstanceSearchRecord) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ResourceInstanceSearchRecord) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


