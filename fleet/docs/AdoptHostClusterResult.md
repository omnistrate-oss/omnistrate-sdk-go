# AdoptHostClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdoptionStatus** | **string** | The status of an operation | 
**AgentInstallationKit** | **string** | The base64 encoded TAR archive containing the agent installation kit | 

## Methods

### NewAdoptHostClusterResult

`func NewAdoptHostClusterResult(adoptionStatus string, agentInstallationKit string, ) *AdoptHostClusterResult`

NewAdoptHostClusterResult instantiates a new AdoptHostClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdoptHostClusterResultWithDefaults

`func NewAdoptHostClusterResultWithDefaults() *AdoptHostClusterResult`

NewAdoptHostClusterResultWithDefaults instantiates a new AdoptHostClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdoptionStatus

`func (o *AdoptHostClusterResult) GetAdoptionStatus() string`

GetAdoptionStatus returns the AdoptionStatus field if non-nil, zero value otherwise.

### GetAdoptionStatusOk

`func (o *AdoptHostClusterResult) GetAdoptionStatusOk() (*string, bool)`

GetAdoptionStatusOk returns a tuple with the AdoptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptionStatus

`func (o *AdoptHostClusterResult) SetAdoptionStatus(v string)`

SetAdoptionStatus sets AdoptionStatus field to given value.


### GetAgentInstallationKit

`func (o *AdoptHostClusterResult) GetAgentInstallationKit() string`

GetAgentInstallationKit returns the AgentInstallationKit field if non-nil, zero value otherwise.

### GetAgentInstallationKitOk

`func (o *AdoptHostClusterResult) GetAgentInstallationKitOk() (*string, bool)`

GetAgentInstallationKitOk returns a tuple with the AgentInstallationKit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstallationKit

`func (o *AdoptHostClusterResult) SetAgentInstallationKit(v string)`

SetAgentInstallationKit sets AgentInstallationKit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


