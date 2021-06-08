package core

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/bububa/baidu-marketing/core/internal/debug"
	"github.com/bububa/baidu-marketing/model"
)

// sdkclient object
type SDKClient struct {
	token    string
	username string
	password string
	debug    bool
}

// init sdk client
func NewSDKClient(token string) *SDKClient {
	return &SDKClient{
		token: token,
	}
}

func (c SDKClient) Token() string {
	return c.token
}

func (c *SDKClient) SetUser(username string, password string) {
	c.username = username
	c.password = password
}

// set debug mode
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// execute api request
func (c *SDKClient) Do(req *model.Request, resp interface{}) error {
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
	reqBytes, _ := json.Marshal(req)
	var reqResp model.Response
	err := c.Post(req.Url(), reqBytes, &reqResp)
	if err != nil {
		return err
	}
	if reqResp.IsError() {
		return reqResp
	}
	if resp != nil {
		err = json.Unmarshal(reqResp.Body, resp)
		if err != nil {
			return err
		}
	}
	return nil
}

// post data through api
func (c *SDKClient) Post(reqUrl string, reqBytes []byte, resp interface{}) error {
	debug.PrintPostJSONRequest(reqUrl, reqBytes, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
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
