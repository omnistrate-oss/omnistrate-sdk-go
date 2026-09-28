# ListManagedArtifactReleasesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Opaque token to retrieve the next page of results. | [optional] 
**Releases** | [**[]ManagedArtifactRelease**](ManagedArtifactRelease.md) |  | 

## Methods

### NewListManagedArtifactReleasesResult

`func NewListManagedArtifactReleasesResult(releases []ManagedArtifactRelease, ) *ListManagedArtifactReleasesResult`

NewListManagedArtifactReleasesResult instantiates a new ListManagedArtifactReleasesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListManagedArtifactReleasesResultWithDefaults

`func NewListManagedArtifactReleasesResultWithDefaults() *ListManagedArtifactReleasesResult`

NewListManagedArtifactReleasesResultWithDefaults instantiates a new ListManagedArtifactReleasesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListManagedArtifactReleasesResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListManagedArtifactReleasesResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListManagedArtifactReleasesResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListManagedArtifactReleasesResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetReleases

`func (o *ListManagedArtifactReleasesResult) GetReleases() []ManagedArtifactRelease`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *ListManagedArtifactReleasesResult) GetReleasesOk() (*[]ManagedArtifactRelease, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *ListManagedArtifactReleasesResult) SetReleases(v []ManagedArtifactRelease)`

SetReleases sets Releases field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


