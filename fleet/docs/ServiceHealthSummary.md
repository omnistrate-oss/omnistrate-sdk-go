# ServiceHealthSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderHealthSummary** | **map[string]interface{}** | The summary of health by cloud provider | 
**HealthyInstances** | **int64** | The number of healthy instances in the region | 
**Message** | **string** | The status message | 
**ServiceEnvironmentID** | **string** | ID of a Service Environment | 
**ServiceID** | **string** | ID of a Service | 
**Status** | **string** | The status of the service | 
**TotalInstances** | **int64** | The total number of instances in the region | 
**UnhealthyInstances** | **int64** | The number of unhealthy instances in the region | 
**UnknownInstances** | Pointer to **int64** | The number of instances with unknown status | [optional] 

## Methods

### NewServiceHealthSummary

`func NewServiceHealthSummary(cloudProviderHealthSummary map[string]interface{}, healthyInstances int64, message string, serviceEnvironmentID string, serviceID string, status string, totalInstances int64, unhealthyInstances int64, ) *ServiceHealthSummary`

NewServiceHealthSummary instantiates a new ServiceHealthSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHealthSummaryWithDefaults

`func NewServiceHealthSummaryWithDefaults() *ServiceHealthSummary`

NewServiceHealthSummaryWithDefaults instantiates a new ServiceHealthSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderHealthSummary

`func (o *ServiceHealthSummary) GetCloudProviderHealthSummary() map[string]interface{}`

GetCloudProviderHealthSummary returns the CloudProviderHealthSummary field if non-nil, zero value otherwise.

### GetCloudProviderHealthSummaryOk

`func (o *ServiceHealthSummary) GetCloudProviderHealthSummaryOk() (*map[string]interface{}, bool)`

GetCloudProviderHealthSummaryOk returns a tuple with the CloudProviderHealthSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderHealthSummary

`func (o *ServiceHealthSummary) SetCloudProviderHealthSummary(v map[string]interface{})`

SetCloudProviderHealthSummary sets CloudProviderHealthSummary field to given value.


### GetHealthyInstances

`func (o *ServiceHealthSummary) GetHealthyInstances() int64`

GetHealthyInstances returns the HealthyInstances field if non-nil, zero value otherwise.

### GetHealthyInstancesOk

`func (o *ServiceHealthSummary) GetHealthyInstancesOk() (*int64, bool)`

GetHealthyInstancesOk returns a tuple with the HealthyInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyInstances

`func (o *ServiceHealthSummary) SetHealthyInstances(v int64)`

SetHealthyInstances sets HealthyInstances field to given value.


### GetMessage

`func (o *ServiceHealthSummary) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceHealthSummary) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceHealthSummary) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetServiceEnvironmentID

`func (o *ServiceHealthSummary) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *ServiceHealthSummary) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *ServiceHealthSummary) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.


### GetServiceID

`func (o *ServiceHealthSummary) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *ServiceHealthSummary) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *ServiceHealthSummary) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.


### GetStatus

`func (o *ServiceHealthSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceHealthSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceHealthSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalInstances

`func (o *ServiceHealthSummary) GetTotalInstances() int64`

GetTotalInstances returns the TotalInstances field if non-nil, zero value otherwise.

### GetTotalInstancesOk

`func (o *ServiceHealthSummary) GetTotalInstancesOk() (*int64, bool)`

GetTotalInstancesOk returns a tuple with the TotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstances

`func (o *ServiceHealthSummary) SetTotalInstances(v int64)`

SetTotalInstances sets TotalInstances field to given value.


### GetUnhealthyInstances

`func (o *ServiceHealthSummary) GetUnhealthyInstances() int64`

GetUnhealthyInstances returns the UnhealthyInstances field if non-nil, zero value otherwise.

### GetUnhealthyInstancesOk

`func (o *ServiceHealthSummary) GetUnhealthyInstancesOk() (*int64, bool)`

GetUnhealthyInstancesOk returns a tuple with the UnhealthyInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyInstances

`func (o *ServiceHealthSummary) SetUnhealthyInstances(v int64)`

SetUnhealthyInstances sets UnhealthyInstances field to given value.


### GetUnknownInstances

`func (o *ServiceHealthSummary) GetUnknownInstances() int64`

GetUnknownInstances returns the UnknownInstances field if non-nil, zero value otherwise.

### GetUnknownInstancesOk

`func (o *ServiceHealthSummary) GetUnknownInstancesOk() (*int64, bool)`

GetUnknownInstancesOk returns a tuple with the UnknownInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownInstances

`func (o *ServiceHealthSummary) SetUnknownInstances(v int64)`

SetUnknownInstances sets UnknownInstances field to given value.

### HasUnknownInstances

`func (o *ServiceHealthSummary) HasUnknownInstances() bool`

HasUnknownInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


