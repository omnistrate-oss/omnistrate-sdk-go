# NodeVMInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** | The architecture of the VM | 
**AzCode** | **string** | The availability zone code of the VM | 
**AzID** | Pointer to **string** | The availability zone ID of the VM | [optional] 
**ExternalIP** | Pointer to **string** | The external IP of the VM | [optional] 
**InstanceType** | **string** | The instance type of the VM | 
**InternalIP** | **string** | The internal IP of the VM | 
**K8sNodeName** | **string** | The k8s node name of the VM | 
**KernelVersion** | Pointer to **string** | The kernel version of the VM | [optional] 
**KubeletVersion** | **string** | The kubelet version of the VM | 
**Os** | **string** | The operating system of the VM | 
**ProviderID** | **string** | The provider ID of the VM | 

## Methods

### NewNodeVMInfoResult

`func NewNodeVMInfoResult(architecture string, azCode string, instanceType string, internalIP string, k8sNodeName string, kubeletVersion string, os string, providerID string, ) *NodeVMInfoResult`

NewNodeVMInfoResult instantiates a new NodeVMInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeVMInfoResultWithDefaults

`func NewNodeVMInfoResultWithDefaults() *NodeVMInfoResult`

NewNodeVMInfoResultWithDefaults instantiates a new NodeVMInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *NodeVMInfoResult) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *NodeVMInfoResult) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *NodeVMInfoResult) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetAzCode

`func (o *NodeVMInfoResult) GetAzCode() string`

GetAzCode returns the AzCode field if non-nil, zero value otherwise.

### GetAzCodeOk

`func (o *NodeVMInfoResult) GetAzCodeOk() (*string, bool)`

GetAzCodeOk returns a tuple with the AzCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCode

`func (o *NodeVMInfoResult) SetAzCode(v string)`

SetAzCode sets AzCode field to given value.


### GetAzID

`func (o *NodeVMInfoResult) GetAzID() string`

GetAzID returns the AzID field if non-nil, zero value otherwise.

### GetAzIDOk

`func (o *NodeVMInfoResult) GetAzIDOk() (*string, bool)`

GetAzIDOk returns a tuple with the AzID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzID

`func (o *NodeVMInfoResult) SetAzID(v string)`

SetAzID sets AzID field to given value.

### HasAzID

`func (o *NodeVMInfoResult) HasAzID() bool`

HasAzID returns a boolean if a field has been set.

### GetExternalIP

`func (o *NodeVMInfoResult) GetExternalIP() string`

GetExternalIP returns the ExternalIP field if non-nil, zero value otherwise.

### GetExternalIPOk

`func (o *NodeVMInfoResult) GetExternalIPOk() (*string, bool)`

GetExternalIPOk returns a tuple with the ExternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIP

`func (o *NodeVMInfoResult) SetExternalIP(v string)`

SetExternalIP sets ExternalIP field to given value.

### HasExternalIP

`func (o *NodeVMInfoResult) HasExternalIP() bool`

HasExternalIP returns a boolean if a field has been set.

### GetInstanceType

`func (o *NodeVMInfoResult) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *NodeVMInfoResult) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *NodeVMInfoResult) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetInternalIP

`func (o *NodeVMInfoResult) GetInternalIP() string`

GetInternalIP returns the InternalIP field if non-nil, zero value otherwise.

### GetInternalIPOk

`func (o *NodeVMInfoResult) GetInternalIPOk() (*string, bool)`

GetInternalIPOk returns a tuple with the InternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIP

`func (o *NodeVMInfoResult) SetInternalIP(v string)`

SetInternalIP sets InternalIP field to given value.


### GetK8sNodeName

`func (o *NodeVMInfoResult) GetK8sNodeName() string`

GetK8sNodeName returns the K8sNodeName field if non-nil, zero value otherwise.

### GetK8sNodeNameOk

`func (o *NodeVMInfoResult) GetK8sNodeNameOk() (*string, bool)`

GetK8sNodeNameOk returns a tuple with the K8sNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNodeName

`func (o *NodeVMInfoResult) SetK8sNodeName(v string)`

SetK8sNodeName sets K8sNodeName field to given value.


### GetKernelVersion

`func (o *NodeVMInfoResult) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *NodeVMInfoResult) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *NodeVMInfoResult) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.

### HasKernelVersion

`func (o *NodeVMInfoResult) HasKernelVersion() bool`

HasKernelVersion returns a boolean if a field has been set.

### GetKubeletVersion

`func (o *NodeVMInfoResult) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *NodeVMInfoResult) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *NodeVMInfoResult) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.


### GetOs

`func (o *NodeVMInfoResult) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *NodeVMInfoResult) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *NodeVMInfoResult) SetOs(v string)`

SetOs sets Os field to given value.


### GetProviderID

`func (o *NodeVMInfoResult) GetProviderID() string`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *NodeVMInfoResult) GetProviderIDOk() (*string, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *NodeVMInfoResult) SetProviderID(v string)`

SetProviderID sets ProviderID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


