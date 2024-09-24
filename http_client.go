package library_http_client_go

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"github.com/harryosmar/http-client-go/ctx"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//go:generate mockgen -destination=mocks/mock_HttpClientRepository.go -package=mocks . HttpClientRepository
type (
	HttpClientRepository interface {
		EnableDebug() HttpClientRepository
		DisableDebug() HttpClientRepository
		SetLogger(loggerCtx ctx.LoggerCtx) HttpClientRepository
		Post(ctx context.Context, url string, payload *bytes.Buffer, headers map[string]string) (*Response, error)
		PostFormUrlEncoded(ctx context.Context, url string, payload url.Values, headers map[string]string) (*Response, error)
		Put(ctx context.Context, url string, payload *bytes.Buffer, headers map[string]string) (*Response, error)
		Delete(ctx context.Context, url string, headers map[string]string) (*Response, error)
		DeleteX(ctx context.Context, url string, data any, headers map[string]string) (*Response, error)
		Get(ctx context.Context, url string, queries map[string][]string, headers map[string]string) (*Response, error)
	}

	httpClientRepository struct {
		client *http.Client
		logger ctx.LoggerCtx
		debug  bool
	}

	Response struct {
		Status   int
		Content  []byte
		Header   http.Header
		Duration int64 // in millisecond
	}
)

func NewHttpClientRepository(client *http.Client) *httpClientRepository {
	return &httpClientRepository{client: client, debug: false, logger: ctx.NewLoggerCtx()}
}

func (v httpClientRepository) Post(ctx context.Context, url string, body *bytes.Buffer, headers map[string]string) (*Response, error) {
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return v.do(
		ctx,
		request,
		func() string {
			return body.String()
		},
		headers,
	)
}

func (v httpClientRepository) PostFormUrlEncoded(ctx context.Context, url string, payload url.Values, headers map[string]string) (*Response, error) {
	encodedPayload := payload.Encode()
	body := strings.NewReader(encodedPayload)
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
		headers["Content-Length"] = strconv.Itoa(len(encodedPayload))
	}

	return v.do(
		ctx,
		request,
		func() string {
			return encodedPayload
		},
		headers,
	)
}

func (v httpClientRepository) Put(ctx context.Context, url string, body *bytes.Buffer, headers map[string]string) (*Response, error) {
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	return v.do(
		ctx,
		request,
		func() string {
			return body.String()
		},
		headers,
	)
}

func (v httpClientRepository) Delete(ctx context.Context, url string, headers map[string]string) (*Response, error) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return v.do(
		ctx,
		request,
		nil,
		headers,
	)
}

func (v httpClientRepository) DeleteX(ctx context.Context, url string, data any, headers map[string]string) (*Response, error) {
	return v.Delete(ctx, url, headers)
}

func (v httpClientRepository) Get(ctx context.Context, url string, queries map[string][]string, headers map[string]string) (*Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if queries != nil {
		reqQueries := request.URL.Query()
		for queryKey, queryValue := range queries {
			for _, qv := range queryValue {
				reqQueries.Add(queryKey, qv)
			}
		}
		request.URL.RawQuery = reqQueries.Encode()
	}

	return v.do(
		ctx,
		request,
		nil,
		headers,
	)
}

func (v httpClientRepository) SetLogger(loggerCtx ctx.LoggerCtx) HttpClientRepository {
	v.logger = loggerCtx
	return v
}

func (v httpClientRepository) EnableDebug() HttpClientRepository {
	v.debug = true
	return v
}

func (v httpClientRepository) DisableDebug() HttpClientRepository {
	v.debug = false
	return v
}

const (
	XRequestIdContext = "X-Request-Id"
)

func (v httpClientRepository) do(ctx context.Context, request *http.Request, getPayload func() string, headers map[string]string) (*Response, error) {
	defer func() {
		if request != nil && request.Body != nil {
			_ = request.Body.Close()
		}
	}()

	if headers != nil && len(headers) > 0 {
		reqHeaders := make(http.Header)
		for headerKey, headerValue := range headers {
			reqHeaders[headerKey] = []string{headerValue}
		}
		request.Header = reqHeaders
	}

	requestId := ctx.Value(XRequestIdContext)
	if requestId == nil {
		requestId = uuid.New().String()
	}
	entry := v.logger.GetLoggerFromContext(ctx).WithField("x-request-id", requestId)
	ctx = v.logger.ContextWithLogger(
		ctx,
		entry,
	)

	v.logRequest(ctx, request, getPayload)

	start := time.Now().UnixNano() / int64(time.Millisecond)
	response, err := v.client.Do(request)
	defer func() {
		if response != nil && response.Body != nil {
			_ = response.Body.Close()
		}
	}()
	end := time.Now().UnixNano() / int64(time.Millisecond)
	if err != nil {
		entry.Errorf("httpClientRepository.do got err %s", err.Error())
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	v.logResponse(ctx, response, content, err)

	return &Response{
		Status:   response.StatusCode,
		Content:  content,
		Header:   response.Header,
		Duration: end - start,
	}, nil
}

func (v httpClientRepository) logRequest(ctx context.Context, req *http.Request, getPayload func() string) {
	entry := v.logger.GetLoggerFromContext(ctx)

	if req != nil {
		entry = entry.WithField("method", req.Method).
			WithField("url", req.URL.String()).
			WithField("headers", req.Header)

		if v.debug {
			entry = entry.WithField("payload", func() string {
				if getPayload == nil {
					return ""
				}
				return getPayload()
			}())
		}
	}

	entry.Infof("httpClientRepository.logRequest")
}

func (v httpClientRepository) logResponse(ctx context.Context, res *http.Response, content []byte, err error) {
	entry := v.logger.GetLoggerFromContext(ctx)
	if res != nil {
		entry = entry.WithField("status_code", func() int {
			if res == nil {
				return 0
			}
			return res.StatusCode
		}()).
			WithField("headers", res.Header)

		if v.debug {
			entry = entry.WithField("content", func() string {
				if res == nil {
					return ""
				}

				return string(content)
			}())
		}
	}

	if err != nil {
		entry.Errorf("httpClientRepository.logResponse got err %s", err.Error())
		return
	}

	entry.Infof("httpClientRepository.logResponse")
}
