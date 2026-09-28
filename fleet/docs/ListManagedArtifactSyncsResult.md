# ListManagedArtifactSyncsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Opaque token to retrieve the next page of results. | [optional] 
**Syncs** | [**[]ManagedArtifactSync**](ManagedArtifactSync.md) |  | 

## Methods

### NewListManagedArtifactSyncsResult

`func NewListManagedArtifactSyncsResult(syncs []ManagedArtifactSync, ) *ListManagedArtifactSyncsResult`

NewListManagedArtifactSyncsResult instantiates a new ListManagedArtifactSyncsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListManagedArtifactSyncsResultWithDefaults

`func NewListManagedArtifactSyncsResultWithDefaults() *ListManagedArtifactSyncsResult`

NewListManagedArtifactSyncsResultWithDefaults instantiates a new ListManagedArtifactSyncsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListManagedArtifactSyncsResult) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListManagedArtifactSyncsResult) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListManagedArtifactSyncsResult) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListManagedArtifactSyncsResult) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetSyncs

`func (o *ListManagedArtifactSyncsResult) GetSyncs() []ManagedArtifactSync`

GetSyncs returns the Syncs field if non-nil, zero value otherwise.

### GetSyncsOk

`func (o *ListManagedArtifactSyncsResult) GetSyncsOk() (*[]ManagedArtifactSync, bool)`

GetSyncsOk returns a tuple with the Syncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncs

`func (o *ListManagedArtifactSyncsResult) SetSyncs(v []ManagedArtifactSync)`

SetSyncs sets Syncs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


