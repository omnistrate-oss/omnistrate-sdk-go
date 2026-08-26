# ManagedReleaseComponentRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | **string** | Component pipeline that owns a managed artifact subset. | 
**SourceCommit** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewManagedReleaseComponentRevision

`func NewManagedReleaseComponentRevision(component string, sourceCommit string, version string, ) *ManagedReleaseComponentRevision`

NewManagedReleaseComponentRevision instantiates a new ManagedReleaseComponentRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedReleaseComponentRevisionWithDefaults

`func NewManagedReleaseComponentRevisionWithDefaults() *ManagedReleaseComponentRevision`

NewManagedReleaseComponentRevisionWithDefaults instantiates a new ManagedReleaseComponentRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ManagedReleaseComponentRevision) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ManagedReleaseComponentRevision) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ManagedReleaseComponentRevision) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetSourceCommit

`func (o *ManagedReleaseComponentRevision) GetSourceCommit() string`

GetSourceCommit returns the SourceCommit field if non-nil, zero value otherwise.

### GetSourceCommitOk

`func (o *ManagedReleaseComponentRevision) GetSourceCommitOk() (*string, bool)`

GetSourceCommitOk returns a tuple with the SourceCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCommit

`func (o *ManagedReleaseComponentRevision) SetSourceCommit(v string)`

SetSourceCommit sets SourceCommit field to given value.


### GetVersion

`func (o *ManagedReleaseComponentRevision) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagedReleaseComponentRevision) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManagedReleaseComponentRevision) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


