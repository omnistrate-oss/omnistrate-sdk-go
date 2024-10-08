# ServiceEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Service environment ID | 
**Name** | **string** | Name of the Service Environment | 
**PromoteStatus** | Pointer to **string** | The status of the promotion | [optional] 
**SaasPortalStatus** | Pointer to **string** | The status of the SaaS portal for this environment type | [optional] 
**SaasPortalUrl** | Pointer to **string** | The URL of the SaaS portal for this environment type | [optional] 
**ServicePlans** | [**[]ServicePlan**](ServicePlan.md) | List of service plans | 
**SourceEnvironmentID** | Pointer to **string** | The source environment ID | [optional] 
**SourceEnvironmentName** | Pointer to **string** | The source environment name | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Visibility** | **string** | Visibility of the service environment | 

## Methods

### NewServiceEnvironment

`func NewServiceEnvironment(id string, name string, servicePlans []ServicePlan, visibility string, ) *ServiceEnvironment`

NewServiceEnvironment instantiates a new ServiceEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceEnvironmentWithDefaults

`func NewServiceEnvironmentWithDefaults() *ServiceEnvironment`

NewServiceEnvironmentWithDefaults instantiates a new ServiceEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceEnvironment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceEnvironment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceEnvironment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ServiceEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetPromoteStatus

`func (o *ServiceEnvironment) GetPromoteStatus() string`

GetPromoteStatus returns the PromoteStatus field if non-nil, zero value otherwise.

### GetPromoteStatusOk

`func (o *ServiceEnvironment) GetPromoteStatusOk() (*string, bool)`

GetPromoteStatusOk returns a tuple with the PromoteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoteStatus

`func (o *ServiceEnvironment) SetPromoteStatus(v string)`

SetPromoteStatus sets PromoteStatus field to given value.

### HasPromoteStatus

`func (o *ServiceEnvironment) HasPromoteStatus() bool`

HasPromoteStatus returns a boolean if a field has been set.

### GetSaasPortalStatus

`func (o *ServiceEnvironment) GetSaasPortalStatus() string`

GetSaasPortalStatus returns the SaasPortalStatus field if non-nil, zero value otherwise.

### GetSaasPortalStatusOk

`func (o *ServiceEnvironment) GetSaasPortalStatusOk() (*string, bool)`

GetSaasPortalStatusOk returns a tuple with the SaasPortalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasPortalStatus

`func (o *ServiceEnvironment) SetSaasPortalStatus(v string)`

SetSaasPortalStatus sets SaasPortalStatus field to given value.

### HasSaasPortalStatus

`func (o *ServiceEnvironment) HasSaasPortalStatus() bool`

HasSaasPortalStatus returns a boolean if a field has been set.

### GetSaasPortalUrl

`func (o *ServiceEnvironment) GetSaasPortalUrl() string`

GetSaasPortalUrl returns the SaasPortalUrl field if non-nil, zero value otherwise.

### GetSaasPortalUrlOk

`func (o *ServiceEnvironment) GetSaasPortalUrlOk() (*string, bool)`

GetSaasPortalUrlOk returns a tuple with the SaasPortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasPortalUrl

`func (o *ServiceEnvironment) SetSaasPortalUrl(v string)`

SetSaasPortalUrl sets SaasPortalUrl field to given value.

### HasSaasPortalUrl

`func (o *ServiceEnvironment) HasSaasPortalUrl() bool`

HasSaasPortalUrl returns a boolean if a field has been set.

### GetServicePlans

`func (o *ServiceEnvironment) GetServicePlans() []ServicePlan`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *ServiceEnvironment) GetServicePlansOk() (*[]ServicePlan, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlans

`func (o *ServiceEnvironment) SetServicePlans(v []ServicePlan)`

SetServicePlans sets ServicePlans field to given value.


### GetSourceEnvironmentID

`func (o *ServiceEnvironment) GetSourceEnvironmentID() string`

GetSourceEnvironmentID returns the SourceEnvironmentID field if non-nil, zero value otherwise.

### GetSourceEnvironmentIDOk

`func (o *ServiceEnvironment) GetSourceEnvironmentIDOk() (*string, bool)`

GetSourceEnvironmentIDOk returns a tuple with the SourceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironmentID

`func (o *ServiceEnvironment) SetSourceEnvironmentID(v string)`

SetSourceEnvironmentID sets SourceEnvironmentID field to given value.

### HasSourceEnvironmentID

`func (o *ServiceEnvironment) HasSourceEnvironmentID() bool`

HasSourceEnvironmentID returns a boolean if a field has been set.

### GetSourceEnvironmentName

`func (o *ServiceEnvironment) GetSourceEnvironmentName() string`

GetSourceEnvironmentName returns the SourceEnvironmentName field if non-nil, zero value otherwise.

### GetSourceEnvironmentNameOk

`func (o *ServiceEnvironment) GetSourceEnvironmentNameOk() (*string, bool)`

GetSourceEnvironmentNameOk returns a tuple with the SourceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnvironmentName

`func (o *ServiceEnvironment) SetSourceEnvironmentName(v string)`

SetSourceEnvironmentName sets SourceEnvironmentName field to given value.

### HasSourceEnvironmentName

`func (o *ServiceEnvironment) HasSourceEnvironmentName() bool`

HasSourceEnvironmentName returns a boolean if a field has been set.

### GetType

`func (o *ServiceEnvironment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceEnvironment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceEnvironment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceEnvironment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *ServiceEnvironment) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ServiceEnvironment) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ServiceEnvironment) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


