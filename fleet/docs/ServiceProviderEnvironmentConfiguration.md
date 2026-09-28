# ServiceProviderEnvironmentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigs** | Pointer to [**[]AccountConfigSummary**](AccountConfigSummary.md) | The account configurations associated with this environment and cloud provider. | [optional] 
**Amenities** | Pointer to [**[]Amenity**](Amenity.md) | The amenities configured for deployment cells in this environment. | [optional] 
**ManagedReleaseVersion** | Pointer to **string** | Optional immutable managed artifact bundle version for this environment and cloud provider. On update, omit it to leave the selection unchanged or send an empty string to clear the pin; an unpinned template uses the READY bundle with the greatest release sequence. | [optional] 
**WorkloadIdentities** | Pointer to [**[]ManagedWorkloadIdentity**](ManagedWorkloadIdentity.md) | The managed workload identities available in the deployment cell. | [optional] 
**WorkloadIdentitiesStatuses** | Pointer to [**map[string]map[string][]string**](map.md) | The account configuration IDs grouped by managed workload identity identifier and workload identity status. | [optional] 

## Methods

### NewServiceProviderEnvironmentConfiguration

`func NewServiceProviderEnvironmentConfiguration() *ServiceProviderEnvironmentConfiguration`

NewServiceProviderEnvironmentConfiguration instantiates a new ServiceProviderEnvironmentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProviderEnvironmentConfigurationWithDefaults

`func NewServiceProviderEnvironmentConfigurationWithDefaults() *ServiceProviderEnvironmentConfiguration`

NewServiceProviderEnvironmentConfigurationWithDefaults instantiates a new ServiceProviderEnvironmentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigs

`func (o *ServiceProviderEnvironmentConfiguration) GetAccountConfigs() []AccountConfigSummary`

GetAccountConfigs returns the AccountConfigs field if non-nil, zero value otherwise.

### GetAccountConfigsOk

`func (o *ServiceProviderEnvironmentConfiguration) GetAccountConfigsOk() (*[]AccountConfigSummary, bool)`

GetAccountConfigsOk returns a tuple with the AccountConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigs

`func (o *ServiceProviderEnvironmentConfiguration) SetAccountConfigs(v []AccountConfigSummary)`

SetAccountConfigs sets AccountConfigs field to given value.

### HasAccountConfigs

`func (o *ServiceProviderEnvironmentConfiguration) HasAccountConfigs() bool`

HasAccountConfigs returns a boolean if a field has been set.

### GetAmenities

`func (o *ServiceProviderEnvironmentConfiguration) GetAmenities() []Amenity`

GetAmenities returns the Amenities field if non-nil, zero value otherwise.

### GetAmenitiesOk

`func (o *ServiceProviderEnvironmentConfiguration) GetAmenitiesOk() (*[]Amenity, bool)`

GetAmenitiesOk returns a tuple with the Amenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenities

`func (o *ServiceProviderEnvironmentConfiguration) SetAmenities(v []Amenity)`

SetAmenities sets Amenities field to given value.

### HasAmenities

`func (o *ServiceProviderEnvironmentConfiguration) HasAmenities() bool`

HasAmenities returns a boolean if a field has been set.

### GetManagedReleaseVersion

`func (o *ServiceProviderEnvironmentConfiguration) GetManagedReleaseVersion() string`

GetManagedReleaseVersion returns the ManagedReleaseVersion field if non-nil, zero value otherwise.

### GetManagedReleaseVersionOk

`func (o *ServiceProviderEnvironmentConfiguration) GetManagedReleaseVersionOk() (*string, bool)`

GetManagedReleaseVersionOk returns a tuple with the ManagedReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedReleaseVersion

`func (o *ServiceProviderEnvironmentConfiguration) SetManagedReleaseVersion(v string)`

SetManagedReleaseVersion sets ManagedReleaseVersion field to given value.

### HasManagedReleaseVersion

`func (o *ServiceProviderEnvironmentConfiguration) HasManagedReleaseVersion() bool`

HasManagedReleaseVersion returns a boolean if a field has been set.

### GetWorkloadIdentities

`func (o *ServiceProviderEnvironmentConfiguration) GetWorkloadIdentities() []ManagedWorkloadIdentity`

GetWorkloadIdentities returns the WorkloadIdentities field if non-nil, zero value otherwise.

### GetWorkloadIdentitiesOk

`func (o *ServiceProviderEnvironmentConfiguration) GetWorkloadIdentitiesOk() (*[]ManagedWorkloadIdentity, bool)`

GetWorkloadIdentitiesOk returns a tuple with the WorkloadIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIdentities

`func (o *ServiceProviderEnvironmentConfiguration) SetWorkloadIdentities(v []ManagedWorkloadIdentity)`

SetWorkloadIdentities sets WorkloadIdentities field to given value.

### HasWorkloadIdentities

`func (o *ServiceProviderEnvironmentConfiguration) HasWorkloadIdentities() bool`

HasWorkloadIdentities returns a boolean if a field has been set.

### GetWorkloadIdentitiesStatuses

`func (o *ServiceProviderEnvironmentConfiguration) GetWorkloadIdentitiesStatuses() map[string]map[string][]string`

GetWorkloadIdentitiesStatuses returns the WorkloadIdentitiesStatuses field if non-nil, zero value otherwise.

### GetWorkloadIdentitiesStatusesOk

`func (o *ServiceProviderEnvironmentConfiguration) GetWorkloadIdentitiesStatusesOk() (*map[string]map[string][]string, bool)`

GetWorkloadIdentitiesStatusesOk returns a tuple with the WorkloadIdentitiesStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIdentitiesStatuses

`func (o *ServiceProviderEnvironmentConfiguration) SetWorkloadIdentitiesStatuses(v map[string]map[string][]string)`

SetWorkloadIdentitiesStatuses sets WorkloadIdentitiesStatuses field to given value.

### HasWorkloadIdentitiesStatuses

`func (o *ServiceProviderEnvironmentConfiguration) HasWorkloadIdentitiesStatuses() bool`

HasWorkloadIdentitiesStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


