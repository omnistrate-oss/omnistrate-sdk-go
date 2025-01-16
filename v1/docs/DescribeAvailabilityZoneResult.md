# DescribeAvailabilityZoneResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**Code** | **string** | Cloud-provider native availability zone code | 
**Description** | **string** | Description of the AvailabilityZone | 
**Id** | **string** | ID of an AZ | 
**RegionCode** | **string** | Cloud-provider native region code | 

## Methods

### NewDescribeAvailabilityZoneResult

`func NewDescribeAvailabilityZoneResult(cloudProviderName string, code string, description string, id string, regionCode string, ) *DescribeAvailabilityZoneResult`

NewDescribeAvailabilityZoneResult instantiates a new DescribeAvailabilityZoneResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAvailabilityZoneResultWithDefaults

`func NewDescribeAvailabilityZoneResultWithDefaults() *DescribeAvailabilityZoneResult`

NewDescribeAvailabilityZoneResultWithDefaults instantiates a new DescribeAvailabilityZoneResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *DescribeAvailabilityZoneResult) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *DescribeAvailabilityZoneResult) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *DescribeAvailabilityZoneResult) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCode

`func (o *DescribeAvailabilityZoneResult) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DescribeAvailabilityZoneResult) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DescribeAvailabilityZoneResult) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *DescribeAvailabilityZoneResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeAvailabilityZoneResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeAvailabilityZoneResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeAvailabilityZoneResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeAvailabilityZoneResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeAvailabilityZoneResult) SetId(v string)`

SetId sets Id field to given value.


### GetRegionCode

`func (o *DescribeAvailabilityZoneResult) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *DescribeAvailabilityZoneResult) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *DescribeAvailabilityZoneResult) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


