# DescribeServiceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time the service was created | 
**Description** | **string** | A brief description of the service | 
**Id** | **string** | The service ID to operate on | 
**Key** | **string** | Unique key of the service | 
**Name** | **string** | Name of the Service | 
**RoleType** | Pointer to **string** | The role type of the caller user | [optional] 
**ServiceEnvironments** | [**[]ServiceEnvironment**](ServiceEnvironment.md) | List of service environments | 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 
**ServiceProviderID** | **string** | The ID of the service provider owning the service | 
**ServiceProviderName** | **string** | The name of the service provider | 

## Methods

### NewDescribeServiceResult

`func NewDescribeServiceResult(createdAt string, description string, id string, key string, name string, serviceEnvironments []ServiceEnvironment, serviceProviderID string, serviceProviderName string, ) *DescribeServiceResult`

NewDescribeServiceResult instantiates a new DescribeServiceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceResultWithDefaults

`func NewDescribeServiceResultWithDefaults() *DescribeServiceResult`

NewDescribeServiceResultWithDefaults instantiates a new DescribeServiceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DescribeServiceResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DescribeServiceResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DescribeServiceResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *DescribeServiceResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeServiceResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeServiceResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeServiceResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeServiceResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeServiceResult) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *DescribeServiceResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DescribeServiceResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DescribeServiceResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *DescribeServiceResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeServiceResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeServiceResult) SetName(v string)`

SetName sets Name field to given value.


### GetRoleType

`func (o *DescribeServiceResult) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *DescribeServiceResult) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *DescribeServiceResult) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *DescribeServiceResult) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetServiceEnvironments

`func (o *DescribeServiceResult) GetServiceEnvironments() []ServiceEnvironment`

GetServiceEnvironments returns the ServiceEnvironments field if non-nil, zero value otherwise.

### GetServiceEnvironmentsOk

`func (o *DescribeServiceResult) GetServiceEnvironmentsOk() (*[]ServiceEnvironment, bool)`

GetServiceEnvironmentsOk returns a tuple with the ServiceEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironments

`func (o *DescribeServiceResult) SetServiceEnvironments(v []ServiceEnvironment)`

SetServiceEnvironments sets ServiceEnvironments field to given value.


### GetServiceLogoURL

`func (o *DescribeServiceResult) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *DescribeServiceResult) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *DescribeServiceResult) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *DescribeServiceResult) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.

### GetServiceProviderID

`func (o *DescribeServiceResult) GetServiceProviderID() string`

GetServiceProviderID returns the ServiceProviderID field if non-nil, zero value otherwise.

### GetServiceProviderIDOk

`func (o *DescribeServiceResult) GetServiceProviderIDOk() (*string, bool)`

GetServiceProviderIDOk returns a tuple with the ServiceProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderID

`func (o *DescribeServiceResult) SetServiceProviderID(v string)`

SetServiceProviderID sets ServiceProviderID field to given value.


### GetServiceProviderName

`func (o *DescribeServiceResult) GetServiceProviderName() string`

GetServiceProviderName returns the ServiceProviderName field if non-nil, zero value otherwise.

### GetServiceProviderNameOk

`func (o *DescribeServiceResult) GetServiceProviderNameOk() (*string, bool)`

GetServiceProviderNameOk returns a tuple with the ServiceProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderName

`func (o *DescribeServiceResult) SetServiceProviderName(v string)`

SetServiceProviderName sets ServiceProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


