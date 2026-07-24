# ListWorkflowExecutionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**ServiceId** | **string** | ID of a Service | 
**StartDate** | Pointer to **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**Status** | Pointer to **string** | Filter executions by status. PENDING|RUNNING|SUCCEEDED|FAILED | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Verb** | Pointer to **string** | Filter executions by the provider-defined verb. | [optional] 

## Methods

### NewListWorkflowExecutionsRequest

`func NewListWorkflowExecutionsRequest(environmentId string, instanceId string, serviceId string, token string, ) *ListWorkflowExecutionsRequest`

NewListWorkflowExecutionsRequest instantiates a new ListWorkflowExecutionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowExecutionsRequestWithDefaults

`func NewListWorkflowExecutionsRequestWithDefaults() *ListWorkflowExecutionsRequest`

NewListWorkflowExecutionsRequestWithDefaults instantiates a new ListWorkflowExecutionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *ListWorkflowExecutionsRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListWorkflowExecutionsRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListWorkflowExecutionsRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListWorkflowExecutionsRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ListWorkflowExecutionsRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ListWorkflowExecutionsRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ListWorkflowExecutionsRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *ListWorkflowExecutionsRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListWorkflowExecutionsRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListWorkflowExecutionsRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetNextPageToken

`func (o *ListWorkflowExecutionsRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListWorkflowExecutionsRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListWorkflowExecutionsRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListWorkflowExecutionsRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListWorkflowExecutionsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListWorkflowExecutionsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListWorkflowExecutionsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListWorkflowExecutionsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetServiceId

`func (o *ListWorkflowExecutionsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListWorkflowExecutionsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListWorkflowExecutionsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStartDate

`func (o *ListWorkflowExecutionsRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListWorkflowExecutionsRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListWorkflowExecutionsRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListWorkflowExecutionsRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *ListWorkflowExecutionsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListWorkflowExecutionsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListWorkflowExecutionsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListWorkflowExecutionsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ListWorkflowExecutionsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListWorkflowExecutionsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListWorkflowExecutionsRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVerb

`func (o *ListWorkflowExecutionsRequest) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *ListWorkflowExecutionsRequest) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *ListWorkflowExecutionsRequest) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *ListWorkflowExecutionsRequest) HasVerb() bool`

HasVerb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


