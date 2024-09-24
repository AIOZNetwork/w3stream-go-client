/*
 * W3STREAM API
 *
 * W3STREAM Service
 *
 * API version: 1.0
 * Contact: support@swagger.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package w3streamsdk

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type ApiKeyApiListRequest struct {
	search  *string
	sortBy  *string
	orderBy *string
	offset  *int32
	limit   *int32
}

func (r ApiKeyApiListRequest) Search(search string) ApiKeyApiListRequest {
	r.search = &search
	return r
}
func (r ApiKeyApiListRequest) SortBy(sortBy string) ApiKeyApiListRequest {
	r.sortBy = &sortBy
	return r
}
func (r ApiKeyApiListRequest) OrderBy(orderBy string) ApiKeyApiListRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiKeyApiListRequest) Offset(offset int32) ApiKeyApiListRequest {
	r.offset = &offset
	return r
}
func (r ApiKeyApiListRequest) Limit(limit int32) ApiKeyApiListRequest {
	r.limit = &limit
	return r
}

type ApiKeyServiceI interface {
	/*
	 * Create Create API key
	 * @return ApiKeyApiCreateRequest
	 */

	Create(request CreateApiKeyRequest) (*CreateApiKeyResponse, error)

	/*
	 * Create Create API key
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiKeyApiCreateRequest
	 */

	CreateWithContext(ctx context.Context, request CreateApiKeyRequest) (*CreateApiKeyResponse, error)

	/*
	 * Update Rename api key
	 * @param id api key id
	 * @return ApiKeyApiUpdateRequest
	 */

	Update(id string, request RenameAPIKeyRequest) (*ResponseSuccess, error)

	/*
	 * Update Rename api key
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id api key id
	 * @return ApiKeyApiUpdateRequest
	 */

	UpdateWithContext(ctx context.Context, id string, request RenameAPIKeyRequest) (*ResponseSuccess, error)

	/*
	 * Delete Delete API key
	 * @param id API key's ID
	 * @return ApiKeyApiDeleteRequest
	 */

	Delete(id string) (*ResponseSuccess, error)

	/*
	 * Delete Delete API key
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id API key's ID
	 * @return ApiKeyApiDeleteRequest
	 */

	DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)

	/*
	 * List Get list API keys
	 * @return ApiKeyApiListRequest
	 */

	List(r ApiKeyApiListRequest) (*GetApiKeysResponse, error)

	/*
	 * List Get list API keys
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiKeyApiListRequest
	 */

	ListWithContext(ctx context.Context, r ApiKeyApiListRequest) (*GetApiKeysResponse, error)
}

// ApiKeyService communicating with the ApiKey
// endpoints of the w3stream API
type ApiKeyService struct {
	client *Client
}

/*
 * Create Create API key
 * This endpoint enables you to create a new API key for a specific project.

 * @return ApiKeyApiCreateRequest
 */

func (s *ApiKeyService) Create(request CreateApiKeyRequest) (*CreateApiKeyResponse, error) {

	return s.CreateWithContext(context.Background(), request)

}

/*
 * Create Create API key
 * This endpoint enables you to create a new API key for a specific project.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiKeyApiCreateRequest
 */

func (s *ApiKeyService) CreateWithContext(ctx context.Context, request CreateApiKeyRequest) (*CreateApiKeyResponse, error) {
	var localVarPostBody interface{}

	localVarPath := "/api_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	// body params
	localVarPostBody = request

	req, err := s.client.prepareRequest(ctx, http.MethodPost, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(CreateApiKeyResponse)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}

/*
 * Update Rename api key
 * This endpoint enables you to rename an API key from a specific project.

 * @param id api key id
 * @return ApiKeyApiUpdateRequest
 */

func (s *ApiKeyService) Update(id string, request RenameAPIKeyRequest) (*ResponseSuccess, error) {

	return s.UpdateWithContext(context.Background(), id, request)

}

/*
 * Update Rename api key
 * This endpoint enables you to rename an API key from a specific project.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id api key id
 * @return ApiKeyApiUpdateRequest
 */

func (s *ApiKeyService) UpdateWithContext(ctx context.Context, id string, request RenameAPIKeyRequest) (*ResponseSuccess, error) {
	var localVarPostBody interface{}

	localVarPath := "/api_keys/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	// body params
	localVarPostBody = request

	req, err := s.client.prepareRequest(ctx, http.MethodPatch, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(ResponseSuccess)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}

/*
 * Delete Delete API key
 * This endpoint enables you to delete an API key from a specific project.

 * @param id API key's ID
 * @return ApiKeyApiDeleteRequest
 */

func (s *ApiKeyService) Delete(id string) (*ResponseSuccess, error) {

	return s.DeleteWithContext(context.Background(), id)

}

/*
 * Delete Delete API key
 * This endpoint enables you to delete an API key from a specific project.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id API key's ID
 * @return ApiKeyApiDeleteRequest
 */

func (s *ApiKeyService) DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error) {
	var localVarPostBody interface{}

	localVarPath := "/api_keys/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	req, err := s.client.prepareRequest(ctx, http.MethodDelete, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(ResponseSuccess)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}

/*
 * List Get list API keys
 * Retrieve a list of all API keys for the current workspace.

 * @return ApiKeyApiListRequest
 */

func (s *ApiKeyService) List(r ApiKeyApiListRequest) (*GetApiKeysResponse, error) {

	return s.ListWithContext(context.Background(), r)

}

/*
 * List Get list API keys
 * Retrieve a list of all API keys for the current workspace.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiKeyApiListRequest
 */

func (s *ApiKeyService) ListWithContext(ctx context.Context, r ApiKeyApiListRequest) (*GetApiKeysResponse, error) {
	var localVarPostBody interface{}

	localVarPath := "/api_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("order_by", parameterToString(*r.orderBy, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}

	req, err := s.client.prepareRequest(ctx, http.MethodGet, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return nil, err
	}

	res := new(GetApiKeysResponse)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}