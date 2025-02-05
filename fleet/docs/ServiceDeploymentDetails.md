# ServiceDeploymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependsOnDeployment** | Pointer to **[]string** | The deployment keys that this deployment depends on | [optional] 
**FailedReason** | Pointer to **string** | The reason why the deployment failed | [optional] 
**InstanceDeploymentAlias** | **string** | The instance deployment alias | 
**InstanceId** | Pointer to **string** | ID of a Resource Instance | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Status** | Pointer to **string** | The status of the service deployment | [optional] 
**StatusMessage** | Pointer to **string** | The status message of the service deployment | [optional] 
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

### GetFailedReason

`func (o *ServiceDeploymentDetails) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *ServiceDeploymentDetails) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *ServiceDeploymentDetails) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *ServiceDeploymentDetails) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

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


### GetStatus

`func (o *ServiceDeploymentDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceDeploymentDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceDeploymentDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceDeploymentDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ServiceDeploymentDetails) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ServiceDeploymentDetails) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ServiceDeploymentDetails) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ServiceDeploymentDetails) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

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


