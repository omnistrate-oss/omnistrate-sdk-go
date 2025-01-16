# ServerlessProxySearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The resource description of the serverless proxy. | 
**Id** | **string** | The Resource ID of the serverless proxy. | 
**Managed** | **bool** | Is the serverless proxy managed by Omnistrate. | 
**ManagedResourceType** | Pointer to **string** | The managed resource type of the serverless proxy. | [optional] 
**Name** | **string** | The resource name of the serverless proxy. | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProductTierName** | **string** | The product tier name of the serverless proxy. | 
**ProxyType** | Pointer to **string** | The proxy type of the serverless proxy. | [optional] 
**ServiceApiId** | **string** | ID of a Service API | 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceEnvironmentName** | **string** | The service environment name of the serverless proxy. | 
**ServiceEnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**ServiceId** | **string** | ID of a Service | 
**ServiceModelId** | **string** | ID of a Service Model | 
**ServiceName** | **string** | The service name of the serverless proxy. | 

## Methods

### NewServerlessProxySearchRecord

`func NewServerlessProxySearchRecord(description string, id string, managed bool, name string, productTierId string, productTierName string, serviceApiId string, serviceEnvironmentId string, serviceEnvironmentName string, serviceId string, serviceModelId string, serviceName string, ) *ServerlessProxySearchRecord`

NewServerlessProxySearchRecord instantiates a new ServerlessProxySearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessProxySearchRecordWithDefaults

`func NewServerlessProxySearchRecordWithDefaults() *ServerlessProxySearchRecord`

NewServerlessProxySearchRecordWithDefaults instantiates a new ServerlessProxySearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServerlessProxySearchRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerlessProxySearchRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerlessProxySearchRecord) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *ServerlessProxySearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerlessProxySearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerlessProxySearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetManaged

`func (o *ServerlessProxySearchRecord) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ServerlessProxySearchRecord) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ServerlessProxySearchRecord) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetManagedResourceType

`func (o *ServerlessProxySearchRecord) GetManagedResourceType() string`

GetManagedResourceType returns the ManagedResourceType field if non-nil, zero value otherwise.

### GetManagedResourceTypeOk

`func (o *ServerlessProxySearchRecord) GetManagedResourceTypeOk() (*string, bool)`

GetManagedResourceTypeOk returns a tuple with the ManagedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedResourceType

`func (o *ServerlessProxySearchRecord) SetManagedResourceType(v string)`

SetManagedResourceType sets ManagedResourceType field to given value.

### HasManagedResourceType

`func (o *ServerlessProxySearchRecord) HasManagedResourceType() bool`

HasManagedResourceType returns a boolean if a field has been set.

### GetName

`func (o *ServerlessProxySearchRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerlessProxySearchRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerlessProxySearchRecord) SetName(v string)`

SetName sets Name field to given value.


### GetProductTierId

`func (o *ServerlessProxySearchRecord) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ServerlessProxySearchRecord) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ServerlessProxySearchRecord) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierName

`func (o *ServerlessProxySearchRecord) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *ServerlessProxySearchRecord) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *ServerlessProxySearchRecord) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetProxyType

`func (o *ServerlessProxySearchRecord) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *ServerlessProxySearchRecord) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *ServerlessProxySearchRecord) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *ServerlessProxySearchRecord) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetServiceApiId

`func (o *ServerlessProxySearchRecord) GetServiceApiId() string`

GetServiceApiId returns the ServiceApiId field if non-nil, zero value otherwise.

### GetServiceApiIdOk

`func (o *ServerlessProxySearchRecord) GetServiceApiIdOk() (*string, bool)`

GetServiceApiIdOk returns a tuple with the ServiceApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiId

`func (o *ServerlessProxySearchRecord) SetServiceApiId(v string)`

SetServiceApiId sets ServiceApiId field to given value.


### GetServiceEnvironmentId

`func (o *ServerlessProxySearchRecord) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *ServerlessProxySearchRecord) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *ServerlessProxySearchRecord) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceEnvironmentName

`func (o *ServerlessProxySearchRecord) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *ServerlessProxySearchRecord) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *ServerlessProxySearchRecord) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceEnvironmentType

`func (o *ServerlessProxySearchRecord) GetServiceEnvironmentType() string`

GetServiceEnvironmentType returns the ServiceEnvironmentType field if non-nil, zero value otherwise.

### GetServiceEnvironmentTypeOk

`func (o *ServerlessProxySearchRecord) GetServiceEnvironmentTypeOk() (*string, bool)`

GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentType

`func (o *ServerlessProxySearchRecord) SetServiceEnvironmentType(v string)`

SetServiceEnvironmentType sets ServiceEnvironmentType field to given value.

### HasServiceEnvironmentType

`func (o *ServerlessProxySearchRecord) HasServiceEnvironmentType() bool`

HasServiceEnvironmentType returns a boolean if a field has been set.

### GetServiceId

`func (o *ServerlessProxySearchRecord) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServerlessProxySearchRecord) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServerlessProxySearchRecord) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceModelId

`func (o *ServerlessProxySearchRecord) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *ServerlessProxySearchRecord) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *ServerlessProxySearchRecord) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetServiceName

`func (o *ServerlessProxySearchRecord) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServerlessProxySearchRecord) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServerlessProxySearchRecord) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


