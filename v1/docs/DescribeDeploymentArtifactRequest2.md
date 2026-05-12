# DescribeDeploymentArtifactRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** | The hash of the deployment artifact to be described. It is used as an additional filter criteria to ensure the id and the hash match for the same deployment artifact. | [optional] 

## Methods

### NewDescribeDeploymentArtifactRequest2

`func NewDescribeDeploymentArtifactRequest2() *DescribeDeploymentArtifactRequest2`

NewDescribeDeploymentArtifactRequest2 instantiates a new DescribeDeploymentArtifactRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDeploymentArtifactRequest2WithDefaults

`func NewDescribeDeploymentArtifactRequest2WithDefaults() *DescribeDeploymentArtifactRequest2`

NewDescribeDeploymentArtifactRequest2WithDefaults instantiates a new DescribeDeploymentArtifactRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *DescribeDeploymentArtifactRequest2) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *DescribeDeploymentArtifactRequest2) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *DescribeDeploymentArtifactRequest2) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *DescribeDeploymentArtifactRequest2) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


