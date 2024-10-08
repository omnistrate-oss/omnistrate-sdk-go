# UpdateServiceEnvironmentRequestBody

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

### NewUpdateServiceEnvironmentRequestBody

`func NewUpdateServiceEnvironmentRequestBody() *UpdateServiceEnvironmentRequestBody`

NewUpdateServiceEnvironmentRequestBody instantiates a new UpdateServiceEnvironmentRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceEnvironmentRequestBodyWithDefaults

`func NewUpdateServiceEnvironmentRequestBodyWithDefaults() *UpdateServiceEnvironmentRequestBody`

NewUpdateServiceEnvironmentRequestBodyWithDefaults instantiates a new UpdateServiceEnvironmentRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *UpdateServiceEnvironmentRequestBody) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *UpdateServiceEnvironmentRequestBody) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *UpdateServiceEnvironmentRequestBody) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *UpdateServiceEnvironmentRequestBody) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetDeploymentConfigId

`func (o *UpdateServiceEnvironmentRequestBody) GetDeploymentConfigId() string`

GetDeploymentConfigId returns the DeploymentConfigId field if non-nil, zero value otherwise.

### GetDeploymentConfigIdOk

`func (o *UpdateServiceEnvironmentRequestBody) GetDeploymentConfigIdOk() (*string, bool)`

GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfigId

`func (o *UpdateServiceEnvironmentRequestBody) SetDeploymentConfigId(v string)`

SetDeploymentConfigId sets DeploymentConfigId field to given value.

### HasDeploymentConfigId

`func (o *UpdateServiceEnvironmentRequestBody) HasDeploymentConfigId() bool`

HasDeploymentConfigId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateServiceEnvironmentRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServiceEnvironmentRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServiceEnvironmentRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServiceEnvironmentRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateServiceEnvironmentRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServiceEnvironmentRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServiceEnvironmentRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServiceEnvironmentRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceAuthPublicKey

`func (o *UpdateServiceEnvironmentRequestBody) GetServiceAuthPublicKey() string`

GetServiceAuthPublicKey returns the ServiceAuthPublicKey field if non-nil, zero value otherwise.

### GetServiceAuthPublicKeyOk

`func (o *UpdateServiceEnvironmentRequestBody) GetServiceAuthPublicKeyOk() (*string, bool)`

GetServiceAuthPublicKeyOk returns a tuple with the ServiceAuthPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAuthPublicKey

`func (o *UpdateServiceEnvironmentRequestBody) SetServiceAuthPublicKey(v string)`

SetServiceAuthPublicKey sets ServiceAuthPublicKey field to given value.

### HasServiceAuthPublicKey

`func (o *UpdateServiceEnvironmentRequestBody) HasServiceAuthPublicKey() bool`

HasServiceAuthPublicKey returns a boolean if a field has been set.

### GetSourceEnvironmentId

`func (o *UpdateServiceEnvironmentRequestBody) GetSourceEnvironmentId() string`

GetSourceEnvironmentId returns the SourceEnvironmentId field if non-nil, zero value otherwise.

### GetSourceEnvironmentIdOk

`func (o *UpdateServiceEnvironmentRequestBody) GetSourceEnvironmentIdOk() (*string, bool)`

GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironmentId

`func (o *UpdateServiceEnvironmentRequestBody) SetSourceEnvironmentId(v string)`

SetSourceEnvironmentId sets SourceEnvironmentId field to given value.

### HasSourceEnvironmentId

`func (o *UpdateServiceEnvironmentRequestBody) HasSourceEnvironmentId() bool`

HasSourceEnvironmentId returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateServiceEnvironmentRequestBody) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateServiceEnvironmentRequestBody) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateServiceEnvironmentRequestBody) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateServiceEnvironmentRequestBody) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


