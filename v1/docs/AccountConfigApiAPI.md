# \AccountConfigApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountConfigApiAccountConfigIdentityID**](AccountConfigApiAPI.md#AccountConfigApiAccountConfigIdentityID) | **Get** /2022-09-01-00/accountconfig/identityid | AccountConfigIdentityID account-config-api
[**AccountConfigApiBulkImportAccountConfigCloudNativeNetworks**](AccountConfigApiAPI.md#AccountConfigApiBulkImportAccountConfigCloudNativeNetworks) | **Post** /2022-09-01-00/accountconfig/{id}/cloud-native-networks/import | BulkImportAccountConfigCloudNativeNetworks account-config-api
[**AccountConfigApiCreateAccountConfig**](AccountConfigApiAPI.md#AccountConfigApiCreateAccountConfig) | **Post** /2022-09-01-00/accountconfig | CreateAccountConfig account-config-api
[**AccountConfigApiDeleteAccountConfig**](AccountConfigApiAPI.md#AccountConfigApiDeleteAccountConfig) | **Delete** /2022-09-01-00/accountconfig/{id} | DeleteAccountConfig account-config-api
[**AccountConfigApiDescribeAccountConfig**](AccountConfigApiAPI.md#AccountConfigApiDescribeAccountConfig) | **Get** /2022-09-01-00/accountconfig/{id} | DescribeAccountConfig account-config-api
[**AccountConfigApiDescribeAccountConfigByAWSAccountID**](AccountConfigApiAPI.md#AccountConfigApiDescribeAccountConfigByAWSAccountID) | **Get** /2022-09-01-00/accountconfig/aws/{awsAccountID} | DescribeAccountConfigByAWSAccountID account-config-api
[**AccountConfigApiDescribeAccountConfigByAzureSubscriptionID**](AccountConfigApiAPI.md#AccountConfigApiDescribeAccountConfigByAzureSubscriptionID) | **Get** /2022-09-01-00/accountconfig/azure/{azureSubscriptionID} | DescribeAccountConfigByAzureSubscriptionID account-config-api
[**AccountConfigApiDescribeAccountConfigByGCPProjectID**](AccountConfigApiAPI.md#AccountConfigApiDescribeAccountConfigByGCPProjectID) | **Get** /2022-09-01-00/accountconfig/gcp/{gcpProjectID} | DescribeAccountConfigByGCPProjectID account-config-api
[**AccountConfigApiDescribeAccountConfigByNebiusTenantID**](AccountConfigApiAPI.md#AccountConfigApiDescribeAccountConfigByNebiusTenantID) | **Get** /2022-09-01-00/accountconfig/nebius/tenant/{nebiusTenantID} | DescribeAccountConfigByNebiusTenantID account-config-api
[**AccountConfigApiDescribeAccountConfigByOCITenancyID**](AccountConfigApiAPI.md#AccountConfigApiDescribeAccountConfigByOCITenancyID) | **Get** /2022-09-01-00/accountconfig/oci/{ociTenancyID} | DescribeAccountConfigByOCITenancyID account-config-api
[**AccountConfigApiImportAccountConfigCloudNativeNetwork**](AccountConfigApiAPI.md#AccountConfigApiImportAccountConfigCloudNativeNetwork) | **Post** /2022-09-01-00/accountconfig/{id}/cloud-native-networks/{cloudNativeNetworkId}/import | ImportAccountConfigCloudNativeNetwork account-config-api
[**AccountConfigApiListAccountConfig**](AccountConfigApiAPI.md#AccountConfigApiListAccountConfig) | **Get** /2022-09-01-00/accountconfig/cloudprovider/{cloudProviderName} | ListAccountConfig account-config-api
[**AccountConfigApiListAccountConfigCloudNativeNetworks**](AccountConfigApiAPI.md#AccountConfigApiListAccountConfigCloudNativeNetworks) | **Get** /2022-09-01-00/accountconfig/{id}/cloud-native-networks | ListAccountConfigCloudNativeNetworks account-config-api
[**AccountConfigApiListBYOAConfig**](AccountConfigApiAPI.md#AccountConfigApiListBYOAConfig) | **Get** /2022-09-01-00/accountconfig/byoa | ListBYOAConfig account-config-api
[**AccountConfigApiSyncAccountConfigCloudNativeNetworks**](AccountConfigApiAPI.md#AccountConfigApiSyncAccountConfigCloudNativeNetworks) | **Post** /2022-09-01-00/accountconfig/{id}/cloud-native-networks/sync | SyncAccountConfigCloudNativeNetworks account-config-api
[**AccountConfigApiUnimportAccountConfigCloudNativeNetwork**](AccountConfigApiAPI.md#AccountConfigApiUnimportAccountConfigCloudNativeNetwork) | **Post** /2022-09-01-00/accountconfig/{id}/cloud-native-networks/{cloudNativeNetworkId}/unimport | UnimportAccountConfigCloudNativeNetwork account-config-api
[**AccountConfigApiUpdateAccountConfig**](AccountConfigApiAPI.md#AccountConfigApiUpdateAccountConfig) | **Put** /2022-09-01-00/accountconfig/{id} | UpdateAccountConfig account-config-api
[**AccountConfigApiVerifyAccountConfig**](AccountConfigApiAPI.md#AccountConfigApiVerifyAccountConfig) | **Post** /2022-09-01-00/accountconfig/verify/{id} | VerifyAccountConfig account-config-api



## AccountConfigApiAccountConfigIdentityID

> AccountConfigIdentityIDResult AccountConfigApiAccountConfigIdentityID(ctx).Execute()

AccountConfigIdentityID account-config-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiAccountConfigIdentityID(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiAccountConfigIdentityID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiAccountConfigIdentityID`: AccountConfigIdentityIDResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiAccountConfigIdentityID`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiAccountConfigIdentityIDRequest struct via the builder pattern


### Return type

[**AccountConfigIdentityIDResult**](AccountConfigIdentityIDResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiBulkImportAccountConfigCloudNativeNetworks

> ListAccountConfigCloudNativeNetworksResult AccountConfigApiBulkImportAccountConfigCloudNativeNetworks(ctx, id).BulkImportAccountConfigCloudNativeNetworksRequest2(bulkImportAccountConfigCloudNativeNetworksRequest2).Execute()

BulkImportAccountConfigCloudNativeNetworks account-config-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "ac-12345678" // string | Account Config ID to operate on
	bulkImportAccountConfigCloudNativeNetworksRequest2 := *openapiclient.NewBulkImportAccountConfigCloudNativeNetworksRequest2([]openapiclient.AccountConfigCloudNativeNetworkOperation{*openapiclient.NewAccountConfigCloudNativeNetworkOperation("vpc-0abc123def456", true)}) // BulkImportAccountConfigCloudNativeNetworksRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiBulkImportAccountConfigCloudNativeNetworks(context.Background(), id).BulkImportAccountConfigCloudNativeNetworksRequest2(bulkImportAccountConfigCloudNativeNetworksRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiBulkImportAccountConfigCloudNativeNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiBulkImportAccountConfigCloudNativeNetworks`: ListAccountConfigCloudNativeNetworksResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiBulkImportAccountConfigCloudNativeNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiBulkImportAccountConfigCloudNativeNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkImportAccountConfigCloudNativeNetworksRequest2** | [**BulkImportAccountConfigCloudNativeNetworksRequest2**](BulkImportAccountConfigCloudNativeNetworksRequest2.md) |  | 

### Return type

[**ListAccountConfigCloudNativeNetworksResult**](ListAccountConfigCloudNativeNetworksResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiCreateAccountConfig

> string AccountConfigApiCreateAccountConfig(ctx).CreateAccountConfigRequest2(createAccountConfigRequest2).Execute()

CreateAccountConfig account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	createAccountConfigRequest2 := *openapiclient.NewCreateAccountConfigRequest2("Omnis nam quisquam modi in.", "An AWS account hosting multiple dev environments", "Dev AWS account") // CreateAccountConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiCreateAccountConfig(context.Background()).CreateAccountConfigRequest2(createAccountConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiCreateAccountConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiCreateAccountConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiCreateAccountConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiCreateAccountConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountConfigRequest2** | [**CreateAccountConfigRequest2**](CreateAccountConfigRequest2.md) |  | 

### Return type

**string**

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiDeleteAccountConfig

> AccountConfigApiDeleteAccountConfig(ctx, id).Execute()

DeleteAccountConfig account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "ac-12345678" // string | Account Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountConfigApiAPI.AccountConfigApiDeleteAccountConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiDeleteAccountConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiDeleteAccountConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiDescribeAccountConfig

> DescribeAccountConfigResult AccountConfigApiDescribeAccountConfig(ctx, id).Execute()

DescribeAccountConfig account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "ac-12345678" // string | Account Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiDescribeAccountConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiDescribeAccountConfig`: DescribeAccountConfigResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiDescribeAccountConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeAccountConfigResult**](DescribeAccountConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiDescribeAccountConfigByAWSAccountID

> DescribeAccountConfigByAWSAccountIDResult AccountConfigApiDescribeAccountConfigByAWSAccountID(ctx, awsAccountID).Execute()

DescribeAccountConfigByAWSAccountID account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	awsAccountID := "123456789012" // string | The AWS account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByAWSAccountID(context.Background(), awsAccountID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByAWSAccountID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiDescribeAccountConfigByAWSAccountID`: DescribeAccountConfigByAWSAccountIDResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByAWSAccountID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsAccountID** | **string** | The AWS account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiDescribeAccountConfigByAWSAccountIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeAccountConfigByAWSAccountIDResult**](DescribeAccountConfigByAWSAccountIDResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiDescribeAccountConfigByAzureSubscriptionID

> DescribeAccountConfigByAzureSubscriptionIDResult AccountConfigApiDescribeAccountConfigByAzureSubscriptionID(ctx, azureSubscriptionID).Execute()

DescribeAccountConfigByAzureSubscriptionID account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	azureSubscriptionID := "12345678-1234-1234-1234-123456789012" // string | The Azure subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByAzureSubscriptionID(context.Background(), azureSubscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByAzureSubscriptionID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiDescribeAccountConfigByAzureSubscriptionID`: DescribeAccountConfigByAzureSubscriptionIDResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByAzureSubscriptionID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureSubscriptionID** | **string** | The Azure subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiDescribeAccountConfigByAzureSubscriptionIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeAccountConfigByAzureSubscriptionIDResult**](DescribeAccountConfigByAzureSubscriptionIDResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiDescribeAccountConfigByGCPProjectID

> DescribeAccountConfigByGCPProjectIDResult AccountConfigApiDescribeAccountConfigByGCPProjectID(ctx, gcpProjectID).Execute()

DescribeAccountConfigByGCPProjectID account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	gcpProjectID := "my-project" // string | The GCP project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByGCPProjectID(context.Background(), gcpProjectID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByGCPProjectID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiDescribeAccountConfigByGCPProjectID`: DescribeAccountConfigByGCPProjectIDResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByGCPProjectID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gcpProjectID** | **string** | The GCP project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiDescribeAccountConfigByGCPProjectIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeAccountConfigByGCPProjectIDResult**](DescribeAccountConfigByGCPProjectIDResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiDescribeAccountConfigByNebiusTenantID

> DescribeAccountConfigByNebiusTenantIDResult AccountConfigApiDescribeAccountConfigByNebiusTenantID(ctx, nebiusTenantID).Execute()

DescribeAccountConfigByNebiusTenantID account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	nebiusTenantID := "tenant-id" // string | The Nebius tenant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByNebiusTenantID(context.Background(), nebiusTenantID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByNebiusTenantID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiDescribeAccountConfigByNebiusTenantID`: DescribeAccountConfigByNebiusTenantIDResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByNebiusTenantID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nebiusTenantID** | **string** | The Nebius tenant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiDescribeAccountConfigByNebiusTenantIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeAccountConfigByNebiusTenantIDResult**](DescribeAccountConfigByNebiusTenantIDResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiDescribeAccountConfigByOCITenancyID

> DescribeAccountConfigByOCITenancyIDResult AccountConfigApiDescribeAccountConfigByOCITenancyID(ctx, ociTenancyID).Execute()

DescribeAccountConfigByOCITenancyID account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	ociTenancyID := "ocid1.tenancy.oc1..aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" // string | The Tenancy OCID for Oracle Cloud Infrastructure

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByOCITenancyID(context.Background(), ociTenancyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByOCITenancyID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiDescribeAccountConfigByOCITenancyID`: DescribeAccountConfigByOCITenancyIDResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiDescribeAccountConfigByOCITenancyID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ociTenancyID** | **string** | The Tenancy OCID for Oracle Cloud Infrastructure | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiDescribeAccountConfigByOCITenancyIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DescribeAccountConfigByOCITenancyIDResult**](DescribeAccountConfigByOCITenancyIDResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiImportAccountConfigCloudNativeNetwork

> ListAccountConfigCloudNativeNetworksResult AccountConfigApiImportAccountConfigCloudNativeNetwork(ctx, id, cloudNativeNetworkId).Execute()

ImportAccountConfigCloudNativeNetwork account-config-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "ac-12345678" // string | Account Config ID to operate on
	cloudNativeNetworkId := "vpc-0abc123def456" // string | The cloud provider network ID (e.g. AWS VPC ID) to import for deployments

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiImportAccountConfigCloudNativeNetwork(context.Background(), id, cloudNativeNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiImportAccountConfigCloudNativeNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiImportAccountConfigCloudNativeNetwork`: ListAccountConfigCloudNativeNetworksResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiImportAccountConfigCloudNativeNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account Config ID to operate on | 
**cloudNativeNetworkId** | **string** | The cloud provider network ID (e.g. AWS VPC ID) to import for deployments | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiImportAccountConfigCloudNativeNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListAccountConfigCloudNativeNetworksResult**](ListAccountConfigCloudNativeNetworksResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiListAccountConfig

> ListAccountConfigResult AccountConfigApiListAccountConfig(ctx, cloudProviderName).Execute()

ListAccountConfig account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	cloudProviderName := "aws" // string | Cloud Provider name to filter on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiListAccountConfig(context.Background(), cloudProviderName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiListAccountConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiListAccountConfig`: ListAccountConfigResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiListAccountConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderName** | **string** | Cloud Provider name to filter on | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiListAccountConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAccountConfigResult**](ListAccountConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiListAccountConfigCloudNativeNetworks

> ListAccountConfigCloudNativeNetworksResult AccountConfigApiListAccountConfigCloudNativeNetworks(ctx, id).Execute()

ListAccountConfigCloudNativeNetworks account-config-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "ac-12345678" // string | Account Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiListAccountConfigCloudNativeNetworks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiListAccountConfigCloudNativeNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiListAccountConfigCloudNativeNetworks`: ListAccountConfigCloudNativeNetworksResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiListAccountConfigCloudNativeNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiListAccountConfigCloudNativeNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAccountConfigCloudNativeNetworksResult**](ListAccountConfigCloudNativeNetworksResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiListBYOAConfig

> ListBYOAConfigResult AccountConfigApiListBYOAConfig(ctx).ListBYOAConfigRequest2(listBYOAConfigRequest2).Execute()

ListBYOAConfig account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	listBYOAConfigRequest2 := *openapiclient.NewListBYOAConfigRequest2("aws|azure|gcp|nebius|oci|byoc-onprem|all") // ListBYOAConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiListBYOAConfig(context.Background()).ListBYOAConfigRequest2(listBYOAConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiListBYOAConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiListBYOAConfig`: ListBYOAConfigResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiListBYOAConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiListBYOAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listBYOAConfigRequest2** | [**ListBYOAConfigRequest2**](ListBYOAConfigRequest2.md) |  | 

### Return type

[**ListBYOAConfigResult**](ListBYOAConfigResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiSyncAccountConfigCloudNativeNetworks

> ListAccountConfigCloudNativeNetworksResult AccountConfigApiSyncAccountConfigCloudNativeNetworks(ctx, id).SyncAccountConfigCloudNativeNetworksRequest2(syncAccountConfigCloudNativeNetworksRequest2).Execute()

SyncAccountConfigCloudNativeNetworks account-config-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "ac-12345678" // string | Account Config ID to operate on
	syncAccountConfigCloudNativeNetworksRequest2 := *openapiclient.NewSyncAccountConfigCloudNativeNetworksRequest2() // SyncAccountConfigCloudNativeNetworksRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiSyncAccountConfigCloudNativeNetworks(context.Background(), id).SyncAccountConfigCloudNativeNetworksRequest2(syncAccountConfigCloudNativeNetworksRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiSyncAccountConfigCloudNativeNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiSyncAccountConfigCloudNativeNetworks`: ListAccountConfigCloudNativeNetworksResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiSyncAccountConfigCloudNativeNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiSyncAccountConfigCloudNativeNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncAccountConfigCloudNativeNetworksRequest2** | [**SyncAccountConfigCloudNativeNetworksRequest2**](SyncAccountConfigCloudNativeNetworksRequest2.md) |  | 

### Return type

[**ListAccountConfigCloudNativeNetworksResult**](ListAccountConfigCloudNativeNetworksResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiUnimportAccountConfigCloudNativeNetwork

> ListAccountConfigCloudNativeNetworksResult AccountConfigApiUnimportAccountConfigCloudNativeNetwork(ctx, id, cloudNativeNetworkId).Execute()

UnimportAccountConfigCloudNativeNetwork account-config-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "ac-12345678" // string | Account Config ID to operate on
	cloudNativeNetworkId := "vpc-0abc123def456" // string | The cloud provider network ID (e.g. AWS VPC ID) to unimport. Rejected with HTTP 400 if the network is currently in use by a host cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiUnimportAccountConfigCloudNativeNetwork(context.Background(), id, cloudNativeNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiUnimportAccountConfigCloudNativeNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiUnimportAccountConfigCloudNativeNetwork`: ListAccountConfigCloudNativeNetworksResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiUnimportAccountConfigCloudNativeNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account Config ID to operate on | 
**cloudNativeNetworkId** | **string** | The cloud provider network ID (e.g. AWS VPC ID) to unimport. Rejected with HTTP 400 if the network is currently in use by a host cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiUnimportAccountConfigCloudNativeNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListAccountConfigCloudNativeNetworksResult**](ListAccountConfigCloudNativeNetworksResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiUpdateAccountConfig

> string AccountConfigApiUpdateAccountConfig(ctx, id).UpdateAccountConfigRequest2(updateAccountConfigRequest2).Execute()

UpdateAccountConfig account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "ac-12345678" // string | Account Config ID to operate on
	updateAccountConfigRequest2 := *openapiclient.NewUpdateAccountConfigRequest2() // UpdateAccountConfigRequest2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiUpdateAccountConfig(context.Background(), id).UpdateAccountConfigRequest2(updateAccountConfigRequest2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiUpdateAccountConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiUpdateAccountConfig`: string
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiUpdateAccountConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiUpdateAccountConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccountConfigRequest2** | [**UpdateAccountConfigRequest2**](UpdateAccountConfigRequest2.md) |  | 

### Return type

**string**

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigApiVerifyAccountConfig

> AccountConfigApiVerifyAccountConfig(ctx, id).Execute()

VerifyAccountConfig account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

func main() {
	id := "ac-12345678" // string | Account Config ID to operate on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountConfigApiAPI.AccountConfigApiVerifyAccountConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiVerifyAccountConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account Config ID to operate on | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiVerifyAccountConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

