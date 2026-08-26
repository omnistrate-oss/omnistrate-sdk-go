# DescribeManagedReleaseRevisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleVersion** | **string** |  | 
**IncludeArtifacts** | Pointer to **bool** |  | [optional] [default to true]
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeManagedReleaseRevisionRequest

`func NewDescribeManagedReleaseRevisionRequest(bundleVersion string, token string, ) *DescribeManagedReleaseRevisionRequest`

NewDescribeManagedReleaseRevisionRequest instantiates a new DescribeManagedReleaseRevisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeManagedReleaseRevisionRequestWithDefaults

`func NewDescribeManagedReleaseRevisionRequestWithDefaults() *DescribeManagedReleaseRevisionRequest`

NewDescribeManagedReleaseRevisionRequestWithDefaults instantiates a new DescribeManagedReleaseRevisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleVersion

`func (o *DescribeManagedReleaseRevisionRequest) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *DescribeManagedReleaseRevisionRequest) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *DescribeManagedReleaseRevisionRequest) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.


### GetIncludeArtifacts

`func (o *DescribeManagedReleaseRevisionRequest) GetIncludeArtifacts() bool`

GetIncludeArtifacts returns the IncludeArtifacts field if non-nil, zero value otherwise.

### GetIncludeArtifactsOk

`func (o *DescribeManagedReleaseRevisionRequest) GetIncludeArtifactsOk() (*bool, bool)`

GetIncludeArtifactsOk returns a tuple with the IncludeArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeArtifacts

`func (o *DescribeManagedReleaseRevisionRequest) SetIncludeArtifacts(v bool)`

SetIncludeArtifacts sets IncludeArtifacts field to given value.

### HasIncludeArtifacts

`func (o *DescribeManagedReleaseRevisionRequest) HasIncludeArtifacts() bool`

HasIncludeArtifacts returns a boolean if a field has been set.

### GetToken

`func (o *DescribeManagedReleaseRevisionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeManagedReleaseRevisionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeManagedReleaseRevisionRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


