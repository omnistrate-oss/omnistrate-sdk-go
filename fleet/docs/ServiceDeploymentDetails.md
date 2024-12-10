# ServiceDeploymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependsOnDeployment** | Pointer to **[]string** | The deployment keys that this deployment depends on | [optional] 
**InstanceDeploymentAlias** | **string** | The instance deployment alias | 
**InstanceId** | Pointer to **string** | ID of a Resource Instance | [optional] 
**ServiceId** | **string** | ID of a Service | 
**SubscriptionId** | **string** | ID of a Subscription | 

## Methods

### NewServiceDeploymentDetails

`func NewServiceDeploymentDetails(instanceDeploymentAlias string, serviceId string, subscriptionId string, ) *ServiceDeploymentDetails`

NewServiceDeploymentDetails instantiates a new ServiceDeploymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDeploymentDetailsWithDefaults

`func NewServiceDeploymentDetailsWithDefaults() *ServiceDeploymentDetails`

NewServiceDeploymentDetailsWithDefaults instantiates a new ServiceDeploymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependsOnDeployment

`func (o *ServiceDeploymentDetails) GetDependsOnDeployment() []string`

GetDependsOnDeployment returns the DependsOnDeployment field if non-nil, zero value otherwise.

### GetDependsOnDeploymentOk

`func (o *ServiceDeploymentDetails) GetDependsOnDeploymentOk() (*[]string, bool)`

GetDependsOnDeploymentOk returns a tuple with the DependsOnDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnDeployment

`func (o *ServiceDeploymentDetails) SetDependsOnDeployment(v []string)`

SetDependsOnDeployment sets DependsOnDeployment field to given value.

### HasDependsOnDeployment

`func (o *ServiceDeploymentDetails) HasDependsOnDeployment() bool`

HasDependsOnDeployment returns a boolean if a field has been set.

### GetInstanceDeploymentAlias

`func (o *ServiceDeploymentDetails) GetInstanceDeploymentAlias() string`

GetInstanceDeploymentAlias returns the InstanceDeploymentAlias field if non-nil, zero value otherwise.

### GetInstanceDeploymentAliasOk

`func (o *ServiceDeploymentDetails) GetInstanceDeploymentAliasOk() (*string, bool)`

GetInstanceDeploymentAliasOk returns a tuple with the InstanceDeploymentAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDeploymentAlias

`func (o *ServiceDeploymentDetails) SetInstanceDeploymentAlias(v string)`

SetInstanceDeploymentAlias sets InstanceDeploymentAlias field to given value.


### GetInstanceId

`func (o *ServiceDeploymentDetails) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ServiceDeploymentDetails) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ServiceDeploymentDetails) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ServiceDeploymentDetails) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetServiceId

`func (o *ServiceDeploymentDetails) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceDeploymentDetails) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceDeploymentDetails) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSubscriptionId

`func (o *ServiceDeploymentDetails) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ServiceDeploymentDetails) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ServiceDeploymentDetails) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


