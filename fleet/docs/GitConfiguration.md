# GitConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | The access token | [optional] 
**CommitSHA** | Pointer to **string** | The commit SHA to checkout | [optional] 
**ReferenceName** | **string** | The reference name of the repository | 
**RepositoryUrl** | **string** | The URL of the repository | 
**UserName** | Pointer to **string** | The name of github user | [optional] 

## Methods

### NewGitConfiguration

`func NewGitConfiguration(referenceName string, repositoryUrl string, ) *GitConfiguration`

NewGitConfiguration instantiates a new GitConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitConfigurationWithDefaults

`func NewGitConfigurationWithDefaults() *GitConfiguration`

NewGitConfigurationWithDefaults instantiates a new GitConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GitConfiguration) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GitConfiguration) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GitConfiguration) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *GitConfiguration) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetCommitSHA

`func (o *GitConfiguration) GetCommitSHA() string`

GetCommitSHA returns the CommitSHA field if non-nil, zero value otherwise.

### GetCommitSHAOk

`func (o *GitConfiguration) GetCommitSHAOk() (*string, bool)`

GetCommitSHAOk returns a tuple with the CommitSHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSHA

`func (o *GitConfiguration) SetCommitSHA(v string)`

SetCommitSHA sets CommitSHA field to given value.

### HasCommitSHA

`func (o *GitConfiguration) HasCommitSHA() bool`

HasCommitSHA returns a boolean if a field has been set.

### GetReferenceName

`func (o *GitConfiguration) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *GitConfiguration) GetReferenceNameOk() (*string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceName

`func (o *GitConfiguration) SetReferenceName(v string)`

SetReferenceName sets ReferenceName field to given value.


### GetRepositoryUrl

`func (o *GitConfiguration) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *GitConfiguration) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *GitConfiguration) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetUserName

`func (o *GitConfiguration) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GitConfiguration) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GitConfiguration) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GitConfiguration) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


