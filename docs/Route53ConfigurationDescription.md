# Route53ConfigurationDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountID** | **string** | The AWS account hosting the the hosted zone for the domain | 
**AwsCloudFormationTemplateURL** | **string** | The URL to the CloudFormation template | 

## Methods

### NewRoute53ConfigurationDescription

`func NewRoute53ConfigurationDescription(awsAccountID string, awsCloudFormationTemplateURL string, ) *Route53ConfigurationDescription`

NewRoute53ConfigurationDescription instantiates a new Route53ConfigurationDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute53ConfigurationDescriptionWithDefaults

`func NewRoute53ConfigurationDescriptionWithDefaults() *Route53ConfigurationDescription`

NewRoute53ConfigurationDescriptionWithDefaults instantiates a new Route53ConfigurationDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountID

`func (o *Route53ConfigurationDescription) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *Route53ConfigurationDescription) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *Route53ConfigurationDescription) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.


### GetAwsCloudFormationTemplateURL

`func (o *Route53ConfigurationDescription) GetAwsCloudFormationTemplateURL() string`

GetAwsCloudFormationTemplateURL returns the AwsCloudFormationTemplateURL field if non-nil, zero value otherwise.

### GetAwsCloudFormationTemplateURLOk

`func (o *Route53ConfigurationDescription) GetAwsCloudFormationTemplateURLOk() (*string, bool)`

GetAwsCloudFormationTemplateURLOk returns a tuple with the AwsCloudFormationTemplateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudFormationTemplateURL

`func (o *Route53ConfigurationDescription) SetAwsCloudFormationTemplateURL(v string)`

SetAwsCloudFormationTemplateURL sets AwsCloudFormationTemplateURL field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


