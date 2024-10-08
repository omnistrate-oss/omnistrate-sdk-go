# \AccountConfigApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountConfigApiAccountConfigIdentityID**](AccountConfigApiAPI.md#AccountConfigApiAccountConfigIdentityID) | **Get** /2022-09-01-00/accountconfig/identityid | AccountConfigIdentityID account-config-api
[**AccountConfigApiCreateAccountConfig**](AccountConfigApiAPI.md#AccountConfigApiCreateAccountConfig) | **Post** /2022-09-01-00/accountconfig | CreateAccountConfig account-config-api
[**AccountConfigApiDeleteAccountConfig**](AccountConfigApiAPI.md#AccountConfigApiDeleteAccountConfig) | **Delete** /2022-09-01-00/accountconfig/{id} | DeleteAccountConfig account-config-api
[**AccountConfigApiDescribeAccountConfig**](AccountConfigApiAPI.md#AccountConfigApiDescribeAccountConfig) | **Get** /2022-09-01-00/accountconfig/{id} | DescribeAccountConfig account-config-api
[**AccountConfigApiDescribeAccountConfigByAWSAccountID**](AccountConfigApiAPI.md#AccountConfigApiDescribeAccountConfigByAWSAccountID) | **Get** /2022-09-01-00/accountconfig/aws/{awsAccountID} | DescribeAccountConfigByAWSAccountID account-config-api
[**AccountConfigApiDescribeAccountConfigByGCPProjectID**](AccountConfigApiAPI.md#AccountConfigApiDescribeAccountConfigByGCPProjectID) | **Get** /2022-09-01-00/accountconfig/gcp/{gcpProjectID} | DescribeAccountConfigByGCPProjectID account-config-api
[**AccountConfigApiListAccountConfig**](AccountConfigApiAPI.md#AccountConfigApiListAccountConfig) | **Get** /2022-09-01-00/accountconfig/cloudprovider/{cloudProviderName} | ListAccountConfig account-config-api
[**AccountConfigApiListBYOAConfig**](AccountConfigApiAPI.md#AccountConfigApiListBYOAConfig) | **Get** /2022-09-01-00/accountconfig/byoa | ListBYOAConfig account-config-api
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
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


## AccountConfigApiCreateAccountConfig

> string AccountConfigApiCreateAccountConfig(ctx).CreateAccountConfigRequestBody(createAccountConfigRequestBody).Execute()

CreateAccountConfig account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	createAccountConfigRequestBody := *openapiclient.NewCreateAccountConfigRequestBody("infra-12345678", "An AWS account hosting multiple dev environments", "Dev AWS account") // CreateAccountConfigRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiCreateAccountConfig(context.Background()).CreateAccountConfigRequestBody(createAccountConfigRequestBody).Execute()
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
 **createAccountConfigRequestBody** | [**CreateAccountConfigRequestBody**](CreateAccountConfigRequestBody.md) |  | 

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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
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


## AccountConfigApiListBYOAConfig

> ListAccountConfigResult AccountConfigApiListBYOAConfig(ctx).ListBYOAConfigRequestBody(listBYOAConfigRequestBody).Execute()

ListBYOAConfig account-config-api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
)

func main() {
	listBYOAConfigRequestBody := *openapiclient.NewListBYOAConfigRequestBody("aws") // ListBYOAConfigRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountConfigApiAPI.AccountConfigApiListBYOAConfig(context.Background()).ListBYOAConfigRequestBody(listBYOAConfigRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountConfigApiAPI.AccountConfigApiListBYOAConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigApiListBYOAConfig`: ListAccountConfigResult
	fmt.Fprintf(os.Stdout, "Response from `AccountConfigApiAPI.AccountConfigApiListBYOAConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigApiListBYOAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listBYOAConfigRequestBody** | [**ListBYOAConfigRequestBody**](ListBYOAConfigRequestBody.md) |  | 

### Return type

[**ListAccountConfigResult**](ListAccountConfigResult.md)

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
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/registration"
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

