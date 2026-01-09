# UploadDeploymentArtifactRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | **[]*os.File** | The deployment artifact file content | 
**ArtifactPath** | **string** | The path to the deployment artifact | 
**ProductTierName** | **string** | The name of the product tier | 
**ServiceName** | **string** | The name of the service | 

## Methods

### NewUploadDeploymentArtifactRequest2

`func NewUploadDeploymentArtifactRequest2(artifact []*os.File, artifactPath string, productTierName string, serviceName string, ) *UploadDeploymentArtifactRequest2`

NewUploadDeploymentArtifactRequest2 instantiates a new UploadDeploymentArtifactRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadDeploymentArtifactRequest2WithDefaults

`func NewUploadDeploymentArtifactRequest2WithDefaults() *UploadDeploymentArtifactRequest2`

NewUploadDeploymentArtifactRequest2WithDefaults instantiates a new UploadDeploymentArtifactRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *UploadDeploymentArtifactRequest2) GetArtifact() []*os.File`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *UploadDeploymentArtifactRequest2) GetArtifactOk() (*[]*os.File, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *UploadDeploymentArtifactRequest2) SetArtifact(v []*os.File)`

SetArtifact sets Artifact field to given value.


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


