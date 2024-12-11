# ResourceSearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The resource description. | 
**Id** | **string** | The Resource ID. | 
**Name** | **string** | The resource name. | 
**ProductTierId** | **string** | The product tier ID of the resource. | 
**ProductTierName** | **string** | The product tier name of the resource. | 
**ResourceType** | Pointer to **string** |  | [optional] 
**ServiceApiId** | **string** | The service API ID of the resource. | 
**ServiceEnvironmentId** | **string** | The service environment ID of the resource. | 
**ServiceEnvironmentName** | **string** | The service environment name of the resource. | 
**ServiceEnvironmentType** | Pointer to **string** | The service environment type of the resource. | [optional] 
**ServiceId** | **string** | The service ID of the resource. | 
**ServiceModelId** | **string** | The service model ID of the resource. | 
**ServiceName** | **string** | The service name of the resource. | 

## Methods

### NewResourceSearchRecord

`func NewResourceSearchRecord(description string, id string, name string, productTierId string, productTierName string, serviceApiId string, serviceEnvironmentId string, serviceEnvironmentName string, serviceId string, serviceModelId string, serviceName string, ) *ResourceSearchRecord`

NewResourceSearchRecord instantiates a new ResourceSearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSearchRecordWithDefaults

`func NewResourceSearchRecordWithDefaults() *ResourceSearchRecord`

NewResourceSearchRecordWithDefaults instantiates a new ResourceSearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ResourceSearchRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceSearchRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceSearchRecord) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *ResourceSearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceSearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceSearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ResourceSearchRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceSearchRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceSearchRecord) SetName(v string)`

SetName sets Name field to given value.


### GetProductTierId

`func (o *ResourceSearchRecord) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ResourceSearchRecord) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ResourceSearchRecord) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierName

`func (o *ResourceSearchRecord) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *ResourceSearchRecord) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *ResourceSearchRecord) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetResourceType

`func (o *ResourceSearchRecord) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceSearchRecord) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceSearchRecord) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceSearchRecord) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetServiceApiId

`func (o *ResourceSearchRecord) GetServiceApiId() string`

GetServiceApiId returns the ServiceApiId field if non-nil, zero value otherwise.

### GetServiceApiIdOk

`func (o *ResourceSearchRecord) GetServiceApiIdOk() (*string, bool)`

GetServiceApiIdOk returns a tuple with the ServiceApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiId

`func (o *ResourceSearchRecord) SetServiceApiId(v string)`

SetServiceApiId sets ServiceApiId field to given value.


### GetServiceEnvironmentId

`func (o *ResourceSearchRecord) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *ResourceSearchRecord) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *ResourceSearchRecord) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceEnvironmentName

`func (o *ResourceSearchRecord) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *ResourceSearchRecord) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *ResourceSearchRecord) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceEnvironmentType

`func (o *ResourceSearchRecord) GetServiceEnvironmentType() string`

GetServiceEnvironmentType returns the ServiceEnvironmentType field if non-nil, zero value otherwise.

### GetServiceEnvironmentTypeOk

`func (o *ResourceSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool)`

GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentType

`func (o *ResourceSearchRecord) SetServiceEnvironmentType(v string)`

SetServiceEnvironmentType sets ServiceEnvironmentType field to given value.

### HasServiceEnvironmentType

`func (o *ResourceSearchRecord) HasServiceEnvironmentType() bool`

HasServiceEnvironmentType returns a boolean if a field has been set.

### GetServiceId

`func (o *ResourceSearchRecord) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ResourceSearchRecord) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ResourceSearchRecord) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceModelId

`func (o *ResourceSearchRecord) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *ResourceSearchRecord) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *ResourceSearchRecord) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetServiceName

`func (o *ResourceSearchRecord) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ResourceSearchRecord) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ResourceSearchRecord) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


