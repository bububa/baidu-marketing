package keyword

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/keyword"
)

// UpdateWord 更新关键词
func UpdateWord(clt *core.SDKClient, auth *model.RequestHeader, keywords []keyword.Keyword) (*model.ResponseHeader, []keyword.Keyword, error) {
	req := &model.Request{
		Header: auth,
		Body: &keyword.UpdateWordRequest{
			KeywordTypes: keywords,
		},
	}
	var resp keyword.UpdateWordResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
