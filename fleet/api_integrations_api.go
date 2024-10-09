/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type IntegrationsApiAPI interface {

	/*
	IntegrationsApiDisableProductTierIntegration DisableProductTierIntegration integrations-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId Service ID
	@param id Product tier ID
	@return ApiIntegrationsApiDisableProductTierIntegrationRequest
	*/
	IntegrationsApiDisableProductTierIntegration(ctx context.Context, serviceId string, id string) ApiIntegrationsApiDisableProductTierIntegrationRequest

	// IntegrationsApiDisableProductTierIntegrationExecute executes the request
	IntegrationsApiDisableProductTierIntegrationExecute(r ApiIntegrationsApiDisableProductTierIntegrationRequest) (*http.Response, error)

	/*
	IntegrationsApiEnableProductTierIntegration EnableProductTierIntegration integrations-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId Service ID
	@param id Product tier ID
	@return ApiIntegrationsApiEnableProductTierIntegrationRequest
	*/
	IntegrationsApiEnableProductTierIntegration(ctx context.Context, serviceId string, id string) ApiIntegrationsApiEnableProductTierIntegrationRequest

	// IntegrationsApiEnableProductTierIntegrationExecute executes the request
	IntegrationsApiEnableProductTierIntegrationExecute(r ApiIntegrationsApiEnableProductTierIntegrationRequest) (*http.Response, error)

	/*
	IntegrationsApiListProductTierIntegrations ListProductTierIntegrations integrations-api

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId Service ID
	@param id Product tier ID
	@return ApiIntegrationsApiListProductTierIntegrationsRequest
	*/
	IntegrationsApiListProductTierIntegrations(ctx context.Context, serviceId string, id string) ApiIntegrationsApiListProductTierIntegrationsRequest

	// IntegrationsApiListProductTierIntegrationsExecute executes the request
	//  @return ListProductTierIntegrationsResult
	IntegrationsApiListProductTierIntegrationsExecute(r ApiIntegrationsApiListProductTierIntegrationsRequest) (*ListProductTierIntegrationsResult, *http.Response, error)
}

// IntegrationsApiAPIService IntegrationsApiAPI service
type IntegrationsApiAPIService service

type ApiIntegrationsApiDisableProductTierIntegrationRequest struct {
	ctx context.Context
	ApiService IntegrationsApiAPI
	serviceId string
	id string
	disableProductTierIntegrationRequestBody *DisableProductTierIntegrationRequestBody
}

func (r ApiIntegrationsApiDisableProductTierIntegrationRequest) DisableProductTierIntegrationRequestBody(disableProductTierIntegrationRequestBody DisableProductTierIntegrationRequestBody) ApiIntegrationsApiDisableProductTierIntegrationRequest {
	r.disableProductTierIntegrationRequestBody = &disableProductTierIntegrationRequestBody
	return r
}

func (r ApiIntegrationsApiDisableProductTierIntegrationRequest) Execute() (*http.Response, error) {
	return r.ApiService.IntegrationsApiDisableProductTierIntegrationExecute(r)
}

/*
IntegrationsApiDisableProductTierIntegration DisableProductTierIntegration integrations-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Service ID
 @param id Product tier ID
 @return ApiIntegrationsApiDisableProductTierIntegrationRequest
*/
func (a *IntegrationsApiAPIService) IntegrationsApiDisableProductTierIntegration(ctx context.Context, serviceId string, id string) ApiIntegrationsApiDisableProductTierIntegrationRequest {
	return ApiIntegrationsApiDisableProductTierIntegrationRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
		id: id,
	}
}

// Execute executes the request
func (a *IntegrationsApiAPIService) IntegrationsApiDisableProductTierIntegrationExecute(r ApiIntegrationsApiDisableProductTierIntegrationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsApiAPIService.IntegrationsApiDisableProductTierIntegration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/fleet/service/{serviceId}/product-tier/{id}/integration"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.disableProductTierIntegrationRequestBody == nil {
		return nil, reportError("disableProductTierIntegrationRequestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.disableProductTierIntegrationRequestBody
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
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 409 {
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

type ApiIntegrationsApiEnableProductTierIntegrationRequest struct {
	ctx context.Context
	ApiService IntegrationsApiAPI
	serviceId string
	id string
	enableProductTierIntegrationRequestBody *EnableProductTierIntegrationRequestBody
}

func (r ApiIntegrationsApiEnableProductTierIntegrationRequest) EnableProductTierIntegrationRequestBody(enableProductTierIntegrationRequestBody EnableProductTierIntegrationRequestBody) ApiIntegrationsApiEnableProductTierIntegrationRequest {
	r.enableProductTierIntegrationRequestBody = &enableProductTierIntegrationRequestBody
	return r
}

func (r ApiIntegrationsApiEnableProductTierIntegrationRequest) Execute() (*http.Response, error) {
	return r.ApiService.IntegrationsApiEnableProductTierIntegrationExecute(r)
}

/*
IntegrationsApiEnableProductTierIntegration EnableProductTierIntegration integrations-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Service ID
 @param id Product tier ID
 @return ApiIntegrationsApiEnableProductTierIntegrationRequest
*/
func (a *IntegrationsApiAPIService) IntegrationsApiEnableProductTierIntegration(ctx context.Context, serviceId string, id string) ApiIntegrationsApiEnableProductTierIntegrationRequest {
	return ApiIntegrationsApiEnableProductTierIntegrationRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
		id: id,
	}
}

// Execute executes the request
func (a *IntegrationsApiAPIService) IntegrationsApiEnableProductTierIntegrationExecute(r ApiIntegrationsApiEnableProductTierIntegrationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsApiAPIService.IntegrationsApiEnableProductTierIntegration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/fleet/service/{serviceId}/product-tier/{id}/integration"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.enableProductTierIntegrationRequestBody == nil {
		return nil, reportError("enableProductTierIntegrationRequestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.enableProductTierIntegrationRequestBody
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
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 409 {
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

type ApiIntegrationsApiListProductTierIntegrationsRequest struct {
	ctx context.Context
	ApiService IntegrationsApiAPI
	serviceId string
	id string
}

func (r ApiIntegrationsApiListProductTierIntegrationsRequest) Execute() (*ListProductTierIntegrationsResult, *http.Response, error) {
	return r.ApiService.IntegrationsApiListProductTierIntegrationsExecute(r)
}

/*
IntegrationsApiListProductTierIntegrations ListProductTierIntegrations integrations-api

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Service ID
 @param id Product tier ID
 @return ApiIntegrationsApiListProductTierIntegrationsRequest
*/
func (a *IntegrationsApiAPIService) IntegrationsApiListProductTierIntegrations(ctx context.Context, serviceId string, id string) ApiIntegrationsApiListProductTierIntegrationsRequest {
	return ApiIntegrationsApiListProductTierIntegrationsRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
		id: id,
	}
}

// Execute executes the request
//  @return ListProductTierIntegrationsResult
func (a *IntegrationsApiAPIService) IntegrationsApiListProductTierIntegrationsExecute(r ApiIntegrationsApiListProductTierIntegrationsRequest) (*ListProductTierIntegrationsResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListProductTierIntegrationsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsApiAPIService.IntegrationsApiListProductTierIntegrations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/2022-09-01-00/fleet/service/{serviceId}/product-tier/{id}/integration"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
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
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 409 {
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