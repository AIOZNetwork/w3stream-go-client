package w3streamsdk

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	BaseURL            *url.URL
	PublicKey          string
	SecretKey          string
	httpClient         Doer
	chunkSize          int64
	applicationName    string
	applicationVersion string
	sdkName            string
	sdkVersion         string
	BearerToken        string

	ApiKey       ApiKeyServiceI
	LiveStream   LiveStreamServiceI
	Players      PlayersServiceI
	Video        VideoServiceI
	VideoChapter VideoChapterServiceI
	Watermark    WatermarkServiceI
	Webhook      WebhookServiceI
}

type ErrorResponse struct {
	Response *http.Response
	Body     []byte
}

func (r *ErrorResponse) Error() string {
	var bodyStr string
	if len(r.Body) > 0 {
		var jsonBody interface{}
		if err := json.Unmarshal(r.Body, &jsonBody); err == nil {
			prettyJSON, err := json.MarshalIndent(jsonBody, "", "  ")
			if err == nil {
				bodyStr = string(prettyJSON)
			}
		}
		if bodyStr == "" {
			bodyStr = string(r.Body)
		}
	} else {
		bodyStr = "Empty body"
	}

	return fmt.Sprintf(
		"[%d]: %v %v\n %s\n",
		r.Response.StatusCode,
		r.Response.Request.Method,
		r.Response.Request.URL,
		bodyStr,
	)
}

type AuthCredentials struct {
	PublicKey   string `json:"public_key"`
	SecretKey   string `json:"secret_key"`
	BearerToken string `json:"bearer_token"`
}

const (
	defaultBaseURL   = "https://api.w3stream.xyz/api/"
	defaultChunkSize = 50 * 1024 * 1024
	minChunkSize     = 5 * 1024 * 1024
	maxChunkSize     = 128 * 1024 * 1024
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Builder struct {
	publicKey          string
	secretKey          string
	baseURL            string
	uploadChunkSize    int64
	httpClient         Doer
	applicationName    string
	applicationVersion string
	sdkName            string
	sdkVersion         string
	bearerToken        string
}

func (cb *Builder) BaseURL(url string) *Builder {
	cb.baseURL = url
	return cb
}

func (cb *Builder) SecretKey(key string) *Builder {
	cb.secretKey = key
	return cb
}

func (cb *Builder) PublicKey(secret string) *Builder {
	cb.publicKey = secret
	return cb
}

func (cb *Builder) BearerToken(token string) *Builder {
	cb.bearerToken = token
	return cb
}

func (cb *Builder) ApplicationName(applicationName string) *Builder {
	cb.applicationName = applicationName
	return cb
}

func (cb *Builder) ApplicationVersion(applicationVersion string) *Builder {
	cb.applicationVersion = applicationVersion
	return cb
}

func (cb *Builder) SdkName(sdkName string) *Builder {
	cb.sdkName = sdkName
	return cb
}

func (cb *Builder) SdkVersion(sdkVersion string) *Builder {
	cb.sdkVersion = sdkVersion
	return cb
}

func (cb *Builder) UploadChunkSize(size int64) *Builder {
	if size < minChunkSize {
		cb.uploadChunkSize = minChunkSize
	} else if size > maxChunkSize {
		cb.uploadChunkSize = maxChunkSize
	} else {
		cb.uploadChunkSize = size
	}
	return cb
}

func (cb *Builder) HTTPClient(httpClient Doer) *Builder {
	cb.httpClient = httpClient
	return cb
}

func ClientBuilder(credentials AuthCredentials) *Builder {
	builder := &Builder{
		baseURL:         defaultBaseURL,
		uploadChunkSize: defaultChunkSize,
	}
	if credentials.BearerToken != "" {
		builder.bearerToken = credentials.BearerToken
	} else {
		builder.secretKey = credentials.SecretKey
		builder.publicKey = credentials.PublicKey
	}

	return builder
}

func (cb *Builder) Build() *Client {

	baseURL, _ := url.Parse(cb.baseURL)

	var httpClient Doer

	if cb.httpClient == nil {
		httpClient = http.DefaultClient
	} else {
		httpClient = cb.httpClient
	}

	c := &Client{
		BaseURL:            baseURL,
		SecretKey:          cb.secretKey,
		PublicKey:          cb.publicKey,
		httpClient:         httpClient,
		chunkSize:          cb.uploadChunkSize,
		applicationName:    cb.applicationName,
		applicationVersion: cb.applicationVersion,
		sdkName:            cb.sdkName,
		sdkVersion:         cb.sdkVersion,
		BearerToken:        cb.bearerToken,
	}

	c.ApiKey = &ApiKeyService{client: c}
	c.LiveStream = &LiveStreamService{client: c}
	c.Players = &PlayersService{client: c}
	c.Video = &VideoService{client: c}
	c.VideoChapter = &VideoChapterService{client: c}
	c.Watermark = &WatermarkService{client: c}
	c.Webhook = &WebhookService{client: c}

	return c
}

// ChunkSize changes chunk size for video upload, by default its 50MB
func (c *Client) ChunkSize(size int64) {
	c.chunkSize = size
}

func (c *Client) UploadVideo(ctx context.Context, id string, fileName string, fileReader io.Reader, fileSize int64) error {
	localVarPath := "/videos/{id}/part"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)

	requests, err := c.prepareRangeRequests(ctx, localVarPath, fileName, fileReader, fileSize, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return err
	}

	res := new(ResponseSuccess)
	for _, req := range requests {
		_, err = c.do(req, res)
		if err != nil {
			return err
		}
	}

	localVarPath = "/videos/{id}/complete"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams = make(map[string]string)
	localVarQueryParams = url.Values{}
	var localVarPostBody interface{}

	req, err := c.prepareRequest(ctx, http.MethodGet, localVarPath, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return err
	}

	res = new(ResponseSuccess)
	_, err = c.do(req, res)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) prepareRequest(
	ctx context.Context,
	method,
	urlStr string,
	body interface{},
	headerParams map[string]string,
	queryParams url.Values) (*http.Request, error) {

	u, err := url.Parse(c.BaseURL.String())
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, urlStr)

	query := u.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}
	u.RawQuery = query.Encode()

	buf := new(bytes.Buffer)
	if body != nil {

		switch v := body.(type) {
		case *bytes.Buffer:
			buf = v
		default:
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	for headerName := range headerParams {
		req.Header.Set(headerName, headerParams[headerName])
	}

	if c.BearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.BearerToken)
	} else {
		req.SetBasicAuth(c.PublicKey, c.SecretKey)
	}
	return req, nil
}

