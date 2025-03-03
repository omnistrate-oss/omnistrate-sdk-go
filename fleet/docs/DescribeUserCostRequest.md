# DescribeUserCostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | 
**EnvironmentType** | **string** | The type of service environment | 
**ExcludeUserIDs** | Pointer to **string** | The user IDs to exclude from the cost data | [optional] 
**IncludeUserIDs** | Pointer to **string** | The user IDs to include in the cost data | [optional] 
**StartDate** | **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | 
**Token** | **string** | JWT token used to perform authorization | 
**TopNInstances** | Pointer to **int64** | The number of top instances to return | [optional] 
**TopNUsers** | Pointer to **int64** | The number of top users to return | [optional] 

## Methods

### NewDescribeUserCostRequest

`func NewDescribeUserCostRequest(endDate time.Time, environmentType string, startDate time.Time, token string, ) *DescribeUserCostRequest`

NewDescribeUserCostRequest instantiates a new DescribeUserCostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUserCostRequestWithDefaults

`func NewDescribeUserCostRequestWithDefaults() *DescribeUserCostRequest`

NewDescribeUserCostRequestWithDefaults instantiates a new DescribeUserCostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *DescribeUserCostRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DescribeUserCostRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DescribeUserCostRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetEnvironmentType

`func (o *DescribeUserCostRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *DescribeUserCostRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *DescribeUserCostRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetExcludeUserIDs

`func (o *DescribeUserCostRequest) GetExcludeUserIDs() string`

GetExcludeUserIDs returns the ExcludeUserIDs field if non-nil, zero value otherwise.

### GetExcludeUserIDsOk

`func (o *DescribeUserCostRequest) GetExcludeUserIDsOk() (*string, bool)`

GetExcludeUserIDsOk returns a tuple with the ExcludeUserIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUserIDs

`func (o *DescribeUserCostRequest) SetExcludeUserIDs(v string)`

SetExcludeUserIDs sets ExcludeUserIDs field to given value.

### HasExcludeUserIDs

`func (o *DescribeUserCostRequest) HasExcludeUserIDs() bool`

HasExcludeUserIDs returns a boolean if a field has been set.

### GetIncludeUserIDs

`func (o *DescribeUserCostRequest) GetIncludeUserIDs() string`

GetIncludeUserIDs returns the IncludeUserIDs field if non-nil, zero value otherwise.

### GetIncludeUserIDsOk

`func (o *DescribeUserCostRequest) GetIncludeUserIDsOk() (*string, bool)`

GetIncludeUserIDsOk returns a tuple with the IncludeUserIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUserIDs

`func (o *DescribeUserCostRequest) SetIncludeUserIDs(v string)`

SetIncludeUserIDs sets IncludeUserIDs field to given value.

### HasIncludeUserIDs

`func (o *DescribeUserCostRequest) HasIncludeUserIDs() bool`

HasIncludeUserIDs returns a boolean if a field has been set.

### GetStartDate

`func (o *DescribeUserCostRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribeUserCostRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribeUserCostRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetToken

`func (o *DescribeUserCostRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeUserCostRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeUserCostRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetTopNInstances

`func (o *DescribeUserCostRequest) GetTopNInstances() int64`

GetTopNInstances returns the TopNInstances field if non-nil, zero value otherwise.

### GetTopNInstancesOk

`func (o *DescribeUserCostRequest) GetTopNInstancesOk() (*int64, bool)`

GetTopNInstancesOk returns a tuple with the TopNInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopNInstances

`func (o *DescribeUserCostRequest) SetTopNInstances(v int64)`

SetTopNInstances sets TopNInstances field to given value.

### HasTopNInstances

`func (o *DescribeUserCostRequest) HasTopNInstances() bool`

HasTopNInstances returns a boolean if a field has been set.

### GetTopNUsers

`func (o *DescribeUserCostRequest) GetTopNUsers() int64`

GetTopNUsers returns the TopNUsers field if non-nil, zero value otherwise.

### GetTopNUsersOk

`func (o *DescribeUserCostRequest) GetTopNUsersOk() (*int64, bool)`

GetTopNUsersOk returns a tuple with the TopNUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopNUsers

`func (o *DescribeUserCostRequest) SetTopNUsers(v int64)`

SetTopNUsers sets TopNUsers field to given value.

### HasTopNUsers

`func (o *DescribeUserCostRequest) HasTopNUsers() bool`

HasTopNUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


