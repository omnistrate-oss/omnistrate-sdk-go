# ListResourceCapabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**[]ResourceCapability**](ResourceCapability.md) | The configurationParameters to enable capabilities for the resource | [optional] 

## Methods

### NewListResourceCapabilitiesResponse

`func NewListResourceCapabilitiesResponse() *ListResourceCapabilitiesResponse`

NewListResourceCapabilitiesResponse instantiates a new ListResourceCapabilitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourceCapabilitiesResponseWithDefaults

`func NewListResourceCapabilitiesResponseWithDefaults() *ListResourceCapabilitiesResponse`

NewListResourceCapabilitiesResponseWithDefaults instantiates a new ListResourceCapabilitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *ListResourceCapabilitiesResponse) GetCapabilities() []ResourceCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ListResourceCapabilitiesResponse) GetCapabilitiesOk() (*[]ResourceCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ListResourceCapabilitiesResponse) SetCapabilities(v []ResourceCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ListResourceCapabilitiesResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


