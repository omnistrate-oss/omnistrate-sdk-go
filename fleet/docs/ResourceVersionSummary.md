# ResourceVersionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**[]ResourceCapability**](ResourceCapability.md) | The capabilities enabled for the resource | [optional] 
**ExternalResource** | Pointer to **bool** | Whether the resource is external. | [optional] 
**LatestVersion** | Pointer to **string** | The latest version of the resource. | [optional] 
**ResourceId** | Pointer to **string** | ID of a resource | [optional] 
**ResourceName** | Pointer to **string** | The name of the resource. | [optional] 
**Version** | Pointer to **string** | The version of the resource deployed for the instance. | [optional] 

## Methods

### NewResourceVersionSummary

`func NewResourceVersionSummary() *ResourceVersionSummary`

NewResourceVersionSummary instantiates a new ResourceVersionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceVersionSummaryWithDefaults

`func NewResourceVersionSummaryWithDefaults() *ResourceVersionSummary`

NewResourceVersionSummaryWithDefaults instantiates a new ResourceVersionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *ResourceVersionSummary) GetCapabilities() []ResourceCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ResourceVersionSummary) GetCapabilitiesOk() (*[]ResourceCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ResourceVersionSummary) SetCapabilities(v []ResourceCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ResourceVersionSummary) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetExternalResource

`func (o *ResourceVersionSummary) GetExternalResource() bool`

GetExternalResource returns the ExternalResource field if non-nil, zero value otherwise.

### GetExternalResourceOk

`func (o *ResourceVersionSummary) GetExternalResourceOk() (*bool, bool)`

GetExternalResourceOk returns a tuple with the ExternalResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalResource

`func (o *ResourceVersionSummary) SetExternalResource(v bool)`

SetExternalResource sets ExternalResource field to given value.

### HasExternalResource

`func (o *ResourceVersionSummary) HasExternalResource() bool`

HasExternalResource returns a boolean if a field has been set.

### GetLatestVersion

`func (o *ResourceVersionSummary) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *ResourceVersionSummary) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *ResourceVersionSummary) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *ResourceVersionSummary) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetResourceId

`func (o *ResourceVersionSummary) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceVersionSummary) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceVersionSummary) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ResourceVersionSummary) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *ResourceVersionSummary) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceVersionSummary) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceVersionSummary) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ResourceVersionSummary) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetVersion

`func (o *ResourceVersionSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResourceVersionSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResourceVersionSummary) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResourceVersionSummary) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


