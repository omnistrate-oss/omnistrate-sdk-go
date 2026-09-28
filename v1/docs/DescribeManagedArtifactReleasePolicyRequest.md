# DescribeManagedArtifactReleasePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Name of the Infra Provider | 
**EnvironmentType** | **string** | The type of service environment | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeManagedArtifactReleasePolicyRequest

`func NewDescribeManagedArtifactReleasePolicyRequest(cloudProvider string, environmentType string, token string, ) *DescribeManagedArtifactReleasePolicyRequest`

NewDescribeManagedArtifactReleasePolicyRequest instantiates a new DescribeManagedArtifactReleasePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeManagedArtifactReleasePolicyRequestWithDefaults

`func NewDescribeManagedArtifactReleasePolicyRequestWithDefaults() *DescribeManagedArtifactReleasePolicyRequest`

NewDescribeManagedArtifactReleasePolicyRequestWithDefaults instantiates a new DescribeManagedArtifactReleasePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DescribeManagedArtifactReleasePolicyRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DescribeManagedArtifactReleasePolicyRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DescribeManagedArtifactReleasePolicyRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetEnvironmentType

`func (o *DescribeManagedArtifactReleasePolicyRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *DescribeManagedArtifactReleasePolicyRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *DescribeManagedArtifactReleasePolicyRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetToken

`func (o *DescribeManagedArtifactReleasePolicyRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeManagedArtifactReleasePolicyRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeManagedArtifactReleasePolicyRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


