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
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type WatermarkApiListRequest struct {
	sortBy  *string
	orderBy *string
	offset  *int32
	limit   *int32
}

func (r WatermarkApiListRequest) SortBy(sortBy string) WatermarkApiListRequest {
	r.sortBy = &sortBy
	return r
}
func (r WatermarkApiListRequest) OrderBy(orderBy string) WatermarkApiListRequest {
	r.orderBy = &orderBy
	return r
}
func (r WatermarkApiListRequest) Offset(offset int32) WatermarkApiListRequest {
	r.offset = &offset
	return r
}
func (r WatermarkApiListRequest) Limit(limit int32) WatermarkApiListRequest {
	r.limit = &limit
	return r
}

type WatermarkServiceI interface {
	/*
	 * Upload Create a new watermark
	 * @return WatermarkApiUploadRequest
	 */
	Upload(fileName string, fileReader io.Reader) (*CreateWatermarkResponse, error)
	/*
	 * Upload Create a new watermark
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return WatermarkApiUploadRequest
	 */
	UploadWithContext(ctx context.Context, fileName string, fileReader io.Reader) (*CreateWatermarkResponse, error)

	/*
	 * Delete Delete a watermark by ID
	 * @param id Watermark ID
	 * @return WatermarkApiDeleteRequest
	 */

	Delete(id string) (*ResponseSuccess, error)

	/*
	 * Delete Delete a watermark by ID
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id Watermark ID
	 * @return WatermarkApiDeleteRequest
	 */

	DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error)

	/*
	 * List List all watermarks
	 * @return WatermarkApiListRequest
	 */

	List(r WatermarkApiListRequest) (*GetAllWatermarkResponse, error)

	/*
	 * List List all watermarks
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return WatermarkApiListRequest
	 */

	ListWithContext(ctx context.Context, r WatermarkApiListRequest) (*GetAllWatermarkResponse, error)
}

// WatermarkService communicating with the Watermark
// endpoints of the w3stream API
type WatermarkService struct {
	client *Client
}

/*
 * Upload Create a new watermark
 * Create a new watermark by uploading a JPG or a PNG image.

 * @return WatermarkApiUploadRequest
 */

func (s *WatermarkService) UploadFile(file *os.File) (*CreateWatermarkResponse, error) {
	return s.UploadFileWithContext(context.Background(), file)
}

/*
 * Upload Create a new watermark
 * Create a new watermark by uploading a JPG or a PNG image.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return WatermarkApiUploadRequest
 */

func (s *WatermarkService) UploadFileWithContext(ctx context.Context, file *os.File) (*CreateWatermarkResponse, error) {
	return s.UploadWithContext(ctx, file.Name(), io.Reader(file))
}

/*
* Upload Create a new watermark
* Create a new watermark by uploading a JPG or a PNG image.

* @return WatermarkApiUploadRequest
 */
func (s *WatermarkService) Upload(fileName string, fileReader io.Reader) (*CreateWatermarkResponse, error) {
	return s.UploadWithContext(context.Background(), fileName, fileReader)
}

/*
 * Upload Create a new watermark
 * Create a new watermark by uploading a JPG or a PNG image.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return WatermarkApiUploadRequest
 */
func (s *WatermarkService) UploadWithContext(ctx context.Context, fileName string, fileReader io.Reader) (*CreateWatermarkResponse, error) {
	localVarPath := "/watermarks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)

	req, err := s.client.prepareUploadRequest(ctx, localVarPath, fileName, fileReader, localVarHeaderParams, localVarQueryParams, localVarFormParams)

	if err != nil {
		return nil, err
	}

	res := new(CreateWatermarkResponse)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}

/*
 * Delete Delete a watermark by ID
 * Delete a watermark.

 * @param id Watermark ID
 * @return WatermarkApiDeleteRequest
 */

func (s *WatermarkService) Delete(id string) (*ResponseSuccess, error) {

	return s.DeleteWithContext(context.Background(), id)

}

/*
 * Delete Delete a watermark by ID
 * Delete a watermark.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Watermark ID
 * @return WatermarkApiDeleteRequest
 */

func (s *WatermarkService) DeleteWithContext(ctx context.Context, id string) (*ResponseSuccess, error) {
	var localVarPostBody interface{}

	localVarPath := "/watermarks/{id}"
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
 * List List all watermarks
 * List all watermarks associated with your workspace.

 * @return WatermarkApiListRequest
 */

func (s *WatermarkService) List(r WatermarkApiListRequest) (*GetAllWatermarkResponse, error) {

	return s.ListWithContext(context.Background(), r)

}

/*
 * List List all watermarks
 * List all watermarks associated with your workspace.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return WatermarkApiListRequest
 */

func (s *WatermarkService) ListWithContext(ctx context.Context, r WatermarkApiListRequest) (*GetAllWatermarkResponse, error) {
	var localVarPostBody interface{}

	localVarPath := "/watermarks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

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

	res := new(GetAllWatermarkResponse)
	_, err = s.client.do(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}