func getOriginHeaderValue(name string, version string) (string, error) {
	if name == "" && version == "" {
		return "", nil
	}
	var re = regexp.MustCompile(`(?m)^[\w-]{1,50}$`)
	if !re.MatchString(name) {
		return "", errors.New("Invalid name value. Allowed characters: A-Z, a-z, 0-9, '-', '_'. Max length: 50.")
	}

	var reVersion = regexp.MustCompile(`(?m)^\d{1,3}(\.\d{1,3}(\.\d{1,3})?)?$`)
	if !reVersion.MatchString(version) {
		return "", errors.New("Invalid version value. The version should match the xxx[.yyy][.zzz] pattern.")
	}

	return name + ":" + version, nil
}

func (c *Client) prepareProgressiveUploadRequest(
	ctx context.Context,
	urlStr string,
	fileName string,
	fileReader io.Reader,
	fileSize int64,
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	part int32,
	isLast bool) (*http.Request, error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for key, val := range formParams {
		err := writer.WriteField(key, val)
		if err != nil {
			return nil, err
		}
	}

	partWriter, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	io.Copy(partWriter, fileReader)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := c.prepareRequest(ctx, http.MethodPost, urlStr, body, headerParams, queryParams)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	var ranges string
	if isLast {
		ranges = fmt.Sprintf("part %d/%d", part, part)
	} else {
		ranges = fmt.Sprintf("part %d/*", part)
	}

	req.Header.Set("Content-Range", ranges)

	return req, nil
}

