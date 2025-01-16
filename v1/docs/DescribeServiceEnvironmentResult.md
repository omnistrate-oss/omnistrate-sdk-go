# DescribeServiceEnvironmentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApproveSubscription** | Pointer to **bool** | Auto approve subscription or not | [optional] 
**DeploymentConfigId** | **string** | ID of a Deployment Config | 
**Description** | **string** | A brief description of the service environment | 
**Id** | **string** | ID of a Service Environment | 
**Key** | **string** | Unique Key of the Service Environment | 
**Name** | **string** | Name of the Service Environment | 
**RoleType** | Pointer to **string** | Type of the role | [optional] 
**SaasPortalStatus** | Pointer to **string** | The status of an operation | [optional] 
**SaasPortalUrl** | Pointer to **string** | The URL of the SaaS portal for this environment type | [optional] 
**ServiceAuthPublicKey** | Pointer to **string** | PEM-encoded Public key part of the key used to sign the JWT tokens for the service control plane APIs | [optional] 
**ServiceId** | **string** | ID of a Service | 
**SourceEnvironmentId** | Pointer to **string** | ID of a Service Environment | [optional] 
**Type** | **string** | The type of service environment | 
**Visibility** | **string** | This parameter is used to configure the visibility of the service control-plane APIs | 

## Methods

### NewDescribeServiceEnvironmentResult

`func NewDescribeServiceEnvironmentResult(deploymentConfigId string, description string, id string, key string, name string, serviceId string, type_ string, visibility string, ) *DescribeServiceEnvironmentResult`

NewDescribeServiceEnvironmentResult instantiates a new DescribeServiceEnvironmentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceEnvironmentResultWithDefaults

`func NewDescribeServiceEnvironmentResultWithDefaults() *DescribeServiceEnvironmentResult`

NewDescribeServiceEnvironmentResultWithDefaults instantiates a new DescribeServiceEnvironmentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApproveSubscription

`func (o *DescribeServiceEnvironmentResult) GetAutoApproveSubscription() bool`

GetAutoApproveSubscription returns the AutoApproveSubscription field if non-nil, zero value otherwise.

### GetAutoApproveSubscriptionOk

`func (o *DescribeServiceEnvironmentResult) GetAutoApproveSubscriptionOk() (*bool, bool)`

GetAutoApproveSubscriptionOk returns a tuple with the AutoApproveSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveSubscription

`func (o *DescribeServiceEnvironmentResult) SetAutoApproveSubscription(v bool)`

SetAutoApproveSubscription sets AutoApproveSubscription field to given value.

### HasAutoApproveSubscription

`func (o *DescribeServiceEnvironmentResult) HasAutoApproveSubscription() bool`

HasAutoApproveSubscription returns a boolean if a field has been set.

### GetDeploymentConfigId

`func (o *DescribeServiceEnvironmentResult) GetDeploymentConfigId() string`

GetDeploymentConfigId returns the DeploymentConfigId field if non-nil, zero value otherwise.

### GetDeploymentConfigIdOk

`func (o *DescribeServiceEnvironmentResult) GetDeploymentConfigIdOk() (*string, bool)`

GetDeploymentConfigIdOk returns a tuple with the DeploymentConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentConfigId

`func (o *DescribeServiceEnvironmentResult) SetDeploymentConfigId(v string)`

SetDeploymentConfigId sets DeploymentConfigId field to given value.


### GetDescription

`func (o *DescribeServiceEnvironmentResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeServiceEnvironmentResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeServiceEnvironmentResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeServiceEnvironmentResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeServiceEnvironmentResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeServiceEnvironmentResult) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *DescribeServiceEnvironmentResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DescribeServiceEnvironmentResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DescribeServiceEnvironmentResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *DescribeServiceEnvironmentResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeServiceEnvironmentResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeServiceEnvironmentResult) SetName(v string)`

SetName sets Name field to given value.


### GetRoleType

