# CancelHostClusterDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Host Cluster | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCancelHostClusterDeploymentRequest

`func NewCancelHostClusterDeploymentRequest(id string, token string, ) *CancelHostClusterDeploymentRequest`

NewCancelHostClusterDeploymentRequest instantiates a new CancelHostClusterDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelHostClusterDeploymentRequestWithDefaults

`func NewCancelHostClusterDeploymentRequestWithDefaults() *CancelHostClusterDeploymentRequest`

NewCancelHostClusterDeploymentRequestWithDefaults instantiates a new CancelHostClusterDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CancelHostClusterDeploymentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CancelHostClusterDeploymentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CancelHostClusterDeploymentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *CancelHostClusterDeploymentRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CancelHostClusterDeploymentRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CancelHostClusterDeploymentRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


