# FleetCreateServicesOrchestrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrchestrationCreateDSL** | **string** | base64 encoded content of service orchestration create DSL | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetCreateServicesOrchestrationRequest

`func NewFleetCreateServicesOrchestrationRequest(orchestrationCreateDSL string, token string, ) *FleetCreateServicesOrchestrationRequest`

NewFleetCreateServicesOrchestrationRequest instantiates a new FleetCreateServicesOrchestrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCreateServicesOrchestrationRequestWithDefaults

`func NewFleetCreateServicesOrchestrationRequestWithDefaults() *FleetCreateServicesOrchestrationRequest`

NewFleetCreateServicesOrchestrationRequestWithDefaults instantiates a new FleetCreateServicesOrchestrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchestrationCreateDSL

`func (o *FleetCreateServicesOrchestrationRequest) GetOrchestrationCreateDSL() string`

GetOrchestrationCreateDSL returns the OrchestrationCreateDSL field if non-nil, zero value otherwise.

### GetOrchestrationCreateDSLOk

`func (o *FleetCreateServicesOrchestrationRequest) GetOrchestrationCreateDSLOk() (*string, bool)`

GetOrchestrationCreateDSLOk returns a tuple with the OrchestrationCreateDSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrationCreateDSL

`func (o *FleetCreateServicesOrchestrationRequest) SetOrchestrationCreateDSL(v string)`

SetOrchestrationCreateDSL sets OrchestrationCreateDSL field to given value.


### GetToken

`func (o *FleetCreateServicesOrchestrationRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetCreateServicesOrchestrationRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetCreateServicesOrchestrationRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


