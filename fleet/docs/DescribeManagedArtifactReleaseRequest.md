# DescribeManagedArtifactReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleVersion** | **string** |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeManagedArtifactReleaseRequest

`func NewDescribeManagedArtifactReleaseRequest(bundleVersion string, token string, ) *DescribeManagedArtifactReleaseRequest`

NewDescribeManagedArtifactReleaseRequest instantiates a new DescribeManagedArtifactReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeManagedArtifactReleaseRequestWithDefaults

`func NewDescribeManagedArtifactReleaseRequestWithDefaults() *DescribeManagedArtifactReleaseRequest`

NewDescribeManagedArtifactReleaseRequestWithDefaults instantiates a new DescribeManagedArtifactReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleVersion

`func (o *DescribeManagedArtifactReleaseRequest) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *DescribeManagedArtifactReleaseRequest) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *DescribeManagedArtifactReleaseRequest) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.


### GetToken

`func (o *DescribeManagedArtifactReleaseRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeManagedArtifactReleaseRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeManagedArtifactReleaseRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


