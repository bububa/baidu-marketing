package core

import (
	"bytes"
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
func (c *SDKClient) Do(req *model.Request, resp interface{}) (*model.ResponseHeader, error) {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return nil, err
	}
	var reqResp model.Response
	err := c.Post(req.Url(), buf.Bytes(), &reqResp)
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

func (c *SDKClient) Conversion(req model.ConversionRequest, resp interface{}) (*model.ResponseHeader, error) {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return nil, err
	}
	var reqResp model.Response
	err := c.Post(req.Url(), buf.Bytes(), &reqResp)
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

func (c *SDKClient) ActionCb(req model.ActionCbRequest) error {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return err
	}
	var reqResp model.ActionCbResponse
	if err := c.Get(req.Url(), &reqResp); err != nil {
		return err
	}
	if reqResp.IsError() {
		return reqResp
	}
	return nil
}

// OAuth execute oauth api request
func (c *SDKClient) OAuth(req model.RequestBody, resp interface{}) error {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return err
	}
	var reqResp model.DataResponse
	err := c.Post(req.Url(), buf.Bytes(), &reqResp)
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
func (c *SDKClient) Post(reqUrl string, bs []byte, resp interface{}) error {
	debug.PrintPostJSONRequest(reqUrl, bs, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json;charset=utf-8")
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	return nil
}

// Get data through api
func (c *SDKClient) Get(reqUrl string, resp interface{}) error {
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	return nil
}
