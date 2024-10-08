/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type GlobalApiAPI interface {

	/*
	GlobalApiConsumptionServiceHealth consumptionServiceHealth global-api

	Consumption service health check endpoint

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGlobalApiConsumptionServiceHealthRequest
	*/
	GlobalApiConsumptionServiceHealth(ctx context.Context) ApiGlobalApiConsumptionServiceHealthRequest

	// GlobalApiConsumptionServiceHealthExecute executes the request
	//  @return OmnistrateServiceHealthResult
	GlobalApiConsumptionServiceHealthExecute(r ApiGlobalApiConsumptionServiceHealthRequest) (*OmnistrateServiceHealthResult, *http.Response, error)

	/*
	GlobalApiConsumptionServiceVersion consumptionServiceVersion global-api

	Consumption service version endpoint

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGlobalApiConsumptionServiceVersionRequest
	*/
	GlobalApiConsumptionServiceVersion(ctx context.Context) ApiGlobalApiConsumptionServiceVersionRequest

	// GlobalApiConsumptionServiceVersionExecute executes the request
	//  @return OmnistrateServiceVersionResult
	GlobalApiConsumptionServiceVersionExecute(r ApiGlobalApiConsumptionServiceVersionRequest) (*OmnistrateServiceVersionResult, *http.Response, error)

	/*
	GlobalApiRegistrationServiceHealth registrationServiceHealth global-api

	Registration service health check endpoint

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGlobalApiRegistrationServiceHealthRequest
	*/
	GlobalApiRegistrationServiceHealth(ctx context.Context) ApiGlobalApiRegistrationServiceHealthRequest

	// GlobalApiRegistrationServiceHealthExecute executes the request
	//  @return OmnistrateServiceHealthResult
	GlobalApiRegistrationServiceHealthExecute(r ApiGlobalApiRegistrationServiceHealthRequest) (*OmnistrateServiceHealthResult, *http.Response, error)

	/*
	GlobalApiRegistrationServiceVersion registrationServiceVersion global-api

	Registration service version endpoint

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGlobalApiRegistrationServiceVersionRequest
	*/
	GlobalApiRegistrationServiceVersion(ctx context.Context) ApiGlobalApiRegistrationServiceVersionRequest

	// GlobalApiRegistrationServiceVersionExecute executes the request
	//  @return OmnistrateServiceVersionResult
	GlobalApiRegistrationServiceVersionExecute(r ApiGlobalApiRegistrationServiceVersionRequest) (*OmnistrateServiceVersionResult, *http.Response, error)
}

// GlobalApiAPIService GlobalApiAPI service
type GlobalApiAPIService service

type ApiGlobalApiConsumptionServiceHealthRequest struct {
	ctx context.Context
	ApiService GlobalApiAPI
}

func (r ApiGlobalApiConsumptionServiceHealthRequest) Execute() (*OmnistrateServiceHealthResult, *http.Response, error) {
	return r.ApiService.GlobalApiConsumptionServiceHealthExecute(r)
}

/*
GlobalApiConsumptionServiceHealth consumptionServiceHealth global-api

Consumption service health check endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGlobalApiConsumptionServiceHealthRequest
*/
func (a *GlobalApiAPIService) GlobalApiConsumptionServiceHealth(ctx context.Context) ApiGlobalApiConsumptionServiceHealthRequest {
	return ApiGlobalApiConsumptionServiceHealthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OmnistrateServiceHealthResult
func (a *GlobalApiAPIService) GlobalApiConsumptionServiceHealthExecute(r ApiGlobalApiConsumptionServiceHealthRequest) (*OmnistrateServiceHealthResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OmnistrateServiceHealthResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalApiAPIService.GlobalApiConsumptionServiceHealth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/resource-instance/health"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiGlobalApiConsumptionServiceVersionRequest struct {
	ctx context.Context
	ApiService GlobalApiAPI
}

func (r ApiGlobalApiConsumptionServiceVersionRequest) Execute() (*OmnistrateServiceVersionResult, *http.Response, error) {
	return r.ApiService.GlobalApiConsumptionServiceVersionExecute(r)
}

/*
GlobalApiConsumptionServiceVersion consumptionServiceVersion global-api

Consumption service version endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGlobalApiConsumptionServiceVersionRequest
*/
func (a *GlobalApiAPIService) GlobalApiConsumptionServiceVersion(ctx context.Context) ApiGlobalApiConsumptionServiceVersionRequest {
	return ApiGlobalApiConsumptionServiceVersionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OmnistrateServiceVersionResult
func (a *GlobalApiAPIService) GlobalApiConsumptionServiceVersionExecute(r ApiGlobalApiConsumptionServiceVersionRequest) (*OmnistrateServiceVersionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OmnistrateServiceVersionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalApiAPIService.GlobalApiConsumptionServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/resource-instance/version"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiGlobalApiRegistrationServiceHealthRequest struct {
	ctx context.Context
	ApiService GlobalApiAPI
}

func (r ApiGlobalApiRegistrationServiceHealthRequest) Execute() (*OmnistrateServiceHealthResult, *http.Response, error) {
	return r.ApiService.GlobalApiRegistrationServiceHealthExecute(r)
}

/*
GlobalApiRegistrationServiceHealth registrationServiceHealth global-api

Registration service health check endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGlobalApiRegistrationServiceHealthRequest
*/
func (a *GlobalApiAPIService) GlobalApiRegistrationServiceHealth(ctx context.Context) ApiGlobalApiRegistrationServiceHealthRequest {
	return ApiGlobalApiRegistrationServiceHealthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OmnistrateServiceHealthResult
func (a *GlobalApiAPIService) GlobalApiRegistrationServiceHealthExecute(r ApiGlobalApiRegistrationServiceHealthRequest) (*OmnistrateServiceHealthResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OmnistrateServiceHealthResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalApiAPIService.GlobalApiRegistrationServiceHealth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/health"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiGlobalApiRegistrationServiceVersionRequest struct {
	ctx context.Context
	ApiService GlobalApiAPI
}

func (r ApiGlobalApiRegistrationServiceVersionRequest) Execute() (*OmnistrateServiceVersionResult, *http.Response, error) {
	return r.ApiService.GlobalApiRegistrationServiceVersionExecute(r)
}

/*
GlobalApiRegistrationServiceVersion registrationServiceVersion global-api

Registration service version endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGlobalApiRegistrationServiceVersionRequest
*/
func (a *GlobalApiAPIService) GlobalApiRegistrationServiceVersion(ctx context.Context) ApiGlobalApiRegistrationServiceVersionRequest {
	return ApiGlobalApiRegistrationServiceVersionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OmnistrateServiceVersionResult
func (a *GlobalApiAPIService) GlobalApiRegistrationServiceVersionExecute(r ApiGlobalApiRegistrationServiceVersionRequest) (*OmnistrateServiceVersionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OmnistrateServiceVersionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalApiAPIService.GlobalApiRegistrationServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/version"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
