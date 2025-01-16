# DescribeServiceOfferingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time the service was created | 
**IsDeprecated** | **bool** | Whether the service offering is deprecated | [default to false]
**Offerings** | [**[]ServiceOffering**](ServiceOffering.md) | The service offerings | 
**ServiceDescription** | **string** | The service description | 
**ServiceId** | **string** | ID of a Service | 
**ServiceName** | **string** | The service name | 
**ServiceOrgId** | **string** | ID of an Org | 
**ServiceProviderId** | **string** | The id of the service provider | 
**ServiceProviderName** | **string** | The name of the service provider | 
**ServiceURLKey** | **string** | The service URL key | 

## Methods

### NewDescribeServiceOfferingResult

`func NewDescribeServiceOfferingResult(createdAt string, isDeprecated bool, offerings []ServiceOffering, serviceDescription string, serviceId string, serviceName string, serviceOrgId string, serviceProviderId string, serviceProviderName string, serviceURLKey string, ) *DescribeServiceOfferingResult`

NewDescribeServiceOfferingResult instantiates a new DescribeServiceOfferingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceOfferingResultWithDefaults

`func NewDescribeServiceOfferingResultWithDefaults() *DescribeServiceOfferingResult`

NewDescribeServiceOfferingResultWithDefaults instantiates a new DescribeServiceOfferingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DescribeServiceOfferingResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DescribeServiceOfferingResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DescribeServiceOfferingResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsDeprecated

`func (o *DescribeServiceOfferingResult) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *DescribeServiceOfferingResult) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *DescribeServiceOfferingResult) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.


### GetOfferings

`func (o *DescribeServiceOfferingResult) GetOfferings() []ServiceOffering`

GetOfferings returns the Offerings field if non-nil, zero value otherwise.

### GetOfferingsOk

`func (o *DescribeServiceOfferingResult) GetOfferingsOk() (*[]ServiceOffering, bool)`

GetOfferingsOk returns a tuple with the Offerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferings

`func (o *DescribeServiceOfferingResult) SetOfferings(v []ServiceOffering)`

SetOfferings sets Offerings field to given value.


### GetServiceDescription

`func (o *DescribeServiceOfferingResult) GetServiceDescription() string`

GetServiceDescription returns the ServiceDescription field if non-nil, zero value otherwise.

### GetServiceDescriptionOk

`func (o *DescribeServiceOfferingResult) GetServiceDescriptionOk() (*string, bool)`

GetServiceDescriptionOk returns a tuple with the ServiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDescription

`func (o *DescribeServiceOfferingResult) SetServiceDescription(v string)`

SetServiceDescription sets ServiceDescription field to given value.


### GetServiceId

`func (o *DescribeServiceOfferingResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeServiceOfferingResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeServiceOfferingResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceName

`func (o *DescribeServiceOfferingResult) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DescribeServiceOfferingResult) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DescribeServiceOfferingResult) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServiceOrgId

`func (o *DescribeServiceOfferingResult) GetServiceOrgId() string`

GetServiceOrgId returns the ServiceOrgId field if non-nil, zero value otherwise.

### GetServiceOrgIdOk

`func (o *DescribeServiceOfferingResult) GetServiceOrgIdOk() (*string, bool)`

GetServiceOrgIdOk returns a tuple with the ServiceOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOrgId

`func (o *DescribeServiceOfferingResult) SetServiceOrgId(v string)`

SetServiceOrgId sets ServiceOrgId field to given value.


### GetServiceProviderId

`func (o *DescribeServiceOfferingResult) GetServiceProviderId() string`

GetServiceProviderId returns the ServiceProviderId field if non-nil, zero value otherwise.

### GetServiceProviderIdOk

`func (o *DescribeServiceOfferingResult) GetServiceProviderIdOk() (*string, bool)`

GetServiceProviderIdOk returns a tuple with the ServiceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderId

`func (o *DescribeServiceOfferingResult) SetServiceProviderId(v string)`

SetServiceProviderId sets ServiceProviderId field to given value.


### GetServiceProviderName

`func (o *DescribeServiceOfferingResult) GetServiceProviderName() string`

GetServiceProviderName returns the ServiceProviderName field if non-nil, zero value otherwise.

### GetServiceProviderNameOk

`func (o *DescribeServiceOfferingResult) GetServiceProviderNameOk() (*string, bool)`

GetServiceProviderNameOk returns a tuple with the ServiceProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderName

`func (o *DescribeServiceOfferingResult) SetServiceProviderName(v string)`

SetServiceProviderName sets ServiceProviderName field to given value.


### GetServiceURLKey

`func (o *DescribeServiceOfferingResult) GetServiceURLKey() string`

GetServiceURLKey returns the ServiceURLKey field if non-nil, zero value otherwise.

### GetServiceURLKeyOk

`func (o *DescribeServiceOfferingResult) GetServiceURLKeyOk() (*string, bool)`

GetServiceURLKeyOk returns a tuple with the ServiceURLKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceURLKey

`func (o *DescribeServiceOfferingResult) SetServiceURLKey(v string)`

SetServiceURLKey sets ServiceURLKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


