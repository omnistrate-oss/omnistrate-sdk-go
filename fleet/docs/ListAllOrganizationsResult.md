# ListAllOrganizationsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token to use for the next request | [optional] 
**Organizations** | Pointer to [**[]Organization**](Organization.md) | List of organizations. | [optional] 

## Methods

### NewListAllOrganizationsResult

`func NewListAllOrganizationsResult() *ListAllOrganizationsResult`

NewListAllOrganizationsResult instantiates a new ListAllOrganizationsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllOrganizationsResultWithDefaults

`func NewListAllOrganizationsResultWithDefaults() *ListAllOrganizationsResult`

NewListAllOrganizationsResultWithDefaults instantiates a new ListAllOrganizationsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListAllOrganizationsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListAllOrganizationsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListAllOrganizationsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListAllOrganizationsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetOrganizations

`func (o *ListAllOrganizationsResult) GetOrganizations() []Organization`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *ListAllOrganizationsResult) GetOrganizationsOk() (*[]Organization, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *ListAllOrganizationsResult) SetOrganizations(v []Organization)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *ListAllOrganizationsResult) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


