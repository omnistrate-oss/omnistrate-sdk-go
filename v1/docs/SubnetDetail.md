# SubnetDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Az** | **string** | The availability zone | 
**Cidr** | Pointer to **string** | The CIDR block of the subnet | [optional] 
**Id** | **string** | The subnet ID | 
**IsPublic** | **bool** | Whether this subnet auto-assigns public IPs | 
**IsTagged** | **bool** | Whether this subnet has the omnistrate managed-by tag (usable for deployment) | 

## Methods

### NewSubnetDetail

`func NewSubnetDetail(az string, id string, isPublic bool, isTagged bool, ) *SubnetDetail`

NewSubnetDetail instantiates a new SubnetDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetDetailWithDefaults

`func NewSubnetDetailWithDefaults() *SubnetDetail`

NewSubnetDetailWithDefaults instantiates a new SubnetDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAz

`func (o *SubnetDetail) GetAz() string`

GetAz returns the Az field if non-nil, zero value otherwise.

### GetAzOk

`func (o *SubnetDetail) GetAzOk() (*string, bool)`

GetAzOk returns a tuple with the Az field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAz

`func (o *SubnetDetail) SetAz(v string)`

SetAz sets Az field to given value.


### GetCidr

`func (o *SubnetDetail) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *SubnetDetail) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *SubnetDetail) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *SubnetDetail) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetId

`func (o *SubnetDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubnetDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubnetDetail) SetId(v string)`

SetId sets Id field to given value.


### GetIsPublic

`func (o *SubnetDetail) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *SubnetDetail) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *SubnetDetail) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetIsTagged

`func (o *SubnetDetail) GetIsTagged() bool`

GetIsTagged returns the IsTagged field if non-nil, zero value otherwise.

### GetIsTaggedOk

`func (o *SubnetDetail) GetIsTaggedOk() (*bool, bool)`

GetIsTaggedOk returns a tuple with the IsTagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTagged

`func (o *SubnetDetail) SetIsTagged(v bool)`

SetIsTagged sets IsTagged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


