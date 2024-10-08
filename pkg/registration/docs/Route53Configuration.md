# Route53Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountID** | **string** | The AWS account hosting the the hosted zone for the domain | 

## Methods

### NewRoute53Configuration

`func NewRoute53Configuration(awsAccountID string, ) *Route53Configuration`

NewRoute53Configuration instantiates a new Route53Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute53ConfigurationWithDefaults

`func NewRoute53ConfigurationWithDefaults() *Route53Configuration`

NewRoute53ConfigurationWithDefaults instantiates a new Route53Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountID

`func (o *Route53Configuration) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *Route53Configuration) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *Route53Configuration) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


