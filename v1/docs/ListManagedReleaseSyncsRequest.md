# ListManagedReleaseSyncsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeArtifactResults** | Pointer to **bool** |  | [optional] [default to false]
**OrganizationId** | **string** |  | 
**ProvisionerTargetId** | **string** |  | 
**Status** | Pointer to **string** | Lifecycle status for a managed release revision or provisioner-target sync. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListManagedReleaseSyncsRequest

`func NewListManagedReleaseSyncsRequest(organizationId string, provisionerTargetId string, token string, ) *ListManagedReleaseSyncsRequest`

NewListManagedReleaseSyncsRequest instantiates a new ListManagedReleaseSyncsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListManagedReleaseSyncsRequestWithDefaults

`func NewListManagedReleaseSyncsRequestWithDefaults() *ListManagedReleaseSyncsRequest`

NewListManagedReleaseSyncsRequestWithDefaults instantiates a new ListManagedReleaseSyncsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeArtifactResults

`func (o *ListManagedReleaseSyncsRequest) GetIncludeArtifactResults() bool`

GetIncludeArtifactResults returns the IncludeArtifactResults field if non-nil, zero value otherwise.

### GetIncludeArtifactResultsOk

`func (o *ListManagedReleaseSyncsRequest) GetIncludeArtifactResultsOk() (*bool, bool)`

GetIncludeArtifactResultsOk returns a tuple with the IncludeArtifactResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeArtifactResults

`func (o *ListManagedReleaseSyncsRequest) SetIncludeArtifactResults(v bool)`

SetIncludeArtifactResults sets IncludeArtifactResults field to given value.

### HasIncludeArtifactResults

`func (o *ListManagedReleaseSyncsRequest) HasIncludeArtifactResults() bool`

HasIncludeArtifactResults returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ListManagedReleaseSyncsRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ListManagedReleaseSyncsRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ListManagedReleaseSyncsRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProvisionerTargetId

`func (o *ListManagedReleaseSyncsRequest) GetProvisionerTargetId() string`

GetProvisionerTargetId returns the ProvisionerTargetId field if non-nil, zero value otherwise.

### GetProvisionerTargetIdOk

`func (o *ListManagedReleaseSyncsRequest) GetProvisionerTargetIdOk() (*string, bool)`

GetProvisionerTargetIdOk returns a tuple with the ProvisionerTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionerTargetId

`func (o *ListManagedReleaseSyncsRequest) SetProvisionerTargetId(v string)`

SetProvisionerTargetId sets ProvisionerTargetId field to given value.


### GetStatus

`func (o *ListManagedReleaseSyncsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListManagedReleaseSyncsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListManagedReleaseSyncsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListManagedReleaseSyncsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ListManagedReleaseSyncsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListManagedReleaseSyncsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListManagedReleaseSyncsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


