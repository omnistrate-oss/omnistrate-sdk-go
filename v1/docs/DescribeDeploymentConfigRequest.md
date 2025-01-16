# DescribeDeploymentConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The deployment configuration ID | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeDeploymentConfigRequest

`func NewDescribeDeploymentConfigRequest(id string, token string, ) *DescribeDeploymentConfigRequest`

NewDescribeDeploymentConfigRequest instantiates a new DescribeDeploymentConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDeploymentConfigRequestWithDefaults

`func NewDescribeDeploymentConfigRequestWithDefaults() *DescribeDeploymentConfigRequest`

NewDescribeDeploymentConfigRequestWithDefaults instantiates a new DescribeDeploymentConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DescribeDeploymentConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeDeploymentConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeDeploymentConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *DescribeDeploymentConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeDeploymentConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeDeploymentConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


