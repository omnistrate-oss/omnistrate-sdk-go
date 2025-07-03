# GitFileConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | The access token | [optional] 
**CommitSHA** | Pointer to **string** | The commit SHA to checkout | [optional] 
**Path** | **string** | The relative file path from repository root | 
**ReferenceName** | **string** | The reference name of the repository | 
**RepositoryUrl** | **string** | The URL of the repository | 
**UserName** | Pointer to **string** | The name of github user | [optional] 

## Methods

### NewGitFileConfiguration

`func NewGitFileConfiguration(path string, referenceName string, repositoryUrl string, ) *GitFileConfiguration`

NewGitFileConfiguration instantiates a new GitFileConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitFileConfigurationWithDefaults

`func NewGitFileConfigurationWithDefaults() *GitFileConfiguration`

NewGitFileConfigurationWithDefaults instantiates a new GitFileConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GitFileConfiguration) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GitFileConfiguration) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GitFileConfiguration) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *GitFileConfiguration) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetCommitSHA

`func (o *GitFileConfiguration) GetCommitSHA() string`

GetCommitSHA returns the CommitSHA field if non-nil, zero value otherwise.

### GetCommitSHAOk

`func (o *GitFileConfiguration) GetCommitSHAOk() (*string, bool)`

GetCommitSHAOk returns a tuple with the CommitSHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSHA

`func (o *GitFileConfiguration) SetCommitSHA(v string)`

SetCommitSHA sets CommitSHA field to given value.

### HasCommitSHA

`func (o *GitFileConfiguration) HasCommitSHA() bool`

HasCommitSHA returns a boolean if a field has been set.

### GetPath

`func (o *GitFileConfiguration) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitFileConfiguration) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitFileConfiguration) SetPath(v string)`

SetPath sets Path field to given value.


### GetReferenceName

`func (o *GitFileConfiguration) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *GitFileConfiguration) GetReferenceNameOk() (*string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceName

`func (o *GitFileConfiguration) SetReferenceName(v string)`

SetReferenceName sets ReferenceName field to given value.


### GetRepositoryUrl

`func (o *GitFileConfiguration) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *GitFileConfiguration) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *GitFileConfiguration) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetUserName

`func (o *GitFileConfiguration) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GitFileConfiguration) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GitFileConfiguration) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GitFileConfiguration) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


