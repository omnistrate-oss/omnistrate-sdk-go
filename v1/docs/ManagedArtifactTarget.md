# ManagedArtifactTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** | Name of the Infra Provider | [optional] 
**Id** | **string** | Tenant-scoped identifier used to filter synchronization results. | 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedArtifactTarget

`func NewManagedArtifactTarget(id string, ) *ManagedArtifactTarget`

NewManagedArtifactTarget instantiates a new ManagedArtifactTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedArtifactTargetWithDefaults

`func NewManagedArtifactTargetWithDefaults() *ManagedArtifactTarget`

NewManagedArtifactTargetWithDefaults instantiates a new ManagedArtifactTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ManagedArtifactTarget) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ManagedArtifactTarget) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ManagedArtifactTarget) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ManagedArtifactTarget) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCloudProvider

`func (o *ManagedArtifactTarget) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ManagedArtifactTarget) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ManagedArtifactTarget) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ManagedArtifactTarget) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetId

`func (o *ManagedArtifactTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedArtifactTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedArtifactTarget) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ManagedArtifactTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedArtifactTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedArtifactTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedArtifactTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *ManagedArtifactTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ManagedArtifactTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ManagedArtifactTarget) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ManagedArtifactTarget) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


