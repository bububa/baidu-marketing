package core

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/bububa/baidu-marketing/core/internal/debug"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// SDKClient  object
type SDKClient struct {
	token     string
	ocpcToken string
	username  string
	password  string
	debug     bool
}

// NewSDKClient init sdk client
func NewSDKClient(token string, ocpcToken string) *SDKClient {
	return &SDKClient{
		token:     token,
		ocpcToken: ocpcToken,
	}
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

// Post data through api
func (c *SDKClient) Post(reqUrl string, bs []byte, resp interface{}) error {
	debug.PrintPostJSONRequest(reqUrl, bs, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json;charset=utf-8")
	httpResp, err := http.DefaultClient.Do(httpReq)
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
