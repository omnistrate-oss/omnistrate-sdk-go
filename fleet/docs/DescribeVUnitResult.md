# DescribeVUnitResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Name of the Infra Provider | 
**Id** | Pointer to **string** | ID of a VUnit | [optional] 
**NetworkIds** | **[]string** | List of network IDs in the given context | 
**Region** | **string** | Region code specific to the cloud-provider | 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**ServiceModelId** | **string** | ID of a Service Model | 

## Methods

### NewDescribeVUnitResult

`func NewDescribeVUnitResult(cloudProvider string, networkIds []string, region string, serviceModelId string, ) *DescribeVUnitResult`

NewDescribeVUnitResult instantiates a new DescribeVUnitResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeVUnitResultWithDefaults

`func NewDescribeVUnitResultWithDefaults() *DescribeVUnitResult`

NewDescribeVUnitResultWithDefaults instantiates a new DescribeVUnitResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DescribeVUnitResult) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DescribeVUnitResult) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DescribeVUnitResult) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetId

`func (o *DescribeVUnitResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeVUnitResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeVUnitResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DescribeVUnitResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkIds

`func (o *DescribeVUnitResult) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *DescribeVUnitResult) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *DescribeVUnitResult) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.


### GetRegion

`func (o *DescribeVUnitResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DescribeVUnitResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DescribeVUnitResult) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetServiceId

`func (o *DescribeVUnitResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeVUnitResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeVUnitResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DescribeVUnitResult) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceModelId

`func (o *DescribeVUnitResult) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *DescribeVUnitResult) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *DescribeVUnitResult) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


