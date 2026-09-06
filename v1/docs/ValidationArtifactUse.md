# ValidationArtifactUse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind of configuration consuming the content, for example terraform, helm, kustomize or operator CRD content. | 
**OnPremPlatform** | Pointer to **string** | OnPrem model platform | [optional] 
**Path** | **string** | JSON pointer into the submitted specification identifying where this use is declared | 
**Provider** | Pointer to **string** | Name of the Infra Provider | [optional] 
**ResourceKey** | **string** | The canonical resource key that consumes the content. This is a specification key, never a persisted resource ID. | 

## Methods

### NewValidationArtifactUse

`func NewValidationArtifactUse(kind string, path string, resourceKey string, ) *ValidationArtifactUse`

NewValidationArtifactUse instantiates a new ValidationArtifactUse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationArtifactUseWithDefaults

`func NewValidationArtifactUseWithDefaults() *ValidationArtifactUse`

NewValidationArtifactUseWithDefaults instantiates a new ValidationArtifactUse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ValidationArtifactUse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ValidationArtifactUse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ValidationArtifactUse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetOnPremPlatform

`func (o *ValidationArtifactUse) GetOnPremPlatform() string`

GetOnPremPlatform returns the OnPremPlatform field if non-nil, zero value otherwise.

### GetOnPremPlatformOk

`func (o *ValidationArtifactUse) GetOnPremPlatformOk() (*string, bool)`

GetOnPremPlatformOk returns a tuple with the OnPremPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremPlatform

`func (o *ValidationArtifactUse) SetOnPremPlatform(v string)`

SetOnPremPlatform sets OnPremPlatform field to given value.

### HasOnPremPlatform

`func (o *ValidationArtifactUse) HasOnPremPlatform() bool`

HasOnPremPlatform returns a boolean if a field has been set.

### GetPath

`func (o *ValidationArtifactUse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ValidationArtifactUse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ValidationArtifactUse) SetPath(v string)`

SetPath sets Path field to given value.


### GetProvider

`func (o *ValidationArtifactUse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ValidationArtifactUse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ValidationArtifactUse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ValidationArtifactUse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetResourceKey

`func (o *ValidationArtifactUse) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *ValidationArtifactUse) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *ValidationArtifactUse) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


