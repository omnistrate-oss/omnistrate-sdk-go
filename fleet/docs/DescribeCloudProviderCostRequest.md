# DescribeCloudProviderCostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | 
**EnvironmentType** | **string** | The type of service environment | 
**ExcludeCloudProviderIDs** | Pointer to **string** | The cloud provider IDs to exclude from the cost data | [optional] 
**Frequency** | **string** | The frequency of the cost data | 
**IncludeCloudProviderIDs** | Pointer to **string** | The cloud provider IDs to include in the cost data | [optional] 
**StartDate** | **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeCloudProviderCostRequest

`func NewDescribeCloudProviderCostRequest(endDate time.Time, environmentType string, frequency string, startDate time.Time, token string, ) *DescribeCloudProviderCostRequest`

NewDescribeCloudProviderCostRequest instantiates a new DescribeCloudProviderCostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCloudProviderCostRequestWithDefaults

`func NewDescribeCloudProviderCostRequestWithDefaults() *DescribeCloudProviderCostRequest`

NewDescribeCloudProviderCostRequestWithDefaults instantiates a new DescribeCloudProviderCostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *DescribeCloudProviderCostRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DescribeCloudProviderCostRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DescribeCloudProviderCostRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetEnvironmentType

`func (o *DescribeCloudProviderCostRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *DescribeCloudProviderCostRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *DescribeCloudProviderCostRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetExcludeCloudProviderIDs

`func (o *DescribeCloudProviderCostRequest) GetExcludeCloudProviderIDs() string`

GetExcludeCloudProviderIDs returns the ExcludeCloudProviderIDs field if non-nil, zero value otherwise.

### GetExcludeCloudProviderIDsOk

`func (o *DescribeCloudProviderCostRequest) GetExcludeCloudProviderIDsOk() (*string, bool)`

GetExcludeCloudProviderIDsOk returns a tuple with the ExcludeCloudProviderIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCloudProviderIDs

`func (o *DescribeCloudProviderCostRequest) SetExcludeCloudProviderIDs(v string)`

SetExcludeCloudProviderIDs sets ExcludeCloudProviderIDs field to given value.

### HasExcludeCloudProviderIDs

`func (o *DescribeCloudProviderCostRequest) HasExcludeCloudProviderIDs() bool`

HasExcludeCloudProviderIDs returns a boolean if a field has been set.

### GetFrequency

`func (o *DescribeCloudProviderCostRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *DescribeCloudProviderCostRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *DescribeCloudProviderCostRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetIncludeCloudProviderIDs

`func (o *DescribeCloudProviderCostRequest) GetIncludeCloudProviderIDs() string`

GetIncludeCloudProviderIDs returns the IncludeCloudProviderIDs field if non-nil, zero value otherwise.

### GetIncludeCloudProviderIDsOk

`func (o *DescribeCloudProviderCostRequest) GetIncludeCloudProviderIDsOk() (*string, bool)`

GetIncludeCloudProviderIDsOk returns a tuple with the IncludeCloudProviderIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCloudProviderIDs

`func (o *DescribeCloudProviderCostRequest) SetIncludeCloudProviderIDs(v string)`

SetIncludeCloudProviderIDs sets IncludeCloudProviderIDs field to given value.

### HasIncludeCloudProviderIDs

`func (o *DescribeCloudProviderCostRequest) HasIncludeCloudProviderIDs() bool`

HasIncludeCloudProviderIDs returns a boolean if a field has been set.

### GetStartDate

`func (o *DescribeCloudProviderCostRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribeCloudProviderCostRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribeCloudProviderCostRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetToken

`func (o *DescribeCloudProviderCostRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeCloudProviderCostRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeCloudProviderCostRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


