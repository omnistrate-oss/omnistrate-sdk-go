/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type CustomNetworkApiAPI interface {

	/*
	CustomNetworkApiCreateCustomNetwork CreateCustomNetwork custom-network-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCustomNetworkApiCreateCustomNetworkRequest
	*/
	CustomNetworkApiCreateCustomNetwork(ctx context.Context) ApiCustomNetworkApiCreateCustomNetworkRequest

	// CustomNetworkApiCreateCustomNetworkExecute executes the request
	//  @return CustomNetwork
	CustomNetworkApiCreateCustomNetworkExecute(r ApiCustomNetworkApiCreateCustomNetworkRequest) (*CustomNetwork, *http.Response, error)

	/*
	CustomNetworkApiDeleteCustomNetwork DeleteCustomNetwork custom-network-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of a custom network
	@return ApiCustomNetworkApiDeleteCustomNetworkRequest
	*/
	CustomNetworkApiDeleteCustomNetwork(ctx context.Context, id string) ApiCustomNetworkApiDeleteCustomNetworkRequest

	// CustomNetworkApiDeleteCustomNetworkExecute executes the request
	CustomNetworkApiDeleteCustomNetworkExecute(r ApiCustomNetworkApiDeleteCustomNetworkRequest) (*http.Response, error)

	/*
	CustomNetworkApiDescribeCustomNetwork DescribeCustomNetwork custom-network-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of a custom network
	@return ApiCustomNetworkApiDescribeCustomNetworkRequest
	*/
	CustomNetworkApiDescribeCustomNetwork(ctx context.Context, id string) ApiCustomNetworkApiDescribeCustomNetworkRequest

	// CustomNetworkApiDescribeCustomNetworkExecute executes the request
	//  @return CustomNetwork
	CustomNetworkApiDescribeCustomNetworkExecute(r ApiCustomNetworkApiDescribeCustomNetworkRequest) (*CustomNetwork, *http.Response, error)

	/*
	CustomNetworkApiListCustomNetworks ListCustomNetworks custom-network-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCustomNetworkApiListCustomNetworksRequest
	*/
	CustomNetworkApiListCustomNetworks(ctx context.Context) ApiCustomNetworkApiListCustomNetworksRequest

	// CustomNetworkApiListCustomNetworksExecute executes the request
	//  @return ListCustomNetworksResult
	CustomNetworkApiListCustomNetworksExecute(r ApiCustomNetworkApiListCustomNetworksRequest) (*ListCustomNetworksResult, *http.Response, error)

	/*
	CustomNetworkApiUpdateCustomNetwork UpdateCustomNetwork custom-network-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of a custom network
	@return ApiCustomNetworkApiUpdateCustomNetworkRequest
	*/
	CustomNetworkApiUpdateCustomNetwork(ctx context.Context, id string) ApiCustomNetworkApiUpdateCustomNetworkRequest

	// CustomNetworkApiUpdateCustomNetworkExecute executes the request
	//  @return CustomNetwork
	CustomNetworkApiUpdateCustomNetworkExecute(r ApiCustomNetworkApiUpdateCustomNetworkRequest) (*CustomNetwork, *http.Response, error)
}

// CustomNetworkApiAPIService CustomNetworkApiAPI service
type CustomNetworkApiAPIService service

type ApiCustomNetworkApiCreateCustomNetworkRequest struct {
	ctx context.Context
	ApiService CustomNetworkApiAPI
	createCustomNetworkRequestBody *CreateCustomNetworkRequestBody
}

func (r ApiCustomNetworkApiCreateCustomNetworkRequest) CreateCustomNetworkRequestBody(createCustomNetworkRequestBody CreateCustomNetworkRequestBody) ApiCustomNetworkApiCreateCustomNetworkRequest {
	r.createCustomNetworkRequestBody = &createCustomNetworkRequestBody
	return r
}

func (r ApiCustomNetworkApiCreateCustomNetworkRequest) Execute() (*CustomNetwork, *http.Response, error) {
	return r.ApiService.CustomNetworkApiCreateCustomNetworkExecute(r)
}

