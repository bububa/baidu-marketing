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
	token      string
	ocpcToken  string
	username   string
	password   string
	debug      bool
}

// NewSDKClient init sdk client
func NewSDKClient(token string, ocpcToken string) *SDKClient {
	return &SDKClient{
		token:      token,
		ocpcToken:  ocpcToken,
		httpClient: defaultHttpClient(),
	}
}

func (c *SDKClient) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

// Token get token
func (c SDKClient) Token() string {
	return c.token
}

// OcpcToken get ocpc token
func (c SDKClient) OcpcToken() string {
	return c.ocpcToken
}

// SetUser set username password
func (c *SDKClient) SetUser(username string, password string) {
	c.username = username
	c.password = password
}

// SetDebug set debug mode
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// Do execute api request
func (c *SDKClient) Do(req *model.Request, resp interface{}) (*model.ResponseHeader, error) {
	if req.Header.Token == "" {
		req.Header.Token = c.token
	}
	if req.Header.AccessToken == "" {
		if req.Header.Username == "" {
			req.Header.Username = c.username
		}
		if req.Header.Password == "" {
			req.Header.Password = c.password
		}
	}
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
	if resp != nil {
		err = json.Unmarshal(reqResp.Body, resp)
		if err != nil {
			return &reqResp.Header, err
		}
	}
	return &reqResp.Header, nil
}

func (c *SDKClient) Conversion(req model.ConversionRequest, resp interface{}) (*model.ResponseHeader, error) {
	if req.OcpcToken() == "" {
		req.SetOcpcToken(c.ocpcToken)
	}
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
	if resp != nil {
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
