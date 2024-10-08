# SaaSPortal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDomain** | Pointer to **string** | The custom domain of the SaaS portal | [optional] 
**CustomDomainStatus** | Pointer to **string** | The custom domain status of the SaaS portal | [optional] 
**DetailedNetworkTopology** | Pointer to [**map[string]ResourceNetworkTopologyResult**](ResourceNetworkTopologyResult.md) | The detailed network topology of the SaaS portal | [optional] 
**EmailConfig** | Pointer to [**SaaSPortalEmailConfig**](SaaSPortalEmailConfig.md) |  | [optional] 
**Endpoint** | Pointer to **string** | The endpoint of the SaaS portal for this environment type | [optional] 
**EnvironmentType** | **string** | The environment type for the SaaS portal | 
**GoogleAnalyticsTagID** | Pointer to **string** | The Google Analytics tag ID for the SaaS portal | [optional] 
**ImageConfig** | Pointer to [**SaaSPortalImageConfig**](SaaSPortalImageConfig.md) |  | [optional] 
**Status** | **string** | The status of the SaaS portal for this environment type | 

## Methods

### NewSaaSPortal

`func NewSaaSPortal(environmentType string, status string, ) *SaaSPortal`

NewSaaSPortal instantiates a new SaaSPortal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaaSPortalWithDefaults

`func NewSaaSPortalWithDefaults() *SaaSPortal`

NewSaaSPortalWithDefaults instantiates a new SaaSPortal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDomain

`func (o *SaaSPortal) GetCustomDomain() string`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *SaaSPortal) GetCustomDomainOk() (*string, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *SaaSPortal) SetCustomDomain(v string)`

SetCustomDomain sets CustomDomain field to given value.

### HasCustomDomain

`func (o *SaaSPortal) HasCustomDomain() bool`

HasCustomDomain returns a boolean if a field has been set.

### GetCustomDomainStatus

`func (o *SaaSPortal) GetCustomDomainStatus() string`

GetCustomDomainStatus returns the CustomDomainStatus field if non-nil, zero value otherwise.

### GetCustomDomainStatusOk

`func (o *SaaSPortal) GetCustomDomainStatusOk() (*string, bool)`

GetCustomDomainStatusOk returns a tuple with the CustomDomainStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomainStatus

`func (o *SaaSPortal) SetCustomDomainStatus(v string)`

SetCustomDomainStatus sets CustomDomainStatus field to given value.

### HasCustomDomainStatus

`func (o *SaaSPortal) HasCustomDomainStatus() bool`

HasCustomDomainStatus returns a boolean if a field has been set.

### GetDetailedNetworkTopology

`func (o *SaaSPortal) GetDetailedNetworkTopology() map[string]ResourceNetworkTopologyResult`

GetDetailedNetworkTopology returns the DetailedNetworkTopology field if non-nil, zero value otherwise.

### GetDetailedNetworkTopologyOk

`func (o *SaaSPortal) GetDetailedNetworkTopologyOk() (*map[string]ResourceNetworkTopologyResult, bool)`

GetDetailedNetworkTopologyOk returns a tuple with the DetailedNetworkTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedNetworkTopology

`func (o *SaaSPortal) SetDetailedNetworkTopology(v map[string]ResourceNetworkTopologyResult)`

SetDetailedNetworkTopology sets DetailedNetworkTopology field to given value.

### HasDetailedNetworkTopology

`func (o *SaaSPortal) HasDetailedNetworkTopology() bool`

HasDetailedNetworkTopology returns a boolean if a field has been set.

### GetEmailConfig

`func (o *SaaSPortal) GetEmailConfig() SaaSPortalEmailConfig`

GetEmailConfig returns the EmailConfig field if non-nil, zero value otherwise.

### GetEmailConfigOk

`func (o *SaaSPortal) GetEmailConfigOk() (*SaaSPortalEmailConfig, bool)`

GetEmailConfigOk returns a tuple with the EmailConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfig

`func (o *SaaSPortal) SetEmailConfig(v SaaSPortalEmailConfig)`

SetEmailConfig sets EmailConfig field to given value.

### HasEmailConfig

`func (o *SaaSPortal) HasEmailConfig() bool`

HasEmailConfig returns a boolean if a field has been set.

### GetEndpoint

`func (o *SaaSPortal) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SaaSPortal) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SaaSPortal) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *SaaSPortal) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *SaaSPortal) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *SaaSPortal) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *SaaSPortal) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetGoogleAnalyticsTagID

`func (o *SaaSPortal) GetGoogleAnalyticsTagID() string`

GetGoogleAnalyticsTagID returns the GoogleAnalyticsTagID field if non-nil, zero value otherwise.

### GetGoogleAnalyticsTagIDOk

`func (o *SaaSPortal) GetGoogleAnalyticsTagIDOk() (*string, bool)`

GetGoogleAnalyticsTagIDOk returns a tuple with the GoogleAnalyticsTagID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAnalyticsTagID

`func (o *SaaSPortal) SetGoogleAnalyticsTagID(v string)`

SetGoogleAnalyticsTagID sets GoogleAnalyticsTagID field to given value.

### HasGoogleAnalyticsTagID

`func (o *SaaSPortal) HasGoogleAnalyticsTagID() bool`

HasGoogleAnalyticsTagID returns a boolean if a field has been set.

### GetImageConfig

`func (o *SaaSPortal) GetImageConfig() SaaSPortalImageConfig`

GetImageConfig returns the ImageConfig field if non-nil, zero value otherwise.

### GetImageConfigOk

`func (o *SaaSPortal) GetImageConfigOk() (*SaaSPortalImageConfig, bool)`

GetImageConfigOk returns a tuple with the ImageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfig

`func (o *SaaSPortal) SetImageConfig(v SaaSPortalImageConfig)`

SetImageConfig sets ImageConfig field to given value.

### HasImageConfig

`func (o *SaaSPortal) HasImageConfig() bool`

HasImageConfig returns a boolean if a field has been set.

### GetStatus

`func (o *SaaSPortal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SaaSPortal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SaaSPortal) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


