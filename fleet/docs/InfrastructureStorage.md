# InfrastructureStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**ClaimTemplates** | Pointer to [**[]InfrastructureClaimTemplate**](InfrastructureClaimTemplate.md) |  | [optional] 
**ClaimTemplatesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Claims** | Pointer to [**[]InfrastructureClaim**](InfrastructureClaim.md) |  | [optional] 
**ClaimsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Mounts** | Pointer to [**[]InfrastructureClaimMount**](InfrastructureClaimMount.md) |  | [optional] 
**MountsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**SortOrder** | Pointer to **string** |  | [optional] 
**StringsTruncated** | Pointer to **bool** |  | [optional] 
**TotalCapacityBytes** | Pointer to **int64** |  | [optional] 

## Methods

### NewInfrastructureStorage

`func NewInfrastructureStorage() *InfrastructureStorage`

NewInfrastructureStorage instantiates a new InfrastructureStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureStorageWithDefaults

`func NewInfrastructureStorageWithDefaults() *InfrastructureStorage`

NewInfrastructureStorageWithDefaults instantiates a new InfrastructureStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *InfrastructureStorage) GetAvailability() InfrastructureAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InfrastructureStorage) GetAvailabilityOk() (*InfrastructureAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InfrastructureStorage) SetAvailability(v InfrastructureAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InfrastructureStorage) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetClaimTemplates

`func (o *InfrastructureStorage) GetClaimTemplates() []InfrastructureClaimTemplate`

GetClaimTemplates returns the ClaimTemplates field if non-nil, zero value otherwise.

### GetClaimTemplatesOk

`func (o *InfrastructureStorage) GetClaimTemplatesOk() (*[]InfrastructureClaimTemplate, bool)`

GetClaimTemplatesOk returns a tuple with the ClaimTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimTemplates

`func (o *InfrastructureStorage) SetClaimTemplates(v []InfrastructureClaimTemplate)`

SetClaimTemplates sets ClaimTemplates field to given value.

### HasClaimTemplates

`func (o *InfrastructureStorage) HasClaimTemplates() bool`

HasClaimTemplates returns a boolean if a field has been set.

### GetClaimTemplatesTruncation

`func (o *InfrastructureStorage) GetClaimTemplatesTruncation() CheckpointTruncation`

GetClaimTemplatesTruncation returns the ClaimTemplatesTruncation field if non-nil, zero value otherwise.

### GetClaimTemplatesTruncationOk

`func (o *InfrastructureStorage) GetClaimTemplatesTruncationOk() (*CheckpointTruncation, bool)`

GetClaimTemplatesTruncationOk returns a tuple with the ClaimTemplatesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimTemplatesTruncation

`func (o *InfrastructureStorage) SetClaimTemplatesTruncation(v CheckpointTruncation)`

SetClaimTemplatesTruncation sets ClaimTemplatesTruncation field to given value.

### HasClaimTemplatesTruncation

`func (o *InfrastructureStorage) HasClaimTemplatesTruncation() bool`

HasClaimTemplatesTruncation returns a boolean if a field has been set.

### GetClaims

`func (o *InfrastructureStorage) GetClaims() []InfrastructureClaim`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *InfrastructureStorage) GetClaimsOk() (*[]InfrastructureClaim, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *InfrastructureStorage) SetClaims(v []InfrastructureClaim)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *InfrastructureStorage) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetClaimsTruncation

`func (o *InfrastructureStorage) GetClaimsTruncation() CheckpointTruncation`

GetClaimsTruncation returns the ClaimsTruncation field if non-nil, zero value otherwise.

### GetClaimsTruncationOk

`func (o *InfrastructureStorage) GetClaimsTruncationOk() (*CheckpointTruncation, bool)`

GetClaimsTruncationOk returns a tuple with the ClaimsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsTruncation

`func (o *InfrastructureStorage) SetClaimsTruncation(v CheckpointTruncation)`

SetClaimsTruncation sets ClaimsTruncation field to given value.

### HasClaimsTruncation

`func (o *InfrastructureStorage) HasClaimsTruncation() bool`

HasClaimsTruncation returns a boolean if a field has been set.

### GetMounts

`func (o *InfrastructureStorage) GetMounts() []InfrastructureClaimMount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *InfrastructureStorage) GetMountsOk() (*[]InfrastructureClaimMount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *InfrastructureStorage) SetMounts(v []InfrastructureClaimMount)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *InfrastructureStorage) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### GetMountsTruncation

`func (o *InfrastructureStorage) GetMountsTruncation() CheckpointTruncation`

GetMountsTruncation returns the MountsTruncation field if non-nil, zero value otherwise.

### GetMountsTruncationOk

`func (o *InfrastructureStorage) GetMountsTruncationOk() (*CheckpointTruncation, bool)`

GetMountsTruncationOk returns a tuple with the MountsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountsTruncation

`func (o *InfrastructureStorage) SetMountsTruncation(v CheckpointTruncation)`

SetMountsTruncation sets MountsTruncation field to given value.

### HasMountsTruncation

`func (o *InfrastructureStorage) HasMountsTruncation() bool`

HasMountsTruncation returns a boolean if a field has been set.

### GetSortOrder

`func (o *InfrastructureStorage) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *InfrastructureStorage) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *InfrastructureStorage) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *InfrastructureStorage) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetStringsTruncated

`func (o *InfrastructureStorage) GetStringsTruncated() bool`

GetStringsTruncated returns the StringsTruncated field if non-nil, zero value otherwise.

### GetStringsTruncatedOk

`func (o *InfrastructureStorage) GetStringsTruncatedOk() (*bool, bool)`

GetStringsTruncatedOk returns a tuple with the StringsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringsTruncated

`func (o *InfrastructureStorage) SetStringsTruncated(v bool)`

SetStringsTruncated sets StringsTruncated field to given value.

### HasStringsTruncated

`func (o *InfrastructureStorage) HasStringsTruncated() bool`

HasStringsTruncated returns a boolean if a field has been set.

### GetTotalCapacityBytes

`func (o *InfrastructureStorage) GetTotalCapacityBytes() int64`

GetTotalCapacityBytes returns the TotalCapacityBytes field if non-nil, zero value otherwise.

### GetTotalCapacityBytesOk

`func (o *InfrastructureStorage) GetTotalCapacityBytesOk() (*int64, bool)`

GetTotalCapacityBytesOk returns a tuple with the TotalCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacityBytes

`func (o *InfrastructureStorage) SetTotalCapacityBytes(v int64)`

SetTotalCapacityBytes sets TotalCapacityBytes field to given value.

### HasTotalCapacityBytes

`func (o *InfrastructureStorage) HasTotalCapacityBytes() bool`

HasTotalCapacityBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


