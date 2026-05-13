# DescribeDeploymentArtifactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** | The hash of the deployment artifact to be described. It is used as an additional filter criteria to ensure the id and the hash match for the same deployment artifact. | [optional] 
**Id** | **string** | ID of a Deployment Artifact | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeDeploymentArtifactRequest

`func NewDescribeDeploymentArtifactRequest(id string, token string, ) *DescribeDeploymentArtifactRequest`

NewDescribeDeploymentArtifactRequest instantiates a new DescribeDeploymentArtifactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDeploymentArtifactRequestWithDefaults

`func NewDescribeDeploymentArtifactRequestWithDefaults() *DescribeDeploymentArtifactRequest`

NewDescribeDeploymentArtifactRequestWithDefaults instantiates a new DescribeDeploymentArtifactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *DescribeDeploymentArtifactRequest) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *DescribeDeploymentArtifactRequest) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *DescribeDeploymentArtifactRequest) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *DescribeDeploymentArtifactRequest) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *DescribeDeploymentArtifactRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeDeploymentArtifactRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeDeploymentArtifactRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *DescribeDeploymentArtifactRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeDeploymentArtifactRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeDeploymentArtifactRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


