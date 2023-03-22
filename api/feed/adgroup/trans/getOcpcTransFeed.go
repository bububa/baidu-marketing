package trans

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/feed/adgroup/trans"
)

// GetOcpcTrans 查询转化追踪
func GetOcpcTrans(clt *core.SDKClient, auth model.RequestHeader, reqBody *trans.OcpcTransRequest) (*model.ResponseHeader, []trans.OcpcTrans, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp trans.OcpcTransResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
