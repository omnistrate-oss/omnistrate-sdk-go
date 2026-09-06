# ArtifactRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalPath** | **string** | The canonical local artifact root that must be supplied. Canonical form removes a leading &#39;./&#39; and represents the source root as &#39;.&#39;. | 
**Uses** | [**[]ValidationArtifactUse**](ValidationArtifactUse.md) | Every place in the candidate that consumes this content. Always nonempty. Content used by several resources, providers or platforms is reported once with all of its uses aggregated. | 

## Methods

### NewArtifactRequirement

`func NewArtifactRequirement(logicalPath string, uses []ValidationArtifactUse, ) *ArtifactRequirement`

NewArtifactRequirement instantiates a new ArtifactRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactRequirementWithDefaults

`func NewArtifactRequirementWithDefaults() *ArtifactRequirement`

NewArtifactRequirementWithDefaults instantiates a new ArtifactRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalPath

`func (o *ArtifactRequirement) GetLogicalPath() string`

GetLogicalPath returns the LogicalPath field if non-nil, zero value otherwise.

### GetLogicalPathOk

`func (o *ArtifactRequirement) GetLogicalPathOk() (*string, bool)`

GetLogicalPathOk returns a tuple with the LogicalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalPath

`func (o *ArtifactRequirement) SetLogicalPath(v string)`

SetLogicalPath sets LogicalPath field to given value.


### GetUses

`func (o *ArtifactRequirement) GetUses() []ValidationArtifactUse`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *ArtifactRequirement) GetUsesOk() (*[]ValidationArtifactUse, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *ArtifactRequirement) SetUses(v []ValidationArtifactUse)`

SetUses sets Uses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


