# DescribeManagedReleaseSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleRevisionId** | **string** |  | 
**IncludeArtifactResults** | Pointer to **bool** |  | [optional] [default to true]
**OrganizationId** | **string** |  | 
**ProvisionerTargetId** | **string** |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeManagedReleaseSyncRequest

`func NewDescribeManagedReleaseSyncRequest(bundleRevisionId string, organizationId string, provisionerTargetId string, token string, ) *DescribeManagedReleaseSyncRequest`

NewDescribeManagedReleaseSyncRequest instantiates a new DescribeManagedReleaseSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeManagedReleaseSyncRequestWithDefaults

`func NewDescribeManagedReleaseSyncRequestWithDefaults() *DescribeManagedReleaseSyncRequest`

NewDescribeManagedReleaseSyncRequestWithDefaults instantiates a new DescribeManagedReleaseSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleRevisionId

`func (o *DescribeManagedReleaseSyncRequest) GetBundleRevisionId() string`

GetBundleRevisionId returns the BundleRevisionId field if non-nil, zero value otherwise.

### GetBundleRevisionIdOk

`func (o *DescribeManagedReleaseSyncRequest) GetBundleRevisionIdOk() (*string, bool)`

GetBundleRevisionIdOk returns a tuple with the BundleRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleRevisionId

`func (o *DescribeManagedReleaseSyncRequest) SetBundleRevisionId(v string)`

SetBundleRevisionId sets BundleRevisionId field to given value.


### GetIncludeArtifactResults

`func (o *DescribeManagedReleaseSyncRequest) GetIncludeArtifactResults() bool`

GetIncludeArtifactResults returns the IncludeArtifactResults field if non-nil, zero value otherwise.

### GetIncludeArtifactResultsOk

`func (o *DescribeManagedReleaseSyncRequest) GetIncludeArtifactResultsOk() (*bool, bool)`

GetIncludeArtifactResultsOk returns a tuple with the IncludeArtifactResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeArtifactResults

`func (o *DescribeManagedReleaseSyncRequest) SetIncludeArtifactResults(v bool)`

SetIncludeArtifactResults sets IncludeArtifactResults field to given value.

### HasIncludeArtifactResults

`func (o *DescribeManagedReleaseSyncRequest) HasIncludeArtifactResults() bool`

HasIncludeArtifactResults returns a boolean if a field has been set.

### GetOrganizationId

`func (o *DescribeManagedReleaseSyncRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DescribeManagedReleaseSyncRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DescribeManagedReleaseSyncRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProvisionerTargetId

`func (o *DescribeManagedReleaseSyncRequest) GetProvisionerTargetId() string`

GetProvisionerTargetId returns the ProvisionerTargetId field if non-nil, zero value otherwise.

### GetProvisionerTargetIdOk

`func (o *DescribeManagedReleaseSyncRequest) GetProvisionerTargetIdOk() (*string, bool)`

GetProvisionerTargetIdOk returns a tuple with the ProvisionerTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionerTargetId

`func (o *DescribeManagedReleaseSyncRequest) SetProvisionerTargetId(v string)`

SetProvisionerTargetId sets ProvisionerTargetId field to given value.


### GetToken

`func (o *DescribeManagedReleaseSyncRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeManagedReleaseSyncRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeManagedReleaseSyncRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


