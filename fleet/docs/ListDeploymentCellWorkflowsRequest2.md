# ListDeploymentCellWorkflowsRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListDeploymentCellWorkflowsRequest2

`func NewListDeploymentCellWorkflowsRequest2() *ListDeploymentCellWorkflowsRequest2`

NewListDeploymentCellWorkflowsRequest2 instantiates a new ListDeploymentCellWorkflowsRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeploymentCellWorkflowsRequest2WithDefaults

`func NewListDeploymentCellWorkflowsRequest2WithDefaults() *ListDeploymentCellWorkflowsRequest2`

NewListDeploymentCellWorkflowsRequest2WithDefaults instantiates a new ListDeploymentCellWorkflowsRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *ListDeploymentCellWorkflowsRequest2) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListDeploymentCellWorkflowsRequest2) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListDeploymentCellWorkflowsRequest2) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListDeploymentCellWorkflowsRequest2) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListDeploymentCellWorkflowsRequest2) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListDeploymentCellWorkflowsRequest2) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListDeploymentCellWorkflowsRequest2) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListDeploymentCellWorkflowsRequest2) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListDeploymentCellWorkflowsRequest2) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListDeploymentCellWorkflowsRequest2) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListDeploymentCellWorkflowsRequest2) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListDeploymentCellWorkflowsRequest2) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStartDate

`func (o *ListDeploymentCellWorkflowsRequest2) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListDeploymentCellWorkflowsRequest2) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListDeploymentCellWorkflowsRequest2) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListDeploymentCellWorkflowsRequest2) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


