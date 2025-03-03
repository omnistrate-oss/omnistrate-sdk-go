# DescribeDeploymentCellCostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | 
**EnvironmentType** | **string** | The type of service environment | 
**ExcludeCloudProviderIDs** | Pointer to **string** | The cloud provider IDs to exclude from the cost data | [optional] 
**ExcludeDeploymentCellIDs** | Pointer to **string** | The deployment cell IDs to exclude from the cost data | [optional] 
**ExcludeInstanceIDs** | Pointer to **string** | The instance IDs to exclude from the cost data | [optional] 
**ExcludeRegionIDs** | Pointer to **string** | The region IDs to exclude from the cost data | [optional] 
**Frequency** | **string** | The frequency of the cost data | 
**IncludeCloudProviderIDs** | Pointer to **string** | The cloud provider IDs to include in the cost data | [optional] 
**IncludeDeploymentCellIDs** | Pointer to **string** | The deployment cell IDs to include in the cost data | [optional] 
**IncludeInstanceIDs** | Pointer to **string** | The instance IDs to include in the cost data | [optional] 
**IncludeRegionIDs** | Pointer to **string** | The region IDs to include in the cost data | [optional] 
**StartDate** | **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | 
**Token** | **string** | JWT token used to perform authorization | 
**TopNInstances** | Pointer to **int64** | The number of top instances to return | [optional] 

## Methods

### NewDescribeDeploymentCellCostRequest

`func NewDescribeDeploymentCellCostRequest(endDate time.Time, environmentType string, frequency string, startDate time.Time, token string, ) *DescribeDeploymentCellCostRequest`

NewDescribeDeploymentCellCostRequest instantiates a new DescribeDeploymentCellCostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDeploymentCellCostRequestWithDefaults

`func NewDescribeDeploymentCellCostRequestWithDefaults() *DescribeDeploymentCellCostRequest`

NewDescribeDeploymentCellCostRequestWithDefaults instantiates a new DescribeDeploymentCellCostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *DescribeDeploymentCellCostRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DescribeDeploymentCellCostRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DescribeDeploymentCellCostRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetEnvironmentType

`func (o *DescribeDeploymentCellCostRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *DescribeDeploymentCellCostRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *DescribeDeploymentCellCostRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetExcludeCloudProviderIDs

`func (o *DescribeDeploymentCellCostRequest) GetExcludeCloudProviderIDs() string`

GetExcludeCloudProviderIDs returns the ExcludeCloudProviderIDs field if non-nil, zero value otherwise.

### GetExcludeCloudProviderIDsOk

`func (o *DescribeDeploymentCellCostRequest) GetExcludeCloudProviderIDsOk() (*string, bool)`

GetExcludeCloudProviderIDsOk returns a tuple with the ExcludeCloudProviderIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCloudProviderIDs

`func (o *DescribeDeploymentCellCostRequest) SetExcludeCloudProviderIDs(v string)`

SetExcludeCloudProviderIDs sets ExcludeCloudProviderIDs field to given value.

### HasExcludeCloudProviderIDs

`func (o *DescribeDeploymentCellCostRequest) HasExcludeCloudProviderIDs() bool`

HasExcludeCloudProviderIDs returns a boolean if a field has been set.

### GetExcludeDeploymentCellIDs

`func (o *DescribeDeploymentCellCostRequest) GetExcludeDeploymentCellIDs() string`

GetExcludeDeploymentCellIDs returns the ExcludeDeploymentCellIDs field if non-nil, zero value otherwise.

### GetExcludeDeploymentCellIDsOk

`func (o *DescribeDeploymentCellCostRequest) GetExcludeDeploymentCellIDsOk() (*string, bool)`

GetExcludeDeploymentCellIDsOk returns a tuple with the ExcludeDeploymentCellIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDeploymentCellIDs

`func (o *DescribeDeploymentCellCostRequest) SetExcludeDeploymentCellIDs(v string)`

SetExcludeDeploymentCellIDs sets ExcludeDeploymentCellIDs field to given value.

### HasExcludeDeploymentCellIDs

`func (o *DescribeDeploymentCellCostRequest) HasExcludeDeploymentCellIDs() bool`

HasExcludeDeploymentCellIDs returns a boolean if a field has been set.

### GetExcludeInstanceIDs

`func (o *DescribeDeploymentCellCostRequest) GetExcludeInstanceIDs() string`

GetExcludeInstanceIDs returns the ExcludeInstanceIDs field if non-nil, zero value otherwise.

### GetExcludeInstanceIDsOk

`func (o *DescribeDeploymentCellCostRequest) GetExcludeInstanceIDsOk() (*string, bool)`

GetExcludeInstanceIDsOk returns a tuple with the ExcludeInstanceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeInstanceIDs

`func (o *DescribeDeploymentCellCostRequest) SetExcludeInstanceIDs(v string)`

SetExcludeInstanceIDs sets ExcludeInstanceIDs field to given value.

### HasExcludeInstanceIDs

`func (o *DescribeDeploymentCellCostRequest) HasExcludeInstanceIDs() bool`

HasExcludeInstanceIDs returns a boolean if a field has been set.

### GetExcludeRegionIDs

`func (o *DescribeDeploymentCellCostRequest) GetExcludeRegionIDs() string`

GetExcludeRegionIDs returns the ExcludeRegionIDs field if non-nil, zero value otherwise.

### GetExcludeRegionIDsOk

`func (o *DescribeDeploymentCellCostRequest) GetExcludeRegionIDsOk() (*string, bool)`

GetExcludeRegionIDsOk returns a tuple with the ExcludeRegionIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeRegionIDs

`func (o *DescribeDeploymentCellCostRequest) SetExcludeRegionIDs(v string)`

SetExcludeRegionIDs sets ExcludeRegionIDs field to given value.

### HasExcludeRegionIDs

`func (o *DescribeDeploymentCellCostRequest) HasExcludeRegionIDs() bool`

HasExcludeRegionIDs returns a boolean if a field has been set.

### GetFrequency

`func (o *DescribeDeploymentCellCostRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *DescribeDeploymentCellCostRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *DescribeDeploymentCellCostRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetIncludeCloudProviderIDs

`func (o *DescribeDeploymentCellCostRequest) GetIncludeCloudProviderIDs() string`

GetIncludeCloudProviderIDs returns the IncludeCloudProviderIDs field if non-nil, zero value otherwise.

### GetIncludeCloudProviderIDsOk

`func (o *DescribeDeploymentCellCostRequest) GetIncludeCloudProviderIDsOk() (*string, bool)`

GetIncludeCloudProviderIDsOk returns a tuple with the IncludeCloudProviderIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCloudProviderIDs

`func (o *DescribeDeploymentCellCostRequest) SetIncludeCloudProviderIDs(v string)`

SetIncludeCloudProviderIDs sets IncludeCloudProviderIDs field to given value.

### HasIncludeCloudProviderIDs

`func (o *DescribeDeploymentCellCostRequest) HasIncludeCloudProviderIDs() bool`

HasIncludeCloudProviderIDs returns a boolean if a field has been set.

### GetIncludeDeploymentCellIDs

`func (o *DescribeDeploymentCellCostRequest) GetIncludeDeploymentCellIDs() string`

GetIncludeDeploymentCellIDs returns the IncludeDeploymentCellIDs field if non-nil, zero value otherwise.

### GetIncludeDeploymentCellIDsOk

`func (o *DescribeDeploymentCellCostRequest) GetIncludeDeploymentCellIDsOk() (*string, bool)`

GetIncludeDeploymentCellIDsOk returns a tuple with the IncludeDeploymentCellIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDeploymentCellIDs

`func (o *DescribeDeploymentCellCostRequest) SetIncludeDeploymentCellIDs(v string)`

SetIncludeDeploymentCellIDs sets IncludeDeploymentCellIDs field to given value.

### HasIncludeDeploymentCellIDs

`func (o *DescribeDeploymentCellCostRequest) HasIncludeDeploymentCellIDs() bool`

HasIncludeDeploymentCellIDs returns a boolean if a field has been set.

### GetIncludeInstanceIDs

`func (o *DescribeDeploymentCellCostRequest) GetIncludeInstanceIDs() string`

GetIncludeInstanceIDs returns the IncludeInstanceIDs field if non-nil, zero value otherwise.

### GetIncludeInstanceIDsOk

`func (o *DescribeDeploymentCellCostRequest) GetIncludeInstanceIDsOk() (*string, bool)`

GetIncludeInstanceIDsOk returns a tuple with the IncludeInstanceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInstanceIDs

`func (o *DescribeDeploymentCellCostRequest) SetIncludeInstanceIDs(v string)`

SetIncludeInstanceIDs sets IncludeInstanceIDs field to given value.

### HasIncludeInstanceIDs

`func (o *DescribeDeploymentCellCostRequest) HasIncludeInstanceIDs() bool`

HasIncludeInstanceIDs returns a boolean if a field has been set.

### GetIncludeRegionIDs

`func (o *DescribeDeploymentCellCostRequest) GetIncludeRegionIDs() string`

GetIncludeRegionIDs returns the IncludeRegionIDs field if non-nil, zero value otherwise.

### GetIncludeRegionIDsOk

`func (o *DescribeDeploymentCellCostRequest) GetIncludeRegionIDsOk() (*string, bool)`

GetIncludeRegionIDsOk returns a tuple with the IncludeRegionIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRegionIDs

`func (o *DescribeDeploymentCellCostRequest) SetIncludeRegionIDs(v string)`

SetIncludeRegionIDs sets IncludeRegionIDs field to given value.

### HasIncludeRegionIDs

`func (o *DescribeDeploymentCellCostRequest) HasIncludeRegionIDs() bool`

HasIncludeRegionIDs returns a boolean if a field has been set.

### GetStartDate

`func (o *DescribeDeploymentCellCostRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribeDeploymentCellCostRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribeDeploymentCellCostRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetToken

`func (o *DescribeDeploymentCellCostRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeDeploymentCellCostRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeDeploymentCellCostRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetTopNInstances

`func (o *DescribeDeploymentCellCostRequest) GetTopNInstances() int64`

GetTopNInstances returns the TopNInstances field if non-nil, zero value otherwise.

### GetTopNInstancesOk

`func (o *DescribeDeploymentCellCostRequest) GetTopNInstancesOk() (*int64, bool)`

GetTopNInstancesOk returns a tuple with the TopNInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopNInstances

`func (o *DescribeDeploymentCellCostRequest) SetTopNInstances(v int64)`

SetTopNInstances sets TopNInstances field to given value.

### HasTopNInstances

`func (o *DescribeDeploymentCellCostRequest) HasTopNInstances() bool`

HasTopNInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


