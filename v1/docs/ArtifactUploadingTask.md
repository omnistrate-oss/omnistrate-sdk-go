# ArtifactUploadingTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigID** | Pointer to **string** | The account config ID associated with the artifact uploading task. This is defined by the service plan spec and is used by the caller to identify which account config this artifact is associated with. | [optional] 
**ArtifactPath** | Pointer to **string** | The path of the artifact to be uploaded. This is defined by the service plan spec and is used by the caller to identify which artifact to upload for this task. | [optional] 
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**ProductTierName** | Pointer to **string** | The name of the product tier. This is defined by the service plan spec and is used by the caller to identify which product tier this artifact is associated with. | [optional] 
**ServiceName** | Pointer to **string** | The name of the service. This is defined by the service plan spec and is used by the caller to identify which service this artifact is associated with. | [optional] 

## Methods

### NewArtifactUploadingTask

`func NewArtifactUploadingTask() *ArtifactUploadingTask`

NewArtifactUploadingTask instantiates a new ArtifactUploadingTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactUploadingTaskWithDefaults

`func NewArtifactUploadingTaskWithDefaults() *ArtifactUploadingTask`

NewArtifactUploadingTaskWithDefaults instantiates a new ArtifactUploadingTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigID

`func (o *ArtifactUploadingTask) GetAccountConfigID() string`

GetAccountConfigID returns the AccountConfigID field if non-nil, zero value otherwise.

### GetAccountConfigIDOk

`func (o *ArtifactUploadingTask) GetAccountConfigIDOk() (*string, bool)`

GetAccountConfigIDOk returns a tuple with the AccountConfigID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigID

`func (o *ArtifactUploadingTask) SetAccountConfigID(v string)`

SetAccountConfigID sets AccountConfigID field to given value.

### HasAccountConfigID

`func (o *ArtifactUploadingTask) HasAccountConfigID() bool`

HasAccountConfigID returns a boolean if a field has been set.

### GetArtifactPath

`func (o *ArtifactUploadingTask) GetArtifactPath() string`

GetArtifactPath returns the ArtifactPath field if non-nil, zero value otherwise.

### GetArtifactPathOk

`func (o *ArtifactUploadingTask) GetArtifactPathOk() (*string, bool)`

GetArtifactPathOk returns a tuple with the ArtifactPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactPath

`func (o *ArtifactUploadingTask) SetArtifactPath(v string)`

SetArtifactPath sets ArtifactPath field to given value.

### HasArtifactPath

`func (o *ArtifactUploadingTask) HasArtifactPath() bool`

HasArtifactPath returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *ArtifactUploadingTask) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *ArtifactUploadingTask) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *ArtifactUploadingTask) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *ArtifactUploadingTask) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetProductTierName

`func (o *ArtifactUploadingTask) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *ArtifactUploadingTask) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *ArtifactUploadingTask) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.

### HasProductTierName

`func (o *ArtifactUploadingTask) HasProductTierName() bool`

HasProductTierName returns a boolean if a field has been set.

### GetServiceName

`func (o *ArtifactUploadingTask) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ArtifactUploadingTask) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ArtifactUploadingTask) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ArtifactUploadingTask) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


