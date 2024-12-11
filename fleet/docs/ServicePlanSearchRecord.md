# ServicePlanSearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentType** | **string** | The Service Model type of the service plan. | 
**Description** | **string** | The Product Tier description of the service plan. | 
**Id** | **string** | The Product Tier ID of the service plan. | 
**Name** | **string** | The Product Tier name of the service plan. | 
**ReleasedAt** | Pointer to **string** | The timestamp when the service plan was released. | [optional] 
**ServiceEnvironmentId** | **string** | The service environment ID of the service plan. | 
**ServiceEnvironmentName** | **string** | The service environment name of the service plan. | 
**ServiceEnvironmentType** | Pointer to **string** | The service environment type of the service plan. | [optional] 
**ServiceId** | **string** | The service ID of the service plan. | 
**ServiceName** | **string** | The service name of the service plan. | 
**TenancyType** | **string** | The Product Tier type of the service plan. | 
**Version** | **string** | The Product Tier version of the service plan. | 
**VersionName** | Pointer to **string** | The Product Tier version name of the service plan. | [optional] 
**VersionSetStatus** | **string** | The Product Tier version set status of the service plan. | 

## Methods

### NewServicePlanSearchRecord

`func NewServicePlanSearchRecord(deploymentType string, description string, id string, name string, serviceEnvironmentId string, serviceEnvironmentName string, serviceId string, serviceName string, tenancyType string, version string, versionSetStatus string, ) *ServicePlanSearchRecord`

NewServicePlanSearchRecord instantiates a new ServicePlanSearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanSearchRecordWithDefaults

`func NewServicePlanSearchRecordWithDefaults() *ServicePlanSearchRecord`

NewServicePlanSearchRecordWithDefaults instantiates a new ServicePlanSearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentType

`func (o *ServicePlanSearchRecord) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *ServicePlanSearchRecord) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *ServicePlanSearchRecord) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.


### GetDescription

`func (o *ServicePlanSearchRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePlanSearchRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePlanSearchRecord) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *ServicePlanSearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicePlanSearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicePlanSearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ServicePlanSearchRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePlanSearchRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePlanSearchRecord) SetName(v string)`

SetName sets Name field to given value.


### GetReleasedAt

`func (o *ServicePlanSearchRecord) GetReleasedAt() string`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *ServicePlanSearchRecord) GetReleasedAtOk() (*string, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *ServicePlanSearchRecord) SetReleasedAt(v string)`

SetReleasedAt sets ReleasedAt field to given value.

### HasReleasedAt

`func (o *ServicePlanSearchRecord) HasReleasedAt() bool`

HasReleasedAt returns a boolean if a field has been set.

### GetServiceEnvironmentId

`func (o *ServicePlanSearchRecord) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *ServicePlanSearchRecord) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *ServicePlanSearchRecord) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceEnvironmentName

`func (o *ServicePlanSearchRecord) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *ServicePlanSearchRecord) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *ServicePlanSearchRecord) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceEnvironmentType

`func (o *ServicePlanSearchRecord) GetServiceEnvironmentType() string`

GetServiceEnvironmentType returns the ServiceEnvironmentType field if non-nil, zero value otherwise.

### GetServiceEnvironmentTypeOk

`func (o *ServicePlanSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool)`

GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentType

`func (o *ServicePlanSearchRecord) SetServiceEnvironmentType(v string)`

SetServiceEnvironmentType sets ServiceEnvironmentType field to given value.

### HasServiceEnvironmentType

`func (o *ServicePlanSearchRecord) HasServiceEnvironmentType() bool`

HasServiceEnvironmentType returns a boolean if a field has been set.

### GetServiceId

`func (o *ServicePlanSearchRecord) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServicePlanSearchRecord) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServicePlanSearchRecord) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceName

`func (o *ServicePlanSearchRecord) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServicePlanSearchRecord) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServicePlanSearchRecord) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetTenancyType

`func (o *ServicePlanSearchRecord) GetTenancyType() string`

GetTenancyType returns the TenancyType field if non-nil, zero value otherwise.

### GetTenancyTypeOk

`func (o *ServicePlanSearchRecord) GetTenancyTypeOk() (*string, bool)`

GetTenancyTypeOk returns a tuple with the TenancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancyType

`func (o *ServicePlanSearchRecord) SetTenancyType(v string)`

SetTenancyType sets TenancyType field to given value.


### GetVersion

`func (o *ServicePlanSearchRecord) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServicePlanSearchRecord) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServicePlanSearchRecord) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersionName

`func (o *ServicePlanSearchRecord) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *ServicePlanSearchRecord) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *ServicePlanSearchRecord) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *ServicePlanSearchRecord) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetVersionSetStatus

`func (o *ServicePlanSearchRecord) GetVersionSetStatus() string`

GetVersionSetStatus returns the VersionSetStatus field if non-nil, zero value otherwise.

### GetVersionSetStatusOk

`func (o *ServicePlanSearchRecord) GetVersionSetStatusOk() (*string, bool)`

GetVersionSetStatusOk returns a tuple with the VersionSetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSetStatus

`func (o *ServicePlanSearchRecord) SetVersionSetStatus(v string)`

SetVersionSetStatus sets VersionSetStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


