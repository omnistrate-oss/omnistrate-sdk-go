# SubscriptionLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateBase64** | Pointer to **string** | Public certificate used to sign the license in Base 64 format | [optional] [default to ""]
**ExpirationDate** | Pointer to **string** | Current expiration of the license in RFC 3339 format. The license will be renewed periodically before this date while the subscription is active. | [optional] 
**LicenseBase64** | Pointer to **string** | License file in Base 64 format | [optional] [default to ""]

## Methods

### NewSubscriptionLicense

`func NewSubscriptionLicense() *SubscriptionLicense`

NewSubscriptionLicense instantiates a new SubscriptionLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionLicenseWithDefaults

`func NewSubscriptionLicenseWithDefaults() *SubscriptionLicense`

NewSubscriptionLicenseWithDefaults instantiates a new SubscriptionLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateBase64

`func (o *SubscriptionLicense) GetCertificateBase64() string`

GetCertificateBase64 returns the CertificateBase64 field if non-nil, zero value otherwise.

### GetCertificateBase64Ok

`func (o *SubscriptionLicense) GetCertificateBase64Ok() (*string, bool)`

GetCertificateBase64Ok returns a tuple with the CertificateBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateBase64

`func (o *SubscriptionLicense) SetCertificateBase64(v string)`

SetCertificateBase64 sets CertificateBase64 field to given value.

### HasCertificateBase64

`func (o *SubscriptionLicense) HasCertificateBase64() bool`

HasCertificateBase64 returns a boolean if a field has been set.

### GetExpirationDate

`func (o *SubscriptionLicense) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SubscriptionLicense) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SubscriptionLicense) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SubscriptionLicense) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLicenseBase64

`func (o *SubscriptionLicense) GetLicenseBase64() string`

GetLicenseBase64 returns the LicenseBase64 field if non-nil, zero value otherwise.

### GetLicenseBase64Ok

`func (o *SubscriptionLicense) GetLicenseBase64Ok() (*string, bool)`

GetLicenseBase64Ok returns a tuple with the LicenseBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseBase64

`func (o *SubscriptionLicense) SetLicenseBase64(v string)`

SetLicenseBase64 sets LicenseBase64 field to given value.

### HasLicenseBase64

`func (o *SubscriptionLicense) HasLicenseBase64() bool`

HasLicenseBase64 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


