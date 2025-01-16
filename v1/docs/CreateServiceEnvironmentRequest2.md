# CreateServiceEnvironmentRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**DeploymentConfigId** | **string** | The deployment configuration ID | 
**Description** | **string** | A brief description of the service environment | 
**Name** | **string** | Name of the Service Environment | 
**ServiceAuthPublicKey** | Pointer to **string** | PEM-encoded Public key part of the key used to sign the JWT tokens for the service control plane APIs | [optional] 
**SourceEnvironmentId** | Pointer to **string** | The ID of the service environment to use for promoting changes to this environment | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** | This parameter is used to configure the visibility of the service control-plane APIs | [optional] 

## Methods

### NewCreateServiceEnvironmentRequest2

`func NewCreateServiceEnvironmentRequest2(deploymentConfigId string, description string, name string, ) *CreateServiceEnvironmentRequest2`

NewCreateServiceEnvironmentRequest2 instantiates a new CreateServiceEnvironmentRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceEnvironmentRequest2WithDefaults

`func NewCreateServiceEnvironmentRequest2WithDefaults() *CreateServiceEnvironmentRequest2`

NewCreateServiceEnvironmentRequest2WithDefaults instantiates a new CreateServiceEnvironmentRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *CreateServiceEnvironmentRequest2) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *CreateServiceEnvironmentRequest2) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *CreateServiceEnvironmentRequest2) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *CreateServiceEnvironmentRequest2) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetDeploymentConfigId

`func (o *CreateServiceEnvironmentRequest2) GetDeploymentConfigId() string`

GetDeploymentConfigId returns the DeploymentConfigId field if non-nil, zero value otherwise.

### GetDeploymentConfigIdOk

`func (o *CreateServiceEnvironmentRequest2) GetDeploymentConfigIdOk() (*string, bool)`

GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfigId

`func (o *CreateServiceEnvironmentRequest2) SetDeploymentConfigId(v string)`

SetDeploymentConfigId sets DeploymentConfigId field to given value.


### GetDescription

`func (o *CreateServiceEnvironmentRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceEnvironmentRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceEnvironmentRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateServiceEnvironmentRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceEnvironmentRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceEnvironmentRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetServiceAuthPublicKey

`func (o *CreateServiceEnvironmentRequest2) GetServiceAuthPublicKey() string`

GetServiceAuthPublicKey returns the ServiceAuthPublicKey field if non-nil, zero value otherwise.

### GetServiceAuthPublicKeyOk

`func (o *CreateServiceEnvironmentRequest2) GetServiceAuthPublicKeyOk() (*string, bool)`

GetServiceAuthPublicKeyOk returns a tuple with the ServiceAuthPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAuthPublicKey

`func (o *CreateServiceEnvironmentRequest2) SetServiceAuthPublicKey(v string)`

SetServiceAuthPublicKey sets ServiceAuthPublicKey field to given value.

### HasServiceAuthPublicKey

`func (o *CreateServiceEnvironmentRequest2) HasServiceAuthPublicKey() bool`

HasServiceAuthPublicKey returns a boolean if a field has been set.

### GetSourceEnvironmentId

`func (o *CreateServiceEnvironmentRequest2) GetSourceEnvironmentId() string`

GetSourceEnvironmentId returns the SourceEnvironmentId field if non-nil, zero value otherwise.

### GetSourceEnvironmentIdOk

`func (o *CreateServiceEnvironmentRequest2) GetSourceEnvironmentIdOk() (*string, bool)`

GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironmentId

`func (o *CreateServiceEnvironmentRequest2) SetSourceEnvironmentId(v string)`

SetSourceEnvironmentId sets SourceEnvironmentId field to given value.

### HasSourceEnvironmentId

`func (o *CreateServiceEnvironmentRequest2) HasSourceEnvironmentId() bool`

HasSourceEnvironmentId returns a boolean if a field has been set.

### GetType

`func (o *CreateServiceEnvironmentRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateServiceEnvironmentRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateServiceEnvironmentRequest2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateServiceEnvironmentRequest2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateServiceEnvironmentRequest2) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateServiceEnvironmentRequest2) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateServiceEnvironmentRequest2) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateServiceEnvironmentRequest2) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


