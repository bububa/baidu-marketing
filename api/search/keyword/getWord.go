package keyword

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/keyword"
)

// GetWord 查询关键词
func GetWord(clt *core.SDKClient, auth *model.RequestHeader, reqBody *keyword.GetWordRequest) (*model.ResponseHeader, []keyword.Keyword, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp keyword.GetWordResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
