# HelmPackageInstallations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HelmPackage** | [**HelmPackage**](HelmPackage.md) |  | 
**HostClusterID** | **string** | The ID of the host cluster where the helm package is installed | 
**Status** | **string** | The status of the Helm package installation | 

## Methods

### NewHelmPackageInstallations

`func NewHelmPackageInstallations(helmPackage HelmPackage, hostClusterID string, status string, ) *HelmPackageInstallations`

NewHelmPackageInstallations instantiates a new HelmPackageInstallations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmPackageInstallationsWithDefaults

`func NewHelmPackageInstallationsWithDefaults() *HelmPackageInstallations`

NewHelmPackageInstallationsWithDefaults instantiates a new HelmPackageInstallations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelmPackage

`func (o *HelmPackageInstallations) GetHelmPackage() HelmPackage`

GetHelmPackage returns the HelmPackage field if non-nil, zero value otherwise.

### GetHelmPackageOk

`func (o *HelmPackageInstallations) GetHelmPackageOk() (*HelmPackage, bool)`

GetHelmPackageOk returns a tuple with the HelmPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmPackage

`func (o *HelmPackageInstallations) SetHelmPackage(v HelmPackage)`

SetHelmPackage sets HelmPackage field to given value.


### GetHostClusterID

`func (o *HelmPackageInstallations) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *HelmPackageInstallations) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *HelmPackageInstallations) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetStatus

`func (o *HelmPackageInstallations) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HelmPackageInstallations) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HelmPackageInstallations) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


