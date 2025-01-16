# UpdateServiceEnvironmentRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**DeploymentConfigId** | Pointer to **string** | The deployment configuration ID | [optional] 
**Description** | Pointer to **string** | A brief description of the service environment | [optional] 
**Name** | Pointer to **string** | Name of the Service Environment | [optional] 
**ServiceAuthPublicKey** | Pointer to **string** | PEM-encoded Public key part of the key used to sign the JWT tokens for the service control plane APIs | [optional] 
**SourceEnvironmentId** | Pointer to **string** | The ID of the service environment to use for promoting changes to this environment | [optional] 
**Visibility** | Pointer to **string** | This parameter is used to configure the visibility of the service control-plane APIs | [optional] 

## Methods

### NewUpdateServiceEnvironmentRequest2

`func NewUpdateServiceEnvironmentRequest2() *UpdateServiceEnvironmentRequest2`

NewUpdateServiceEnvironmentRequest2 instantiates a new UpdateServiceEnvironmentRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceEnvironmentRequest2WithDefaults

`func NewUpdateServiceEnvironmentRequest2WithDefaults() *UpdateServiceEnvironmentRequest2`

NewUpdateServiceEnvironmentRequest2WithDefaults instantiates a new UpdateServiceEnvironmentRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *UpdateServiceEnvironmentRequest2) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *UpdateServiceEnvironmentRequest2) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *UpdateServiceEnvironmentRequest2) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *UpdateServiceEnvironmentRequest2) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetDeploymentConfigId

`func (o *UpdateServiceEnvironmentRequest2) GetDeploymentConfigId() string`

GetDeploymentConfigId returns the DeploymentConfigId field if non-nil, zero value otherwise.

### GetDeploymentConfigIdOk

`func (o *UpdateServiceEnvironmentRequest2) GetDeploymentConfigIdOk() (*string, bool)`

GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfigId

`func (o *UpdateServiceEnvironmentRequest2) SetDeploymentConfigId(v string)`

SetDeploymentConfigId sets DeploymentConfigId field to given value.

### HasDeploymentConfigId

`func (o *UpdateServiceEnvironmentRequest2) HasDeploymentConfigId() bool`

HasDeploymentConfigId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateServiceEnvironmentRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServiceEnvironmentRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServiceEnvironmentRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServiceEnvironmentRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateServiceEnvironmentRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServiceEnvironmentRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServiceEnvironmentRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServiceEnvironmentRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceAuthPublicKey

`func (o *UpdateServiceEnvironmentRequest2) GetServiceAuthPublicKey() string`

GetServiceAuthPublicKey returns the ServiceAuthPublicKey field if non-nil, zero value otherwise.

### GetServiceAuthPublicKeyOk

`func (o *UpdateServiceEnvironmentRequest2) GetServiceAuthPublicKeyOk() (*string, bool)`

GetServiceAuthPublicKeyOk returns a tuple with the ServiceAuthPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAuthPublicKey

`func (o *UpdateServiceEnvironmentRequest2) SetServiceAuthPublicKey(v string)`

SetServiceAuthPublicKey sets ServiceAuthPublicKey field to given value.

### HasServiceAuthPublicKey

`func (o *UpdateServiceEnvironmentRequest2) HasServiceAuthPublicKey() bool`

HasServiceAuthPublicKey returns a boolean if a field has been set.

### GetSourceEnvironmentId

`func (o *UpdateServiceEnvironmentRequest2) GetSourceEnvironmentId() string`

GetSourceEnvironmentId returns the SourceEnvironmentId field if non-nil, zero value otherwise.

### GetSourceEnvironmentIdOk

`func (o *UpdateServiceEnvironmentRequest2) GetSourceEnvironmentIdOk() (*string, bool)`

GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironmentId

`func (o *UpdateServiceEnvironmentRequest2) SetSourceEnvironmentId(v string)`

SetSourceEnvironmentId sets SourceEnvironmentId field to given value.

### HasSourceEnvironmentId

`func (o *UpdateServiceEnvironmentRequest2) HasSourceEnvironmentId() bool`

HasSourceEnvironmentId returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateServiceEnvironmentRequest2) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateServiceEnvironmentRequest2) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateServiceEnvironmentRequest2) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateServiceEnvironmentRequest2) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


