# FleetListServicesOrchestrationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | **string** | The type of service environment | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetListServicesOrchestrationsRequest

`func NewFleetListServicesOrchestrationsRequest(environmentType string, token string, ) *FleetListServicesOrchestrationsRequest`

NewFleetListServicesOrchestrationsRequest instantiates a new FleetListServicesOrchestrationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListServicesOrchestrationsRequestWithDefaults

`func NewFleetListServicesOrchestrationsRequestWithDefaults() *FleetListServicesOrchestrationsRequest`

NewFleetListServicesOrchestrationsRequestWithDefaults instantiates a new FleetListServicesOrchestrationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *FleetListServicesOrchestrationsRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *FleetListServicesOrchestrationsRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *FleetListServicesOrchestrationsRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetToken

`func (o *FleetListServicesOrchestrationsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetListServicesOrchestrationsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetListServicesOrchestrationsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


