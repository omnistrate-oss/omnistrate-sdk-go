# ProxyInstanceSearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Name of the Infra Provider | 
**Description** | **string** | The instance description. | 
**Id** | **string** | The proxy instance ID. | 
**Managed** | **bool** | Is the proxy managed by Omnistrate. | 
**ManagedResourceType** | Pointer to **string** | The managed resource type of the proxy instance. | [optional] 
**PortsRegistrationStatus** | **map[string][]int64** | The ports registration status of the ports based proxy instance. | 
**ProxyType** | Pointer to **string** | The proxy type. | [optional] 
**RegionCode** | **string** | The region code where the instance is hosted. | 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceEnvironmentName** | **string** | The service environment name of the instance. | 
**ServiceEnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**ServiceId** | **string** | ID of a Service | 
**ServiceName** | **string** | The service name of the instance. | 
**Status** | **string** | The status of an operation | 
**StatusDescription** | **string** | The instance status description. | 
**TargetResourceName** | **string** | The name of the target resource of the proxy instance. | 

## Methods

### NewProxyInstanceSearchRecord

`func NewProxyInstanceSearchRecord(cloudProvider string, description string, id string, managed bool, portsRegistrationStatus map[string][]int64, regionCode string, serviceEnvironmentId string, serviceEnvironmentName string, serviceId string, serviceName string, status string, statusDescription string, targetResourceName string, ) *ProxyInstanceSearchRecord`

NewProxyInstanceSearchRecord instantiates a new ProxyInstanceSearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyInstanceSearchRecordWithDefaults

`func NewProxyInstanceSearchRecordWithDefaults() *ProxyInstanceSearchRecord`

NewProxyInstanceSearchRecordWithDefaults instantiates a new ProxyInstanceSearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ProxyInstanceSearchRecord) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ProxyInstanceSearchRecord) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ProxyInstanceSearchRecord) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetDescription

`func (o *ProxyInstanceSearchRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProxyInstanceSearchRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProxyInstanceSearchRecord) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *ProxyInstanceSearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProxyInstanceSearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProxyInstanceSearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetManaged

`func (o *ProxyInstanceSearchRecord) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ProxyInstanceSearchRecord) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ProxyInstanceSearchRecord) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetManagedResourceType

`func (o *ProxyInstanceSearchRecord) GetManagedResourceType() string`

GetManagedResourceType returns the ManagedResourceType field if non-nil, zero value otherwise.

### GetManagedResourceTypeOk

`func (o *ProxyInstanceSearchRecord) GetManagedResourceTypeOk() (*string, bool)`

GetManagedResourceTypeOk returns a tuple with the ManagedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedResourceType

`func (o *ProxyInstanceSearchRecord) SetManagedResourceType(v string)`

SetManagedResourceType sets ManagedResourceType field to given value.

### HasManagedResourceType

`func (o *ProxyInstanceSearchRecord) HasManagedResourceType() bool`

HasManagedResourceType returns a boolean if a field has been set.

### GetPortsRegistrationStatus

`func (o *ProxyInstanceSearchRecord) GetPortsRegistrationStatus() map[string][]int64`

GetPortsRegistrationStatus returns the PortsRegistrationStatus field if non-nil, zero value otherwise.

### GetPortsRegistrationStatusOk

`func (o *ProxyInstanceSearchRecord) GetPortsRegistrationStatusOk() (*map[string][]int64, bool)`

GetPortsRegistrationStatusOk returns a tuple with the PortsRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsRegistrationStatus

`func (o *ProxyInstanceSearchRecord) SetPortsRegistrationStatus(v map[string][]int64)`

SetPortsRegistrationStatus sets PortsRegistrationStatus field to given value.


### GetProxyType

`func (o *ProxyInstanceSearchRecord) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *ProxyInstanceSearchRecord) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *ProxyInstanceSearchRecord) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *ProxyInstanceSearchRecord) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetRegionCode

`func (o *ProxyInstanceSearchRecord) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ProxyInstanceSearchRecord) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ProxyInstanceSearchRecord) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


### GetServiceEnvironmentId

`func (o *ProxyInstanceSearchRecord) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *ProxyInstanceSearchRecord) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *ProxyInstanceSearchRecord) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceEnvironmentName

`func (o *ProxyInstanceSearchRecord) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *ProxyInstanceSearchRecord) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *ProxyInstanceSearchRecord) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceEnvironmentType

`func (o *ProxyInstanceSearchRecord) GetServiceEnvironmentType() string`

GetServiceEnvironmentType returns the ServiceEnvironmentType field if non-nil, zero value otherwise.

### GetServiceEnvironmentTypeOk

`func (o *ProxyInstanceSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool)`

GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentType

`func (o *ProxyInstanceSearchRecord) SetServiceEnvironmentType(v string)`

SetServiceEnvironmentType sets ServiceEnvironmentType field to given value.

### HasServiceEnvironmentType

`func (o *ProxyInstanceSearchRecord) HasServiceEnvironmentType() bool`

HasServiceEnvironmentType returns a boolean if a field has been set.

### GetServiceId

`func (o *ProxyInstanceSearchRecord) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ProxyInstanceSearchRecord) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ProxyInstanceSearchRecord) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceName

`func (o *ProxyInstanceSearchRecord) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ProxyInstanceSearchRecord) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ProxyInstanceSearchRecord) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStatus

`func (o *ProxyInstanceSearchRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProxyInstanceSearchRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProxyInstanceSearchRecord) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *ProxyInstanceSearchRecord) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *ProxyInstanceSearchRecord) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *ProxyInstanceSearchRecord) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetTargetResourceName

`func (o *ProxyInstanceSearchRecord) GetTargetResourceName() string`

GetTargetResourceName returns the TargetResourceName field if non-nil, zero value otherwise.

### GetTargetResourceNameOk

`func (o *ProxyInstanceSearchRecord) GetTargetResourceNameOk() (*string, bool)`

GetTargetResourceNameOk returns a tuple with the TargetResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceName

`func (o *ProxyInstanceSearchRecord) SetTargetResourceName(v string)`

SetTargetResourceName sets TargetResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


