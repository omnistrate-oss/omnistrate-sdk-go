# CreateServiceEnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**DeploymentConfigId** | **string** | ID of a Deployment Config | 
**Description** | **string** | A brief description of the service environment | 
**Name** | **string** | Name of the Service Environment | 
**ServiceAuthPublicKey** | Pointer to **string** | PEM-encoded Public key part of the key used to sign the JWT tokens for the service control plane APIs | [optional] 
**ServiceId** | **string** | ID of a Service | 
**SourceEnvironmentId** | Pointer to **string** | ID of a Service Environment | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Type** | Pointer to **string** | The type of service environment | [optional] 
**Visibility** | Pointer to **string** | This parameter is used to configure the visibility of the service control-plane APIs | [optional] 

## Methods

### NewCreateServiceEnvironmentRequest

`func NewCreateServiceEnvironmentRequest(deploymentConfigId string, description string, name string, serviceId string, token string, ) *CreateServiceEnvironmentRequest`

NewCreateServiceEnvironmentRequest instantiates a new CreateServiceEnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceEnvironmentRequestWithDefaults

`func NewCreateServiceEnvironmentRequestWithDefaults() *CreateServiceEnvironmentRequest`

NewCreateServiceEnvironmentRequestWithDefaults instantiates a new CreateServiceEnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *CreateServiceEnvironmentRequest) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *CreateServiceEnvironmentRequest) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *CreateServiceEnvironmentRequest) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *CreateServiceEnvironmentRequest) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetDeploymentConfigId

`func (o *CreateServiceEnvironmentRequest) GetDeploymentConfigId() string`

GetDeploymentConfigId returns the DeploymentConfigId field if non-nil, zero value otherwise.

### GetDeploymentConfigIdOk

`func (o *CreateServiceEnvironmentRequest) GetDeploymentConfigIdOk() (*string, bool)`

GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfigId

`func (o *CreateServiceEnvironmentRequest) SetDeploymentConfigId(v string)`

SetDeploymentConfigId sets DeploymentConfigId field to given value.


### GetDescription

`func (o *CreateServiceEnvironmentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceEnvironmentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceEnvironmentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateServiceEnvironmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceEnvironmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceEnvironmentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServiceAuthPublicKey

`func (o *CreateServiceEnvironmentRequest) GetServiceAuthPublicKey() string`

GetServiceAuthPublicKey returns the ServiceAuthPublicKey field if non-nil, zero value otherwise.

### GetServiceAuthPublicKeyOk

`func (o *CreateServiceEnvironmentRequest) GetServiceAuthPublicKeyOk() (*string, bool)`

GetServiceAuthPublicKeyOk returns a tuple with the ServiceAuthPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAuthPublicKey

`func (o *CreateServiceEnvironmentRequest) SetServiceAuthPublicKey(v string)`

SetServiceAuthPublicKey sets ServiceAuthPublicKey field to given value.

### HasServiceAuthPublicKey

`func (o *CreateServiceEnvironmentRequest) HasServiceAuthPublicKey() bool`

HasServiceAuthPublicKey returns a boolean if a field has been set.

### GetServiceId

`func (o *CreateServiceEnvironmentRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateServiceEnvironmentRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateServiceEnvironmentRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSourceEnvironmentId

`func (o *CreateServiceEnvironmentRequest) GetSourceEnvironmentId() string`

GetSourceEnvironmentId returns the SourceEnvironmentId field if non-nil, zero value otherwise.

### GetSourceEnvironmentIdOk

`func (o *CreateServiceEnvironmentRequest) GetSourceEnvironmentIdOk() (*string, bool)`

GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironmentId

`func (o *CreateServiceEnvironmentRequest) SetSourceEnvironmentId(v string)`

SetSourceEnvironmentId sets SourceEnvironmentId field to given value.

### HasSourceEnvironmentId

`func (o *CreateServiceEnvironmentRequest) HasSourceEnvironmentId() bool`

HasSourceEnvironmentId returns a boolean if a field has been set.

### GetToken

`func (o *CreateServiceEnvironmentRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateServiceEnvironmentRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateServiceEnvironmentRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *CreateServiceEnvironmentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateServiceEnvironmentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateServiceEnvironmentRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateServiceEnvironmentRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateServiceEnvironmentRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateServiceEnvironmentRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateServiceEnvironmentRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateServiceEnvironmentRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


