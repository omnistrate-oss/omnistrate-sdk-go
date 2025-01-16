# CreateServicesOrchestrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrchestrationCreateDSL** | **string** | base64 encoded content of service orchestration create DSL | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateServicesOrchestrationRequest

`func NewCreateServicesOrchestrationRequest(orchestrationCreateDSL string, token string, ) *CreateServicesOrchestrationRequest`

NewCreateServicesOrchestrationRequest instantiates a new CreateServicesOrchestrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServicesOrchestrationRequestWithDefaults

`func NewCreateServicesOrchestrationRequestWithDefaults() *CreateServicesOrchestrationRequest`

NewCreateServicesOrchestrationRequestWithDefaults instantiates a new CreateServicesOrchestrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchestrationCreateDSL

`func (o *CreateServicesOrchestrationRequest) GetOrchestrationCreateDSL() string`

GetOrchestrationCreateDSL returns the OrchestrationCreateDSL field if non-nil, zero value otherwise.

### GetOrchestrationCreateDSLOk

`func (o *CreateServicesOrchestrationRequest) GetOrchestrationCreateDSLOk() (*string, bool)`

GetOrchestrationCreateDSLOk returns a tuple with the OrchestrationCreateDSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrationCreateDSL

`func (o *CreateServicesOrchestrationRequest) SetOrchestrationCreateDSL(v string)`

SetOrchestrationCreateDSL sets OrchestrationCreateDSL field to given value.


### GetToken

`func (o *CreateServicesOrchestrationRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateServicesOrchestrationRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateServicesOrchestrationRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


