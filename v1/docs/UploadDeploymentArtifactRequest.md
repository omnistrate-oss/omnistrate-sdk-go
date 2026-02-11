# UploadDeploymentArtifactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigID** | **string** | The account config ID associated with the deployment artifact | 
**ArtifactPath** | **string** | The path to the deployment artifact | 
**Base64EncodedArtifact** | **string** | The deployment artifact file content with base64 encoding, and expected to be a .tar.gz file | 
**EnvironmentType** | **string** | The type of service environment | 
**ProductTierName** | **string** | The name of the product tier | 
**ServiceName** | **string** | The name of the service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUploadDeploymentArtifactRequest

`func NewUploadDeploymentArtifactRequest(accountConfigID string, artifactPath string, base64EncodedArtifact string, environmentType string, productTierName string, serviceName string, token string, ) *UploadDeploymentArtifactRequest`

NewUploadDeploymentArtifactRequest instantiates a new UploadDeploymentArtifactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadDeploymentArtifactRequestWithDefaults

`func NewUploadDeploymentArtifactRequestWithDefaults() *UploadDeploymentArtifactRequest`

NewUploadDeploymentArtifactRequestWithDefaults instantiates a new UploadDeploymentArtifactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigID

`func (o *UploadDeploymentArtifactRequest) GetAccountConfigID() string`

GetAccountConfigID returns the AccountConfigID field if non-nil, zero value otherwise.

### GetAccountConfigIDOk

`func (o *UploadDeploymentArtifactRequest) GetAccountConfigIDOk() (*string, bool)`

GetAccountConfigIDOk returns a tuple with the AccountConfigID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigID

`func (o *UploadDeploymentArtifactRequest) SetAccountConfigID(v string)`

SetAccountConfigID sets AccountConfigID field to given value.


### GetArtifactPath

`func (o *UploadDeploymentArtifactRequest) GetArtifactPath() string`

GetArtifactPath returns the ArtifactPath field if non-nil, zero value otherwise.

### GetArtifactPathOk

`func (o *UploadDeploymentArtifactRequest) GetArtifactPathOk() (*string, bool)`

GetArtifactPathOk returns a tuple with the ArtifactPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactPath

`func (o *UploadDeploymentArtifactRequest) SetArtifactPath(v string)`

SetArtifactPath sets ArtifactPath field to given value.


### GetBase64EncodedArtifact

`func (o *UploadDeploymentArtifactRequest) GetBase64EncodedArtifact() string`

GetBase64EncodedArtifact returns the Base64EncodedArtifact field if non-nil, zero value otherwise.

### GetBase64EncodedArtifactOk

`func (o *UploadDeploymentArtifactRequest) GetBase64EncodedArtifactOk() (*string, bool)`

GetBase64EncodedArtifactOk returns a tuple with the Base64EncodedArtifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64EncodedArtifact

`func (o *UploadDeploymentArtifactRequest) SetBase64EncodedArtifact(v string)`

SetBase64EncodedArtifact sets Base64EncodedArtifact field to given value.


### GetEnvironmentType

`func (o *UploadDeploymentArtifactRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *UploadDeploymentArtifactRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *UploadDeploymentArtifactRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetProductTierName

`func (o *UploadDeploymentArtifactRequest) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *UploadDeploymentArtifactRequest) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *UploadDeploymentArtifactRequest) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetServiceName

`func (o *UploadDeploymentArtifactRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *UploadDeploymentArtifactRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *UploadDeploymentArtifactRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetToken

`func (o *UploadDeploymentArtifactRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UploadDeploymentArtifactRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UploadDeploymentArtifactRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


