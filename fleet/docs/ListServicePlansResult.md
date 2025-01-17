# ListServicePlansResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token to use for the next page | [optional] 
**ServicePlans** | [**[]GetServicePlanResult**](GetServicePlanResult.md) | List of service plans | 

## Methods

### NewListServicePlansResult

`func NewListServicePlansResult(servicePlans []GetServicePlanResult, ) *ListServicePlansResult`

NewListServicePlansResult instantiates a new ListServicePlansResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServicePlansResultWithDefaults

`func NewListServicePlansResultWithDefaults() *ListServicePlansResult`

NewListServicePlansResultWithDefaults instantiates a new ListServicePlansResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListServicePlansResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListServicePlansResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListServicePlansResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListServicePlansResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetServicePlans

`func (o *ListServicePlansResult) GetServicePlans() []GetServicePlanResult`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *ListServicePlansResult) GetServicePlansOk() (*[]GetServicePlanResult, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlans

`func (o *ListServicePlansResult) SetServicePlans(v []GetServicePlanResult)`

SetServicePlans sets ServicePlans field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


