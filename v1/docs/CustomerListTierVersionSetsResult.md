# CustomerListTierVersionSetsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**TierVersionSets** | [**[]TierVersionSet**](TierVersionSet.md) | List of product-tier version sets. | 

## Methods

### NewCustomerListTierVersionSetsResult

`func NewCustomerListTierVersionSetsResult(tierVersionSets []TierVersionSet, ) *CustomerListTierVersionSetsResult`

NewCustomerListTierVersionSetsResult instantiates a new CustomerListTierVersionSetsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerListTierVersionSetsResultWithDefaults

`func NewCustomerListTierVersionSetsResultWithDefaults() *CustomerListTierVersionSetsResult`

NewCustomerListTierVersionSetsResultWithDefaults instantiates a new CustomerListTierVersionSetsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *CustomerListTierVersionSetsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CustomerListTierVersionSetsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CustomerListTierVersionSetsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CustomerListTierVersionSetsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTierVersionSets

`func (o *CustomerListTierVersionSetsResult) GetTierVersionSets() []TierVersionSet`

GetTierVersionSets returns the TierVersionSets field if non-nil, zero value otherwise.

### GetTierVersionSetsOk

`func (o *CustomerListTierVersionSetsResult) GetTierVersionSetsOk() (*[]TierVersionSet, bool)`

GetTierVersionSetsOk returns a tuple with the TierVersionSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierVersionSets

`func (o *CustomerListTierVersionSetsResult) SetTierVersionSets(v []TierVersionSet)`

SetTierVersionSets sets TierVersionSets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


