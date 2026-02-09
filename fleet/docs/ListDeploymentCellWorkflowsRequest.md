# ListDeploymentCellWorkflowsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**HostClusterID** | Pointer to **string** | ID of the Host Cluster to filter Deployment Cell Workflows | [optional] 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**StartDate** | Pointer to **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListDeploymentCellWorkflowsRequest

`func NewListDeploymentCellWorkflowsRequest(token string, ) *ListDeploymentCellWorkflowsRequest`

NewListDeploymentCellWorkflowsRequest instantiates a new ListDeploymentCellWorkflowsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeploymentCellWorkflowsRequestWithDefaults

`func NewListDeploymentCellWorkflowsRequestWithDefaults() *ListDeploymentCellWorkflowsRequest`

NewListDeploymentCellWorkflowsRequestWithDefaults instantiates a new ListDeploymentCellWorkflowsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *ListDeploymentCellWorkflowsRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListDeploymentCellWorkflowsRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListDeploymentCellWorkflowsRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListDeploymentCellWorkflowsRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetHostClusterID

`func (o *ListDeploymentCellWorkflowsRequest) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *ListDeploymentCellWorkflowsRequest) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *ListDeploymentCellWorkflowsRequest) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.

### HasHostClusterID

`func (o *ListDeploymentCellWorkflowsRequest) HasHostClusterID() bool`

HasHostClusterID returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListDeploymentCellWorkflowsRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListDeploymentCellWorkflowsRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListDeploymentCellWorkflowsRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListDeploymentCellWorkflowsRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListDeploymentCellWorkflowsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListDeploymentCellWorkflowsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListDeploymentCellWorkflowsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListDeploymentCellWorkflowsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStartDate

`func (o *ListDeploymentCellWorkflowsRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListDeploymentCellWorkflowsRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListDeploymentCellWorkflowsRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListDeploymentCellWorkflowsRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetToken

`func (o *ListDeploymentCellWorkflowsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListDeploymentCellWorkflowsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListDeploymentCellWorkflowsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


