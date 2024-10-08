# \GlobalApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalApiConsumptionServiceHealth**](GlobalApiAPI.md#GlobalApiConsumptionServiceHealth) | **Get** /2022-09-01-00/resource-instance/health | consumptionServiceHealth global-api
[**GlobalApiConsumptionServiceVersion**](GlobalApiAPI.md#GlobalApiConsumptionServiceVersion) | **Get** /2022-09-01-00/resource-instance/version | consumptionServiceVersion global-api
[**GlobalApiRegistrationServiceHealth**](GlobalApiAPI.md#GlobalApiRegistrationServiceHealth) | **Get** /2022-09-01-00/health | registrationServiceHealth global-api
[**GlobalApiRegistrationServiceVersion**](GlobalApiAPI.md#GlobalApiRegistrationServiceVersion) | **Get** /2022-09-01-00/version | registrationServiceVersion global-api



## GlobalApiConsumptionServiceHealth

> OmnistrateServiceHealthResult GlobalApiConsumptionServiceHealth(ctx).Execute()

consumptionServiceHealth global-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalApiAPI.GlobalApiConsumptionServiceHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalApiAPI.GlobalApiConsumptionServiceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalApiConsumptionServiceHealth`: OmnistrateServiceHealthResult
	fmt.Fprintf(os.Stdout, "Response from `GlobalApiAPI.GlobalApiConsumptionServiceHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalApiConsumptionServiceHealthRequest struct via the builder pattern


### Return type

[**OmnistrateServiceHealthResult**](OmnistrateServiceHealthResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalApiConsumptionServiceVersion

> OmnistrateServiceVersionResult GlobalApiConsumptionServiceVersion(ctx).Execute()

consumptionServiceVersion global-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalApiAPI.GlobalApiConsumptionServiceVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalApiAPI.GlobalApiConsumptionServiceVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalApiConsumptionServiceVersion`: OmnistrateServiceVersionResult
	fmt.Fprintf(os.Stdout, "Response from `GlobalApiAPI.GlobalApiConsumptionServiceVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalApiConsumptionServiceVersionRequest struct via the builder pattern


### Return type

[**OmnistrateServiceVersionResult**](OmnistrateServiceVersionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalApiRegistrationServiceHealth

> OmnistrateServiceHealthResult GlobalApiRegistrationServiceHealth(ctx).Execute()

registrationServiceHealth global-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalApiAPI.GlobalApiRegistrationServiceHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalApiAPI.GlobalApiRegistrationServiceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalApiRegistrationServiceHealth`: OmnistrateServiceHealthResult
	fmt.Fprintf(os.Stdout, "Response from `GlobalApiAPI.GlobalApiRegistrationServiceHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalApiRegistrationServiceHealthRequest struct via the builder pattern


### Return type

[**OmnistrateServiceHealthResult**](OmnistrateServiceHealthResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalApiRegistrationServiceVersion

> OmnistrateServiceVersionResult GlobalApiRegistrationServiceVersion(ctx).Execute()

registrationServiceVersion global-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/omnistrate/omnistrate-sdk-go/v1"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalApiAPI.GlobalApiRegistrationServiceVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalApiAPI.GlobalApiRegistrationServiceVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GlobalApiRegistrationServiceVersion`: OmnistrateServiceVersionResult
	fmt.Fprintf(os.Stdout, "Response from `GlobalApiAPI.GlobalApiRegistrationServiceVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalApiRegistrationServiceVersionRequest struct via the builder pattern


### Return type

[**OmnistrateServiceVersionResult**](OmnistrateServiceVersionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

