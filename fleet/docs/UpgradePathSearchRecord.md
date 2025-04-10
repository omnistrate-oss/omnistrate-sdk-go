# UpgradePathSearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of an Upgrade Path | 
**NotifyCustomer** | **bool** | Whether to notify the end customer about the upgrade progress. | 
**ProductTierID** | **string** | ID of a Product Tier | 
**ProductTierName** | **string** | The product tier name of the upgrade path. | 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceEnvironmentName** | **string** | The service environment name of the upgrade path. | 
**ServiceEnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**ServiceId** | **string** | ID of a Service | 
**ServiceName** | **string** | The service name of the upgrade path. | 
**Status** | **string** | The status of the upgrade path. | 

## Methods

### NewUpgradePathSearchRecord

`func NewUpgradePathSearchRecord(id string, notifyCustomer bool, productTierID string, productTierName string, serviceEnvironmentId string, serviceEnvironmentName string, serviceId string, serviceName string, status string, ) *UpgradePathSearchRecord`

NewUpgradePathSearchRecord instantiates a new UpgradePathSearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradePathSearchRecordWithDefaults

`func NewUpgradePathSearchRecordWithDefaults() *UpgradePathSearchRecord`

NewUpgradePathSearchRecordWithDefaults instantiates a new UpgradePathSearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpgradePathSearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpgradePathSearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpgradePathSearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetNotifyCustomer

`func (o *UpgradePathSearchRecord) GetNotifyCustomer() bool`

GetNotifyCustomer returns the NotifyCustomer field if non-nil, zero value otherwise.

### GetNotifyCustomerOk

`func (o *UpgradePathSearchRecord) GetNotifyCustomerOk() (*bool, bool)`

GetNotifyCustomerOk returns a tuple with the NotifyCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCustomer

`func (o *UpgradePathSearchRecord) SetNotifyCustomer(v bool)`

SetNotifyCustomer sets NotifyCustomer field to given value.


### GetProductTierID

`func (o *UpgradePathSearchRecord) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *UpgradePathSearchRecord) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *UpgradePathSearchRecord) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.


### GetProductTierName

`func (o *UpgradePathSearchRecord) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *UpgradePathSearchRecord) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *UpgradePathSearchRecord) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetServiceEnvironmentId

`func (o *UpgradePathSearchRecord) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *UpgradePathSearchRecord) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *UpgradePathSearchRecord) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceEnvironmentName

`func (o *UpgradePathSearchRecord) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *UpgradePathSearchRecord) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *UpgradePathSearchRecord) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceEnvironmentType

`func (o *UpgradePathSearchRecord) GetServiceEnvironmentType() string`

GetServiceEnvironmentType returns the ServiceEnvironmentType field if non-nil, zero value otherwise.

### GetServiceEnvironmentTypeOk

`func (o *UpgradePathSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool)`

GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentType

`func (o *UpgradePathSearchRecord) SetServiceEnvironmentType(v string)`

SetServiceEnvironmentType sets ServiceEnvironmentType field to given value.

### HasServiceEnvironmentType

`func (o *UpgradePathSearchRecord) HasServiceEnvironmentType() bool`

HasServiceEnvironmentType returns a boolean if a field has been set.

### GetServiceId

`func (o *UpgradePathSearchRecord) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpgradePathSearchRecord) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpgradePathSearchRecord) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceName

`func (o *UpgradePathSearchRecord) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *UpgradePathSearchRecord) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *UpgradePathSearchRecord) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStatus

`func (o *UpgradePathSearchRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpgradePathSearchRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpgradePathSearchRecord) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


