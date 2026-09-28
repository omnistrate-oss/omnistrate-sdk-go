# ManagedArtifactSyncWorkflowDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactResultCount** | **int64** | Number of artifact results persisted for this synchronization. | 
**BundleRevisionId** | **string** | Identifier of the managed artifact bundle revision being synchronized. | 
**BundleVersion** | **string** | Release version of the managed artifact bundle being synchronized. | 
**LastTransitionTime** | Pointer to **string** | Time of the latest synchronization state transition in RFC3339 format. | [optional] 
**ProvisionerHostClusterId** | **string** | ID of a Host Cluster | 
**SyncId** | **string** | Stable identifier of the managed artifact synchronization record. | 

## Methods

### NewManagedArtifactSyncWorkflowDetail

`func NewManagedArtifactSyncWorkflowDetail(artifactResultCount int64, bundleRevisionId string, bundleVersion string, provisionerHostClusterId string, syncId string, ) *ManagedArtifactSyncWorkflowDetail`

NewManagedArtifactSyncWorkflowDetail instantiates a new ManagedArtifactSyncWorkflowDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedArtifactSyncWorkflowDetailWithDefaults

`func NewManagedArtifactSyncWorkflowDetailWithDefaults() *ManagedArtifactSyncWorkflowDetail`

NewManagedArtifactSyncWorkflowDetailWithDefaults instantiates a new ManagedArtifactSyncWorkflowDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactResultCount

`func (o *ManagedArtifactSyncWorkflowDetail) GetArtifactResultCount() int64`

GetArtifactResultCount returns the ArtifactResultCount field if non-nil, zero value otherwise.

### GetArtifactResultCountOk

`func (o *ManagedArtifactSyncWorkflowDetail) GetArtifactResultCountOk() (*int64, bool)`

GetArtifactResultCountOk returns a tuple with the ArtifactResultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactResultCount

`func (o *ManagedArtifactSyncWorkflowDetail) SetArtifactResultCount(v int64)`

SetArtifactResultCount sets ArtifactResultCount field to given value.


### GetBundleRevisionId

`func (o *ManagedArtifactSyncWorkflowDetail) GetBundleRevisionId() string`

GetBundleRevisionId returns the BundleRevisionId field if non-nil, zero value otherwise.

### GetBundleRevisionIdOk

`func (o *ManagedArtifactSyncWorkflowDetail) GetBundleRevisionIdOk() (*string, bool)`

GetBundleRevisionIdOk returns a tuple with the BundleRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleRevisionId

`func (o *ManagedArtifactSyncWorkflowDetail) SetBundleRevisionId(v string)`

SetBundleRevisionId sets BundleRevisionId field to given value.


### GetBundleVersion

`func (o *ManagedArtifactSyncWorkflowDetail) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *ManagedArtifactSyncWorkflowDetail) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *ManagedArtifactSyncWorkflowDetail) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.


### GetLastTransitionTime

`func (o *ManagedArtifactSyncWorkflowDetail) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *ManagedArtifactSyncWorkflowDetail) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *ManagedArtifactSyncWorkflowDetail) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *ManagedArtifactSyncWorkflowDetail) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetProvisionerHostClusterId

`func (o *ManagedArtifactSyncWorkflowDetail) GetProvisionerHostClusterId() string`

GetProvisionerHostClusterId returns the ProvisionerHostClusterId field if non-nil, zero value otherwise.

### GetProvisionerHostClusterIdOk

`func (o *ManagedArtifactSyncWorkflowDetail) GetProvisionerHostClusterIdOk() (*string, bool)`

GetProvisionerHostClusterIdOk returns a tuple with the ProvisionerHostClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionerHostClusterId

`func (o *ManagedArtifactSyncWorkflowDetail) SetProvisionerHostClusterId(v string)`

SetProvisionerHostClusterId sets ProvisionerHostClusterId field to given value.


### GetSyncId

`func (o *ManagedArtifactSyncWorkflowDetail) GetSyncId() string`

GetSyncId returns the SyncId field if non-nil, zero value otherwise.

### GetSyncIdOk

`func (o *ManagedArtifactSyncWorkflowDetail) GetSyncIdOk() (*string, bool)`

GetSyncIdOk returns a tuple with the SyncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncId

`func (o *ManagedArtifactSyncWorkflowDetail) SetSyncId(v string)`

SetSyncId sets SyncId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