`func (o *DescribeServiceEnvironmentResult) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *DescribeServiceEnvironmentResult) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *DescribeServiceEnvironmentResult) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *DescribeServiceEnvironmentResult) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetSaasPortalStatus

`func (o *DescribeServiceEnvironmentResult) GetSaasPortalStatus() string`

GetSaasPortalStatus returns the SaasPortalStatus field if non-nil, zero value otherwise.

### GetSaasPortalStatusOk

`func (o *DescribeServiceEnvironmentResult) GetSaasPortalStatusOk() (*string, bool)`

GetSaasPortalStatusOk returns a tuple with the SaasPortalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasPortalStatus

`func (o *DescribeServiceEnvironmentResult) SetSaasPortalStatus(v string)`

SetSaasPortalStatus sets SaasPortalStatus field to given value.

### HasSaasPortalStatus

`func (o *DescribeServiceEnvironmentResult) HasSaasPortalStatus() bool`

HasSaasPortalStatus returns a boolean if a field has been set.

### GetSaasPortalUrl

`func (o *DescribeServiceEnvironmentResult) GetSaasPortalUrl() string`

GetSaasPortalUrl returns the SaasPortalUrl field if non-nil, zero value otherwise.

### GetSaasPortalUrlOk

`func (o *DescribeServiceEnvironmentResult) GetSaasPortalUrlOk() (*string, bool)`

GetSaasPortalUrlOk returns a tuple with the SaasPortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasPortalUrl

`func (o *DescribeServiceEnvironmentResult) SetSaasPortalUrl(v string)`

SetSaasPortalUrl sets SaasPortalUrl field to given value.

### HasSaasPortalUrl

`func (o *DescribeServiceEnvironmentResult) HasSaasPortalUrl() bool`

HasSaasPortalUrl returns a boolean if a field has been set.

### GetServiceAuthPublicKey

`func (o *DescribeServiceEnvironmentResult) GetServiceAuthPublicKey() string`

GetServiceAuthPublicKey returns the ServiceAuthPublicKey field if non-nil, zero value otherwise.

### GetServiceAuthPublicKeyOk

`func (o *DescribeServiceEnvironmentResult) GetServiceAuthPublicKeyOk() (*string, bool)`

GetServiceAuthPublicKeyOk returns a tuple with the ServiceAuthPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAuthPublicKey

`func (o *DescribeServiceEnvironmentResult) SetServiceAuthPublicKey(v string)`

SetServiceAuthPublicKey sets ServiceAuthPublicKey field to given value.

### HasServiceAuthPublicKey

`func (o *DescribeServiceEnvironmentResult) HasServiceAuthPublicKey() bool`

HasServiceAuthPublicKey returns a boolean if a field has been set.

### GetServiceId

`func (o *DescribeServiceEnvironmentResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeServiceEnvironmentResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeServiceEnvironmentResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSourceEnvironmentId

`func (o *DescribeServiceEnvironmentResult) GetSourceEnvironmentId() string`

GetSourceEnvironmentId returns the SourceEnvironmentId field if non-nil, zero value otherwise.

### GetSourceEnvironmentIdOk

`func (o *DescribeServiceEnvironmentResult) GetSourceEnvironmentIdOk() (*string, bool)`

GetSourceEnvironmentIdOk returns a tuple with the SourceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironmentId

`func (o *DescribeServiceEnvironmentResult) SetSourceEnvironmentId(v string)`

SetSourceEnvironmentId sets SourceEnvironmentId field to given value.

### HasSourceEnvironmentId

`func (o *DescribeServiceEnvironmentResult) HasSourceEnvironmentId() bool`

HasSourceEnvironmentId returns a boolean if a field has been set.

### GetType

`func (o *DescribeServiceEnvironmentResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DescribeServiceEnvironmentResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DescribeServiceEnvironmentResult) SetType(v string)`

SetType sets Type field to given value.


### GetVisibility

`func (o *DescribeServiceEnvironmentResult) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *DescribeServiceEnvironmentResult) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *DescribeServiceEnvironmentResult) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


