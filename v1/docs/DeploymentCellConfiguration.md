# DeploymentCellConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amenities** | Pointer to [**[]Amenity**](Amenity.md) | The amenities available in the deployment cell. | [optional] 
**WorkloadIdentities** | Pointer to [**[]ManagedWorkloadIdentity**](ManagedWorkloadIdentity.md) | The managed workload identities available in the deployment cell. | [optional] 

## Methods

### NewDeploymentCellConfiguration

`func NewDeploymentCellConfiguration() *DeploymentCellConfiguration`

NewDeploymentCellConfiguration instantiates a new DeploymentCellConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellConfigurationWithDefaults

`func NewDeploymentCellConfigurationWithDefaults() *DeploymentCellConfiguration`

NewDeploymentCellConfigurationWithDefaults instantiates a new DeploymentCellConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmenities

`func (o *DeploymentCellConfiguration) GetAmenities() []Amenity`

GetAmenities returns the Amenities field if non-nil, zero value otherwise.

### GetAmenitiesOk

`func (o *DeploymentCellConfiguration) GetAmenitiesOk() (*[]Amenity, bool)`

GetAmenitiesOk returns a tuple with the Amenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenities

`func (o *DeploymentCellConfiguration) SetAmenities(v []Amenity)`

SetAmenities sets Amenities field to given value.

### HasAmenities

`func (o *DeploymentCellConfiguration) HasAmenities() bool`

HasAmenities returns a boolean if a field has been set.

### GetWorkloadIdentities

`func (o *DeploymentCellConfiguration) GetWorkloadIdentities() []ManagedWorkloadIdentity`

GetWorkloadIdentities returns the WorkloadIdentities field if non-nil, zero value otherwise.

### GetWorkloadIdentitiesOk

`func (o *DeploymentCellConfiguration) GetWorkloadIdentitiesOk() (*[]ManagedWorkloadIdentity, bool)`

GetWorkloadIdentitiesOk returns a tuple with the WorkloadIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIdentities

`func (o *DeploymentCellConfiguration) SetWorkloadIdentities(v []ManagedWorkloadIdentity)`

SetWorkloadIdentities sets WorkloadIdentities field to given value.

### HasWorkloadIdentities

`func (o *DeploymentCellConfiguration) HasWorkloadIdentities() bool`

HasWorkloadIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


