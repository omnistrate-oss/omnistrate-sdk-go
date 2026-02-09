# UploadDeploymentArtifactRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigID** | **string** | The account config ID associated with the deployment artifact | 
**ArtifactPath** | **string** | The path to the deployment artifact | 
**Base64EncodedArtifact** | **string** | The deployment artifact file content with base64 encoding, and expected to be a .tar.gz file | 
**ProductTierName** | **string** | The name of the product tier | 
**ServiceName** | **string** | The name of the service | 

## Methods

### NewUploadDeploymentArtifactRequest2

`func NewUploadDeploymentArtifactRequest2(accountConfigID string, artifactPath string, base64EncodedArtifact string, productTierName string, serviceName string, ) *UploadDeploymentArtifactRequest2`

NewUploadDeploymentArtifactRequest2 instantiates a new UploadDeploymentArtifactRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadDeploymentArtifactRequest2WithDefaults

`func NewUploadDeploymentArtifactRequest2WithDefaults() *UploadDeploymentArtifactRequest2`

NewUploadDeploymentArtifactRequest2WithDefaults instantiates a new UploadDeploymentArtifactRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigID

`func (o *UploadDeploymentArtifactRequest2) GetAccountConfigID() string`

GetAccountConfigID returns the AccountConfigID field if non-nil, zero value otherwise.

### GetAccountConfigIDOk

`func (o *UploadDeploymentArtifactRequest2) GetAccountConfigIDOk() (*string, bool)`

GetAccountConfigIDOk returns a tuple with the AccountConfigID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigID

`func (o *UploadDeploymentArtifactRequest2) SetAccountConfigID(v string)`

SetAccountConfigID sets AccountConfigID field to given value.


### GetArtifactPath

`func (o *UploadDeploymentArtifactRequest2) GetArtifactPath() string`

GetArtifactPath returns the ArtifactPath field if non-nil, zero value otherwise.

### GetArtifactPathOk

`func (o *UploadDeploymentArtifactRequest2) GetArtifactPathOk() (*string, bool)`

GetArtifactPathOk returns a tuple with the ArtifactPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactPath

`func (o *UploadDeploymentArtifactRequest2) SetArtifactPath(v string)`

SetArtifactPath sets ArtifactPath field to given value.


### GetBase64EncodedArtifact

`func (o *UploadDeploymentArtifactRequest2) GetBase64EncodedArtifact() string`

GetBase64EncodedArtifact returns the Base64EncodedArtifact field if non-nil, zero value otherwise.

### GetBase64EncodedArtifactOk

`func (o *UploadDeploymentArtifactRequest2) GetBase64EncodedArtifactOk() (*string, bool)`

GetBase64EncodedArtifactOk returns a tuple with the Base64EncodedArtifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64EncodedArtifact

`func (o *UploadDeploymentArtifactRequest2) SetBase64EncodedArtifact(v string)`

SetBase64EncodedArtifact sets Base64EncodedArtifact field to given value.


### GetProductTierName

`func (o *UploadDeploymentArtifactRequest2) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *UploadDeploymentArtifactRequest2) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *UploadDeploymentArtifactRequest2) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetServiceName

`func (o *UploadDeploymentArtifactRequest2) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *UploadDeploymentArtifactRequest2) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *UploadDeploymentArtifactRequest2) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


