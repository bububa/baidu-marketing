package core

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/bububa/baidu-marketing/core/internal/debug"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
)

func defaultHttpClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// SDKClient  object
type SDKClient struct {
	httpClient *http.Client
	tracer     *Otel
	appID      string
	secret     string
	debug      bool
}

// NewSDKClient init sdk client
func NewSDKClient(appID string, secret string) *SDKClient {
	return &SDKClient{
		httpClient: defaultHttpClient(),
	}
}

func (c *SDKClient) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *SDKClient) WithTracer(namespace string) {
	c.tracer = NewOtel(namespace, c.AppID())
}

// SetDebug set debug mode
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

func (c *SDKClient) AppID() string {
	return c.appID
}

func (c *SDKClient) Secret() string {
	return c.secret
}

// Do execute api request
func (c *SDKClient) Do(ctx context.Context, req *model.Request, resp interface{}) (*model.ResponseHeader, error) {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return nil, err
	}
	var reqResp model.Response
	err := c.Post(ctx, req.Url(), buf.Bytes(), &reqResp)
	if err != nil {
		return nil, err
	}
	if reqResp.IsError() {
		if reqResp.Header.Status == 1 && resp != nil {
			err = json.Unmarshal(reqResp.Body, resp)
			if err != nil {
				return &reqResp.Header, err
			}
		}
		return &reqResp.Header, reqResp
	}
	if resp != nil && reqResp.Body != nil {
		err = json.Unmarshal(reqResp.Body, resp)
		if err != nil {
			return &reqResp.Header, err
		}
	}
	return &reqResp.Header, nil
}

func (c *SDKClient) Conversion(ctx context.Context, req model.ConversionRequest, resp interface{}) (*model.ResponseHeader, error) {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return nil, err
	}
	var reqResp model.Response
	err := c.Post(ctx, req.Url(), buf.Bytes(), &reqResp)
	if err != nil {
		return nil, err
	}
	if reqResp.IsError() {
		if reqResp.Header.Status == 1 && resp != nil {
			err = json.Unmarshal(reqResp.Body, resp)
			if err != nil {
				return &reqResp.Header, err
			}
		}
		return &reqResp.Header, reqResp
	}
	if resp != nil && reqResp.Body != nil {
		err = json.Unmarshal(reqResp.Body, resp)
		if err != nil {
			return &reqResp.Header, err
		}
	}
	return &reqResp.Header, nil
}

func (c *SDKClient) ActionCb(ctx context.Context, req model.ActionCbRequest) error {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return err
	}
	var reqResp model.ActionCbResponse
	if err := c.Get(ctx, req.Url(), &reqResp); err != nil {
		return err
	}
	if reqResp.IsError() {
		return reqResp
	}
	return nil
}

// OAuth execute oauth api request
func (c *SDKClient) OAuth(ctx context.Context, req model.RequestBody, resp interface{}) error {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return err
	}
	var reqResp model.DataResponse
	err := c.Post(ctx, req.Url(), buf.Bytes(), &reqResp)
	if err != nil {
		return err
	}
	if reqResp.IsError() {
		return reqResp
	}
	if resp != nil && reqResp.Data != nil {
		err = json.Unmarshal(reqResp.Data, resp)
		if err != nil {
			return err
		}
	}
	return nil
}

// Post data through api
func (c *SDKClient) Post(ctx context.Context, reqUrl string, bs []byte, resp interface{}) error {
	debug.PrintPostJSONRequest(reqUrl, bs, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json;charset=utf-8")
	return c.WithSpan(ctx, httpReq, resp, bs, c.fetch)
}

// Get data through api
func (c *SDKClient) Get(ctx context.Context, reqUrl string, resp interface{}) error {
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}
	return c.WithSpan(ctx, httpReq, resp, nil, c.fetch)
}

func (c *SDKClient) fetch(httpReq *http.Request, resp interface{}) (*http.Response, error) {
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return httpResp, err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		return httpResp, nil
	}
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return httpResp, err
	}
	return httpResp, nil
}

func (c *SDKClient) WithSpan(ctx context.Context, req *http.Request, resp interface{}, payload []byte, fn func(*http.Request, interface{}) (*http.Response, error)) error {
	if c.tracer == nil {
		_, err := fn(req, resp)
		return err
	}
	return c.tracer.WithSpan(ctx, req, resp, payload, fn)
}
