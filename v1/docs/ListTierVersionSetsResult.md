# ListTierVersionSetsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**TierVersionSets** | [**[]TierVersionSet**](TierVersionSet.md) | List of product-tier version sets. | 

## Methods

### NewListTierVersionSetsResult

`func NewListTierVersionSetsResult(tierVersionSets []TierVersionSet, ) *ListTierVersionSetsResult`

NewListTierVersionSetsResult instantiates a new ListTierVersionSetsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTierVersionSetsResultWithDefaults

`func NewListTierVersionSetsResultWithDefaults() *ListTierVersionSetsResult`

NewListTierVersionSetsResultWithDefaults instantiates a new ListTierVersionSetsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListTierVersionSetsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListTierVersionSetsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListTierVersionSetsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListTierVersionSetsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTierVersionSets

`func (o *ListTierVersionSetsResult) GetTierVersionSets() []TierVersionSet`

GetTierVersionSets returns the TierVersionSets field if non-nil, zero value otherwise.

### GetTierVersionSetsOk

`func (o *ListTierVersionSetsResult) GetTierVersionSetsOk() (*[]TierVersionSet, bool)`

GetTierVersionSetsOk returns a tuple with the TierVersionSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierVersionSets

`func (o *ListTierVersionSetsResult) SetTierVersionSets(v []TierVersionSet)`

SetTierVersionSets sets TierVersionSets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


