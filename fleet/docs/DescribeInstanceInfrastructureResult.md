# DescribeInstanceInfrastructureResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **string** | ID of a Resource Instance | [optional] 
**ResourcesInfrastructure** | Pointer to [**map[string]InstanceInfrastructure**](InstanceInfrastructure.md) | Live infrastructure by instance resource. Each requested dimension reports its own availability, reason and observation time; unrequested dimensions are omitted. | [optional] 

## Methods

### NewDescribeInstanceInfrastructureResult

`func NewDescribeInstanceInfrastructureResult() *DescribeInstanceInfrastructureResult`

NewDescribeInstanceInfrastructureResult instantiates a new DescribeInstanceInfrastructureResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeInstanceInfrastructureResultWithDefaults

`func NewDescribeInstanceInfrastructureResultWithDefaults() *DescribeInstanceInfrastructureResult`

NewDescribeInstanceInfrastructureResultWithDefaults instantiates a new DescribeInstanceInfrastructureResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *DescribeInstanceInfrastructureResult) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *DescribeInstanceInfrastructureResult) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *DescribeInstanceInfrastructureResult) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *DescribeInstanceInfrastructureResult) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetResourcesInfrastructure

`func (o *DescribeInstanceInfrastructureResult) GetResourcesInfrastructure() map[string]InstanceInfrastructure`

GetResourcesInfrastructure returns the ResourcesInfrastructure field if non-nil, zero value otherwise.

### GetResourcesInfrastructureOk

`func (o *DescribeInstanceInfrastructureResult) GetResourcesInfrastructureOk() (*map[string]InstanceInfrastructure, bool)`

GetResourcesInfrastructureOk returns a tuple with the ResourcesInfrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesInfrastructure

`func (o *DescribeInstanceInfrastructureResult) SetResourcesInfrastructure(v map[string]InstanceInfrastructure)`

SetResourcesInfrastructure sets ResourcesInfrastructure field to given value.

### HasResourcesInfrastructure

`func (o *DescribeInstanceInfrastructureResult) HasResourcesInfrastructure() bool`

HasResourcesInfrastructure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


