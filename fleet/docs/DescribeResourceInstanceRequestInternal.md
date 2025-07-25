# DescribeResourceInstanceRequestInternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **bool** | Whether to include detailed information about the resource instance | [optional] [default to false]
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeResourceInstanceRequestInternal

`func NewDescribeResourceInstanceRequestInternal(environmentId string, instanceId string, serviceId string, token string, ) *DescribeResourceInstanceRequestInternal`

NewDescribeResourceInstanceRequestInternal instantiates a new DescribeResourceInstanceRequestInternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceInstanceRequestInternalWithDefaults

`func NewDescribeResourceInstanceRequestInternalWithDefaults() *DescribeResourceInstanceRequestInternal`

NewDescribeResourceInstanceRequestInternalWithDefaults instantiates a new DescribeResourceInstanceRequestInternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *DescribeResourceInstanceRequestInternal) GetDetail() bool`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DescribeResourceInstanceRequestInternal) GetDetailOk() (*bool, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DescribeResourceInstanceRequestInternal) SetDetail(v bool)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *DescribeResourceInstanceRequestInternal) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DescribeResourceInstanceRequestInternal) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DescribeResourceInstanceRequestInternal) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DescribeResourceInstanceRequestInternal) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *DescribeResourceInstanceRequestInternal) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *DescribeResourceInstanceRequestInternal) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *DescribeResourceInstanceRequestInternal) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetServiceId

`func (o *DescribeResourceInstanceRequestInternal) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeResourceInstanceRequestInternal) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeResourceInstanceRequestInternal) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *DescribeResourceInstanceRequestInternal) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeResourceInstanceRequestInternal) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeResourceInstanceRequestInternal) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


