# \CostApiAPI

All URIs are relative to *https://api.omnistrate.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CostApiDescribeCloudProviderCost**](CostApiAPI.md#CostApiDescribeCloudProviderCost) | **Get** /2022-09-01-00/fleet/cost/cloud-provider | DescribeCloudProviderCost cost-api
[**CostApiDescribeDeploymentCellCost**](CostApiAPI.md#CostApiDescribeDeploymentCellCost) | **Get** /2022-09-01-00/fleet/cost/deployment-cell | DescribeDeploymentCellCost cost-api
[**CostApiDescribeRegionCost**](CostApiAPI.md#CostApiDescribeRegionCost) | **Get** /2022-09-01-00/fleet/cost/region | DescribeRegionCost cost-api
[**CostApiDescribeUserCost**](CostApiAPI.md#CostApiDescribeUserCost) | **Get** /2022-09-01-00/fleet/cost/user | DescribeUserCost cost-api



## CostApiDescribeCloudProviderCost

> DescribeCloudProviderCostResult CostApiDescribeCloudProviderCost(ctx).StartDate(startDate).EndDate(endDate).EnvironmentType(environmentType).Frequency(frequency).IncludeCloudProviderIDs(includeCloudProviderIDs).ExcludeCloudProviderIDs(excludeCloudProviderIDs).Execute()

DescribeCloudProviderCost cost-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	startDate := time.Now() // time.Time | The start date of the range
	endDate := time.Now() // time.Time | The end date of the range
	environmentType := "DEV" // string | The type of environment to filter costs by
	frequency := "DAILY|MONTHLY" // string | The frequency of the cost data
	includeCloudProviderIDs := "infra-12345678,infra-23456789" // string | The cloud provider IDs to include in the cost data (optional)
	excludeCloudProviderIDs := "infra-12345678,infra-23456789" // string | The cloud provider IDs to exclude from the cost data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostApiAPI.CostApiDescribeCloudProviderCost(context.Background()).StartDate(startDate).EndDate(endDate).EnvironmentType(environmentType).Frequency(frequency).IncludeCloudProviderIDs(includeCloudProviderIDs).ExcludeCloudProviderIDs(excludeCloudProviderIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostApiAPI.CostApiDescribeCloudProviderCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostApiDescribeCloudProviderCost`: DescribeCloudProviderCostResult
	fmt.Fprintf(os.Stdout, "Response from `CostApiAPI.CostApiDescribeCloudProviderCost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostApiDescribeCloudProviderCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | The start date of the range | 
 **endDate** | **time.Time** | The end date of the range | 
 **environmentType** | **string** | The type of environment to filter costs by | 
 **frequency** | **string** | The frequency of the cost data | 
 **includeCloudProviderIDs** | **string** | The cloud provider IDs to include in the cost data | 
 **excludeCloudProviderIDs** | **string** | The cloud provider IDs to exclude from the cost data | 

### Return type

[**DescribeCloudProviderCostResult**](DescribeCloudProviderCostResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostApiDescribeDeploymentCellCost

> DescribeDeploymentCellCostResult CostApiDescribeDeploymentCellCost(ctx).StartDate(startDate).EndDate(endDate).EnvironmentType(environmentType).Frequency(frequency).IncludeCloudProviderIDs(includeCloudProviderIDs).ExcludeCloudProviderIDs(excludeCloudProviderIDs).IncludeRegionIDs(includeRegionIDs).ExcludeRegionIDs(excludeRegionIDs).IncludeDeploymentCellIDs(includeDeploymentCellIDs).ExcludeDeploymentCellIDs(excludeDeploymentCellIDs).IncludeInstanceIDs(includeInstanceIDs).ExcludeInstanceIDs(excludeInstanceIDs).TopNInstances(topNInstances).Execute()

DescribeDeploymentCellCost cost-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	startDate := time.Now() // time.Time | The start date of the range
	endDate := time.Now() // time.Time | The end date of the range
	environmentType := "DEV" // string | The type of environment to filter costs by
	frequency := "DAILY|MONTHLY" // string | The frequency of the cost data
	includeCloudProviderIDs := "infra-12345678,infra-23456789" // string | The cloud provider IDs to include in the cost data (optional)
	excludeCloudProviderIDs := "infra-12345678,infra-23456789" // string | The cloud provider IDs to exclude from the cost data (optional)
	includeRegionIDs := "region-12345678,region-23456789" // string | The region IDs to include in the cost data (optional)
	excludeRegionIDs := "region-12345678,region-23456789" // string | The region IDs to exclude from the cost data (optional)
	includeDeploymentCellIDs := "hc-12345678,hc-23456789" // string | The deployment cell IDs to include in the cost data (optional)
	excludeDeploymentCellIDs := "hc-12345678,hc-23456789" // string | The deployment cell IDs to exclude from the cost data (optional)
	includeInstanceIDs := "instance-12345678,instance-23456789" // string | The instance IDs to include in the cost data (optional)
	excludeInstanceIDs := "instance-12345678,instance-23456789" // string | The instance IDs to exclude from the cost data (optional)
	topNInstances := int64(10) // int64 | The number of top instances to include in the cost data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostApiAPI.CostApiDescribeDeploymentCellCost(context.Background()).StartDate(startDate).EndDate(endDate).EnvironmentType(environmentType).Frequency(frequency).IncludeCloudProviderIDs(includeCloudProviderIDs).ExcludeCloudProviderIDs(excludeCloudProviderIDs).IncludeRegionIDs(includeRegionIDs).ExcludeRegionIDs(excludeRegionIDs).IncludeDeploymentCellIDs(includeDeploymentCellIDs).ExcludeDeploymentCellIDs(excludeDeploymentCellIDs).IncludeInstanceIDs(includeInstanceIDs).ExcludeInstanceIDs(excludeInstanceIDs).TopNInstances(topNInstances).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostApiAPI.CostApiDescribeDeploymentCellCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostApiDescribeDeploymentCellCost`: DescribeDeploymentCellCostResult
	fmt.Fprintf(os.Stdout, "Response from `CostApiAPI.CostApiDescribeDeploymentCellCost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostApiDescribeDeploymentCellCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | The start date of the range | 
 **endDate** | **time.Time** | The end date of the range | 
 **environmentType** | **string** | The type of environment to filter costs by | 
 **frequency** | **string** | The frequency of the cost data | 
 **includeCloudProviderIDs** | **string** | The cloud provider IDs to include in the cost data | 
 **excludeCloudProviderIDs** | **string** | The cloud provider IDs to exclude from the cost data | 
 **includeRegionIDs** | **string** | The region IDs to include in the cost data | 
 **excludeRegionIDs** | **string** | The region IDs to exclude from the cost data | 
 **includeDeploymentCellIDs** | **string** | The deployment cell IDs to include in the cost data | 
 **excludeDeploymentCellIDs** | **string** | The deployment cell IDs to exclude from the cost data | 
 **includeInstanceIDs** | **string** | The instance IDs to include in the cost data | 
 **excludeInstanceIDs** | **string** | The instance IDs to exclude from the cost data | 
 **topNInstances** | **int64** | The number of top instances to include in the cost data | 

### Return type

[**DescribeDeploymentCellCostResult**](DescribeDeploymentCellCostResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostApiDescribeRegionCost

> DescribeRegionCostResult CostApiDescribeRegionCost(ctx).StartDate(startDate).EndDate(endDate).EnvironmentType(environmentType).Frequency(frequency).IncludeCloudProviderIDs(includeCloudProviderIDs).ExcludeCloudProviderIDs(excludeCloudProviderIDs).IncludeRegionIDs(includeRegionIDs).ExcludeRegionIDs(excludeRegionIDs).Execute()

DescribeRegionCost cost-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	startDate := time.Now() // time.Time | The start date of the range
	endDate := time.Now() // time.Time | The end date of the range
	environmentType := "DEV" // string | The type of environment to filter costs by
	frequency := "DAILY|MONTHLY" // string | The frequency of the cost data
	includeCloudProviderIDs := "infra-12345678,infra-23456789" // string | The cloud provider IDs to include in the cost data (optional)
	excludeCloudProviderIDs := "infra-12345678,infra-23456789" // string | The cloud provider IDs to exclude from the cost data (optional)
	includeRegionIDs := "region-12345678,region-23456789" // string | The region IDs to include in the cost data (optional)
	excludeRegionIDs := "region-12345678,region-23456789" // string | The region IDs to exclude from the cost data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostApiAPI.CostApiDescribeRegionCost(context.Background()).StartDate(startDate).EndDate(endDate).EnvironmentType(environmentType).Frequency(frequency).IncludeCloudProviderIDs(includeCloudProviderIDs).ExcludeCloudProviderIDs(excludeCloudProviderIDs).IncludeRegionIDs(includeRegionIDs).ExcludeRegionIDs(excludeRegionIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostApiAPI.CostApiDescribeRegionCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostApiDescribeRegionCost`: DescribeRegionCostResult
	fmt.Fprintf(os.Stdout, "Response from `CostApiAPI.CostApiDescribeRegionCost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostApiDescribeRegionCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | The start date of the range | 
 **endDate** | **time.Time** | The end date of the range | 
 **environmentType** | **string** | The type of environment to filter costs by | 
 **frequency** | **string** | The frequency of the cost data | 
 **includeCloudProviderIDs** | **string** | The cloud provider IDs to include in the cost data | 
 **excludeCloudProviderIDs** | **string** | The cloud provider IDs to exclude from the cost data | 
 **includeRegionIDs** | **string** | The region IDs to include in the cost data | 
 **excludeRegionIDs** | **string** | The region IDs to exclude from the cost data | 

### Return type

[**DescribeRegionCostResult**](DescribeRegionCostResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostApiDescribeUserCost

> DescribeUserCostResult CostApiDescribeUserCost(ctx).StartDate(startDate).EndDate(endDate).EnvironmentType(environmentType).IncludeUserIDs(includeUserIDs).ExcludeUserIDs(excludeUserIDs).TopNUsers(topNUsers).TopNInstances(topNInstances).Execute()

DescribeUserCost cost-api



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/omnistrate-oss/omnistrate-sdk-go/fleet"
)

func main() {
	startDate := time.Now() // time.Time | The start date of the range
	endDate := time.Now() // time.Time | The end date of the range
	environmentType := "DEV" // string | The type of environment to filter costs by
	includeUserIDs := "user-12345678,user-23456789" // string | The user IDs to include in the cost data (optional)
	excludeUserIDs := "user-23456789" // string | The user IDs to exclude from the cost data (optional)
	topNUsers := int64(10) // int64 | The number of top users to include in the cost data (optional)
	topNInstances := int64(10) // int64 | The number of top instances to include in the cost data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostApiAPI.CostApiDescribeUserCost(context.Background()).StartDate(startDate).EndDate(endDate).EnvironmentType(environmentType).IncludeUserIDs(includeUserIDs).ExcludeUserIDs(excludeUserIDs).TopNUsers(topNUsers).TopNInstances(topNInstances).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostApiAPI.CostApiDescribeUserCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostApiDescribeUserCost`: DescribeUserCostResult
	fmt.Fprintf(os.Stdout, "Response from `CostApiAPI.CostApiDescribeUserCost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostApiDescribeUserCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | The start date of the range | 
 **endDate** | **time.Time** | The end date of the range | 
 **environmentType** | **string** | The type of environment to filter costs by | 
 **includeUserIDs** | **string** | The user IDs to include in the cost data | 
 **excludeUserIDs** | **string** | The user IDs to exclude from the cost data | 
 **topNUsers** | **int64** | The number of top users to include in the cost data | 
 **topNInstances** | **int64** | The number of top instances to include in the cost data | 

### Return type

[**DescribeUserCostResult**](DescribeUserCostResult.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