func (c *Client) prepareRangeRequests(ctx context.Context, urlStr string, fileName string, fileReader io.Reader, fileSize int64, headerParams map[string]string, queryParams url.Values, formParams map[string]string) ([]*http.Request, error) {
	chunkSize := int64(50 * 1024 * 1024)
	var requests []*http.Request

	partsCount := int(math.Ceil(float64(fileSize) / float64(chunkSize)))
	startByte := int64(0)
	if partsCount == 1 {
		endByte := startByte + chunkSize - 1
		if endByte >= fileSize {
			endByte = fileSize - 1
		}
		bytesToRead := endByte - startByte + 1

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		for key, val := range formParams {
			err := writer.WriteField(key, val)
			if err != nil {
				return nil, err
			}
		}

		partWriter, err := writer.CreateFormFile("file", fileName)
		if err != nil {
			return nil, err
		}

		hash := md5.New()
		multiWriter := io.MultiWriter(partWriter, hash)

		limitedReader := io.LimitReader(fileReader, bytesToRead)
		_, err = io.Copy(multiWriter, limitedReader)
		if err != nil {
			return nil, err
		}

		md5Hash := hex.EncodeToString(hash.Sum(nil))
		if err := writer.WriteField("hash", md5Hash); err != nil {
			return nil, err
		}
		if err := writer.WriteField("index", "1"); err != nil {
			return nil, err
		}

		err = writer.Close()
		if err != nil {
			return nil, err
		}

		req, err := c.prepareRequest(ctx, http.MethodPost, urlStr, body, headerParams, queryParams)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())
		ranges := fmt.Sprintf("bytes %d-%d/%d", startByte, endByte, fileSize)
		req.Header.Set("Content-Range", ranges)

		requests = append(requests, req)
	} else {
		for part := 1; part <= partsCount; part++ {
			var endByte int64
			if part == partsCount {
				endByte = fileSize - 1
			} else {
				endByte = startByte + chunkSize - 1
				if endByte >= fileSize {
					endByte = fileSize - 1
				}
			}
			bytesToRead := endByte - startByte + 1

			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)

			for key, val := range formParams {
				err := writer.WriteField(key, val)
				if err != nil {
					return nil, err
				}
			}

			partWriter, err := writer.CreateFormFile("file", fileName)
			if err != nil {
				return nil, err
			}

			hash := md5.New()
			multiWriter := io.MultiWriter(partWriter, hash)

			limitedReader := io.LimitReader(fileReader, bytesToRead)
			_, err = io.Copy(multiWriter, limitedReader)
			if err != nil {
				return nil, err
			}

			md5Hash := hex.EncodeToString(hash.Sum(nil))
			if err := writer.WriteField("hash", md5Hash); err != nil {
				return nil, err
			}
			if err := writer.WriteField("index", strconv.Itoa(part)); err != nil {
				return nil, err
			}

			err = writer.Close()
			if err != nil {
				return nil, err
			}

			req, err := c.prepareRequest(ctx, http.MethodPost, urlStr, body, headerParams, queryParams)
			if err != nil {
				return nil, err
			}

			req.Header.Set("Content-Type", writer.FormDataContentType())
			ranges := fmt.Sprintf("bytes %d-%d/%d", startByte, endByte, fileSize)
			req.Header.Set("Content-Range", ranges)

			requests = append(requests, req)

			startByte = endByte + 1

			if startByte >= fileSize {
				break
			}
		}
	}

	return requests, nil
}

func (c *Client) prepareUploadRequest(
	ctx context.Context,
	urlStr string,
	fileName string,
	fileReader io.Reader,
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string) (*http.Request, error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	partWriter, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(partWriter, fileReader)
	if err != nil {
		return nil, err
	}

	for key, val := range formParams {
		err = writer.WriteField(key, val)
		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := c.prepareRequest(ctx, http.MethodPost, urlStr, body, headerParams, queryParams)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	defer r.Body.Close()

	errorResponse := &ErrorResponse{
		Response: r,
		Body:     body,
	}
	return errorResponse
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

func addDeepQueryParams(filter interface{}, prefix string, queryParams url.Values) {
	v := reflect.ValueOf(filter)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := t.Field(i)

			if field.Kind() == reflect.Ptr && !field.IsNil() {
				field = field.Elem()
			}

			lowercaseFirstChar := strings.ToLower(string(fieldType.Name[0]))
			restOfString := fieldType.Name[1:]
			lowercaseName := lowercaseFirstChar + restOfString

			if field.Kind() == reflect.Slice {
				for j := 0; j < field.Len(); j++ {
					item := field.Index(j)
					queryParams.Add(fmt.Sprintf("%s[%s][%d]", prefix, lowercaseName, j), item.String())
				}
			} else if field.Kind() == reflect.String {
				queryParams.Add(fmt.Sprintf("%s[%s]", prefix, lowercaseName), field.String())
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			queryParams.Add(fmt.Sprintf("%s[%s]", prefix, key.String()), fmt.Sprintf("%v", val))
		}
	default:
		fmt.Println("Unsupported type")
	}
}
