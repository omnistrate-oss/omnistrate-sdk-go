# ModifyServicesOrchestrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Services Orchestration | 
**OrchestrationModifyDSL** | **string** | base64 encoded content of services orchestration modify DSL | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewModifyServicesOrchestrationRequest

`func NewModifyServicesOrchestrationRequest(id string, orchestrationModifyDSL string, token string, ) *ModifyServicesOrchestrationRequest`

NewModifyServicesOrchestrationRequest instantiates a new ModifyServicesOrchestrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyServicesOrchestrationRequestWithDefaults

`func NewModifyServicesOrchestrationRequestWithDefaults() *ModifyServicesOrchestrationRequest`

NewModifyServicesOrchestrationRequestWithDefaults instantiates a new ModifyServicesOrchestrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModifyServicesOrchestrationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyServicesOrchestrationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyServicesOrchestrationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOrchestrationModifyDSL

`func (o *ModifyServicesOrchestrationRequest) GetOrchestrationModifyDSL() string`

GetOrchestrationModifyDSL returns the OrchestrationModifyDSL field if non-nil, zero value otherwise.

### GetOrchestrationModifyDSLOk

`func (o *ModifyServicesOrchestrationRequest) GetOrchestrationModifyDSLOk() (*string, bool)`

GetOrchestrationModifyDSLOk returns a tuple with the OrchestrationModifyDSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrationModifyDSL

`func (o *ModifyServicesOrchestrationRequest) SetOrchestrationModifyDSL(v string)`

SetOrchestrationModifyDSL sets OrchestrationModifyDSL field to given value.


### GetToken

`func (o *ModifyServicesOrchestrationRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ModifyServicesOrchestrationRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ModifyServicesOrchestrationRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


