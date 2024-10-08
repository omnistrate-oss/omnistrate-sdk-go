# ListCustomDomainResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDomains** | Pointer to [**[]DescribeCustomDomainResult**](DescribeCustomDomainResult.md) | The list of custom domains | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**NextPageToken** | Pointer to **string** | Token to use for the next page | [optional] 

## Methods

### NewListCustomDomainResult

`func NewListCustomDomainResult() *ListCustomDomainResult`

NewListCustomDomainResult instantiates a new ListCustomDomainResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomDomainResultWithDefaults

`func NewListCustomDomainResultWithDefaults() *ListCustomDomainResult`

NewListCustomDomainResultWithDefaults instantiates a new ListCustomDomainResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDomains

`func (o *ListCustomDomainResult) GetCustomDomains() []DescribeCustomDomainResult`

GetCustomDomains returns the CustomDomains field if non-nil, zero value otherwise.

### GetCustomDomainsOk

`func (o *ListCustomDomainResult) GetCustomDomainsOk() (*[]DescribeCustomDomainResult, bool)`

GetCustomDomainsOk returns a tuple with the CustomDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomains

`func (o *ListCustomDomainResult) SetCustomDomains(v []DescribeCustomDomainResult)`

SetCustomDomains sets CustomDomains field to given value.

### HasCustomDomains

`func (o *ListCustomDomainResult) HasCustomDomains() bool`

HasCustomDomains returns a boolean if a field has been set.

### GetIds

`func (o *ListCustomDomainResult) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListCustomDomainResult) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListCustomDomainResult) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListCustomDomainResult) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListCustomDomainResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListCustomDomainResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListCustomDomainResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListCustomDomainResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


