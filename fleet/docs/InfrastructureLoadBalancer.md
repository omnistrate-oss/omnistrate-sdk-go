# InfrastructureLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]InfrastructureAddress**](InfrastructureAddress.md) |  | [optional] 
**AddressesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Annotations** | Pointer to [**[]InfrastructureKeyValue**](InfrastructureKeyValue.md) |  | [optional] 
**AnnotationsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Identity** | Pointer to **string** |  | [optional] 
**IdentityAvailability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**IdentityProvenance** | Pointer to **string** |  | [optional] 
**PlacementAvailability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**Source** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**TargetHealthAvailability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 

## Methods

### NewInfrastructureLoadBalancer

`func NewInfrastructureLoadBalancer() *InfrastructureLoadBalancer`

NewInfrastructureLoadBalancer instantiates a new InfrastructureLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureLoadBalancerWithDefaults

`func NewInfrastructureLoadBalancerWithDefaults() *InfrastructureLoadBalancer`

NewInfrastructureLoadBalancerWithDefaults instantiates a new InfrastructureLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *InfrastructureLoadBalancer) GetAddresses() []InfrastructureAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InfrastructureLoadBalancer) GetAddressesOk() (*[]InfrastructureAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InfrastructureLoadBalancer) SetAddresses(v []InfrastructureAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InfrastructureLoadBalancer) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAddressesTruncation

`func (o *InfrastructureLoadBalancer) GetAddressesTruncation() CheckpointTruncation`

GetAddressesTruncation returns the AddressesTruncation field if non-nil, zero value otherwise.

### GetAddressesTruncationOk

`func (o *InfrastructureLoadBalancer) GetAddressesTruncationOk() (*CheckpointTruncation, bool)`

GetAddressesTruncationOk returns a tuple with the AddressesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressesTruncation

`func (o *InfrastructureLoadBalancer) SetAddressesTruncation(v CheckpointTruncation)`

SetAddressesTruncation sets AddressesTruncation field to given value.

### HasAddressesTruncation

`func (o *InfrastructureLoadBalancer) HasAddressesTruncation() bool`

HasAddressesTruncation returns a boolean if a field has been set.

### GetAnnotations

`func (o *InfrastructureLoadBalancer) GetAnnotations() []InfrastructureKeyValue`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *InfrastructureLoadBalancer) GetAnnotationsOk() (*[]InfrastructureKeyValue, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *InfrastructureLoadBalancer) SetAnnotations(v []InfrastructureKeyValue)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *InfrastructureLoadBalancer) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAnnotationsTruncation

`func (o *InfrastructureLoadBalancer) GetAnnotationsTruncation() CheckpointTruncation`

GetAnnotationsTruncation returns the AnnotationsTruncation field if non-nil, zero value otherwise.

### GetAnnotationsTruncationOk

`func (o *InfrastructureLoadBalancer) GetAnnotationsTruncationOk() (*CheckpointTruncation, bool)`

GetAnnotationsTruncationOk returns a tuple with the AnnotationsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationsTruncation

`func (o *InfrastructureLoadBalancer) SetAnnotationsTruncation(v CheckpointTruncation)`

SetAnnotationsTruncation sets AnnotationsTruncation field to given value.

### HasAnnotationsTruncation

`func (o *InfrastructureLoadBalancer) HasAnnotationsTruncation() bool`

HasAnnotationsTruncation returns a boolean if a field has been set.

### GetIdentity

`func (o *InfrastructureLoadBalancer) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *InfrastructureLoadBalancer) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *InfrastructureLoadBalancer) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *InfrastructureLoadBalancer) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIdentityAvailability

`func (o *InfrastructureLoadBalancer) GetIdentityAvailability() InfrastructureAvailability`

GetIdentityAvailability returns the IdentityAvailability field if non-nil, zero value otherwise.

### GetIdentityAvailabilityOk

`func (o *InfrastructureLoadBalancer) GetIdentityAvailabilityOk() (*InfrastructureAvailability, bool)`

GetIdentityAvailabilityOk returns a tuple with the IdentityAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAvailability

`func (o *InfrastructureLoadBalancer) SetIdentityAvailability(v InfrastructureAvailability)`

SetIdentityAvailability sets IdentityAvailability field to given value.

### HasIdentityAvailability

`func (o *InfrastructureLoadBalancer) HasIdentityAvailability() bool`

HasIdentityAvailability returns a boolean if a field has been set.

### GetIdentityProvenance

`func (o *InfrastructureLoadBalancer) GetIdentityProvenance() string`

GetIdentityProvenance returns the IdentityProvenance field if non-nil, zero value otherwise.

### GetIdentityProvenanceOk

`func (o *InfrastructureLoadBalancer) GetIdentityProvenanceOk() (*string, bool)`

GetIdentityProvenanceOk returns a tuple with the IdentityProvenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvenance

`func (o *InfrastructureLoadBalancer) SetIdentityProvenance(v string)`

SetIdentityProvenance sets IdentityProvenance field to given value.

### HasIdentityProvenance

`func (o *InfrastructureLoadBalancer) HasIdentityProvenance() bool`

HasIdentityProvenance returns a boolean if a field has been set.

### GetPlacementAvailability

`func (o *InfrastructureLoadBalancer) GetPlacementAvailability() InfrastructureAvailability`

GetPlacementAvailability returns the PlacementAvailability field if non-nil, zero value otherwise.

### GetPlacementAvailabilityOk

`func (o *InfrastructureLoadBalancer) GetPlacementAvailabilityOk() (*InfrastructureAvailability, bool)`

GetPlacementAvailabilityOk returns a tuple with the PlacementAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementAvailability

`func (o *InfrastructureLoadBalancer) SetPlacementAvailability(v InfrastructureAvailability)`

SetPlacementAvailability sets PlacementAvailability field to given value.

### HasPlacementAvailability

`func (o *InfrastructureLoadBalancer) HasPlacementAvailability() bool`

HasPlacementAvailability returns a boolean if a field has been set.

### GetSource

`func (o *InfrastructureLoadBalancer) GetSource() InfrastructureObjectReference`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InfrastructureLoadBalancer) GetSourceOk() (*InfrastructureObjectReference, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InfrastructureLoadBalancer) SetSource(v InfrastructureObjectReference)`

SetSource sets Source field to given value.

### HasSource

`func (o *InfrastructureLoadBalancer) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTargetHealthAvailability

`func (o *InfrastructureLoadBalancer) GetTargetHealthAvailability() InfrastructureAvailability`

GetTargetHealthAvailability returns the TargetHealthAvailability field if non-nil, zero value otherwise.

### GetTargetHealthAvailabilityOk

`func (o *InfrastructureLoadBalancer) GetTargetHealthAvailabilityOk() (*InfrastructureAvailability, bool)`

GetTargetHealthAvailabilityOk returns a tuple with the TargetHealthAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHealthAvailability

`func (o *InfrastructureLoadBalancer) SetTargetHealthAvailability(v InfrastructureAvailability)`

SetTargetHealthAvailability sets TargetHealthAvailability field to given value.

### HasTargetHealthAvailability

`func (o *InfrastructureLoadBalancer) HasTargetHealthAvailability() bool`

HasTargetHealthAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


