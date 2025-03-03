# DescribeRegionCostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | 
**EnvironmentType** | **string** | The type of service environment | 
**ExcludeCloudProviderIDs** | Pointer to **string** | The cloud provider IDs to exclude from the cost data | [optional] 
**ExcludeRegionIDs** | Pointer to **string** | The region IDs to exclude from the cost data | [optional] 
**Frequency** | **string** | The frequency of the cost data | 
**IncludeCloudProviderIDs** | Pointer to **string** | The cloud provider IDs to include in the cost data | [optional] 
**IncludeRegionIDs** | Pointer to **string** | The region IDs to include in the cost data | [optional] 
**StartDate** | **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeRegionCostRequest

`func NewDescribeRegionCostRequest(endDate time.Time, environmentType string, frequency string, startDate time.Time, token string, ) *DescribeRegionCostRequest`

NewDescribeRegionCostRequest instantiates a new DescribeRegionCostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeRegionCostRequestWithDefaults

`func NewDescribeRegionCostRequestWithDefaults() *DescribeRegionCostRequest`

NewDescribeRegionCostRequestWithDefaults instantiates a new DescribeRegionCostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *DescribeRegionCostRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DescribeRegionCostRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DescribeRegionCostRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetEnvironmentType

`func (o *DescribeRegionCostRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *DescribeRegionCostRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *DescribeRegionCostRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetExcludeCloudProviderIDs

`func (o *DescribeRegionCostRequest) GetExcludeCloudProviderIDs() string`

GetExcludeCloudProviderIDs returns the ExcludeCloudProviderIDs field if non-nil, zero value otherwise.

### GetExcludeCloudProviderIDsOk

`func (o *DescribeRegionCostRequest) GetExcludeCloudProviderIDsOk() (*string, bool)`

GetExcludeCloudProviderIDsOk returns a tuple with the ExcludeCloudProviderIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCloudProviderIDs

`func (o *DescribeRegionCostRequest) SetExcludeCloudProviderIDs(v string)`

SetExcludeCloudProviderIDs sets ExcludeCloudProviderIDs field to given value.

### HasExcludeCloudProviderIDs

`func (o *DescribeRegionCostRequest) HasExcludeCloudProviderIDs() bool`

HasExcludeCloudProviderIDs returns a boolean if a field has been set.

### GetExcludeRegionIDs

`func (o *DescribeRegionCostRequest) GetExcludeRegionIDs() string`

GetExcludeRegionIDs returns the ExcludeRegionIDs field if non-nil, zero value otherwise.

### GetExcludeRegionIDsOk

`func (o *DescribeRegionCostRequest) GetExcludeRegionIDsOk() (*string, bool)`

GetExcludeRegionIDsOk returns a tuple with the ExcludeRegionIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeRegionIDs

`func (o *DescribeRegionCostRequest) SetExcludeRegionIDs(v string)`

SetExcludeRegionIDs sets ExcludeRegionIDs field to given value.

### HasExcludeRegionIDs

`func (o *DescribeRegionCostRequest) HasExcludeRegionIDs() bool`

HasExcludeRegionIDs returns a boolean if a field has been set.

### GetFrequency

`func (o *DescribeRegionCostRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *DescribeRegionCostRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *DescribeRegionCostRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetIncludeCloudProviderIDs

`func (o *DescribeRegionCostRequest) GetIncludeCloudProviderIDs() string`

GetIncludeCloudProviderIDs returns the IncludeCloudProviderIDs field if non-nil, zero value otherwise.

### GetIncludeCloudProviderIDsOk

`func (o *DescribeRegionCostRequest) GetIncludeCloudProviderIDsOk() (*string, bool)`

GetIncludeCloudProviderIDsOk returns a tuple with the IncludeCloudProviderIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCloudProviderIDs

`func (o *DescribeRegionCostRequest) SetIncludeCloudProviderIDs(v string)`

SetIncludeCloudProviderIDs sets IncludeCloudProviderIDs field to given value.

### HasIncludeCloudProviderIDs

`func (o *DescribeRegionCostRequest) HasIncludeCloudProviderIDs() bool`

HasIncludeCloudProviderIDs returns a boolean if a field has been set.

### GetIncludeRegionIDs

`func (o *DescribeRegionCostRequest) GetIncludeRegionIDs() string`

GetIncludeRegionIDs returns the IncludeRegionIDs field if non-nil, zero value otherwise.

### GetIncludeRegionIDsOk

`func (o *DescribeRegionCostRequest) GetIncludeRegionIDsOk() (*string, bool)`

GetIncludeRegionIDsOk returns a tuple with the IncludeRegionIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRegionIDs

`func (o *DescribeRegionCostRequest) SetIncludeRegionIDs(v string)`

SetIncludeRegionIDs sets IncludeRegionIDs field to given value.

### HasIncludeRegionIDs

`func (o *DescribeRegionCostRequest) HasIncludeRegionIDs() bool`

HasIncludeRegionIDs returns a boolean if a field has been set.

### GetStartDate

`func (o *DescribeRegionCostRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribeRegionCostRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribeRegionCostRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetToken

`func (o *DescribeRegionCostRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeRegionCostRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeRegionCostRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


