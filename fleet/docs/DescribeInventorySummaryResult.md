# DescribeInventorySummaryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentCount** | **int64** | The number of environments in the service. | 
**ImagesCount** | **int64** | The number of images in the service. | 
**InfraConfigCount** | **int64** | The number of infrastructure configurations in the service. | 
**InstancesCount** | **int64** | The number of instances in the service. | 
**OrganizationCount** | **int64** | The number of active organizations using the service. | 
**ResourceCount** | **int64** | The number of resources in the service. | 
**UserCount** | **int64** | The number of active users using the service. | 
**EnvironmentId** | **string** | The service environment ID this workflow belongs to. | 
**ServiceId** | **string** | The service ID this workflow belongs to. | 

## Methods

### NewDescribeInventorySummaryResult

`func NewDescribeInventorySummaryResult(environmentCount int64, imagesCount int64, infraConfigCount int64, instancesCount int64, organizationCount int64, resourceCount int64, userCount int64, environmentId string, serviceId string, ) *DescribeInventorySummaryResult`

NewDescribeInventorySummaryResult instantiates a new DescribeInventorySummaryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeInventorySummaryResultWithDefaults

`func NewDescribeInventorySummaryResultWithDefaults() *DescribeInventorySummaryResult`

NewDescribeInventorySummaryResultWithDefaults instantiates a new DescribeInventorySummaryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentCount

`func (o *DescribeInventorySummaryResult) GetEnvironmentCount() int64`

GetEnvironmentCount returns the EnvironmentCount field if non-nil, zero value otherwise.

### GetEnvironmentCountOk

`func (o *DescribeInventorySummaryResult) GetEnvironmentCountOk() (*int64, bool)`

GetEnvironmentCountOk returns a tuple with the EnvironmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCount

`func (o *DescribeInventorySummaryResult) SetEnvironmentCount(v int64)`

SetEnvironmentCount sets EnvironmentCount field to given value.


### GetImagesCount

`func (o *DescribeInventorySummaryResult) GetImagesCount() int64`

GetImagesCount returns the ImagesCount field if non-nil, zero value otherwise.

### GetImagesCountOk

`func (o *DescribeInventorySummaryResult) GetImagesCountOk() (*int64, bool)`

GetImagesCountOk returns a tuple with the ImagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagesCount

`func (o *DescribeInventorySummaryResult) SetImagesCount(v int64)`

SetImagesCount sets ImagesCount field to given value.


### GetInfraConfigCount

`func (o *DescribeInventorySummaryResult) GetInfraConfigCount() int64`

GetInfraConfigCount returns the InfraConfigCount field if non-nil, zero value otherwise.

### GetInfraConfigCountOk

`func (o *DescribeInventorySummaryResult) GetInfraConfigCountOk() (*int64, bool)`

GetInfraConfigCountOk returns a tuple with the InfraConfigCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraConfigCount

`func (o *DescribeInventorySummaryResult) SetInfraConfigCount(v int64)`

SetInfraConfigCount sets InfraConfigCount field to given value.


### GetInstancesCount

`func (o *DescribeInventorySummaryResult) GetInstancesCount() int64`

GetInstancesCount returns the InstancesCount field if non-nil, zero value otherwise.

### GetInstancesCountOk

`func (o *DescribeInventorySummaryResult) GetInstancesCountOk() (*int64, bool)`

GetInstancesCountOk returns a tuple with the InstancesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesCount

`func (o *DescribeInventorySummaryResult) SetInstancesCount(v int64)`

SetInstancesCount sets InstancesCount field to given value.


### GetOrganizationCount

`func (o *DescribeInventorySummaryResult) GetOrganizationCount() int64`

GetOrganizationCount returns the OrganizationCount field if non-nil, zero value otherwise.

### GetOrganizationCountOk

`func (o *DescribeInventorySummaryResult) GetOrganizationCountOk() (*int64, bool)`

GetOrganizationCountOk returns a tuple with the OrganizationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCount

`func (o *DescribeInventorySummaryResult) SetOrganizationCount(v int64)`

SetOrganizationCount sets OrganizationCount field to given value.


### GetResourceCount

`func (o *DescribeInventorySummaryResult) GetResourceCount() int64`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *DescribeInventorySummaryResult) GetResourceCountOk() (*int64, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *DescribeInventorySummaryResult) SetResourceCount(v int64)`

SetResourceCount sets ResourceCount field to given value.


### GetUserCount

`func (o *DescribeInventorySummaryResult) GetUserCount() int64`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *DescribeInventorySummaryResult) GetUserCountOk() (*int64, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *DescribeInventorySummaryResult) SetUserCount(v int64)`

SetUserCount sets UserCount field to given value.


### GetEnvironmentId

`func (o *DescribeInventorySummaryResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DescribeInventorySummaryResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DescribeInventorySummaryResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetServiceId

`func (o *DescribeInventorySummaryResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeInventorySummaryResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeInventorySummaryResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