/*
CustomNetworkApiCreateCustomNetwork CreateCustomNetwork custom-network-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCustomNetworkApiCreateCustomNetworkRequest
*/
func (a *CustomNetworkApiAPIService) CustomNetworkApiCreateCustomNetwork(ctx context.Context) ApiCustomNetworkApiCreateCustomNetworkRequest {
	return ApiCustomNetworkApiCreateCustomNetworkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomNetwork
func (a *CustomNetworkApiAPIService) CustomNetworkApiCreateCustomNetworkExecute(r ApiCustomNetworkApiCreateCustomNetworkRequest) (*CustomNetwork, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomNetworkApiAPIService.CustomNetworkApiCreateCustomNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/resource-instance/custom-network"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createCustomNetworkRequestBody == nil {
		return localVarReturnValue, nil, reportError("createCustomNetworkRequestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.goa.error"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createCustomNetworkRequestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCustomNetworkApiDeleteCustomNetworkRequest struct {
	ctx context.Context
	ApiService CustomNetworkApiAPI
	id string
}

func (r ApiCustomNetworkApiDeleteCustomNetworkRequest) Execute() (*http.Response, error) {
	return r.ApiService.CustomNetworkApiDeleteCustomNetworkExecute(r)
}

/*
CustomNetworkApiDeleteCustomNetwork DeleteCustomNetwork custom-network-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of a custom network
 @return ApiCustomNetworkApiDeleteCustomNetworkRequest
*/
func (a *CustomNetworkApiAPIService) CustomNetworkApiDeleteCustomNetwork(ctx context.Context, id string) ApiCustomNetworkApiDeleteCustomNetworkRequest {
	return ApiCustomNetworkApiDeleteCustomNetworkRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *CustomNetworkApiAPIService) CustomNetworkApiDeleteCustomNetworkExecute(r ApiCustomNetworkApiDeleteCustomNetworkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomNetworkApiAPIService.CustomNetworkApiDeleteCustomNetwork")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/resource-instance/custom-network/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.goa.error"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCustomNetworkApiDescribeCustomNetworkRequest struct {
	ctx context.Context
	ApiService CustomNetworkApiAPI
	id string
}

func (r ApiCustomNetworkApiDescribeCustomNetworkRequest) Execute() (*CustomNetwork, *http.Response, error) {
	return r.ApiService.CustomNetworkApiDescribeCustomNetworkExecute(r)
}

/*
CustomNetworkApiDescribeCustomNetwork DescribeCustomNetwork custom-network-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of a custom network
 @return ApiCustomNetworkApiDescribeCustomNetworkRequest
*/
func (a *CustomNetworkApiAPIService) CustomNetworkApiDescribeCustomNetwork(ctx context.Context, id string) ApiCustomNetworkApiDescribeCustomNetworkRequest {
	return ApiCustomNetworkApiDescribeCustomNetworkRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CustomNetwork
func (a *CustomNetworkApiAPIService) CustomNetworkApiDescribeCustomNetworkExecute(r ApiCustomNetworkApiDescribeCustomNetworkRequest) (*CustomNetwork, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomNetworkApiAPIService.CustomNetworkApiDescribeCustomNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/resource-instance/custom-network/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.goa.error"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCustomNetworkApiListCustomNetworksRequest struct {
	ctx context.Context
	ApiService CustomNetworkApiAPI
	cloudProviderName *string
	cloudProviderRegion *string
	customNetworksOnly *bool
}

// The name of the cloud provider that custom network should be created in
func (r ApiCustomNetworkApiListCustomNetworksRequest) CloudProviderName(cloudProviderName string) ApiCustomNetworkApiListCustomNetworksRequest {
	r.cloudProviderName = &cloudProviderName
	return r
}

// The region of the cloud provider that the network should be created in
func (r ApiCustomNetworkApiListCustomNetworksRequest) CloudProviderRegion(cloudProviderRegion string) ApiCustomNetworkApiListCustomNetworksRequest {
	r.cloudProviderRegion = &cloudProviderRegion
	return r
}

// Flag indicating whether to return only custom networks, or to include default and imported networks as well
func (r ApiCustomNetworkApiListCustomNetworksRequest) CustomNetworksOnly(customNetworksOnly bool) ApiCustomNetworkApiListCustomNetworksRequest {
	r.customNetworksOnly = &customNetworksOnly
	return r
}

func (r ApiCustomNetworkApiListCustomNetworksRequest) Execute() (*ListCustomNetworksResult, *http.Response, error) {
	return r.ApiService.CustomNetworkApiListCustomNetworksExecute(r)
}

/*
CustomNetworkApiListCustomNetworks ListCustomNetworks custom-network-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCustomNetworkApiListCustomNetworksRequest
*/
func (a *CustomNetworkApiAPIService) CustomNetworkApiListCustomNetworks(ctx context.Context) ApiCustomNetworkApiListCustomNetworksRequest {
	return ApiCustomNetworkApiListCustomNetworksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListCustomNetworksResult
func (a *CustomNetworkApiAPIService) CustomNetworkApiListCustomNetworksExecute(r ApiCustomNetworkApiListCustomNetworksRequest) (*ListCustomNetworksResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCustomNetworksResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomNetworkApiAPIService.CustomNetworkApiListCustomNetworks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/resource-instance/custom-network"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cloudProviderName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cloudProviderName", r.cloudProviderName, "form", "")
	}
	if r.cloudProviderRegion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cloudProviderRegion", r.cloudProviderRegion, "form", "")
	}
	if r.customNetworksOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customNetworksOnly", r.customNetworksOnly, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.goa.error"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCustomNetworkApiUpdateCustomNetworkRequest struct {
	ctx context.Context
	ApiService CustomNetworkApiAPI
	id string
	updateCustomNetworkRequestBody *UpdateCustomNetworkRequestBody
}

func (r ApiCustomNetworkApiUpdateCustomNetworkRequest) UpdateCustomNetworkRequestBody(updateCustomNetworkRequestBody UpdateCustomNetworkRequestBody) ApiCustomNetworkApiUpdateCustomNetworkRequest {
	r.updateCustomNetworkRequestBody = &updateCustomNetworkRequestBody
	return r
}

func (r ApiCustomNetworkApiUpdateCustomNetworkRequest) Execute() (*CustomNetwork, *http.Response, error) {
	return r.ApiService.CustomNetworkApiUpdateCustomNetworkExecute(r)
}

/*
CustomNetworkApiUpdateCustomNetwork UpdateCustomNetwork custom-network-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of a custom network
 @return ApiCustomNetworkApiUpdateCustomNetworkRequest
*/
func (a *CustomNetworkApiAPIService) CustomNetworkApiUpdateCustomNetwork(ctx context.Context, id string) ApiCustomNetworkApiUpdateCustomNetworkRequest {
	return ApiCustomNetworkApiUpdateCustomNetworkRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CustomNetwork
func (a *CustomNetworkApiAPIService) CustomNetworkApiUpdateCustomNetworkExecute(r ApiCustomNetworkApiUpdateCustomNetworkRequest) (*CustomNetwork, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomNetworkApiAPIService.CustomNetworkApiUpdateCustomNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/resource-instance/custom-network/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateCustomNetworkRequestBody == nil {
		return localVarReturnValue, nil, reportError("updateCustomNetworkRequestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/vnd.goa.error"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateCustomNetworkRequestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
